// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/v2/pkg/models/shared"
	"net/http"
)

type V2GetWorkflowRequest struct {
	// The flow id
	FlowID string `pathParam:"style=simple,explode=false,name=flowId"`
}

func (o *V2GetWorkflowRequest) GetFlowID() string {
	if o == nil {
		return ""
	}
	return o.FlowID
}

type V2GetWorkflowResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The workflow
	V2GetWorkflowResponse *shared.V2GetWorkflowResponse
}

func (o *V2GetWorkflowResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *V2GetWorkflowResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *V2GetWorkflowResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *V2GetWorkflowResponse) GetV2GetWorkflowResponse() *shared.V2GetWorkflowResponse {
	if o == nil {
		return nil
	}
	return o.V2GetWorkflowResponse
}
