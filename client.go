package ipgeolocation

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"sync/atomic"
	"time"
)

// Client is the main client for the IPGeolocation.io IP Location API.
//
// Homepage: https://ipgeolocation.io
//
// IP Location API: https://ipgeolocation.io/ip-location-api.html
//
// Documentation: https://ipgeolocation.io/documentation/ip-location-api.html
type Client struct {
	clientSettings
	transport transport
	closed    uint32
}

func NewClient(config *Config) (*Client, error) {
	return newClientWithTransport(config, nil)
}

func newClientWithTransport(config *Config, transport transport) (*Client, error) {
	normalized, err := normalizeConfig(config)
	if err != nil {
		return nil, err
	}

	if transport == nil {
		transport = newHTTPTransport(normalized.ConnectTimeout, DefaultMaxResponseBodyBytes)
	}

	return &Client{
		clientSettings: clientSettings{config: normalized},
		transport:      transport,
	}, nil
}

func (c *Client) Close() {
	if atomic.CompareAndSwapUint32(&c.closed, 0, 1) && c.transport != nil {
		c.transport.Close()
	}
}

func (c *Client) LookupIPGeolocation(ctx context.Context, request *LookupRequest) (*APIResponse[*IPGeolocationResponse], error) {
	return lookupTyped(c, ctx, request, parseResponse)
}

func (c *Client) LookupIPGeolocationDetailed(ctx context.Context, request *LookupRequest) (*APIResponse[*DetailedIPGeolocationResponse], error) {
	return lookupTyped(c, ctx, request, parseDetailedResponse)
}

func lookupTyped[T any](c *Client, ctx context.Context, request *LookupRequest, parse func(map[string]json.RawMessage) (*T, error)) (*APIResponse[*T], error) {
	raw, err := c.single(ctx, request, true)
	if err != nil {
		return nil, err
	}
	data, err := parseSingleLookup(raw.Data, parse)
	if err != nil {
		return nil, err
	}
	return &APIResponse[*T]{Data: data, Metadata: raw.Metadata}, nil
}

func (c *Client) LookupIPGeolocationRaw(ctx context.Context, request *LookupRequest) (*APIResponse[string], error) {
	return c.single(ctx, request, false)
}

func (c *Client) single(ctx context.Context, request *LookupRequest, typed bool) (*APIResponse[string], error) {
	if err := c.ensureReady(); err != nil {
		return nil, err
	}
	if c.config.APIKey == "" && c.config.RequestOrigin == "" {
		return nil, &ValidationError{Message: "single lookup requires apiKey or requestOrigin in client config"}
	}
	normalized, err := normalizeLookupRequest(request)
	if err != nil {
		return nil, err
	}
	if typed {
		if err := validateJSONOutput(normalized.Output); err != nil {
			return nil, err
		}
	}
	return c.execute(ctx, c.buildLookupHTTPRequest(normalized))
}

func (c *Client) BulkLookupIPGeolocation(ctx context.Context, request *BulkLookupRequest) (*APIResponse[[]BulkLookupResult], error) {
	return lookupBulkTyped(c, ctx, request, parseResponse, makeBulkResult)
}

func makeBulkResult(data *IPGeolocationResponse, err *BulkLookupError) BulkLookupResult {
	return BulkLookupResult{Data: data, Error: err}
}

func (c *Client) BulkLookupIPGeolocationDetailed(ctx context.Context, request *BulkLookupRequest) (*APIResponse[[]DetailedBulkLookupResult], error) {
	return lookupBulkTyped(c, ctx, request, parseDetailedResponse, func(data *DetailedIPGeolocationResponse, err *BulkLookupError) DetailedBulkLookupResult {
		return DetailedBulkLookupResult{Data: data, Error: err}
	})
}

func lookupBulkTyped[T, R any](c *Client, ctx context.Context, request *BulkLookupRequest, parse func(map[string]json.RawMessage) (*T, error), row func(*T, *BulkLookupError) R) (*APIResponse[[]R], error) {
	raw, count, err := c.bulk(ctx, request, true)
	if err != nil {
		return nil, err
	}
	data, err := parseBulkLookup(raw.Data, count, c.config.APIKey, parse, row)
	if err != nil {
		return nil, err
	}
	return &APIResponse[[]R]{Data: data, Metadata: raw.Metadata}, nil
}

