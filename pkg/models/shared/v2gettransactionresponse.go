// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type V2GetTransactionResponse struct {
	Data V2Transaction `json:"data"`
}

func (o *V2GetTransactionResponse) GetData() V2Transaction {
	if o == nil {
		return V2Transaction{}
	}
	return o.Data
}
