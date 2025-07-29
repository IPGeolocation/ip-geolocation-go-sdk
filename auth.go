package ipgeolocationsdk

import (
	"context"
)

// WithAPIKey sets the API key in the context. It will be used for authentication
// of all requests made with the returned context.
func WithAPIKey(ctx context.Context, key string) context.Context {
	return context.WithValue(ctx, ContextAPIKeys, map[string]APIKey{
		"ApiKeyAuth": {
			Key: key,
		},
	})
}
