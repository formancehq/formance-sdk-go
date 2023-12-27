// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/formancehq/formance-sdk-go/pkg/utils"
	"math/big"
	"time"
)

type OrchestrationTransaction struct {
	ID        *big.Int          `json:"id"`
	Metadata  map[string]string `json:"metadata"`
	Postings  []Posting         `json:"postings"`
	Reference *string           `json:"reference,omitempty"`
	Reverted  bool              `json:"reverted"`
	Timestamp time.Time         `json:"timestamp"`
}

func (o OrchestrationTransaction) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(o, "", false)
}

func (o *OrchestrationTransaction) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &o, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *OrchestrationTransaction) GetID() *big.Int {
	if o == nil {
		return big.NewInt(0)
	}
	return o.ID
}

func (o *OrchestrationTransaction) GetMetadata() map[string]string {
	if o == nil {
		return map[string]string{}
	}
	return o.Metadata
}

func (o *OrchestrationTransaction) GetPostings() []Posting {
	if o == nil {
		return []Posting{}
	}
	return o.Postings
}

func (o *OrchestrationTransaction) GetReference() *string {
	if o == nil {
		return nil
	}
	return o.Reference
}

func (o *OrchestrationTransaction) GetReverted() bool {
	if o == nil {
		return false
	}
	return o.Reverted
}

func (o *OrchestrationTransaction) GetTimestamp() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.Timestamp
}
