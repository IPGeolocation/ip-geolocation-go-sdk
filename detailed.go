package ipgeolocation

import (
	"encoding/json"
	"strconv"
)

type DetailedIPGeolocationResponse struct {
	IPGeolocationResponse
	Network  *DetailedNetwork  `json:"network,omitempty"`
	Security *DetailedSecurity `json:"security,omitempty"`
}

type DetailedNetwork struct {
	*Network
	IsCDN           *bool   `json:"is_cdn,omitempty"`
	CDNProviderName *string `json:"cdn_provider_name,omitempty"`
}

type DetailedSecurity struct {
	*Security
	ProxyDaysSeen7d              *int     `json:"proxy_days_seen_7d,omitempty"`
	ProxyDaysSeen30d             *int     `json:"proxy_days_seen_30d,omitempty"`
	ProxyProviderCount30d        *int     `json:"proxy_provider_count_30d,omitempty"`
	BotType                      *string  `json:"bot_type,omitempty"`
	BotLastSeen                  *string  `json:"bot_last_seen,omitempty"`
	BotConfidenceScore           *float64 `json:"bot_confidence_score,omitempty"`
	BotOperatorName              *string  `json:"bot_operator_name,omitempty"`
	IsKnownGoodBot               *bool    `json:"is_known_good_bot,omitempty"`
	IsCorporateGateway           *bool    `json:"is_corporate_gateway,omitempty"`
	CorporateGatewayType         *string  `json:"corporate_gateway_type,omitempty"`
	CorporateGatewayProviderName *string  `json:"corporate_gateway_provider_name,omitempty"`
}

type DetailedBulkLookupResult struct {
	Data  *DetailedIPGeolocationResponse `json:"data,omitempty"`
	Error *BulkLookupError               `json:"error,omitempty"`
}

func (r DetailedBulkLookupResult) IsSuccess() bool { return r.Data != nil && r.Error == nil }

func parseDetailedResponse(object map[string]json.RawMessage) (*DetailedIPGeolocationResponse, error) {
	base, err := parseResponse(object)
	if err != nil {
		return nil, err
	}
	network, err := readObject(object, "network", func(n map[string]json.RawMessage) (DetailedNetwork, error) {
		return DetailedNetwork{Network: base.Network, IsCDN: readBool(n, "is_cdn"), CDNProviderName: readString(n, "cdn_provider_name")}, nil
	})
	if err != nil {
		return nil, err
	}
	security, err := readObject(object, "security", func(s map[string]json.RawMessage) (DetailedSecurity, error) {
		return DetailedSecurity{
			Security:                     base.Security,
			ProxyDaysSeen7d:              readCount(s, "proxy_days_seen_7d"),
			ProxyDaysSeen30d:             readCount(s, "proxy_days_seen_30d"),
			ProxyProviderCount30d:        readCount(s, "proxy_provider_count_30d"),
			BotType:                      readString(s, "bot_type"),
			BotLastSeen:                  readString(s, "bot_last_seen"),
			BotConfidenceScore:           readFloat(s, "bot_confidence_score"),
			BotOperatorName:              readString(s, "bot_operator_name"),
			IsKnownGoodBot:               readBool(s, "is_known_good_bot"),
			IsCorporateGateway:           readBool(s, "is_corporate_gateway"),
			CorporateGatewayType:         readString(s, "corporate_gateway_type"),
			CorporateGatewayProviderName: readString(s, "corporate_gateway_provider_name"),
		}, nil
	})
	if err != nil {
		return nil, err
	}
	return &DetailedIPGeolocationResponse{IPGeolocationResponse: *base, Network: network, Security: security}, nil
}

func readCount(object map[string]json.RawMessage, key string) *int {
	v := readString(object, key)
	if v == nil {
		return nil
	}
	n, err := strconv.Atoi(*v)
	if err != nil {
		return nil
	}
	return &n
}
