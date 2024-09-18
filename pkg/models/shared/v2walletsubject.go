// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type V2WalletSubject struct {
	Balance    *string `json:"balance,omitempty"`
	Identifier string  `json:"identifier"`
	Type       string  `json:"type"`
}

func (o *V2WalletSubject) GetBalance() *string {
	if o == nil {
		return nil
	}
	return o.Balance
}

func (o *V2WalletSubject) GetIdentifier() string {
	if o == nil {
		return ""
	}
	return o.Identifier
}

func (o *V2WalletSubject) GetType() string {
	if o == nil {
		return ""
	}
	return o.Type
}
