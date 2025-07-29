# Network

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asn** | Pointer to [**NetworkAsn**](NetworkAsn.md) |  | [optional] 
**ConnectionType** | Pointer to **string** |  | [optional] 
**Company** | Pointer to [**NetworkCompany**](NetworkCompany.md) |  | [optional] 

## Methods

### NewNetwork

`func NewNetwork() *Network`

NewNetwork instantiates a new Network object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkWithDefaults

`func NewNetworkWithDefaults() *Network`

NewNetworkWithDefaults instantiates a new Network object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsn

`func (o *Network) GetAsn() NetworkAsn`

GetAsn returns the Asn field if non-nil, zero value otherwise.

### GetAsnOk

`func (o *Network) GetAsnOk() (*NetworkAsn, bool)`

GetAsnOk returns a tuple with the Asn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsn

`func (o *Network) SetAsn(v NetworkAsn)`

SetAsn sets Asn field to given value.

### HasAsn

`func (o *Network) HasAsn() bool`

HasAsn returns a boolean if a field has been set.

### GetConnectionType

`func (o *Network) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *Network) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *Network) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.

### HasConnectionType

`func (o *Network) HasConnectionType() bool`

HasConnectionType returns a boolean if a field has been set.

### GetCompany

`func (o *Network) GetCompany() NetworkCompany`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *Network) GetCompanyOk() (*NetworkCompany, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *Network) SetCompany(v NetworkCompany)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *Network) HasCompany() bool`

HasCompany returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


