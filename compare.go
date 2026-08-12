package main

type CompareResult struct {
	OldScore int      `json:"old_score"`
	NewScore int      `json:"new_score"`
	Changes  []string `json:"changes"`
}

func compareResults(old *Result, new *Result) CompareResult {
	result := CompareResult{
		OldScore: old.Score.Score,
		NewScore: new.Score.Score,
		Changes:  []string{},
	}

	if new.Score.Score > old.Score.Score {
		result.Changes = append(
			result.Changes,
			"Security score improved",
		)
	}

	if new.Score.Score < old.Score.Score {
		result.Changes = append(
			result.Changes,
			"Security score decreased",
		)
	}

	if old.TLSVersion != new.TLSVersion {
		result.Changes = append(
			result.Changes,
			"TLS version changed",
		)
	}

	return result
}
