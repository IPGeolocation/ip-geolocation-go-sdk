# GeolocationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | Pointer to **string** |  | [optional] 
**Hostname** | Pointer to **string** |  | [optional] 
**Domain** | Pointer to **string** |  | [optional] 
**Location** | Pointer to [**Location**](Location.md) |  | [optional] 
**CountryMetadata** | Pointer to [**CountryMetadata**](CountryMetadata.md) |  | [optional] 
**Network** | Pointer to [**Network**](Network.md) |  | [optional] 
**Currency** | Pointer to [**Currency**](Currency.md) |  | [optional] 
**Security** | Pointer to [**Security**](Security.md) |  | [optional] 
**Abuse** | Pointer to [**Abuse**](Abuse.md) |  | [optional] 
**TimeZone** | Pointer to [**TimeZone**](TimeZone.md) |  | [optional] 
**UserAgent** | Pointer to [**UserAgentData**](UserAgentData.md) |  | [optional] 

## Methods

### NewGeolocationResponse

`func NewGeolocationResponse() *GeolocationResponse`

NewGeolocationResponse instantiates a new GeolocationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeolocationResponseWithDefaults

`func NewGeolocationResponseWithDefaults() *GeolocationResponse`

NewGeolocationResponseWithDefaults instantiates a new GeolocationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *GeolocationResponse) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *GeolocationResponse) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *GeolocationResponse) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *GeolocationResponse) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetHostname

`func (o *GeolocationResponse) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *GeolocationResponse) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *GeolocationResponse) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *GeolocationResponse) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetDomain

`func (o *GeolocationResponse) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *GeolocationResponse) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *GeolocationResponse) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *GeolocationResponse) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetLocation

`func (o *GeolocationResponse) GetLocation() Location`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *GeolocationResponse) GetLocationOk() (*Location, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *GeolocationResponse) SetLocation(v Location)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *GeolocationResponse) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetCountryMetadata

`func (o *GeolocationResponse) GetCountryMetadata() CountryMetadata`

GetCountryMetadata returns the CountryMetadata field if non-nil, zero value otherwise.

### GetCountryMetadataOk

`func (o *GeolocationResponse) GetCountryMetadataOk() (*CountryMetadata, bool)`

GetCountryMetadataOk returns a tuple with the CountryMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryMetadata

`func (o *GeolocationResponse) SetCountryMetadata(v CountryMetadata)`

SetCountryMetadata sets CountryMetadata field to given value.

### HasCountryMetadata

`func (o *GeolocationResponse) HasCountryMetadata() bool`

HasCountryMetadata returns a boolean if a field has been set.

### GetNetwork

`func (o *GeolocationResponse) GetNetwork() Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *GeolocationResponse) GetNetworkOk() (*Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *GeolocationResponse) SetNetwork(v Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *GeolocationResponse) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetCurrency

`func (o *GeolocationResponse) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GeolocationResponse) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GeolocationResponse) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *GeolocationResponse) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetSecurity

`func (o *GeolocationResponse) GetSecurity() Security`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *GeolocationResponse) GetSecurityOk() (*Security, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *GeolocationResponse) SetSecurity(v Security)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *GeolocationResponse) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetAbuse

`func (o *GeolocationResponse) GetAbuse() Abuse`

GetAbuse returns the Abuse field if non-nil, zero value otherwise.

### GetAbuseOk

`func (o *GeolocationResponse) GetAbuseOk() (*Abuse, bool)`

GetAbuseOk returns a tuple with the Abuse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbuse

`func (o *GeolocationResponse) SetAbuse(v Abuse)`

SetAbuse sets Abuse field to given value.

### HasAbuse

`func (o *GeolocationResponse) HasAbuse() bool`

HasAbuse returns a boolean if a field has been set.

### GetTimeZone

`func (o *GeolocationResponse) GetTimeZone() TimeZone`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *GeolocationResponse) GetTimeZoneOk() (*TimeZone, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *GeolocationResponse) SetTimeZone(v TimeZone)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *GeolocationResponse) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetUserAgent

`func (o *GeolocationResponse) GetUserAgent() UserAgentData`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *GeolocationResponse) GetUserAgentOk() (*UserAgentData, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *GeolocationResponse) SetUserAgent(v UserAgentData)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *GeolocationResponse) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


