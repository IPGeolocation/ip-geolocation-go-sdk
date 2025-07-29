# Astronomy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeZone** | Pointer to **string** |  | [optional] 
**Date** | Pointer to **string** |  | [optional] 
**CurrentTime** | Pointer to **string** |  | [optional] 
**MidNight** | Pointer to **string** |  | [optional] 
**NightEnd** | Pointer to **string** |  | [optional] 
**Morning** | Pointer to [**AstronomyMorning**](AstronomyMorning.md) |  | [optional] 
**Sunrise** | Pointer to **string** |  | [optional] 
**Sunset** | Pointer to **string** |  | [optional] 
**Evening** | Pointer to [**AstronomyEvening**](AstronomyEvening.md) |  | [optional] 
**NightBegin** | Pointer to **string** |  | [optional] 
**SunStatus** | Pointer to **string** |  | [optional] 
**SolarNoon** | Pointer to **string** |  | [optional] 
**DayLength** | Pointer to **string** |  | [optional] 
**SunAltitude** | Pointer to **float32** |  | [optional] 
**SunDistance** | Pointer to **float32** |  | [optional] 
**SunAzimuth** | Pointer to **float32** |  | [optional] 
**MoonPhase** | Pointer to **string** |  | [optional] 
**Moonrise** | Pointer to **string** |  | [optional] 
**Moonset** | Pointer to **string** |  | [optional] 
**MoonStatus** | Pointer to **string** |  | [optional] 
**MoonAltitude** | Pointer to **float32** |  | [optional] 
**MoonDistance** | Pointer to **float32** |  | [optional] 
**MoonAzimuth** | Pointer to **float32** |  | [optional] 
**MoonParallacticAngle** | Pointer to **float32** |  | [optional] 
**MoonIlluminationPercentage** | Pointer to **string** |  | [optional] 
**MoonAngle** | Pointer to **float32** |  | [optional] 

## Methods

### NewAstronomy

`func NewAstronomy() *Astronomy`

NewAstronomy instantiates a new Astronomy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAstronomyWithDefaults

`func NewAstronomyWithDefaults() *Astronomy`

NewAstronomyWithDefaults instantiates a new Astronomy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeZone

`func (o *Astronomy) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *Astronomy) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *Astronomy) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *Astronomy) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetDate

`func (o *Astronomy) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *Astronomy) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *Astronomy) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *Astronomy) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetCurrentTime

`func (o *Astronomy) GetCurrentTime() string`

GetCurrentTime returns the CurrentTime field if non-nil, zero value otherwise.

### GetCurrentTimeOk

`func (o *Astronomy) GetCurrentTimeOk() (*string, bool)`

GetCurrentTimeOk returns a tuple with the CurrentTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentTime

`func (o *Astronomy) SetCurrentTime(v string)`

SetCurrentTime sets CurrentTime field to given value.

### HasCurrentTime

`func (o *Astronomy) HasCurrentTime() bool`

HasCurrentTime returns a boolean if a field has been set.

### GetMidNight

`func (o *Astronomy) GetMidNight() string`

GetMidNight returns the MidNight field if non-nil, zero value otherwise.

### GetMidNightOk

`func (o *Astronomy) GetMidNightOk() (*string, bool)`

GetMidNightOk returns a tuple with the MidNight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMidNight

`func (o *Astronomy) SetMidNight(v string)`

SetMidNight sets MidNight field to given value.

### HasMidNight

`func (o *Astronomy) HasMidNight() bool`

HasMidNight returns a boolean if a field has been set.

### GetNightEnd

`func (o *Astronomy) GetNightEnd() string`

GetNightEnd returns the NightEnd field if non-nil, zero value otherwise.

### GetNightEndOk

`func (o *Astronomy) GetNightEndOk() (*string, bool)`

GetNightEndOk returns a tuple with the NightEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNightEnd

`func (o *Astronomy) SetNightEnd(v string)`

SetNightEnd sets NightEnd field to given value.

### HasNightEnd

`func (o *Astronomy) HasNightEnd() bool`

HasNightEnd returns a boolean if a field has been set.

### GetMorning

`func (o *Astronomy) GetMorning() AstronomyMorning`

GetMorning returns the Morning field if non-nil, zero value otherwise.

