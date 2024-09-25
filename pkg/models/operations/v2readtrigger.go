// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/shared"
	"net/http"
)

type V2ReadTriggerRequest struct {
	// The trigger id
	TriggerID string `pathParam:"style=simple,explode=false,name=triggerID"`
}

func (o *V2ReadTriggerRequest) GetTriggerID() string {
	if o == nil {
		return ""
	}
	return o.TriggerID
}

type V2ReadTriggerResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// A specific trigger
	V2ReadTriggerResponse *shared.V2ReadTriggerResponse
}

func (o *V2ReadTriggerResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *V2ReadTriggerResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *V2ReadTriggerResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *V2ReadTriggerResponse) GetV2ReadTriggerResponse() *shared.V2ReadTriggerResponse {
	if o == nil {
		return nil
	}
	return o.V2ReadTriggerResponse
}
