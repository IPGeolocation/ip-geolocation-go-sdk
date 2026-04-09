package ipgeolocation

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func validateJSONOutput(output ResponseFormat) error {
	if output != ResponseFormatJSON {
		return &ValidationError{Message: "typed methods support JSON only"}
	}

	return nil
}

func buildQueryString(values map[string]string) string {
	query := url.Values{}
	for key, value := range values {
		if value == "" {
			continue
		}

		query.Set(key, value)
	}

	encoded := query.Encode()
	if encoded == "" {
		return ""
	}

	return "?" + encoded
}

func resolveUserAgent(requestUserAgent string, headers map[string]string) string {
	if requestUserAgent != "" {
		return requestUserAgent
	}

	if value := headers["User-Agent"]; strings.TrimSpace(value) != "" {
		return value
	}

	return DefaultUserAgent()
}

func toMetadata(statusCode int, durationMS int, headers http.Header) ResponseMetadata {
	metadata := ResponseMetadata{
		StatusCode: statusCode,
		DurationMS: durationMS,
		RawHeaders: cloneHeader(headers),
	}

	if value := headers.Get("X-Credits-Charged"); value != "" {
		if parsed, err := strconv.Atoi(value); err == nil {
			metadata.CreditsCharged = intPtr(parsed)
		}
	}

	if value := headers.Get("X-Successful-Records"); value != "" {
		if parsed, err := strconv.Atoi(value); err == nil {
			metadata.SuccessfulRecords = intPtr(parsed)
		}
	}

	return metadata
}

func toAPIError(statusCode int, body string) error {
	message := extractAPIMessage(body)
	if message == "" {
		message = apiErrorMessage(statusCode)
	}

	base := APIError{
		StatusCode: statusCode,
		Message:    message,
	}

	switch statusCode {
	case 400:
		return &BadRequestError{APIError: base}
	case 401:
		return &UnauthorizedError{APIError: base}
	case 403:
		return &ForbiddenError{APIError: base}
	case 404:
		return &NotFoundError{APIError: base}
	case 405:
		return &MethodNotAllowedError{APIError: base}
	case 413:
		return &ContentTooLargeError{APIError: base}
	case 415:
		return &UnsupportedMediaTypeError{APIError: base}
	case 423:
		return &LockedError{APIError: base}
	case 429:
		return &TooManyRequestsError{APIError: base}
	case 499:
		return &CustomStatus499Error{APIError: base}
	case 500:
		return &InternalServerError{APIError: base}
	default:
		return &base
	}
}

func extractAPIMessage(body string) string {
	if strings.TrimSpace(body) == "" {
		return ""
	}

	var object map[string]json.RawMessage
	if err := json.Unmarshal([]byte(body), &object); err != nil {
		return strings.TrimSpace(body)
	}

	if message := readString(object, "message"); message != nil {
		return *message
	}

	if nested, ok := object["error"]; ok {
		errorObject, err := decodeObject(nested)
		if err == nil {
			if message := readString(errorObject, "message"); message != nil {
				return *message
			}
		}
	}

	return strings.TrimSpace(body)
}

