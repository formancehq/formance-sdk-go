// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/formancehq/formance-sdk-go/v2/v2/pkg/utils"
	"time"
)

type V2WorkflowInstance struct {
	CreatedAt    time.Time       `json:"createdAt"`
	Error        *string         `json:"error,omitempty"`
	ID           string          `json:"id"`
	Status       []V2StageStatus `json:"status,omitempty"`
	Terminated   bool            `json:"terminated"`
	TerminatedAt *time.Time      `json:"terminatedAt,omitempty"`
	UpdatedAt    time.Time       `json:"updatedAt"`
	WorkflowID   string          `json:"workflowID"`
}

func (v V2WorkflowInstance) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(v, "", false)
}

func (v *V2WorkflowInstance) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &v, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *V2WorkflowInstance) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *V2WorkflowInstance) GetError() *string {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *V2WorkflowInstance) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *V2WorkflowInstance) GetStatus() []V2StageStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *V2WorkflowInstance) GetTerminated() bool {
	if o == nil {
		return false
	}
	return o.Terminated
}

func (o *V2WorkflowInstance) GetTerminatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.TerminatedAt
}

func (o *V2WorkflowInstance) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}

func (o *V2WorkflowInstance) GetWorkflowID() string {
	if o == nil {
		return ""
	}
	return o.WorkflowID
}
