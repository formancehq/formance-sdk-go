// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type TransferResponse struct {
	ID *string `json:"id,omitempty"`
}

func (o *TransferResponse) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}
