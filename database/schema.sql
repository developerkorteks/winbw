-- Database schema for WinbuNET API Dashboard
-- Created: 2026-02-12

-- Configuration table - stores dynamic configuration
CREATE TABLE IF NOT EXISTS config (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    key VARCHAR(100) UNIQUE NOT NULL,
    value TEXT NOT NULL,
    description TEXT,
    category VARCHAR(50) DEFAULT 'general',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_by VARCHAR(100) DEFAULT 'system'
);

-- Create index for faster lookups
CREATE INDEX IF NOT EXISTS idx_config_key ON config(key);
CREATE INDEX IF NOT EXISTS idx_config_category ON config(category);

-- API metrics table - stores request/response metrics
CREATE TABLE IF NOT EXISTS api_metrics (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    endpoint VARCHAR(255) NOT NULL,
    method VARCHAR(10) NOT NULL,
    status_code INTEGER NOT NULL,
    response_time_ms INTEGER NOT NULL,
    error_message TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Indexes for metrics
CREATE INDEX IF NOT EXISTS idx_metrics_endpoint ON api_metrics(endpoint, created_at);
CREATE INDEX IF NOT EXISTS idx_metrics_created ON api_metrics(created_at);
CREATE INDEX IF NOT EXISTS idx_metrics_status ON api_metrics(status_code);

-- Health checks table - stores scraper health check results
CREATE TABLE IF NOT EXISTS health_checks (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    scraper_name VARCHAR(100) NOT NULL,
    status VARCHAR(20) NOT NULL, -- success, warning, error
    items_found INTEGER DEFAULT 0,
    confidence_score REAL DEFAULT 0.0,
    response_time_ms INTEGER DEFAULT 0,
    error_message TEXT,
    details TEXT, -- JSON string for additional details
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Indexes for health checks
CREATE INDEX IF NOT EXISTS idx_health_scraper ON health_checks(scraper_name, created_at);
CREATE INDEX IF NOT EXISTS idx_health_status ON health_checks(status);

-- Users table - for dashboard authentication
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username VARCHAR(50) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    role VARCHAR(20) DEFAULT 'admin',
    is_active BOOLEAN DEFAULT 1,
    last_login TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Insert default configurations
INSERT OR IGNORE INTO config (key, value, description, category) VALUES 
    ('base_url', 'https://winbu.net', 'Target website base URL', 'scraping'),
    ('timeout', '30s', 'HTTP request timeout', 'scraping'),
    ('rate_limit', '1s', 'Delay between requests', 'scraping'),
    ('max_retries', '3', 'Maximum retry attempts', 'scraping'),
    ('cache_enabled', 'true', 'Enable/disable cache', 'cache'),
    ('cache_ttl', '5m', 'Cache time-to-live', 'cache'),
    ('user_agent', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36', 'HTTP User-Agent header', 'scraping');

-- Insert default admin user (password: admin123 - HARUS DIUBAH!)
-- Password hash for 'admin123' using bcrypt
INSERT OR IGNORE INTO users (username, password_hash, role) VALUES 
    ('admin', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', 'admin');
