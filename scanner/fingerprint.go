package scanner

type Fingerprint struct {
	Name       string   `json:"name"`
	Confidence string   `json:"confidence"`
	Sources    []string `json:"sources"`
}

func MergeTechnologies(items []TechInfo) []Fingerprint {
	result := []Fingerprint{}

	index := map[string]int{}

	for _, tech := range items {
		pos, ok := index[tech.Name]

		if !ok {
			index[tech.Name] = len(result)

			result = append(result, Fingerprint{
				Name:    tech.Name,
				Sources: []string{tech.Source},
			})

			continue
		}

		exists := false

		for _, s := range result[pos].Sources {
			if s == tech.Source {
				exists = true
				break
			}
		}

		if !exists {
			result[pos].Sources = append(
				result[pos].Sources,
				tech.Source,
			)
		}
	}

	for i := range result {
		count := len(result[i].Sources)

		switch {
		case count >= 3:
			result[i].Confidence = "high"
		case count == 2:
			result[i].Confidence = "medium"
		default:
			result[i].Confidence = "low"
		}
	}

	return result
}

func FingerprintFromJS(js []JSInfo) []TechInfo {
	result := []TechInfo{}

	for _, item := range js {
		for _, lib := range item.Libraries {
			result = append(result, TechInfo{
				Name:   lib,
				Source: "javascript library",
			})
		}
	}

	return result
}
