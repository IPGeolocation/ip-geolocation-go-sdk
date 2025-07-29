# TimezoneDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Offset** | Pointer to **int32** |  | [optional] 
**OffsetWithDst** | Pointer to **int32** |  | [optional] 
**Date** | Pointer to **string** |  | [optional] 
**DateTime** | Pointer to **string** |  | [optional] 
**DateTimeTxt** | Pointer to **string** |  | [optional] 
**DateTimeWti** | Pointer to **string** |  | [optional] 
**DateTimeYmd** | Pointer to **string** |  | [optional] 
**DateTimeUnix** | Pointer to **float64** |  | [optional] 
**Time24** | Pointer to **string** |  | [optional] 
**Time12** | Pointer to **string** |  | [optional] 
**Week** | Pointer to **int32** |  | [optional] 
**Month** | Pointer to **int32** |  | [optional] 
**Year** | Pointer to **int32** |  | [optional] 
**YearAbbr** | Pointer to **string** |  | [optional] 
**IsDst** | Pointer to **bool** |  | [optional] 
**DstSavings** | Pointer to **int32** |  | [optional] 
**DstExists** | Pointer to **bool** |  | [optional] 
**DstStart** | Pointer to [**TimezoneDetailDstStart**](TimezoneDetailDstStart.md) |  | [optional] 
**DstEnd** | Pointer to [**TimezoneDetailDstEnd**](TimezoneDetailDstEnd.md) |  | [optional] 

## Methods

### NewTimezoneDetail

`func NewTimezoneDetail() *TimezoneDetail`

NewTimezoneDetail instantiates a new TimezoneDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimezoneDetailWithDefaults

`func NewTimezoneDetailWithDefaults() *TimezoneDetail`

NewTimezoneDetailWithDefaults instantiates a new TimezoneDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TimezoneDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TimezoneDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TimezoneDetail) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TimezoneDetail) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOffset

`func (o *TimezoneDetail) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *TimezoneDetail) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *TimezoneDetail) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *TimezoneDetail) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetOffsetWithDst

`func (o *TimezoneDetail) GetOffsetWithDst() int32`

GetOffsetWithDst returns the OffsetWithDst field if non-nil, zero value otherwise.

### GetOffsetWithDstOk

`func (o *TimezoneDetail) GetOffsetWithDstOk() (*int32, bool)`

GetOffsetWithDstOk returns a tuple with the OffsetWithDst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffsetWithDst

`func (o *TimezoneDetail) SetOffsetWithDst(v int32)`

SetOffsetWithDst sets OffsetWithDst field to given value.

### HasOffsetWithDst

`func (o *TimezoneDetail) HasOffsetWithDst() bool`

HasOffsetWithDst returns a boolean if a field has been set.

### GetDate

`func (o *TimezoneDetail) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *TimezoneDetail) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *TimezoneDetail) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *TimezoneDetail) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetDateTime

`func (o *TimezoneDetail) GetDateTime() string`

GetDateTime returns the DateTime field if non-nil, zero value otherwise.

### GetDateTimeOk

`func (o *TimezoneDetail) GetDateTimeOk() (*string, bool)`

GetDateTimeOk returns a tuple with the DateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTime

`func (o *TimezoneDetail) SetDateTime(v string)`

SetDateTime sets DateTime field to given value.

### HasDateTime

`func (o *TimezoneDetail) HasDateTime() bool`

HasDateTime returns a boolean if a field has been set.

### GetDateTimeTxt

`func (o *TimezoneDetail) GetDateTimeTxt() string`

GetDateTimeTxt returns the DateTimeTxt field if non-nil, zero value otherwise.

### GetDateTimeTxtOk

`func (o *TimezoneDetail) GetDateTimeTxtOk() (*string, bool)`

GetDateTimeTxtOk returns a tuple with the DateTimeTxt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTimeTxt

`func (o *TimezoneDetail) SetDateTimeTxt(v string)`

SetDateTimeTxt sets DateTimeTxt field to given value.

### HasDateTimeTxt

`func (o *TimezoneDetail) HasDateTimeTxt() bool`

