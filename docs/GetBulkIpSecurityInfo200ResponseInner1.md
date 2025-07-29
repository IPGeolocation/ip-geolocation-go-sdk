# GetBulkIpSecurityInfo200ResponseInner1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | Pointer to **string** |  | [optional] 
**Hostname** | Pointer to **string** |  | [optional] 
**Security** | Pointer to [**Security**](Security.md) |  | [optional] 
**Location** | Pointer to [**LocationMinimal**](LocationMinimal.md) |  | [optional] 
**Network** | Pointer to [**NetworkMinimal**](NetworkMinimal.md) |  | [optional] 
**TimeZone** | Pointer to [**TimeZone**](TimeZone.md) |  | [optional] 
**UserAgent** | Pointer to [**UserAgentData**](UserAgentData.md) |  | [optional] 
**CountryMetadata** | Pointer to [**CountryMetadata**](CountryMetadata.md) |  | [optional] 
**Currency** | Pointer to [**Currency**](Currency.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewGetBulkIpSecurityInfo200ResponseInner1

`func NewGetBulkIpSecurityInfo200ResponseInner1() *GetBulkIpSecurityInfo200ResponseInner1`

NewGetBulkIpSecurityInfo200ResponseInner1 instantiates a new GetBulkIpSecurityInfo200ResponseInner1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBulkIpSecurityInfo200ResponseInner1WithDefaults

`func NewGetBulkIpSecurityInfo200ResponseInner1WithDefaults() *GetBulkIpSecurityInfo200ResponseInner1`

NewGetBulkIpSecurityInfo200ResponseInner1WithDefaults instantiates a new GetBulkIpSecurityInfo200ResponseInner1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *GetBulkIpSecurityInfo200ResponseInner1) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *GetBulkIpSecurityInfo200ResponseInner1) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *GetBulkIpSecurityInfo200ResponseInner1) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *GetBulkIpSecurityInfo200ResponseInner1) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetHostname

`func (o *GetBulkIpSecurityInfo200ResponseInner1) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *GetBulkIpSecurityInfo200ResponseInner1) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *GetBulkIpSecurityInfo200ResponseInner1) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *GetBulkIpSecurityInfo200ResponseInner1) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetSecurity

`func (o *GetBulkIpSecurityInfo200ResponseInner1) GetSecurity() Security`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *GetBulkIpSecurityInfo200ResponseInner1) GetSecurityOk() (*Security, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *GetBulkIpSecurityInfo200ResponseInner1) SetSecurity(v Security)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *GetBulkIpSecurityInfo200ResponseInner1) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetLocation

`func (o *GetBulkIpSecurityInfo200ResponseInner1) GetLocation() LocationMinimal`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *GetBulkIpSecurityInfo200ResponseInner1) GetLocationOk() (*LocationMinimal, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *GetBulkIpSecurityInfo200ResponseInner1) SetLocation(v LocationMinimal)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *GetBulkIpSecurityInfo200ResponseInner1) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetNetwork

`func (o *GetBulkIpSecurityInfo200ResponseInner1) GetNetwork() NetworkMinimal`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *GetBulkIpSecurityInfo200ResponseInner1) GetNetworkOk() (*NetworkMinimal, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *GetBulkIpSecurityInfo200ResponseInner1) SetNetwork(v NetworkMinimal)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *GetBulkIpSecurityInfo200ResponseInner1) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetTimeZone

`func (o *GetBulkIpSecurityInfo200ResponseInner1) GetTimeZone() TimeZone`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *GetBulkIpSecurityInfo200ResponseInner1) GetTimeZoneOk() (*TimeZone, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *GetBulkIpSecurityInfo200ResponseInner1) SetTimeZone(v TimeZone)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *GetBulkIpSecurityInfo200ResponseInner1) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetUserAgent

`func (o *GetBulkIpSecurityInfo200ResponseInner1) GetUserAgent() UserAgentData`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *GetBulkIpSecurityInfo200ResponseInner1) GetUserAgentOk() (*UserAgentData, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *GetBulkIpSecurityInfo200ResponseInner1) SetUserAgent(v UserAgentData)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *GetBulkIpSecurityInfo200ResponseInner1) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.

### GetCountryMetadata

`func (o *GetBulkIpSecurityInfo200ResponseInner1) GetCountryMetadata() CountryMetadata`

GetCountryMetadata returns the CountryMetadata field if non-nil, zero value otherwise.

### GetCountryMetadataOk

`func (o *GetBulkIpSecurityInfo200ResponseInner1) GetCountryMetadataOk() (*CountryMetadata, bool)`

GetCountryMetadataOk returns a tuple with the CountryMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryMetadata

`func (o *GetBulkIpSecurityInfo200ResponseInner1) SetCountryMetadata(v CountryMetadata)`

SetCountryMetadata sets CountryMetadata field to given value.

### HasCountryMetadata

`func (o *GetBulkIpSecurityInfo200ResponseInner1) HasCountryMetadata() bool`

HasCountryMetadata returns a boolean if a field has been set.

### GetCurrency

`func (o *GetBulkIpSecurityInfo200ResponseInner1) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GetBulkIpSecurityInfo200ResponseInner1) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GetBulkIpSecurityInfo200ResponseInner1) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *GetBulkIpSecurityInfo200ResponseInner1) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetMessage

`func (o *GetBulkIpSecurityInfo200ResponseInner1) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetBulkIpSecurityInfo200ResponseInner1) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetBulkIpSecurityInfo200ResponseInner1) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetBulkIpSecurityInfo200ResponseInner1) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


