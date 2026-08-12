package main

import (
	"crypto/tls"
	"net/http"
)

type HeaderResult struct {
	Name    string `json:"name"`
	Value   string `json:"value"`
	Present bool   `json:"present"`
}

type SecuritySummary struct {
	Present int            `json:"present"`
	Total   int            `json:"total"`
	Score   int            `json:"score"`
	Headers []HeaderResult `json:"headers"`
}

func analyzeSecurityHeaders(h http.Header) SecuritySummary {
	names := []string{
		"Content-Security-Policy",
		"Strict-Transport-Security",
		"X-Content-Type-Options",
		"X-Frame-Options",
		"Referrer-Policy",
		"Permissions-Policy",
	}

	result := SecuritySummary{
		Total: len(names),
	}

	for _, name := range names {
		value := h.Get(name)
		present := value != ""

		result.Headers = append(result.Headers, HeaderResult{
			Name:    name,
			Value:   value,
			Present: present,
		})

		if present {
			result.Present++
		}
	}

	if result.Total > 0 {
		result.Score = result.Present * 100 / result.Total
	}

	return result
}

func tlsVersionName(v uint16) string {
	switch v {
	case tls.VersionTLS13:
		return "TLS 1.3"
	case tls.VersionTLS12:
		return "TLS 1.2"
	case tls.VersionTLS11:
		return "TLS 1.1"
	case tls.VersionTLS10:
		return "TLS 1.0"
	default:
		return "Unknown"
	}
}
