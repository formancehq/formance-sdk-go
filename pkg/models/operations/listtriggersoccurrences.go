// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/shared"
	"net/http"
)

type ListTriggersOccurrencesRequest struct {
	// The trigger id
	TriggerID string `pathParam:"style=simple,explode=false,name=triggerID"`
}

func (o *ListTriggersOccurrencesRequest) GetTriggerID() string {
	if o == nil {
		return ""
	}
	return o.TriggerID
}

type ListTriggersOccurrencesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// List of triggers occurrences
	ListTriggersOccurrencesResponse *shared.ListTriggersOccurrencesResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ListTriggersOccurrencesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListTriggersOccurrencesResponse) GetListTriggersOccurrencesResponse() *shared.ListTriggersOccurrencesResponse {
	if o == nil {
		return nil
	}
	return o.ListTriggersOccurrencesResponse
}

func (o *ListTriggersOccurrencesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListTriggersOccurrencesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
