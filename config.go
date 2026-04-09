package ipgeolocation

import (
	"encoding/json"
	"net/url"
	"strings"
	"time"
)

const (
	DefaultBaseURL              = "https://api.ipgeolocation.io"
	DefaultMaxResponseBodyBytes = 134217728
)

var (
	DefaultConnectTimeout = 10 * time.Second
	DefaultReadTimeout    = 30 * time.Second
)

type Config struct {
	APIKey         string        `json:"-"`
	RequestOrigin  string        `json:"requestOrigin,omitempty"`
	BaseURL        string        `json:"baseUrl,omitempty"`
	ConnectTimeout time.Duration `json:"connectTimeout,omitempty"`
	ReadTimeout    time.Duration `json:"readTimeout,omitempty"`
}

type normalizedConfig struct {
	APIKey         string
	RequestOrigin  string
	BaseURL        string
	ConnectTimeout time.Duration
	ReadTimeout    time.Duration
}

func DefaultConfig() *Config {
	return &Config{
		BaseURL:        DefaultBaseURL,
		ConnectTimeout: DefaultConnectTimeout,
		ReadTimeout:    DefaultReadTimeout,
	}
}

func (c *Config) Validate() error {
	_, err := normalizeConfig(c)
	return err
}

func (c *Config) SafeMap() map[string]interface{} {
	cfg := c
	if cfg == nil {
		cfg = DefaultConfig()
	}

	result := map[string]interface{}{
		"apiKey":         nil,
		"requestOrigin":  strings.TrimSpace(cfg.RequestOrigin),
		"baseUrl":        strings.TrimSpace(cfg.BaseURL),
		"connectTimeout": cfg.ConnectTimeout.Seconds(),
		"readTimeout":    cfg.ReadTimeout.Seconds(),
	}

	if strings.TrimSpace(cfg.APIKey) != "" {
		result["apiKey"] = "[REDACTED]"
	}

	if normalized, err := normalizeConfig(cfg); err == nil {
		result["requestOrigin"] = emptyStringAsNil(normalized.RequestOrigin)
		result["baseUrl"] = normalized.BaseURL
		result["connectTimeout"] = normalized.ConnectTimeout.Seconds()
		result["readTimeout"] = normalized.ReadTimeout.Seconds()
	}

	return result
}

func (c *Config) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.SafeMap())
}

func normalizeConfig(config *Config) (normalizedConfig, error) {
	if config == nil {
		config = DefaultConfig()
	}

	baseURL := config.BaseURL
	if strings.TrimSpace(baseURL) == "" {
		baseURL = DefaultBaseURL
	}

	connectTimeout := config.ConnectTimeout
	if connectTimeout == 0 {
		connectTimeout = DefaultConnectTimeout
	}

	readTimeout := config.ReadTimeout
	if readTimeout == 0 {
		readTimeout = DefaultReadTimeout
	}

	apiKey, err := normalizeAPIKey(config.APIKey)
	if err != nil {
		return normalizedConfig{}, err
	}

	requestOrigin, err := normalizeRequestOrigin(config.RequestOrigin)
	if err != nil {
		return normalizedConfig{}, err
	}

	normalizedBaseURL, err := normalizeBaseURL(baseURL)
	if err != nil {
		return normalizedConfig{}, err
	}

	normalizedConnectTimeout, err := normalizeTimeout(connectTimeout, "connectTimeout")
	if err != nil {
		return normalizedConfig{}, err
	}

	normalizedReadTimeout, err := normalizeTimeout(readTimeout, "readTimeout")
	if err != nil {
		return normalizedConfig{}, err
	}

	return normalizedConfig{
		APIKey:         apiKey,
		RequestOrigin:  requestOrigin,
		BaseURL:        normalizedBaseURL,
		ConnectTimeout: normalizedConnectTimeout,
		ReadTimeout:    normalizedReadTimeout,
	}, nil
}

func normalizeAPIKey(value string) (string, error) {
	normalized := strings.TrimSpace(value)
	if value != "" && normalized == "" {
		return "", &ValidationError{Message: "apiKey must not be blank"}
	}

	if normalized == "" {
		return "", nil
	}

	return normalized, nil
}

func normalizeRequestOrigin(value string) (string, error) {
	normalized := strings.TrimSpace(value)
	if value != "" && normalized == "" {
		return "", &ValidationError{Message: "requestOrigin must not be blank"}
	}

	if normalized == "" {
		return "", nil
	}

	if strings.ContainsAny(normalized, "\r\n") {
		return "", &ValidationError{Message: "requestOrigin must not contain CR or LF"}
	}

	parsed, err := url.Parse(normalized)
	if err != nil || !parsed.IsAbs() || (parsed.Scheme != "http" && parsed.Scheme != "https") {
		return "", &ValidationError{Message: "requestOrigin must be an absolute http or https origin"}
	}

	if parsed.User != nil {
		return "", &ValidationError{Message: "requestOrigin must not include userinfo"}
	}

	if parsed.RawQuery != "" || parsed.Fragment != "" {
		return "", &ValidationError{Message: "requestOrigin must not include query or fragment"}
	}

	if parsed.Path != "" && parsed.Path != "/" {
		return "", &ValidationError{Message: "requestOrigin must not include a path"}
	}

	return parsed.Scheme + "://" + parsed.Host, nil
}

func normalizeBaseURL(value string) (string, error) {
	normalized := strings.TrimRight(strings.TrimSpace(value), "/")
	if normalized == "" {
		return "", &ValidationError{Message: "baseUrl must not be blank"}
	}

	parsed, err := url.Parse(normalized)
	if err != nil || !parsed.IsAbs() || (parsed.Scheme != "http" && parsed.Scheme != "https") {
		return "", &ValidationError{Message: "baseUrl must be an absolute http or https URL"}
	}

	if parsed.User != nil {
		return "", &ValidationError{Message: "baseUrl must not include userinfo"}
	}

	if parsed.RawQuery != "" || parsed.Fragment != "" {
		return "", &ValidationError{Message: "baseUrl must not include query or fragment"}
	}

	return normalized, nil
}

func normalizeTimeout(value time.Duration, fieldName string) (time.Duration, error) {
	if value <= 0 {
		return 0, &ValidationError{Message: fieldName + " must be greater than zero"}
	}

	return value, nil
}

func emptyStringAsNil(value string) interface{} {
	if value == "" {
		return nil
	}

	return value
}
