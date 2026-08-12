package scanner

import (
	"regexp"
	"strings"
)

type AssetInfo struct {
	Scripts []string `json:"scripts"`
	Styles  []string `json:"styles"`
}

func ExtractAssets(body string) AssetInfo {
	result := AssetInfo{
		Scripts: []string{},
		Styles:  []string{},
	}

	scriptRegex := regexp.MustCompile(`(?i)<script[^>]+src=["']([^"']+)`)
	styleRegex := regexp.MustCompile(`(?i)<link[^>]+href=["']([^"']+\.css[^"']*)`)

	for _, match := range scriptRegex.FindAllStringSubmatch(body, -1) {
		if len(match) > 1 {
			result.Scripts = appendUnique(result.Scripts, match[1])
		}
	}

	for _, match := range styleRegex.FindAllStringSubmatch(body, -1) {
		if len(match) > 1 {
			result.Styles = appendUnique(result.Styles, match[1])
		}
	}

	return result
}

func appendUnique(list []string, value string) []string {
	for _, v := range list {
		if v == value {
			return list
		}
	}

	return append(list, value)
}

func detectFrameworkHint(body string) []string {
	found := []string{}

	hints := map[string]string{
		"React":   "react",
		"Vue":     "vue",
		"Angular": "angular",
		"jQuery":  "jquery",
		"Next.js": "_next",
	}

	lower := strings.ToLower(body)

	for name, key := range hints {
		if strings.Contains(lower, strings.ToLower(key)) {
			found = append(found, name)
		}
	}

	return found
}
