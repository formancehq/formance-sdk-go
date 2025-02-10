// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/formancehq/formance-sdk-go/v3/pkg/utils"
	"time"
)

type TaskCurrencyCloudDescriptor struct {
	Name *string `json:"name,omitempty"`
}

func (o *TaskCurrencyCloudDescriptor) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type TaskCurrencyCloudState struct {
}

type TaskCurrencyCloud struct {
	ConnectorID string                      `json:"connectorID"`
	CreatedAt   time.Time                   `json:"createdAt"`
	Descriptor  TaskCurrencyCloudDescriptor `json:"descriptor"`
	Error       *string                     `json:"error,omitempty"`
	ID          string                      `json:"id"`
	State       *TaskCurrencyCloudState     `json:"state,omitempty"`
	Status      TaskStatus                  `json:"status"`
	UpdatedAt   time.Time                   `json:"updatedAt"`
}

func (t TaskCurrencyCloud) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TaskCurrencyCloud) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *TaskCurrencyCloud) GetConnectorID() string {
	if o == nil {
		return ""
	}
	return o.ConnectorID
}

func (o *TaskCurrencyCloud) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *TaskCurrencyCloud) GetDescriptor() TaskCurrencyCloudDescriptor {
	if o == nil {
		return TaskCurrencyCloudDescriptor{}
	}
	return o.Descriptor
}

func (o *TaskCurrencyCloud) GetError() *string {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *TaskCurrencyCloud) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *TaskCurrencyCloud) GetState() *TaskCurrencyCloudState {
	if o == nil {
		return nil
	}
	return o.State
}

func (o *TaskCurrencyCloud) GetStatus() TaskStatus {
	if o == nil {
		return TaskStatus("")
	}
	return o.Status
}

func (o *TaskCurrencyCloud) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}
