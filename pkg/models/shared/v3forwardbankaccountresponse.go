// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type V3ForwardBankAccountResponseData struct {
	// Since this call is asynchronous, the response will contain the ID of the task that was created to forward the bank account to the PSP. You can use the task API to check the status of the task and get the resulting bank account ID.
	//
	TaskID string `json:"taskID"`
}

func (o *V3ForwardBankAccountResponseData) GetTaskID() string {
	if o == nil {
		return ""
	}
	return o.TaskID
}

type V3ForwardBankAccountResponse struct {
	Data V3ForwardBankAccountResponseData `json:"data"`
}

func (o *V3ForwardBankAccountResponse) GetData() V3ForwardBankAccountResponseData {
	if o == nil {
		return V3ForwardBankAccountResponseData{}
	}
	return o.Data
}