HasDateTimeTxt returns a boolean if a field has been set.

### GetDateTimeWti

`func (o *TimezoneDetail) GetDateTimeWti() string`

GetDateTimeWti returns the DateTimeWti field if non-nil, zero value otherwise.

### GetDateTimeWtiOk

`func (o *TimezoneDetail) GetDateTimeWtiOk() (*string, bool)`

GetDateTimeWtiOk returns a tuple with the DateTimeWti field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTimeWti

`func (o *TimezoneDetail) SetDateTimeWti(v string)`

SetDateTimeWti sets DateTimeWti field to given value.

### HasDateTimeWti

`func (o *TimezoneDetail) HasDateTimeWti() bool`

HasDateTimeWti returns a boolean if a field has been set.

### GetDateTimeYmd

`func (o *TimezoneDetail) GetDateTimeYmd() string`

GetDateTimeYmd returns the DateTimeYmd field if non-nil, zero value otherwise.

### GetDateTimeYmdOk

`func (o *TimezoneDetail) GetDateTimeYmdOk() (*string, bool)`

GetDateTimeYmdOk returns a tuple with the DateTimeYmd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTimeYmd

`func (o *TimezoneDetail) SetDateTimeYmd(v string)`

SetDateTimeYmd sets DateTimeYmd field to given value.

### HasDateTimeYmd

`func (o *TimezoneDetail) HasDateTimeYmd() bool`

HasDateTimeYmd returns a boolean if a field has been set.

### GetDateTimeUnix

`func (o *TimezoneDetail) GetDateTimeUnix() float64`

GetDateTimeUnix returns the DateTimeUnix field if non-nil, zero value otherwise.

### GetDateTimeUnixOk

`func (o *TimezoneDetail) GetDateTimeUnixOk() (*float64, bool)`

GetDateTimeUnixOk returns a tuple with the DateTimeUnix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTimeUnix

`func (o *TimezoneDetail) SetDateTimeUnix(v float64)`

SetDateTimeUnix sets DateTimeUnix field to given value.

### HasDateTimeUnix

`func (o *TimezoneDetail) HasDateTimeUnix() bool`

HasDateTimeUnix returns a boolean if a field has been set.

### GetTime24

`func (o *TimezoneDetail) GetTime24() string`

GetTime24 returns the Time24 field if non-nil, zero value otherwise.

### GetTime24Ok

`func (o *TimezoneDetail) GetTime24Ok() (*string, bool)`

GetTime24Ok returns a tuple with the Time24 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime24

`func (o *TimezoneDetail) SetTime24(v string)`

SetTime24 sets Time24 field to given value.

### HasTime24

`func (o *TimezoneDetail) HasTime24() bool`

HasTime24 returns a boolean if a field has been set.

### GetTime12

`func (o *TimezoneDetail) GetTime12() string`

GetTime12 returns the Time12 field if non-nil, zero value otherwise.

### GetTime12Ok

`func (o *TimezoneDetail) GetTime12Ok() (*string, bool)`

GetTime12Ok returns a tuple with the Time12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime12

`func (o *TimezoneDetail) SetTime12(v string)`

SetTime12 sets Time12 field to given value.

### HasTime12

`func (o *TimezoneDetail) HasTime12() bool`

HasTime12 returns a boolean if a field has been set.

### GetWeek

`func (o *TimezoneDetail) GetWeek() int32`

GetWeek returns the Week field if non-nil, zero value otherwise.

### GetWeekOk

`func (o *TimezoneDetail) GetWeekOk() (*int32, bool)`

GetWeekOk returns a tuple with the Week field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeek

`func (o *TimezoneDetail) SetWeek(v int32)`

SetWeek sets Week field to given value.

### HasWeek

`func (o *TimezoneDetail) HasWeek() bool`

HasWeek returns a boolean if a field has been set.

### GetMonth

