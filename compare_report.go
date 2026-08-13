package main

type CompareSummary struct {
	HasPrevious bool
	OldScore    int
	NewScore    int
	Changes     []string
}

func compareLatest(r *Result) CompareSummary {
	result := CompareSummary{
		NewScore: r.Score.Score,
		Changes:  []string{},
	}

	history := loadHistory()

	if len(history) < 2 {
		return result
	}

	old := history[len(history)-2]

	result.HasPrevious = true
	result.OldScore = old.Score

	if r.Score.Score > old.Score {
		result.Changes = append(
			result.Changes,
			"Security score meningkat",
		)
	}

	if r.Score.Score < old.Score {
		result.Changes = append(
			result.Changes,
			"Security score menurun",
		)
	}

	if len(r.Findings) > 0 {
		result.Changes = append(
			result.Changes,
			"Jumlah findings: ",
		)
	}

	return result
}
