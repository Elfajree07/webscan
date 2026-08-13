package main

import (
	"bufio"
	"net/url"
	"os"
	"strings"
)

func loadScope(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var scopes []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// skip kosong dan komentar
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		scopes = append(scopes, line)
	}

	return scopes, scanner.Err()
}

func inScope(target string, scopes []string) bool {
	u, err := url.Parse(target)
	if err == nil && u.Hostname() != "" {
		target = u.Hostname()
	}

	target = strings.ToLower(target)

	for _, scope := range scopes {
		scope = strings.ToLower(strings.TrimSpace(scope))

		if strings.HasPrefix(scope, "*.") {
			base := strings.TrimPrefix(scope, "*.")

			if strings.HasSuffix(target, base) {
				return true
			}
		}

		if target == scope {
			return true
		}
	}

	return false
}
