// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package shared

type V2CreateTriggerResponse struct {
	Data V2Trigger `json:"data"`
}

func (o *V2CreateTriggerResponse) GetData() V2Trigger {
	if o == nil {
		return V2Trigger{}
	}
	return o.Data
}
