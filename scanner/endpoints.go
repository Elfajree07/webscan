package scanner

import (
	"net/url"
	"regexp"
)

type EndpointInfo struct {
	Path   string `json:"path"`
	Source string `json:"source"`
}

func ExtractEndpoints(base string, body string) []EndpointInfo {
	result := []EndpointInfo{}
	seen := map[string]bool{}

	regex := regexp.MustCompile(`(?i)(href|src)=["']([^"']+)`)

	for _, match := range regex.FindAllStringSubmatch(body, -1) {
		if len(match) < 3 {
			continue
		}

		value := match[2]

		u, err := url.Parse(value)
		if err != nil {
			continue
		}

		path := u.Path

		if path == "" || seen[path] {
			continue
		}

		seen[path] = true

		result = append(result, EndpointInfo{
			Path:   path,
			Source: "html",
		})
	}

	return result
}
