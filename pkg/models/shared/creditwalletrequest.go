// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/formancehq/formance-sdk-go/v2/pkg/utils"
	"time"
)

type CreditWalletRequest struct {
	Amount Monetary `json:"amount"`
	// The balance to credit
	Balance *string `json:"balance,omitempty"`
	// Metadata associated with the wallet.
	Metadata  map[string]string `json:"metadata,omitempty"`
	Reference *string           `json:"reference,omitempty"`
	Sources   []Subject         `json:"sources,omitempty"`
	Timestamp *time.Time        `json:"timestamp,omitempty"`
}

func (c CreditWalletRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreditWalletRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreditWalletRequest) GetAmount() Monetary {
	if o == nil {
		return Monetary{}
	}
	return o.Amount
}

func (o *CreditWalletRequest) GetBalance() *string {
	if o == nil {
		return nil
	}
	return o.Balance
}

func (o *CreditWalletRequest) GetMetadata() map[string]string {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *CreditWalletRequest) GetReference() *string {
	if o == nil {
		return nil
	}
	return o.Reference
}

func (o *CreditWalletRequest) GetSources() []Subject {
	if o == nil {
		return nil
	}
	return o.Sources
}

func (o *CreditWalletRequest) GetTimestamp() *time.Time {
	if o == nil {
		return nil
	}
	return o.Timestamp
}
