// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package shared

type GetHoldResponse struct {
	Data ExpandedDebitHold `json:"data"`
}

func (o *GetHoldResponse) GetData() ExpandedDebitHold {
	if o == nil {
		return ExpandedDebitHold{}
	}
	return o.Data
}
