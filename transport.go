package ipgeolocation

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"
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
}

func newHTTPTransport(connectTimeout time.Duration, maxResponseBodyBytes int) *httpTransport {
	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.DialContext = (&net.Dialer{
		Timeout:   connectTimeout,
		KeepAlive: 30 * time.Second,
	}).DialContext
	transport.ResponseHeaderTimeout = connectTimeout
	transport.TLSHandshakeTimeout = connectTimeout

	return &httpTransport{
		client: &http.Client{
			Transport: transport,
		},
		maxResponseBodyBytes: maxResponseBodyBytes,
	}
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

	response, err := t.client.Do(httpRequest)
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
	if err != nil {
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
