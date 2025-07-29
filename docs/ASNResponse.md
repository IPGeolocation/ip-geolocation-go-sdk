# ASNResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | Pointer to **string** | IP address for which ASN information is returned. Present if the &#39;ip&#39; query parameter is used or no IP is specified (defaults to requester&#39;s IP). | [optional] 
**Asn** | Pointer to [**ASNResponseAsn**](ASNResponseAsn.md) |  | [optional] 

## Methods

### NewASNResponse

`func NewASNResponse() *ASNResponse`

NewASNResponse instantiates a new ASNResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewASNResponseWithDefaults

`func NewASNResponseWithDefaults() *ASNResponse`

NewASNResponseWithDefaults instantiates a new ASNResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *ASNResponse) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *ASNResponse) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *ASNResponse) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *ASNResponse) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetAsn

`func (o *ASNResponse) GetAsn() ASNResponseAsn`

GetAsn returns the Asn field if non-nil, zero value otherwise.

### GetAsnOk

`func (o *ASNResponse) GetAsnOk() (*ASNResponseAsn, bool)`

GetAsnOk returns a tuple with the Asn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsn

`func (o *ASNResponse) SetAsn(v ASNResponseAsn)`

SetAsn sets Asn field to given value.

### HasAsn

`func (o *ASNResponse) HasAsn() bool`

HasAsn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


