package dashboard

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nabilulilalbab/winbu.tv/config"
	"github.com/nabilulilalbab/winbu.tv/database"
)

// Handler manages dashboard and admin endpoints
type Handler struct {
	dynamicConfig *config.DynamicConfig
}

// NewHandler creates a new dashboard handler
func NewHandler(dc *config.DynamicConfig) *Handler {
	return &Handler{
		dynamicConfig: dc,
	}
}

// GetConfig returns all configuration
func (h *Handler) GetConfig(c *gin.Context) {
	configs, err := database.GetAllConfigs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": true,
			"message": "Failed to get configurations: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"message": "Success",
		"data": configs,
	})
}

// GetConfigByKey returns specific configuration
func (h *Handler) GetConfigByKey(c *gin.Context) {
	key := c.Param("key")
	if key == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"message": "Config key is required",
		})
		return
	}

	value, err := database.GetConfig(key)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": true,
			"message": "Config not found: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"message": "Success",
		"data": gin.H{
			"key": key,
			"value": value,
		},
	})
}

// UpdateConfig updates a configuration
func (h *Handler) UpdateConfig(c *gin.Context) {
	key := c.Param("key")
	if key == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"message": "Config key is required",
		})
		return
	}

	var req struct {
		Value string `json:"value" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"message": "Invalid request: " + err.Error(),
		})
		return
	}

	// Get username from context (will be set by auth middleware)
	username := c.GetString("username")
	if username == "" {
		username = "admin" // Default for now
	}

	// Update config in database and reload
	err := h.dynamicConfig.UpdateConfig(key, req.Value, username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": true,
			"message": "Failed to update config: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"message": "Configuration updated successfully",
		"data": gin.H{
			"key": key,
			"value": req.Value,
		},
	})
}

// ReloadConfig reloads configuration from database
func (h *Handler) ReloadConfig(c *gin.Context) {
	err := h.dynamicConfig.Reload()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": true,
			"message": "Failed to reload config: " + err.Error(),
		})
		return
	}

	// Get current config to return
	cfg := h.dynamicConfig.Get()

	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"message": "Configuration reloaded successfully",
		"data": gin.H{
			"base_url": cfg.BaseURL,
			"timeout": cfg.Timeout.String(),
			"rate_limit": cfg.RateLimit.String(),
			"max_retries": cfg.MaxRetries,
			"cache_enabled": cfg.CacheEnabled,
			"cache_ttl": cfg.CacheTTL.String(),
		},
	})
}

// GetCurrentConfig returns the currently active configuration
func (h *Handler) GetCurrentConfig(c *gin.Context) {
	cfg := h.dynamicConfig.Get()

	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"message": "Success",
		"data": gin.H{
			"base_url": cfg.BaseURL,
			"timeout": cfg.Timeout.String(),
			"rate_limit": cfg.RateLimit.String(),
			"max_retries": cfg.MaxRetries,
			"cache_enabled": cfg.CacheEnabled,
			"cache_ttl": cfg.CacheTTL.String(),
			"user_agent": cfg.UserAgent,
		},
	})
}

// GetMetrics returns API metrics
func (h *Handler) GetMetrics(c *gin.Context) {
	// Get limit from query param (default 100)
	limitStr := c.DefaultQuery("limit", "100")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 100
	}

	// Query metrics from database
	rows, err := database.DB.Query(`
		SELECT endpoint, method, status_code, response_time_ms, created_at
		FROM api_metrics
		ORDER BY created_at DESC
		LIMIT ?
	`, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": true,
			"message": "Failed to get metrics: " + err.Error(),
		})
		return
	}
	defer rows.Close()

	var metrics []gin.H
	for rows.Next() {
		var endpoint, method string
		var statusCode, responseTimeMs int
		var createdAt time.Time

		if err := rows.Scan(&endpoint, &method, &statusCode, &responseTimeMs, &createdAt); err != nil {
			continue
		}

		metrics = append(metrics, gin.H{
			"endpoint": endpoint,
			"method": method,
			"status_code": statusCode,
			"response_time_ms": responseTimeMs,
			"created_at": createdAt.Format("2006-01-02 15:04:05"),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"message": "Success",
		"count": len(metrics),
		"data": metrics,
	})
}

// GetMetricsSummary returns aggregated metrics
func (h *Handler) GetMetricsSummary(c *gin.Context) {
	// Total requests
	var totalRequests int
	database.DB.QueryRow("SELECT COUNT(*) FROM api_metrics").Scan(&totalRequests)

	// Average response time
	var avgResponseTime float64
	database.DB.QueryRow("SELECT AVG(response_time_ms) FROM api_metrics").Scan(&avgResponseTime)

	// Success rate (2xx status codes)
	var successCount int
	database.DB.QueryRow("SELECT COUNT(*) FROM api_metrics WHERE status_code >= 200 AND status_code < 300").Scan(&successCount)

	var successRate float64
	if totalRequests > 0 {
		successRate = (float64(successCount) / float64(totalRequests)) * 100
	}

	// Top endpoints
	rows, _ := database.DB.Query(`
		SELECT endpoint, COUNT(*) as count
		FROM api_metrics
		GROUP BY endpoint
		ORDER BY count DESC
		LIMIT 5
	`)
	defer rows.Close()

	var topEndpoints []gin.H
	for rows.Next() {
		var endpoint string
		var count int
		if err := rows.Scan(&endpoint, &count); err == nil {
			topEndpoints = append(topEndpoints, gin.H{
				"endpoint": endpoint,
				"count": count,
			})
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"message": "Success",
		"data": gin.H{
			"total_requests": totalRequests,
			"avg_response_time_ms": avgResponseTime,
			"success_rate": successRate,
			"top_endpoints": topEndpoints,
		},
	})
}

// GetHealthChecks returns health check history
func (h *Handler) GetHealthChecks(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "50")
	limit, _ := strconv.Atoi(limitStr)
	if limit < 1 {
		limit = 50
	}

	rows, err := database.DB.Query(`
		SELECT scraper_name, status, items_found, confidence_score, response_time_ms, created_at
		FROM health_checks
		ORDER BY created_at DESC
		LIMIT ?
	`, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": true,
			"message": "Failed to get health checks: " + err.Error(),
		})
		return
	}
	defer rows.Close()

	var checks []gin.H
	for rows.Next() {
		var scraperName, status string
		var itemsFound, responseTimeMs int
		var confidenceScore float64
		var createdAt time.Time

		if err := rows.Scan(&scraperName, &status, &itemsFound, &confidenceScore, &responseTimeMs, &createdAt); err != nil {
			continue
		}

		checks = append(checks, gin.H{
			"scraper_name": scraperName,
			"status": status,
			"items_found": itemsFound,
			"confidence_score": confidenceScore,
			"response_time_ms": responseTimeMs,
			"created_at": createdAt.Format("2006-01-02 15:04:05"),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"message": "Success",
		"count": len(checks),
		"data": checks,
	})
}
