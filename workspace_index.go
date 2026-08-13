package main

import (
	"encoding/json"
	"os"
	"time"
)

type WorkspaceEntry struct {
	Target string `json:"target"`
	Time   string `json:"time"`
}

type WorkspaceIndex struct {
	Targets []WorkspaceEntry `json:"targets"`
}

func updateWorkspaceIndex(target string) error {

	file := "workspace/index.json"

	var index WorkspaceIndex

	data, err := os.ReadFile(file)

	if err == nil {
		json.Unmarshal(data, &index)
	}

	index.Targets = append(index.Targets, WorkspaceEntry{
		Target: target,
		Time:   time.Now().Format(time.RFC3339),
	})

	out, err := json.MarshalIndent(
		index,
		"",
		"  ",
	)

	if err != nil {
		return err
	}

	return os.WriteFile(
		file,
		out,
		0644,
	)
}
