// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type StatsResponse struct {
	Data Stats `json:"data"`
}

func (o *StatsResponse) GetData() Stats {
	if o == nil {
		return Stats{}
	}
	return o.Data
}
