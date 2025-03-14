// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/formancehq/formance-sdk-go/v3/pkg/utils"
	"math/big"
)

type WalletSummary struct {
	AvailableFunds map[string]*big.Int `json:"availableFunds"`
	Balances       []BalanceWithAssets `json:"balances"`
	ExpirableFunds map[string]*big.Int `json:"expirableFunds"`
	ExpiredFunds   map[string]*big.Int `json:"expiredFunds"`
	HoldFunds      map[string]*big.Int `json:"holdFunds"`
}

func (w WalletSummary) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(w, "", false)
}

func (w *WalletSummary) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &w, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *WalletSummary) GetAvailableFunds() map[string]*big.Int {
	if o == nil {
		return map[string]*big.Int{}
	}
	return o.AvailableFunds
}

func (o *WalletSummary) GetBalances() []BalanceWithAssets {
	if o == nil {
		return []BalanceWithAssets{}
	}
	return o.Balances
}

func (o *WalletSummary) GetExpirableFunds() map[string]*big.Int {
	if o == nil {
		return map[string]*big.Int{}
	}
	return o.ExpirableFunds
}

func (o *WalletSummary) GetExpiredFunds() map[string]*big.Int {
	if o == nil {
		return map[string]*big.Int{}
	}
	return o.ExpiredFunds
}

func (o *WalletSummary) GetHoldFunds() map[string]*big.Int {
	if o == nil {
		return map[string]*big.Int{}
	}
	return o.HoldFunds
}
