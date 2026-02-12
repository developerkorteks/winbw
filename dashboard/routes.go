package dashboard

import (
	"github.com/gin-gonic/gin"
	"github.com/nabilulilalbab/winbu.tv/config"
)

// SetupRoutes sets up dashboard and admin routes
func SetupRoutes(r *gin.RouterGroup, dc *config.DynamicConfig) {
	handler := NewHandler(dc)

	// Admin API routes
	admin := r.Group("/admin")
	{
		// Configuration management
		admin.GET("/config", handler.GetConfig)
		admin.GET("/config/current", handler.GetCurrentConfig)
		admin.GET("/config/:key", handler.GetConfigByKey)
		admin.PUT("/config/:key", handler.UpdateConfig)
		admin.POST("/config/reload", handler.ReloadConfig)

		// Metrics
		admin.GET("/metrics", handler.GetMetrics)
		admin.GET("/metrics/summary", handler.GetMetricsSummary)

		// Health checks
		admin.GET("/health-checks", handler.GetHealthChecks)
	}
}

// SetupWebRoutes sets up web dashboard routes
func SetupWebRoutes(r *gin.Engine, dc *config.DynamicConfig) error {
	webHandler, err := NewWebHandler(dc)
	if err != nil {
		return err
	}

	// Web dashboard routes
	r.GET("/dashboard", webHandler.ShowDashboard)
	r.GET("/dashboard/config", webHandler.ShowConfig)
	r.GET("/dashboard/health", webHandler.ShowHealth)

	return nil
}
