# BulkIpSecurity

Represents the unified response for IP Security lookups. May contain either full IP security details or an error message, depending on API outcome.

## Output

### Security Response

| Name                | Type                                                | Description                  | Notes       |
|---------------------|-----------------------------------------------------|------------------------------|-------------|
| **Ip**              | Pointer to **string**                               | IP address queried           | [optional]  |
| **Hostname**        | Pointer to **string**                               | Hostname associated          | [optional]  |
| **Security**        | Pointer to [**Security**](Security.md)              | Security information         | [optional]  |
| **Location**        | Pointer to [**LocationMinimal**](LocationMinimal.md)| Minimal location data        | [optional]  |
| **Network**         | Pointer to [**NetworkMinimal**](NetworkMinimal.md)  | Minimal network data         | [optional]  |
| **TimeZone**        | Pointer to [**TimeZone**](TimeZone.md)              | Time zone of IP              | [optional]  |
| **UserAgent**       | Pointer to [**UserAgentData**](UserAgentData.md)    | User agent metadata          | [optional]  |
| **CountryMetadata** | Pointer to [**CountryMetadata**](CountryMetadata.md)| Country metadata             | [optional]  |
| **Currency**        | Pointer to [**Currency**](Currency.md)              | Currency data                | [optional]  |
---

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
        fmt.Println("Success:", *result.SecurityResponse.Ip)
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


