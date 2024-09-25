// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"fmt"
)

type ErrorCode string

const (
	ErrorCodeValidation ErrorCode = "VALIDATION"
	ErrorCodeNotFound   ErrorCode = "NOT_FOUND"
	ErrorCodeInternal   ErrorCode = "INTERNAL"
)

func (e ErrorCode) ToPointer() *ErrorCode {
	return &e
}
func (e *ErrorCode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "VALIDATION":
		fallthrough
	case "NOT_FOUND":
		fallthrough
	case "INTERNAL":
		*e = ErrorCode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ErrorCode: %v", v)
	}
}

// Error - General error
type Error struct {
	ErrorCode    ErrorCode `json:"errorCode"`
	ErrorMessage string    `json:"errorMessage"`
}

var _ error = &Error{}

func (e *Error) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
