// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type ForwardBankAccountRequest struct {
	ConnectorID string `json:"connectorID"`
}

func (o *ForwardBankAccountRequest) GetConnectorID() string {
	if o == nil {
		return ""
	}
	return o.ConnectorID
}
