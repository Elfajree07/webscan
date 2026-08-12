package scanner

import (
	"regexp"
	"strings"
)

type FormInfo struct {
	Method string   `json:"method"`
	Action string   `json:"action"`
	Inputs []string `json:"inputs"`
}

func ExtractForms(body string) []FormInfo {
	forms := []FormInfo{}

	formRegex := regexp.MustCompile(`(?is)<form([^>]*)>(.*?)</form>`)

	matches := formRegex.FindAllStringSubmatch(body, -1)

	for _, form := range matches {
		attr := form[1]
		content := form[2]

		method := "GET"
		action := ""

		if m := regexp.MustCompile(`(?i)method=["']?([^"'\s>]+)`).FindStringSubmatch(attr); len(m) > 1 {
			method = strings.ToUpper(m[1])
		}

		if a := regexp.MustCompile(`(?i)action=["']?([^"'\s>]+)`).FindStringSubmatch(attr); len(a) > 1 {
			action = a[1]
		}

		inputs := []string{}

		inputRegex := regexp.MustCompile(`(?i)<input[^>]*name=["']?([^"'\s>]+)`)

		for _, input := range inputRegex.FindAllStringSubmatch(content, -1) {
			if len(input) > 1 {
				inputs = append(inputs, input[1])
			}
		}

		forms = append(forms, FormInfo{
			Method: method,
			Action: action,
			Inputs: inputs,
		})
	}

	return forms
}