### GetMorningOk

`func (o *Astronomy) GetMorningOk() (*AstronomyMorning, bool)`

GetMorningOk returns a tuple with the Morning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMorning

`func (o *Astronomy) SetMorning(v AstronomyMorning)`

SetMorning sets Morning field to given value.

### HasMorning

`func (o *Astronomy) HasMorning() bool`

HasMorning returns a boolean if a field has been set.

### GetSunrise

`func (o *Astronomy) GetSunrise() string`

GetSunrise returns the Sunrise field if non-nil, zero value otherwise.

### GetSunriseOk

`func (o *Astronomy) GetSunriseOk() (*string, bool)`

GetSunriseOk returns a tuple with the Sunrise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSunrise

`func (o *Astronomy) SetSunrise(v string)`

SetSunrise sets Sunrise field to given value.

### HasSunrise

`func (o *Astronomy) HasSunrise() bool`

HasSunrise returns a boolean if a field has been set.

### GetSunset

`func (o *Astronomy) GetSunset() string`

GetSunset returns the Sunset field if non-nil, zero value otherwise.

### GetSunsetOk

`func (o *Astronomy) GetSunsetOk() (*string, bool)`

GetSunsetOk returns a tuple with the Sunset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSunset

`func (o *Astronomy) SetSunset(v string)`

SetSunset sets Sunset field to given value.

### HasSunset

`func (o *Astronomy) HasSunset() bool`

HasSunset returns a boolean if a field has been set.

### GetEvening

`func (o *Astronomy) GetEvening() AstronomyEvening`

GetEvening returns the Evening field if non-nil, zero value otherwise.

### GetEveningOk

`func (o *Astronomy) GetEveningOk() (*AstronomyEvening, bool)`

GetEveningOk returns a tuple with the Evening field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvening

`func (o *Astronomy) SetEvening(v AstronomyEvening)`

SetEvening sets Evening field to given value.

### HasEvening

`func (o *Astronomy) HasEvening() bool`

HasEvening returns a boolean if a field has been set.

### GetNightBegin

`func (o *Astronomy) GetNightBegin() string`

GetNightBegin returns the NightBegin field if non-nil, zero value otherwise.

### GetNightBeginOk

`func (o *Astronomy) GetNightBeginOk() (*string, bool)`

GetNightBeginOk returns a tuple with the NightBegin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNightBegin

`func (o *Astronomy) SetNightBegin(v string)`

SetNightBegin sets NightBegin field to given value.

### HasNightBegin

`func (o *Astronomy) HasNightBegin() bool`

HasNightBegin returns a boolean if a field has been set.

### GetSunStatus

`func (o *Astronomy) GetSunStatus() string`

GetSunStatus returns the SunStatus field if non-nil, zero value otherwise.

### GetSunStatusOk

`func (o *Astronomy) GetSunStatusOk() (*string, bool)`

GetSunStatusOk returns a tuple with the SunStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSunStatus

`func (o *Astronomy) SetSunStatus(v string)`

SetSunStatus sets SunStatus field to given value.

### HasSunStatus

`func (o *Astronomy) HasSunStatus() bool`

HasSunStatus returns a boolean if a field has been set.

### GetSolarNoon

`func (o *Astronomy) GetSolarNoon() string`

GetSolarNoon returns the SolarNoon field if non-nil, zero value otherwise.

### GetSolarNoonOk

`func (o *Astronomy) GetSolarNoonOk() (*string, bool)`

GetSolarNoonOk returns a tuple with the SolarNoon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolarNoon

`func (o *Astronomy) SetSolarNoon(v string)`

SetSolarNoon sets SolarNoon field to given value.

### HasSolarNoon

`func (o *Astronomy) HasSolarNoon() bool`

HasSolarNoon returns a boolean if a field has been set.

### GetDayLength

`func (o *Astronomy) GetDayLength() string`

GetDayLength returns the DayLength field if non-nil, zero value otherwise.

### GetDayLengthOk

`func (o *Astronomy) GetDayLengthOk() (*string, bool)`

GetDayLengthOk returns a tuple with the DayLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayLength

`func (o *Astronomy) SetDayLength(v string)`

SetDayLength sets DayLength field to given value.

### HasDayLength

