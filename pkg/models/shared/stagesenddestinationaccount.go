// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type StageSendDestinationAccount struct {
	ID     string  `json:"id"`
	Ledger *string `json:"ledger,omitempty"`
}

func (o *StageSendDestinationAccount) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *StageSendDestinationAccount) GetLedger() *string {
	if o == nil {
		return nil
	}
	return o.Ledger
}
