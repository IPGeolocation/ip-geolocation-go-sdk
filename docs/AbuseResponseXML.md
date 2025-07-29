# AbuseResponseXML

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | Pointer to **string** |  | [optional] 
**Abuse** | Pointer to [**Abuse**](Abuse.md) |  | [optional] 

## Methods

### NewAbuseResponseXML

`func NewAbuseResponseXML() *AbuseResponseXML`

NewAbuseResponseXML instantiates a new AbuseResponseXML object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAbuseResponseXMLWithDefaults

`func NewAbuseResponseXMLWithDefaults() *AbuseResponseXML`

NewAbuseResponseXMLWithDefaults instantiates a new AbuseResponseXML object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *AbuseResponseXML) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *AbuseResponseXML) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *AbuseResponseXML) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *AbuseResponseXML) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetAbuse

`func (o *AbuseResponseXML) GetAbuse() Abuse`

GetAbuse returns the Abuse field if non-nil, zero value otherwise.

### GetAbuseOk

`func (o *AbuseResponseXML) GetAbuseOk() (*Abuse, bool)`

GetAbuseOk returns a tuple with the Abuse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbuse

`func (o *AbuseResponseXML) SetAbuse(v Abuse)`

SetAbuse sets Abuse field to given value.

### HasAbuse

`func (o *AbuseResponseXML) HasAbuse() bool`

HasAbuse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


