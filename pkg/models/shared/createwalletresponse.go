// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type CreateWalletResponse struct {
	Data Wallet `json:"data"`
}

func (o *CreateWalletResponse) GetData() Wallet {
	if o == nil {
		return Wallet{}
	}
	return o.Data
}
