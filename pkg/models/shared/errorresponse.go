// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ErrorResponse - Error
type ErrorResponse struct {
	Details      *string     `json:"details,omitempty"`
	ErrorCode    *ErrorsEnum `json:"errorCode,omitempty"`
	ErrorMessage *string     `json:"errorMessage,omitempty"`
}