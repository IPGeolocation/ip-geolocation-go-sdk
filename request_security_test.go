package ipgeolocation

import (
	"context"
	"crypto/x509"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"syscall"
	"testing"
	"time"
)

func TestRedirectsKeepRequestDataOnOriginalServer(t *testing.T) {
	const key = "example-redirect-key"
	var received int32
	destination := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&received, 1)
	}))
	defer destination.Close()
	for _, code := range []int{301, 302, 303, 307, 308} {
		t.Run(strconv.Itoa(code), func(t *testing.T) {
			source := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				http.Redirect(w, r, destination.URL, code)
			}))
			defer source.Close()
			c, err := NewClient(&Config{APIKey: key, BaseURL: source.URL})
			if err != nil {
				t.Fatal(err)
			}
			defer c.Close()
			h := map[string]string{"X-Partner": "example-partner"}
			_, singleErr := c.LookupIPGeolocationRaw(context.Background(), &LookupRequest{IP: "192.0.2.1", Headers: h})
			_, bulkErr := c.BulkLookupIPGeolocationRaw(context.Background(), &BulkLookupRequest{IPs: []string{"192.0.2.1"}, Headers: h})
			for _, e := range []error{singleErr, bulkErr} {
				var api *APIError
				if !errors.As(e, &api) || api.StatusCode != code {
					t.Fatalf("expected HTTP %d, got %v", code, e)
				}
			}
		})
	}
	if atomic.LoadInt32(&received) != 0 {
		t.Fatal("a redirected request reached the other server")
	}
}

func TestRedirectsOnOriginalServer(t *testing.T) {
	const key = "example-same-host"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/final" {
			if r.URL.Query().Get("apiKey") != key {
				t.Error("missing key on original request")
			}
			http.Redirect(w, r, "/final", 307)
			return
		}
		data, _ := io.ReadAll(r.Body)
		if r.Method != "POST" || !strings.Contains(string(data), "192.0.2.2") || r.Header.Get("X-Partner") != "example" || r.Referer() != "" {
			t.Error("unexpected redirected request")
		}
		fmt.Fprint(w, "[{}]")
	}))
	defer server.Close()
	c, err := NewClient(&Config{APIKey: key, BaseURL: server.URL})
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()
	_, err = c.BulkLookupIPGeolocation(context.Background(), &BulkLookupRequest{IPs: []string{"192.0.2.2"}, Headers: map[string]string{"X-Partner": "example"}})
	if err != nil {
		t.Fatal(err)
	}
}

func TestRedirectOriginAndLimit(t *testing.T) {
	for _, pair := range []struct {
		from, to string
		allow    bool
	}{
		{"https://EXAMPLE.test/first", "https://example.test:443/second", true},
		{"http://example.test:80", "http://example.test/", true},
		{"https://example.test", "http://example.test", false},
		{"https://example.test", "https://example.test:444", false},
		{"https://example.test", "https://else.example.test", false},
		{"https://example.test", "https://user@example.test/", false},
		{"https://σ.invalid", "https://ς.invalid", false},
		{"http://[fe80::1%25en0]", "http://[fe80::1%25EN0]", false},
	} {
		old, _ := http.NewRequest("GET", pair.from, nil)
		next, _ := http.NewRequest("GET", pair.to, nil)
		err := checkRedirect(next, []*http.Request{old})
		if (err == nil) != pair.allow {
			t.Errorf("%s to %s: %v", pair.from, pair.to, err)
		}
	}
	req, _ := http.NewRequest("GET", "https://example.test", nil)
	via := make([]*http.Request, 10)
	via[0] = req
	if err := checkRedirect(req, via); err == nil {
		t.Fatal("redirect loop accepted")
	}
}

