package ipgeolocation

import (
	"context"
	"net/http"
)

type stubTransport struct {
	send func(context.Context, *httpRequestData) (*httpResponseData, error)
}

func (s stubTransport) Send(ctx context.Context, request *httpRequestData) (*httpResponseData, error) {
	return s.send(ctx, request)
}

func (s stubTransport) Close() {}

func mustNewTestClient(config *Config, transport transport) *Client {
	client, err := newClientWithTransport(config, transport)
	if err != nil {
		panic(err)
	}

	return client
}

func header(key, value string) http.Header {
	return http.Header{key: []string{value}}
}
