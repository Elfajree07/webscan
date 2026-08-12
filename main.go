package main

import (
	"crypto/tls"
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const version = "1.3.0"

type Result struct {
	Target        string            `json:"target"`
	Status        int               `json:"status"`
	StatusText    string            `json:"status_text"`
	FinalURL      string            `json:"final_url"`
	Server        string            `json:"server"`
	ContentType   string            `json:"content_type"`
	ContentLength int64             `json:"content_length"`
	ResponseTime  string            `json:"response_time"`
	HTTPS         bool              `json:"https"`
	TLSVersion    string            `json:"tls_version,omitempty"`
	DNS           []string          `json:"dns,omitempty"`
	Robots        string            `json:"robots"`
	Sitemap       string            `json:"sitemap"`
	Headers       map[string]string `json:"security_headers"`
	Security      SecuritySummary   `json:"security"`
	TLSInfo       *TLSInfo          `json:"tls_info,omitempty"`
	Redirects     []string          `json:"redirects,omitempty"`
	ScanTime      string            `json:"scan_time"`
	ScanDuration  string            `json:"scan_duration"`
	Summary       ScanSummary       `json:"summary"`
}

type ScanSummary struct {
	StatusCode      int    `json:"status_code"`
	HTTPS           bool   `json:"https"`
	TLSVersion      string `json:"tls_version,omitempty"`
	SecurityHeaders string `json:"security_headers"`
	RedirectCount   int    `json:"redirect_count"`
	DNSCount        int    `json:"dns_count"`
	Robots          string `json:"robots"`
	Sitemap         string `json:"sitemap"`
}

var securityHeaders = []string{
	"Content-Security-Policy",
	"Strict-Transport-Security",
	"X-Content-Type-Options",
	"X-Frame-Options",
	"Referrer-Policy",
	"Permissions-Policy",
}

func main() {
	jsonMode := flag.Bool("json", false, "output JSON")
	timeout := flag.Duration("timeout", 10*time.Second, "HTTP timeout")
	showVersion := flag.Bool("version", false, "show version")
	report := flag.Bool("report", false, "buat HTML report")
	output := flag.String("output", "webscan-report.html", "nama file HTML report")

	flag.Usage = func() {
		fmt.Println("WebScan v" + version)
		fmt.Println()
		fmt.Println("Usage:")
		fmt.Println("  webscan <URL>")
		fmt.Println()
		fmt.Println("Options:")
		flag.PrintDefaults()
		fmt.Println()
		fmt.Println("Examples/contoh:")
		fmt.Println("  webscan https://example.com")
		fmt.Println("  webscan --json https://example.com")
		fmt.Println("  webscan --report https://example.com")
		fmt.Println("  webscan --report --output report.html https://example.com")
	}

	flag.Parse()

	if *showVersion {
		fmt.Println("WebScan v" + version)
		return
	}

	if flag.NArg() < 1 {
		flag.Usage()
		return
	}

	target := flag.Arg(0)

	result, err := scan(target, *timeout)
	if err != nil {
		fmt.Fprintln(os.Stderr, "[!] Error:", err)
		os.Exit(1)
	}

	if *report {
		if err := writeHTMLReport(result, *output); err != nil {
			fmt.Fprintln(os.Stderr, "[!] Report error:", err)
			os.Exit(1)
		}

		fmt.Println("[+] Report dibuat:", *output)
		return
	}

	if *jsonMode {
		printJSON(result)
	} else {
		printResult(result)
	}

}

func scan(target string, timeout time.Duration) (*Result, error) {
	scanStart := time.Now()

	if !strings.HasPrefix(target, "http://") &&
		!strings.HasPrefix(target, "https://") {
		target = "https://" + target
	}

	parsed, err := url.Parse(target)
	if err != nil || parsed.Hostname() == "" {
		return nil, fmt.Errorf("URL tidak valid")
	}

	result := &Result{
		Target:  target,
		HTTPS:   parsed.Scheme == "https",
		Headers: make(map[string]string),
	}

	// DNS lookup.
	ips, err := net.LookupHost(parsed.Hostname())
	if err == nil {
		result.DNS = ips
	}

	client := &http.Client{
		Timeout: timeout,

		// Redirect tetap diikuti, tapi dibatasi.
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			if len(via) >= 5 {
				return http.ErrUseLastResponse
			}
			return nil
		},
	}

	req, err := http.NewRequest(http.MethodGet, target, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "WebScan/1.0")
	req.Header.Set("Accept", "text/html,application/xhtml+xml")

	start := time.Now()

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	result.ResponseTime = time.Since(start).Round(time.Millisecond).String()
	result.Status = resp.StatusCode
	result.StatusText = resp.Status
	result.FinalURL = resp.Request.URL.String()
	result.Server = resp.Header.Get("Server")
	result.ContentType = resp.Header.Get("Content-Type")
	result.ContentLength = resp.ContentLength
	result.Security = analyzeSecurityHeaders(resp.Header)

	for _, header := range securityHeaders {
		if value := resp.Header.Get(header); value != "" {
			result.Headers[header] = value
		} else {
			result.Headers[header] = "MISSING"
		}
	}

	if resp.TLS != nil {
		result.TLSVersion = tlsVersion(resp.TLS.Version)
		result.TLSInfo = getTLSInfo(resp.TLS)
	}

	// Passive checks terhadap resource standar.
	result.Robots = checkResource(client, resp.Request.URL, "/robots.txt")
	result.Sitemap = checkResource(client, resp.Request.URL, "/sitemap.xml")

	chain, err := redirectChain(target, timeout)
	if err == nil {
		result.Redirects = chain
	}

	result.ScanTime = scanStart.Format(time.RFC3339)
	result.ScanDuration = time.Since(scanStart).Round(time.Millisecond).String()

	result.Summary = ScanSummary{
		StatusCode:      result.Status,
		HTTPS:           result.HTTPS,
		TLSVersion:      result.TLSVersion,
		SecurityHeaders: fmt.Sprintf("%d/%d", result.Security.Present, result.Security.Total),
		RedirectCount:   max(0, len(result.Redirects)-1),
		DNSCount:        len(result.DNS),
		Robots:          result.Robots,
		Sitemap:         result.Sitemap,
	}

	return result, nil
}

