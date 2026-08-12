package main

import (
	"encoding/json"
	"os"
)

type MultiScanSummary struct {
	Total   int       `json:"total"`
	Success int       `json:"success"`
	Failed  int       `json:"failed"`
	Results []*Result `json:"results"`
}

func writeSummary(filename string, results []*Result) error {
	summary := MultiScanSummary{
		Total:   len(results),
		Results: results,
	}

	for _, r := range results {
		if r != nil && r.Status > 0 {
			summary.Success++
		} else {
			summary.Failed++
		}
	}

	data, err := json.MarshalIndent(summary, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}
