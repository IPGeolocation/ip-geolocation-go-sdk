# BulkIPGeolocation

## Output 

### GeolocationResponse (Successful)

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | Pointer to **string** |  | [optional] 
**Hostname** | Pointer to **string** |  | [optional] 
**Domain** | Pointer to **string** |  | [optional] 
**Location** | Pointer to [**Location**](Location.md) |  | [optional] 
**CountryMetadata** | Pointer to [**CountryMetadata**](CountryMetadata.md) |  | [optional] 
**Network** | Pointer to [**Network**](Network.md) |  | [optional] 
**Currency** | Pointer to [**Currency**](Currency.md) |  | [optional] 
**Security** | Pointer to [**Security**](Security.md) |  | [optional] 
**Abuse** | Pointer to [**Abuse**](Abuse.md) |  | [optional] 
**TimeZone** | Pointer to [**TimeZone**](TimeZone.md) |  | [optional] 
**UserAgent** | Pointer to [**UserAgentData**](UserAgentData.md) |  | [optional] 

### Error Response (if any) 

| Name        | Type                                                | Description                  | Notes       |
|-------------|-----------------------------------------------------|------------------------------|-------------|
| **Message** | Pointer to **string**                               | Error message                | [optional]  |
---


## Usage Examples

### ✅ Checking if a result is successful

```go
for _, result := range bulkResults {
    if result.IsSuccess() {
        fmt.Println("Success:", *result.GeolocationResponse.Ip)
    }
}
```

### 🚨 Checking if a result is an error

```go
for _, result := range bulkResults {
    if result.IsError() {
        fmt.Println("Error:", *result.Error.Message)
    }
}
```


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


