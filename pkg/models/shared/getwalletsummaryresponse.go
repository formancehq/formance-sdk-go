// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"math/big"
)

// GetWalletSummaryResponse - Wallet summary
type GetWalletSummaryResponse struct {
	AvailableFunds map[string]*big.Int `json:"availableFunds"`
	Balances       []BalanceWithAssets `json:"balances"`
	ExpirableFunds map[string]*big.Int `json:"expirableFunds"`
	ExpiredFunds   map[string]*big.Int `json:"expiredFunds"`
	HoldFunds      map[string]*big.Int `json:"holdFunds"`
}
