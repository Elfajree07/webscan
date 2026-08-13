package main

type ReconSummary struct {
	Target       string   `json:"target"`
	Technologies []string `json:"technologies"`
	Endpoints    []string `json:"endpoints"`
	Assets       []string `json:"assets"`
	Findings     []string `json:"findings"`
}

func buildRecon(r *Result) ReconSummary {

	out := ReconSummary{
		Target: r.Target,
	}

	for _, t := range r.Technologies {
		out.Technologies = append(
			out.Technologies,
			t.Name,
		)
	}

	for _, e := range r.Endpoints {
		out.Endpoints = append(
			out.Endpoints,
			e.Path,
		)
	}

	return out
}
