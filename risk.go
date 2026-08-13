package main

type RiskSummary struct {
	Score int    `json:"score"`
	Level string `json:"level"`
}

func calculateRisk(findings []Finding) RiskSummary {
	score := 0

	for _, f := range findings {
		switch f.Severity {
		case "critical":
			score += 25
		case "high":
			score += 15
		case "medium":
			score += 8
		case "low":
			score += 3
		}
	}

	level := "low"

	switch {
	case score >= 50:
		level = "high"
	case score >= 25:
		level = "medium"
	default:
		level = "low"
	}

	return RiskSummary{
		Score: score,
		Level: level,
	}
}
