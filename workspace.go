package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
)

func sanitizeTarget(target string) string {
	target = strings.TrimPrefix(target, "https://")
	target = strings.TrimPrefix(target, "http://")

	target = strings.Split(target, "/")[0]

	return target
}

func saveWorkspace(target string, name string, data interface{}) error {
	dir := filepath.Join(
		"workspace",
		sanitizeTarget(target),
	)

	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return err
	}

	file, err := os.Create(
		filepath.Join(dir, name),
	)

	if err != nil {
		return err
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	return encoder.Encode(data)
}

func createNotes(target string) error {
	dir := filepath.Join(
		"workspace",
		sanitizeTarget(target),
	)

	content := `# Recon Notes

Target:

Date:

Observations:

- 

Findings:

- 

`

	return os.WriteFile(
		filepath.Join(dir, "notes.md"),
		[]byte(content),
		0644,
	)
}
