// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type WebhooksErrorsEnum string

const (
	WebhooksErrorsEnumInternal   WebhooksErrorsEnum = "INTERNAL"
	WebhooksErrorsEnumValidation WebhooksErrorsEnum = "VALIDATION"
	WebhooksErrorsEnumNotFound   WebhooksErrorsEnum = "NOT_FOUND"
)

func (e WebhooksErrorsEnum) ToPointer() *WebhooksErrorsEnum {
	return &e
}
func (e *WebhooksErrorsEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "INTERNAL":
		fallthrough
	case "VALIDATION":
		fallthrough
	case "NOT_FOUND":
		*e = WebhooksErrorsEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for WebhooksErrorsEnum: %v", v)
	}
}
