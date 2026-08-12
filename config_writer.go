package main

import (
	"encoding/json"
	"os"
)

func createDefaultConfig(filename string) error {
	cfg := Config{
		Profile: "full",
		Timeout: 10 * 1000000000,
		Report:  false,
		Output:  "webscan-report.html",
	}

	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}
