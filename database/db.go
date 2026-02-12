package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// Initialize initializes the database connection and creates tables
func Initialize(dbPath string) error {
	// Create database directory if not exists
	dir := filepath.Dir(dbPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create database directory: %v", err)
	}

	// Open database connection
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return fmt.Errorf("failed to open database: %v", err)
	}

	// Test connection
	if err := db.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %v", err)
	}

	// Set connection pool settings
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	DB = db

	// Run migrations
	if err := runMigrations(); err != nil {
		return fmt.Errorf("failed to run migrations: %v", err)
	}

	log.Println("✓ Database initialized successfully")
	return nil
}

// runMigrations executes the schema.sql file
func runMigrations() error {
	// Read schema file
	schemaPath := filepath.Join("database", "schema.sql")
	schema, err := os.ReadFile(schemaPath)
	if err != nil {
		return fmt.Errorf("failed to read schema file: %v", err)
	}

	// Execute schema
	if _, err := DB.Exec(string(schema)); err != nil {
		return fmt.Errorf("failed to execute schema: %v", err)
	}

	log.Println("✓ Database migrations completed")
	return nil
}

// Close closes the database connection
func Close() error {
	if DB != nil {
		return DB.Close()
	}
	return nil
}

// GetConfig retrieves a configuration value from database
func GetConfig(key string) (string, error) {
	var value string
	err := DB.QueryRow("SELECT value FROM config WHERE key = ?", key).Scan(&value)
	if err != nil {
		return "", err
	}
	return value, nil
}

// SetConfig sets or updates a configuration value
func SetConfig(key, value, updatedBy string) error {
	_, err := DB.Exec(`
		INSERT INTO config (key, value, updated_at, updated_by) 
		VALUES (?, ?, CURRENT_TIMESTAMP, ?)
		ON CONFLICT(key) DO UPDATE SET 
			value = excluded.value,
			updated_at = CURRENT_TIMESTAMP,
			updated_by = excluded.updated_by
	`, key, value, updatedBy)
	return err
}

// GetAllConfigs retrieves all configurations
func GetAllConfigs() (map[string]string, error) {
	rows, err := DB.Query("SELECT key, value FROM config ORDER BY category, key")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	configs := make(map[string]string)
	for rows.Next() {
		var key, value string
		if err := rows.Scan(&key, &value); err != nil {
			return nil, err
		}
		configs[key] = value
	}

	return configs, nil
}

// RecordMetric records an API metric
func RecordMetric(endpoint, method string, statusCode, responseTimeMs int, errorMsg string) error {
	_, err := DB.Exec(`
		INSERT INTO api_metrics (endpoint, method, status_code, response_time_ms, error_message)
		VALUES (?, ?, ?, ?, ?)
	`, endpoint, method, statusCode, responseTimeMs, errorMsg)
	return err
}

// RecordHealthCheck records a health check result
func RecordHealthCheck(scraperName, status string, itemsFound int, confidenceScore float64, responseTimeMs int, errorMsg, details string) error {
	_, err := DB.Exec(`
		INSERT INTO health_checks (scraper_name, status, items_found, confidence_score, response_time_ms, error_message, details)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`, scraperName, status, itemsFound, confidenceScore, responseTimeMs, errorMsg, details)
	return err
}
