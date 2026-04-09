package ipgeolocation

import (
	"context"
	"strings"
	"testing"
)

func TestLiveSingleLookup(t *testing.T) {
	freeKey, _ := requireLiveTests(t)

	client, err := NewClient(&Config{APIKey: freeKey})
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}
	defer client.Close()

	response, err := client.LookupIPGeolocation(context.Background(), &LookupRequest{IP: "8.8.8.8"})
	if err != nil {
		t.Fatalf("LookupIPGeolocation() error = %v", err)
	}

	if response.Data == nil || response.Data.IP == nil || *response.Data.IP == "" {
		t.Fatalf("response.Data = %#v", response.Data)
	}
}

func TestLiveSingleRawXML(t *testing.T) {
	freeKey, _ := requireLiveTests(t)

	client, err := NewClient(&Config{APIKey: freeKey})
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}
	defer client.Close()

	response, err := client.LookupIPGeolocationRaw(context.Background(), &LookupRequest{
		IP:     "8.8.8.8",
		Output: ResponseFormatXML,
	})
	if err != nil {
		t.Fatalf("LookupIPGeolocationRaw() error = %v", err)
	}

	if !strings.Contains(response.Data, "<ip>") {
		t.Fatalf("response.Data = %q", response.Data)
	}
}

func TestLiveBulkLookup(t *testing.T) {
	_, paidKey := requireLiveTests(t)

	client, err := NewClient(&Config{APIKey: paidKey})
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}
	defer client.Close()

	response, err := client.BulkLookupIPGeolocation(context.Background(), &BulkLookupRequest{
		IPs: []string{"8.8.8.8", "1.1.1.1"},
	})
	if err != nil {
		t.Fatalf("BulkLookupIPGeolocation() error = %v", err)
	}

	if len(response.Data) != 2 {
		t.Fatalf("len(response.Data) = %d", len(response.Data))
	}

	if response.Data[0].Data == nil || response.Data[0].Data.IP == nil {
		t.Fatalf("response.Data[0] = %#v", response.Data[0])
	}
}
