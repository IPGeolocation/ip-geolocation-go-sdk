package ipgeolocation

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
	"sync"
	"testing"
	"time"
)

func TestHTTPTransportReturnsResponseDataAndHeaders(t *testing.T) {
	client := &http.Client{
		Transport: stubRoundTripper{
			roundTrip: func(request *http.Request) (*http.Response, error) {
				response := &http.Response{
					StatusCode: 200,
					Header:     header("X-Credits-Charged", "1"),
					Body:       io.NopCloser(strings.NewReader(`{"ip":"8.8.8.8"}`)),
				}
				return response, nil
			},
		},
	}

	transport := newHTTPTransportWithClient(client, DefaultMaxResponseBodyBytes)
	response, err := transport.Send(context.Background(), &httpRequestData{
		URL:         "https://api.ipgeolocation.io/v3/ipgeo?apiKey=test-key",
		Method:      http.MethodGet,
		Headers:     http.Header{"Accept": []string{"application/json"}},
		ReadTimeout: 30 * time.Second,
	})
	if err != nil {
		t.Fatalf("Send() error = %v", err)
	}

	if response.StatusCode != 200 {
		t.Fatalf("response.StatusCode = %d", response.StatusCode)
	}

	if response.Body != `{"ip":"8.8.8.8"}` {
		t.Fatalf("response.Body = %q", response.Body)
	}

	if response.Headers.Get("X-Credits-Charged") != "1" {
		t.Fatalf("response.Headers = %#v", response.Headers)
	}
}

func TestHTTPTransportRejectsOversizedResponseBody(t *testing.T) {
	client := &http.Client{
		Transport: stubRoundTripper{
			roundTrip: func(request *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 200,
					Body:       io.NopCloser(strings.NewReader("12345")),
					Header:     http.Header{},
				}, nil
			},
		},
	}

	transport := newHTTPTransportWithClient(client, 4)
	_, err := transport.Send(context.Background(), &httpRequestData{
		URL:         "https://api.ipgeolocation.io/v3/ipgeo",
		Method:      http.MethodGet,
		Headers:     http.Header{},
		ReadTimeout: 30 * time.Second,
	})
	if err == nil || err.Error() != "Response body exceeded max size of 4 bytes" {
		t.Fatalf("Send() error = %v", err)
	}
}

func TestHTTPTransportWrapsMidBodyReadFailures(t *testing.T) {
	client := &http.Client{
		Transport: stubRoundTripper{
			roundTrip: func(request *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 200,
					Header:     http.Header{},
					Body: &faultingReadCloser{
						prefix: []byte(`{"ip":"8.8.8.8"}`),
						err:    errors.New("stream read failed"),
					},
				}, nil
			},
		},
	}

	transport := newHTTPTransportWithClient(client, DefaultMaxResponseBodyBytes)
	_, err := transport.Send(context.Background(), &httpRequestData{
		URL:         "https://api.ipgeolocation.io/v3/ipgeo",
		Method:      http.MethodGet,
		Headers:     http.Header{},
		ReadTimeout: 30 * time.Second,
	})
	var transportErr *TransportError
	if !errors.As(err, &transportErr) || transportErr.Message != "HTTP transport error while reading the response body" {
		t.Fatalf("Send() error = %v", err)
	}
}

func TestHTTPTransportWrapsReadTimeoutWhileReadingBody(t *testing.T) {
	client := &http.Client{
		Transport: stubRoundTripper{
			roundTrip: func(request *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 200,
					Header:     http.Header{},
					Body:       newBlockingReadCloser(),
				}, nil
			},
		},
	}

	transport := newHTTPTransportWithClient(client, DefaultMaxResponseBodyBytes)
	_, err := transport.Send(context.Background(), &httpRequestData{
		URL:         "https://api.ipgeolocation.io/v3/ipgeo",
		Method:      http.MethodGet,
		Headers:     http.Header{},
		ReadTimeout: 20 * time.Millisecond,
	})
	var timeoutErr *RequestTimeoutError
	if !errors.As(err, &timeoutErr) || timeoutErr.Message != "HTTP request timed out while waiting for response body data" {
		t.Fatalf("Send() error = %v", err)
	}
}

type stubRoundTripper struct {
	roundTrip func(*http.Request) (*http.Response, error)
}

func (s stubRoundTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	return s.roundTrip(request)
}

type faultingReadCloser struct {
	prefix []byte
	err    error
	index  int
	failed bool
}

func (r *faultingReadCloser) Read(buffer []byte) (int, error) {
	if r.index < len(r.prefix) {
		bytesRead := copy(buffer, r.prefix[r.index:])
		r.index += bytesRead
		return bytesRead, nil
	}

	if r.failed {
		return 0, io.EOF
	}

	r.failed = true
	return 0, r.err
}

func (r *faultingReadCloser) Close() error {
	return nil
}

type blockingReadCloser struct {
	closed chan struct{}
	once   sync.Once
}

func newBlockingReadCloser() *blockingReadCloser {
	return &blockingReadCloser{
		closed: make(chan struct{}),
	}
}

func (r *blockingReadCloser) Read(buffer []byte) (int, error) {
	<-r.closed
	return 0, io.EOF
}

func (r *blockingReadCloser) Close() error {
	r.once.Do(func() {
		close(r.closed)
	})

	return nil
}
