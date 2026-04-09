package ipgeolocation

import "net/http"

type APIResponse[T any] struct {
	Data     T
	Metadata ResponseMetadata
}

type ResponseMetadata struct {
	CreditsCharged    *int
	SuccessfulRecords *int
	StatusCode        int
	DurationMS        int
	RawHeaders        http.Header
}

func (m ResponseMetadata) HeaderValues(name string) []string {
	if name == "" {
		return nil
	}

	values := m.RawHeaders.Values(name)
	if len(values) == 0 {
		return nil
	}

	cloned := make([]string, len(values))
	copy(cloned, values)
	return cloned
}

func (m ResponseMetadata) FirstHeaderValue(name string) string {
	values := m.HeaderValues(name)
	if len(values) == 0 {
		return ""
	}

	return values[0]
}