`func (o *TimezoneDetail) GetMonth() int32`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *TimezoneDetail) GetMonthOk() (*int32, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *TimezoneDetail) SetMonth(v int32)`

SetMonth sets Month field to given value.

### HasMonth

`func (o *TimezoneDetail) HasMonth() bool`

HasMonth returns a boolean if a field has been set.

### GetYear

`func (o *TimezoneDetail) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *TimezoneDetail) GetYearOk() (*int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *TimezoneDetail) SetYear(v int32)`

SetYear sets Year field to given value.

### HasYear

`func (o *TimezoneDetail) HasYear() bool`

HasYear returns a boolean if a field has been set.

### GetYearAbbr

`func (o *TimezoneDetail) GetYearAbbr() string`

GetYearAbbr returns the YearAbbr field if non-nil, zero value otherwise.

### GetYearAbbrOk

`func (o *TimezoneDetail) GetYearAbbrOk() (*string, bool)`

GetYearAbbrOk returns a tuple with the YearAbbr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYearAbbr

`func (o *TimezoneDetail) SetYearAbbr(v string)`

SetYearAbbr sets YearAbbr field to given value.

### HasYearAbbr

`func (o *TimezoneDetail) HasYearAbbr() bool`

HasYearAbbr returns a boolean if a field has been set.

### GetIsDst

`func (o *TimezoneDetail) GetIsDst() bool`

GetIsDst returns the IsDst field if non-nil, zero value otherwise.

### GetIsDstOk

`func (o *TimezoneDetail) GetIsDstOk() (*bool, bool)`

GetIsDstOk returns a tuple with the IsDst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDst

`func (o *TimezoneDetail) SetIsDst(v bool)`

SetIsDst sets IsDst field to given value.

### HasIsDst

`func (o *TimezoneDetail) HasIsDst() bool`

HasIsDst returns a boolean if a field has been set.

### GetDstSavings

`func (o *TimezoneDetail) GetDstSavings() int32`

GetDstSavings returns the DstSavings field if non-nil, zero value otherwise.

### GetDstSavingsOk

`func (o *TimezoneDetail) GetDstSavingsOk() (*int32, bool)`

GetDstSavingsOk returns a tuple with the DstSavings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstSavings

`func (o *TimezoneDetail) SetDstSavings(v int32)`

SetDstSavings sets DstSavings field to given value.

### HasDstSavings

`func (o *TimezoneDetail) HasDstSavings() bool`

HasDstSavings returns a boolean if a field has been set.

### GetDstExists

`func (o *TimezoneDetail) GetDstExists() bool`

GetDstExists returns the DstExists field if non-nil, zero value otherwise.

### GetDstExistsOk

`func (o *TimezoneDetail) GetDstExistsOk() (*bool, bool)`

GetDstExistsOk returns a tuple with the DstExists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstExists

`func (o *TimezoneDetail) SetDstExists(v bool)`

SetDstExists sets DstExists field to given value.

### HasDstExists

`func (o *TimezoneDetail) HasDstExists() bool`

HasDstExists returns a boolean if a field has been set.

### GetDstStart

`func (o *TimezoneDetail) GetDstStart() TimezoneDetailDstStart`

GetDstStart returns the DstStart field if non-nil, zero value otherwise.

### GetDstStartOk

`func (o *TimezoneDetail) GetDstStartOk() (*TimezoneDetailDstStart, bool)`

GetDstStartOk returns a tuple with the DstStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstStart

`func (o *TimezoneDetail) SetDstStart(v TimezoneDetailDstStart)`

SetDstStart sets DstStart field to given value.

### HasDstStart

`func (o *TimezoneDetail) HasDstStart() bool`

HasDstStart returns a boolean if a field has been set.

### GetDstEnd

`func (o *TimezoneDetail) GetDstEnd() TimezoneDetailDstEnd`

GetDstEnd returns the DstEnd field if non-nil, zero value otherwise.

### GetDstEndOk

`func (o *TimezoneDetail) GetDstEndOk() (*TimezoneDetailDstEnd, bool)`

GetDstEndOk returns a tuple with the DstEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstEnd

`func (o *TimezoneDetail) SetDstEnd(v TimezoneDetailDstEnd)`

SetDstEnd sets DstEnd field to given value.

### HasDstEnd

`func (o *TimezoneDetail) HasDstEnd() bool`

HasDstEnd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


