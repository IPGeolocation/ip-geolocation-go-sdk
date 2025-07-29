# ASNDetails

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
**NumOfIpv4Routes** | Pointer to **string** |  | [optional] 
**NumOfIpv6Routes** | Pointer to **string** |  | [optional] 
**Rir** | Pointer to **string** |  | [optional] 
**Routes** | Pointer to **[]string** | It will only be included in the response, if you set include&#x3D;routes in the request | [optional] 
**Upstreams** | Pointer to [**[]ASNConnection**](ASNConnection.md) | It will only be included in the response, if you set include&#x3D;upstreams in the request | [optional] 
**Downstreams** | Pointer to [**[]ASNConnection**](ASNConnection.md) | It will only be included in the response, if you set include&#x3D;downstreams in the request | [optional] 
**Peers** | Pointer to [**[]ASNConnection**](ASNConnection.md) | It will only be included in the response, if you set include&#x3D;peers in the request | [optional] 
**WhoisResponse** | Pointer to **string** | It will only be included in the response, if you set include&#x3D;whois_response in the request | [optional] 

## Methods

### NewASNDetails

`func NewASNDetails() *ASNDetails`

NewASNDetails instantiates a new ASNDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewASNDetailsWithDefaults

`func NewASNDetailsWithDefaults() *ASNDetails`

NewASNDetailsWithDefaults instantiates a new ASNDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsNumber

`func (o *ASNDetails) GetAsNumber() string`

GetAsNumber returns the AsNumber field if non-nil, zero value otherwise.

### GetAsNumberOk

`func (o *ASNDetails) GetAsNumberOk() (*string, bool)`

GetAsNumberOk returns a tuple with the AsNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsNumber

`func (o *ASNDetails) SetAsNumber(v string)`

SetAsNumber sets AsNumber field to given value.

### HasAsNumber

`func (o *ASNDetails) HasAsNumber() bool`

HasAsNumber returns a boolean if a field has been set.

### GetOrganization

`func (o *ASNDetails) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *ASNDetails) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *ASNDetails) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *ASNDetails) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetCountry

`func (o *ASNDetails) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *ASNDetails) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *ASNDetails) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *ASNDetails) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetAsnName

`func (o *ASNDetails) GetAsnName() string`

GetAsnName returns the AsnName field if non-nil, zero value otherwise.

### GetAsnNameOk

`func (o *ASNDetails) GetAsnNameOk() (*string, bool)`

GetAsnNameOk returns a tuple with the AsnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsnName

`func (o *ASNDetails) SetAsnName(v string)`

SetAsnName sets AsnName field to given value.

### HasAsnName

`func (o *ASNDetails) HasAsnName() bool`

HasAsnName returns a boolean if a field has been set.

### GetType

`func (o *ASNDetails) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ASNDetails) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ASNDetails) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ASNDetails) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDomain

`func (o *ASNDetails) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *ASNDetails) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *ASNDetails) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *ASNDetails) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetDateAllocated

`func (o *ASNDetails) GetDateAllocated() string`

GetDateAllocated returns the DateAllocated field if non-nil, zero value otherwise.

### GetDateAllocatedOk

`func (o *ASNDetails) GetDateAllocatedOk() (*string, bool)`

GetDateAllocatedOk returns a tuple with the DateAllocated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAllocated

`func (o *ASNDetails) SetDateAllocated(v string)`

SetDateAllocated sets DateAllocated field to given value.

### HasDateAllocated

`func (o *ASNDetails) HasDateAllocated() bool`

HasDateAllocated returns a boolean if a field has been set.

### GetAllocationStatus

`func (o *ASNDetails) GetAllocationStatus() string`

GetAllocationStatus returns the AllocationStatus field if non-nil, zero value otherwise.

### GetAllocationStatusOk

`func (o *ASNDetails) GetAllocationStatusOk() (*string, bool)`

GetAllocationStatusOk returns a tuple with the AllocationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationStatus

`func (o *ASNDetails) SetAllocationStatus(v string)`

SetAllocationStatus sets AllocationStatus field to given value.

### HasAllocationStatus

`func (o *ASNDetails) HasAllocationStatus() bool`

HasAllocationStatus returns a boolean if a field has been set.

### GetNumOfIpv4Routes

`func (o *ASNDetails) GetNumOfIpv4Routes() string`

GetNumOfIpv4Routes returns the NumOfIpv4Routes field if non-nil, zero value otherwise.

### GetNumOfIpv4RoutesOk

`func (o *ASNDetails) GetNumOfIpv4RoutesOk() (*string, bool)`

