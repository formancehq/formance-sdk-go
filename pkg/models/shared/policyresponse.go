// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// PolicyResponse - OK
type PolicyResponse struct {
	Data Policy `json:"data"`
}

func (o *PolicyResponse) GetData() Policy {
	if o == nil {
		return Policy{}
	}
	return o.Data
}
