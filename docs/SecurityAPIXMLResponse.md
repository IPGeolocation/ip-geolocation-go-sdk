# SecurityAPIXMLResponse

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

## Methods

### NewSecurityAPIXMLResponse

`func NewSecurityAPIXMLResponse() *SecurityAPIXMLResponse`

NewSecurityAPIXMLResponse instantiates a new SecurityAPIXMLResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityAPIXMLResponseWithDefaults

`func NewSecurityAPIXMLResponseWithDefaults() *SecurityAPIXMLResponse`

NewSecurityAPIXMLResponseWithDefaults instantiates a new SecurityAPIXMLResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *SecurityAPIXMLResponse) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *SecurityAPIXMLResponse) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *SecurityAPIXMLResponse) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *SecurityAPIXMLResponse) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetHostname

`func (o *SecurityAPIXMLResponse) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *SecurityAPIXMLResponse) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *SecurityAPIXMLResponse) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *SecurityAPIXMLResponse) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetSecurity

`func (o *SecurityAPIXMLResponse) GetSecurity() Security`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *SecurityAPIXMLResponse) GetSecurityOk() (*Security, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *SecurityAPIXMLResponse) SetSecurity(v Security)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *SecurityAPIXMLResponse) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetLocation

`func (o *SecurityAPIXMLResponse) GetLocation() LocationMinimal`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *SecurityAPIXMLResponse) GetLocationOk() (*LocationMinimal, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *SecurityAPIXMLResponse) SetLocation(v LocationMinimal)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *SecurityAPIXMLResponse) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetNetwork

`func (o *SecurityAPIXMLResponse) GetNetwork() NetworkMinimal`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *SecurityAPIXMLResponse) GetNetworkOk() (*NetworkMinimal, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *SecurityAPIXMLResponse) SetNetwork(v NetworkMinimal)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *SecurityAPIXMLResponse) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetTimeZone

`func (o *SecurityAPIXMLResponse) GetTimeZone() TimeZone`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *SecurityAPIXMLResponse) GetTimeZoneOk() (*TimeZone, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *SecurityAPIXMLResponse) SetTimeZone(v TimeZone)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *SecurityAPIXMLResponse) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetUserAgent

`func (o *SecurityAPIXMLResponse) GetUserAgent() UserAgentData`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *SecurityAPIXMLResponse) GetUserAgentOk() (*UserAgentData, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *SecurityAPIXMLResponse) SetUserAgent(v UserAgentData)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *SecurityAPIXMLResponse) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.

### GetCountryMetadata

`func (o *SecurityAPIXMLResponse) GetCountryMetadata() CountryMetadata`

GetCountryMetadata returns the CountryMetadata field if non-nil, zero value otherwise.

### GetCountryMetadataOk

`func (o *SecurityAPIXMLResponse) GetCountryMetadataOk() (*CountryMetadata, bool)`

GetCountryMetadataOk returns a tuple with the CountryMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryMetadata

`func (o *SecurityAPIXMLResponse) SetCountryMetadata(v CountryMetadata)`

SetCountryMetadata sets CountryMetadata field to given value.

### HasCountryMetadata

`func (o *SecurityAPIXMLResponse) HasCountryMetadata() bool`

HasCountryMetadata returns a boolean if a field has been set.

### GetCurrency

`func (o *SecurityAPIXMLResponse) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *SecurityAPIXMLResponse) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *SecurityAPIXMLResponse) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *SecurityAPIXMLResponse) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


