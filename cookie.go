package main

import "net/http"

func sameSiteName(s http.SameSite) string {
	switch s {
	case http.SameSiteStrictMode:
		return "Strict"
	case http.SameSiteLaxMode:
		return "Lax"
	case http.SameSiteNoneMode:
		return "None"
	default:
		return "Default"
	}
}

func analyzeCookies(cookies []*http.Cookie) []CookieInfo {
	result := make([]CookieInfo, 0)

	for _, c := range cookies {
		result = append(result, CookieInfo{
			Name:     c.Name,
			Secure:   c.Secure,
			HttpOnly: c.HttpOnly,
			SameSite: sameSiteName(c.SameSite),
		})
	}

	return result
}
