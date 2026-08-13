package main

type Evidence struct {
	Source string `json:"source"`
	Value  string `json:"value"`
}

type Intelligence struct {
	Technology string     `json:"technology"`
	Confidence int        `json:"confidence"`
	Evidence   []Evidence `json:"evidence"`
}

func buildIntelligence(r *Result) []Intelligence {
	result := []Intelligence{}

	// Cloudflare detection
	for _, tech := range r.Technologies {
		result = append(result, Intelligence{
			Technology: tech.Name,
			Confidence: 80,
			Evidence: []Evidence{
				{
					Source: tech.Source,
					Value:  tech.Name,
				},
			},
		})
	}

	// TLS evidence
	if r.HTTPS && r.TLSVersion != "" {
		result = append(result, Intelligence{
			Technology: "HTTPS/TLS",
			Confidence: 100,
			Evidence: []Evidence{
				{
					Source: "TLS",
					Value:  r.TLSVersion,
				},
			},
		})
	}

	return result
}
