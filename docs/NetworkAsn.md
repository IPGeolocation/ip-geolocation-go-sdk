# NetworkAsn

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

## Methods

### NewNetworkAsn

`func NewNetworkAsn() *NetworkAsn`

NewNetworkAsn instantiates a new NetworkAsn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkAsnWithDefaults

`func NewNetworkAsnWithDefaults() *NetworkAsn`

NewNetworkAsnWithDefaults instantiates a new NetworkAsn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsNumber

`func (o *NetworkAsn) GetAsNumber() string`

GetAsNumber returns the AsNumber field if non-nil, zero value otherwise.

### GetAsNumberOk

`func (o *NetworkAsn) GetAsNumberOk() (*string, bool)`

GetAsNumberOk returns a tuple with the AsNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsNumber

`func (o *NetworkAsn) SetAsNumber(v string)`

SetAsNumber sets AsNumber field to given value.

### HasAsNumber

`func (o *NetworkAsn) HasAsNumber() bool`

HasAsNumber returns a boolean if a field has been set.

### GetOrganization

`func (o *NetworkAsn) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *NetworkAsn) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *NetworkAsn) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *NetworkAsn) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetCountry

`func (o *NetworkAsn) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *NetworkAsn) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *NetworkAsn) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *NetworkAsn) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetAsnName

`func (o *NetworkAsn) GetAsnName() string`

GetAsnName returns the AsnName field if non-nil, zero value otherwise.

### GetAsnNameOk

`func (o *NetworkAsn) GetAsnNameOk() (*string, bool)`

GetAsnNameOk returns a tuple with the AsnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsnName

`func (o *NetworkAsn) SetAsnName(v string)`

SetAsnName sets AsnName field to given value.

### HasAsnName

`func (o *NetworkAsn) HasAsnName() bool`

HasAsnName returns a boolean if a field has been set.

### GetType

`func (o *NetworkAsn) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NetworkAsn) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NetworkAsn) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NetworkAsn) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDomain

`func (o *NetworkAsn) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *NetworkAsn) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *NetworkAsn) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *NetworkAsn) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetDateAllocated

`func (o *NetworkAsn) GetDateAllocated() string`

GetDateAllocated returns the DateAllocated field if non-nil, zero value otherwise.

### GetDateAllocatedOk

`func (o *NetworkAsn) GetDateAllocatedOk() (*string, bool)`

GetDateAllocatedOk returns a tuple with the DateAllocated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAllocated

`func (o *NetworkAsn) SetDateAllocated(v string)`

SetDateAllocated sets DateAllocated field to given value.

### HasDateAllocated

`func (o *NetworkAsn) HasDateAllocated() bool`

HasDateAllocated returns a boolean if a field has been set.

### GetAllocationStatus

`func (o *NetworkAsn) GetAllocationStatus() string`

GetAllocationStatus returns the AllocationStatus field if non-nil, zero value otherwise.

### GetAllocationStatusOk

`func (o *NetworkAsn) GetAllocationStatusOk() (*string, bool)`

GetAllocationStatusOk returns a tuple with the AllocationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationStatus

`func (o *NetworkAsn) SetAllocationStatus(v string)`

SetAllocationStatus sets AllocationStatus field to given value.

### HasAllocationStatus

`func (o *NetworkAsn) HasAllocationStatus() bool`

HasAllocationStatus returns a boolean if a field has been set.

### GetNumOfIpv4Routes

`func (o *NetworkAsn) GetNumOfIpv4Routes() string`

GetNumOfIpv4Routes returns the NumOfIpv4Routes field if non-nil, zero value otherwise.

### GetNumOfIpv4RoutesOk

`func (o *NetworkAsn) GetNumOfIpv4RoutesOk() (*string, bool)`

GetNumOfIpv4RoutesOk returns a tuple with the NumOfIpv4Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumOfIpv4Routes

`func (o *NetworkAsn) SetNumOfIpv4Routes(v string)`

SetNumOfIpv4Routes sets NumOfIpv4Routes field to given value.

### HasNumOfIpv4Routes

`func (o *NetworkAsn) HasNumOfIpv4Routes() bool`

HasNumOfIpv4Routes returns a boolean if a field has been set.

### GetNumOfIpv6Routes

`func (o *NetworkAsn) GetNumOfIpv6Routes() string`

GetNumOfIpv6Routes returns the NumOfIpv6Routes field if non-nil, zero value otherwise.

### GetNumOfIpv6RoutesOk

`func (o *NetworkAsn) GetNumOfIpv6RoutesOk() (*string, bool)`

GetNumOfIpv6RoutesOk returns a tuple with the NumOfIpv6Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumOfIpv6Routes

`func (o *NetworkAsn) SetNumOfIpv6Routes(v string)`

SetNumOfIpv6Routes sets NumOfIpv6Routes field to given value.

### HasNumOfIpv6Routes

`func (o *NetworkAsn) HasNumOfIpv6Routes() bool`

HasNumOfIpv6Routes returns a boolean if a field has been set.

### GetRir

`func (o *NetworkAsn) GetRir() string`

GetRir returns the Rir field if non-nil, zero value otherwise.

### GetRirOk

`func (o *NetworkAsn) GetRirOk() (*string, bool)`

GetRirOk returns a tuple with the Rir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRir

`func (o *NetworkAsn) SetRir(v string)`

SetRir sets Rir field to given value.

### HasRir

`func (o *NetworkAsn) HasRir() bool`

HasRir returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


