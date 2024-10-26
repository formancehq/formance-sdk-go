// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"fmt"
)

type SchemasWalletsErrorResponseErrorCode string

const (
	SchemasWalletsErrorResponseErrorCodeValidation       SchemasWalletsErrorResponseErrorCode = "VALIDATION"
	SchemasWalletsErrorResponseErrorCodeInternalError    SchemasWalletsErrorResponseErrorCode = "INTERNAL_ERROR"
	SchemasWalletsErrorResponseErrorCodeInsufficientFund SchemasWalletsErrorResponseErrorCode = "INSUFFICIENT_FUND"
	SchemasWalletsErrorResponseErrorCodeHoldClosed       SchemasWalletsErrorResponseErrorCode = "HOLD_CLOSED"
)

func (e SchemasWalletsErrorResponseErrorCode) ToPointer() *SchemasWalletsErrorResponseErrorCode {
	return &e
}
func (e *SchemasWalletsErrorResponseErrorCode) UnmarshalJSON(data []byte) error {
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
		*e = SchemasWalletsErrorResponseErrorCode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SchemasWalletsErrorResponseErrorCode: %v", v)
	}
}

type WalletsErrorResponse struct {
	ErrorCode    SchemasWalletsErrorResponseErrorCode `json:"errorCode"`
	ErrorMessage string                               `json:"errorMessage"`
}

var _ error = &WalletsErrorResponse{}

func (e *WalletsErrorResponse) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
