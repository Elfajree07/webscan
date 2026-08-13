package scanner

import (
	"io"
	"net/http"
	"strings"
	"time"
)

func FetchJSContent(url string) string {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	req, err := http.NewRequest(
		http.MethodGet,
		url,
		nil,
	)

	if err != nil {
		return ""
	}

	req.Header.Set(
		"User-Agent",
		"WebScan/2.1",
	)

	resp, err := client.Do(req)

	if err != nil {
		return ""
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return ""
	}

	return string(body)
}

func IsJavaScript(url string) bool {
	lower := strings.ToLower(url)

	return strings.Contains(lower, ".js")
}

func AnalyzeJSAssets(assets AssetInfo) []JSEndpoint {
	result := []JSEndpoint{}

	for _, script := range assets.Scripts {

		if !IsJavaScript(script) {
			continue
		}

		content := FetchJSContent(script)

		if content == "" {
			continue
		}

		result = append(
			result,
			ExtractJSEndpoints(content)...,
		)
	}

	return result
}
