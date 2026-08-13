package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type HistoryItem struct {
	Target string `json:"target"`
	Score  int    `json:"score"`
	Time   string `json:"scan_time"`
}

func loadHistory() []HistoryItem {
	items := []HistoryItem{}

	files, err := filepath.Glob("history/*.json")
	if err != nil {
		return items
	}

	for _, file := range files {
		data, err := os.ReadFile(file)
		if err != nil {
			continue
		}

		var r Result

		if json.Unmarshal(data, &r) != nil {
			continue
		}

		items = append(items, HistoryItem{
			Target: r.Target,
			Score:  r.Score.Score,
			Time:   r.ScanTime,
		})
	}

	return items
}
