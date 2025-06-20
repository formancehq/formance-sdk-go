// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/formancehq/formance-sdk-go/v3/pkg/utils"
	"math/big"
)

type ExpandedDebitHold struct {
	Asset       string   `json:"asset"`
	Description string   `json:"description"`
	Destination *Subject `json:"destination,omitempty"`
	// The unique ID of the hold.
	ID string `json:"id"`
	// Metadata associated with the hold.
	Metadata map[string]string `json:"metadata"`
	// Original amount on hold
	OriginalAmount *big.Int `json:"originalAmount"`
	// Remaining amount on hold
	Remaining *big.Int `json:"remaining"`
	// The ID of the wallet the hold is associated with.
	WalletID string `json:"walletID"`
}

func (e ExpandedDebitHold) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(e, "", false)
}

func (e *ExpandedDebitHold) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &e, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ExpandedDebitHold) GetAsset() string {
	if o == nil {
		return ""
	}
	return o.Asset
}

func (o *ExpandedDebitHold) GetDescription() string {
	if o == nil {
		return ""
	}
	return o.Description
}

func (o *ExpandedDebitHold) GetDestination() *Subject {
	if o == nil {
		return nil
	}
	return o.Destination
}

func (o *ExpandedDebitHold) GetDestinationAccount() *LedgerAccountSubject {
	if v := o.GetDestination(); v != nil {
		return v.LedgerAccountSubject
	}
	return nil
}

func (o *ExpandedDebitHold) GetDestinationWallet() *WalletSubject {
	if v := o.GetDestination(); v != nil {
		return v.WalletSubject
	}
	return nil
}

func (o *ExpandedDebitHold) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ExpandedDebitHold) GetMetadata() map[string]string {
	if o == nil {
		return map[string]string{}
	}
	return o.Metadata
}

func (o *ExpandedDebitHold) GetOriginalAmount() *big.Int {
	if o == nil {
		return big.NewInt(0)
	}
	return o.OriginalAmount
}

func (o *ExpandedDebitHold) GetRemaining() *big.Int {
	if o == nil {
		return big.NewInt(0)
	}
	return o.Remaining
}

func (o *ExpandedDebitHold) GetWalletID() string {
	if o == nil {
		return ""
	}
	return o.WalletID
}
