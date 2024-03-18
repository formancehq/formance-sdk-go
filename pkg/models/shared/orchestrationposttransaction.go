// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/formancehq/formance-sdk-go/v2/v2/pkg/utils"
	"time"
)

type OrchestrationPostTransactionScript struct {
	Plain string                 `json:"plain"`
	Vars  map[string]interface{} `json:"vars,omitempty"`
}

func (o *OrchestrationPostTransactionScript) GetPlain() string {
	if o == nil {
		return ""
	}
	return o.Plain
}

func (o *OrchestrationPostTransactionScript) GetVars() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.Vars
}

type OrchestrationPostTransaction struct {
	Metadata  map[string]string                   `json:"metadata"`
	Postings  []Posting                           `json:"postings,omitempty"`
	Reference *string                             `json:"reference,omitempty"`
	Script    *OrchestrationPostTransactionScript `json:"script,omitempty"`
	Timestamp *time.Time                          `json:"timestamp,omitempty"`
}

func (o OrchestrationPostTransaction) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(o, "", false)
}

func (o *OrchestrationPostTransaction) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &o, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *OrchestrationPostTransaction) GetMetadata() map[string]string {
	if o == nil {
		return map[string]string{}
	}
	return o.Metadata
}

func (o *OrchestrationPostTransaction) GetPostings() []Posting {
	if o == nil {
		return nil
	}
	return o.Postings
}

func (o *OrchestrationPostTransaction) GetReference() *string {
	if o == nil {
		return nil
	}
	return o.Reference
}

func (o *OrchestrationPostTransaction) GetScript() *OrchestrationPostTransactionScript {
	if o == nil {
		return nil
	}
	return o.Script
}

func (o *OrchestrationPostTransaction) GetTimestamp() *time.Time {
	if o == nil {
		return nil
	}
	return o.Timestamp
}
