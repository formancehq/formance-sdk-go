// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type CreateWorkflowResponse struct {
	Data Workflow `json:"data"`
}

func (o *CreateWorkflowResponse) GetData() Workflow {
	if o == nil {
		return Workflow{}
	}
	return o.Data
}
