// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type ListWorkflowsResponse struct {
	Data []Workflow `json:"data"`
}

func (o *ListWorkflowsResponse) GetData() []Workflow {
	if o == nil {
		return []Workflow{}
	}
	return o.Data
}
