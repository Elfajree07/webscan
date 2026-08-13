package scanner

import (
	"regexp"
)

type JSEndpoint struct {
	Path   string `json:"path"`
	Source string `json:"source"`
}

func ExtractJSEndpoints(jsBody string) []JSEndpoint {
	result := []JSEndpoint{}
	seen := map[string]bool{}

	pattern := regexp.MustCompile(`["'](\/[a-zA-Z0-9_\-\/\.]+)["']`)

	matches := pattern.FindAllStringSubmatch(jsBody, -1)

	for _, match := range matches {
		if len(match) < 2 {
			continue
		}

		path := match[1]

		if seen[path] {
			continue
		}

		seen[path] = true

		result = append(result, JSEndpoint{
			Path:   path,
			Source: "javascript",
		})
	}

	return result
}
