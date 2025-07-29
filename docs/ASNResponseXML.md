# ASNResponseXML

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | Pointer to **string** |  | [optional] 
**Asn** | Pointer to [**ASNResponseXMLAsn**](ASNResponseXMLAsn.md) |  | [optional] 

## Methods

### NewASNResponseXML

`func NewASNResponseXML() *ASNResponseXML`

NewASNResponseXML instantiates a new ASNResponseXML object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewASNResponseXMLWithDefaults

`func NewASNResponseXMLWithDefaults() *ASNResponseXML`

NewASNResponseXMLWithDefaults instantiates a new ASNResponseXML object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *ASNResponseXML) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *ASNResponseXML) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *ASNResponseXML) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *ASNResponseXML) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetAsn

`func (o *ASNResponseXML) GetAsn() ASNResponseXMLAsn`

GetAsn returns the Asn field if non-nil, zero value otherwise.

### GetAsnOk

`func (o *ASNResponseXML) GetAsnOk() (*ASNResponseXMLAsn, bool)`

GetAsnOk returns a tuple with the Asn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsn

`func (o *ASNResponseXML) SetAsn(v ASNResponseXMLAsn)`

SetAsn sets Asn field to given value.

### HasAsn

`func (o *ASNResponseXML) HasAsn() bool`

HasAsn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


