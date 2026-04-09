package ipgeolocation

import (
	"fmt"
	"testing"
)

func TestLookupRequestNormalizesAndDedupesValues(t *testing.T) {
	request := &LookupRequest{
		IP:       " 8.8.8.8 ",
		Lang:     "EN",
		Include:  []string{"security", "security", "abuse"},
		Fields:   []string{"location.country_name"},
		Excludes: []string{"currency"},
		Headers:  map[string]string{"user-agent": "custom"},
	}

	normalized, err := normalizeLookupRequest(request)
	if err != nil {
		t.Fatalf("normalizeLookupRequest() error = %v", err)
	}

	if normalized.IP != "8.8.8.8" {
		t.Fatalf("normalized.IP = %q", normalized.IP)
	}

	if normalized.Lang != "en" {
		t.Fatalf("normalized.Lang = %q", normalized.Lang)
	}

	if len(normalized.Include) != 2 || normalized.Include[0] != "security" || normalized.Include[1] != "abuse" {
		t.Fatalf("normalized.Include = %#v", normalized.Include)
	}

	if normalized.Headers["User-Agent"] != "custom" {
		t.Fatalf("normalized.Headers = %#v", normalized.Headers)
	}
}

func TestLookupRequestRejectsBlankHeaderNames(t *testing.T) {
	request := &LookupRequest{
		Headers: map[string]string{"   ": "value"},
	}

	err := request.Validate()
	if err == nil || err.Error() != "headers must not contain blank names" {
		t.Fatalf("Validate() error = %v", err)
	}
}

func TestBulkLookupRequestRequiresIPs(t *testing.T) {
	request := &BulkLookupRequest{}

	err := request.Validate()
	if err == nil || err.Error() != "ips must contain at least one IP address or domain" {
		t.Fatalf("Validate() error = %v", err)
	}
}

func TestBulkLookupRequestRejectsTooManyIPs(t *testing.T) {
	ips := make([]string, 50001)
	for index := range ips {
		ips[index] = fmt.Sprintf("192.0.2.%d", index)
	}

	request := &BulkLookupRequest{IPs: ips}

	err := request.Validate()
	if err == nil || err.Error() != "ips must not contain more than 50000 entries" {
		t.Fatalf("Validate() error = %v", err)
	}
}

func TestBulkLookupRequestPreservesDuplicateIPsAndOrder(t *testing.T) {
	request := &BulkLookupRequest{
		IPs: []string{" 8.8.8.8 ", "1.1.1.1", "8.8.8.8", "example.com", "example.com"},
	}

	normalized, err := normalizeBulkLookupRequest(request)
	if err != nil {
		t.Fatalf("normalizeBulkLookupRequest() error = %v", err)
	}

	expected := []string{"8.8.8.8", "1.1.1.1", "8.8.8.8", "example.com", "example.com"}
	if len(normalized.IPs) != len(expected) {
		t.Fatalf("len(normalized.IPs) = %d", len(normalized.IPs))
	}

	for index, value := range expected {
		if normalized.IPs[index] != value {
			t.Fatalf("normalized.IPs[%d] = %q", index, normalized.IPs[index])
		}
	}
}
