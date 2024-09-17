// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/shared"
	"net/http"
)

type TestTriggerRequest struct {
	RequestBody map[string]any `request:"mediaType=application/json"`
	// The trigger id
	TriggerID string `pathParam:"style=simple,explode=false,name=triggerID"`
}

func (o *TestTriggerRequest) GetRequestBody() map[string]any {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *TestTriggerRequest) GetTriggerID() string {
	if o == nil {
		return ""
	}
	return o.TriggerID
}

type TestTriggerResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Test a trigger
	V2TestTriggerResponse *shared.V2TestTriggerResponse
}

func (o *TestTriggerResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *TestTriggerResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *TestTriggerResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *TestTriggerResponse) GetV2TestTriggerResponse() *shared.V2TestTriggerResponse {
	if o == nil {
		return nil
	}
	return o.V2TestTriggerResponse
}
