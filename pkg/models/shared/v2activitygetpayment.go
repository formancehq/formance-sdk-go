// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package shared

type V2ActivityGetPayment struct {
	ID string `json:"id"`
}

func (o *V2ActivityGetPayment) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}
