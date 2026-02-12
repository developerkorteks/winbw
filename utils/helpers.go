package utils

import (
	"regexp"
	"strings"
)

// ExtractDomain extracts domain from URL (e.g., "https://winbu.net" -> "winbu.net")
func ExtractDomain(baseURL string) string {
	domain := baseURL
	if len(domain) > 8 && domain[:8] == "https://" {
		domain = domain[8:]
	} else if len(domain) > 7 && domain[:7] == "http://" {
		domain = domain[7:]
	}
	if len(domain) > 0 && domain[len(domain)-1] == '/' {
		domain = domain[:len(domain)-1]
	}
	for i := 0; i < len(domain); i++ {
		if domain[i] == '/' {
			domain = domain[:i]
			break
		}
	}
	return domain
}

// CleanText removes extra whitespace and trims text
func CleanText(text string) string {
	re := regexp.MustCompile(`\s+`)
	return strings.TrimSpace(re.ReplaceAllString(text, " "))
}

// ExtractSlugFromURL extracts anime slug from URL
func ExtractSlugFromURL(url string) string {
	// Extract slug from URL like https://winbu.tv/anime/one-piece/
	parts := strings.Split(strings.TrimSuffix(url, "/"), "/")
	if len(parts) >= 2 {
		return parts[len(parts)-1]
	}
	return ""
}

// CalculateConfidenceScore calculates confidence score based on data completeness
func CalculateConfidenceScore(totalFields, filledFields int) float64 {
	if totalFields == 0 {
		return 0.0
	}
	score := float64(filledFields) / float64(totalFields)
	if score > 1.0 {
		return 1.0
	}
	return score
}

// IsValidURL checks if URL is valid
func IsValidURL(url string) bool {
	return strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://")
}
