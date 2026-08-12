package main

import (
	"crypto/tls"
	"net/url"
	"strings"
	"time"
)

type TLSInfo struct {
	Version  string   `json:"version"`
	Subject  string   `json:"subject,omitempty"`
	Issuer   string   `json:"issuer,omitempty"`
	Expires  string   `json:"expires,omitempty"`
	DaysLeft int      `json:"days_left,omitempty"`
	DNSNames []string `json:"dns_names,omitempty"`
}

func getTLSInfo(respTLS *tls.ConnectionState) *TLSInfo {
	if respTLS == nil || len(respTLS.PeerCertificates) == 0 {
		return nil
	}

	cert := respTLS.PeerCertificates[0]
	daysLeft := int(time.Until(cert.NotAfter).Hours() / 24)

	return &TLSInfo{
		Version:  tlsVersionName(respTLS.Version),
		Subject:  cert.Subject.CommonName,
		Issuer:   cert.Issuer.CommonName,
		Expires:  cert.NotAfter.Format(time.RFC3339),
		DaysLeft: daysLeft,
		DNSNames: cert.DNSNames,
	}
}

func normalizeURL(raw string) string {
	if !strings.HasPrefix(raw, "http://") &&
		!strings.HasPrefix(raw, "https://") {
		return "https://" + raw
	}

	return raw
}

func redirectChain(target string, timeout time.Duration) ([]string, error) {
	target = normalizeURL(target)

	parsed, err := url.Parse(target)
	if err != nil {
		return nil, err
	}

	_ = parsed

	var chain []string

	client := newHTTPClient(timeout, &chain)

	req, err := newRequest(parsed.String())
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return chain, err
	}
	defer resp.Body.Close()

	if len(chain) == 0 {
		chain = append(chain, resp.Request.URL.String())
	}

	return chain, nil
}
