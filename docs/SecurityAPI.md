# \SecurityAPI

All URIs are relative to *https://api.ipgeolocation.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBulkIpSecurityInfo**](SecurityAPI.md#GetBulkIpSecurityInfo) | **Post** /security-bulk | 
[**GetIpSecurityInfo**](SecurityAPI.md#GetIpSecurityInfo) | **Get** /security | 



## GetBulkIpSecurityInfo

> []GetBulkIpSecurityInfo200ResponseInner GetBulkIpSecurityInfo(ctx).GetBulkIpGeolocationRequest(getBulkIpGeolocationRequest).Include(include).Fields(fields).Excludes(excludes).Output(output).Lang(lang).Execute()





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
	include := "location,network,currency,time_zone,user_agent,country_metadata,hostname" // string | Include optional objects like `location`, `network`.  Use comma-separated values. Example: include=location,network  (optional)
	fields := "security.is_tor, location," // string | Get specific fields, objects from the response. (optional)
	excludes := "security.is_cloud_provider" // string | Exclude specific fields, objects from the response. (optional)
	output := "json" // string | Desired output format. (optional)
	lang := "en" // string | By default, the API responds in English. You can change the response language by passing the language code as a query parameter `lang`. Multi language feature is available only for `paid users`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.GetBulkIpSecurityInfo(context.Background()).GetBulkIpGeolocationRequest(getBulkIpGeolocationRequest).Include(include).Fields(fields).Excludes(excludes).Output(output).Lang(lang).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.GetBulkIpSecurityInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBulkIpSecurityInfo`: []GetBulkIpSecurityInfo200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.GetBulkIpSecurityInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBulkIpSecurityInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getBulkIpGeolocationRequest** | [**GetBulkIpGeolocationRequest**](GetBulkIpGeolocationRequest.md) |  | 
 **include** | **string** | Include optional objects like &#x60;location&#x60;, &#x60;network&#x60;.  Use comma-separated values. Example: include&#x3D;location,network  | 
 **fields** | **string** | Get specific fields, objects from the response. | 
 **excludes** | **string** | Exclude specific fields, objects from the response. | 
 **output** | **string** | Desired output format. | 
 **lang** | **string** | By default, the API responds in English. You can change the response language by passing the language code as a query parameter &#x60;lang&#x60;. Multi language feature is available only for &#x60;paid users&#x60;. | 

### Return type

[**[]GetBulkIpSecurityInfo200ResponseInner**](GetBulkIpSecurityInfo200ResponseInner.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIpSecurityInfo

> SecurityAPIResponse GetIpSecurityInfo(ctx).Ip(ip).Include(include).Fields(fields).Excludes(excludes).Output(output).Lang(lang).Execute()





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
	ip := "8.8.8.8" // string | query parameter 'ip'. If not provided, will be your network IP (optional)
	include := "location,network,currency,time_zone,user_agent,country_metadata,hostname" // string | Include optional details like location and/or network. (optional)
	fields := "security.is_tor, location," // string | Get specific fields, objects from the response. (optional)
	excludes := "security.is_cloud_provider" // string | Exclude specific fields, objects from the response. (optional)
	output := "json" // string | Desired output format (json or xml). (optional)
	lang := "en" // string | By default, the API responds in English. You can change the response language by passing the language code as a query parameter `lang`. Multi language feature is available only for `paid users`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.GetIpSecurityInfo(context.Background()).Ip(ip).Include(include).Fields(fields).Excludes(excludes).Output(output).Lang(lang).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.GetIpSecurityInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIpSecurityInfo`: SecurityAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.GetIpSecurityInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIpSecurityInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ip** | **string** | query parameter &#39;ip&#39;. If not provided, will be your network IP | 
 **include** | **string** | Include optional details like location and/or network. | 
 **fields** | **string** | Get specific fields, objects from the response. | 
 **excludes** | **string** | Exclude specific fields, objects from the response. | 
 **output** | **string** | Desired output format (json or xml). | 
 **lang** | **string** | By default, the API responds in English. You can change the response language by passing the language code as a query parameter &#x60;lang&#x60;. Multi language feature is available only for &#x60;paid users&#x60;. | 

### Return type

[**SecurityAPIResponse**](SecurityAPIResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

