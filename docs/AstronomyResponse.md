# AstronomyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | Pointer to **string** |  | [optional] 
**Location** | Pointer to [**AstronomyLocation**](AstronomyLocation.md) |  | [optional] 
**Astronomy** | Pointer to [**Astronomy**](Astronomy.md) |  | [optional] 

## Methods

### NewAstronomyResponse

`func NewAstronomyResponse() *AstronomyResponse`

NewAstronomyResponse instantiates a new AstronomyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAstronomyResponseWithDefaults

`func NewAstronomyResponseWithDefaults() *AstronomyResponse`

NewAstronomyResponseWithDefaults instantiates a new AstronomyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *AstronomyResponse) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *AstronomyResponse) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *AstronomyResponse) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *AstronomyResponse) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetLocation

`func (o *AstronomyResponse) GetLocation() AstronomyLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *AstronomyResponse) GetLocationOk() (*AstronomyLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *AstronomyResponse) SetLocation(v AstronomyLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *AstronomyResponse) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetAstronomy

`func (o *AstronomyResponse) GetAstronomy() Astronomy`

GetAstronomy returns the Astronomy field if non-nil, zero value otherwise.

### GetAstronomyOk

`func (o *AstronomyResponse) GetAstronomyOk() (*Astronomy, bool)`

GetAstronomyOk returns a tuple with the Astronomy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAstronomy

`func (o *AstronomyResponse) SetAstronomy(v Astronomy)`

SetAstronomy sets Astronomy field to given value.

### HasAstronomy

`func (o *AstronomyResponse) HasAstronomy() bool`

HasAstronomy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


