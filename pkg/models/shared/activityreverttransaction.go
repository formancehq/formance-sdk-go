// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ActivityRevertTransaction struct {
	ID     string `json:"id"`
	Ledger string `json:"ledger"`
}

func (o *ActivityRevertTransaction) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ActivityRevertTransaction) GetLedger() string {
	if o == nil {
		return ""
	}
	return o.Ledger
}