func parseSingleLookup(body string) (*IPGeolocationResponse, error) {
	object, err := decodeObject([]byte(body))
	if err != nil {
		return nil, &SerializationError{Message: "Failed to parse single lookup response body", Err: err}
	}

	response, err := parseResponse(object)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func parseBulkLookup(body string) ([]BulkLookupResult, error) {
	var items []json.RawMessage
	if err := json.Unmarshal([]byte(body), &items); err != nil {
		return nil, &SerializationError{Message: "Failed to parse bulk lookup response body", Err: err}
	}

	results := make([]BulkLookupResult, 0, len(items))
	for _, item := range items {
		object, err := decodeObject(item)
		if err != nil {
			return nil, &SerializationError{Message: "Failed to parse bulk lookup response body", Err: err}
		}

		if isBulkError(object) {
			results = append(results, BulkLookupResult{
				Error: &BulkLookupError{Message: readBulkErrorMessage(object)},
			})
			continue
		}

		response, err := parseResponse(object)
		if err != nil {
			return nil, err
		}

		results = append(results, BulkLookupResult{Data: response})
	}

	return results, nil
}

func isBulkError(object map[string]json.RawMessage) bool {
	_, hasError := object["error"]
	_, hasMessage := object["message"]
	_, hasIP := object["ip"]
	return hasError || (hasMessage && !hasIP)
}

func readBulkErrorMessage(object map[string]json.RawMessage) string {
	if message := readString(object, "message"); message != nil {
		return *message
	}

	if nested, ok := object["error"]; ok {
		errorObject, err := decodeObject(nested)
		if err == nil {
			if message := readString(errorObject, "message"); message != nil {
				return *message
			}
		}
	}

	return "Bulk lookup item failed"
}

func parseResponse(object map[string]json.RawMessage) (*IPGeolocationResponse, error) {
	location, err := readObject(object, "location", parseLocation)
	if err != nil {
		return nil, err
	}

	countryMetadata, err := readObject(object, "country_metadata", parseCountryMetadata)
	if err != nil {
		return nil, err
	}

	currency, err := readObject(object, "currency", parseCurrency)
	if err != nil {
		return nil, err
	}

	network, err := readObject(object, "network", parseNetwork)
	if err != nil {
		return nil, err
	}

	asn, err := readObject(object, "asn", parseASN)
	if err != nil {
		return nil, err
	}

	company, err := readObject(object, "company", parseCompany)
	if err != nil {
		return nil, err
	}

	timeZone, err := readObject(object, "time_zone", parseTimeZone)
	if err != nil {
		return nil, err
	}

	security, err := readObject(object, "security", parseSecurity)
	if err != nil {
		return nil, err
	}

	userAgent, err := readObject(object, "user_agent", parseUserAgent)
	if err != nil {
		return nil, err
	}

	abuse, err := readObject(object, "abuse", parseAbuse)
	if err != nil {
		return nil, err
	}

	return &IPGeolocationResponse{
		IP:              readString(object, "ip"),
		Hostname:        readString(object, "hostname"),
		Domain:          readString(object, "domain"),
		Location:        location,
		CountryMetadata: countryMetadata,
		Currency:        currency,
		Network:         network,
		ASN:             asn,
		Company:         company,
		TimeZone:        timeZone,
		Security:        security,
		UserAgent:       userAgent,
		Abuse:           abuse,
	}, nil
}

func parseLocation(object map[string]json.RawMessage) (Location, error) {
	return Location{
		ContinentCode:       readString(object, "continent_code"),
		ContinentName:       readString(object, "continent_name"),
		CountryCode2:        readString(object, "country_code2"),
		CountryCode3:        readString(object, "country_code3"),
		CountryName:         readString(object, "country_name"),
		CountryNameOfficial: readString(object, "country_name_official"),
		CountryCapital:      readString(object, "country_capital"),
		StateProv:           readString(object, "state_prov"),
		StateCode:           readString(object, "state_code"),
		District:            readString(object, "district"),
		City:                readString(object, "city"),
		Locality:            readString(object, "locality"),
		AccuracyRadius:      readString(object, "accuracy_radius"),
		Confidence:          readString(object, "confidence"),
		DMACode:             readString(object, "dma_code"),
		Zipcode:             readString(object, "zipcode"),
		Latitude:            readString(object, "latitude"),
		Longitude:           readString(object, "longitude"),
		IsEU:                readBool(object, "is_eu"),
		CountryFlag:         readString(object, "country_flag"),
		GeonameID:           readString(object, "geoname_id"),
		CountryEmoji:        readString(object, "country_emoji"),
	}, nil
}

func parseCountryMetadata(object map[string]json.RawMessage) (CountryMetadata, error) {
	return CountryMetadata{
		CallingCode: readString(object, "calling_code"),
		TLD:         readString(object, "tld"),
		Languages:   readStringArray(object, "languages"),
	}, nil
}

func parseCurrency(object map[string]json.RawMessage) (Currency, error) {
	return Currency{
		Code:   readString(object, "code"),
		Name:   readString(object, "name"),
		Symbol: readString(object, "symbol"),
	}, nil
}

func parseNetwork(object map[string]json.RawMessage) (Network, error) {
	return Network{
		ConnectionType: readString(object, "connection_type"),
		Route:          readString(object, "route"),
		IsAnycast:      readBool(object, "is_anycast"),
	}, nil
}

func parseASN(object map[string]json.RawMessage) (ASNInfo, error) {
	return ASNInfo{
		ASNumber:      readString(object, "as_number"),
		Organization:  readString(object, "organization"),
		Country:       readString(object, "country"),
		Type:          readString(object, "type"),
		Domain:        readString(object, "domain"),
		DateAllocated: readString(object, "date_allocated"),
		RIR:           readString(object, "rir"),
	}, nil
}

func parseCompany(object map[string]json.RawMessage) (Company, error) {
	return Company{
		Name:   readString(object, "name"),
		Type:   readString(object, "type"),
		Domain: readString(object, "domain"),
	}, nil
}

func parseTimeZone(object map[string]json.RawMessage) (TimeZone, error) {
	dstStart, err := readObject(object, "dst_start", parseDSTTransition)
	if err != nil {
		return TimeZone{}, err
	}

	dstEnd, err := readObject(object, "dst_end", parseDSTTransition)
	if err != nil {
		return TimeZone{}, err
	}

	return TimeZone{
		Name:                         readString(object, "name"),
		Offset:                       readFloat(object, "offset"),
		OffsetWithDST:                readFloat(object, "offset_with_dst"),
		CurrentTime:                  readString(object, "current_time"),
		CurrentTimeUnix:              readFloat(object, "current_time_unix"),
		CurrentTimeZoneAbbreviation:  readFirstString(object, "current_tz_abbreviation", "current_timezone_abbreviation"),
		CurrentTimeZoneFullName:      readFirstString(object, "current_tz_full_name", "current_timezone_name"),
		StandardTimeZoneAbbreviation: readFirstString(object, "standard_tz_abbreviation", "timezone_abbreviation"),
		StandardTimeZoneFullName:     readFirstString(object, "standard_tz_full_name", "timezone_name"),
		IsDST:                        readBool(object, "is_dst"),
		DSTSavings:                   readFloat(object, "dst_savings"),
		DSTExists:                    readBool(object, "dst_exists"),
		DSTTimeZoneAbbreviation:      readFirstString(object, "dst_tz_abbreviation", "dst_timezone_abbreviation"),
		DSTTimeZoneFullName:          readFirstString(object, "dst_tz_full_name", "dst_timezone_name"),
		DSTStart:                     dstStart,
		DSTEnd:                       dstEnd,
	}, nil
}

func parseDSTTransition(object map[string]json.RawMessage) (DSTTransition, error) {
	return DSTTransition{
		UTCTime:        readString(object, "utc_time"),
		Duration:       readString(object, "duration"),
		Gap:            readBool(object, "gap"),
		DateTimeAfter:  readString(object, "date_time_after"),
		DateTimeBefore: readString(object, "date_time_before"),
		Overlap:        readBool(object, "overlap"),
	}, nil
}

func parseSecurity(object map[string]json.RawMessage) (Security, error) {
	return Security{
		ThreatScore:          readFloat(object, "threat_score"),
		IsTor:                readBool(object, "is_tor"),
		IsProxy:              readBool(object, "is_proxy"),
		ProxyProviderNames:   readStringArray(object, "proxy_provider_names"),
		ProxyConfidenceScore: readFloat(object, "proxy_confidence_score"),
		ProxyLastSeen:        readString(object, "proxy_last_seen"),
		IsResidentialProxy:   readBool(object, "is_residential_proxy"),
		IsVPN:                readBool(object, "is_vpn"),
		VPNProviderNames:     readStringArray(object, "vpn_provider_names"),
		VPNConfidenceScore:   readFloat(object, "vpn_confidence_score"),
		VPNLastSeen:          readString(object, "vpn_last_seen"),
		IsRelay:              readBool(object, "is_relay"),
		RelayProviderName:    readString(object, "relay_provider_name"),
		IsAnonymous:          readBool(object, "is_anonymous"),
		IsKnownAttacker:      readBool(object, "is_known_attacker"),
		IsBot:                readBool(object, "is_bot"),
		IsSpam:               readBool(object, "is_spam"),
		IsCloudProvider:      readBool(object, "is_cloud_provider"),
		CloudProviderName:    readString(object, "cloud_provider_name"),
	}, nil
}

func parseAbuse(object map[string]json.RawMessage) (Abuse, error) {
	return Abuse{
		Route:        readString(object, "route"),
		Country:      readString(object, "country"),
		Name:         readString(object, "name"),
		Organization: readString(object, "organization"),
		Kind:         readString(object, "kind"),
		Address:      readString(object, "address"),
		Emails:       readStringArray(object, "emails"),
		PhoneNumbers: readStringArray(object, "phone_numbers"),
	}, nil
}

func parseUserAgent(object map[string]json.RawMessage) (UserAgent, error) {
	device, err := readObject(object, "device", parseUserAgentDevice)
	if err != nil {
		return UserAgent{}, err
	}

	engine, err := readObject(object, "engine", parseUserAgentEngine)
	if err != nil {
		return UserAgent{}, err
	}

	operatingSystem, err := readObject(object, "operating_system", parseUserAgentOperatingSystem)
	if err != nil {
		return UserAgent{}, err
	}

	return UserAgent{
		Name:            readString(object, "name"),
		Type:            readString(object, "type"),
		Version:         readString(object, "version"),
		VersionMajor:    readString(object, "version_major"),
		Device:          device,
		Engine:          engine,
		OperatingSystem: operatingSystem,
	}, nil
}

func parseUserAgentDevice(object map[string]json.RawMessage) (UserAgentDevice, error) {
	return UserAgentDevice{
		Name:  readString(object, "name"),
		Type:  readString(object, "type"),
		Brand: readString(object, "brand"),
		CPU:   readString(object, "cpu"),
	}, nil
}

func parseUserAgentEngine(object map[string]json.RawMessage) (UserAgentEngine, error) {
	return UserAgentEngine{
		Name:         readString(object, "name"),
		Type:         readString(object, "type"),
		Version:      readString(object, "version"),
		VersionMajor: readString(object, "version_major"),
	}, nil
}

func parseUserAgentOperatingSystem(object map[string]json.RawMessage) (UserAgentOperatingSystem, error) {
	return UserAgentOperatingSystem{
		Name:    readString(object, "name"),
		Version: readString(object, "version"),
	}, nil
}

func decodeObject(data []byte) (map[string]json.RawMessage, error) {
	var object map[string]json.RawMessage
	if err := json.Unmarshal(data, &object); err != nil {
		return nil, err
	}

	if object == nil {
		return nil, fmt.Errorf("expected JSON object")
	}

	return object, nil
}

func readString(object map[string]json.RawMessage, key string) *string {
	raw, ok := object[key]
	if !ok || isNull(raw) {
		return nil
	}

	var value string
	if err := json.Unmarshal(raw, &value); err == nil {
		return &value
	}

	var number json.Number
	if err := json.Unmarshal(raw, &number); err == nil {
		value = number.String()
		return &value
	}

	var boolean bool
	if err := json.Unmarshal(raw, &boolean); err == nil {
		if boolean {
			value = "true"
		} else {
			value = "false"
		}
		return &value
	}

	return nil
}

func readFirstString(object map[string]json.RawMessage, keys ...string) *string {
	for _, key := range keys {
		if value := readString(object, key); value != nil {
			return value
		}
	}

	return nil
}

func readBool(object map[string]json.RawMessage, key string) *bool {
	raw, ok := object[key]
	if !ok || isNull(raw) {
		return nil
	}

	var value bool
	if err := json.Unmarshal(raw, &value); err == nil {
		return &value
	}

	var stringValue string
	if err := json.Unmarshal(raw, &stringValue); err == nil {
		parsed, err := strconv.ParseBool(strings.ToLower(stringValue))
		if err == nil {
			return &parsed
		}
	}

	return nil
}

func readFloat(object map[string]json.RawMessage, key string) *float64 {
	raw, ok := object[key]
	if !ok || isNull(raw) {
		return nil
	}

	var value float64
	if err := json.Unmarshal(raw, &value); err == nil {
		return &value
	}

	var stringValue string
	if err := json.Unmarshal(raw, &stringValue); err == nil {
		parsed, err := strconv.ParseFloat(stringValue, 64)
		if err == nil {
			return &parsed
		}
	}

	return nil
}

func readStringArray(object map[string]json.RawMessage, key string) []string {
	raw, ok := object[key]
	if !ok || isNull(raw) {
		return nil
	}

	var values []json.RawMessage
	if err := json.Unmarshal(raw, &values); err != nil {
		return nil
	}

	result := make([]string, 0, len(values))
	for _, value := range values {
		if parsed := readString(map[string]json.RawMessage{"value": value}, "value"); parsed != nil && *parsed != "" {
			result = append(result, *parsed)
		}
	}

	return result
}

func readObject[T any](object map[string]json.RawMessage, key string, parser func(map[string]json.RawMessage) (T, error)) (*T, error) {
	raw, ok := object[key]
	if !ok || isNull(raw) {
		return nil, nil
	}

	if string(raw) == "[]" {
		return nil, nil
	}

	decoded, err := decodeObject(raw)
	if err != nil {
		return nil, &SerializationError{Message: "Failed to parse " + key + " object", Err: err}
	}

	value, err := parser(decoded)
	if err != nil {
		return nil, err
	}

	return &value, nil
}

func isNull(raw json.RawMessage) bool {
	return len(raw) == 0 || strings.TrimSpace(string(raw)) == "null"
}

func intPtr(value int) *int {
	return &value
}