`func (o *Astronomy) HasDayLength() bool`

HasDayLength returns a boolean if a field has been set.

### GetSunAltitude

`func (o *Astronomy) GetSunAltitude() float32`

GetSunAltitude returns the SunAltitude field if non-nil, zero value otherwise.

### GetSunAltitudeOk

`func (o *Astronomy) GetSunAltitudeOk() (*float32, bool)`

GetSunAltitudeOk returns a tuple with the SunAltitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSunAltitude

`func (o *Astronomy) SetSunAltitude(v float32)`

SetSunAltitude sets SunAltitude field to given value.

### HasSunAltitude

`func (o *Astronomy) HasSunAltitude() bool`

HasSunAltitude returns a boolean if a field has been set.

### GetSunDistance

`func (o *Astronomy) GetSunDistance() float32`

GetSunDistance returns the SunDistance field if non-nil, zero value otherwise.

### GetSunDistanceOk

`func (o *Astronomy) GetSunDistanceOk() (*float32, bool)`

GetSunDistanceOk returns a tuple with the SunDistance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSunDistance

`func (o *Astronomy) SetSunDistance(v float32)`

SetSunDistance sets SunDistance field to given value.

### HasSunDistance

`func (o *Astronomy) HasSunDistance() bool`

HasSunDistance returns a boolean if a field has been set.

### GetSunAzimuth

`func (o *Astronomy) GetSunAzimuth() float32`

GetSunAzimuth returns the SunAzimuth field if non-nil, zero value otherwise.

### GetSunAzimuthOk

`func (o *Astronomy) GetSunAzimuthOk() (*float32, bool)`

GetSunAzimuthOk returns a tuple with the SunAzimuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSunAzimuth

`func (o *Astronomy) SetSunAzimuth(v float32)`

SetSunAzimuth sets SunAzimuth field to given value.

### HasSunAzimuth

`func (o *Astronomy) HasSunAzimuth() bool`

HasSunAzimuth returns a boolean if a field has been set.

### GetMoonPhase

`func (o *Astronomy) GetMoonPhase() string`

GetMoonPhase returns the MoonPhase field if non-nil, zero value otherwise.

### GetMoonPhaseOk

`func (o *Astronomy) GetMoonPhaseOk() (*string, bool)`

GetMoonPhaseOk returns a tuple with the MoonPhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoonPhase

`func (o *Astronomy) SetMoonPhase(v string)`

SetMoonPhase sets MoonPhase field to given value.

### HasMoonPhase

`func (o *Astronomy) HasMoonPhase() bool`

HasMoonPhase returns a boolean if a field has been set.

### GetMoonrise

`func (o *Astronomy) GetMoonrise() string`

GetMoonrise returns the Moonrise field if non-nil, zero value otherwise.

### GetMoonriseOk

`func (o *Astronomy) GetMoonriseOk() (*string, bool)`

GetMoonriseOk returns a tuple with the Moonrise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoonrise

`func (o *Astronomy) SetMoonrise(v string)`

SetMoonrise sets Moonrise field to given value.

### HasMoonrise

`func (o *Astronomy) HasMoonrise() bool`

HasMoonrise returns a boolean if a field has been set.

### GetMoonset

`func (o *Astronomy) GetMoonset() string`

GetMoonset returns the Moonset field if non-nil, zero value otherwise.

### GetMoonsetOk

`func (o *Astronomy) GetMoonsetOk() (*string, bool)`

GetMoonsetOk returns a tuple with the Moonset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoonset

`func (o *Astronomy) SetMoonset(v string)`

SetMoonset sets Moonset field to given value.

### HasMoonset

`func (o *Astronomy) HasMoonset() bool`

HasMoonset returns a boolean if a field has been set.

### GetMoonStatus

`func (o *Astronomy) GetMoonStatus() string`

GetMoonStatus returns the MoonStatus field if non-nil, zero value otherwise.

### GetMoonStatusOk

`func (o *Astronomy) GetMoonStatusOk() (*string, bool)`

GetMoonStatusOk returns a tuple with the MoonStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoonStatus

`func (o *Astronomy) SetMoonStatus(v string)`

SetMoonStatus sets MoonStatus field to given value.

### HasMoonStatus

`func (o *Astronomy) HasMoonStatus() bool`

