package ipgeolocation

import (
	"context"
	"encoding/json"
	"testing"
)

func TestLiveSingleTypedAndRawShareKeyFields(t *testing.T) {
	paidKey := requireLiveHardening(t)

	client, err := NewClient(&Config{APIKey: paidKey})
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}
	defer client.Close()

	typedResponse, err := client.LookupIPGeolocation(context.Background(), &LookupRequest{
		IP:      "8.8.8.8",
		Include: []string{"security", "abuse", "user_agent", "hostname"},
	})
	if err != nil {
		t.Fatalf("LookupIPGeolocation() error = %v", err)
	}

	rawResponse, err := client.LookupIPGeolocationRaw(context.Background(), &LookupRequest{
		IP:      "8.8.8.8",
		Include: []string{"security", "abuse", "user_agent", "hostname"},
	})
	if err != nil {
		t.Fatalf("LookupIPGeolocationRaw() error = %v", err)
	}

	var raw map[string]interface{}
	if err := json.Unmarshal([]byte(rawResponse.Data), &raw); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}

	if typedResponse.Data == nil || typedResponse.Data.IP == nil || *typedResponse.Data.IP != raw["ip"].(string) {
		t.Fatalf("typed IP = %#v raw = %#v", typedResponse.Data, raw["ip"])
	}

	rawASN := raw["asn"].(map[string]interface{})
	if typedResponse.Data.ASN == nil || typedResponse.Data.ASN.ASNumber == nil || *typedResponse.Data.ASN.ASNumber != rawASN["as_number"].(string) {
		t.Fatalf("typed ASN = %#v raw ASN = %#v", typedResponse.Data.ASN, rawASN)
	}

	rawCompany := raw["company"].(map[string]interface{})
	if typedResponse.Data.Company == nil || typedResponse.Data.Company.Name == nil || *typedResponse.Data.Company.Name != rawCompany["name"].(string) {
		t.Fatalf("typed Company = %#v raw Company = %#v", typedResponse.Data.Company, rawCompany)
	}

	rawLocation := requireRawObject(t, raw, "location")
	if typedResponse.Data.Location == nil || typedResponse.Data.Location.CountryName == nil || *typedResponse.Data.Location.CountryName != requireRawString(t, rawLocation, "country_name") {
		t.Fatalf("typed Location = %#v raw Location = %#v", typedResponse.Data.Location, rawLocation)
	}

	rawNetwork := requireRawObject(t, raw, "network")
	if typedResponse.Data.Network == nil || typedResponse.Data.Network.Route == nil || *typedResponse.Data.Network.Route != requireRawString(t, rawNetwork, "route") {
		t.Fatalf("typed Network = %#v raw Network = %#v", typedResponse.Data.Network, rawNetwork)
	}

	rawTimeZone := requireRawObject(t, raw, "time_zone")
	if typedResponse.Data.TimeZone == nil || typedResponse.Data.TimeZone.CurrentTimeZoneAbbreviation == nil || *typedResponse.Data.TimeZone.CurrentTimeZoneAbbreviation != requireRawString(t, rawTimeZone, "current_tz_abbreviation") {
		t.Fatalf("typed TimeZone = %#v raw TimeZone = %#v", typedResponse.Data.TimeZone, rawTimeZone)
	}

	if typedResponse.Data.TimeZone.CurrentTimeZoneFullName == nil || *typedResponse.Data.TimeZone.CurrentTimeZoneFullName != requireRawString(t, rawTimeZone, "current_tz_full_name") {
		t.Fatalf("typed TimeZone = %#v raw TimeZone = %#v", typedResponse.Data.TimeZone, rawTimeZone)
	}

	if typedResponse.Data.TimeZone.StandardTimeZoneAbbreviation == nil || *typedResponse.Data.TimeZone.StandardTimeZoneAbbreviation != requireRawString(t, rawTimeZone, "standard_tz_abbreviation") {
		t.Fatalf("typed TimeZone = %#v raw TimeZone = %#v", typedResponse.Data.TimeZone, rawTimeZone)
	}

	if typedResponse.Data.TimeZone.StandardTimeZoneFullName == nil || *typedResponse.Data.TimeZone.StandardTimeZoneFullName != requireRawString(t, rawTimeZone, "standard_tz_full_name") {
		t.Fatalf("typed TimeZone = %#v raw TimeZone = %#v", typedResponse.Data.TimeZone, rawTimeZone)
	}

	rawSecurity := requireRawObject(t, raw, "security")
	if typedResponse.Data.Security == nil || typedResponse.Data.Security.ThreatScore == nil || *typedResponse.Data.Security.ThreatScore != requireRawFloat(t, rawSecurity, "threat_score") {
		t.Fatalf("typed Security = %#v raw Security = %#v", typedResponse.Data.Security, rawSecurity)
	}
}

