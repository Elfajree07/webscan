package main

import (
	"net/http"
	"time"
)

func newHTTPClient(timeout time.Duration, chain *[]string) *http.Client {
	return &http.Client{
		Timeout: timeout,

		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			*chain = append(*chain, req.URL.String())

			if len(via) >= 5 {
				return http.ErrUseLastResponse
			}

			return nil
		},
	}
}

func newRequest(target string) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodGet, target, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "WebScan/1.2")
	req.Header.Set("Accept", "text/html,application/xhtml+xml")

	return req, nil
}
