package main

import (
	"testing"
)

const (
	URL = "http://www.omdbapi.com"
)

func TestBuildParams(t *testing.T) {
	for _, c := range testCasesBuildParams {
		got := buildURLwithParams(c.input, URL)
		if got != c.expected {
			t.Fatalf("FAIL: %q, want %q, got %q", c.input, c.expected, got)
		}
		t.Logf("PASS: %q", c.description)
	}
}