func (c *Client) BulkLookupIPGeolocationRaw(ctx context.Context, request *BulkLookupRequest) (*APIResponse[string], error) {
	raw, _, err := c.bulk(ctx, request, false)
	return raw, err
}

func (c *Client) bulk(ctx context.Context, request *BulkLookupRequest, typed bool) (*APIResponse[string], int, error) {
	if err := c.ensureReady(); err != nil {
		return nil, 0, err
	}
	if c.config.APIKey == "" {
		return nil, 0, &ValidationError{Message: "bulk lookup requires apiKey in client config"}
	}
	normalized, err := normalizeBulkLookupRequest(request)
	if err != nil {
		return nil, 0, err
	}
	if typed {
		if err := validateJSONOutput(normalized.Output); err != nil {
			return nil, 0, err
		}
	}
	req, err := c.buildBulkHTTPRequest(normalized)
	if err != nil {
		return nil, 0, err
	}
	raw, err := c.execute(ctx, req)
	return raw, len(normalized.IPs), err
}

func (c *Client) ensureReady() error {
	if atomic.LoadUint32(&c.closed) != 0 {
		return &ValidationError{Message: "client is closed"}
	}
	return nil
}

func (c *Client) buildLookupHTTPRequest(request normalizedLookupRequest) *httpRequestData {
	query := buildQueryString(map[string]string{
		"apiKey":   c.config.APIKey,
		"ip":       request.IP,
		"lang":     request.Lang,
		"include":  strings.Join(request.Include, ","),
		"fields":   strings.Join(request.Fields, ","),
		"excludes": strings.Join(request.Excludes, ","),
		"output":   string(request.Output),
	})

	return &httpRequestData{
		URL:         c.config.BaseURL + "/v3/ipgeo" + query,
		Method:      http.MethodGet,
		Headers:     c.buildHeaders(request.Headers, request.UserAgent, request.Output, false),
		ReadTimeout: c.config.ReadTimeout,
	}
}

func (c *Client) buildBulkHTTPRequest(request normalizedBulkLookupRequest) (*httpRequestData, error) {
	body, err := json.Marshal(map[string][]string{
		"ips": request.IPs,
	})
	if err != nil {
		return nil, &SerializationError{Message: "Failed to serialize bulk lookup request body", Err: err}
	}

	query := buildQueryString(map[string]string{
		"apiKey":   c.config.APIKey,
		"lang":     request.Lang,
		"include":  strings.Join(request.Include, ","),
		"fields":   strings.Join(request.Fields, ","),
		"excludes": strings.Join(request.Excludes, ","),
		"output":   string(request.Output),
	})

	return &httpRequestData{
		URL:         c.config.BaseURL + "/v3/ipgeo-bulk" + query,
		Method:      http.MethodPost,
		Headers:     c.buildHeaders(request.Headers, request.UserAgent, request.Output, true),
		Body:        body,
		ReadTimeout: c.config.ReadTimeout,
	}, nil
}

func (c *Client) buildHeaders(requestHeaders map[string]string, requestUserAgent string, output ResponseFormat, includeContentType bool) http.Header {
	headers := http.Header{}
	for name, value := range requestHeaders {
		headers.Set(name, value)
	}

	headers.Set("User-Agent", resolveUserAgent(requestUserAgent, requestHeaders))
	if output == ResponseFormatXML {
		headers.Set("Accept", "application/xml")
	} else {
		headers.Set("Accept", "application/json")
	}

	if c.config.RequestOrigin != "" {
		headers.Set("Origin", c.config.RequestOrigin)
	}

	if includeContentType {
		headers.Set("Content-Type", "application/json")
	}

	return headers
}

func (c *Client) execute(ctx context.Context, request *httpRequestData) (*APIResponse[string], error) {
	startedAt := time.Now()
	response, err := c.transport.Send(ctx, request)
	if err != nil {
		return nil, safeTransportError(err, c.config.APIKey)
	}
	if !isSuccessStatus(response.StatusCode) {
		return nil, toAPIError(response.StatusCode, response.Body, c.config.APIKey)
	}
	return &APIResponse[string]{Data: response.Body, Metadata: toMetadata(response.StatusCode, int(time.Since(startedAt)/time.Millisecond), response.Headers)}, nil
}

func isSuccessStatus(statusCode int) bool {
	return statusCode >= 200 && statusCode <= 299
}
