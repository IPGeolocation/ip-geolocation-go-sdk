package ipgeolocation

import "strings"

type ResponseFormat string

const (
	ResponseFormatJSON ResponseFormat = "json"
	ResponseFormatXML  ResponseFormat = "xml"
)

func normalizeResponseFormat(value ResponseFormat) (ResponseFormat, error) {
	if value == "" {
		return ResponseFormatJSON, nil
	}

	switch strings.ToLower(string(value)) {
	case "json":
		return ResponseFormatJSON, nil
	case "xml":
		return ResponseFormatXML, nil
	default:
		return "", &ValidationError{Message: "output must be json or xml"}
	}
}
