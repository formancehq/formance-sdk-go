// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/formancehq/formance-sdk-go/v3/pkg/utils"
	"math/big"
)

type PoolBalance struct {
	Amount *big.Int `json:"amount"`
	Asset  string   `json:"asset"`
}

func (p PoolBalance) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PoolBalance) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *PoolBalance) GetAmount() *big.Int {
	if o == nil {
		return big.NewInt(0)
	}
	return o.Amount
}

func (o *PoolBalance) GetAsset() string {
	if o == nil {
		return ""
	}
	return o.Asset
}
