package ipgeolocation

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
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
	config    normalizedConfig
	transport transport
	closed    bool
}

func NewClient(config *Config) (*Client, error) {
	normalized, err := normalizeConfig(config)
	if err != nil {
		return nil, err
	}

	return &Client{
		config:    normalized,
		transport: newHTTPTransport(normalized.ConnectTimeout, DefaultMaxResponseBodyBytes),
	}, nil
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
		config:    normalized,
		transport: transport,
	}, nil
}

func (c *Client) Close() {
	if c.closed {
		return
	}

	c.closed = true
	c.transport.Close()
}

func (c *Client) LookupIPGeolocation(ctx context.Context, request *LookupRequest) (*APIResponse[*IPGeolocationResponse], error) {
	if err := c.ensureReady(); err != nil {
		return nil, err
	}

	if c.config.APIKey == "" && c.config.RequestOrigin == "" {
		return nil, &ValidationError{Message: "single lookup requires apiKey or requestOrigin in client config"}
	}

	normalizedRequest, err := normalizeLookupRequest(request)
	if err != nil {
		return nil, err
	}

	if err := validateJSONOutput(normalizedRequest.Output); err != nil {
		return nil, err
	}

	response, durationMS, err := c.execute(ctx, c.buildLookupHTTPRequest(normalizedRequest))
	if err != nil {
		return nil, err
	}

	if !isSuccessStatus(response.StatusCode) {
		return nil, toAPIError(response.StatusCode, response.Body)
	}

	data, err := parseSingleLookup(response.Body)
	if err != nil {
		return nil, err
	}

	return &APIResponse[*IPGeolocationResponse]{
		Data:     data,
		Metadata: toMetadata(response.StatusCode, durationMS, response.Headers),
	}, nil
}

func (c *Client) LookupIPGeolocationRaw(ctx context.Context, request *LookupRequest) (*APIResponse[string], error) {
	if err := c.ensureReady(); err != nil {
		return nil, err
	}

	if c.config.APIKey == "" && c.config.RequestOrigin == "" {
		return nil, &ValidationError{Message: "single lookup requires apiKey or requestOrigin in client config"}
	}

	normalizedRequest, err := normalizeLookupRequest(request)
	if err != nil {
		return nil, err
	}

	response, durationMS, err := c.execute(ctx, c.buildLookupHTTPRequest(normalizedRequest))
	if err != nil {
		return nil, err
	}

	if !isSuccessStatus(response.StatusCode) {
		return nil, toAPIError(response.StatusCode, response.Body)
	}

	return &APIResponse[string]{
		Data:     response.Body,
		Metadata: toMetadata(response.StatusCode, durationMS, response.Headers),
	}, nil
}

func (c *Client) BulkLookupIPGeolocation(ctx context.Context, request *BulkLookupRequest) (*APIResponse[[]BulkLookupResult], error) {
	if err := c.ensureReady(); err != nil {
		return nil, err
	}

	if c.config.APIKey == "" {
		return nil, &ValidationError{Message: "bulk lookup requires apiKey in client config"}
	}

	normalizedRequest, err := normalizeBulkLookupRequest(request)
	if err != nil {
		return nil, err
	}

	if err := validateJSONOutput(normalizedRequest.Output); err != nil {
		return nil, err
	}

	bulkRequestData, err := c.buildBulkHTTPRequest(normalizedRequest)
	if err != nil {
		return nil, err
	}

	response, durationMS, err := c.execute(ctx, bulkRequestData)
	if err != nil {
		return nil, err
	}

	if !isSuccessStatus(response.StatusCode) {
		return nil, toAPIError(response.StatusCode, response.Body)
	}

	data, err := parseBulkLookup(response.Body)
	if err != nil {
		return nil, err
	}

	return &APIResponse[[]BulkLookupResult]{
		Data:     data,
		Metadata: toMetadata(response.StatusCode, durationMS, response.Headers),
	}, nil
}

func (c *Client) BulkLookupIPGeolocationRaw(ctx context.Context, request *BulkLookupRequest) (*APIResponse[string], error) {
	if err := c.ensureReady(); err != nil {
		return nil, err
	}

	if c.config.APIKey == "" {
		return nil, &ValidationError{Message: "bulk lookup requires apiKey in client config"}
	}

	normalizedRequest, err := normalizeBulkLookupRequest(request)
	if err != nil {
		return nil, err
	}

	bulkRequestData, err := c.buildBulkHTTPRequest(normalizedRequest)
	if err != nil {
		return nil, err
	}

	response, durationMS, err := c.execute(ctx, bulkRequestData)
	if err != nil {
		return nil, err
	}

	if !isSuccessStatus(response.StatusCode) {
		return nil, toAPIError(response.StatusCode, response.Body)
	}

	return &APIResponse[string]{
		Data:     response.Body,
		Metadata: toMetadata(response.StatusCode, durationMS, response.Headers),
	}, nil
}

func (c *Client) ensureReady() error {
	if c.closed {
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

func (c *Client) execute(ctx context.Context, request *httpRequestData) (*httpResponseData, int, error) {
	startedAt := time.Now()
	response, err := c.transport.Send(ctx, request)
	if err != nil {
		return nil, 0, err
	}

	durationMS := int(time.Since(startedAt) / time.Millisecond)
	return response, durationMS, nil
}

func isSuccessStatus(statusCode int) bool {
	return statusCode >= 200 && statusCode <= 299
}
