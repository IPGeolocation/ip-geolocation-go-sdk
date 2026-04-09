package ipgeolocation

import (
	"os"
	"strings"
	"testing"
)

func requireLiveTests(t *testing.T) (string, string) {
	t.Helper()

	if !strings.EqualFold(os.Getenv("IPGEO_RUN_LIVE_TESTS"), "true") {
		t.Skip("set IPGEO_RUN_LIVE_TESTS=true to run live Go tests")
	}

	freeKey := strings.TrimSpace(os.Getenv("IPGEO_FREE_KEY"))
	paidKey := strings.TrimSpace(os.Getenv("IPGEO_PAID_KEY"))
	if freeKey == "" || paidKey == "" {
		t.Skip("set IPGEO_FREE_KEY and IPGEO_PAID_KEY to run live Go tests")
	}

	return freeKey, paidKey
}

func requireLiveHardening(t *testing.T) string {
	t.Helper()

	if !strings.EqualFold(os.Getenv("IPGEO_RUN_LIVE_HARDENING"), "true") {
		t.Skip("set IPGEO_RUN_LIVE_HARDENING=true to run live Go parity tests")
	}

	paidKey := strings.TrimSpace(os.Getenv("IPGEO_PAID_KEY"))
	if paidKey == "" {
		t.Skip("set IPGEO_PAID_KEY to run live Go parity tests")
	}

	return paidKey
}
