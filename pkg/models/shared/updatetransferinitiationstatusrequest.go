// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type Status string

const (
	StatusRejected  Status = "REJECTED"
	StatusValidated Status = "VALIDATED"
)

func (e Status) ToPointer() *Status {
	return &e
}
func (e *Status) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "REJECTED":
		fallthrough
	case "VALIDATED":
		*e = Status(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Status: %v", v)
	}
}

type UpdateTransferInitiationStatusRequest struct {
	Status Status `json:"status"`
}

func (o *UpdateTransferInitiationStatusRequest) GetStatus() Status {
	if o == nil {
		return Status("")
	}
	return o.Status
}
