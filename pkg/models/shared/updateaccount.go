// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type UpdateAccount struct {
	ID       string            `json:"id"`
	Ledger   string            `json:"ledger"`
	Metadata map[string]string `json:"metadata"`
}

func (o *UpdateAccount) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateAccount) GetLedger() string {
	if o == nil {
		return ""
	}
	return o.Ledger
}

func (o *UpdateAccount) GetMetadata() map[string]string {
	if o == nil {
		return map[string]string{}
	}
	return o.Metadata
}
