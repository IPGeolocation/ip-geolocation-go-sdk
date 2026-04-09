package ipgeolocation

import (
	"encoding/json"
	"testing"
)

func TestParseBulkLookupMixedSuccessAndError(t *testing.T) {
	body := `[
	  { "ip": "8.8.8.8", "location": { "country_name": "United States" } },
	  { "message": "Invalid IP address." },
	  { "error": { "message": "Quota exceeded." } }
	]`

	result, err := parseBulkLookup(body)
	if err != nil {
		t.Fatalf("parseBulkLookup() error = %v", err)
	}

	if len(result) != 3 {
		t.Fatalf("len(result) = %d", len(result))
	}

	if result[0].Data == nil || result[0].Data.IP == nil || *result[0].Data.IP != "8.8.8.8" {
		t.Fatalf("result[0] = %#v", result[0])
	}

	if result[1].Error == nil || result[1].Error.Message != "Invalid IP address." {
		t.Fatalf("result[1] = %#v", result[1])
	}

	if result[2].Error == nil || result[2].Error.Message != "Quota exceeded." {
		t.Fatalf("result[2] = %#v", result[2])
	}
}

func TestExtractAPIMessagePrefersNestedErrorMessage(t *testing.T) {
	message := extractAPIMessage(`{"error":{"message":"Invalid IP address."}}`)

	if message != "Invalid IP address." {
		t.Fatalf("extractAPIMessage() = %q", message)
	}
}

func TestParseSingleLookupDoesNotInventMissingFields(t *testing.T) {
	response, err := parseSingleLookup(`{
	  "ip": "8.8.8.8",
	  "location": { "country_name": "United States" },
	  "security": { "threat_score": 1.5 }
	}`)
	if err != nil {
		t.Fatalf("parseSingleLookup() error = %v", err)
	}

	if response.IP == nil || *response.IP != "8.8.8.8" {
		t.Fatalf("response.IP = %#v", response.IP)
	}

	if response.Location == nil || response.Location.CountryName == nil || *response.Location.CountryName != "United States" {
		t.Fatalf("response.Location = %#v", response.Location)
	}

	if response.Security == nil || response.Security.ThreatScore == nil || *response.Security.ThreatScore != 1.5 {
		t.Fatalf("response.Security = %#v", response.Security)
	}

	if response.TimeZone != nil {
		t.Fatalf("response.TimeZone = %#v", response.TimeZone)
	}

	if response.Network != nil {
		t.Fatalf("response.Network = %#v", response.Network)
	}
}

func TestParseSingleLookupTopLevelASNAndCompany(t *testing.T) {
	response, err := parseSingleLookup(`{
	  "ip": "8.8.8.8",
	  "network": {
	    "connection_type": "Corporate",
	    "route": "8.8.8.0/24",
	    "is_anycast": true
	  },
	  "asn": {
	    "as_number": "AS15169",
	    "organization": "Google LLC"
	  },
	  "company": {
	    "name": "Google LLC",
	    "domain": "google.com"
	  }
	}`)
	if err != nil {
		t.Fatalf("parseSingleLookup() error = %v", err)
	}

	if response.Network == nil || response.Network.ConnectionType == nil || *response.Network.ConnectionType != "Corporate" {
		t.Fatalf("response.Network = %#v", response.Network)
	}

	if response.ASN == nil || response.ASN.ASNumber == nil || *response.ASN.ASNumber != "AS15169" {
		t.Fatalf("response.ASN = %#v", response.ASN)
	}

	if response.Company == nil || response.Company.Name == nil || *response.Company.Name != "Google LLC" {
		t.Fatalf("response.Company = %#v", response.Company)
	}
}

