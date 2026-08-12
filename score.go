package main

func calculateScore(r *Result) SecurityScore {
	score := SecurityScore{
		HTTPS:   r.HTTPS,
		TLS:     r.TLSVersion != "",
		Robots:  r.Robots == "FOUND",
		Sitemap: r.Sitemap == "FOUND",
	}

	if score.HTTPS {
		score.Score += 25
	}

	if score.TLS {
		score.Score += 25
	}

	for _, h := range r.Headers {
		if h != "MISSING" {
			score.Headers++
		}
	}

	score.TotalHeaders = len(r.Headers)

	if score.TotalHeaders > 0 {
		score.Score += (score.Headers * 20) / score.TotalHeaders
	}

	if score.Robots {
		score.Score += 15
	}

	if score.Sitemap {
		score.Score += 15
	}

	if score.Score > 100 {
		score.Score = 100
	}

	return score
}
