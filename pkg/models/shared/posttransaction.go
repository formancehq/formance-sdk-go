// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/formancehq/formance-sdk-go/v2/v2/pkg/utils"
	"time"
)

type PostTransactionScript struct {
	Plain string                 `json:"plain"`
	Vars  map[string]interface{} `json:"vars,omitempty"`
}

func (o *PostTransactionScript) GetPlain() string {
	if o == nil {
		return ""
	}
	return o.Plain
}

func (o *PostTransactionScript) GetVars() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.Vars
}

type PostTransaction struct {
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
	Postings  []Posting              `json:"postings,omitempty"`
	Reference *string                `json:"reference,omitempty"`
	Script    *PostTransactionScript `json:"script,omitempty"`
	Timestamp *time.Time             `json:"timestamp,omitempty"`
}

func (p PostTransaction) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PostTransaction) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *PostTransaction) GetMetadata() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *PostTransaction) GetPostings() []Posting {
	if o == nil {
		return nil
	}
	return o.Postings
}

func (o *PostTransaction) GetReference() *string {
	if o == nil {
		return nil
	}
	return o.Reference
}

func (o *PostTransaction) GetScript() *PostTransactionScript {
	if o == nil {
		return nil
	}
	return o.Script
}

func (o *PostTransaction) GetTimestamp() *time.Time {
	if o == nil {
		return nil
	}
	return o.Timestamp
}
