// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/formancehq/formance-sdk-go/v3/pkg/utils"
	"time"
)

type TaskDummyPayDescriptor struct {
	FileName *string `json:"fileName,omitempty"`
	Key      *string `json:"key,omitempty"`
	Name     *string `json:"name,omitempty"`
}

func (o *TaskDummyPayDescriptor) GetFileName() *string {
	if o == nil {
		return nil
	}
	return o.FileName
}

func (o *TaskDummyPayDescriptor) GetKey() *string {
	if o == nil {
		return nil
	}
	return o.Key
}

func (o *TaskDummyPayDescriptor) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type TaskDummyPayState struct {
}

type TaskDummyPay struct {
	ConnectorID string                 `json:"connectorID"`
	CreatedAt   time.Time              `json:"createdAt"`
	Descriptor  TaskDummyPayDescriptor `json:"descriptor"`
	Error       *string                `json:"error,omitempty"`
	ID          string                 `json:"id"`
	State       *TaskDummyPayState     `json:"state,omitempty"`
	Status      TaskStatus             `json:"status"`
	UpdatedAt   time.Time              `json:"updatedAt"`
}

func (t TaskDummyPay) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TaskDummyPay) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *TaskDummyPay) GetConnectorID() string {
	if o == nil {
		return ""
	}
	return o.ConnectorID
}

func (o *TaskDummyPay) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *TaskDummyPay) GetDescriptor() TaskDummyPayDescriptor {
	if o == nil {
		return TaskDummyPayDescriptor{}
	}
	return o.Descriptor
}

func (o *TaskDummyPay) GetError() *string {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *TaskDummyPay) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *TaskDummyPay) GetState() *TaskDummyPayState {
	if o == nil {
		return nil
	}
	return o.State
}

func (o *TaskDummyPay) GetStatus() TaskStatus {
	if o == nil {
		return TaskStatus("")
	}
	return o.Status
}

func (o *TaskDummyPay) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}
