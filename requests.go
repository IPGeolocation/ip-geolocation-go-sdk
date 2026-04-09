package ipgeolocation

import (
	"net/textproto"
	"strings"
)

type LookupRequest struct {
	IP        string
	Lang      string
	Include   []string
	Fields    []string
	Excludes  []string
	UserAgent string
	Headers   map[string]string
	Output    ResponseFormat
}

type BulkLookupRequest struct {
	IPs       []string
	Lang      string
	Include   []string
	Fields    []string
	Excludes  []string
	UserAgent string
	Headers   map[string]string
	Output    ResponseFormat
}

type normalizedLookupRequest struct {
	IP        string
	Lang      string
	Include   []string
	Fields    []string
	Excludes  []string
	UserAgent string
	Headers   map[string]string
	Output    ResponseFormat
}

type normalizedBulkLookupRequest struct {
	IPs       []string
	Lang      string
	Include   []string
	Fields    []string
	Excludes  []string
	UserAgent string
	Headers   map[string]string
	Output    ResponseFormat
}

func (r *LookupRequest) Validate() error {
	_, err := normalizeLookupRequest(r)
	return err
}

func (r *BulkLookupRequest) Validate() error {
	_, err := normalizeBulkLookupRequest(r)
	return err
}

func normalizeLookupRequest(request *LookupRequest) (normalizedLookupRequest, error) {
	if request == nil {
		return normalizedLookupRequest{
			Headers: map[string]string{},
			Output:  ResponseFormatJSON,
		}, nil
	}

	output, err := normalizeResponseFormat(request.Output)
	if err != nil {
		return normalizedLookupRequest{}, err
	}

	ip, err := normalizeLookupIP(request.IP)
	if err != nil {
		return normalizedLookupRequest{}, err
	}

	lang, err := normalizeLanguage(request.Lang)
	if err != nil {
		return normalizedLookupRequest{}, err
	}

	include, err := normalizeStringList(request.Include, "include")
	if err != nil {
		return normalizedLookupRequest{}, err
	}

	fields, err := normalizeStringList(request.Fields, "fields")
	if err != nil {
		return normalizedLookupRequest{}, err
	}

	excludes, err := normalizeStringList(request.Excludes, "excludes")
	if err != nil {
		return normalizedLookupRequest{}, err
	}

	userAgent, err := normalizeOptionalHeaderValue(request.UserAgent, "userAgent")
	if err != nil {
		return normalizedLookupRequest{}, err
	}

	headers, err := normalizeHeaders(request.Headers)
	if err != nil {
		return normalizedLookupRequest{}, err
	}

	return normalizedLookupRequest{
		IP:        ip,
		Lang:      lang,
		Include:   include,
		Fields:    fields,
		Excludes:  excludes,
		UserAgent: userAgent,
		Headers:   headers,
		Output:    output,
	}, nil
}

func normalizeBulkLookupRequest(request *BulkLookupRequest) (normalizedBulkLookupRequest, error) {
	if request == nil {
		return normalizedBulkLookupRequest{}, &ValidationError{Message: "bulk lookup request is required"}
	}

	output, err := normalizeResponseFormat(request.Output)
	if err != nil {
		return normalizedBulkLookupRequest{}, err
	}

	ips, err := normalizeIPList(request.IPs)
	if err != nil {
		return normalizedBulkLookupRequest{}, err
	}

	lang, err := normalizeLanguage(request.Lang)
	if err != nil {
		return normalizedBulkLookupRequest{}, err
	}

	include, err := normalizeStringList(request.Include, "include")
	if err != nil {
		return normalizedBulkLookupRequest{}, err
	}

	fields, err := normalizeStringList(request.Fields, "fields")
	if err != nil {
		return normalizedBulkLookupRequest{}, err
	}

	excludes, err := normalizeStringList(request.Excludes, "excludes")
	if err != nil {
		return normalizedBulkLookupRequest{}, err
	}

	userAgent, err := normalizeOptionalHeaderValue(request.UserAgent, "userAgent")
	if err != nil {
		return normalizedBulkLookupRequest{}, err
	}

	headers, err := normalizeHeaders(request.Headers)
	if err != nil {
		return normalizedBulkLookupRequest{}, err
	}

	return normalizedBulkLookupRequest{
		IPs:       ips,
		Lang:      lang,
		Include:   include,
		Fields:    fields,
		Excludes:  excludes,
		UserAgent: userAgent,
		Headers:   headers,
		Output:    output,
	}, nil
}

