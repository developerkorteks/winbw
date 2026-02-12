package database

import (
	"os"
	"testing"
)

func TestDatabaseInitialization(t *testing.T) {
	// Use temporary database for testing
	dbPath := "test_winbu.db"
	defer os.Remove(dbPath)

	// Test initialization
	err := Initialize(dbPath)
	if err != nil {
		t.Fatalf("Failed to initialize database: %v", err)
	}

	// Test database is accessible
	if DB == nil {
		t.Fatal("Database connection is nil")
	}

	// Test ping
	if err := DB.Ping(); err != nil {
		t.Fatalf("Failed to ping database: %v", err)
	}

	// Test config table exists
	var count int
	err = DB.QueryRow("SELECT COUNT(*) FROM config").Scan(&count)
	if err != nil {
		t.Fatalf("Config table not found: %v", err)
	}

	if count == 0 {
		t.Error("Config table is empty, expected default values")
	}

	t.Logf("✓ Database initialized with %d config entries", count)

	// Close database
	if err := Close(); err != nil {
		t.Fatalf("Failed to close database: %v", err)
	}
}

func TestConfigOperations(t *testing.T) {
	dbPath := "test_winbu.db"
	defer os.Remove(dbPath)

	if err := Initialize(dbPath); err != nil {
		t.Fatalf("Failed to initialize database: %v", err)
	}
	defer Close()

	// Test GetConfig
	t.Run("GetConfig", func(t *testing.T) {
		value, err := GetConfig("base_url")
		if err != nil {
			t.Fatalf("Failed to get config: %v", err)
		}

		if value != "https://winbu.net" {
			t.Errorf("Expected 'https://winbu.net', got '%s'", value)
		}

		t.Logf("✓ GetConfig works: %s", value)
	})

	// Test SetConfig
	t.Run("SetConfig", func(t *testing.T) {
		err := SetConfig("base_url", "https://new-domain.com", "test")
		if err != nil {
			t.Fatalf("Failed to set config: %v", err)
		}

		// Verify update
		value, err := GetConfig("base_url")
		if err != nil {
			t.Fatalf("Failed to get updated config: %v", err)
		}

		if value != "https://new-domain.com" {
			t.Errorf("Expected 'https://new-domain.com', got '%s'", value)
		}

		t.Logf("✓ SetConfig works: %s", value)
	})

	// Test GetAllConfigs
	t.Run("GetAllConfigs", func(t *testing.T) {
		configs, err := GetAllConfigs()
		if err != nil {
			t.Fatalf("Failed to get all configs: %v", err)
		}

		if len(configs) == 0 {
			t.Error("Expected configs, got empty map")
		}

		t.Logf("✓ GetAllConfigs works: %d configs found", len(configs))
		for k, v := range configs {
			t.Logf("  - %s: %s", k, v)
		}
	})
}

func TestMetricsRecording(t *testing.T) {
	dbPath := "test_winbu.db"
	defer os.Remove(dbPath)

	if err := Initialize(dbPath); err != nil {
		t.Fatalf("Failed to initialize database: %v", err)
	}
	defer Close()

	// Test RecordMetric
	err := RecordMetric("/api/v1/home", "GET", 200, 150, "")
	if err != nil {
		t.Fatalf("Failed to record metric: %v", err)
	}

	// Verify metric was recorded
	var count int
	err = DB.QueryRow("SELECT COUNT(*) FROM api_metrics").Scan(&count)
	if err != nil {
		t.Fatalf("Failed to query metrics: %v", err)
	}

	if count == 0 {
		t.Error("Expected metric to be recorded")
	}

	t.Logf("✓ Metrics recording works: %d metrics recorded", count)
}

func TestHealthCheckRecording(t *testing.T) {
	dbPath := "test_winbu.db"
	defer os.Remove(dbPath)

	if err := Initialize(dbPath); err != nil {
		t.Fatalf("Failed to initialize database: %v", err)
	}
	defer Close()

	// Test RecordHealthCheck
	err := RecordHealthCheck("anime_scraper", "success", 20, 1.0, 250, "", "{}")
	if err != nil {
		t.Fatalf("Failed to record health check: %v", err)
	}

	// Verify health check was recorded
	var count int
	err = DB.QueryRow("SELECT COUNT(*) FROM health_checks").Scan(&count)
	if err != nil {
		t.Fatalf("Failed to query health checks: %v", err)
	}

	if count == 0 {
		t.Error("Expected health check to be recorded")
	}

	t.Logf("✓ Health check recording works: %d checks recorded", count)
}
