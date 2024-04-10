// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/v2/pkg/models/shared"
	"net/http"
)

type GetHoldRequest struct {
	// The hold ID
	HoldID string `pathParam:"style=simple,explode=false,name=holdID"`
}

func (o *GetHoldRequest) GetHoldID() string {
	if o == nil {
		return ""
	}
	return o.HoldID
}

type GetHoldResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Holds
	GetHoldResponse *shared.GetHoldResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetHoldResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetHoldResponse) GetGetHoldResponse() *shared.GetHoldResponse {
	if o == nil {
		return nil
	}
	return o.GetHoldResponse
}

func (o *GetHoldResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetHoldResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
