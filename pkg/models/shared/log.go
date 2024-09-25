// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/formancehq/formance-sdk-go/v3/pkg/utils"
	"time"
)

type Type string

const (
	TypeNewTransaction Type = "NEW_TRANSACTION"
	TypeSetMetadata    Type = "SET_METADATA"
)

func (e Type) ToPointer() *Type {
	return &e
}
func (e *Type) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "NEW_TRANSACTION":
		fallthrough
	case "SET_METADATA":
		*e = Type(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Type: %v", v)
	}
}

type Log struct {
	Data map[string]any `json:"data"`
	Date time.Time      `json:"date"`
	Hash string         `json:"hash"`
	ID   int64          `json:"id"`
	Type Type           `json:"type"`
}

func (l Log) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *Log) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Log) GetData() map[string]any {
	if o == nil {
		return map[string]any{}
	}
	return o.Data
}

func (o *Log) GetDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.Date
}

func (o *Log) GetHash() string {
	if o == nil {
		return ""
	}
	return o.Hash
}

func (o *Log) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

func (o *Log) GetType() Type {
	if o == nil {
		return Type("")
	}
	return o.Type
}
