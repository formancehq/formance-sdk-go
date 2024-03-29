// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"fmt"
)

type SchemasErrorCode string

const (
	SchemasErrorCodeValidation       SchemasErrorCode = "VALIDATION"
	SchemasErrorCodeInternalError    SchemasErrorCode = "INTERNAL_ERROR"
	SchemasErrorCodeInsufficientFund SchemasErrorCode = "INSUFFICIENT_FUND"
	SchemasErrorCodeHoldClosed       SchemasErrorCode = "HOLD_CLOSED"
)

func (e SchemasErrorCode) ToPointer() *SchemasErrorCode {
	return &e
}

func (e *SchemasErrorCode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "VALIDATION":
		fallthrough
	case "INTERNAL_ERROR":
		fallthrough
	case "INSUFFICIENT_FUND":
		fallthrough
	case "HOLD_CLOSED":
		*e = SchemasErrorCode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SchemasErrorCode: %v", v)
	}
}

// WalletsErrorResponse - Error
type WalletsErrorResponse struct {
	ErrorCode    SchemasErrorCode `json:"errorCode"`
	ErrorMessage string           `json:"errorMessage"`
}

var _ error = &WalletsErrorResponse{}

func (e *WalletsErrorResponse) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
