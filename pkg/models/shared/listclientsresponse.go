// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package shared

type ListClientsResponse struct {
	Data []Client `json:"data,omitempty"`
}

func (o *ListClientsResponse) GetData() []Client {
	if o == nil {
		return nil
	}
	return o.Data
}
