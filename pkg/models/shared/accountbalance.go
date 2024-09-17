// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/formancehq/formance-sdk-go/v3/pkg/utils"
	"math/big"
	"time"
)

type AccountBalance struct {
	AccountID string    `json:"accountId"`
	Asset     string    `json:"asset"`
	Balance   *big.Int  `json:"balance"`
	CreatedAt time.Time `json:"createdAt"`
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	Currency      string    `json:"currency"`
	LastUpdatedAt time.Time `json:"lastUpdatedAt"`
}

func (a AccountBalance) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AccountBalance) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AccountBalance) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *AccountBalance) GetAsset() string {
	if o == nil {
		return ""
	}
	return o.Asset
}

func (o *AccountBalance) GetBalance() *big.Int {
	if o == nil {
		return big.NewInt(0)
	}
	return o.Balance
}

func (o *AccountBalance) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *AccountBalance) GetCurrency() string {
	if o == nil {
		return ""
	}
	return o.Currency
}

func (o *AccountBalance) GetLastUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.LastUpdatedAt
}
