package scanner

import (
	"strings"
)

type JSInfo struct {
	File      string   `json:"file"`
	Type      string   `json:"type"`
	Libraries []string `json:"libraries,omitempty"`
}

func AnalyzeJavaScript(assets AssetInfo) []JSInfo {
	result := []JSInfo{}

	for _, script := range assets.Scripts {
		info := JSInfo{
			File: script,
			Type: "javascript",
		}

		info.Libraries = DetectJSLibraries(script)

		result = append(result, info)
	}

	return result
}

func DetectJSLibraries(file string) []string {
	found := []string{}

	checks := map[string]string{
		"jQuery":    "jquery",
		"React":     "react",
		"Vue":       "vue",
		"Angular":   "angular",
		"Bootstrap": "bootstrap",
		"Lodash":    "lodash",
		"Alpine.js": "alpine",
	}

	lower := strings.ToLower(file)

	for name, pattern := range checks {
		if strings.Contains(lower, pattern) {
			found = append(found, name)
		}
	}

	return found
}
