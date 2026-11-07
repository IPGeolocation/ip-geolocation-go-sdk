//go:build !go1.20

package ipgeolocation

func safeTLSVerificationError(error, string) (error, bool) { return nil, false }