HasMoonStatus returns a boolean if a field has been set.

### GetMoonAltitude

`func (o *Astronomy) GetMoonAltitude() float32`

GetMoonAltitude returns the MoonAltitude field if non-nil, zero value otherwise.

### GetMoonAltitudeOk

`func (o *Astronomy) GetMoonAltitudeOk() (*float32, bool)`

GetMoonAltitudeOk returns a tuple with the MoonAltitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoonAltitude

`func (o *Astronomy) SetMoonAltitude(v float32)`

SetMoonAltitude sets MoonAltitude field to given value.

### HasMoonAltitude

`func (o *Astronomy) HasMoonAltitude() bool`

HasMoonAltitude returns a boolean if a field has been set.

### GetMoonDistance

`func (o *Astronomy) GetMoonDistance() float32`

GetMoonDistance returns the MoonDistance field if non-nil, zero value otherwise.

### GetMoonDistanceOk

`func (o *Astronomy) GetMoonDistanceOk() (*float32, bool)`

GetMoonDistanceOk returns a tuple with the MoonDistance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoonDistance

`func (o *Astronomy) SetMoonDistance(v float32)`

SetMoonDistance sets MoonDistance field to given value.

### HasMoonDistance

`func (o *Astronomy) HasMoonDistance() bool`

HasMoonDistance returns a boolean if a field has been set.

### GetMoonAzimuth

`func (o *Astronomy) GetMoonAzimuth() float32`

GetMoonAzimuth returns the MoonAzimuth field if non-nil, zero value otherwise.

### GetMoonAzimuthOk

`func (o *Astronomy) GetMoonAzimuthOk() (*float32, bool)`

GetMoonAzimuthOk returns a tuple with the MoonAzimuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoonAzimuth

`func (o *Astronomy) SetMoonAzimuth(v float32)`

SetMoonAzimuth sets MoonAzimuth field to given value.

### HasMoonAzimuth

`func (o *Astronomy) HasMoonAzimuth() bool`

HasMoonAzimuth returns a boolean if a field has been set.

### GetMoonParallacticAngle

`func (o *Astronomy) GetMoonParallacticAngle() float32`

GetMoonParallacticAngle returns the MoonParallacticAngle field if non-nil, zero value otherwise.

### GetMoonParallacticAngleOk

`func (o *Astronomy) GetMoonParallacticAngleOk() (*float32, bool)`

GetMoonParallacticAngleOk returns a tuple with the MoonParallacticAngle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoonParallacticAngle

`func (o *Astronomy) SetMoonParallacticAngle(v float32)`

SetMoonParallacticAngle sets MoonParallacticAngle field to given value.

### HasMoonParallacticAngle

`func (o *Astronomy) HasMoonParallacticAngle() bool`

HasMoonParallacticAngle returns a boolean if a field has been set.

### GetMoonIlluminationPercentage

`func (o *Astronomy) GetMoonIlluminationPercentage() string`

GetMoonIlluminationPercentage returns the MoonIlluminationPercentage field if non-nil, zero value otherwise.

### GetMoonIlluminationPercentageOk

`func (o *Astronomy) GetMoonIlluminationPercentageOk() (*string, bool)`

GetMoonIlluminationPercentageOk returns a tuple with the MoonIlluminationPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoonIlluminationPercentage

`func (o *Astronomy) SetMoonIlluminationPercentage(v string)`

SetMoonIlluminationPercentage sets MoonIlluminationPercentage field to given value.

### HasMoonIlluminationPercentage

`func (o *Astronomy) HasMoonIlluminationPercentage() bool`

HasMoonIlluminationPercentage returns a boolean if a field has been set.

### GetMoonAngle

`func (o *Astronomy) GetMoonAngle() float32`

GetMoonAngle returns the MoonAngle field if non-nil, zero value otherwise.

### GetMoonAngleOk

`func (o *Astronomy) GetMoonAngleOk() (*float32, bool)`

GetMoonAngleOk returns a tuple with the MoonAngle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoonAngle

`func (o *Astronomy) SetMoonAngle(v float32)`

SetMoonAngle sets MoonAngle field to given value.

### HasMoonAngle

`func (o *Astronomy) HasMoonAngle() bool`

HasMoonAngle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


