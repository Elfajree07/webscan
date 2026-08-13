package scanner

import (
	"regexp"
	"strings"
)

type JSDependency struct {
	Name    string `json:"name"`
	Version string `json:"version,omitempty"`
	Source  string `json:"source"`
}

func DetectJSDependencies(body string) []JSDependency {
	result := []JSDependency{}

	checks := []struct {
		name    string
		pattern string
	}{
		{
			name:    "jQuery",
			pattern: `jQuery v([0-9.]+)`,
		},
		{
			name:    "React",
			pattern: `React v([0-9.]+)`,
		},
		{
			name:    "Vue",
			pattern: `Vue\.js v([0-9.]+)`,
		},
	}

	for _, item := range checks {

		re := regexp.MustCompile(
			item.pattern,
		)

		match := re.FindStringSubmatch(body)

		if len(match) > 1 {
			result = append(result, JSDependency{
				Name:    item.name,
				Version: match[1],
				Source:  "javascript content",
			})

			continue
		}

		if strings.Contains(
			strings.ToLower(body),
			strings.ToLower(item.name),
		) {
			result = append(result, JSDependency{
				Name:   item.name,
				Source: "javascript content",
			})
		}
	}

	return result
}
