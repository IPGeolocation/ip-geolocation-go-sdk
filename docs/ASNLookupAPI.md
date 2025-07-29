# \ASNLookupAPI

All URIs are relative to *https://api.ipgeolocation.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAsnInfo**](ASNLookupAPI.md#GetAsnInfo) | **Get** /asn | 



## GetAsnInfo

> ASNResponse GetAsnInfo(ctx).Ip(ip).Asn(asn).Include(include).Excludes(excludes).Fields(fields).Execute()





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
	asn := int32(1) // int32 | query paramter 'asn'. (optional)
	include := "peers" // string | This parameter can have four options: a) peers b) downstreams c) upstreams d) routes e) whois_response. You may add any of them in comma-separated way. In order to get the ASN details with peering data, pass peers string in the include parameter like mentioned below. (optional)
	excludes := "asn.date_allocated,asn.allocation_status" // string | You can exclude fields from the response according to you requirement with the exception of ip field. For example, you want to remove date_allocated and allocation_status from api response, you can put the keys in excludes parameter like this. (optional)
	fields := "asn.organization,asn.country,asn.downstreams" // string | You can filter out only those fields which you want to see in the response by using the fields parameter. To retrieve only the AS organization, its country and downstreams in api response, you can put the keys in fields parameter like this. API will combine these fields with the default ASN response. Note: Parameters `peers, downstreams, upstreams, routes, whois_response` can be used in both `include` , and `fields`. If you use include and fields at the same time, fields parameter will be considered only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ASNLookupAPI.GetAsnInfo(context.Background()).Ip(ip).Asn(asn).Include(include).Excludes(excludes).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ASNLookupAPI.GetAsnInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAsnInfo`: ASNResponse
	fmt.Fprintf(os.Stdout, "Response from `ASNLookupAPI.GetAsnInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAsnInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ip** | **string** | query parameter &#39;ip&#39;. | 
 **asn** | **int32** | query paramter &#39;asn&#39;. | 
 **include** | **string** | This parameter can have four options: a) peers b) downstreams c) upstreams d) routes e) whois_response. You may add any of them in comma-separated way. In order to get the ASN details with peering data, pass peers string in the include parameter like mentioned below. | 
 **excludes** | **string** | You can exclude fields from the response according to you requirement with the exception of ip field. For example, you want to remove date_allocated and allocation_status from api response, you can put the keys in excludes parameter like this. | 
 **fields** | **string** | You can filter out only those fields which you want to see in the response by using the fields parameter. To retrieve only the AS organization, its country and downstreams in api response, you can put the keys in fields parameter like this. API will combine these fields with the default ASN response. Note: Parameters &#x60;peers, downstreams, upstreams, routes, whois_response&#x60; can be used in both &#x60;include&#x60; , and &#x60;fields&#x60;. If you use include and fields at the same time, fields parameter will be considered only. | 

### Return type

[**ASNResponse**](ASNResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

