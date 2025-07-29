package ipgeolocationsdk

import (
	"encoding/json"
	"fmt"
)

// BulkIPGeolocation represents a single response in a bulk IP geolocation request.
// It can either be a GeolocationResponse (success) or an ErrorResponse (failure).
type BulkIPGeolocation struct {
	ErrorResponse       *ErrorResponse       
	GeolocationResponse *GeolocationResponse
}

// ErrorResponseAsBulkIPGeolocation creates a BulkIPGeolocation with an error
func ErrorResponseAsBulkIPGeolocation(v *ErrorResponse) BulkIPGeolocation {
	return BulkIPGeolocation{
		ErrorResponse: v,
	}
}

// GeolocationResponseAsBulkIPGeolocation creates a BulkIPGeolocation with a successful result
func GeolocationResponseAsBulkIPGeolocation(v *GeolocationResponse) BulkIPGeolocation {
	return BulkIPGeolocation{
		GeolocationResponse: v,
	}
}

// UnmarshalJSON unmarshals the correct object type from the raw JSON
func (dst *BulkIPGeolocation) UnmarshalJSON(data []byte) error {
	var err error
	match := 0

	err = json.Unmarshal(data, &dst.ErrorResponse)
	if err == nil && dst.ErrorResponse != nil && dst.ErrorResponse.Message != nil {
		match++
	} else {
		dst.ErrorResponse = nil
	}

	err = json.Unmarshal(data, &dst.GeolocationResponse)
	if err == nil && dst.GeolocationResponse != nil && dst.GeolocationResponse.Ip != nil {
		match++
	} else {
		dst.GeolocationResponse = nil
	}

	if match > 1 {
		dst.ErrorResponse = nil
		dst.GeolocationResponse = nil
		return fmt.Errorf("data matches more than one schema in oneOf(BulkIPGeolocation)")
	} else if match == 1 {
		return nil
	}
	return fmt.Errorf("data failed to match schemas in oneOf(BulkIPGeolocation)")
}

// MarshalJSON serializes only the set field
func (src BulkIPGeolocation) MarshalJSON() ([]byte, error) {
	if src.ErrorResponse != nil {
		return json.Marshal(&src.ErrorResponse)
	}
	if src.GeolocationResponse != nil {
		return json.Marshal(&src.GeolocationResponse)
	}
	return nil, nil
}

// IsError returns true if the response is an error
func (r *BulkIPGeolocation) IsError() bool {
	return r != nil && r.ErrorResponse != nil
}

// IsSuccess returns true if the response is a success
func (r *BulkIPGeolocation) IsSuccess() bool {
	return r != nil && r.GeolocationResponse != nil
}

// AsResult returns the two possible inner structs
func (r *BulkIPGeolocation) AsResult() (*GeolocationResponse, *ErrorResponse) {
	if r == nil {
		return nil, nil
	}
	return r.GeolocationResponse, r.ErrorResponse
}
type NullableBulkIPGeolocation struct {
	value *BulkIPGeolocation
	isSet bool
}

func (v NullableBulkIPGeolocation) Get() *BulkIPGeolocation {
	return v.value
}

func (v *NullableBulkIPGeolocation) Set(val *BulkIPGeolocation) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkIPGeolocation) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkIPGeolocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkIPGeolocation(val *BulkIPGeolocation) *NullableBulkIPGeolocation {
	return &NullableBulkIPGeolocation{value: val, isSet: true}
}

func (v NullableBulkIPGeolocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkIPGeolocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
