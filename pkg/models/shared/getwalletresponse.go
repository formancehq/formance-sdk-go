// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type GetWalletResponse struct {
	Data WalletWithBalances `json:"data"`
}

func (o *GetWalletResponse) GetData() WalletWithBalances {
	if o == nil {
		return WalletWithBalances{}
	}
	return o.Data
}