func TestKeyHiddenInFormattingAndErrors(t *testing.T) {
	const key = "synthetic +/suffix"
	cfg := Config{APIKey: key}
	c, err := NewClient(&cfg)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()
	for _, v := range []interface{}{cfg, &cfg, c, *c} {
		for _, format := range []string{"%v", "%+v", "%#v", "%s", "%q", "%d", "%x"} {
			got := fmt.Sprintf(format, v)
			if strings.Contains(got, key) || strings.Contains(got, fmt.Sprintf("%x", key)) {
				t.Errorf("format %s exposed the key", format)
			}
		}
	}
	for _, message := range []string{key, url.QueryEscape(key), url.PathEscape(key), "quota 90%: synthetic+%2b%2fsuffix", "quoted synthetic%20%2b%2fsuffix"} {
		c := mustNewTestClient(&cfg, stubTransport{send: func(_ context.Context, req *httpRequestData) (*httpResponseData, error) {
			body, _ := json.Marshal(map[string]string{"message": message})
			if req.Method == "POST" {
				return &httpResponseData{StatusCode: 200, Body: "[" + string(body) + "]"}, nil
			}
			return &httpResponseData{StatusCode: 403, Body: string(body)}, nil
		}})
		_, err := c.LookupIPGeolocationRaw(nil, nil)
		var forbidden *ForbiddenError
		if !errors.As(err, &forbidden) || !strings.Contains(forbidden.Message, "[REDACTED]") || strings.Contains(forbidden.Message, "suffix") {
			t.Errorf("unexpected HTTP error: %T", err)
		}
		req := &BulkLookupRequest{IPs: []string{"192.0.2.4"}}
		old, err := c.BulkLookupIPGeolocation(nil, req)
		if err != nil {
			t.Fatal(err)
		}
		full, err := c.BulkLookupIPGeolocationDetailed(nil, req)
		if err != nil {
			t.Fatal(err)
		}
		for _, row := range []*BulkLookupError{old.Data[0].Error, full.Data[0].Error} {
			if row == nil || !strings.Contains(row.Message, "[REDACTED]") || strings.Contains(row.Message, "suffix") {
				t.Error("typed bulk error exposed the key")
			}
		}
		raw, err := c.BulkLookupIPGeolocationRaw(nil, req)
		if err != nil || !strings.Contains(raw.Data, message) {
			t.Error("successful raw response changed")
		}
	}
	site := httptest.NewServer(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {}))
	site.Close()
	c, err = NewClient(&Config{APIKey: key, BaseURL: site.URL})
	if err != nil {
		t.Fatal(err)
	}
	_, err = c.LookupIPGeolocationRaw(nil, nil)
	var cause *url.Error
	if !errors.As(err, &cause) || strings.Contains(cause.URL, "apiKey") || !errors.Is(err, syscall.ECONNREFUSED) {
		t.Errorf("unexpected connection error: %v", err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err = c.LookupIPGeolocationRaw(ctx, nil)
	if !errors.Is(err, context.Canceled) {
		t.Error("lost caller cancellation")
	}
}

func TestHTTPStatusSurvivesBodyFailure(t *testing.T) {
	for _, tc := range []struct {
		name   string
		status int
		body   func() io.ReadCloser
	}{
		{"oversized", 429, func() io.ReadCloser { return io.NopCloser(strings.NewReader("12345")) }},
		{"short", 403, func() io.ReadCloser { return &faultingReadCloser{prefix: []byte("abc"), err: io.ErrUnexpectedEOF} }},
		{"timeout", 429, func() io.ReadCloser { return newBlockingReadCloser() }},
	} {
		t.Run(tc.name, func(t *testing.T) {
			tr := newHTTPTransportWithClient(&http.Client{Transport: stubRoundTripper{roundTrip: func(*http.Request) (*http.Response, error) {
				return &http.Response{StatusCode: tc.status, Body: tc.body()}, nil
			}}}, 4)
			c := mustNewTestClient(&Config{APIKey: "example", ReadTimeout: 20 * time.Millisecond}, tr)
			_, err := c.LookupIPGeolocation(nil, nil)
			var rate *TooManyRequestsError
			var forbidden *ForbiddenError
			if tc.status == 429 && (!errors.As(err, &rate) || rate.StatusCode != 429) || tc.status == 403 && (!errors.As(err, &forbidden) || forbidden.StatusCode != 403) {
				t.Fatalf("lost HTTP %d: %v", tc.status, err)
			}
		})
	}
}

func TestCustomDefaultTransportAndTimeout(t *testing.T) {
	saved := http.DefaultTransport
	defer func() { http.DefaultTransport = saved }()
	http.DefaultTransport = stubRoundTripper{roundTrip: func(r *http.Request) (*http.Response, error) {
		if r.URL.Query().Get("ip") == "slow" {
			<-r.Context().Done()
			return nil, r.Context().Err()
		}
		return &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader("{}"))}, nil
	}}
	c, err := NewClient(&Config{APIKey: "example", ConnectTimeout: 40 * time.Millisecond})
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()
	if _, err = c.LookupIPGeolocation(nil, nil); err != nil {
		t.Fatal(err)
	}
	_, err = c.LookupIPGeolocation(nil, &LookupRequest{IP: "slow"})
	var timeout *RequestTimeoutError
	if !errors.As(err, &timeout) || !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("lost custom connection timeout: %v", err)
	}
}

func TestConcurrentCloseAndLookup(t *testing.T) {
	c := mustNewTestClient(&Config{APIKey: "example"}, stubTransport{send: func(context.Context, *httpRequestData) (*httpResponseData, error) {
		return &httpResponseData{StatusCode: 200, Body: "{}"}, nil
	}})
	copied := *c
	copied.Close()
	if _, err := c.LookupIPGeolocation(nil, nil); err != nil {
		t.Fatalf("closing a copy closed the original: %v", err)
	}
	start := make(chan struct{})
	var workers sync.WaitGroup
	for i := 0; i < 32; i++ {
		workers.Add(1)
		go func(i int) {
			defer workers.Done()
			<-start
			if i%2 == 0 {
				c.Close()
			} else {
				_, _ = c.LookupIPGeolocation(nil, nil)
				_ = fmt.Sprint(c)
			}
		}(i)
	}
	close(start)
	workers.Wait()
	_, err := c.LookupIPGeolocation(nil, nil)
	var closed *ValidationError
	if !errors.As(err, &closed) {
		t.Fatal("closed client accepted lookup")
	}
}

func TestCertificateErrorsAndInvalidAddresses(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(http.ResponseWriter, *http.Request) { t.Error("untrusted server received a request") }))
	defer server.Close()
	server.Config.ErrorLog = log.New(io.Discard, "", 0)
	c, err := NewClient(&Config{APIKey: "example-tls-key", BaseURL: server.URL})
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()
	_, err = c.LookupIPGeolocation(nil, nil)
	var trust x509.UnknownAuthorityError
	var u *url.Error
	if !errors.As(err, &trust) || trust.Cert == nil || len(trust.Cert.Raw) == 0 || !errors.As(err, &u) || strings.Contains(u.URL, "apiKey") {
		t.Fatalf("lost certificate classification or exposed request URL: %v", err)
	}
	for _, u := range []string{"https:///path", "https:opaque", "https://example.test?", "https://example.test#"} {
		if _, err := NewClient(&Config{BaseURL: u}); err == nil {
			t.Errorf("accepted invalid base address %q", u)
		}
	}
	if _, err := NewClient(&Config{BaseURL: "https://example.test/path%23id"}); err != nil {
		t.Fatal(err)
	}
}
