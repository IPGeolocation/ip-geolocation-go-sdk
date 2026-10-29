# IPGeolocation Go SDK

Official Go SDK for the IPGeolocation.io IP Geolocation API.

Look up IPv4, IPv6, and domains with `/v3/ipgeo` and `/v3/ipgeo-bulk`. Get geolocation, company, ASN, timezone, network, hostname, abuse, user-agent, and security data from one API.

- Go module with `v2` import path
- Typed responses plus raw JSON and XML methods
- Built on the Go standard library

## Table of Contents

- [Install](#install)
- [Quick Start](#quick-start)
- [At a Glance](#at-a-glance)
- [Get Your API Key](#get-your-api-key)
- [Authentication](#authentication)
- [Plan Behavior](#plan-behavior)
- [Client Configuration](#client-configuration)
- [Available Methods](#available-methods)
- [Request Options](#request-options)
- [Examples](#examples)
- [Response Metadata](#response-metadata)
- [Errors](#errors)
- [Troubleshooting](#troubleshooting)
- [Frequently Asked Questions](#frequently-asked-questions)
- [Links](#links)

## Install

```bash
go get github.com/IPGeolocation/ip-geolocation-go-sdk/v2
```

```go
import ipgeolocation "github.com/IPGeolocation/ip-geolocation-go-sdk/v2"
```

Go module: `github.com/IPGeolocation/ip-geolocation-go-sdk/v2`
Package page: <https://pkg.go.dev/github.com/IPGeolocation/ip-geolocation-go-sdk/v2>
GitHub repository: <https://github.com/IPGeolocation/ip-geolocation-go-sdk>

## Quick Start

```go
package main

import (
	"context"
	"fmt"
	"os"

	ipgeolocation "github.com/IPGeolocation/ip-geolocation-go-sdk/v2"
)

func main() {
	apiKey := os.Getenv("IPGEO_API_KEY")
	if apiKey == "" {
		panic("set IPGEO_API_KEY first")
	}

	client, err := ipgeolocation.NewClient(&ipgeolocation.Config{
		APIKey: apiKey,
	})
	if err != nil {
		panic(err)
	}
	defer client.Close()

	response, err := client.LookupIPGeolocation(context.Background(), &ipgeolocation.LookupRequest{
		IP: "8.8.8.8",
	})
	if err != nil {
		panic(err)
	}

	data := response.Data

	if data.IP != nil {
		fmt.Println(*data.IP) // 8.8.8.8
	}

	if data.Location != nil && data.Location.CountryName != nil {
		fmt.Println(*data.Location.CountryName) // United States
	}

	if data.Location != nil && data.Location.City != nil {
		fmt.Println(*data.Location.City)
	}

	if data.TimeZone != nil && data.TimeZone.Name != nil {
		fmt.Println(*data.TimeZone.Name)
	}

	if response.Metadata.CreditsCharged != nil {
		fmt.Println(*response.Metadata.CreditsCharged)
	}
}
```

## At a Glance

| Item | Value |
|------|-------|
| Module | `github.com/IPGeolocation/ip-geolocation-go-sdk/v2` |
| Package | `ipgeolocation` |
| Supported Endpoints | `/v3/ipgeo`, `/v3/ipgeo-bulk` |
| Supported Inputs | IPv4, IPv6, domain |
| Main Data Returned | Geolocation, company, ASN, timezone, network, hostname, abuse, user-agent, currency, security |
| Authentication | API key, request-origin auth for `/v3/ipgeo` only |
| Response Formats | Structured JSON, raw JSON, raw XML |
| Bulk Limit | Up to 50,000 IPs or domains per request |
| Transport | Go standard library `net/http` |

## Get Your API Key

Create an IPGeolocation account and copy an API key from your dashboard.

1. Sign up: <https://app.ipgeolocation.io/signup>
2. Verify your email if prompted
3. Sign in: <https://app.ipgeolocation.io/login>
4. Open your dashboard: <https://app.ipgeolocation.io/dashboard>
5. Copy an API key from the `API Keys` section

For server-side code, keep the API key in an environment variable or secret manager. For browser-based single lookups on paid plans, use request-origin auth instead of exposing an API key in frontend code.

## Authentication

### API Key

```go
client, err := ipgeolocation.NewClient(&ipgeolocation.Config{
	APIKey: os.Getenv("IPGEO_API_KEY"),
})
```

### Request-Origin Auth

```go
client, err := ipgeolocation.NewClient(&ipgeolocation.Config{
	RequestOrigin: "https://app.example.com",
})
```

`RequestOrigin` must be an absolute `http` or `https` origin with no path, query string, fragment, or userinfo.

> [!IMPORTANT]
> Request-origin auth does not work with `/v3/ipgeo-bulk`. Bulk lookup always requires `APIKey`.

> [!NOTE]
> If you set both `APIKey` and `RequestOrigin`, single lookup still uses the API key. The API key is sent as the `apiKey` query parameter, so avoid logging full request URLs.

## Plan Behavior

Feature availability depends on your plan and request parameters.

| Capability | Free | Paid |
|------------|------|------|
| Single IPv4 and IPv6 lookup | Supported | Supported |
| Domain lookup | Not supported | Supported |
| Bulk lookup | Not supported | Supported |
| Non-English `lang` | Not supported | Supported |
| Request-origin auth | Not supported | Supported for `/v3/ipgeo` only |
| Optional modules via `include` | Not supported | Supported |
| `include: ["*"]` | Base response only | All plan-available modules |

Paid plans still need `include` for optional modules. `fields` and `excludes` only trim the response. They do not turn modules on or unlock paid data.

## Client Configuration

| Field | Type | Default | Notes |
|-------|------|---------|-------|
| `APIKey` | `string` | unset | Required for bulk lookup. Optional for single lookup if `RequestOrigin` is set. |
| `RequestOrigin` | `string` | unset | Must be an absolute `http` or `https` origin. |
| `BaseURL` | `string` | `https://api.ipgeolocation.io` | Override the API base URL. |
| `ConnectTimeout` | `time.Duration` | `10 * time.Second` | Time to open the connection or wait for response headers. |
| `ReadTimeout` | `time.Duration` | `30 * time.Second` | Time to wait while reading the response body. |

Create a client with `NewClient`. Config values are validated when the client is created. Request values are validated before each request is sent.

## Available Methods

| Method | Returns | Notes |
|--------|---------|-------|
| `LookupIPGeolocation(ctx, request)` | `*APIResponse[*IPGeolocationResponse]` | Single lookup. Typed JSON response. |
| `LookupIPGeolocationRaw(ctx, request)` | `*APIResponse[string]` | Single lookup. Raw JSON or XML string. |
| `BulkLookupIPGeolocation(ctx, request)` | `*APIResponse[[]BulkLookupResult]` | Bulk lookup. Typed JSON response. |
| `BulkLookupIPGeolocationRaw(ctx, request)` | `*APIResponse[string]` | Bulk lookup. Raw JSON or XML string. |
| `Close()` | `void` | Closes idle HTTP connections and marks the client closed. |

> [!NOTE]
> Typed methods support JSON only. Use the raw methods when you need XML output.

## Request Options

| Field | Applies To | Notes |
|-------|------------|-------|
| `IP` | Single lookup | IPv4, IPv6, or domain. Leave it empty for caller IP lookup. |
| `IPs` | Bulk lookup | Collection of 1 to 50,000 IPs or domains. |
| `Lang` | Single and bulk | One of `en`, `de`, `ru`, `ja`, `fr`, `cn`, `es`, `cs`, `it`, `ko`, `fa`, `pt`. |
| `Include` | Single and bulk | Collection of module names such as `security`, `abuse`, `user_agent`, `hostname`, `liveHostname`, `hostnameFallbackLive`, `geo_accuracy`, `dma_code`, or `*`. |
| `Fields` | Single and bulk | Collection of field paths to keep, for example `location.country_name` or `security.threat_score`. |
| `Excludes` | Single and bulk | Collection of field paths to remove from the response. |
| `UserAgent` | Single and bulk | Overrides the outbound `User-Agent` header. |
| `Headers` | Single and bulk | Extra request headers. |
| `Output` | Single and bulk | `ResponseFormatJSON` or `ResponseFormatXML`. Typed methods require JSON. |

## Examples

The examples below assume you already have a configured client in scope:

```go
apiKey := os.Getenv("IPGEO_API_KEY")

client, err := ipgeolocation.NewClient(&ipgeolocation.Config{
	APIKey: apiKey,
})
if err != nil {
	panic(err)
}
defer client.Close()
```

### Caller IP

Leave `IP` empty to look up the public IP of the machine making the request.

```go
response, err := client.LookupIPGeolocation(context.Background(), nil)
if err != nil {
	panic(err)
}

if response.Data.IP != nil {
	fmt.Println(*response.Data.IP)
}
```

### Domain Lookup

Domain lookup is a paid-plan feature.

```go
response, err := client.LookupIPGeolocation(context.Background(), &ipgeolocation.LookupRequest{
	IP: "ipgeolocation.io",
})
if err != nil {
	panic(err)
}

if response.Data.Domain != nil {
	fmt.Println(*response.Data.Domain) // ipgeolocation.io
}
```

### Security and Abuse

```go
response, err := client.LookupIPGeolocation(context.Background(), &ipgeolocation.LookupRequest{
	IP:      "9.9.9.9",
	Include: []string{"security", "abuse"},
})
if err != nil {
	panic(err)
}

if response.Data.Security != nil && response.Data.Security.ThreatScore != nil {
	fmt.Println(*response.Data.Security.ThreatScore)
}

if response.Data.Abuse != nil && len(response.Data.Abuse.Emails) > 0 {
	fmt.Println(response.Data.Abuse.Emails[0])
}
```

### User-Agent Parsing

To parse a visitor user-agent string, pass `Include: []string{"user_agent"}` and send the visitor string in the request `User-Agent` header.

```go
visitorUA := "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_2) AppleWebKit/601.3.9 (KHTML, like Gecko) Version/9.0.2 Safari/601.3.9"

response, err := client.LookupIPGeolocation(context.Background(), &ipgeolocation.LookupRequest{
	IP:      "115.240.90.163",
	Include: []string{"user_agent"},
	Headers: map[string]string{
		"User-Agent": visitorUA,
	},
})
if err != nil {
	panic(err)
}

if response.Data.UserAgent != nil && response.Data.UserAgent.Name != nil {
	fmt.Println(*response.Data.UserAgent.Name)
}
```

> [!NOTE]
> The `UserAgent` request field overrides the SDK's default outbound `User-Agent` header. It takes precedence over `Headers["User-Agent"]`.
> `response.Data.UserAgent` is different. That field is the parsed visitor user-agent data returned by the API.

### Filtered Response

```go
response, err := client.LookupIPGeolocation(context.Background(), &ipgeolocation.LookupRequest{
	IP:       "8.8.8.8",
	Include:  []string{"security"},
	Fields:   []string{"location.country_name", "security.threat_score", "security.is_vpn"},
	Excludes: []string{"currency"},
})
if err != nil {
	panic(err)
}

if response.Data.Location != nil && response.Data.Location.CountryName != nil {
	fmt.Println(*response.Data.Location.CountryName)
}

if response.Data.Security != nil && response.Data.Security.IsVPN != nil {
	fmt.Println(*response.Data.Security.IsVPN)
}

fmt.Println(response.Data.Currency == nil)
```

### Raw XML

```go
response, err := client.LookupIPGeolocationRaw(context.Background(), &ipgeolocation.LookupRequest{
	IP:     "8.8.8.8",
	Output: ipgeolocation.ResponseFormatXML,
})
if err != nil {
	panic(err)
}

fmt.Println(response.Data)
```

### Bulk Lookup

```go
response, err := client.BulkLookupIPGeolocation(context.Background(), &ipgeolocation.BulkLookupRequest{
	IPs: []string{"8.8.8.8", "1.1.1.1"},
})
if err != nil {
	panic(err)
}

for _, result := range response.Data {
	if result.IsSuccess() && result.Data != nil && result.Data.IP != nil {
		fmt.Println(*result.Data.IP)
		continue
	}

	if result.Error != nil {
		fmt.Println(result.Error.Message)
	}
}
```

### Raw Bulk JSON

```go
response, err := client.BulkLookupIPGeolocationRaw(context.Background(), &ipgeolocation.BulkLookupRequest{
	IPs: []string{"8.8.8.8", "1.1.1.1"},
})
if err != nil {
	panic(err)
}

fmt.Println(response.Data)
```

## Response Metadata

Every method returns `APIResponse[T]`, where:

- `Data` contains the typed object or raw response string
- `Metadata` contains response details such as:
  - `CreditsCharged`
  - `SuccessfulRecords`
  - `StatusCode`
  - `DurationMS`
  - `RawHeaders`

Example:

```go
fmt.Println(response.Metadata.StatusCode)
fmt.Println(response.Metadata.DurationMS)
fmt.Println(response.Metadata.FirstHeaderValue("content-type"))
```

## Errors

The SDK returns typed errors instead of leaving you to parse error bodies yourself.

| Error Type | When it happens |
|------------|-----------------|
| `*ValidationError` | Invalid config, invalid request values, or typed XML request |
| `*RequestTimeoutError` | Connect timeout or read timeout |
| `*TransportError` | Network or transport failure |
| `*SerializationError` | Request or response serialization failure |
| `*APIError` | API returned a non-2xx response |

`*APIError` has status-specific types:

- `*BadRequestError`
- `*UnauthorizedError`
- `*ForbiddenError`
- `*NotFoundError`
- `*MethodNotAllowedError`
- `*ContentTooLargeError`
- `*UnsupportedMediaTypeError`
- `*LockedError`
- `*TooManyRequestsError`
- `*CustomStatus499Error`
- `*InternalServerError`

Example:

```go
var validationErr *ipgeolocation.ValidationError

_, err := client.LookupIPGeolocation(context.Background(), &ipgeolocation.LookupRequest{
	Output: ipgeolocation.ResponseFormatXML,
})
if err != nil {
	if errors.As(err, &validationErr) {
		fmt.Println(validationErr.Message)
		return
	}

	fmt.Println(err)
}
```

## Troubleshooting

- Bulk lookup always requires `APIKey`. `RequestOrigin` is not enough.
- Typed methods only support JSON. Use the raw methods for XML.
- If you need security, abuse, user-agent, or hostname data, include those modules explicitly.
- `Fields` and `Excludes` filter the response. They do not unlock paid data.
- `RequestOrigin` must be an origin only. Do not include a path, query string, fragment, or userinfo.
- Response fields are pointers so omitted fields stay omitted.

## Frequently Asked Questions

<details>
<summary>Can I use this SDK without an API key?</summary>

Only for single lookup with paid-plan request-origin auth. Bulk lookup always requires an API key.
</details>

<details>
<summary>Can I request XML and still get typed models?</summary>

No. Typed methods only support JSON. Use `LookupIPGeolocationRaw` or `BulkLookupIPGeolocationRaw` for XML.
</details>

<details>
<summary>Why are many response fields pointers?</summary>

Pointers let the SDK preserve omitted fields instead of inventing zero values for data the API did not send.
</details>

<details>
<summary>Does domain lookup work on the free plan?</summary>

No. Domain lookup is a paid-plan feature.
</details>

## Links

- Homepage: <https://ipgeolocation.io>
- IP Geolocation API product page: <https://ipgeolocation.io/ip-location-api.html>
- Documentation: <https://ipgeolocation.io/documentation/ip-location-api.html>
- Go package: <https://pkg.go.dev/github.com/IPGeolocation/ip-geolocation-go-sdk/v2>
- GitHub repository: <https://github.com/IPGeolocation/ip-geolocation-go-sdk>
