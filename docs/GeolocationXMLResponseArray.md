# GeolocationXMLResponseArray

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

### NewGeolocationXMLResponseArray

`func NewGeolocationXMLResponseArray() *GeolocationXMLResponseArray`

NewGeolocationXMLResponseArray instantiates a new GeolocationXMLResponseArray object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeolocationXMLResponseArrayWithDefaults

`func NewGeolocationXMLResponseArrayWithDefaults() *GeolocationXMLResponseArray`

NewGeolocationXMLResponseArrayWithDefaults instantiates a new GeolocationXMLResponseArray object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *GeolocationXMLResponseArray) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *GeolocationXMLResponseArray) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *GeolocationXMLResponseArray) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *GeolocationXMLResponseArray) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetHostname

`func (o *GeolocationXMLResponseArray) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *GeolocationXMLResponseArray) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *GeolocationXMLResponseArray) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *GeolocationXMLResponseArray) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetDomain

`func (o *GeolocationXMLResponseArray) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *GeolocationXMLResponseArray) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *GeolocationXMLResponseArray) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *GeolocationXMLResponseArray) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetLocation

`func (o *GeolocationXMLResponseArray) GetLocation() Location`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *GeolocationXMLResponseArray) GetLocationOk() (*Location, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *GeolocationXMLResponseArray) SetLocation(v Location)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *GeolocationXMLResponseArray) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetCountryMetadata

`func (o *GeolocationXMLResponseArray) GetCountryMetadata() CountryMetadata`

GetCountryMetadata returns the CountryMetadata field if non-nil, zero value otherwise.

### GetCountryMetadataOk

`func (o *GeolocationXMLResponseArray) GetCountryMetadataOk() (*CountryMetadata, bool)`

GetCountryMetadataOk returns a tuple with the CountryMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryMetadata

`func (o *GeolocationXMLResponseArray) SetCountryMetadata(v CountryMetadata)`

SetCountryMetadata sets CountryMetadata field to given value.

### HasCountryMetadata

`func (o *GeolocationXMLResponseArray) HasCountryMetadata() bool`

HasCountryMetadata returns a boolean if a field has been set.

### GetNetwork

`func (o *GeolocationXMLResponseArray) GetNetwork() Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *GeolocationXMLResponseArray) GetNetworkOk() (*Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *GeolocationXMLResponseArray) SetNetwork(v Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *GeolocationXMLResponseArray) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetCurrency

`func (o *GeolocationXMLResponseArray) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GeolocationXMLResponseArray) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GeolocationXMLResponseArray) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *GeolocationXMLResponseArray) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetSecurity

`func (o *GeolocationXMLResponseArray) GetSecurity() Security`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *GeolocationXMLResponseArray) GetSecurityOk() (*Security, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *GeolocationXMLResponseArray) SetSecurity(v Security)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *GeolocationXMLResponseArray) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetAbuse

`func (o *GeolocationXMLResponseArray) GetAbuse() Abuse`

GetAbuse returns the Abuse field if non-nil, zero value otherwise.

### GetAbuseOk

`func (o *GeolocationXMLResponseArray) GetAbuseOk() (*Abuse, bool)`

GetAbuseOk returns a tuple with the Abuse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbuse

`func (o *GeolocationXMLResponseArray) SetAbuse(v Abuse)`

SetAbuse sets Abuse field to given value.

### HasAbuse

`func (o *GeolocationXMLResponseArray) HasAbuse() bool`

HasAbuse returns a boolean if a field has been set.

### GetTimeZone

`func (o *GeolocationXMLResponseArray) GetTimeZone() TimeZone`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *GeolocationXMLResponseArray) GetTimeZoneOk() (*TimeZone, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *GeolocationXMLResponseArray) SetTimeZone(v TimeZone)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *GeolocationXMLResponseArray) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetUserAgent

`func (o *GeolocationXMLResponseArray) GetUserAgent() UserAgentData`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *GeolocationXMLResponseArray) GetUserAgentOk() (*UserAgentData, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *GeolocationXMLResponseArray) SetUserAgent(v UserAgentData)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *GeolocationXMLResponseArray) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


