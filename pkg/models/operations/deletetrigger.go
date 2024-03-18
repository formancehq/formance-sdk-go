// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/v2/v2/pkg/models/shared"
	"net/http"
)

type DeleteTriggerRequest struct {
	// The trigger id
	TriggerID string `pathParam:"style=simple,explode=false,name=triggerID"`
}

func (o *DeleteTriggerRequest) GetTriggerID() string {
	if o == nil {
		return ""
	}
	return o.TriggerID
}

type DeleteTriggerResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// General error
	Error *shared.Error
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *DeleteTriggerResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteTriggerResponse) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *DeleteTriggerResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteTriggerResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