func normalizeLookupIP(value string) (string, error) {
	normalized := strings.TrimSpace(value)
	if value != "" && normalized == "" {
		return "", &ValidationError{Message: "ip must not be blank"}
	}

	if normalized == "" {
		return "", nil
	}

	if strings.ContainsAny(normalized, "\r\n") {
		return "", &ValidationError{Message: "ip must not contain CR or LF"}
	}

	return normalized, nil
}

func normalizeIPList(values []string) ([]string, error) {
	if values == nil {
		return nil, &ValidationError{Message: "ips must contain at least one IP address or domain"}
	}

	if len(values) == 0 {
		return nil, &ValidationError{Message: "ips must contain at least one IP address or domain"}
	}

	if len(values) > 50000 {
		return nil, &ValidationError{Message: "ips must not contain more than 50000 entries"}
	}

	result := make([]string, 0, len(values))
	for _, value := range values {
		normalized, err := normalizeLookupIP(value)
		if err != nil {
			return nil, err
		}

		if normalized == "" {
			return nil, &ValidationError{Message: "ips must not contain blank values"}
		}

		result = append(result, normalized)
	}

	return result, nil
}

func normalizeLanguage(value string) (string, error) {
	normalized := strings.TrimSpace(value)
	if value != "" && normalized == "" {
		return "", &ValidationError{Message: "lang must not be blank"}
	}

	if normalized == "" {
		return "", nil
	}

	switch strings.ToLower(normalized) {
	case "en", "de", "ru", "ja", "fr", "cn", "es", "cs", "it", "ko", "fa", "pt":
		return strings.ToLower(normalized), nil
	default:
		return "", &ValidationError{Message: "lang must be one of: en, de, ru, ja, fr, cn, es, cs, it, ko, fa, pt"}
	}
}

func normalizeStringList(values []string, fieldName string) ([]string, error) {
	if len(values) == 0 {
		return nil, nil
	}

	seen := map[string]struct{}{}
	result := make([]string, 0, len(values))
	for _, value := range values {
		normalized := strings.TrimSpace(value)
		if normalized == "" {
			return nil, &ValidationError{Message: fieldName + " must not contain blank values"}
		}

		if _, ok := seen[normalized]; ok {
			continue
		}

		seen[normalized] = struct{}{}
		result = append(result, normalized)
	}

	return result, nil
}

func normalizeOptionalHeaderValue(value string, fieldName string) (string, error) {
	normalized := strings.TrimSpace(value)
	if value != "" && normalized == "" {
		return "", &ValidationError{Message: fieldName + " must not be blank"}
	}

	if normalized == "" {
		return "", nil
	}

	if strings.ContainsAny(normalized, "\r\n") {
		return "", &ValidationError{Message: fieldName + " must not contain CR or LF"}
	}

	return normalized, nil
}

func normalizeHeaders(headers map[string]string) (map[string]string, error) {
	if len(headers) == 0 {
		return map[string]string{}, nil
	}

	result := make(map[string]string, len(headers))
	for key, value := range headers {
		name := textproto.CanonicalMIMEHeaderKey(strings.TrimSpace(key))
		if name == "" {
			return nil, &ValidationError{Message: "headers must not contain blank names"}
		}

		normalizedValue := strings.TrimSpace(value)
		if normalizedValue == "" {
			return nil, &ValidationError{Message: "headers must not contain blank values"}
		}

		if strings.ContainsAny(name, "\r\n") || strings.ContainsAny(normalizedValue, "\r\n") {
			return nil, &ValidationError{Message: "headers must not contain CR or LF"}
		}

		result[name] = normalizedValue
	}

	return result, nil
}
