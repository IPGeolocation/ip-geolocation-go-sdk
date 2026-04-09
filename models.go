package ipgeolocation

type IPGeolocationResponse struct {
	IP              *string          `json:"ip,omitempty"`
	Hostname        *string          `json:"hostname,omitempty"`
	Domain          *string          `json:"domain,omitempty"`
	Location        *Location        `json:"location,omitempty"`
	CountryMetadata *CountryMetadata `json:"country_metadata,omitempty"`
	Currency        *Currency        `json:"currency,omitempty"`
	Network         *Network         `json:"network,omitempty"`
	ASN             *ASNInfo         `json:"asn,omitempty"`
	Company         *Company         `json:"company,omitempty"`
	TimeZone        *TimeZone        `json:"time_zone,omitempty"`
	Security        *Security        `json:"security,omitempty"`
	UserAgent       *UserAgent       `json:"user_agent,omitempty"`
	Abuse           *Abuse           `json:"abuse,omitempty"`
}

type Location struct {
	ContinentCode       *string `json:"continent_code,omitempty"`
	ContinentName       *string `json:"continent_name,omitempty"`
	CountryCode2        *string `json:"country_code2,omitempty"`
	CountryCode3        *string `json:"country_code3,omitempty"`
	CountryName         *string `json:"country_name,omitempty"`
	CountryNameOfficial *string `json:"country_name_official,omitempty"`
	CountryCapital      *string `json:"country_capital,omitempty"`
	StateProv           *string `json:"state_prov,omitempty"`
	StateCode           *string `json:"state_code,omitempty"`
	District            *string `json:"district,omitempty"`
	City                *string `json:"city,omitempty"`
	Locality            *string `json:"locality,omitempty"`
	AccuracyRadius      *string `json:"accuracy_radius,omitempty"`
	Confidence          *string `json:"confidence,omitempty"`
	DMACode             *string `json:"dma_code,omitempty"`
	Zipcode             *string `json:"zipcode,omitempty"`
	Latitude            *string `json:"latitude,omitempty"`
	Longitude           *string `json:"longitude,omitempty"`
	IsEU                *bool   `json:"is_eu,omitempty"`
	CountryFlag         *string `json:"country_flag,omitempty"`
	GeonameID           *string `json:"geoname_id,omitempty"`
	CountryEmoji        *string `json:"country_emoji,omitempty"`
}

type CountryMetadata struct {
	CallingCode *string  `json:"calling_code,omitempty"`
	TLD         *string  `json:"tld,omitempty"`
	Languages   []string `json:"languages,omitempty"`
}

