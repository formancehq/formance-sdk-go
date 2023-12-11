// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/formancehq/formance-sdk-go/pkg/utils"
	"time"
)

type TaskMangoPayDescriptor struct {
	Key    *string `json:"key,omitempty"`
	Name   *string `json:"name,omitempty"`
	UserID *string `json:"userID,omitempty"`
}

func (o *TaskMangoPayDescriptor) GetKey() *string {
	if o == nil {
		return nil
	}
	return o.Key
}

func (o *TaskMangoPayDescriptor) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *TaskMangoPayDescriptor) GetUserID() *string {
	if o == nil {
		return nil
	}
	return o.UserID
}

type TaskMangoPayState struct {
}

type TaskMangoPay struct {
	ConnectorID string                 `json:"connectorId"`
	CreatedAt   time.Time              `json:"createdAt"`
	Descriptor  TaskMangoPayDescriptor `json:"descriptor"`
	Error       *string                `json:"error,omitempty"`
	ID          string                 `json:"id"`
	State       TaskMangoPayState      `json:"state"`
	Status      PaymentStatus          `json:"status"`
	UpdatedAt   time.Time              `json:"updatedAt"`
}

func (t TaskMangoPay) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TaskMangoPay) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *TaskMangoPay) GetConnectorID() string {
	if o == nil {
		return ""
	}
	return o.ConnectorID
}

func (o *TaskMangoPay) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *TaskMangoPay) GetDescriptor() TaskMangoPayDescriptor {
	if o == nil {
		return TaskMangoPayDescriptor{}
	}
	return o.Descriptor
}

func (o *TaskMangoPay) GetError() *string {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *TaskMangoPay) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *TaskMangoPay) GetState() TaskMangoPayState {
	if o == nil {
		return TaskMangoPayState{}
	}
	return o.State
}

func (o *TaskMangoPay) GetStatus() PaymentStatus {
	if o == nil {
		return PaymentStatus("")
	}
	return o.Status
}

func (o *TaskMangoPay) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}
