# \TimezoneAPI

All URIs are relative to *https://api.ipgeolocation.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTimezoneInfo**](TimezoneAPI.md#GetTimezoneInfo) | **Get** /timezone | Timezone information details



## GetTimezoneInfo

> TimeZoneDetailedResponse GetTimezoneInfo(ctx).Tz(tz).Location(location).Lat(lat).Long(long).Ip(ip).IataCode(iataCode).IcaoCode(icaoCode).LoCode(loCode).Output(output).Lang(lang).Execute()

Timezone information details



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/ipgeolocationsdk"
)

func main() {
	tz := "America/Los_Angeles" // string | pass a valid time zone name as a query parameter tz to get the time zone information. (optional)
	location := "London, UK" // string | pass any address of a location as the query parameter location to get the time zone information. (optional)
	lat := float32(-27.4748) // float32 | pass the latitude of a location as query parameters to get the time zone information. (optional)
	long := float32(153.017) // float32 | pass the longitude of a location as query parameters to get the time zone information. (optional)
	ip := "1.1.1.1" // string | You can pass any IPv4 or IPv6 address as a query parameter ip to get the regional timezone information. (optional)
	iataCode := "DXB" // string | pass any 3 letter IATA code as a query paramter iata_code to get the comprehensive airport details along with the time zone information, in which that airport exists. (optional)
	icaoCode := "KATL" // string | pass any 4 letter ICAO code as a query paramter icao_code to get the comprehensive airport details along with the time zone information, in which that airport exists. (optional)
	loCode := "DEBER" // string | pass any 5 letter UNLOCODE as a query paramter lo_code to get the comprehensive lo code/city details along with the time zone information of the concerned city. (optional)
	output := "json" // string | Desired output format (json or xml). (optional)
	lang := "en" // string | By default, the API responds in English. You can change the response language by passing the language code as a query parameter `lang`. Multi language feature is available only for `paid users`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TimezoneAPI.GetTimezoneInfo(context.Background()).Tz(tz).Location(location).Lat(lat).Long(long).Ip(ip).IataCode(iataCode).IcaoCode(icaoCode).LoCode(loCode).Output(output).Lang(lang).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TimezoneAPI.GetTimezoneInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTimezoneInfo`: TimeZoneDetailedResponse
	fmt.Fprintf(os.Stdout, "Response from `TimezoneAPI.GetTimezoneInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTimezoneInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tz** | **string** | pass a valid time zone name as a query parameter tz to get the time zone information. | 
 **location** | **string** | pass any address of a location as the query parameter location to get the time zone information. | 
 **lat** | **float32** | pass the latitude of a location as query parameters to get the time zone information. | 
 **long** | **float32** | pass the longitude of a location as query parameters to get the time zone information. | 
 **ip** | **string** | You can pass any IPv4 or IPv6 address as a query parameter ip to get the regional timezone information. | 
 **iataCode** | **string** | pass any 3 letter IATA code as a query paramter iata_code to get the comprehensive airport details along with the time zone information, in which that airport exists. | 
 **icaoCode** | **string** | pass any 4 letter ICAO code as a query paramter icao_code to get the comprehensive airport details along with the time zone information, in which that airport exists. | 
 **loCode** | **string** | pass any 5 letter UNLOCODE as a query paramter lo_code to get the comprehensive lo code/city details along with the time zone information of the concerned city. | 
 **output** | **string** | Desired output format (json or xml). | 
 **lang** | **string** | By default, the API responds in English. You can change the response language by passing the language code as a query parameter &#x60;lang&#x60;. Multi language feature is available only for &#x60;paid users&#x60;. | 

### Return type

[**TimeZoneDetailedResponse**](TimeZoneDetailedResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

