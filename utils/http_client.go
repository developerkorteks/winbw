package utils

import (
	"log"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/nabilulilalbab/winbu.tv/config"
)

// CreateCollector creates a new colly collector with standard settings
func CreateCollector(cfg *config.Config) *colly.Collector {
	// Extract domain from base URL
	domain := ExtractDomain(cfg.BaseURL)
	
	c := colly.NewCollector(
		colly.AllowedDomains(domain),
		colly.UserAgent(cfg.UserAgent),
		colly.CacheDir(""), // Disable cache to ensure fresh requests
	)

	// Set timeout
	c.SetRequestTimeout(cfg.Timeout)

	// Add rate limiting
	c.Limit(&colly.LimitRule{
		DomainGlob:  "*" + domain + "*",
		Parallelism: 1,
		Delay:       cfg.RateLimit,
	})

	// Add debug logging
	c.OnRequest(func(r *colly.Request) {
		log.Printf("[Scraper] Visiting %s with domain whitelist: %s", r.URL.String(), domain)
	})
	
	c.OnError(func(r *colly.Response, err error) {
		log.Printf("[Scraper] Error visiting %s: %v", r.Request.URL, err)
	})

	return c
}


// CreateCollectorWithRetry creates a collector with retry logic
func CreateCollectorWithRetry(cfg *config.Config) *colly.Collector {
	c := CreateCollector(cfg)

	// Add retry logic
	retryCount := 0
	c.OnError(func(r *colly.Response, err error) {
		if retryCount < cfg.MaxRetries {
			retryCount++
			time.Sleep(time.Duration(retryCount) * time.Second) // Exponential backoff
			r.Request.Retry()
		}
	})

	return c
}