func TestParseSingleLookupTimeZoneUsesCurrentAPIKeys(t *testing.T) {
	response, err := parseSingleLookup(`{
	  "ip": "8.8.8.8",
	  "time_zone": {
	    "name": "America/Chicago",
	    "current_tz_abbreviation": "CDT",
	    "current_tz_full_name": "Central Daylight Time",
	    "standard_tz_abbreviation": "CST",
	    "standard_tz_full_name": "Central Standard Time",
	    "dst_tz_abbreviation": "CDT",
	    "dst_tz_full_name": "Central Daylight Time"
	  }
	}`)
	if err != nil {
		t.Fatalf("parseSingleLookup() error = %v", err)
	}

	if response.TimeZone == nil || response.TimeZone.CurrentTimeZoneAbbreviation == nil || *response.TimeZone.CurrentTimeZoneAbbreviation != "CDT" {
		t.Fatalf("response.TimeZone = %#v", response.TimeZone)
	}

	if response.TimeZone.CurrentTimeZoneFullName == nil || *response.TimeZone.CurrentTimeZoneFullName != "Central Daylight Time" {
		t.Fatalf("response.TimeZone = %#v", response.TimeZone)
	}

	if response.TimeZone.StandardTimeZoneAbbreviation == nil || *response.TimeZone.StandardTimeZoneAbbreviation != "CST" {
		t.Fatalf("response.TimeZone = %#v", response.TimeZone)
	}

	if response.TimeZone.StandardTimeZoneFullName == nil || *response.TimeZone.StandardTimeZoneFullName != "Central Standard Time" {
		t.Fatalf("response.TimeZone = %#v", response.TimeZone)
	}

	if response.TimeZone.DSTTimeZoneAbbreviation == nil || *response.TimeZone.DSTTimeZoneAbbreviation != "CDT" {
		t.Fatalf("response.TimeZone = %#v", response.TimeZone)
	}

	if response.TimeZone.DSTTimeZoneFullName == nil || *response.TimeZone.DSTTimeZoneFullName != "Central Daylight Time" {
		t.Fatalf("response.TimeZone = %#v", response.TimeZone)
	}
}

func TestParseSingleLookupTimeZoneFallsBackToLegacyKeys(t *testing.T) {
	response, err := parseSingleLookup(`{
	  "ip": "8.8.8.8",
	  "time_zone": {
	    "current_timezone_abbreviation": "CDT",
	    "current_timezone_name": "Central Daylight Time",
	    "timezone_abbreviation": "CST",
	    "timezone_name": "Central Standard Time",
	    "dst_timezone_abbreviation": "CDT",
	    "dst_timezone_name": "Central Daylight Time"
	  }
	}`)
	if err != nil {
		t.Fatalf("parseSingleLookup() error = %v", err)
	}

	if response.TimeZone == nil || response.TimeZone.CurrentTimeZoneAbbreviation == nil || *response.TimeZone.CurrentTimeZoneAbbreviation != "CDT" {
		t.Fatalf("response.TimeZone = %#v", response.TimeZone)
	}

	if response.TimeZone.CurrentTimeZoneFullName == nil || *response.TimeZone.CurrentTimeZoneFullName != "Central Daylight Time" {
		t.Fatalf("response.TimeZone = %#v", response.TimeZone)
	}

	if response.TimeZone.StandardTimeZoneAbbreviation == nil || *response.TimeZone.StandardTimeZoneAbbreviation != "CST" {
		t.Fatalf("response.TimeZone = %#v", response.TimeZone)
	}

	if response.TimeZone.StandardTimeZoneFullName == nil || *response.TimeZone.StandardTimeZoneFullName != "Central Standard Time" {
		t.Fatalf("response.TimeZone = %#v", response.TimeZone)
	}

	if response.TimeZone.DSTTimeZoneAbbreviation == nil || *response.TimeZone.DSTTimeZoneAbbreviation != "CDT" {
		t.Fatalf("response.TimeZone = %#v", response.TimeZone)
	}

	if response.TimeZone.DSTTimeZoneFullName == nil || *response.TimeZone.DSTTimeZoneFullName != "Central Daylight Time" {
		t.Fatalf("response.TimeZone = %#v", response.TimeZone)
	}
}

func TestBulkLookupResultMarshalJSONUsesDataAndErrorKeys(t *testing.T) {
	ip := "8.8.8.8"

	body, err := json.Marshal([]BulkLookupResult{
		{
			Data: &IPGeolocationResponse{
				IP: &ip,
			},
		},
		{
			Error: &BulkLookupError{Message: "Invalid IP address."},
		},
	})
	if err != nil {
		t.Fatalf("json.Marshal() error = %v", err)
	}

	expected := `[{"data":{"ip":"8.8.8.8"}},{"error":{"message":"Invalid IP address."}}]`
	if string(body) != expected {
		t.Fatalf("json.Marshal() = %s", body)
	}
}
