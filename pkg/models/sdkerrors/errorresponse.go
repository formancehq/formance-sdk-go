// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/shared"
)

// ErrorResponse - Error
type ErrorResponse struct {
	Details      *string           `json:"details,omitempty"`
	ErrorCode    shared.ErrorsEnum `json:"errorCode"`
	ErrorMessage string            `json:"errorMessage"`
}

var _ error = &ErrorResponse{}

func (e *ErrorResponse) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
