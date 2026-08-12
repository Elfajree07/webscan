package scanner

import (
	"net/url"
	"regexp"
)

type ResourceInfo struct {
	Links   []string `json:"links"`
	Scripts []string `json:"scripts"`
	Styles  []string `json:"styles"`
}

func ExtractResources(baseURL string, body string) ResourceInfo {
	result := ResourceInfo{
		Links:   []string{},
		Scripts: []string{},
		Styles:  []string{},
	}

	linkRegex := regexp.MustCompile(`href="([^"]+)"`)
	scriptRegex := regexp.MustCompile(`<script[^>]+src="([^"]+)"`)
	styleRegex := regexp.MustCompile(`<link[^>]+href="([^"]+\.css[^"]*)"`)

	for _, m := range linkRegex.FindAllStringSubmatch(body, -1) {
		result.Links = append(result.Links, normalizeURL(baseURL, m[1]))
	}

	for _, m := range scriptRegex.FindAllStringSubmatch(body, -1) {
		result.Scripts = append(result.Scripts, normalizeURL(baseURL, m[1]))
	}

	for _, m := range styleRegex.FindAllStringSubmatch(body, -1) {
		result.Styles = append(result.Styles, normalizeURL(baseURL, m[1]))
	}

	return result
}

func normalizeURL(base, target string) string {
	u, err := url.Parse(target)
	if err != nil {
		return target
	}

	b, err := url.Parse(base)
	if err != nil {
		return target
	}

	return b.ResolveReference(u).String()
}
