// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package shared

type PaymentResponse struct {
	Data Payment `json:"data"`
}

func (o *PaymentResponse) GetData() Payment {
	if o == nil {
		return Payment{}
	}
	return o.Data
}
