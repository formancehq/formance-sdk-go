// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/v2/pkg/models/shared"
	"net/http"
)

type ReadClientRequest struct {
	// Client ID
	ClientID string `pathParam:"style=simple,explode=false,name=clientId"`
}

func (o *ReadClientRequest) GetClientID() string {
	if o == nil {
		return ""
	}
	return o.ClientID
}

type ReadClientResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Retrieved client
	ReadClientResponse *shared.ReadClientResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ReadClientResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ReadClientResponse) GetReadClientResponse() *shared.ReadClientResponse {
	if o == nil {
		return nil
	}
	return o.ReadClientResponse
}

func (o *ReadClientResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ReadClientResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
