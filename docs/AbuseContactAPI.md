# \AbuseContactAPI

All URIs are relative to *https://api.ipgeolocation.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAbuseContactInfo**](AbuseContactAPI.md#GetAbuseContactInfo) | **Get** /abuse | 



## GetAbuseContactInfo

> AbuseResponse GetAbuseContactInfo(ctx).Ip(ip).Excludes(excludes).Fields(fields).Execute()





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
	ip := "8.8.8.8" // string | query parameter 'ip'. (optional)
	excludes := "abuse.handle,abuse.emails" // string | You can exclude specific fields from the API response (except the ip field) by listing them in the excludes parameter as a comma-separated list. For example, you want to remove emails and handle from api response, you can put the keys in excludes parameter like this. (optional)
	fields := "abuse.role,abuse.emails" // string | You can customize the API response by using the fields parameter to include only the specific data you need. For example, to retrieve only the role and emails, specify these keys in the fields parameter as shown below. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AbuseContactAPI.GetAbuseContactInfo(context.Background()).Ip(ip).Excludes(excludes).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AbuseContactAPI.GetAbuseContactInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAbuseContactInfo`: AbuseResponse
	fmt.Fprintf(os.Stdout, "Response from `AbuseContactAPI.GetAbuseContactInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAbuseContactInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ip** | **string** | query parameter &#39;ip&#39;. | 
 **excludes** | **string** | You can exclude specific fields from the API response (except the ip field) by listing them in the excludes parameter as a comma-separated list. For example, you want to remove emails and handle from api response, you can put the keys in excludes parameter like this. | 
 **fields** | **string** | You can customize the API response by using the fields parameter to include only the specific data you need. For example, to retrieve only the role and emails, specify these keys in the fields parameter as shown below. | 

### Return type

[**AbuseResponse**](AbuseResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

