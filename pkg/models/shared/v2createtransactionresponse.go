// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type V2CreateTransactionResponse struct {
	Data V2Transaction `json:"data"`
}

func (o *V2CreateTransactionResponse) GetData() V2Transaction {
	if o == nil {
		return V2Transaction{}
	}
	return o.Data
}
