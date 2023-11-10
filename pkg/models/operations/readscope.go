// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/pkg/models/shared"
	"net/http"
)

type ReadScopeRequest struct {
	// Scope ID
	ScopeID string `pathParam:"style=simple,explode=false,name=scopeId"`
}

func (o *ReadScopeRequest) GetScopeID() string {
	if o == nil {
		return ""
	}
	return o.ScopeID
}

type ReadScopeResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Retrieved scope
	ReadScopeResponse *shared.ReadScopeResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ReadScopeResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ReadScopeResponse) GetReadScopeResponse() *shared.ReadScopeResponse {
	if o == nil {
		return nil
	}
	return o.ReadScopeResponse
}

func (o *ReadScopeResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ReadScopeResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
