package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func loadTargets(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var targets []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		target := strings.TrimSpace(scanner.Text())

		if target == "" || strings.HasPrefix(target, "#") {
			continue
		}

		targets = append(targets, target)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return targets, nil
}

func printMultiSummary(results []*Result) {
	fmt.Println()
	fmt.Println("Multi Target Summary")
	fmt.Println("--------------------------------------------")

	for i, r := range results {
		fmt.Printf(
			"[%d/%d] %-35s %d %s\n",
			i+1,
			len(results),
			r.Target,
			r.Status,
			r.StatusText,
		)
	}
}
