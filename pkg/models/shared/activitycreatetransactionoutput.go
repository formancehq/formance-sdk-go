// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ActivityCreateTransactionOutput struct {
	Data []Transaction `json:"data"`
}

func (o *ActivityCreateTransactionOutput) GetData() []Transaction {
	if o == nil {
		return []Transaction{}
	}
	return o.Data
}
