// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type AttemptResponse struct {
	Data Attempt `json:"data"`
}

func (o *AttemptResponse) GetData() Attempt {
	if o == nil {
		return Attempt{}
	}
	return o.Data
}