GetNumOfIpv4RoutesOk returns a tuple with the NumOfIpv4Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumOfIpv4Routes

`func (o *ASNDetails) SetNumOfIpv4Routes(v string)`

SetNumOfIpv4Routes sets NumOfIpv4Routes field to given value.

### HasNumOfIpv4Routes

`func (o *ASNDetails) HasNumOfIpv4Routes() bool`

HasNumOfIpv4Routes returns a boolean if a field has been set.

### GetNumOfIpv6Routes

`func (o *ASNDetails) GetNumOfIpv6Routes() string`

GetNumOfIpv6Routes returns the NumOfIpv6Routes field if non-nil, zero value otherwise.

### GetNumOfIpv6RoutesOk

`func (o *ASNDetails) GetNumOfIpv6RoutesOk() (*string, bool)`

GetNumOfIpv6RoutesOk returns a tuple with the NumOfIpv6Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumOfIpv6Routes

`func (o *ASNDetails) SetNumOfIpv6Routes(v string)`

SetNumOfIpv6Routes sets NumOfIpv6Routes field to given value.

### HasNumOfIpv6Routes

`func (o *ASNDetails) HasNumOfIpv6Routes() bool`

HasNumOfIpv6Routes returns a boolean if a field has been set.

### GetRir

`func (o *ASNDetails) GetRir() string`

GetRir returns the Rir field if non-nil, zero value otherwise.

### GetRirOk

`func (o *ASNDetails) GetRirOk() (*string, bool)`

GetRirOk returns a tuple with the Rir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRir

`func (o *ASNDetails) SetRir(v string)`

SetRir sets Rir field to given value.

### HasRir

`func (o *ASNDetails) HasRir() bool`

HasRir returns a boolean if a field has been set.

### GetRoutes

`func (o *ASNDetails) GetRoutes() []string`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *ASNDetails) GetRoutesOk() (*[]string, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *ASNDetails) SetRoutes(v []string)`

SetRoutes sets Routes field to given value.

### HasRoutes

`func (o *ASNDetails) HasRoutes() bool`

HasRoutes returns a boolean if a field has been set.

### GetUpstreams

`func (o *ASNDetails) GetUpstreams() []ASNConnection`

GetUpstreams returns the Upstreams field if non-nil, zero value otherwise.

### GetUpstreamsOk

`func (o *ASNDetails) GetUpstreamsOk() (*[]ASNConnection, bool)`

GetUpstreamsOk returns a tuple with the Upstreams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreams

`func (o *ASNDetails) SetUpstreams(v []ASNConnection)`

SetUpstreams sets Upstreams field to given value.

### HasUpstreams

`func (o *ASNDetails) HasUpstreams() bool`

HasUpstreams returns a boolean if a field has been set.

### GetDownstreams

`func (o *ASNDetails) GetDownstreams() []ASNConnection`

GetDownstreams returns the Downstreams field if non-nil, zero value otherwise.

### GetDownstreamsOk

`func (o *ASNDetails) GetDownstreamsOk() (*[]ASNConnection, bool)`

GetDownstreamsOk returns a tuple with the Downstreams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownstreams

`func (o *ASNDetails) SetDownstreams(v []ASNConnection)`

SetDownstreams sets Downstreams field to given value.

### HasDownstreams

`func (o *ASNDetails) HasDownstreams() bool`

HasDownstreams returns a boolean if a field has been set.

### GetPeers

`func (o *ASNDetails) GetPeers() []ASNConnection`

GetPeers returns the Peers field if non-nil, zero value otherwise.

### GetPeersOk

`func (o *ASNDetails) GetPeersOk() (*[]ASNConnection, bool)`

GetPeersOk returns a tuple with the Peers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeers

`func (o *ASNDetails) SetPeers(v []ASNConnection)`

SetPeers sets Peers field to given value.

### HasPeers

`func (o *ASNDetails) HasPeers() bool`

HasPeers returns a boolean if a field has been set.

### GetWhoisResponse

`func (o *ASNDetails) GetWhoisResponse() string`

GetWhoisResponse returns the WhoisResponse field if non-nil, zero value otherwise.

### GetWhoisResponseOk

`func (o *ASNDetails) GetWhoisResponseOk() (*string, bool)`

GetWhoisResponseOk returns a tuple with the WhoisResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhoisResponse

`func (o *ASNDetails) SetWhoisResponse(v string)`

SetWhoisResponse sets WhoisResponse field to given value.

### HasWhoisResponse

`func (o *ASNDetails) HasWhoisResponse() bool`

HasWhoisResponse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


