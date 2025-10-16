package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	testHeader := http.Header{"Authorization": []string{"ApiKey secretkey"}}
	_, err := GetAPIKey(testHeader)
	if err != nil {
		t.Fatalf("expected err=nil, got %s", err)
	}
}

func TestGetAPIKeyNil(t *testing.T) {
	testHeader := http.Header{"Authorization": []string{""}}
	_, err := GetAPIKey(testHeader)
	if err != nil {
		return
	}
	t.Fatalf("expected err=nil, got %s", err)
}
