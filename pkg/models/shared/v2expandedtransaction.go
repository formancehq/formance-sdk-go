// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/formancehq/formance-sdk-go/v2/pkg/utils"
	"math/big"
	"time"
)

type V2ExpandedTransaction struct {
	ID                *big.Int                       `json:"id"`
	Metadata          map[string]string              `json:"metadata"`
	PostCommitVolumes map[string]map[string]V2Volume `json:"postCommitVolumes,omitempty"`
	Postings          []V2Posting                    `json:"postings"`
	PreCommitVolumes  map[string]map[string]V2Volume `json:"preCommitVolumes,omitempty"`
	Reference         *string                        `json:"reference,omitempty"`
	Reverted          bool                           `json:"reverted"`
	Timestamp         time.Time                      `json:"timestamp"`
}

func (v V2ExpandedTransaction) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(v, "", false)
}

func (v *V2ExpandedTransaction) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &v, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *V2ExpandedTransaction) GetID() *big.Int {
	if o == nil {
		return big.NewInt(0)
	}
	return o.ID
}

func (o *V2ExpandedTransaction) GetMetadata() map[string]string {
	if o == nil {
		return map[string]string{}
	}
	return o.Metadata
}

func (o *V2ExpandedTransaction) GetPostCommitVolumes() map[string]map[string]V2Volume {
	if o == nil {
		return nil
	}
	return o.PostCommitVolumes
}

func (o *V2ExpandedTransaction) GetPostings() []V2Posting {
	if o == nil {
		return []V2Posting{}
	}
	return o.Postings
}

func (o *V2ExpandedTransaction) GetPreCommitVolumes() map[string]map[string]V2Volume {
	if o == nil {
		return nil
	}
	return o.PreCommitVolumes
}

func (o *V2ExpandedTransaction) GetReference() *string {
	if o == nil {
		return nil
	}
	return o.Reference
}

func (o *V2ExpandedTransaction) GetReverted() bool {
	if o == nil {
		return false
	}
	return o.Reverted
}

func (o *V2ExpandedTransaction) GetTimestamp() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.Timestamp
}
