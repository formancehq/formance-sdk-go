// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type V2BulkResponse struct {
	Data []V2BulkElementResult `json:"data"`
}

func (o *V2BulkResponse) GetData() []V2BulkElementResult {
	if o == nil {
		return []V2BulkElementResult{}
	}
	return o.Data
}
