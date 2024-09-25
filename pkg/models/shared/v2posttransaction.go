// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/formancehq/formance-sdk-go/v3/pkg/utils"
	"time"
)

type V2PostTransactionScript struct {
	Plain string         `json:"plain"`
	Vars  map[string]any `json:"vars,omitempty"`
}

func (o *V2PostTransactionScript) GetPlain() string {
	if o == nil {
		return ""
	}
	return o.Plain
}

func (o *V2PostTransactionScript) GetVars() map[string]any {
	if o == nil {
		return nil
	}
	return o.Vars
}

type V2PostTransaction struct {
	Metadata  map[string]string        `json:"metadata"`
	Postings  []V2Posting              `json:"postings,omitempty"`
	Reference *string                  `json:"reference,omitempty"`
	Script    *V2PostTransactionScript `json:"script,omitempty"`
	Timestamp *time.Time               `json:"timestamp,omitempty"`
}

func (v V2PostTransaction) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(v, "", false)
}

func (v *V2PostTransaction) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &v, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *V2PostTransaction) GetMetadata() map[string]string {
	if o == nil {
		return map[string]string{}
	}
	return o.Metadata
}

func (o *V2PostTransaction) GetPostings() []V2Posting {
	if o == nil {
		return nil
	}
	return o.Postings
}

func (o *V2PostTransaction) GetReference() *string {
	if o == nil {
		return nil
	}
	return o.Reference
}

func (o *V2PostTransaction) GetScript() *V2PostTransactionScript {
	if o == nil {
		return nil
	}
	return o.Script
}

func (o *V2PostTransaction) GetTimestamp() *time.Time {
	if o == nil {
		return nil
	}
	return o.Timestamp
}
