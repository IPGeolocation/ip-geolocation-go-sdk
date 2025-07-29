# TimeZoneDetailedResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | Pointer to **string** |  | [optional] 
**AirportDetails** | Pointer to [**TimezoneAirport**](TimezoneAirport.md) |  | [optional] 
**LoCodeDetails** | Pointer to [**TimezoneLocode**](TimezoneLocode.md) |  | [optional] 
**Location** | Pointer to [**TimezoneLocation**](TimezoneLocation.md) |  | [optional] 
**TimeZone** | Pointer to [**TimezoneDetail**](TimezoneDetail.md) |  | [optional] 

## Methods

### NewTimeZoneDetailedResponse

`func NewTimeZoneDetailedResponse() *TimeZoneDetailedResponse`

NewTimeZoneDetailedResponse instantiates a new TimeZoneDetailedResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeZoneDetailedResponseWithDefaults

`func NewTimeZoneDetailedResponseWithDefaults() *TimeZoneDetailedResponse`

NewTimeZoneDetailedResponseWithDefaults instantiates a new TimeZoneDetailedResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *TimeZoneDetailedResponse) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *TimeZoneDetailedResponse) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *TimeZoneDetailedResponse) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *TimeZoneDetailedResponse) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetAirportDetails

`func (o *TimeZoneDetailedResponse) GetAirportDetails() TimezoneAirport`

GetAirportDetails returns the AirportDetails field if non-nil, zero value otherwise.

### GetAirportDetailsOk

`func (o *TimeZoneDetailedResponse) GetAirportDetailsOk() (*TimezoneAirport, bool)`

GetAirportDetailsOk returns a tuple with the AirportDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirportDetails

`func (o *TimeZoneDetailedResponse) SetAirportDetails(v TimezoneAirport)`

SetAirportDetails sets AirportDetails field to given value.

### HasAirportDetails

`func (o *TimeZoneDetailedResponse) HasAirportDetails() bool`

HasAirportDetails returns a boolean if a field has been set.

### GetLoCodeDetails

`func (o *TimeZoneDetailedResponse) GetLoCodeDetails() TimezoneLocode`

GetLoCodeDetails returns the LoCodeDetails field if non-nil, zero value otherwise.

### GetLoCodeDetailsOk

`func (o *TimeZoneDetailedResponse) GetLoCodeDetailsOk() (*TimezoneLocode, bool)`

GetLoCodeDetailsOk returns a tuple with the LoCodeDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoCodeDetails

`func (o *TimeZoneDetailedResponse) SetLoCodeDetails(v TimezoneLocode)`

SetLoCodeDetails sets LoCodeDetails field to given value.

### HasLoCodeDetails

`func (o *TimeZoneDetailedResponse) HasLoCodeDetails() bool`

HasLoCodeDetails returns a boolean if a field has been set.

### GetLocation

`func (o *TimeZoneDetailedResponse) GetLocation() TimezoneLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *TimeZoneDetailedResponse) GetLocationOk() (*TimezoneLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *TimeZoneDetailedResponse) SetLocation(v TimezoneLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *TimeZoneDetailedResponse) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetTimeZone

`func (o *TimeZoneDetailedResponse) GetTimeZone() TimezoneDetail`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *TimeZoneDetailedResponse) GetTimeZoneOk() (*TimezoneDetail, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *TimeZoneDetailedResponse) SetTimeZone(v TimezoneDetail)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *TimeZoneDetailedResponse) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


