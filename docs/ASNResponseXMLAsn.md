# ASNResponseXMLAsn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AsNumber** | Pointer to **string** |  | [optional] 
**Organization** | Pointer to **string** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**AsnName** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Domain** | Pointer to **string** |  | [optional] 
**DateAllocated** | Pointer to **string** |  | [optional] 
**AllocationStatus** | Pointer to **string** |  | [optional] 
**NumOfIpv4Routes** | Pointer to **int32** |  | [optional] 
**NumOfIpv6Routes** | Pointer to **int32** |  | [optional] 
**Rir** | Pointer to **string** |  | [optional] 
**Routes** | Pointer to **[]string** | It will only be included in the response, if you set include&#x3D;routes in the request | [optional] 
**Upstreams** | Pointer to [**[]ASNConnection**](ASNConnection.md) |  | [optional] 
**Downstreams** | Pointer to [**[]ASNConnection**](ASNConnection.md) |  | [optional] 
**Peers** | Pointer to [**[]ASNConnection**](ASNConnection.md) |  | [optional] 
**WhoisResponse** | Pointer to **string** |  | [optional] 

## Methods

### NewASNResponseXMLAsn

`func NewASNResponseXMLAsn() *ASNResponseXMLAsn`

NewASNResponseXMLAsn instantiates a new ASNResponseXMLAsn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewASNResponseXMLAsnWithDefaults

`func NewASNResponseXMLAsnWithDefaults() *ASNResponseXMLAsn`

NewASNResponseXMLAsnWithDefaults instantiates a new ASNResponseXMLAsn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsNumber

`func (o *ASNResponseXMLAsn) GetAsNumber() string`

GetAsNumber returns the AsNumber field if non-nil, zero value otherwise.

### GetAsNumberOk

`func (o *ASNResponseXMLAsn) GetAsNumberOk() (*string, bool)`

GetAsNumberOk returns a tuple with the AsNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsNumber

`func (o *ASNResponseXMLAsn) SetAsNumber(v string)`

SetAsNumber sets AsNumber field to given value.

### HasAsNumber

`func (o *ASNResponseXMLAsn) HasAsNumber() bool`

HasAsNumber returns a boolean if a field has been set.

### GetOrganization

`func (o *ASNResponseXMLAsn) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *ASNResponseXMLAsn) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *ASNResponseXMLAsn) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *ASNResponseXMLAsn) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetCountry

`func (o *ASNResponseXMLAsn) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *ASNResponseXMLAsn) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *ASNResponseXMLAsn) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *ASNResponseXMLAsn) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetAsnName

`func (o *ASNResponseXMLAsn) GetAsnName() string`

GetAsnName returns the AsnName field if non-nil, zero value otherwise.

### GetAsnNameOk

`func (o *ASNResponseXMLAsn) GetAsnNameOk() (*string, bool)`

GetAsnNameOk returns a tuple with the AsnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsnName

`func (o *ASNResponseXMLAsn) SetAsnName(v string)`

SetAsnName sets AsnName field to given value.

### HasAsnName

`func (o *ASNResponseXMLAsn) HasAsnName() bool`

HasAsnName returns a boolean if a field has been set.

### GetType

`func (o *ASNResponseXMLAsn) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ASNResponseXMLAsn) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ASNResponseXMLAsn) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ASNResponseXMLAsn) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDomain

`func (o *ASNResponseXMLAsn) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *ASNResponseXMLAsn) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *ASNResponseXMLAsn) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *ASNResponseXMLAsn) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetDateAllocated

`func (o *ASNResponseXMLAsn) GetDateAllocated() string`

GetDateAllocated returns the DateAllocated field if non-nil, zero value otherwise.

### GetDateAllocatedOk

`func (o *ASNResponseXMLAsn) GetDateAllocatedOk() (*string, bool)`

GetDateAllocatedOk returns a tuple with the DateAllocated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAllocated

`func (o *ASNResponseXMLAsn) SetDateAllocated(v string)`

SetDateAllocated sets DateAllocated field to given value.

### HasDateAllocated

`func (o *ASNResponseXMLAsn) HasDateAllocated() bool`

HasDateAllocated returns a boolean if a field has been set.

### GetAllocationStatus

`func (o *ASNResponseXMLAsn) GetAllocationStatus() string`

GetAllocationStatus returns the AllocationStatus field if non-nil, zero value otherwise.

### GetAllocationStatusOk

`func (o *ASNResponseXMLAsn) GetAllocationStatusOk() (*string, bool)`

GetAllocationStatusOk returns a tuple with the AllocationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationStatus

`func (o *ASNResponseXMLAsn) SetAllocationStatus(v string)`

SetAllocationStatus sets AllocationStatus field to given value.

### HasAllocationStatus

`func (o *ASNResponseXMLAsn) HasAllocationStatus() bool`

HasAllocationStatus returns a boolean if a field has been set.

### GetNumOfIpv4Routes

`func (o *ASNResponseXMLAsn) GetNumOfIpv4Routes() int32`

GetNumOfIpv4Routes returns the NumOfIpv4Routes field if non-nil, zero value otherwise.

### GetNumOfIpv4RoutesOk

