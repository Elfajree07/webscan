package scanner

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

func DetectTechnologies(h http.Header, body string) []TechInfo {
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

func ExtractMetadata(html string) HTMLMetadata {
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

func DetectFromAssets(assets AssetInfo) []TechInfo {
	result := []TechInfo{}

	checks := map[string][]string{
		"jQuery": {
			"jquery",
		},
		"React": {
			"react",
		},
		"Vue": {
			"vue",
		},
		"Angular": {
			"angular",
		},
		"Next.js": {
			"_next",
		},
		"Bootstrap": {
			"bootstrap",
		},
	}

	for name, patterns := range checks {
		for _, script := range assets.Scripts {
			lower := strings.ToLower(script)

			for _, pattern := range patterns {
				if strings.Contains(lower, pattern) {
					result = append(result, TechInfo{
						Name:   name,
						Source: "javascript asset",
					})
					break
				}
			}
		}
	}

	return result
}

func DetectFromHTML(body string) []TechInfo {
	result := []TechInfo{}

	patterns := map[string][]string{
		"WordPress": {
			"wp-content",
			"wp-includes",
		},
		"Next.js": {
			"__next_data__",
			"/_next/",
		},
		"React": {
			"data-reactroot",
			"react",
		},
		"Angular": {
			"ng-version",
		},
		"Bootstrap": {
			"bootstrap",
		},
		"jQuery": {
			"jquery",
		},
		"Shopify": {
			"cdn.shopify.com",
		},
	}

	lower := strings.ToLower(body)

	for name, keys := range patterns {
		for _, key := range keys {
			if strings.Contains(lower, strings.ToLower(key)) {
				result = append(result, TechInfo{
					Name:   name,
					Source: "html pattern",
				})
				break
			}
		}
	}

	return result
}
