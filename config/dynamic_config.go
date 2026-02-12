package config

import (
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"
)

// DynamicConfig manages configuration that can be updated at runtime
type DynamicConfig struct {
	mu     sync.RWMutex
	config *Config
	db     ConfigStore
}

// ConfigStore interface for database operations
type ConfigStore interface {
	GetConfig(key string) (string, error)
	GetAllConfigs() (map[string]string, error)
	SetConfig(key, value, updatedBy string) error
}

var (
	dynamicConfig *DynamicConfig
	once          sync.Once
)

// InitDynamic initializes the dynamic configuration with database
func InitDynamic(db ConfigStore) (*DynamicConfig, error) {
	var err error
	once.Do(func() {
		dynamicConfig = &DynamicConfig{
			db: db,
		}
		err = dynamicConfig.loadFromDatabase()
	})
	return dynamicConfig, err
}

// GetDynamic returns the singleton dynamic config instance
func GetDynamic() *DynamicConfig {
	return dynamicConfig
}

// loadFromDatabase loads configuration from database
func (dc *DynamicConfig) loadFromDatabase() error {
	dc.mu.Lock()
	defer dc.mu.Unlock()

	// Get all configs from database
	configs, err := dc.db.GetAllConfigs()
	if err != nil {
		return fmt.Errorf("failed to load configs from database: %v", err)
	}

	// Create config object
	cfg := &Config{
		Environment: getEnv("ENVIRONMENT", "development"),
		Port:        getEnv("PORT", "59123"),
	}

	// Load from database with fallback to defaults
	cfg.BaseURL = getConfigValue(configs, "base_url", "https://winbu.net")
	cfg.UserAgent = getConfigValue(configs, "user_agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36")
	
	// Parse timeout
	timeoutStr := getConfigValue(configs, "timeout", "30s")
	timeout, err := time.ParseDuration(timeoutStr)
	if err != nil {
		log.Printf("Warning: Invalid timeout value '%s', using default 30s", timeoutStr)
		timeout = 30 * time.Second
	}
	cfg.Timeout = timeout

	// Parse rate limit
	rateLimitStr := getConfigValue(configs, "rate_limit", "1s")
	rateLimit, err := time.ParseDuration(rateLimitStr)
	if err != nil {
		log.Printf("Warning: Invalid rate_limit value '%s', using default 1s", rateLimitStr)
		rateLimit = 1 * time.Second
	}
	cfg.RateLimit = rateLimit

	// Parse max retries
	maxRetriesStr := getConfigValue(configs, "max_retries", "3")
	maxRetries, err := strconv.Atoi(maxRetriesStr)
	if err != nil {
		log.Printf("Warning: Invalid max_retries value '%s', using default 3", maxRetriesStr)
		maxRetries = 3
	}
	cfg.MaxRetries = maxRetries

	// Parse cache enabled
	cacheEnabledStr := getConfigValue(configs, "cache_enabled", "true")
	cfg.CacheEnabled = cacheEnabledStr == "true"

	// Parse cache TTL
	cacheTTLStr := getConfigValue(configs, "cache_ttl", "5m")
	cacheTTL, err := time.ParseDuration(cacheTTLStr)
	if err != nil {
		log.Printf("Warning: Invalid cache_ttl value '%s', using default 5m", cacheTTLStr)
		cacheTTL = 5 * time.Minute
	}
	cfg.CacheTTL = cacheTTL

	dc.config = cfg
	log.Println("✓ Configuration loaded from database")
	return nil
}

// Reload reloads configuration from database
func (dc *DynamicConfig) Reload() error {
	return dc.loadFromDatabase()
}

// Get returns a copy of the current configuration
func (dc *DynamicConfig) Get() *Config {
	dc.mu.RLock()
	defer dc.mu.RUnlock()

	// Return a copy to prevent external modifications
	cfg := *dc.config
	return &cfg
}

// UpdateBaseURL updates the base URL and reloads config
func (dc *DynamicConfig) UpdateBaseURL(url, updatedBy string) error {
	if err := dc.db.SetConfig("base_url", url, updatedBy); err != nil {
		return err
	}
	return dc.Reload()
}

// UpdateTimeout updates the timeout and reloads config
func (dc *DynamicConfig) UpdateTimeout(timeout time.Duration, updatedBy string) error {
	if err := dc.db.SetConfig("timeout", timeout.String(), updatedBy); err != nil {
		return err
	}
	return dc.Reload()
}

// UpdateRateLimit updates the rate limit and reloads config
func (dc *DynamicConfig) UpdateRateLimit(rateLimit time.Duration, updatedBy string) error {
	if err := dc.db.SetConfig("rate_limit", rateLimit.String(), updatedBy); err != nil {
		return err
	}
	return dc.Reload()
}

// UpdateMaxRetries updates the max retries and reloads config
func (dc *DynamicConfig) UpdateMaxRetries(maxRetries int, updatedBy string) error {
	if err := dc.db.SetConfig("max_retries", strconv.Itoa(maxRetries), updatedBy); err != nil {
		return err
	}
	return dc.Reload()
}

// UpdateCacheEnabled updates cache enabled flag and reloads config
func (dc *DynamicConfig) UpdateCacheEnabled(enabled bool, updatedBy string) error {
	value := "false"
	if enabled {
		value = "true"
	}
	if err := dc.db.SetConfig("cache_enabled", value, updatedBy); err != nil {
		return err
	}
	return dc.Reload()
}

// UpdateCacheTTL updates cache TTL and reloads config
func (dc *DynamicConfig) UpdateCacheTTL(ttl time.Duration, updatedBy string) error {
	if err := dc.db.SetConfig("cache_ttl", ttl.String(), updatedBy); err != nil {
		return err
	}
	return dc.Reload()
}

// UpdateConfig updates any config key and reloads
func (dc *DynamicConfig) UpdateConfig(key, value, updatedBy string) error {
	if err := dc.db.SetConfig(key, value, updatedBy); err != nil {
		return err
	}
	return dc.Reload()
}

// Helper function to get config value with fallback
func getConfigValue(configs map[string]string, key, defaultValue string) string {
	if value, ok := configs[key]; ok && value != "" {
		return value
	}
	return defaultValue
}
