package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	header := http.Header{}
	header.Set("Authorization", "ApiKey wrong-key")

	gotKey, gotErr := GetAPIKey(header)

	if gotErr != nil {
		t.Fatalf("got unknown error: %v", gotErr)
	}

	if gotKey != "my-api-key" {
		t.Errorf("expected: %q, got: %q", "my-api-key", gotKey)
	}
}

func TestGetAPIKeyMissingHeader(t *testing.T) {
	header := http.Header{}

	gotKey, gotErr := GetAPIKey(header)

	if !errors.Is(gotErr, ErrNoAuthHeaderIncluded) {
		t.Errorf("err: %v, expected: %v", gotErr, ErrNoAuthHeaderIncluded)
	}

	if gotKey != "" {
		t.Errorf("expected: \"empty string\", got: %q", gotKey)
	}
}
