# TimeZone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Offset** | Pointer to **int32** |  | [optional] 
**OffsetWithDst** | Pointer to **int32** |  | [optional] 
**CurrentTime** | Pointer to **string** |  | [optional] 
**CurrentTimeUnix** | Pointer to **float32** |  | [optional] 
**IsDst** | Pointer to **bool** |  | [optional] 
**DstSavings** | Pointer to **int32** |  | [optional] 
**DstExists** | Pointer to **bool** |  | [optional] 
**DstStart** | Pointer to [**TimeZoneDstStart**](TimeZoneDstStart.md) |  | [optional] 
**DstEnd** | Pointer to [**TimeZoneDstEnd**](TimeZoneDstEnd.md) |  | [optional] 

## Methods

### NewTimeZone

`func NewTimeZone() *TimeZone`

NewTimeZone instantiates a new TimeZone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeZoneWithDefaults

`func NewTimeZoneWithDefaults() *TimeZone`

NewTimeZoneWithDefaults instantiates a new TimeZone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TimeZone) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TimeZone) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TimeZone) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TimeZone) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOffset

`func (o *TimeZone) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *TimeZone) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *TimeZone) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *TimeZone) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetOffsetWithDst

`func (o *TimeZone) GetOffsetWithDst() int32`

GetOffsetWithDst returns the OffsetWithDst field if non-nil, zero value otherwise.

### GetOffsetWithDstOk

`func (o *TimeZone) GetOffsetWithDstOk() (*int32, bool)`

GetOffsetWithDstOk returns a tuple with the OffsetWithDst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffsetWithDst

`func (o *TimeZone) SetOffsetWithDst(v int32)`

SetOffsetWithDst sets OffsetWithDst field to given value.

### HasOffsetWithDst

`func (o *TimeZone) HasOffsetWithDst() bool`

HasOffsetWithDst returns a boolean if a field has been set.

### GetCurrentTime

`func (o *TimeZone) GetCurrentTime() string`

GetCurrentTime returns the CurrentTime field if non-nil, zero value otherwise.

### GetCurrentTimeOk

`func (o *TimeZone) GetCurrentTimeOk() (*string, bool)`

GetCurrentTimeOk returns a tuple with the CurrentTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentTime

`func (o *TimeZone) SetCurrentTime(v string)`

SetCurrentTime sets CurrentTime field to given value.

### HasCurrentTime

`func (o *TimeZone) HasCurrentTime() bool`

HasCurrentTime returns a boolean if a field has been set.

### GetCurrentTimeUnix

`func (o *TimeZone) GetCurrentTimeUnix() float32`

GetCurrentTimeUnix returns the CurrentTimeUnix field if non-nil, zero value otherwise.

### GetCurrentTimeUnixOk

`func (o *TimeZone) GetCurrentTimeUnixOk() (*float32, bool)`

GetCurrentTimeUnixOk returns a tuple with the CurrentTimeUnix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentTimeUnix

`func (o *TimeZone) SetCurrentTimeUnix(v float32)`

SetCurrentTimeUnix sets CurrentTimeUnix field to given value.

### HasCurrentTimeUnix

`func (o *TimeZone) HasCurrentTimeUnix() bool`

HasCurrentTimeUnix returns a boolean if a field has been set.

### GetIsDst

`func (o *TimeZone) GetIsDst() bool`

GetIsDst returns the IsDst field if non-nil, zero value otherwise.

### GetIsDstOk

`func (o *TimeZone) GetIsDstOk() (*bool, bool)`

GetIsDstOk returns a tuple with the IsDst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDst

`func (o *TimeZone) SetIsDst(v bool)`

SetIsDst sets IsDst field to given value.

### HasIsDst

`func (o *TimeZone) HasIsDst() bool`

HasIsDst returns a boolean if a field has been set.

### GetDstSavings

`func (o *TimeZone) GetDstSavings() int32`

GetDstSavings returns the DstSavings field if non-nil, zero value otherwise.

### GetDstSavingsOk

`func (o *TimeZone) GetDstSavingsOk() (*int32, bool)`

GetDstSavingsOk returns a tuple with the DstSavings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstSavings

`func (o *TimeZone) SetDstSavings(v int32)`

SetDstSavings sets DstSavings field to given value.

### HasDstSavings

`func (o *TimeZone) HasDstSavings() bool`

HasDstSavings returns a boolean if a field has been set.

### GetDstExists

`func (o *TimeZone) GetDstExists() bool`

GetDstExists returns the DstExists field if non-nil, zero value otherwise.

### GetDstExistsOk

`func (o *TimeZone) GetDstExistsOk() (*bool, bool)`

GetDstExistsOk returns a tuple with the DstExists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstExists

`func (o *TimeZone) SetDstExists(v bool)`

SetDstExists sets DstExists field to given value.

### HasDstExists

`func (o *TimeZone) HasDstExists() bool`

HasDstExists returns a boolean if a field has been set.

### GetDstStart

`func (o *TimeZone) GetDstStart() TimeZoneDstStart`

GetDstStart returns the DstStart field if non-nil, zero value otherwise.

### GetDstStartOk

`func (o *TimeZone) GetDstStartOk() (*TimeZoneDstStart, bool)`

GetDstStartOk returns a tuple with the DstStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstStart

`func (o *TimeZone) SetDstStart(v TimeZoneDstStart)`

SetDstStart sets DstStart field to given value.

### HasDstStart

`func (o *TimeZone) HasDstStart() bool`

HasDstStart returns a boolean if a field has been set.

### GetDstEnd

`func (o *TimeZone) GetDstEnd() TimeZoneDstEnd`

GetDstEnd returns the DstEnd field if non-nil, zero value otherwise.

### GetDstEndOk

`func (o *TimeZone) GetDstEndOk() (*TimeZoneDstEnd, bool)`

GetDstEndOk returns a tuple with the DstEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstEnd

`func (o *TimeZone) SetDstEnd(v TimeZoneDstEnd)`

SetDstEnd sets DstEnd field to given value.

### HasDstEnd

`func (o *TimeZone) HasDstEnd() bool`

HasDstEnd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


