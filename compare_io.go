package main

import (
	"encoding/json"
	"os"
)

func loadResultJSON(filename string) (*Result, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var r Result

	err = json.Unmarshal(data, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}
