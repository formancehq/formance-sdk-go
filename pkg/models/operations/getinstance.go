// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/v2/pkg/models/shared"
	"net/http"
)

type GetInstanceRequest struct {
	// The instance id
	InstanceID string `pathParam:"style=simple,explode=false,name=instanceID"`
}

func (o *GetInstanceRequest) GetInstanceID() string {
	if o == nil {
		return ""
	}
	return o.InstanceID
}

type GetInstanceResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// The workflow instance
	GetWorkflowInstanceResponse *shared.GetWorkflowInstanceResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetInstanceResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetInstanceResponse) GetGetWorkflowInstanceResponse() *shared.GetWorkflowInstanceResponse {
	if o == nil {
		return nil
	}
	return o.GetWorkflowInstanceResponse
}

func (o *GetInstanceResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetInstanceResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
