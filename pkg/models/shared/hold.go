// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type Hold struct {
	Asset       string   `json:"asset"`
	Description string   `json:"description"`
	Destination *Subject `json:"destination,omitempty"`
	// The unique ID of the hold.
	ID string `json:"id"`
	// Metadata associated with the hold.
	Metadata map[string]string `json:"metadata"`
	// The ID of the wallet the hold is associated with.
	WalletID string `json:"walletID"`
}

func (o *Hold) GetAsset() string {
	if o == nil {
		return ""
	}
	return o.Asset
}

func (o *Hold) GetDescription() string {
	if o == nil {
		return ""
	}
	return o.Description
}

func (o *Hold) GetDestination() *Subject {
	if o == nil {
		return nil
	}
	return o.Destination
}

func (o *Hold) GetDestinationAccount() *LedgerAccountSubject {
	if v := o.GetDestination(); v != nil {
		return v.LedgerAccountSubject
	}
	return nil
}

func (o *Hold) GetDestinationWallet() *WalletSubject {
	if v := o.GetDestination(); v != nil {
		return v.WalletSubject
	}
	return nil
}

func (o *Hold) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *Hold) GetMetadata() map[string]string {
	if o == nil {
		return map[string]string{}
	}
	return o.Metadata
}

func (o *Hold) GetWalletID() string {
	if o == nil {
		return ""
	}
	return o.WalletID
}
