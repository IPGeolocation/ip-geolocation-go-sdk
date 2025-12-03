# UserAgentAPI

All URIs are relative to *https://api.ipgeolocation.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUserAgentDetails**](UserAgentAPI.md#GetUserAgentDetails) | **Get** /user-agent | Get details of user-agent
[**ParseBulkUserAgentStrings**](UserAgentAPI.md#ParseBulkUserAgentStrings) | **Post** /user-agent-bulk | Handle multiple user-agent string lookups
[**ParseUserAgentString**](UserAgentAPI.md#ParseUserAgentString) | **Post** /user-agent | Handle single User-Agent string



## GetUserAgentDetails

> UserAgentData GetUserAgentDetails(ctx).UserAgent(userAgent).Output(output).Execute()

Get details of user-agent



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
	userAgent := "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_2) AppleWebKit/601.3.9 (KHTML, like Gecko) Version/9.0.2 Safari/601.3.9" // string |  (optional)
	output := "json" // string | Desired output format (json or xml). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAgentAPI.GetUserAgentDetails(context.Background()).UserAgent(userAgent).Output(output).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAgentAPI.GetUserAgentDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserAgentDetails`: UserAgentData
	fmt.Fprintf(os.Stdout, "Response from `UserAgentAPI.GetUserAgentDetails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserAgentDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userAgent** | **string** |  | 
 **output** | **string** | Desired output format (json or xml). | 

### Return type

[**UserAgentData**](UserAgentData.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ParseBulkUserAgentStrings

> []UserAgentData ParseBulkUserAgentStrings(ctx).Output(output).ParseBulkUserAgentStringsRequest(parseBulkUserAgentStringsRequest).Execute()

Handle multiple user-agent string lookups



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
	output := "json" // string | Desired output format (json or xml). (optional)
	parseBulkUserAgentStringsRequest := *openapiclient.NewParseBulkUserAgentStringsRequest() // ParseBulkUserAgentStringsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAgentAPI.ParseBulkUserAgentStrings(context.Background()).Output(output).ParseBulkUserAgentStringsRequest(parseBulkUserAgentStringsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAgentAPI.ParseBulkUserAgentStrings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ParseBulkUserAgentStrings`: []UserAgentData
	fmt.Fprintf(os.Stdout, "Response from `UserAgentAPI.ParseBulkUserAgentStrings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiParseBulkUserAgentStringsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **output** | **string** | Desired output format (json or xml). | 
 **parseBulkUserAgentStringsRequest** | [**ParseBulkUserAgentStringsRequest**](ParseBulkUserAgentStringsRequest.md) |  | 

### Return type

[**[]UserAgentData**](UserAgentData.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ParseUserAgentString

> UserAgentData ParseUserAgentString(ctx).Output(output).ParseUserAgentStringRequest(parseUserAgentStringRequest).Execute()

Handle single User-Agent string



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
	output := "json" // string | Desired output format (json or xml). (optional)
	parseUserAgentStringRequest := *openapiclient.NewParseUserAgentStringRequest() // ParseUserAgentStringRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAgentAPI.ParseUserAgentString(context.Background()).Output(output).ParseUserAgentStringRequest(parseUserAgentStringRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAgentAPI.ParseUserAgentString``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ParseUserAgentString`: UserAgentData
	fmt.Fprintf(os.Stdout, "Response from `UserAgentAPI.ParseUserAgentString`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiParseUserAgentStringRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **output** | **string** | Desired output format (json or xml). | 
 **parseUserAgentStringRequest** | [**ParseUserAgentStringRequest**](ParseUserAgentStringRequest.md) |  | 

### Return type

[**UserAgentData**](UserAgentData.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