`func (o *ASNResponseXMLAsn) GetNumOfIpv4RoutesOk() (*int32, bool)`

GetNumOfIpv4RoutesOk returns a tuple with the NumOfIpv4Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumOfIpv4Routes

`func (o *ASNResponseXMLAsn) SetNumOfIpv4Routes(v int32)`

SetNumOfIpv4Routes sets NumOfIpv4Routes field to given value.

### HasNumOfIpv4Routes

`func (o *ASNResponseXMLAsn) HasNumOfIpv4Routes() bool`

HasNumOfIpv4Routes returns a boolean if a field has been set.

### GetNumOfIpv6Routes

`func (o *ASNResponseXMLAsn) GetNumOfIpv6Routes() int32`

GetNumOfIpv6Routes returns the NumOfIpv6Routes field if non-nil, zero value otherwise.

### GetNumOfIpv6RoutesOk

`func (o *ASNResponseXMLAsn) GetNumOfIpv6RoutesOk() (*int32, bool)`

GetNumOfIpv6RoutesOk returns a tuple with the NumOfIpv6Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumOfIpv6Routes

`func (o *ASNResponseXMLAsn) SetNumOfIpv6Routes(v int32)`

SetNumOfIpv6Routes sets NumOfIpv6Routes field to given value.

### HasNumOfIpv6Routes

`func (o *ASNResponseXMLAsn) HasNumOfIpv6Routes() bool`

HasNumOfIpv6Routes returns a boolean if a field has been set.

### GetRir

`func (o *ASNResponseXMLAsn) GetRir() string`

GetRir returns the Rir field if non-nil, zero value otherwise.

### GetRirOk

`func (o *ASNResponseXMLAsn) GetRirOk() (*string, bool)`

GetRirOk returns a tuple with the Rir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRir

`func (o *ASNResponseXMLAsn) SetRir(v string)`

SetRir sets Rir field to given value.

### HasRir

`func (o *ASNResponseXMLAsn) HasRir() bool`

HasRir returns a boolean if a field has been set.

### GetRoutes

`func (o *ASNResponseXMLAsn) GetRoutes() []string`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *ASNResponseXMLAsn) GetRoutesOk() (*[]string, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *ASNResponseXMLAsn) SetRoutes(v []string)`

SetRoutes sets Routes field to given value.

### HasRoutes

`func (o *ASNResponseXMLAsn) HasRoutes() bool`

HasRoutes returns a boolean if a field has been set.

### GetUpstreams

`func (o *ASNResponseXMLAsn) GetUpstreams() []ASNConnection`

GetUpstreams returns the Upstreams field if non-nil, zero value otherwise.

### GetUpstreamsOk

`func (o *ASNResponseXMLAsn) GetUpstreamsOk() (*[]ASNConnection, bool)`

GetUpstreamsOk returns a tuple with the Upstreams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreams

`func (o *ASNResponseXMLAsn) SetUpstreams(v []ASNConnection)`

SetUpstreams sets Upstreams field to given value.

### HasUpstreams

`func (o *ASNResponseXMLAsn) HasUpstreams() bool`

HasUpstreams returns a boolean if a field has been set.

### GetDownstreams

`func (o *ASNResponseXMLAsn) GetDownstreams() []ASNConnection`

GetDownstreams returns the Downstreams field if non-nil, zero value otherwise.

### GetDownstreamsOk

`func (o *ASNResponseXMLAsn) GetDownstreamsOk() (*[]ASNConnection, bool)`

GetDownstreamsOk returns a tuple with the Downstreams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownstreams

`func (o *ASNResponseXMLAsn) SetDownstreams(v []ASNConnection)`

SetDownstreams sets Downstreams field to given value.

### HasDownstreams

`func (o *ASNResponseXMLAsn) HasDownstreams() bool`

HasDownstreams returns a boolean if a field has been set.

### GetPeers

`func (o *ASNResponseXMLAsn) GetPeers() []ASNConnection`

GetPeers returns the Peers field if non-nil, zero value otherwise.

### GetPeersOk

`func (o *ASNResponseXMLAsn) GetPeersOk() (*[]ASNConnection, bool)`

GetPeersOk returns a tuple with the Peers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeers

`func (o *ASNResponseXMLAsn) SetPeers(v []ASNConnection)`

SetPeers sets Peers field to given value.

### HasPeers

`func (o *ASNResponseXMLAsn) HasPeers() bool`

HasPeers returns a boolean if a field has been set.

### GetWhoisResponse

`func (o *ASNResponseXMLAsn) GetWhoisResponse() string`

GetWhoisResponse returns the WhoisResponse field if non-nil, zero value otherwise.

### GetWhoisResponseOk

`func (o *ASNResponseXMLAsn) GetWhoisResponseOk() (*string, bool)`

GetWhoisResponseOk returns a tuple with the WhoisResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhoisResponse

`func (o *ASNResponseXMLAsn) SetWhoisResponse(v string)`

SetWhoisResponse sets WhoisResponse field to given value.

### HasWhoisResponse

`func (o *ASNResponseXMLAsn) HasWhoisResponse() bool`

HasWhoisResponse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


