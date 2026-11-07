package ipgeolocation

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

const detailedRecord = `{"ip":"192.0.2.8","location":{"country_name":"اليابان"},"network":{"route":"192.0.2.0/24","is_cdn":false,"cdn_provider_name":"Example CDN"},"security":{"is_proxy":true,"proxy_days_seen_7d":0,"proxy_days_seen_30d":3,"proxy_provider_count_30d":2,"bot_type":"crawler","bot_last_seen":"2026-01-03","bot_confidence_score":0.5,"bot_operator_name":"Example","is_known_good_bot":false,"is_corporate_gateway":true,"corporate_gateway_type":"proxy","corporate_gateway_provider_name":"Example Gateway"}}`

func TestArabicAndDetailedResults(t *testing.T) {
	record := detailedRecord
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("lang") != "ar" {
			t.Error("lang=ar missing from request")
		}
		if r.URL.Query().Get("apiKey") != "example" {
			t.Error("missing authentication")
		}
		switch r.URL.Path {
		case "/v3/ipgeo":
			fmt.Fprint(w, record)
		case "/v3/ipgeo-bulk":
			var body struct {
				IPs []string `json:"ips"`
			}
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil || !reflect.DeepEqual(body.IPs, []string{"192.0.2.8", "bad", "192.0.2.8"}) {
				t.Error("wrong bulk inputs")
			}
			fmt.Fprintf(w, "[%s,{\"message\":\"invalid IP\"},%s]", record, record)
		default:
			t.Error("unexpected endpoint")
			w.WriteHeader(404)
		}
	}))
	defer srv.Close()
	c, err := NewClient(&Config{APIKey: "example", BaseURL: srv.URL})
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()
	req := &LookupRequest{IP: "192.0.2.8", Lang: " AR "}
	single, err := c.LookupIPGeolocation(nil, req)
	if err != nil || *single.Data.Location.CountryName != "اليابان" {
		t.Fatalf("ordinary Arabic lookup: %v", err)
	}
	full, err := c.LookupIPGeolocationDetailed(nil, req)
	if err != nil {
		t.Fatal(err)
	}
	checkDetailed := func(d *DetailedIPGeolocationResponse) {
		t.Helper()
		got, err := json.Marshal(d)
		if err != nil {
			t.Fatal(err)
		}
		var expected, actual map[string]interface{}
		if json.Unmarshal([]byte(record), &expected) != nil || json.Unmarshal(got, &actual) != nil || !reflect.DeepEqual(expected, actual) {
			t.Errorf("missing or changed detail: %s", string(got))
		}
		if d.Network.Network != d.IPGeolocationResponse.Network || d.Security.Security != d.IPGeolocationResponse.Security {
			t.Error("old and detailed fields differ")
		}
	}
	checkDetailed(full.Data)
	bulkReq := &BulkLookupRequest{IPs: []string{"192.0.2.8", "bad", "192.0.2.8"}, Lang: "ar"}
	bulk, err := c.BulkLookupIPGeolocation(nil, bulkReq)
	if err != nil || len(bulk.Data) != 3 || bulk.Data[1].Error == nil {
		t.Fatalf("ordinary Arabic bulk: %v", err)
	}
	detailed, err := c.BulkLookupIPGeolocationDetailed(nil, bulkReq)
	if err != nil || len(detailed.Data) != 3 || !detailed.Data[0].IsSuccess() || detailed.Data[1].IsSuccess() || detailed.Data[1].Error.Message != "invalid IP" {
		t.Fatalf("detailed bulk: %v", err)
	}
	checkDetailed(detailed.Data[0].Data)
	checkDetailed(detailed.Data[2].Data)
}

func TestBulkResponseValidationAndRawOutput(t *testing.T) {
	req := &BulkLookupRequest{IPs: []string{"192.0.2.8", "192.0.2.8"}}
	for _, body := range []string{"null", "{}", "[]", "[{}]", "[{},{},{}]", "[{},null]", "[{},[]]", "[{},1]", "[{},{]", "[{},{}]true", "[{},{}]\u00a0"} {
		c := mustNewTestClient(&Config{APIKey: "example"}, stubTransport{send: func(context.Context, *httpRequestData) (*httpResponseData, error) {
			return &httpResponseData{StatusCode: 200, Body: body}, nil
		}})
		_, err := c.BulkLookupIPGeolocation(nil, req)
		var ser *SerializationError
		if !errors.As(err, &ser) {
			t.Errorf("ordinary lookup accepted %q", body)
		}
		_, err = c.BulkLookupIPGeolocationDetailed(nil, req)
		if !errors.As(err, &ser) {
			t.Errorf("detailed lookup accepted %q", body)
		}
		raw, err := c.BulkLookupIPGeolocationRaw(nil, req)
		if err != nil || raw.Data != body {
			t.Errorf("raw lookup changed %q", body)
		}
	}
}

func TestDetailedLookupOmittedFieldsAndFormats(t *testing.T) {
	for _, body := range []string{"{}", `{"network":null,"security":[]}`} {
		c := mustNewTestClient(&Config{APIKey: "example"}, stubTransport{send: func(context.Context, *httpRequestData) (*httpResponseData, error) {
			return &httpResponseData{StatusCode: 200, Body: body}, nil
		}})
		result, err := c.LookupIPGeolocationDetailed(nil, nil)
		if err != nil || result.Data.Network != nil || result.Data.Security != nil {
			t.Fatalf("omitted fields: %#v %v", result, err)
		}
	}
	c := mustNewTestClient(&Config{RequestOrigin: "https://example.test"}, stubTransport{send: func(context.Context, *httpRequestData) (*httpResponseData, error) {
		return &httpResponseData{StatusCode: 200, Body: "{}"}, nil
	}})
	if _, err := c.LookupIPGeolocationDetailed(nil, nil); err != nil {
		t.Fatal(err)
	}
	if _, err := c.BulkLookupIPGeolocationDetailed(nil, &BulkLookupRequest{IPs: []string{"192.0.2.8"}}); err == nil {
		t.Fatal("bulk authentication missing")
	}
	if _, err := c.LookupIPGeolocationDetailed(nil, &LookupRequest{Output: ResponseFormatXML}); err == nil {
		t.Fatal("typed XML accepted")
	}
}

func TestMaximumBulkSize(t *testing.T) {
	const count = 50000
	ips := make([]string, count)
	for i := range ips {
		ips[i] = "192.0.2.8"
	}
	body := "[" + strings.Repeat("{},", count-1) + `{"ip":"192.0.2.8"}]`
	c := mustNewTestClient(&Config{APIKey: "example"}, stubTransport{send: func(context.Context, *httpRequestData) (*httpResponseData, error) {
		return &httpResponseData{StatusCode: 200, Body: body}, nil
	}})
	r, err := c.BulkLookupIPGeolocation(nil, &BulkLookupRequest{IPs: ips})
	if err != nil || len(r.Data) != count || *r.Data[count-1].Data.IP != ips[count-1] {
		t.Fatalf("maximum bulk lookup: %v", err)
	}
}
