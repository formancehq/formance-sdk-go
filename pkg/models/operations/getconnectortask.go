// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/v2/pkg/models/shared"
	"net/http"
)

type GetConnectorTaskRequest struct {
	// The name of the connector.
	Connector shared.Connector `pathParam:"style=simple,explode=false,name=connector"`
	// The task ID.
	TaskID string `pathParam:"style=simple,explode=false,name=taskId"`
}

func (o *GetConnectorTaskRequest) GetConnector() shared.Connector {
	if o == nil {
		return shared.Connector("")
	}
	return o.Connector
}

func (o *GetConnectorTaskRequest) GetTaskID() string {
	if o == nil {
		return ""
	}
	return o.TaskID
}

type GetConnectorTaskResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	TaskResponse *shared.TaskResponse
}

func (o *GetConnectorTaskResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetConnectorTaskResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetConnectorTaskResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetConnectorTaskResponse) GetTaskResponse() *shared.TaskResponse {
	if o == nil {
		return nil
	}
	return o.TaskResponse
}
