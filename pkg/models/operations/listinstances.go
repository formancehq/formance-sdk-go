// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/v2/v2/pkg/models/shared"
	"net/http"
)

type ListInstancesRequest struct {
	// Filter running instances
	Running *bool `queryParam:"style=form,explode=true,name=running"`
	// A workflow id
	WorkflowID *string `queryParam:"style=form,explode=true,name=workflowID"`
}

func (o *ListInstancesRequest) GetRunning() *bool {
	if o == nil {
		return nil
	}
	return o.Running
}

func (o *ListInstancesRequest) GetWorkflowID() *string {
	if o == nil {
		return nil
	}
	return o.WorkflowID
}

type ListInstancesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// General error
	Error *shared.Error
	// List of workflow instances
	ListRunsResponse *shared.ListRunsResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ListInstancesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListInstancesResponse) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *ListInstancesResponse) GetListRunsResponse() *shared.ListRunsResponse {
	if o == nil {
		return nil
	}
	return o.ListRunsResponse
}

func (o *ListInstancesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListInstancesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
