# \IPLocationAPI

All URIs are relative to *https://api.ipgeolocation.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBulkIpGeolocation**](IPLocationAPI.md#GetBulkIpGeolocation) | **Post** /ipgeo-bulk | 
[**GetIpGeolocation**](IPLocationAPI.md#GetIpGeolocation) | **Get** /ipgeo | 



## GetBulkIpGeolocation

> []GetBulkIpGeolocation200ResponseInner GetBulkIpGeolocation(ctx).GetBulkIpGeolocationRequest(getBulkIpGeolocationRequest).Lang(lang).Fields(fields).Excludes(excludes).Include(include).Output(output).Execute()





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
	getBulkIpGeolocationRequest := *openapiclient.NewGetBulkIpGeolocationRequest() // GetBulkIpGeolocationRequest | 
	lang := "en" // string | By default, the API responds in English. You can change the response language by passing the language code as a query parameter `lang`. Multi language feature is available only for `paid users`. (optional)
	fields := "location, location.country_name, network.asn.organization" // string | you can filter the API response by specifying the fields that you need, instead of getting the full response. To do this, pass the desired field names using the `fields` query parameter with each field represented as a dot-separated path. (optional)
	excludes := "location.continent_code,currency,network" // string | you can also filter the API response by excluding specific fields (except the IP address) that you don't need. To do this, pass the unwanted field names using the excludes query parameter, with each field represented as a dot-separated path (optional)
	include := "security" // string | IP Geolocation API also provides IP-Security, abuse, timezone, user-agent and DMA (Designated Market Area) code, which is specifically used in the US for marketing and regional targeting information on Advanced API subscription, but doesn't respond it by default. To get these information along with the geolocation information, you must pass the `include=security` or `include=abuse` or `include=dma` or `include=time_zone` or `include=user-agent` or you can fetch multiples by adding values in comma-separated way. In addition to that, IPGeolocation API also provide hostname lookup for an IP address on all the paid API subscriptions (STANDARD and ADVANCED), but doesn't respond it by default. To get the hostname for an IP address, you can pass one of the three values `hostname, liveHostname, hostnameFallbackLive` as a URL parameter `include=`. (optional)
	output := "json" // string | Desired output format(json or xml). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IPLocationAPI.GetBulkIpGeolocation(context.Background()).GetBulkIpGeolocationRequest(getBulkIpGeolocationRequest).Lang(lang).Fields(fields).Excludes(excludes).Include(include).Output(output).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPLocationAPI.GetBulkIpGeolocation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBulkIpGeolocation`: []GetBulkIpGeolocation200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `IPLocationAPI.GetBulkIpGeolocation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBulkIpGeolocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getBulkIpGeolocationRequest** | [**GetBulkIpGeolocationRequest**](GetBulkIpGeolocationRequest.md) |  | 
 **lang** | **string** | By default, the API responds in English. You can change the response language by passing the language code as a query parameter &#x60;lang&#x60;. Multi language feature is available only for &#x60;paid users&#x60;. | 
 **fields** | **string** | you can filter the API response by specifying the fields that you need, instead of getting the full response. To do this, pass the desired field names using the &#x60;fields&#x60; query parameter with each field represented as a dot-separated path. | 
 **excludes** | **string** | you can also filter the API response by excluding specific fields (except the IP address) that you don&#39;t need. To do this, pass the unwanted field names using the excludes query parameter, with each field represented as a dot-separated path | 
 **include** | **string** | IP Geolocation API also provides IP-Security, abuse, timezone, user-agent and DMA (Designated Market Area) code, which is specifically used in the US for marketing and regional targeting information on Advanced API subscription, but doesn&#39;t respond it by default. To get these information along with the geolocation information, you must pass the &#x60;include&#x3D;security&#x60; or &#x60;include&#x3D;abuse&#x60; or &#x60;include&#x3D;dma&#x60; or &#x60;include&#x3D;time_zone&#x60; or &#x60;include&#x3D;user-agent&#x60; or you can fetch multiples by adding values in comma-separated way. In addition to that, IPGeolocation API also provide hostname lookup for an IP address on all the paid API subscriptions (STANDARD and ADVANCED), but doesn&#39;t respond it by default. To get the hostname for an IP address, you can pass one of the three values &#x60;hostname, liveHostname, hostnameFallbackLive&#x60; as a URL parameter &#x60;include&#x3D;&#x60;. | 
 **output** | **string** | Desired output format(json or xml). | 

### Return type

[**[]GetBulkIpGeolocation200ResponseInner**](GetBulkIpGeolocation200ResponseInner.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIpGeolocation

> GeolocationResponse GetIpGeolocation(ctx).Ip(ip).Lang(lang).Fields(fields).Excludes(excludes).Include(include).Output(output).Execute()





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
	ip := "8.8.8.8 OR dns.google.com" // string | In order to find geolocation information about an IP (ipv4 ipv6) address or a domain name, pass it as a query parameter ip. When this endpoint is queried without an IP address, it returns the geolocation information of the device/client which is calling it (optional)
	lang := "en" // string | By default, the API responds in English. You can change the response language by passing the language code as a query parameter `lang`. Multi language feature is available only for `paid users`. (optional)
	fields := "location, location.country_name, network.asn.organization" // string | you can filter the API response by specifying the fields that you need, instead of getting the full response. To do this, pass the desired field names using the `fields` query parameter with each field represented as a dot-separated path. (optional)
	excludes := "location.continent_code,currency,network" // string | you can also filter the API response by excluding specific fields (except the IP address) that you don't need. To do this, pass the unwanted field names using the excludes query parameter, with each field represented as a dot-separated path (optional)
	include := "security" // string | IP Geolocation API also provides IP-Security, abuse, timezone, user-agent and DMA (Designated Market Area) code, which is specifically used in the US for marketing and regional targeting information on Advanced API subscription, but doesn't respond it by default. To get these information along with the geolocation information, you must pass the `include=security` or `include=abuse` or `include=dma` or `include=time_zone` or `include=user-agent` or you can fetch multiples by adding values in comma-separated way. In addition to that, IPGeolocation API also provide hostname lookup for an IP address on all the paid API subscriptions (STANDARD and ADVANCED), but doesn't respond it by default. To get the hostname for an IP address, you can pass one of the three values `hostname, liveHostname, hostnameFallbackLive` as a URL parameter `include=`. (optional)
	output := "json" // string | Desired output format (json or xml). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IPLocationAPI.GetIpGeolocation(context.Background()).Ip(ip).Lang(lang).Fields(fields).Excludes(excludes).Include(include).Output(output).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPLocationAPI.GetIpGeolocation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIpGeolocation`: GeolocationResponse
	fmt.Fprintf(os.Stdout, "Response from `IPLocationAPI.GetIpGeolocation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIpGeolocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ip** | **string** | In order to find geolocation information about an IP (ipv4 ipv6) address or a domain name, pass it as a query parameter ip. When this endpoint is queried without an IP address, it returns the geolocation information of the device/client which is calling it | 
 **lang** | **string** | By default, the API responds in English. You can change the response language by passing the language code as a query parameter &#x60;lang&#x60;. Multi language feature is available only for &#x60;paid users&#x60;. | 
 **fields** | **string** | you can filter the API response by specifying the fields that you need, instead of getting the full response. To do this, pass the desired field names using the &#x60;fields&#x60; query parameter with each field represented as a dot-separated path. | 
 **excludes** | **string** | you can also filter the API response by excluding specific fields (except the IP address) that you don&#39;t need. To do this, pass the unwanted field names using the excludes query parameter, with each field represented as a dot-separated path | 
 **include** | **string** | IP Geolocation API also provides IP-Security, abuse, timezone, user-agent and DMA (Designated Market Area) code, which is specifically used in the US for marketing and regional targeting information on Advanced API subscription, but doesn&#39;t respond it by default. To get these information along with the geolocation information, you must pass the &#x60;include&#x3D;security&#x60; or &#x60;include&#x3D;abuse&#x60; or &#x60;include&#x3D;dma&#x60; or &#x60;include&#x3D;time_zone&#x60; or &#x60;include&#x3D;user-agent&#x60; or you can fetch multiples by adding values in comma-separated way. In addition to that, IPGeolocation API also provide hostname lookup for an IP address on all the paid API subscriptions (STANDARD and ADVANCED), but doesn&#39;t respond it by default. To get the hostname for an IP address, you can pass one of the three values &#x60;hostname, liveHostname, hostnameFallbackLive&#x60; as a URL parameter &#x60;include&#x3D;&#x60;. | 
 **output** | **string** | Desired output format (json or xml). | 

### Return type

[**GeolocationResponse**](GeolocationResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

