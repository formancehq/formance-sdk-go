// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/formancehq/formance-sdk-go/v2/pkg/utils"
	"math/big"
)

type V2Posting struct {
	Amount      *big.Int `json:"amount"`
	Asset       string   `json:"asset"`
	Destination string   `json:"destination"`
	Source      string   `json:"source"`
}

func (v V2Posting) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(v, "", false)
}

func (v *V2Posting) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &v, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *V2Posting) GetAmount() *big.Int {
	if o == nil {
		return big.NewInt(0)
	}
	return o.Amount
}

func (o *V2Posting) GetAsset() string {
	if o == nil {
		return ""
	}
	return o.Asset
}

func (o *V2Posting) GetDestination() string {
	if o == nil {
		return ""
	}
	return o.Destination
}

func (o *V2Posting) GetSource() string {
	if o == nil {
		return ""
	}
	return o.Source
}
