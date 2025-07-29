# Go API client for IPGeolocation.io

Ipgeolocation provides a set of APIs to make ip based decisions.

- Supports IPGeolocation API version: 2.0
- Package version: 1.0.1

## Table of Contents

1. [Installation](#installation)
2. [Authentication Setup](#authentication-setup)
3. [Custom HTTP Client Timeout](#custom-http-client-timeout)
4. [API Endpoints](#api-endpoints)
5. [Accessing Fields in Response Models](#accessing-fields-in-response-models)
  - [1. Recommended: Use Getter Methods](#recommended-use-getter-methods)
  - [2. If You Still Want to Use Struct Fields](#if-you-still-want-to-use-struct-fields)
  - [3. Applies to All Models](#applies-to-all-models)
6. [IP Geolocation Examples](#ip-geolocation-examples)
   - [1. Basic Plan Examples](#basic-plan-examples)
   - [2. Standard Plan Examples](#2-standard-plan-examples)
   - [3. Advanced Plan Examples](#3-advanced-plan-examples)
   - [Bulk IP Geolocation Example](#bulk-ip-geolocation-example)

4. [IP Security Examples](#ip-security-examples)
   - [Basic Request (Minimal Setup)](#basic-request-minimal-setup)
   - [Include Multiple Optional Fields](#include-multiple-optional-fields)
   - [Request with Field Filtering](#request-with-field-filtering)
   - [Bulk IP Security Request](#bulk-ip-security-request)

5. [ASN API Examples](#asn-api-examples)
   - [Get ASN Information by IP Address](#get-asn-information-by-ip-address)
   - [Get ASN Information by ASN Number](#get-asn-information-by-asn-number)
   - [Combine All objects using Include](#combine-all-objects-using-include)

6. [Timezone API Examples](#timezone-api-examples)
   - [Get Timezone by IP Address](#get-timezone-by-ip-address)
   - [Get Timezone by Timezone Name](#get-timezone-by-timezone-name)
   - [Get Timezone from Any Address](#get-timezone-from-any-address)
   - [Get Timezone from Location Coordinates](#get-timezone-from-location-coordinates)
   - [Get Timezone and Airport Details from IATA Code](#get-timezone-and-airport-details-from-iata-code)
   - [Get Timezone and City Details from UN/LOCODE](#get-timezone-and-city-details-from-unlocode)

7. [Timezone Converter Examples](#timezone-converter-examples)
   - [Convert Current Time from One Timezone to Another](#convert-current-time-from-one-timezone-to-another)

8. [User Agent API Examples](#user-agent-api-examples)
   - [Parse a Basic User Agent String](#parse-a-basic-user-agent-string)
   - [Bulk User Agent Parsing Example](#bulk-user-agent-parsing-example)
9. [Astronomy API Examples](#astronomy-api-examples)
   - [Astronomy by Coordinates](#astronomy-by-coordinates)
   - [Astronomy by IP Address](#astronomy-by-ip-address)
   - [Astronomy by Location String](#astronomy-by-location-string)
   - [Astronomy for Specific Date](#astronomy-for-specific-date)
   - [Astronomy in Different Language](#astronomy-in-different-language)

10. [Documentation for Models](#documentation-for-models)


# Installation

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
go get github.com/IPGeolocation/ip-geolocation-go-sdk
```

Put the package under your project folder and add the following in import:

```go
import ipgeolocation "github.com/IPGeolocation/ip-geolocation-go-sdk/sdk"
```


# Authentication Setup

The API key for your IPGeolocation account is required for all requests. You may pass the api key as follows: 

```go
ctx := ipgeolocation.WithAPIKey(context.Background(), "<YOUR_API_KEY>")
configuration := ipgeolocation.NewConfiguration()
apiClient := ipgeolocation.NewAPIClient(configuration)
```

## Custom HTTP Client Timeout

You can configure the SDK to use a custom HTTP client — useful for setting timeouts, proxies, or other advanced behaviors.

### Example:

```go
import (
    "net/http"
    "time"
)

cfg := ipgeolocationsdk.NewConfiguration()

// Set custom timeout of 5 seconds
cfg.HTTPClient = &http.Client{
    Timeout: 5 * time.Second,
}
```

## API Endpoints

All URIs are relative to *https://api.ipgeolocation.io/v2*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ASNLookupAPI* | [**GetAsnInfo**](docs/ASNLookupAPI.md#getasninfo) | **Get** /asn | Get details of any ASN number
*AbuseContactAPI* | [**GetAbuseContactInfo**](docs/AbuseContactAPI.md#getabusecontactinfo) | **Get** /abuse | Retrieve abuse contact data for an IP address
*AstronomyAPI* | [**GetAstronomyDetails**](docs/AstronomyAPI.md#getastronomydetails) | **Get** /astronomy | Get sun and moon timings and positions
*IPGeolocationAPI* | [**GetBulkIpGeolocation**](docs/IPLocationAPI.md#getbulkipgeolocation) | **Post** /ipgeo-bulk | 	Get geolocation data for multiple IP addresses
*IPGeolocationAPI* | [**GetIpGeolocation**](docs/IPLocationAPI.md#getipgeolocation) | **Get** /ipgeo | Get geolocation data for a single IP address
*IPSecurityAPI* | [**GetBulkIpSecurityInfo**](docs/SecurityAPI.md#getbulkipsecurityinfo) | **Post** /security-bulk | Get threat intelligence for multiple IP addresses 
*IPSecurityAPI* | [**GetIpSecurityInfo**](docs/SecurityAPI.md#getipsecurityinfo) | **Get** /security | 	Get threat intelligence for a single IP address
*TimeConversionAPI* | [**ConvertTimeBetweenTimezones**](docs/TimeConversionAPI.md#converttimebetweentimezones) | **Get** /timezone/convert | Convert time from one timezone to another
*TimezoneAPI* | [**GetTimezoneInfo**](docs/TimezoneAPI.md#gettimezoneinfo) | **Get** /timezone | Get timezone information based on IP, coordinates, or name
*UserAgentAPI* | [**GetUserAgentDetails**](docs/UserAgentAPI.md#getuseragentdetails) | **Get** /user-agent | Parse a single user-agent string
*UserAgentAPI* | [**ParseBulkUserAgentStrings**](docs/UserAgentAPI.md#parsebulkuseragentstrings) | **Post** /user-agent-bulk | Parse multiple user-agent strings

### Accessing Fields in Response Models

When working with responses returned by the SDK, most fields are represented as **pointers** (`*string`, `*int`, or custom types like `*Location`). This allows for precise nullability handling, but it also means that:

- You can't access values like `resp.Ip` directly without dereferencing (`*resp.Ip`)
- This can lead to verbose or error-prone code if not handled properly

#### Recommended: Use Getter Methods

All response models include `GetX()` methods for safe, nil-checked access.

**Example:**

```go
resp, _, err := apiClient.IPGeolocationAPI.
	GetIpGeolocation(ctx).
	Ip("8.8.8.8").
	Execute()

if err != nil {
	log.Fatalf("Error fetching IP details: %v", err)
}

// Safely access values
fmt.Println("IP:", resp.GetIp())
fmt.Println("City:", resp.Location.GetCity())
```

These getter methods will:

- Return the value directly if it’s present
- Return the zero value (`""`, `0`, etc.) if the pointer is `nil`

---

#### If You Still Want to Use Struct Fields

You can directly access struct fields using pointer dereferencing, but you must check for nil:

```go
if resp.Ip != nil {
	fmt.Println("IP:", *resp.Ip)
}

if resp.Location != nil && resp.Location.City != nil {
	fmt.Println("City:", *resp.Location.City)
}
```

⚠️ This approach can panic if you skip nil checks, so it's recommended only if you need more control.

---

#### Applies to All Models

This pattern of using `*Type` fields + `GetX()` methods applies across **all models in the SDK**, including:

- `IpGeolocation`
- `Location`
- `Timezone`
- etc.

Wherever you see fields like `*string` or `*Location`, you can use the corresponding `.GetField()` method for safe and clean access.


# Example Usage

## IP Geolocation Examples

This section provides usage examples of the `getIPGeolocation()` method from the SDK across Free, Standard, and Advanced subscription tiers. Each example highlights different combinations of parameters: `fields`, `include`, and `excludes`.

### Parameters

#### `fields`
Use this parameter to include specific fields in the response.

#### `excludes`
Use this parameter to omit specific fields from the response.

#### `include`
Use this parameter to add optional modules to the response, such as:
- `security`
- `user_agent`
- `hostname`
- `liveHostname`
- `hostnameFallbackLive`
- `abuse`
- `dma`
- `timezone`


For complete details, refer to the official documentation: [IP Geolocation API Documentation](https://ipgeolocation.io/ip-location-api.html#documentation-overview)

The `ip` parameter in the SDK can accept any valid IPv4 address, IPv6 address, or domain name. If the `ip()` method is not used or the parameter is omitted, the API will return information about the public IP address of the device or server where the SDK is executed.

### 1. Basic Plan Examples

#### Default Fields

```go 
respGeolocation, res, err := apiClient.IPGeolocationAPI.
	GetIpGeolocation(ctx).
	Ip("8.8.8.8").
	Execute()

if err != nil {
	fmt.Fprintf(os.Stderr, "Error calling `IPGeolocationAPI.GetIpGeolocation`: %v\n", err)
	if res != nil {
		fmt.Fprintf(os.Stderr, "HTTP response: %v\n", res.Status)
	}
	return
}

responseJson, _ := json.MarshalIndent(respGeolocation, "", "  ")
fmt.Println(string(responseJson))
```

Sample Response:

```json

{
  "ip": "8.8.8.8",
  "location": {
    "city": "Mountain View",
    "continent_code": "NA",
    "continent_name": "North America",
    "country_capital": "Washington, D.C.",
    "country_code2": "US",
    "country_code3": "USA",
    "country_emoji": "🇺🇸",
    "country_flag": "https://ipgeolocation.io/static/flags/us_64.png",
    "country_name": "United States",
    "country_name_official": "United States of America",
    "district": "Santa Clara",
    "geoname_id": "6301403",
    "is_eu": false,
    "latitude": "37.42240",
    "longitude": "-122.08421",
    "state_code": "US-CA",
    "state_prov": "California",
    "zipcode": "94043-1351"
  }
  "country_metadata": {
    "calling_code": "+1",
    "languages": [
      "en-US",
      "es-US",
      "haw",
      "fr"
    ],
    "tld": ".us"
  },
  "currency": {
    "code": "USD",
    "name": "US Dollar",
    "symbol": "$"
  },
}
```

Filtering Specific Fields from the Response (Use of 'exclude' and 'fields')
```go
respGeolocation, res, err := apiClient.IPGeolocationAPI.
		GetIpGeolocation(ctx).
		Ip("8.8.8.8").
		Fields("location").
		Excludes("location.continent_code,location.continent_name").
		Execute()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error calling `IPGeolocationAPI.GetIpGeolocation`: %v\n", err)
		if res != nil {
			fmt.Fprintf(os.Stderr, "HTTP response: %v\n", res.Status)
		}
		return
	}

	responseJson, _ := json.MarshalIndent(respGeolocation, "", "  ")
	fmt.Println(string(responseJson))
```
Sample Response
```
model GeolocationResponse {
    ip: 8.8.4.4
    location: model Location {
        countryCode2: US
        countryCode3: USA
        countryName: United States
        countryNameOfficial: United States of America
        countryCapital: Washington, D.C.
        stateProv: California
        stateCode: US-CA
        district: Santa Clara
        city: Mountain View
        zipcode: 94043-1351
        latitude: 37.42240
        longitude: -122.08421
        isEu: false
        countryFlag: https://ipgeolocation.io/static/flags/us_64.png
        geonameId: 6301403
        countryEmoji: 🇺🇸
    }
}
```

### 2. Standard Plan Examples
#### Default Fields

```go
respGeolocation, res, err := apiClient.IPGeolocationAPI.
	GetIpGeolocation(ctx).
	Ip("8.8.8.8").
	Execute()

if err != nil {
	fmt.Fprintf(os.Stderr, "Error calling `IPGeolocationAPI.GetIpGeolocation`: %v\n", err)
	if res != nil {
		fmt.Fprintf(os.Stderr, "HTTP response: %v\n", res.Status)
	}
	return
}

responseJson, _ := json.MarshalIndent(respGeolocation, "", "  ")
fmt.Println(string(responseJson))
```
Sample Response:
```
model GeolocationResponse {
    ip: 8.8.8.8
    location: model Location {
        continentCode: NA
        continentName: North America
        countryCode2: US
        countryCode3: USA
        countryName: United States
        countryNameOfficial: United States of America
        countryCapital: Washington, D.C.
        stateProv: California
        stateCode: US-CA
        district: Santa Clara
        city: Mountain View
        zipcode: 94043-1351
        latitude: 37.42240
        longitude: -122.08421
        isEu: false
        countryFlag: https://ipgeolocation.io/static/flags/us_64.png
        geonameId: 6301403
        countryEmoji: 🇺🇸
    }
    countryMetadata: model CountryMetadata {
        callingCode: +1
        tld: .us
        languages: [en-US, es-US, haw, fr]
    }
    network: model Network {
        asn: model NetworkAsn {
            asNumber: AS15169
            organization: Google LLC
            country: US
        }
        company: model NetworkCompany {
            name: Google LLC
        }
    }
    currency: model Currency {
        code: USD
        name: US Dollar
        symbol: $
    }
}
```
### Retrieving Geolocation Data in Multiple Languages
Here is an example to get the geolocation data for IP address '2001:4230:4890::1' in French language:
```go
respGeolocation, res, err := apiClient.IPGeolocationAPI.
	GetIpGeolocation(ctx).
	Ip("2001:4230:4890::1").
	Language("fr").
	Execute()

if err != nil {
	fmt.Fprintf(os.Stderr, "Error calling `IPGeolocationAPI.GetIpGeolocation`: %v\n", err)
	if res != nil {
		fmt.Fprintf(os.Stderr, "HTTP response: %v\n", res.Status)
	}
	return
}

responseJson, _ := json.MarshalIndent(respGeolocation, "", "  ")
fmt.Println(string(responseJson))
```

Sample Response
```
model GeolocationResponse {
    ip: 2001:4230:4890:0:0:0:0:1
    location: model Location {
        continentCode: AF
        continentName: Afrique
        countryCode2: MU
        countryCode3: MUS
        countryName: Maurice
        countryNameOfficial: 
        countryCapital: Port Louis
        stateProv: Wilhems des plaines
        stateCode: MU-PW
        district: Quatre Bornes
        city: Quatre Bornes
        zipcode: 72201
        latitude: -20.24304
        longitude: 57.49631
        isEu: false
        countryFlag: https://ipgeolocation.io/static/flags/mu_64.png
        geonameId: 1106777
        countryEmoji: 🇲🇺
    }
    countryMetadata: model CountryMetadata {
        callingCode: +230
        tld: .mu
        languages: [en-MU, bho, fr]
    }
    network: model Network {
        asn: model NetworkAsn {
            asNumber: AS0
            organization: 
            country:
        }
        company: model NetworkCompany {
            name: African Network Information Center AfriNIC Ltd
        }
    }
    currency: model Currency {
        code: MUR
        name: Mauritius Rupee
        symbol: ₨
    }
}
```

#### Include HostName, Timezone and User-Agent
```go
respGeolocation, res, err := apiClient.IPGeolocationAPI.
	GetIpGeolocation(ctx).
	Ip("4.5.6.7").
	Fields("location.country_name,location.country_capital").
    Include("user_agent, time_zone, hostnameFallbackLive").
	Execute()

if err != nil {
	fmt.Fprintf(os.Stderr, "Error calling `IPGeolocationAPI.GetIpGeolocation`: %v\n", err)
	if res != nil {
		fmt.Fprintf(os.Stderr, "HTTP response: %v\n", res.Status)
	}
	return
}

responseJson, _ := json.MarshalIndent(respGeolocation, "", "  ")
fmt.Println(string(responseJson))
```
Sample Response
```
model GeolocationResponse {
    ip: 4.5.6.7
    hostname: 4.5.6.7
    location: model Location {
        countryName: United States
        countryCapital: Washington, D.C.
    }
    timeZone: model TimeZone {
        name: America/Chicago
        offset: -6
        offsetWithDst: -5
        currentTime: 2025-05-28 06:52:16.748-0500
        currentTimeUnix: 1748433136.748
        isDst: true
        dstSavings: 1
        dstExists: true
        dstStart: model TimeZoneDstStart {
            utcTime: 2025-03-09 TIME 08
            duration: +1H
            gap: true
            dateTimeAfter: 2025-03-09 TIME 03
            dateTimeBefore: 2025-03-09 TIME 02
            overlap: false
        }
        dstEnd: model TimeZoneDstEnd {
            utcTime: 2025-11-02 TIME 07
            duration: -1H
            gap: false
            dateTimeAfter: 2025-11-02 TIME 01
            dateTimeBefore: 2025-11-02 TIME 02
            overlap: true
        }
    }
    userAgent: model UserAgentData {
        userAgentString: IPGeolocation/2.0.0/java
        name: IPGeolocation Java SDK
        type: Special
        version: 2.0.0
        versionMajor: 1
        device: model UserAgentDataDevice {
            name: Unknown
            type: Unknown
            brand: Unknown
            cpu: Unknown
        }
        engine: model UserAgentDataEngine {
            name: Unknown
            type: Unknown
            version: ??
            versionMajor: ??
        }
        operatingSystem: model UserAgentDataOperatingSystem {
            name: Unknown
            type: Unknown
            version: ??
            versionMajor: ??
            build: ??
        }
    }
}
```
**Note on Hostname Parameters**

The IP Geolocation API supports hostname lookup for all paid subscriptions. However, this is not included by default. To enable hostname resolution, use the `include` parameter with one of the following options:

- `hostname`: Performs a quick lookup using the internal hostname database. If no match is found, the IP is returned as-is. This is fast but may produce incomplete results.
- `liveHostname`: Queries live sources for accurate hostname resolution. This may increase response time.
- `hostnameFallbackLive`: Attempts the internal database first, and falls back to live sources if no result is found. This option provides a balance of speed and reliability.

### 3. Advanced Plan Examples
#### Include DMA, Abuse and Security

```go
respGeolocation, res, err := apiClient.IPGeolocationAPI.
		GetIpGeolocation(ctx).
		Ip("8.8.8.8").
		Include("dma, abuse, security").
		Execute()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error calling `IPGeolocationAPI.GetIpGeolocation`: %v\n", err)
		if res != nil {
			fmt.Fprintf(os.Stderr, "HTTP response: %v\n", res.Status)
		}
		return
	}

	responseJson, _ := json.MarshalIndent(respGeolocation, "", "  ")
	fmt.Println(string(responseJson))
```
Sample Response:
```
model GeolocationResponse {
    ip: 8.8.8.8
    location: model Location {
        continentCode: NA
        continentName: North America
        countryCode2: US
        countryCode3: USA
        countryName: United States
        countryNameOfficial: United States of America
        countryCapital: Washington, D.C.
        stateProv: California
        stateCode: US-CA
        district: Santa Clara
        city: Mountain View
        zipcode: 94043-1351
        latitude: 37.42240
        longitude: -122.08421
        isEu: false
        countryFlag: null
        geonameId: 6301403
        countryEmoji: null
        accuracyRadius: 
        locality: Mountain View
        dmaCode: 807
    }
    countryMetadata: model CountryMetadata {
        callingCode: +1
        tld: .us
        languages: [en-US, es-US, haw, fr]
    }
    network: model Network {
        asn: model NetworkAsn {
            asNumber: AS15169
            organization: Google LLC
            country: US
            asnName: GOOGLE
            type: BUSINESS
            domain: about.google
            dateAllocated: 
            allocationStatus: assigned
            numOfIpv4Routes: 965
            numOfIpv6Routes: 104
            rir: ARIN
        }
        connectionType: 
        company: model NetworkCompany {
            name: Google LLC
            type: Business
            domain: googlellc.com
        }
    }
    currency: model Currency {
        code: USD
        name: US Dollar
        symbol: $
    }
    security: model Security {
        threatScore: 0
        isTor: false
        isProxy: false
        proxyType: 
        proxyProvider: 
        isAnonymous: false
        isKnownAttacker: false
        isSpam: false
        isBot: false
        isCloudProvider: false
        cloudProvider: 
    }
    abuse: model Abuse {
        route: 8.8.8.0/24
        country: 
        handle: ABUSE5250-ARIN
        name: Abuse
        organization: Abuse
        role: abuse
        kind: group
        address: 1600 Amphitheatre Parkway
        Mountain View
        CA
        94043
        United States
        emails: [network-abuse@google.com]
        phoneNumbers: [+1-650-253-0000]
    }
}
```
These examples demonstrate typical usage of the IP Geolocation API with different subscription tiers. Use `fields` to specify exactly which data to receive, `include` for optional data like security and user agent, and `excludes` to omit specific keys from the response.

**Note:** All features available in the Free plan are also included in the Standard and Advanced plans. Similarly, all features of the Standard plan are available in the Advanced plan.

## Bulk IP Geolocation Example
The SDK also supports bulk IP geolocation requests using the `getBulkIpGeolocation()` method. All parameters like `fields`, `include`, and `excludes` can also be used in bulk requests.

```go

respGeolocation, res, err := apiClient.IPGeolocationAPI.
	GetBulkIpGeolocation(ctx).
	Ips([]string{"8.8.8.8", "asdasdasd"}).
	Include("security,time_zone,user_agent,abuse,dma,hostname").
	Execute()

if err != nil {
	fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.GetSecurityInfo`: %v\n", err)
	fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", res)
	return
}

 // full JSON response
responseJson, _ := json.MarshalIndent(respGeolocation, "", "  ")
fmt.Println(string(responseJson))

// Loop through results and check for errors
for _, item := range respGeolocation {
	if item.IsError() { // Check for error response
		fmt.Println("❌", item.ErrorResponse.GetMessage())
	} else { //GeolocationResponse
		fmt.Println("✅", item.GeolocationResponse.GetIp())
	}
}

```

## IP Security Examples

This section provides usage examples of the `getIPSecurity()` method from the SDK across various subscription tiers. Each example demonstrates different ways to query threat intelligence and risk metadata using parameters like fields, excludes, and optional modules.

For full API specifications, refer to the [official IP Security API documentation](https://ipgeolocation.io/ip-security-api.html#documentation-overview).

---

### Basic Request (Minimal Setup)

```go
respSecurity, res, err := apiClient.IPSecurityAPI.GetIpSecurityInfo(ctx).Ip("2.56.188.34").Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.GetSecurityInfo`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", res)
		return
	}

	responseJson, _ := json.MarshalIndent(respSecurity, "", "  ")
	fmt.Println(string(responseJson))
```

Sample Response
```
model SecurityAPIResponse {
    ip: 2.56.188.34
    security: model Security {
        threatScore: 80
        isTor: false
        isProxy: true
        proxyType: VPN
        proxyProvider: Nord VPN
        isAnonymous: true
        isKnownAttacker: true
        isSpam: false
        isBot: false
        isCloudProvider: true
        cloudProvider: Packethub S.A.
    }
}
```

### Include Multiple Optional Fields
```go
respSecurity, res, err := apiClient.IPSecurityAPI.
		GetIpSecurityInfo(ctx).
		Ip("2.56.188.34").
		Include("location,network,currency,time_zone,user_agent,country_metadata,hostname").
		Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.GetSecurityInfo`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", res)
		return
	}

	responseJson, _ := json.MarshalIndent(respSecurity, "", "  ")
	fmt.Println(string(responseJson))
```

Sample Response
```
model SecurityAPIResponse {
    ip: 2.56.188.34
    hostname: 2.56.188.34
    security: model Security {
        threatScore: 80
        isTor: false
        isProxy: true
        proxyType: VPN
        proxyProvider: Nord VPN
        isAnonymous: true
        isKnownAttacker: true
        isSpam: false
        isBot: false
        isCloudProvider: true
        cloudProvider: Packethub S.A.
    }
    location: model LocationMinimal {
        continentCode: NA
        continentName: North America
        countryCode2: US
        countryCode3: USA
        countryName: United States
        countryNameOfficial: United States of America
        countryCapital: Washington, D.C.
        stateProv: Texas
        stateCode: US-TX
        district: Dallas County
        city: Dallas
        zipcode: 75207
        latitude: 32.78916
        longitude: -96.82170
        isEu: false
        countryFlag: https://ipgeolocation.io/static/flags/us_64.png
        geonameId: 7181768
        countryEmoji: 🇺🇸
    }
    network: model NetworkMinimal {
        asn: model NetworkMinimalAsn {
            asNumber: AS62240
            organization: Clouvider Limited
            country: GB
        }
        company: model NetworkMinimalCompany {
            name: Packethub S.A.
        }
    }
    timeZone: model TimeZone {
        name: America/Chicago
        offset: -6
        offsetWithDst: -5
        currentTime: 2025-05-29 08:27:44.939-0500
        currentTimeUnix: 1748525264.939
        isDst: true
        dstSavings: 1
        dstExists: true
        dstStart: model TimeZoneDstStart {
            utcTime: 2025-03-09 TIME 08
            duration: +1H
            gap: true
            dateTimeAfter: 2025-03-09 TIME 03
            dateTimeBefore: 2025-03-09 TIME 02
            overlap: false
        }
        dstEnd: model TimeZoneDstEnd {
            utcTime: 2025-11-02 TIME 07
            duration: -1H
            gap: false
            dateTimeAfter: 2025-11-02 TIME 01
            dateTimeBefore: 2025-11-02 TIME 02
            overlap: true
        }
    }
    userAgent: model UserAgentData {
        userAgentString: IPGeolocation/2.0.0/java
        name: IPGeolocation Java SDK
        type: Special
        version: 2.0.0
        versionMajor: 1
        device: model UserAgentDataDevice {
            name: Unknown
            type: Unknown
            brand: Unknown
            cpu: Unknown
        }
        engine: model UserAgentDataEngine {
            name: Unknown
            type: Unknown
            version: ??
            versionMajor: ??
        }
        operatingSystem: model UserAgentDataOperatingSystem {
            name: Unknown
            type: Unknown
            version: ??
            versionMajor: ??
            build: ??
        }
    }
    countryMetadata: model CountryMetadata {
        callingCode: +1
        tld: .us
        languages: [en-US, es-US, haw, fr]
    }
    currency: model Currency {
        code: USD
        name: US Dollar
        symbol: $
    }
}
```
### Request with Field Filtering

```go
respSecurity, res, err := apiClient.IPSecurityAPI.
		GetIpSecurityInfo(ctx).
		Ip("2.56.188.34").
		Fields("security.is_tor,security.is_proxy,security.is_bot,security.is_spam").
		Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.GetSecurityInfo`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", res)
		return
	}

	responseJson, _ := json.MarshalIndent(respSecurity, "", "  ")
	fmt.Println(string(responseJson))
```
Sample Response
```
model SecurityAPIResponse {
    ip: 195.154.221.54
    security: model Security {
        isTor: false
        isProxy: true
        isSpam: false
        isBot: false
    }
}
```

## Bulk IP Security Request
The SDK also supports bulk IP Security requests using the `getBulkIPSecurity()` method. All parameters like `fields`, `include`, and `excludes` can also be used in bulk requests.

```go
respSecurity, r, err := apiClient.IPSecurityAPI.GetBulkIpSecurityInfo(ctx).Ips([]string{"8.8.8.8", "asdasd"}).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.GetSecurityInfo`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		return
	}

	responseJson, _ := json.MarshalIndent(respSecurity, "", "  ")
	fmt.Println(string(responseJson))

	for _, item := range respSecurity {
		if item.IsError() {
			fmt.Println("❌", *item.Error.Message)
		} else {
			fmt.Println("✅", *item.SecurityResponse.Security.IsAnonymous)
		}
	}
```

## ASN API Examples

This section provides usage examples of the `getAsnDetails()` method from the SDK. These methods allow developers to retrieve detailed ASN-level network data either by ASN number or by IP address. Note that ASN API is only available in the Advanced subscription plans.

Refer to the [ASN API documentation](https://ipgeolocation.io/asn-api.html#documentation-overview) for a detailed list of supported fields and behaviors.

### Get ASN Information by IP Address

```go
respASN, r, err := apiClient.ASNLookupAPI.
		GetAsnInfo(ctx).Ip("8.8.8.8").
		Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ASNAPI.GetASNInfo`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		return
	}

	responseJson, _ := json.MarshalIndent(respASN, "", "  ")
	fmt.Println(string(responseJson))
```
Sample Response
```
model ASNResponse {
    ip: 8.8.8.8
    asn: model ASNDetails {
        asNumber: AS15169
        organization: Google LLC
        country: US
        asnName: GOOGLE
        type: BUSINESS
        domain: about.google
        dateAllocated: 
        allocationStatus: assigned
        numOfIpv4Routes: 983
        numOfIpv6Routes: 104
        rir: ARIN
    }
}
```
### Get ASN Information by ASN Number

```go
respASN, r, err := apiClient.ASNLookupAPI.
		GetAsnInfo(ctx).Asn(15169).
		Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ASNAPI.GetASNInfo`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		return
	}

	responseJson, _ := json.MarshalIndent(respASN, "", "  ")
	fmt.Println(string(responseJson))
```
Sample Response
```
model ASNResponse {
    asn: model ASNDetails {
        asNumber: AS15169
        organization: Google LLC
        country: US
        asnName: GOOGLE
        type: BUSINESS
        domain: about.google
        dateAllocated: 
        allocationStatus: assigned
        numOfIpv4Routes: 983
        numOfIpv6Routes: 104
        rir: ARIN
    }
}
```

### Combine All objects using Include
```go
respASN, r, err := apiClient.ASNLookupAPI.
		GetAsnInfo(ctx).Asn(12).
		Include("peers,downstreams,upstreams,routes,whois_response").
		Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ASNAPI.GetASNInfo`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		return
	}

	responseJson, _ := json.MarshalIndent(respASN, "", "  ")
	fmt.Println(string(responseJson))
```
Sample Response
```
model ASNResponse {
    ip: null
    asn: model ASNDetails {
        asNumber: AS12
        organization: New York University
        country: US
        asnName: NYU-DOMAIN
        type: EDUCATION
        domain: nyu.edu
        dateAllocated: 
        allocationStatus: assigned
        numOfIpv4Routes: 11
        numOfIpv6Routes: 1
        rir: ARIN
        routes: [192.76.177.0/24, 216.165.96.0/20, 216.165.89.0/24, 216.165.0.0/18, 216.165.112.0/21, 128.122.0.0/16, 2607:f600::/32, 216.165.102.0/24, 216.165.64.0/19, 216.165.120.0/22, 192.86.139.0/24, 216.165.103.0/24]
        upstreams: [model ASNConnection {
            asNumber: AS3269
            description: Telecom Italia S.p.A.
            country: IT
        }, model ASNConnection {
            asNumber: AS8220
            description: COLT Technology Services Group Limited
            country: GB
        }, model ASNConnection {
            asNumber: AS286
            description: GTT Communications Inc.
            country: US
        }, model ASNConnection {
            asNumber: AS3257
            description: GTT Communications Inc.
            country: US
        }, model ASNConnection {
            asNumber: AS3754
            description: NYSERNet
            country: US
        }, model ASNConnection {
            asNumber: AS3356
            description: Level 3 Parent, LLC
            country: US
        }, model ASNConnection {
            asNumber: AS6461
            description: Zayo Bandwidth
            country: US
        }, model ASNConnection {
            asNumber: AS137
            description: Consortium GARR
            country: IT
        }]
        downstreams: [model ASNConnection {
            asNumber: AS394666
            description: NYU Langone Health
            country: US
        }, model ASNConnection {
            asNumber: AS54965
            description: Polytechnic Institute of NYU
            country: US
        }]
        peers: [model ASNConnection {
            asNumber: AS3269
            description: Telecom Italia S.p.A.
            country: IT
        }, model ASNConnection {
            asNumber: AS8220
            description: COLT Technology Services Group Limited
            country: GB
        }, model ASNConnection {
            asNumber: AS394666
            description: NYU Langone Health
            country: US
        }, model ASNConnection {
            asNumber: AS286
            description: GTT Communications Inc.
            country: NL
        }, model ASNConnection {
            asNumber: AS286
            description: GTT Communications Inc.
            country: US
        }, model ASNConnection {
            asNumber: AS3257
            description: GTT Communications Inc.
            country: US
        }, model ASNConnection {
            asNumber: AS3754
            description: NYSERNet
            country: US
        }, model ASNConnection {
            asNumber: AS3356
            description: Level 3 Parent, LLC
            country: US
        }, model ASNConnection {
            asNumber: AS6461
            description: Zayo Bandwidth
            country: US
        }, model ASNConnection {
            asNumber: AS137
            description: Consortium GARR
            country: IT
        }, model ASNConnection {
            asNumber: AS54965
            description: Polytechnic Institute of NYU
            country: US
        }]
        whoisResponse: 
        
        
        ASNumber:       12
        ASName:         NYU-DOMAIN
        ASHandle:       AS12
        RegDate:        1984-07-05
        Updated:        2023-05-25    
        Ref:            https://rdap.arin.net/registry/autnum/12
        
        
        OrgName:        New York University
        OrgId:          NYU-Z
        Address:        726 Broadway, 8th Floor - ITS
        City:           New York
        StateProv:      NY
        PostalCode:     10003
        Country:        US
        RegDate:        2023-05-15
        Updated:        2023-05-15
        Ref:            https://rdap.arin.net/registry/entity/NYU-Z
        
        
        OrgAbuseHandle: OIS9-ARIN
        OrgAbuseName:   Office of Information Security
        OrgAbusePhone:  +1-212-998-3333 
        OrgAbuseEmail:  abuse@nyu.edu
        OrgAbuseRef:    https://rdap.arin.net/registry/entity/OIS9-ARIN
        
        OrgNOCHandle: COSI-ARIN
        OrgNOCName:   Communications Operations Services - ITS
        OrgNOCPhone:  +1-212-998-3444 
        OrgNOCEmail:  noc-cosi-arin@nyu.edu
        OrgNOCRef:    https://rdap.arin.net/registry/entity/COSI-ARIN
        
        OrgTechHandle: COSI-ARIN
        OrgTechName:   Communications Operations Services - ITS
        OrgTechPhone:  +1-212-998-3444 
        OrgTechEmail:  noc-cosi-arin@nyu.edu
        OrgTechRef:    https://rdap.arin.net/registry/entity/COSI-ARIN
        
        RTechHandle: COSI-ARIN
        RTechName:   Communications Operations Services - ITS
        RTechPhone:  +1-212-998-3444 
        RTechEmail:  noc-cosi-arin@nyu.edu
        RTechRef:    https://rdap.arin.net/registry/entity/COSI-ARIN
        
        RNOCHandle: COSI-ARIN
        RNOCName:   Communications Operations Services - ITS
        RNOCPhone:  +1-212-998-3444 
        RNOCEmail:  noc-cosi-arin@nyu.edu
        RNOCRef:    https://rdap.arin.net/registry/entity/COSI-ARIN
               
    }
}
```

## Timezone API Examples

This section provides usage examples of the `GetTimezoneInfo()` method from the SDK, showcasing how to fetch timezone and time-related data using different query types — IP address, latitude/longitude, and timezone ID.

For full API specifications, refer to the [Timezone API documentation](https://ipgeolocation.io/timezone-api.html#documentation-overview).

### Get Timezone by IP Address

```go
respTimezone, r, err := apiClient.TimezoneAPI.GetTimezoneInfo(ctx).Ip("8.8.8.8").Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TimezoneAPI.GetTimezoneInfo`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		return
	}
	
	responseJson, _ := json.MarshalIndent(respTimezone, "", "  ")
	fmt.Println(string(responseJson))
```
Sample Response
```
model TimeZoneResponse {
    ip: 8.8.8.8
    location: model TimezoneLocation {
        continentCode: NA
        continentName: North America
        countryCode2: US
        countryCode3: USA
        countryName: United States
        countryNameOfficial: United States of America
        isEu: false
        stateProv: California
        stateCode: US-CA
        district: Santa Clara
        city: Mountain View
        locality: null
        zipcode: 94043-1351
        latitude: 37.42240
        longitude: -122.08421
    }
    timeZone: model TimezoneDetail {
        name: America/Los_Angeles
        offset: -8
        offsetWithDst: -7
        date: 2025-06-23
        dateTime: 2025-06-23 02:15:25
        dateTimeTxt: Monday, June 23, 2025 02:15:25
        dateTimeWti: Mon, 23 Jun 2025 02:15:25 -0700
        dateTimeYmd: 2025-06-23T02:15:25-0700
        dateTimeUnix: 1.750670125437E9
        time24: 02:15:25
        time12: 02:15:25 AM
        week: 26
        month: 6
        year: 2025
        yearAbbr: 25
        isDst: true
        dstSavings: 1
        dstExists: true
        dstStart: model TimezoneDetailDstStart {
            utcTime: 2025-03-09 TIME 10
            duration: +1H
            gap: true
            dateTimeAfter: 2025-03-09 TIME 03
            dateTimeBefore: 2025-03-09 TIME 02
            overlap: false
        }
        dstEnd: model TimezoneDetailDstEnd {
            utcTime: 2025-11-02 TIME 09
            duration: -1H
            gap: false
            dateTimeAfter: 2025-11-02 TIME 01
            dateTimeBefore: 2025-11-02 TIME 02
            overlap: true
        }
    }
}
```
### Get Timezone by Timezone Name

```go
respTimezone, r, err := apiClient.TimezoneAPI.GetTimezoneInfo(ctx).Tz("Europe/London").Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TimezoneAPI.GetTimezoneInfo`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		return
	}

	responseJson, _ := json.MarshalIndent(respTimezone, "", "  ")
	fmt.Println(string(responseJson))

```
Sample Response
``` 
model TimeZoneResponse {
    timeZone: model TimezoneDetail {
        name: Europe/London
        offset: 0
        offsetWithDst: 1
        date: 2025-06-23
        dateTime: 2025-06-23 10:25:01
        dateTimeTxt: Monday, June 23, 2025 10:25:01
        dateTimeWti: Mon, 23 Jun 2025 10:25:01 +0100
        dateTimeYmd: 2025-06-23T10:25:01+0100
        dateTimeUnix: 1.750670701706E9
        time24: 10:25:01
        time12: 10:25:01 AM
        week: 26
        month: 6
        year: 2025
        yearAbbr: 25
        isDst: true
        dstSavings: 1
        dstExists: true
        dstStart: model TimezoneDetailDstStart {
            utcTime: 2025-03-30 TIME 01
            duration: +1H
            gap: true
            dateTimeAfter: 2025-03-30 TIME 02
            dateTimeBefore: 2025-03-30 TIME 01
            overlap: false
        }
        dstEnd: model TimezoneDetailDstEnd {
            utcTime: 2025-10-26 TIME 01
            duration: -1H
            gap: false
            dateTimeAfter: 2025-10-26 TIME 01
            dateTimeBefore: 2025-10-26 TIME 02
            overlap: true
        }
    }
}
```
### Get Timezone from Any Address

```go
respTimezone, r, err := apiClient.TimezoneAPI.GetTimezoneInfo(ctx).Location("Munich, Germany").Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TimezoneAPI.GetTimezoneInfo`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		return
	}

	responseJson, _ := json.MarshalIndent(respTimezone, "", "  ")
	fmt.Println(string(responseJson))
```
Sample Response
```
model TimeZoneResponse {
    location: model TimezoneLocation {
        locationString: Munich, Germany
        countryName: Germany
        stateProv: Bavaria
        city: Munich
        locality: 
        latitude: 48.13711
        longitude: 11.57538
    }
    timeZone: model TimezoneDetail {
        name: Europe/Berlin
        offset: 1
        offsetWithDst: 2
        date: 2025-06-23
        dateTime: 2025-06-23 11:35:23
        dateTimeTxt: Monday, June 23, 2025 11:35:23
        dateTimeWti: Mon, 23 Jun 2025 11:35:23 +0200
        dateTimeYmd: 2025-06-23T11:35:23+0200
        dateTimeUnix: 1.750671323755E9
        time24: 11:35:23
        time12: 11:35:23 AM
        week: 26
        month: 6
        year: 2025
        yearAbbr: 25
        isDst: true
        dstSavings: 1
        dstExists: true
        dstStart: model TimezoneDetailDstStart {
            utcTime: 2025-03-30 TIME 01
            duration: +1H
            gap: true
            dateTimeAfter: 2025-03-30 TIME 03
            dateTimeBefore: 2025-03-30 TIME 02
            overlap: false
        }
        dstEnd: model TimezoneDetailDstEnd {
            utcTime: 2025-10-26 TIME 01
            duration: -1H
            gap: false
            dateTimeAfter: 2025-10-26 TIME 02
            dateTimeBefore: 2025-10-26 TIME 03
            overlap: true
        }
    }
}
```
### Get Timezone from Location Coordinates

```go
respTimezone, r, err := apiClient.TimezoneAPI.GetTimezoneInfo(ctx).Lat(48.8566).Long(2.3522).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TimezoneAPI.GetTimezoneInfo`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		return
	}

	responseJson, _ := json.MarshalIndent(respTimezone, "", "  ")
	fmt.Println(string(responseJson))
```
Sample Response
```
model TimeZoneResponse {
    timeZone: model TimezoneDetail {
        name: Europe/Paris
        offset: 1
        offsetWithDst: 2
        date: 2025-06-23
        dateTime: 2025-06-23 11:53:31
        dateTimeTxt: Monday, June 23, 2025 11:53:31
        dateTimeWti: Mon, 23 Jun 2025 11:53:31 +0200
        dateTimeYmd: 2025-06-23T11:53:31+0200
        dateTimeUnix: 1.750672411295E9
        time24: 11:53:31
        time12: 11:53:31 AM
        week: 26
        month: 6
        year: 2025
        yearAbbr: 25
        isDst: true
        dstSavings: 1
        dstExists: true
        dstStart: model TimezoneDetailDstStart {
            utcTime: 2025-03-30 TIME 01
            duration: +1H
            gap: true
            dateTimeAfter: 2025-03-30 TIME 03
            dateTimeBefore: 2025-03-30 TIME 02
            overlap: false
        }
        dstEnd: model TimezoneDetailDstEnd {
            utcTime: 2025-10-26 TIME 01
            duration: -1H
            gap: false
            dateTimeAfter: 2025-10-26 TIME 02
            dateTimeBefore: 2025-10-26 TIME 03
            overlap: true
        }
    }
}
```
### Get Timezone and Airport Details from IATA Code

```go
respTimezone, r, err := apiClient.TimezoneAPI.GetTimezoneInfo(ctx).IataCode("ZRH").Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TimezoneAPI.GetTimezoneInfo`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		return
	}

	responseJson, _ := json.MarshalIndent(respTimezone, "", "  ")
	fmt.Println(string(responseJson))
```
Sample Response
```
model TimeZoneResponse {
    airportDetails: model TimezoneAirport {
        type: large_airport
        name: Zurich Airport
        latitude: 47.45806
        longitude: 8.54806
        elevationFt: 1417
        continentCode: EU
        countryCode: CH
        stateCode: CH-ZH
        city: Zurich
        iataCode: ZRH
        icaoCode: LSZH
        faaCode: 
    }
    timeZone: model TimezoneDetail {
        name: Europe/Zurich
        offset: 1
        offsetWithDst: 2
        date: 2025-06-23
        dateTime: 2025-06-23 12:24:08
        dateTimeTxt: Monday, June 23, 2025 12:24:08
        dateTimeWti: Mon, 23 Jun 2025 12:24:08 +0200
        dateTimeYmd: 2025-06-23T12:24:08+0200
        dateTimeUnix: 1.750674248242E9
        time24: 12:24:08
        time12: 12:24:08 PM
        week: 26
        month: 6
        year: 2025
        yearAbbr: 25
        isDst: true
        dstSavings: 1
        dstExists: true
        dstStart: model TimezoneDetailDstStart {
            utcTime: 2025-03-30 TIME 01
            duration: +1H
            gap: true
            dateTimeAfter: 2025-03-30 TIME 03
            dateTimeBefore: 2025-03-30 TIME 02
            overlap: false
        }
        dstEnd: model TimezoneDetailDstEnd {
            utcTime: 2025-10-26 TIME 01
            duration: -1H
            gap: false
            dateTimeAfter: 2025-10-26 TIME 02
            dateTimeBefore: 2025-10-26 TIME 03
            overlap: true
        }
    }
}
```
Similarly, you can fetch Airport Details and Timezone from using any ICAO code as well

### Get Timezone and City Details from UN/LOCODE

```go
respTimezone, r, err := apiClient.TimezoneAPI.GetTimezoneInfo(ctx).LoCode("ESBCN").Execute()
if err != nil {
	fmt.Fprintf(os.Stderr, "Error when calling `TimezoneAPI.GetTimezoneInfo`: %v\n", err)
	fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	return
}

responseJson, _ := json.MarshalIndent(respTimezone, "", "  ")
fmt.Println(string(responseJson))
```
Sample Response
```
model TimeZoneResponse {
    loCodeDetails: model TimezoneLocode {
        loCode: ESBCN
        city: Barcelona
        stateCode: 
        countryCode: ES
        countryName: 
        locationType: Port, Rail Terminal, Road Terminal, Airport, Postal Exchange
        latitude: 41.38289
        longitude: 2.17743
    }
    timeZone: model TimezoneDetail {
        name: Europe/Madrid
        offset: 1
        offsetWithDst: 2
        date: 2025-06-23
        dateTime: 2025-06-23 12:32:55
        dateTimeTxt: Monday, June 23, 2025 12:32:55
        dateTimeWti: Mon, 23 Jun 2025 12:32:55 +0200
        dateTimeYmd: 2025-06-23T12:32:55+0200
        dateTimeUnix: 1.750674775033E9
        time24: 12:32:55
        time12: 12:32:55 PM
        week: 26
        month: 6
        year: 2025
        yearAbbr: 25
        isDst: true
        dstSavings: 1
        dstExists: true
        dstStart: model TimezoneDetailDstStart {
            utcTime: 2025-03-30 TIME 01
            duration: +1H
            gap: true
            dateTimeAfter: 2025-03-30 TIME 03
            dateTimeBefore: 2025-03-30 TIME 02
            overlap: false
        }
        dstEnd: model TimezoneDetailDstEnd {
            utcTime: 2025-10-26 TIME 01
            duration: -1H
            gap: false
            dateTimeAfter: 2025-10-26 TIME 02
            dateTimeBefore: 2025-10-26 TIME 03
            overlap: true
        }
    }
}
```
## Timezone Converter Examples

This section provides usage examples of the `ConvertTimeBetweenTimezones()` method from the SDK. The Timezone Converter API allows you to convert a specific time from one timezone to another using timezone identifiers and optional date/time inputs.

For more details, refer to official documentation: [Timezone Converter API](https://ipgeolocation.io/timezone-api.html#convert-time-bw-time-zones).

### Convert Current Time from One Timezone to Another

```go
respTimeConversion, res, err := apiClient.TimeConversionAPI.ConvertTimeBetweenTimezones(ctx).
		TzFrom("America/New_York").
		TzTo("Asia/Tokyo").
		Execute()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error calling `TimeConversionAPI.ConvertTimeBetweenTimezones`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", res)
		return
	}
	responseJson, _ := json.MarshalIndent(respTimeConversion, "", "  ")
	fmt.Println(string(responseJson))

```
Sample Response
```
model TimeConversionResponse {
    originalTime: 2024-12-08 11:00
    convertedTime: 2024-12-09 01:00:00
    diffHour: 13.0
    diffMin: 780
}
```
Similarly, you can convert time from any timezone to another timezone using location coordinates (Latitude and Longitude), location addresses, IATA codes, ICAO codes and UN/LOCODE.

## User Agent API Examples

This section provides usage examples of the `GetUserAgentDetails()` method from the SDK. The User Agent API extracts and modelifies information from user agent strings, including browser, engine, device, OS, and type metadata.

For full explanation, visit the [User Agent API documentation](https://ipgeolocation.io/user-agent-api.html#documentation-overview).

### Parse a Basic User Agent String

```go
respUserAgent, r, err := apiClient.UserAgentAPI.GetUserAgentDetails(ctx).UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36").Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAgentAPI.GetUserAgentDetails`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		return
	}

	responseJson, _ := json.MarshalIndent(respUserAgent, "", "  ")
	fmt.Println(string(responseJson))
```
Sample Response
```
model UserAgentData {
    userAgentString: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36
    name: Chrome
    type: Browser
    version: 125
    versionMajor: 125
    device: model UserAgentDataDevice {
        name: Desktop
        type: Desktop
        brand: Unknown
        cpu: Intel x86_64
    }
    engine: model UserAgentDataEngine {
        name: Blink
        type: Browser
        version: 125
        versionMajor: 125
    }
    operatingSystem: model UserAgentDataOperatingSystem {
        name: Windows NT
        type: Desktop
        version: ??
        versionMajor: ??
        build: ??
    }
}
```
If you don't pass any userAgentString, the API will return the data of device's user agent.

## Bulk User Agent Parsing Example

The SDK also supports bulk User Agent parsing using the `ParseBulkUserAgentStrings()` method. This allows parsing multiple user agent strings in a single request. All fields available in single-user-agent parsing are returned per entry.

```go
respBulkUserAgents, r, err := apiClient.UserAgentAPI.ParseBulkUserAgentStrings(ctx).UserAgents([]string{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_2) AppleWebKit/601.3.9 (KHTML, like Gecko) Version/9.0.2 Safari/601.3.9"}).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAgentAPI.ParseBulkUserAgentStrings`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		return
	}
	responseJson, _ := json.MarshalIndent(respBulkUserAgents, "", "  ")
	fmt.Println(string(responseJson))
```
## Astronomy API Examples

This section provides usage examples of the `GetAstronomyDetails()` method from the SDK, allowing developers to fetch sun and moon timings and position data based on coordinates, IP, or location string.

Refer to the [official Astronomy API documentation](https://ipgeolocation.io/astronomy-api.html#documentation-overview) for more details.

### Astronomy by Coordinates

```go
respAstronomy, r, err := apiClient.AstronomyAPI.GetAstronomyDetails(ctx).Lat("40.71280").Long("-74.00084").Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AstronomyAPI.GetAstronomyDetails`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		return
	}

	responseJson, _ := json.MarshalIndent(respAstronomy, "", "  ")
	fmt.Println(string(responseJson))
```
Sample Response
```
model AstronomyResponse {
    location: model AstronomyLocation {
        countryName: United States
        stateProv: New York
        city: New York
        locality:  
        latitude: 40.71280
        longitude: -74.00600
        elevation: 6.0
    }
    astronomy: model Astronomy {
        date: 2025-07-22
        currentTime: 05:34:17.046
        midNight: 01:02
        nightEnd: 03:48
        morning: model AstronomyMorning {
            astronomicalTwilightBegin: 03:48
            astronomicalTwilightEnd: 04:32
            nauticalTwilightBegin: 04:32
            nauticalTwilightEnd: 05:12
            civilTwilightBegin: 05:12
            civilTwilightEnd: 05:43
            blueHourBegin: 04:59
            blueHourEnd: 05:24
            goldenHourBegin: 05:24
            goldenHourEnd: 06:23
        }
        sunrise: 05:43
        sunset: 20:21
        evening: model AstronomyEvening {
            goldenHourBegin: 19:41
            goldenHourEnd: 20:40
            blueHourBegin: 20:40
            blueHourEnd: 21:05
            civilTwilightBegin: 20:21
            civilTwilightEnd: 20:52
            nauticalTwilightBegin: 20:52
            nauticalTwilightEnd: 21:31
            astronomicalTwilightBegin: 21:31
            astronomicalTwilightEnd: 22:16
        }
        nightBegin: 22:16
        sunStatus: -
        solarNoon: 13:02
        dayLength: 14:37
        sunAltitude: -2.4240905951150817
        sunDistance: 152012050.75662628
        sunAzimuth: 60.53270916713848
        moonPhase: WANING_CRESCENT
        moonrise: 02:48
        moonset: 19:10
        moonStatus: -
        moonAltitude: 26.687264834949556
        moonDistance: 369857.6483476412
        moonAzimuth: 74.22460131532307
        moonParallacticAngle: -56.08124322972331
        moonIlluminationPercentage: -7.41
        moonAngle: 328.4181377849406
    }
}
```
### Astronomy by IP Address
```go
respAstronomy, r, err := apiClient.AstronomyAPI.GetAstronomyDetails(ctx).Ip("8.8.8.8").Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AstronomyAPI.GetAstronomyDetails`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		return
	}

	responseJson, _ := json.MarshalIndent(respAstronomy, "", "  ")
	fmt.Println(string(responseJson))
```
Sample Response
```
model AstronomyResponse {
    ip: 8.8.8.8
    location: model AstronomyLocation {
        continentCode: NA
        continentName: North America
        countryCode2: US
        countryCode3: USA
        countryName: United States
        countryNameOfficial: United States of America
        isEu: false
        stateProv: California
        stateCode: US-CA
        district: Santa Clara
        city: Mountain View
        locality: Charleston Terrace
        zipcode: 94043-1351
        latitude: 37.42240
        longitude: -122.08421
        elevation: 3.0
    }
    astronomy: model Astronomy {
        date: 2025-07-22
        currentTime: 02:36:01.027
        midNight: 01:15
        nightEnd: 04:18
        morning: model AstronomyMorning {
            astronomicalTwilightBegin: 04:18
            astronomicalTwilightEnd: 04:58
            nauticalTwilightBegin: 04:58
            nauticalTwilightEnd: 05:35
            civilTwilightBegin: 05:35
            civilTwilightEnd: 06:04
            blueHourBegin: 05:23
            blueHourEnd: 05:47
            goldenHourBegin: 05:47
            goldenHourEnd: 06:42
        }
        sunrise: 06:04
        sunset: 20:24
        evening: model AstronomyEvening {
            goldenHourBegin: 19:46
            goldenHourEnd: 20:42
            blueHourBegin: 20:42
            blueHourEnd: 21:05
            civilTwilightBegin: 20:24
            civilTwilightEnd: 20:54
            nauticalTwilightBegin: 20:54
            nauticalTwilightEnd: 21:30
            astronomicalTwilightBegin: 21:30
            astronomicalTwilightEnd: 22:10
        }
        nightBegin: 22:10
        sunStatus: -
        solarNoon: 13:14
        dayLength: 14:20
        sunAltitude: -29.312204242565592
        sunDistance: 152012050.7566263
        sunAzimuth: 21.915241201213632
        moonPhase: WANING_CRESCENT
        moonrise: 03:23
        moonset: 19:16
        moonStatus: -
        moonAltitude: -6.780866431657464
        moonDistance: 369859.5847016905
        moonAzimuth: 45.928379972251605
        moonParallacticAngle: -40.47546867785306
        moonIlluminationPercentage: -7.40
        moonAngle: 328.43423626935555
    }
}
```

### Astronomy by Location String
```go
respAstronomy, r, err := apiClient.AstronomyAPI.GetAstronomyDetails(ctx).Location("Milan, Italy").Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AstronomyAPI.GetAstronomyDetails`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		return
	}

	responseJson, _ := json.MarshalIndent(respAstronomy, "", "  ")
	fmt.Println(string(responseJson))
```
Sample Response
```
model AstronomyResponse {
    location: model AstronomyLocation {
        locationString: Milan, Italy
        countryName: Italy
        stateProv: Lombardy
        city: Milan
        locality: 
        latitude: 45.46419
        longitude: 9.18963
        elevation: 122.0
    }
    astronomy: model Astronomy {
        date: 2025-07-22
        currentTime: 11:37:28.787
        midNight: 01:29
        nightEnd: 03:39
        morning: model AstronomyMorning {
            astronomicalTwilightBegin: 03:39
            astronomicalTwilightEnd: 04:35
            nauticalTwilightBegin: 04:35
            nauticalTwilightEnd: 05:21
            civilTwilightBegin: 05:21
            civilTwilightEnd: 05:54
            blueHourBegin: 05:06
            blueHourEnd: 05:35
            goldenHourBegin: 05:35
            goldenHourEnd: 06:40
        }
        sunrise: 05:54
        sunset: 21:04
        evening: model AstronomyEvening {
            goldenHourBegin: 20:19
            goldenHourEnd: 21:24
            blueHourBegin: 21:24
            blueHourEnd: 21:52
            civilTwilightBegin: 21:04
            civilTwilightEnd: 21:38
            nauticalTwilightBegin: 21:38
            nauticalTwilightEnd: 22:23
            astronomicalTwilightBegin: 22:23
            astronomicalTwilightEnd: 23:18
        }
        nightBegin: 23:18
        sunStatus: -
        solarNoon: 13:29
        dayLength: 15:10
        sunAltitude: 55.76507063803926
        sunDistance: 152012050.7566263
        sunAzimuth: 128.26574664275847
        moonPhase: WANING_CRESCENT
        moonrise: 02:36
        moonset: 19:49
        moonStatus: -
        moonAltitude: 72.39158071193661
        moonDistance: 369861.22005060845
        moonAzimuth: 197.31311454833428
        moonParallacticAngle: 13.735730743087668
        moonIlluminationPercentage: -7.39
        moonAngle: 328.44782327106236
    }
}
```
### Astronomy for Specific Date
```go
respAstronomy, r, err := apiClient.AstronomyAPI.GetAstronomyDetails(ctx).Lat("-27.47").Long("153.03").Date("2025-01-01").Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AstronomyAPI.GetAstronomyDetails`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		return
	}

	responseJson, _ := json.MarshalIndent(respAstronomy, "", "  ")
	fmt.Println(string(responseJson))
```
Sample Response
```
model AstronomyResponse {
    location: model AstronomyLocation {
        countryName: Australia
        stateProv: Queensland
        city: Brisbane
        locality: Brisbane
        latitude: -27.47000
        longitude: 153.02000
        elevation: 
    }
    astronomy: model Astronomy {
        date: 2025-01-01
        currentTime: 19:45:17.561
        midNight: 23:51
        nightEnd: 03:24
        morning: model AstronomyMorning {
            astronomicalTwilightBegin: 03:24
            astronomicalTwilightEnd: 03:57
            nauticalTwilightBegin: 03:57
            nauticalTwilightEnd: 04:29
            civilTwilightBegin: 04:29
            civilTwilightEnd: 04:56
            blueHourBegin: 04:19
            blueHourEnd: 04:40
            goldenHourBegin: 04:40
            goldenHourEnd: 05:30
        }
        sunrise: 04:56
        sunset: 18:46
        evening: model AstronomyEvening {
            goldenHourBegin: 18:12
            goldenHourEnd: 19:02
            blueHourBegin: 19:02
            blueHourEnd: 19:23
            civilTwilightBegin: 18:46
            civilTwilightEnd: 19:13
            nauticalTwilightBegin: 19:13
            nauticalTwilightEnd: 19:45
            astronomicalTwilightBegin: 19:45
            astronomicalTwilightEnd: 20:18
        }
        nightBegin: 20:18
        sunStatus: -
        solarNoon: 11:51
        dayLength: 13:50
        sunAltitude: -12.059617608402677
        sunDistance: 147102938.88036567
        sunAzimuth: 235.897971324645
        moonPhase: NEW_MOON
        moonrise: 05:42
        moonset: 20:08
        moonStatus: -
        moonAltitude: 4.6701693782344345
        moonDistance: 380596.5823950267
        moonAzimuth: 244.56945849604378
        moonParallacticAngle: 118.21976701203934
        moonIlluminationPercentage: 2.49
        moonAngle: 18.156495178599695
    }
}
```

### Astronomy in Different Language
You can also get Astronomy Data in other languages as well. Only paid subscriptions can access this feature.
```go
respAstronomy, r, err := apiClient.AstronomyAPI.GetAstronomyDetails(ctx).Ip("1.1.1.1").Lang("fr").Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AstronomyAPI.GetAstronomyDetails`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		return
	}

	responseJson, _ := json.MarshalIndent(respAstronomy, "", "  ")
	fmt.Println(string(responseJson))
```
Sample Response
```
model AstronomyResponse {
    ip: 1.1.1.1
    location: model AstronomyLocation {
        continentCode: OC
        continentName: Océanie
        countryCode2: AU
        countryCode3: AUS
        countryName: Australie
        countryNameOfficial: 
        isEu: false
        stateProv: Queensland
        stateCode: AU-QLD
        district: Brisbane
        city: Brisbane Sud
        locality: 
        zipcode: 4101
        latitude: -27.47306
        longitude: 153.01421
        elevation: 
    }
    astronomy: model Astronomy {
        date: 2025-07-22
        currentTime: 19:54:32.920
        midNight: 23:54
        nightEnd: 05:13
        morning: model AstronomyMorning {
            astronomicalTwilightBegin: 05:13
            astronomicalTwilightEnd: 05:41
            nauticalTwilightBegin: 05:41
            nauticalTwilightEnd: 06:09
            civilTwilightBegin: 06:09
            civilTwilightEnd: 06:34
            blueHourBegin: 06:00
            blueHourEnd: 06:19
            goldenHourBegin: 06:19
            goldenHourEnd: 07:08
        }
        sunrise: 06:34
        sunset: 17:14
        evening: model AstronomyEvening {
            goldenHourBegin: 16:40
            goldenHourEnd: 17:29
            blueHourBegin: 17:29
            blueHourEnd: 17:49
            civilTwilightBegin: 17:14
            civilTwilightEnd: 17:39
            nauticalTwilightBegin: 17:39
            nauticalTwilightEnd: 18:07
            astronomicalTwilightBegin: 18:07
            astronomicalTwilightEnd: 18:35
        }
        nightBegin: 18:35
        sunStatus: -
        solarNoon: 11:54
        dayLength: 10:39
        sunAltitude: -35.15165719378359
        sunDistance: 152012050.75662628
        sunAzimuth: 276.2757088601843
        moonPhase: WANING_CRESCENT
        moonrise: 04:04
        moonset: 14:19
        moonStatus: -
        moonAltitude: -66.8771626746063
        moonDistance: 369880.37618917384
        moonAzimuth: 278.66762618741274
        moonParallacticAngle: 93.79636599869248
        moonIlluminationPercentage: -7.32
        moonAngle: 328.6063710418327
    }
}
```

## Abuse Contact API Examples
This section demonstrates how to use the `GetAbuseContactInfo()` method of the AbuseContact API. This API helps security teams, hosting providers, and compliance professionals quickly identify the correct abuse reporting contacts for any IPv4 or IPv6 address. You can retrieve data like the responsible organization, role, contact emails, phone numbers, and address to take appropriate mitigation action against abusive or malicious activity.
> **Note**: Abuse Contact API is only available in Advanced Plan

Refer to the official [Abuse Contact API documentation](https://ipgeolocation.io/ip-abuse-contact-api.html#documentation-overview) for details on all available fields.
### Lookup Abuse Contact by IP
```go
respAbuse, res, err := apiClient.AbuseContactAPI.GetAbuseContactInfo(ctx).
		Ip("1.0.0.0").
		Execute()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AbuseAPI.GetAbuseDetails` with full params: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", res)
		return
	}

	responseJson, _ := json.MarshalIndent(respAbuse, "", "  ")
	fmt.Println(string(responseJson))
```
Sample Response:
```
model AbuseResponse {
    ip: 1.0.0.0
    abuse: model Abuse {
        route: 1.0.0.0/24
        country: AU
        handle: IRT-APNICRANDNET-AU
        name: IRT-APNICRANDNET-AU
        organization: 
        role: abuse
        kind: group
        address: PO Box 3646
        South Brisbane, QLD 4101
        Australia
        emails: [helpdesk@apnic.net]
        phoneNumbers: [+61 7 3858 3100]
    }
}
```

### Lookup Abuse Contact with Specific Fields
```go
respAbuse, res, err := apiClient.AbuseContactAPI.GetAbuseContactInfo(ctx).
		Ip("1.2.3.4").
		Fields("abuse.role,abuse.emails").
		Execute()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AbuseAPI.GetAbuseDetails` with full params: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", res)
		return
	}

	responseJson, _ := json.MarshalIndent(respAbuse, "", "  ")
	fmt.Println(string(responseJson))
```
Sample Response:
```
model AbuseResponse {
    ip: 1.2.3.4
    abuse: model Abuse {
        role: abuse
        emails: [helpdesk@apnic.net]
    }
}
```
### Lookup Abuse Contact while Excluding Fields
```go
respAbuse, res, err := apiClient.AbuseContactAPI.GetAbuseContactInfo(ctx).
		Ip("9.9.9.9").
		Excludes("abuse.handle,abuse.emails").
		Execute()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AbuseAPI.GetAbuseDetails` with full params: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", res)
		return
	}

	responseJson, _ := json.MarshalIndent(respAbuse, "", "  ")
	fmt.Println(string(responseJson))
```
Sample Response:
```
model AbuseResponse {
    ip: 9.9.9.9
    abuse: model Abuse {
        route: 9.9.9.0/24
        country:
        name: Quad9 Abuse
        organization: Quad9 Abuse
        role: abuse
        kind: group
        address: 1442 A Walnut Street Ste 501
        Berkeley
        CA
        94709
        United States
        phoneNumbers: [+1-415-831-3129]
    }
}
```

## Documentation For Models

 - [ASNConnection](docs/ASNConnection.md)
 - [ASNResponse](docs/ASNResponse.md)
 - [ASNDetails](docs/ASNDetails.md)
 - [Abuse](docs/Abuse.md)
 - [AbuseResponse](docs/AbuseResponse.md)
 - [Astronomy](docs/Astronomy.md)
 - [AstronomyEvening](docs/AstronomyEvening.md)
 - [AstronomyLocation](docs/AstronomyLocation.md)
 - [AstronomyMorning](docs/AstronomyMorning.md)
 - [AstronomyResponse](docs/AstronomyResponse.md)
 - [BulkIpSecurity](docs/BulkIpSecurity.md)
 - [BulkIPGeolocation](docs/BulkIPGeolocation.md)
 - [CountryMetadata](docs/CountryMetadata.md)
 - [Currency](docs/Currency.md)
 - [ErrorResponse](docs/ErrorResponse.md)
 - [GeolocationResponse](docs/GeolocationResponse.md)
 - [GetBulkIpGeolocationRequest](docs/GetBulkIpGeolocationRequest.md)
 - [Location](docs/Location.md)
 - [LocationMinimal](docs/LocationMinimal.md)
 - [Network](docs/Network.md)
 - [NetworkAsn](docs/NetworkAsn.md)
 - [NetworkCompany](docs/NetworkCompany.md)
 - [NetworkMinimal](docs/NetworkMinimal.md)
 - [NetworkMinimalAsn](docs/NetworkMinimalAsn.md)
 - [NetworkMinimalCompany](docs/NetworkMinimalCompany.md)
 - [ParseBulkUserAgentStringsRequest](docs/ParseBulkUserAgentStringsRequest.md)
 - [ParseUserAgentStringRequest](docs/ParseUserAgentStringRequest.md)
 - [Security](docs/Security.md)
 - [SecurityAPIResponse](docs/SecurityAPIResponse.md)
 - [TimeConversionXMLResponse](docs/TimeConversionXMLResponse.md)
 - [TimeZone](docs/TimeZone.md)
 - [TimeZoneDetailedResponse](docs/TimeZoneDetailedResponse.md)
 - [TimeZoneDstEnd](docs/TimeZoneDstEnd.md)
 - [TimeZoneDstStart](docs/TimeZoneDstStart.md)
 - [TimezoneAirport](docs/TimezoneAirport.md)
 - [TimezoneDetail](docs/TimezoneDetail.md)
 - [TimezoneDetailDstEnd](docs/TimezoneDetailDstEnd.md)
 - [TimezoneDetailDstStart](docs/TimezoneDetailDstStart.md)
 - [TimezoneLocation](docs/TimezoneLocation.md)
 - [TimezoneLocode](docs/TimezoneLocode.md)
 - [UserAgentData](docs/UserAgentData.md)
 - [UserAgentDataDevice](docs/UserAgentDataDevice.md)
 - [UserAgentDataEngine](docs/UserAgentDataEngine.md)
 - [UserAgentDataOperatingSystem](docs/UserAgentDataOperatingSystem.md)




