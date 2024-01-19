// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/formancehq/formance-sdk-go/v2/pkg/utils"
	"math/big"
)

type V2Volume struct {
	Balance *big.Int `json:"balance,omitempty"`
	Input   *big.Int `json:"input"`
	Output  *big.Int `json:"output"`
}

func (v V2Volume) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(v, "", false)
}

func (v *V2Volume) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &v, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *V2Volume) GetBalance() *big.Int {
	if o == nil {
		return nil
	}
	return o.Balance
}

func (o *V2Volume) GetInput() *big.Int {
	if o == nil {
		return big.NewInt(0)
	}
	return o.Input
}

func (o *V2Volume) GetOutput() *big.Int {
	if o == nil {
		return big.NewInt(0)
	}
	return o.Output
}
