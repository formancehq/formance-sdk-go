// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type V2CreditWalletRequest struct {
	Amount V2Monetary `json:"amount"`
	// The balance to credit
	Balance *string `json:"balance,omitempty"`
	// Metadata associated with the wallet.
	Metadata  map[string]string `json:"metadata"`
	Reference *string           `json:"reference,omitempty"`
	Sources   []V2Subject       `json:"sources"`
}

func (o *V2CreditWalletRequest) GetAmount() V2Monetary {
	if o == nil {
		return V2Monetary{}
	}
	return o.Amount
}

func (o *V2CreditWalletRequest) GetBalance() *string {
	if o == nil {
		return nil
	}
	return o.Balance
}

func (o *V2CreditWalletRequest) GetMetadata() map[string]string {
	if o == nil {
		return map[string]string{}
	}
	return o.Metadata
}

func (o *V2CreditWalletRequest) GetReference() *string {
	if o == nil {
		return nil
	}
	return o.Reference
}

func (o *V2CreditWalletRequest) GetSources() []V2Subject {
	if o == nil {
		return []V2Subject{}
	}
	return o.Sources
}
