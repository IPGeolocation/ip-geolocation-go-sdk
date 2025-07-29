# \AstronomyAPI

All URIs are relative to *https://api.ipgeolocation.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAstronomyDetails**](AstronomyAPI.md#GetAstronomyDetails) | **Get** /astronomy | 



## GetAstronomyDetails

> AstronomyResponse GetAstronomyDetails(ctx).Ip(ip).Location(location).Lat(lat).Long(long).Date(date).Elevation(elevation).Output(output).Lang(lang).Execute()





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
	ip := "8.8.8.8" // string | query paramter 'ip'. If not provided, will be your network IP (optional)
	location := "New York, US" // string | query paramter 'location'. If not provided, will be your ip location (optional)
	lat := "40.76473" // string | query paramter 'lat'. (optional)
	long := "-74.00084" // string | query paramter 'long'. (optional)
	date := "2025-04-22" // string | The date (YYYY-MM-DD) to lookup for. default will be the current date (optional)
	elevation := float64(9) // float64 | query parameter 'elevation' (optional)
	output := "json" // string | Desired output format. (optional)
	lang := "en" // string | By default, the API responds in English. You can change the response language by passing the language code as a query parameter `lang`. Multi language feature is available only for `paid users`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AstronomyAPI.GetAstronomyDetails(context.Background()).Ip(ip).Location(location).Lat(lat).Long(long).Date(date).Elevation(elevation).Output(output).Lang(lang).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AstronomyAPI.GetAstronomyDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAstronomyDetails`: AstronomyResponse
	fmt.Fprintf(os.Stdout, "Response from `AstronomyAPI.GetAstronomyDetails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAstronomyDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ip** | **string** | query paramter &#39;ip&#39;. If not provided, will be your network IP | 
 **location** | **string** | query paramter &#39;location&#39;. If not provided, will be your ip location | 
 **lat** | **string** | query paramter &#39;lat&#39;. | 
 **long** | **string** | query paramter &#39;long&#39;. | 
 **date** | **string** | The date (YYYY-MM-DD) to lookup for. default will be the current date | 
 **elevation** | **float64** | query parameter &#39;elevation&#39; | 
 **output** | **string** | Desired output format. | 
 **lang** | **string** | By default, the API responds in English. You can change the response language by passing the language code as a query parameter &#x60;lang&#x60;. Multi language feature is available only for &#x60;paid users&#x60;. | 

### Return type

[**AstronomyResponse**](AstronomyResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

