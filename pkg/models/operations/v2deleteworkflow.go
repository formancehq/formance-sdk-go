// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"net/http"
)

type V2DeleteWorkflowRequest struct {
	// The flow id
	FlowID string `pathParam:"style=simple,explode=false,name=flowId"`
}

func (o *V2DeleteWorkflowRequest) GetFlowID() string {
	if o == nil {
		return ""
	}
	return o.FlowID
}

type V2DeleteWorkflowResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *V2DeleteWorkflowResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *V2DeleteWorkflowResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *V2DeleteWorkflowResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
