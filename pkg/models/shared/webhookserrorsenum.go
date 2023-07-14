// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type WebhooksErrorsEnum string

const (
	WebhooksErrorsEnumInternal          WebhooksErrorsEnum = "INTERNAL"
	WebhooksErrorsEnumInsufficientFund  WebhooksErrorsEnum = "INSUFFICIENT_FUND"
	WebhooksErrorsEnumValidation        WebhooksErrorsEnum = "VALIDATION"
	WebhooksErrorsEnumConflict          WebhooksErrorsEnum = "CONFLICT"
	WebhooksErrorsEnumNoScript          WebhooksErrorsEnum = "NO_SCRIPT"
	WebhooksErrorsEnumCompilationFailed WebhooksErrorsEnum = "COMPILATION_FAILED"
	WebhooksErrorsEnumMetadataOverride  WebhooksErrorsEnum = "METADATA_OVERRIDE"
	WebhooksErrorsEnumNotFound          WebhooksErrorsEnum = "NOT_FOUND"
	WebhooksErrorsEnumContextCancelled  WebhooksErrorsEnum = "CONTEXT_CANCELLED"
	WebhooksErrorsEnumStore             WebhooksErrorsEnum = "STORE"
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
	case "INSUFFICIENT_FUND":
		fallthrough
	case "VALIDATION":
		fallthrough
	case "CONFLICT":
		fallthrough
	case "NO_SCRIPT":
		fallthrough
	case "COMPILATION_FAILED":
		fallthrough
	case "METADATA_OVERRIDE":
		fallthrough
	case "NOT_FOUND":
		fallthrough
	case "CONTEXT_CANCELLED":
		fallthrough
	case "STORE":
		*e = WebhooksErrorsEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for WebhooksErrorsEnum: %v", v)
	}
}