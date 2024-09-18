// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type AccountType string

const (
	AccountTypeUnknown  AccountType = "UNKNOWN"
	AccountTypeInternal AccountType = "INTERNAL"
	AccountTypeExternal AccountType = "EXTERNAL"
)

func (e AccountType) ToPointer() *AccountType {
	return &e
}
func (e *AccountType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "UNKNOWN":
		fallthrough
	case "INTERNAL":
		fallthrough
	case "EXTERNAL":
		*e = AccountType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AccountType: %v", v)
	}
}
