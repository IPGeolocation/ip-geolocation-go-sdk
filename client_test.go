package ipgeolocation

import (
	"context"
	"errors"
	"net/http"
	"testing"
)

func TestLookupIPGeolocationRejectsTypedXML(t *testing.T) {
	client := mustNewTestClient(&Config{APIKey: "test-key"}, stubTransport{
		send: func(ctx context.Context, request *httpRequestData) (*httpResponseData, error) {
			t.Fatal("transport should not be called")
			return nil, nil
		},
	})

	_, err := client.LookupIPGeolocation(context.Background(), &LookupRequest{Output: ResponseFormatXML})
	var validationErr *ValidationError
	if !errors.As(err, &validationErr) || validationErr.Message != "typed methods support JSON only" {
		t.Fatalf("LookupIPGeolocation() error = %v", err)
	}
}

func TestLookupIPGeolocationRequiresAPIKeyOrRequestOrigin(t *testing.T) {
	client := mustNewTestClient(&Config{}, stubTransport{
		send: func(ctx context.Context, request *httpRequestData) (*httpResponseData, error) {
			t.Fatal("transport should not be called")
			return nil, nil
		},
	})

	_, err := client.LookupIPGeolocation(context.Background(), nil)
	var validationErr *ValidationError
	if !errors.As(err, &validationErr) || validationErr.Message != "single lookup requires apiKey or requestOrigin in client config" {
		t.Fatalf("LookupIPGeolocation() error = %v", err)
	}
}

func TestBulkLookupIPGeolocationRequiresAPIKey(t *testing.T) {
	client := mustNewTestClient(&Config{}, stubTransport{
		send: func(ctx context.Context, request *httpRequestData) (*httpResponseData, error) {
			t.Fatal("transport should not be called")
			return nil, nil
		},
	})

	_, err := client.BulkLookupIPGeolocation(context.Background(), &BulkLookupRequest{IPs: []string{"8.8.8.8"}})
	var validationErr *ValidationError
	if !errors.As(err, &validationErr) || validationErr.Message != "bulk lookup requires apiKey in client config" {
		t.Fatalf("BulkLookupIPGeolocation() error = %v", err)
	}
}

func TestLookupIPGeolocationReturnsTypedDataAndMetadata(t *testing.T) {
	client := mustNewTestClient(&Config{APIKey: "test-key"}, stubTransport{
		send: func(ctx context.Context, request *httpRequestData) (*httpResponseData, error) {
			if request.URL != "https://api.ipgeolocation.io/v3/ipgeo?apiKey=test-key&ip=8.8.8.8&output=json" {
				t.Fatalf("request.URL = %q", request.URL)
			}

			if request.Headers.Get("User-Agent") != DefaultUserAgent() {
				t.Fatalf("request.Headers = %#v", request.Headers)
			}

			return &httpResponseData{
				StatusCode: 200,
				Body:       `{"ip":"8.8.8.8","location":{"country_name":"United States"}}`,
				Headers: http.Header{
					"X-Credits-Charged": []string{"1"},
				},
			}, nil
		},
	})

	response, err := client.LookupIPGeolocation(context.Background(), &LookupRequest{IP: "8.8.8.8"})
	if err != nil {
		t.Fatalf("LookupIPGeolocation() error = %v", err)
	}

	if response.Data == nil || response.Data.IP == nil || *response.Data.IP != "8.8.8.8" {
		t.Fatalf("response.Data = %#v", response.Data)
	}

	if response.Metadata.CreditsCharged == nil || *response.Metadata.CreditsCharged != 1 {
		t.Fatalf("response.Metadata = %#v", response.Metadata)
	}
}

func TestBulkLookupIPGeolocationReturnsDataAndErrors(t *testing.T) {
	client := mustNewTestClient(&Config{APIKey: "paid-key"}, stubTransport{
		send: func(ctx context.Context, request *httpRequestData) (*httpResponseData, error) {
			if request.Method != http.MethodPost {
				t.Fatalf("request.Method = %q", request.Method)
			}

			return &httpResponseData{
				StatusCode: 200,
				Body:       `[{"ip":"8.8.8.8"},{"error":{"message":"Invalid IP address."}}]`,
				Headers:    http.Header{},
			}, nil
		},
	})

	response, err := client.BulkLookupIPGeolocation(context.Background(), &BulkLookupRequest{IPs: []string{"8.8.8.8", "bad value"}})
	if err != nil {
		t.Fatalf("BulkLookupIPGeolocation() error = %v", err)
	}

	if len(response.Data) != 2 {
		t.Fatalf("len(response.Data) = %d", len(response.Data))
	}

	if response.Data[0].Data == nil || response.Data[0].Data.IP == nil || *response.Data[0].Data.IP != "8.8.8.8" {
		t.Fatalf("response.Data[0] = %#v", response.Data[0])
	}

	if response.Data[1].Error == nil || response.Data[1].Error.Message != "Invalid IP address." {
		t.Fatalf("response.Data[1] = %#v", response.Data[1])
	}
}

func TestBulkLookupIPGeolocationPreservesDuplicateInputsInRequestBody(t *testing.T) {
	client := mustNewTestClient(&Config{APIKey: "paid-key"}, stubTransport{
		send: func(ctx context.Context, request *httpRequestData) (*httpResponseData, error) {
			expectedBody := `{"ips":["8.8.8.8","1.1.1.1","8.8.8.8","example.com","example.com"]}`
			if string(request.Body) != expectedBody {
				t.Fatalf("request.Body = %s", request.Body)
			}

			return &httpResponseData{
				StatusCode: 200,
				Body:       `[{"ip":"8.8.8.8"},{"ip":"1.1.1.1"},{"ip":"8.8.8.8"},{"ip":"example.com"},{"ip":"example.com"}]`,
				Headers:    http.Header{},
			}, nil
		},
	})

	response, err := client.BulkLookupIPGeolocation(context.Background(), &BulkLookupRequest{
		IPs: []string{" 8.8.8.8 ", "1.1.1.1", "8.8.8.8", "example.com", "example.com"},
	})
	if err != nil {
		t.Fatalf("BulkLookupIPGeolocation() error = %v", err)
	}

	if len(response.Data) != 5 {
		t.Fatalf("len(response.Data) = %d", len(response.Data))
	}

	if response.Data[2].Data == nil || response.Data[2].Data.IP == nil || *response.Data[2].Data.IP != "8.8.8.8" {
		t.Fatalf("response.Data[2] = %#v", response.Data[2])
	}

	if response.Data[4].Data == nil || response.Data[4].Data.IP == nil || *response.Data[4].Data.IP != "example.com" {
		t.Fatalf("response.Data[4] = %#v", response.Data[4])
	}
}
