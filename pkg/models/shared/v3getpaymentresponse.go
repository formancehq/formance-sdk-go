// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type V3GetPaymentResponse struct {
	Data V3Payment `json:"data"`
}

func (o *V3GetPaymentResponse) GetData() V3Payment {
	if o == nil {
		return V3Payment{}
	}
	return o.Data
}
