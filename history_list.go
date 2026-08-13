package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func listHistory(target string) error {
	files, err := filepath.Glob("history/*")
	if err != nil {
		return err
	}

	if len(files) == 0 {
		fmt.Println("Belum ada history scan")
		return nil
	}

	fmt.Println("WebScan History")
	fmt.Println("--------------------------------")

	filter := strings.ReplaceAll(target, "https://", "")
	filter = strings.ReplaceAll(filter, "http://", "")

	found := false

	for _, file := range files {
		if strings.Contains(file, filter) {
			fmt.Println(filepath.Base(file))
			found = true
		}
	}

	if !found {
		fmt.Println("Tidak ada history untuk:", target)
	}

	return nil
}
