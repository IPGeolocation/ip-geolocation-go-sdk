# \TimeConversionAPI

All URIs are relative to *https://api.ipgeolocation.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConvertTimeBetweenTimezones**](TimeConversionAPI.md#ConvertTimeBetweenTimezones) | **Get** /timezone/convert | 



## ConvertTimeBetweenTimezones

> TimeConversionResponse ConvertTimeBetweenTimezones(ctx).Time(time).TzFrom(tzFrom).TzTo(tzTo).LatFrom(latFrom).LongFrom(longFrom).LatTo(latTo).LongTo(longTo).LocationFrom(locationFrom).LocationTo(locationTo).IcaoFrom(icaoFrom).IcaoTo(icaoTo).IataFrom(iataFrom).IataTo(iataTo).LocodeFrom(locodeFrom).LocodeTo(locodeTo).Execute()





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
	time := "2025-02-28 9:00:00" // string | time parameter takes the input in the following two formats: i) 'yyyy-MM-dd HH:mm', and ii) 'yyyy-MM-dd HH:mm:ss'. This parameter is optional and you can omit it to convert the current time between two coordinates, time zones or locations. (optional)
	tzFrom := "America/Argentina/Catamarca" // string | timezone to convert from (optional)
	tzTo := "Asia/Kabul" // string | timezone to convert to (optional)
	latFrom := float32(34.0207305) // float32 | latitude to convert from (optional)
	longFrom := float32(-118.691916) // float32 | longitude to convert from (optional)
	latTo := float32(53.473682) // float32 | latitude to convert to (optional)
	longTo := float32(-77.397706) // float32 | longitude to convert to (optional)
	locationFrom := "New York, USA" // string | location to convert from (optional)
	locationTo := "Lahore, Pakistan" // string | location to convert to (optional)
	icaoFrom := "YSSY" // string | location to convert from (optional)
	icaoTo := "ZBAA" // string | location to convert to (optional)
	iataFrom := "DXB" // string | location to convert from (optional)
	iataTo := "LHR" // string | location to convert to (optional)
	locodeFrom := "PKISB" // string | location to convert from (optional)
	locodeTo := "USNYC" // string | location to convert to (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TimeConversionAPI.ConvertTimeBetweenTimezones(context.Background()).Time(time).TzFrom(tzFrom).TzTo(tzTo).LatFrom(latFrom).LongFrom(longFrom).LatTo(latTo).LongTo(longTo).LocationFrom(locationFrom).LocationTo(locationTo).IcaoFrom(icaoFrom).IcaoTo(icaoTo).IataFrom(iataFrom).IataTo(iataTo).LocodeFrom(locodeFrom).LocodeTo(locodeTo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TimeConversionAPI.ConvertTimeBetweenTimezones``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConvertTimeBetweenTimezones`: TimeConversionResponse
	fmt.Fprintf(os.Stdout, "Response from `TimeConversionAPI.ConvertTimeBetweenTimezones`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConvertTimeBetweenTimezonesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **time** | **string** | time parameter takes the input in the following two formats: i) &#39;yyyy-MM-dd HH:mm&#39;, and ii) &#39;yyyy-MM-dd HH:mm:ss&#39;. This parameter is optional and you can omit it to convert the current time between two coordinates, time zones or locations. | 
 **tzFrom** | **string** | timezone to convert from | 
 **tzTo** | **string** | timezone to convert to | 
 **latFrom** | **float32** | latitude to convert from | 
 **longFrom** | **float32** | longitude to convert from | 
 **latTo** | **float32** | latitude to convert to | 
 **longTo** | **float32** | longitude to convert to | 
 **locationFrom** | **string** | location to convert from | 
 **locationTo** | **string** | location to convert to | 
 **icaoFrom** | **string** | location to convert from | 
 **icaoTo** | **string** | location to convert to | 
 **iataFrom** | **string** | location to convert from | 
 **iataTo** | **string** | location to convert to | 
 **locodeFrom** | **string** | location to convert from | 
 **locodeTo** | **string** | location to convert to | 

### Return type

[**TimeConversionResponse**](TimeConversionResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

