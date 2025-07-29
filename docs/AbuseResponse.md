# AbuseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | Pointer to **string** |  | [optional] 
**Abuse** | Pointer to [**Abuse**](Abuse.md) |  | [optional] 

## Methods

### NewAbuseResponse

`func NewAbuseResponse() *AbuseResponse`

NewAbuseResponse instantiates a new AbuseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAbuseResponseWithDefaults

`func NewAbuseResponseWithDefaults() *AbuseResponse`

NewAbuseResponseWithDefaults instantiates a new AbuseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *AbuseResponse) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *AbuseResponse) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *AbuseResponse) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *AbuseResponse) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetAbuse

`func (o *AbuseResponse) GetAbuse() Abuse`

GetAbuse returns the Abuse field if non-nil, zero value otherwise.

### GetAbuseOk

`func (o *AbuseResponse) GetAbuseOk() (*Abuse, bool)`

GetAbuseOk returns a tuple with the Abuse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbuse

`func (o *AbuseResponse) SetAbuse(v Abuse)`

SetAbuse sets Abuse field to given value.

### HasAbuse

`func (o *AbuseResponse) HasAbuse() bool`

HasAbuse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


