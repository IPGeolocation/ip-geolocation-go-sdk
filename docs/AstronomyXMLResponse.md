# AstronomyXMLResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | Pointer to **string** |  | [optional] 
**Location** | Pointer to [**TimezoneLocation**](TimezoneLocation.md) |  | [optional] 
**Astronomy** | Pointer to [**Astronomy**](Astronomy.md) |  | [optional] 

## Methods

### NewAstronomyXMLResponse

`func NewAstronomyXMLResponse() *AstronomyXMLResponse`

NewAstronomyXMLResponse instantiates a new AstronomyXMLResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAstronomyXMLResponseWithDefaults

`func NewAstronomyXMLResponseWithDefaults() *AstronomyXMLResponse`

NewAstronomyXMLResponseWithDefaults instantiates a new AstronomyXMLResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *AstronomyXMLResponse) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *AstronomyXMLResponse) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *AstronomyXMLResponse) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *AstronomyXMLResponse) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetLocation

`func (o *AstronomyXMLResponse) GetLocation() TimezoneLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *AstronomyXMLResponse) GetLocationOk() (*TimezoneLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *AstronomyXMLResponse) SetLocation(v TimezoneLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *AstronomyXMLResponse) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetAstronomy

`func (o *AstronomyXMLResponse) GetAstronomy() Astronomy`

GetAstronomy returns the Astronomy field if non-nil, zero value otherwise.

### GetAstronomyOk

`func (o *AstronomyXMLResponse) GetAstronomyOk() (*Astronomy, bool)`

GetAstronomyOk returns a tuple with the Astronomy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAstronomy

`func (o *AstronomyXMLResponse) SetAstronomy(v Astronomy)`

SetAstronomy sets Astronomy field to given value.

### HasAstronomy

`func (o *AstronomyXMLResponse) HasAstronomy() bool`

HasAstronomy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


