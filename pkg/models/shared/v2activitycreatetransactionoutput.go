// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type V2ActivityCreateTransactionOutput struct {
	Data []OrchestrationV2Transaction `json:"data"`
}

func (o *V2ActivityCreateTransactionOutput) GetData() []OrchestrationV2Transaction {
	if o == nil {
		return []OrchestrationV2Transaction{}
	}
	return o.Data
}
