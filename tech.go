package main

import (
	"net/http"
	"regexp"
	"strings"
)

type TechInfo struct {
	Name   string `json:"name"`
	Source string `json:"source"`
}

type HTMLMetadata struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Generator   string `json:"generator"`
}

func detectTechnologies(h http.Header, body string) []TechInfo {
	tech := make([]TechInfo, 0)

	server := strings.ToLower(h.Get("Server"))

	if strings.Contains(server, "cloudflare") {
		tech = append(tech, TechInfo{
			Name:   "Cloudflare",
			Source: "server header",
		})
	}

	if strings.Contains(strings.ToLower(body), "wp-content") {
		tech = append(tech, TechInfo{
			Name:   "WordPress",
			Source: "html body",
		})
	}

	if powered := h.Get("X-Powered-By"); powered != "" {
		tech = append(tech, TechInfo{
			Name:   powered,
			Source: "x-powered-by header",
		})
	}

	return tech
}

func extractMetadata(html string) HTMLMetadata {
	meta := HTMLMetadata{}

	title := regexp.MustCompile(`(?i)<title>(.*?)</title>`)
	desc := regexp.MustCompile(`(?i)<meta name="description" content="(.*?)"`)

	if m := title.FindStringSubmatch(html); len(m) > 1 {
		meta.Title = m[1]
	}

	if m := desc.FindStringSubmatch(html); len(m) > 1 {
		meta.Description = m[1]
	}

	return meta
}
