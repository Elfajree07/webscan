package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Config struct {
	Profile string        `json:"profile"`
	Timeout time.Duration `json:"timeout"`
	Report  bool          `json:"report"`
	Output  string        `json:"output"`
}

func defaultConfig() Config {
	return Config{
		Profile: "full",
		Timeout: 10 * time.Second,
		Report:  false,
		Output:  "webscan-report.html",
	}
}

func loadConfig(filename string) (Config, error) {
	cfg := defaultConfig()

	data, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return cfg, nil
		}
		return cfg, err
	}

	var raw struct {
		Profile string `json:"profile"`
		Timeout string `json:"timeout"`
		Report  bool   `json:"report"`
		Output  string `json:"output"`
	}

	if err := json.Unmarshal(data, &raw); err != nil {
		return cfg, fmt.Errorf("config tidak valid: %w", err)
	}

	if raw.Profile != "" {
		cfg.Profile = raw.Profile
	}

	if raw.Timeout != "" {
		duration, err := time.ParseDuration(raw.Timeout)
		if err != nil {
			return cfg, fmt.Errorf("timeout tidak valid: %w", err)
		}
		cfg.Timeout = duration
	}

	cfg.Report = raw.Report

	if raw.Output != "" {
		cfg.Output = raw.Output
	}

	if cfg.Profile != "quick" && cfg.Profile != "full" {
		return cfg, fmt.Errorf("profile harus quick atau full")
	}

	if cfg.Timeout <= 0 {
		return cfg, fmt.Errorf("timeout harus lebih besar dari 0")
	}

	return cfg, nil
}
