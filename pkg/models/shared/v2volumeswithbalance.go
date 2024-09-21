// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/formancehq/formance-sdk-go/v2/pkg/utils"
	"math/big"
)

type V2VolumesWithBalance struct {
	Account string   `json:"account"`
	Asset   string   `json:"asset"`
	Balance *big.Int `json:"balance"`
	Input   *big.Int `json:"input"`
	Output  *big.Int `json:"output"`
}

func (v V2VolumesWithBalance) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(v, "", false)
}

func (v *V2VolumesWithBalance) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &v, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *V2VolumesWithBalance) GetAccount() string {
	if o == nil {
		return ""
	}
	return o.Account
}

func (o *V2VolumesWithBalance) GetAsset() string {
	if o == nil {
		return ""
	}
	return o.Asset
}

func (o *V2VolumesWithBalance) GetBalance() *big.Int {
	if o == nil {
		return big.NewInt(0)
	}
	return o.Balance
}

func (o *V2VolumesWithBalance) GetInput() *big.Int {
	if o == nil {
		return big.NewInt(0)
	}
	return o.Input
}

func (o *V2VolumesWithBalance) GetOutput() *big.Int {
	if o == nil {
		return big.NewInt(0)
	}
	return o.Output
}
