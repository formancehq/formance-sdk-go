// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type GetWorkflowInstanceHistoryStageResponse struct {
	Data []WorkflowInstanceHistoryStage `json:"data"`
}

func (o *GetWorkflowInstanceHistoryStageResponse) GetData() []WorkflowInstanceHistoryStage {
	if o == nil {
		return []WorkflowInstanceHistoryStage{}
	}
	return o.Data
}
