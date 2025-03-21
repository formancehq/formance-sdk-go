// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type V3PaymentInitiationTypeEnum string

const (
	V3PaymentInitiationTypeEnumUnknown  V3PaymentInitiationTypeEnum = "UNKNOWN"
	V3PaymentInitiationTypeEnumTransfer V3PaymentInitiationTypeEnum = "TRANSFER"
	V3PaymentInitiationTypeEnumPayout   V3PaymentInitiationTypeEnum = "PAYOUT"
)

func (e V3PaymentInitiationTypeEnum) ToPointer() *V3PaymentInitiationTypeEnum {
	return &e
}
func (e *V3PaymentInitiationTypeEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "UNKNOWN":
		fallthrough
	case "TRANSFER":
		fallthrough
	case "PAYOUT":
		*e = V3PaymentInitiationTypeEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for V3PaymentInitiationTypeEnum: %v", v)
	}
}
