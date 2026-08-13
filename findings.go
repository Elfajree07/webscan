package main

type Finding struct {
	Title    string `json:"title"`
	Severity string `json:"severity"`
	Detail   string `json:"detail"`
}

func generateFindings(r *Result) []Finding {
	findings := []Finding{}

	// HTTPS check
	if !r.HTTPS {
		findings = append(findings, Finding{
			Title:    "HTTPS tidak aktif",
			Severity: "high",
			Detail:   "Target tidak menggunakan HTTPS",
		})
	}

	// Security headers
	if r.Security.Present < r.Security.Total {
		findings = append(findings, Finding{
			Title:    "Security header belum lengkap",
			Severity: "medium",
			Detail:   "Beberapa header keamanan belum ditemukan",
		})
	}

	// TLS
	if r.TLSVersion == "TLS 1.0" ||
		r.TLSVersion == "TLS 1.1" {
		findings = append(findings, Finding{
			Title:    "TLS versi lama",
			Severity: "medium",
			Detail:   r.TLSVersion,
		})
	}

	// Cookies
	for _, c := range r.Cookies {
		if !c.Secure {
			findings = append(findings, Finding{
				Title:    "Cookie tanpa Secure flag",
				Severity: "low",
				Detail:   c.Name,
			})
		}

		if !c.HttpOnly {
			findings = append(findings, Finding{
				Title:    "Cookie tanpa HttpOnly flag",
				Severity: "low",
				Detail:   c.Name,
			})
		}
	}

	return findings
}
