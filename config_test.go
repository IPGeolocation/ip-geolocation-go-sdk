package ipgeolocation

import (
	"encoding/json"
	"testing"
	"time"
)

func TestConfigNormalizesRequestOriginAndDefaults(t *testing.T) {
	config := &Config{
		APIKey:         "test-key",
		RequestOrigin:  "https://app.example.com/",
		BaseURL:        "https://api.ipgeolocation.io/",
		ConnectTimeout: 10 * time.Second,
		ReadTimeout:    30 * time.Second,
	}

	normalized, err := normalizeConfig(config)
	if err != nil {
		t.Fatalf("normalizeConfig() error = %v", err)
	}

	if normalized.APIKey != "test-key" {
		t.Fatalf("normalized.APIKey = %q", normalized.APIKey)
	}

	if normalized.RequestOrigin != "https://app.example.com" {
		t.Fatalf("normalized.RequestOrigin = %q", normalized.RequestOrigin)
	}

	if normalized.BaseURL != "https://api.ipgeolocation.io" {
		t.Fatalf("normalized.BaseURL = %q", normalized.BaseURL)
	}
}

func TestConfigRejectsRequestOriginPath(t *testing.T) {
	config := &Config{RequestOrigin: "https://app.example.com/path"}

	err := config.Validate()
	if err == nil || err.Error() != "requestOrigin must not include a path" {
		t.Fatalf("Validate() error = %v", err)
	}
}

func TestConfigSafeMapAndJSONHideAPIKey(t *testing.T) {
	config := &Config{APIKey: "test-key"}

	safeMap := config.SafeMap()
	if safeMap["apiKey"] != "[REDACTED]" {
		t.Fatalf("SafeMap apiKey = %#v", safeMap["apiKey"])
	}

	body, err := json.Marshal(config)
	if err != nil {
		t.Fatalf("json.Marshal() error = %v", err)
	}

	if string(body) == "" {
		t.Fatal("json.Marshal() returned empty body")
	}

	if jsonContains := string(body); jsonContains == "test-key" {
		t.Fatalf("json body leaked API key: %s", jsonContains)
	}

	if string(body) == "{\"apiKey\":\"test-key\"}" {
		t.Fatalf("json body leaked apiKey field: %s", string(body))
	}
}
