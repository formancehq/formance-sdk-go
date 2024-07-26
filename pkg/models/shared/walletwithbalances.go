// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package shared

import (
	"github.com/formancehq/formance-sdk-go/v2/pkg/utils"
	"time"
)

type WalletWithBalancesBalances struct {
	Main AssetHolder `json:"main"`
}

func (o *WalletWithBalancesBalances) GetMain() AssetHolder {
	if o == nil {
		return AssetHolder{}
	}
	return o.Main
}

type WalletWithBalances struct {
	Balances  WalletWithBalancesBalances `json:"balances"`
	CreatedAt time.Time                  `json:"createdAt"`
	// The unique ID of the wallet.
	ID     string `json:"id"`
	Ledger string `json:"ledger"`
	// Metadata associated with the wallet.
	Metadata map[string]string `json:"metadata"`
	Name     string            `json:"name"`
}

func (w WalletWithBalances) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(w, "", false)
}

func (w *WalletWithBalances) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &w, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *WalletWithBalances) GetBalances() WalletWithBalancesBalances {
	if o == nil {
		return WalletWithBalancesBalances{}
	}
	return o.Balances
}

func (o *WalletWithBalances) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *WalletWithBalances) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *WalletWithBalances) GetLedger() string {
	if o == nil {
		return ""
	}
	return o.Ledger
}

func (o *WalletWithBalances) GetMetadata() map[string]string {
	if o == nil {
		return map[string]string{}
	}
	return o.Metadata
}

func (o *WalletWithBalances) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}
