package ipgeolocation

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
	"unicode"
)

type transport interface {
	Send(ctx context.Context, request *httpRequestData) (*httpResponseData, error)
	Close()
}

type httpRequestData struct {
	URL         string
	Method      string
	Headers     http.Header
	Body        []byte
	ReadTimeout time.Duration
}

type httpResponseData struct {
	StatusCode int
	Body       string
	Headers    http.Header
}

type httpTransport struct {
	client               *http.Client
	maxResponseBodyBytes int
	customHeaderTimeout  time.Duration
}

func newHTTPTransport(connectTimeout time.Duration, maxResponseBodyBytes int) *httpTransport {
	base := http.DefaultTransport
	var customHeaderTimeout time.Duration
	if tr, ok := base.(*http.Transport); ok || base == nil {
		if tr == nil {
			tr = &http.Transport{Proxy: http.ProxyFromEnvironment, ForceAttemptHTTP2: true}
		}
		owned := tr.Clone()
		owned.DialContext = (&net.Dialer{Timeout: connectTimeout, KeepAlive: 30 * time.Second}).DialContext
		owned.ResponseHeaderTimeout = connectTimeout
		owned.TLSHandshakeTimeout = connectTimeout
		base = owned
	} else {
		customHeaderTimeout = connectTimeout
	}
	return &httpTransport{client: &http.Client{Transport: base, CheckRedirect: checkRedirect}, maxResponseBodyBytes: maxResponseBodyBytes, customHeaderTimeout: customHeaderTimeout}
}

func checkRedirect(req *http.Request, via []*http.Request) error {
	first := via[0].URL
	if req.URL.User != nil || first.Scheme != req.URL.Scheme || !sameHostname(first.Hostname(), req.URL.Hostname()) || originPort(first) != originPort(req.URL) {
		return http.ErrUseLastResponse
	}
	if len(via) >= 10 {
		return errors.New("stopped after 10 redirects")
	}
	req.Header.Del("Referer")
	return nil
}

func sameHostname(a, b string) bool {
	if a == b {
		return true
	}
	if strings.Contains(a, "%") || strings.Contains(b, "%") || strings.IndexFunc(a+b, func(r rune) bool { return r > unicode.MaxASCII }) >= 0 {
		return false
	}
	return strings.EqualFold(a, b)
}

func originPort(u *url.URL) string {
	if p := u.Port(); p != "" {
		return p
	}
	if u.Scheme == "https" {
		return "443"
	}
	return "80"
}

func newHTTPTransportWithClient(client *http.Client, maxResponseBodyBytes int) *httpTransport {
	if client == nil {
		client = &http.Client{}
	}

	return &httpTransport{
		client:               client,
		maxResponseBodyBytes: maxResponseBodyBytes,
	}
}

func (t *httpTransport) Close() {
	t.client.CloseIdleConnections()
}

func (t *httpTransport) Send(ctx context.Context, request *httpRequestData) (*httpResponseData, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	var body io.Reader
	if len(request.Body) > 0 {
		body = bytes.NewReader(request.Body)
	}

	httpRequest, err := http.NewRequestWithContext(ctx, request.Method, request.URL, body)
	if err != nil {
		return nil, &TransportError{Message: "HTTP transport error", Err: err}
	}

	for name, values := range request.Headers {
		for _, value := range values {
			httpRequest.Header.Add(name, value)
		}
	}

	var timer *time.Timer
	if t.customHeaderTimeout > 0 {
		headerCtx, cancel := context.WithCancel(ctx)
		defer cancel()
		httpRequest = httpRequest.WithContext(headerCtx)
		timer = time.AfterFunc(t.customHeaderTimeout, cancel)
	}
	response, err := t.client.Do(httpRequest)
	if timer != nil && !timer.Stop() && ctx.Err() == nil {
		if response != nil {
			response.Body.Close()
		}
		err = context.DeadlineExceeded
	}
	if err != nil {
		if ctx.Err() == nil && isTimeoutError(err) {
			return nil, &RequestTimeoutError{
				TransportError: TransportError{
					Message: "HTTP request timed out while opening the connection or waiting for response headers",
					Err:     err,
				},
			}
		}

		return nil, &TransportError{Message: "HTTP transport error", Err: err}
	}

	bodyBytes, err := t.readBody(response.Body, request.ReadTimeout)
	if err != nil && isSuccessStatus(response.StatusCode) {
		return nil, err
	}

	return &httpResponseData{
		StatusCode: response.StatusCode,
		Body:       string(bodyBytes),
		Headers:    cloneHeader(response.Header),
	}, nil
}

func (t *httpTransport) readBody(body io.ReadCloser, readTimeout time.Duration) ([]byte, error) {
	defer body.Close()

	buffer := bytes.NewBuffer(nil)
	chunk := make([]byte, 8192)

	for {
		bytesRead, err := t.readChunk(body, chunk, readTimeout)
		if bytesRead > 0 {
			if buffer.Len()+bytesRead > t.maxResponseBodyBytes {
				return nil, &TransportError{
					Message: fmt.Sprintf("Response body exceeded max size of %d bytes", t.maxResponseBodyBytes),
				}
			}

			if _, writeErr := buffer.Write(chunk[:bytesRead]); writeErr != nil {
				return nil, &TransportError{Message: "HTTP transport error while buffering the response body", Err: writeErr}
			}
		}

		if err == nil {
			continue
		}

		if errors.Is(err, io.EOF) {
			break
		}

		return nil, err
	}

	return buffer.Bytes(), nil
}

func (t *httpTransport) readChunk(body io.ReadCloser, chunk []byte, readTimeout time.Duration) (int, error) {
	type readResult struct {
		BytesRead int
		Err       error
	}

	resultCh := make(chan readResult, 1)
	go func() {
		n, err := body.Read(chunk)
		resultCh <- readResult{BytesRead: n, Err: err}
	}()

	timer := time.NewTimer(readTimeout)
	defer timer.Stop()

	select {
	case result := <-resultCh:
		if result.Err == nil || errors.Is(result.Err, io.EOF) {
			return result.BytesRead, result.Err
		}

		if isTimeoutError(result.Err) {
			return result.BytesRead, &RequestTimeoutError{
				TransportError: TransportError{
					Message: "HTTP request timed out while waiting for response body data",
					Err:     result.Err,
				},
			}
		}

		return result.BytesRead, &TransportError{
			Message: "HTTP transport error while reading the response body",
			Err:     result.Err,
		}
	case <-timer.C:
		_ = body.Close()
		return 0, &RequestTimeoutError{
			TransportError: TransportError{
				Message: "HTTP request timed out while waiting for response body data",
				Err:     context.DeadlineExceeded,
			},
		}
	}
}

func cloneHeader(header http.Header) http.Header {
	if header == nil {
		return http.Header{}
	}

	return header.Clone()
}

func isTimeoutError(err error) bool {
	if errors.Is(err, context.DeadlineExceeded) {
		return true
	}

	var netError net.Error
	return errors.As(err, &netError) && netError.Timeout()
}