type Currency struct {
	Code   *string `json:"code,omitempty"`
	Name   *string `json:"name,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
}

type Network struct {
	ConnectionType *string `json:"connection_type,omitempty"`
	Route          *string `json:"route,omitempty"`
	IsAnycast      *bool   `json:"is_anycast,omitempty"`
}

type ASNInfo struct {
	ASNumber      *string `json:"as_number,omitempty"`
	Organization  *string `json:"organization,omitempty"`
	Country       *string `json:"country,omitempty"`
	Type          *string `json:"type,omitempty"`
	Domain        *string `json:"domain,omitempty"`
	DateAllocated *string `json:"date_allocated,omitempty"`
	RIR           *string `json:"rir,omitempty"`
}

type Company struct {
	Name   *string `json:"name,omitempty"`
	Type   *string `json:"type,omitempty"`
	Domain *string `json:"domain,omitempty"`
}

type TimeZone struct {
	Name                         *string        `json:"name,omitempty"`
	Offset                       *float64       `json:"offset,omitempty"`
	OffsetWithDST                *float64       `json:"offset_with_dst,omitempty"`
	CurrentTime                  *string        `json:"current_time,omitempty"`
	CurrentTimeUnix              *float64       `json:"current_time_unix,omitempty"`
	CurrentTimeZoneAbbreviation  *string        `json:"current_tz_abbreviation,omitempty"`
	CurrentTimeZoneFullName      *string        `json:"current_tz_full_name,omitempty"`
	StandardTimeZoneAbbreviation *string        `json:"standard_tz_abbreviation,omitempty"`
	StandardTimeZoneFullName     *string        `json:"standard_tz_full_name,omitempty"`
	IsDST                        *bool          `json:"is_dst,omitempty"`
	DSTSavings                   *float64       `json:"dst_savings,omitempty"`
	DSTExists                    *bool          `json:"dst_exists,omitempty"`
	DSTTimeZoneAbbreviation      *string        `json:"dst_tz_abbreviation,omitempty"`
	DSTTimeZoneFullName          *string        `json:"dst_tz_full_name,omitempty"`
	DSTStart                     *DSTTransition `json:"dst_start,omitempty"`
	DSTEnd                       *DSTTransition `json:"dst_end,omitempty"`
}

type DSTTransition struct {
	UTCTime        *string `json:"utc_time,omitempty"`
	Duration       *string `json:"duration,omitempty"`
	Gap            *bool   `json:"gap,omitempty"`
	DateTimeAfter  *string `json:"date_time_after,omitempty"`
	DateTimeBefore *string `json:"date_time_before,omitempty"`
	Overlap        *bool   `json:"overlap,omitempty"`
}

type Security struct {
	ThreatScore          *float64 `json:"threat_score,omitempty"`
	IsTor                *bool    `json:"is_tor,omitempty"`
	IsProxy              *bool    `json:"is_proxy,omitempty"`
	ProxyProviderNames   []string `json:"proxy_provider_names,omitempty"`
	ProxyConfidenceScore *float64 `json:"proxy_confidence_score,omitempty"`
	ProxyLastSeen        *string  `json:"proxy_last_seen,omitempty"`
	IsResidentialProxy   *bool    `json:"is_residential_proxy,omitempty"`
	IsVPN                *bool    `json:"is_vpn,omitempty"`
	VPNProviderNames     []string `json:"vpn_provider_names,omitempty"`
	VPNConfidenceScore   *float64 `json:"vpn_confidence_score,omitempty"`
	VPNLastSeen          *string  `json:"vpn_last_seen,omitempty"`
	IsRelay              *bool    `json:"is_relay,omitempty"`
	RelayProviderName    *string  `json:"relay_provider_name,omitempty"`
	IsAnonymous          *bool    `json:"is_anonymous,omitempty"`
	IsKnownAttacker      *bool    `json:"is_known_attacker,omitempty"`
	IsBot                *bool    `json:"is_bot,omitempty"`
	IsSpam               *bool    `json:"is_spam,omitempty"`
	IsCloudProvider      *bool    `json:"is_cloud_provider,omitempty"`
	CloudProviderName    *string  `json:"cloud_provider_name,omitempty"`
}

type Abuse struct {
	Route        *string  `json:"route,omitempty"`
	Country      *string  `json:"country,omitempty"`
	Name         *string  `json:"name,omitempty"`
	Organization *string  `json:"organization,omitempty"`
	Kind         *string  `json:"kind,omitempty"`
	Address      *string  `json:"address,omitempty"`
	Emails       []string `json:"emails,omitempty"`
	PhoneNumbers []string `json:"phone_numbers,omitempty"`
}

type UserAgent struct {
	Name            *string                   `json:"name,omitempty"`
	Type            *string                   `json:"type,omitempty"`
	Version         *string                   `json:"version,omitempty"`
	VersionMajor    *string                   `json:"version_major,omitempty"`
	Device          *UserAgentDevice          `json:"device,omitempty"`
	Engine          *UserAgentEngine          `json:"engine,omitempty"`
	OperatingSystem *UserAgentOperatingSystem `json:"operating_system,omitempty"`
}

type UserAgentDevice struct {
	Name  *string `json:"name,omitempty"`
	Type  *string `json:"type,omitempty"`
	Brand *string `json:"brand,omitempty"`
	CPU   *string `json:"cpu,omitempty"`
}

type UserAgentEngine struct {
	Name         *string `json:"name,omitempty"`
	Type         *string `json:"type,omitempty"`
	Version      *string `json:"version,omitempty"`
	VersionMajor *string `json:"version_major,omitempty"`
}

type UserAgentOperatingSystem struct {
	Name    *string `json:"name,omitempty"`
	Version *string `json:"version,omitempty"`
}

// BulkLookupResult holds one bulk lookup item.
//
// Successful items use Data. Failed items use Error.
type BulkLookupResult struct {
	Data  *IPGeolocationResponse `json:"data,omitempty"`
	Error *BulkLookupError       `json:"error,omitempty"`
}

// IsSuccess reports whether the bulk lookup item succeeded.
func (r BulkLookupResult) IsSuccess() bool {
	return r.Data != nil && r.Error == nil
}

type BulkLookupError struct {
	Message string `json:"message"`
}
