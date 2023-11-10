// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/formancehq/formance-sdk-go/pkg/utils"
	"math/big"
)

type WalletsPosting struct {
	Amount      *big.Int `json:"amount"`
	Asset       string   `json:"asset"`
	Destination string   `json:"destination"`
	Source      string   `json:"source"`
}

func (w WalletsPosting) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(w, "", false)
}

func (w *WalletsPosting) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &w, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *WalletsPosting) GetAmount() *big.Int {
	if o == nil {
		return big.NewInt(0)
	}
	return o.Amount
}

func (o *WalletsPosting) GetAsset() string {
	if o == nil {
		return ""
	}
	return o.Asset
}

func (o *WalletsPosting) GetDestination() string {
	if o == nil {
		return ""
	}
	return o.Destination
}

func (o *WalletsPosting) GetSource() string {
	if o == nil {
		return ""
	}
	return o.Source
}
