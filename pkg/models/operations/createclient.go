// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/shared"
	"net/http"
)

type CreateClientResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Client created
	CreateClientResponse *shared.CreateClientResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateClientResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateClientResponse) GetCreateClientResponse() *shared.CreateClientResponse {
	if o == nil {
		return nil
	}
	return o.CreateClientResponse
}

func (o *CreateClientResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateClientResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
