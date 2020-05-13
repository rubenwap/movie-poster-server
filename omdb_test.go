package main

import (
	"github.com/jarcoal/httpmock"
	"testing"
	"reflect"
)

const (
	URL = "http://www.omdbapi.com"
)

func TestBuildURLwithParams(t *testing.T) {
	for _, c := range testCasesBuildParams {
		got := buildURLwithParams(c.input, URL)
		if got != c.expected {
			t.Fatalf("FAIL: %q, want %q, got %q", c.input, c.expected, got)
		}
		t.Logf("PASS: %q", c.description)
	}
}

func TestGetMovieResponse(t *testing.T) {
	httpmock.Activate()
	for _, c := range testCasesGetMovieResponse {
		httpmock.RegisterResponder("GET", c.input,
			httpmock.NewStringResponder(200, c.expected))
		got := getMovieResponse(c.input)
		if string(got) != c.expected {
			t.Fatalf("FAIL: %q, want %q, got %q", c.input, c.expected, got)
		}
		t.Logf("PASS: %q", c.description)
	}
}

func TestUnmarshalMovieIntoStruct(t *testing.T) {
	for _, c := range testCasesUnmarshalMovie {
		got := unmarshalMovieIntoStruct([]byte(c.input))
		if !reflect.DeepEqual(got, c.expected) {
			t.Fatalf("FAIL: %q, want %q, got %q", c.input, c.expected, got)
		}
		t.Logf("PASS: %q", c.description)
	}
}
