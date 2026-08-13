package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func saveHistory(r *Result) error {
	dir := "history"

	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	host := strings.ReplaceAll(r.FinalURL, "https://", "")
	host = strings.ReplaceAll(host, "http://", "")
	host = strings.ReplaceAll(host, "/", "_")

	filename := fmt.Sprintf(
		"%s_%s.json",
		host,
		time.Now().Format("20060102_150405"),
	)

	path := filepath.Join(dir, filename)

	data, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}
