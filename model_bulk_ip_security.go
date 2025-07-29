package ipgeolocationsdk

import (
	"encoding/json"
	"fmt"
)

type BulkIpSecurity struct {
	Error    *ErrorResponse
	SecurityResponse *SecurityAPIResponse
}

// Helper to wrap Error
func WrapError(err *ErrorResponse) BulkIpSecurity {
	return BulkIpSecurity{
		Error: err,
	}
}

// Helper to wrap Success
func WrapSecurity(sec *SecurityAPIResponse) BulkIpSecurity {
	return BulkIpSecurity{
		SecurityResponse: sec,
	}
}

// JSON unmarshal
func (b *BulkIpSecurity) UnmarshalJSON(data []byte) error {
	var err error
	match := 0

	// Try Error
	errResp := new(ErrorResponse)
	err = json.Unmarshal(data, errResp)
	if err == nil && !isEmptyStruct(errResp) {
		b.Error = errResp
		match++
	}

	// Try Security
	secResp := new(SecurityAPIResponse)
	err = json.Unmarshal(data, secResp)
	if err == nil && !isEmptyStruct(secResp) {
		b.SecurityResponse= secResp
		match++
	}

	if match == 1 {
		return nil
	}

	// Reset
	b.Error = nil
	b.SecurityResponse = nil

	if match > 1 {
		return fmt.Errorf("data matches more than one schema in oneOf(BulkIpSecurity)")
	}
	return fmt.Errorf("data failed to match schemas in oneOf(BulkIpSecurity)")
}

// JSON marshal
func (b BulkIpSecurity) MarshalJSON() ([]byte, error) {
	if b.Error != nil {
		return json.Marshal(b.Error)
	}
	if b.SecurityResponse != nil {
		return json.Marshal(b.SecurityResponse)
	}
	return nil, nil
}

// Helpers
func (b *BulkIpSecurity) IsError() bool {
	return b != nil && b.Error != nil
}

func (b *BulkIpSecurity) IsSuccess() bool {
	return b != nil && b.SecurityResponse != nil
}

func (b *BulkIpSecurity) AsResult() (*SecurityAPIResponse, *ErrorResponse) {
	if b == nil {
		return nil, nil
	}
	return b.SecurityResponse, b.Error
}

func isEmptyStruct(v interface{}) bool {
	b, err := json.Marshal(v)
	return err == nil && string(b) == "{}"
}