func TestLiveBulkTypedAndRawShareSuccessIPValues(t *testing.T) {
	paidKey := requireLiveHardening(t)

	client, err := NewClient(&Config{APIKey: paidKey})
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}
	defer client.Close()

	typedResponse, err := client.BulkLookupIPGeolocation(context.Background(), &BulkLookupRequest{
		IPs:     []string{"8.8.8.8", "1.1.1.1"},
		Include: []string{"security"},
	})
	if err != nil {
		t.Fatalf("BulkLookupIPGeolocation() error = %v", err)
	}

	rawResponse, err := client.BulkLookupIPGeolocationRaw(context.Background(), &BulkLookupRequest{
		IPs:     []string{"8.8.8.8", "1.1.1.1"},
		Include: []string{"security"},
	})
	if err != nil {
		t.Fatalf("BulkLookupIPGeolocationRaw() error = %v", err)
	}

	var raw []map[string]interface{}
	if err := json.Unmarshal([]byte(rawResponse.Data), &raw); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}

	if len(typedResponse.Data) != len(raw) {
		t.Fatalf("len(typedResponse.Data) = %d len(raw) = %d", len(typedResponse.Data), len(raw))
	}

	for index, result := range typedResponse.Data {
		rawItem := raw[index]
		if result.Data == nil {
			t.Fatalf("typed result[%d] = %#v", index, result)
		}

		if result.Data.IP == nil || *result.Data.IP != rawItem["ip"].(string) {
			t.Fatalf("typed result[%d] = %#v raw = %#v", index, result, rawItem)
		}

		rawASN := rawItem["asn"].(map[string]interface{})
		if result.Data.ASN == nil || result.Data.ASN.ASNumber == nil || *result.Data.ASN.ASNumber != rawASN["as_number"].(string) {
			t.Fatalf("typed result[%d] ASN = %#v raw ASN = %#v", index, result.Data.ASN, rawASN)
		}

		rawLocation := requireRawObject(t, rawItem, "location")
		if result.Data.Location == nil || result.Data.Location.CountryName == nil || *result.Data.Location.CountryName != requireRawString(t, rawLocation, "country_name") {
			t.Fatalf("typed result[%d] Location = %#v raw Location = %#v", index, result.Data.Location, rawLocation)
		}

		rawNetwork := requireRawObject(t, rawItem, "network")
		if result.Data.Network == nil || result.Data.Network.Route == nil || *result.Data.Network.Route != requireRawString(t, rawNetwork, "route") {
			t.Fatalf("typed result[%d] Network = %#v raw Network = %#v", index, result.Data.Network, rawNetwork)
		}

		rawTimeZone := requireRawObject(t, rawItem, "time_zone")
		if result.Data.TimeZone == nil || result.Data.TimeZone.CurrentTimeZoneAbbreviation == nil || *result.Data.TimeZone.CurrentTimeZoneAbbreviation != requireRawString(t, rawTimeZone, "current_tz_abbreviation") {
			t.Fatalf("typed result[%d] TimeZone = %#v raw TimeZone = %#v", index, result.Data.TimeZone, rawTimeZone)
		}

		if result.Data.TimeZone.StandardTimeZoneAbbreviation == nil || *result.Data.TimeZone.StandardTimeZoneAbbreviation != requireRawString(t, rawTimeZone, "standard_tz_abbreviation") {
			t.Fatalf("typed result[%d] TimeZone = %#v raw TimeZone = %#v", index, result.Data.TimeZone, rawTimeZone)
		}

		rawSecurity := requireRawObject(t, rawItem, "security")
		if result.Data.Security == nil || result.Data.Security.ThreatScore == nil || *result.Data.Security.ThreatScore != requireRawFloat(t, rawSecurity, "threat_score") {
			t.Fatalf("typed result[%d] Security = %#v raw Security = %#v", index, result.Data.Security, rawSecurity)
		}
	}
}

func requireRawObject(t *testing.T, object map[string]interface{}, key string) map[string]interface{} {
	t.Helper()

	value, ok := object[key]
	if !ok {
		t.Fatalf("raw object missing %q: %#v", key, object)
	}

	nested, ok := value.(map[string]interface{})
	if !ok {
		t.Fatalf("raw object %q = %#v", key, value)
	}

	return nested
}

func requireRawString(t *testing.T, object map[string]interface{}, key string) string {
	t.Helper()

	value, ok := object[key]
	if !ok {
		t.Fatalf("raw string missing %q: %#v", key, object)
	}

	parsed, ok := value.(string)
	if !ok {
		t.Fatalf("raw string %q = %#v", key, value)
	}

	return parsed
}

func requireRawFloat(t *testing.T, object map[string]interface{}, key string) float64 {
	t.Helper()

	value, ok := object[key]
	if !ok {
		t.Fatalf("raw float missing %q: %#v", key, object)
	}

	parsed, ok := value.(float64)
	if !ok {
		t.Fatalf("raw float %q = %#v", key, value)
	}

	return parsed
}
