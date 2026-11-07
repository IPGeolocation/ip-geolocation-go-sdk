//go:build go1.23

package ipgeolocation

import (
	"context"
	"errors"
	"net"
	"testing"
)

func TestDNSErrorKeepsCancellation(t *testing.T) {
	for _, cause := range []error{context.Canceled, context.DeadlineExceeded, errors.New("custom example-secret diagnostic")} {
		dns := &net.DNSError{Err: "lookup failed", Name: "example-secret", UnwrapErr: cause}
		got := safeCause(dns, "example-secret")
		var sanitized *net.DNSError
		if !errors.As(got, &sanitized) || sanitized.Name != "[REDACTED]" {
			t.Fatal("lost DNS error details")
		}
		if cause == context.Canceled || cause == context.DeadlineExceeded {
			if !errors.Is(got, cause) {
				t.Fatal("lost DNS cancellation or deadline")
			}
		} else if errors.Unwrap(sanitized) != nil {
			t.Fatal("retained custom DNS error details")
		}
	}
}