func checkResource(client *http.Client, base *url.URL, path string) string {
	u := *base
	u.Path = path
	u.RawQuery = ""

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return "ERROR"
	}

	req.Header.Set("User-Agent", "WebScan/1.0")

	resp, err := client.Do(req)
	if err != nil {
		return "ERROR"
	}
	defer resp.Body.Close()

	switch {
	case resp.StatusCode >= 200 && resp.StatusCode < 300:
		return "FOUND"
	case resp.StatusCode == http.StatusNotFound:
		return "NOT FOUND"
	default:
		return resp.Status
	}
}

func tlsVersion(v uint16) string {
	switch v {
	case tls.VersionTLS10:
		return "TLS 1.0"
	case tls.VersionTLS11:
		return "TLS 1.1"
	case tls.VersionTLS12:
		return "TLS 1.2"
	case tls.VersionTLS13:
		return "TLS 1.3"
	default:
		return "UNKNOWN"
	}
}

func printResult(r *Result) {
	fmt.Println()
	fmt.Println("╔════════════════════════════════════════════╗")
	fmt.Println("║              WEBSCAN v1.0                  ║")
	fmt.Println("╚════════════════════════════════════════════╝")

	fmt.Println()
	fmt.Println("Target       :", r.Target)
	fmt.Println("Final URL    :", r.FinalURL)
	fmt.Println("Status       :", r.StatusText)
	fmt.Println("Server       :", valueOrDash(r.Server))
	fmt.Println("Content-Type :", valueOrDash(r.ContentType))
	fmt.Println("Size         :", r.ContentLength)
	fmt.Println("Response     :", r.ResponseTime)
	fmt.Println("Scan Time    :", r.ScanTime)
	fmt.Println("Scan Duration:", r.ScanDuration)
	fmt.Println("HTTPS        :", r.HTTPS)
	fmt.Println("TLS          :", valueOrDash(r.TLSVersion))

	fmt.Println()
	fmt.Println("DNS")
	fmt.Println("--------------------------------------------")

	if len(r.DNS) == 0 {
		fmt.Println("No address found")
	} else {
		for _, ip := range r.DNS {
			fmt.Println(ip)
		}
	}

	fmt.Println()
	fmt.Println("Security Header Score")
	fmt.Println("--------------------------------------------")

	fmt.Printf("Score: %d%% (%d/%d)\n",
		r.Security.Score,
		r.Security.Present,
		r.Security.Total,
	)

	for _, header := range r.Security.Headers {
		if header.Present {
			fmt.Printf("[+] %-30s PRESENT\n", header.Name)
		} else {
			fmt.Printf("[-] %-30s MISSING\n", header.Name)
		}
	}

	fmt.Println()
	fmt.Println("Standard Resources")
	fmt.Println("--------------------------------------------")
	fmt.Println("robots.txt :", r.Robots)
	fmt.Println("sitemap.xml:", r.Sitemap)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func valueOrDash(s string) string {
	if s == "" {
		return "-"
	}
	return s
}

func printJSON(r *Result) {
	out, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		fmt.Fprintln(os.Stderr, "[!] JSON error:", err)
		return
	}

	fmt.Println(string(out))
}
