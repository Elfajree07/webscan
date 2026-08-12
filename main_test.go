package main

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"testing"
)

func TestFormatScanError_Nil(t *testing.T) {
	got := formatScanError(nil)

	if got != "" {
		t.Fatalf("expected empty string, got %q", got)
	}
}

func TestFormatScanError_Timeout(t *testing.T) {
	got := formatScanError(context.DeadlineExceeded)

	if !strings.Contains(got, "timeout") {
		t.Fatalf("expected timeout message, got %q", got)
	}
}

func TestFormatScanError_Generic(t *testing.T) {
	err := errors.New("test error")
	got := formatScanError(err)

	if got != "test error" {
		t.Fatalf("expected %q, got %q", "test error", got)
	}
}

func TestAnalyzeSecurityHeaders_AllPresent(t *testing.T) {
	headers := http.Header{}

	for _, name := range securityHeaders {
		headers.Set(name, "test-value")
	}

	result := analyzeSecurityHeaders(headers)

	if result.Present != result.Total {
		t.Fatalf("expected all headers present: %d/%d",
			result.Present, result.Total)
	}

	if result.Score != 100 {
		t.Fatalf("expected score 100, got %d", result.Score)
	}
}

func TestAnalyzeSecurityHeaders_NonePresent(t *testing.T) {
	headers := http.Header{}

	result := analyzeSecurityHeaders(headers)

	if result.Present != 0 {
		t.Fatalf("expected 0 headers present, got %d", result.Present)
	}

	if result.Score != 0 {
		t.Fatalf("expected score 0, got %d", result.Score)
	}
}
