// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package shared

type GetWorkflowResponse struct {
	Data Workflow `json:"data"`
}

func (o *GetWorkflowResponse) GetData() Workflow {
	if o == nil {
		return Workflow{}
	}
	return o.Data
}
