// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type V2GetLedgerResponse struct {
	Data V2Ledger `json:"data"`
}

func (o *V2GetLedgerResponse) GetData() V2Ledger {
	if o == nil {
		return V2Ledger{}
	}
	return o.Data
}
