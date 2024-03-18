// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/formancehq/formance-sdk-go/v2/v2/pkg/utils"
	"math/big"
)

type WalletsVolume struct {
	Balance *big.Int `json:"balance"`
	Input   *big.Int `json:"input"`
	Output  *big.Int `json:"output"`
}

func (w WalletsVolume) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(w, "", false)
}

func (w *WalletsVolume) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &w, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *WalletsVolume) GetBalance() *big.Int {
	if o == nil {
		return big.NewInt(0)
	}
	return o.Balance
}

func (o *WalletsVolume) GetInput() *big.Int {
	if o == nil {
		return big.NewInt(0)
	}
	return o.Input
}

func (o *WalletsVolume) GetOutput() *big.Int {
	if o == nil {
		return big.NewInt(0)
	}
	return o.Output
}
