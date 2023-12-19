// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/pkg/models/shared"
	"net/http"
)

type DeleteConfigRequest struct {
	// Config ID
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *DeleteConfigRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type DeleteConfigResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Error
	WebhooksErrorResponse *shared.WebhooksErrorResponse
}

func (o *DeleteConfigResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteConfigResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteConfigResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteConfigResponse) GetWebhooksErrorResponse() *shared.WebhooksErrorResponse {
	if o == nil {
		return nil
	}
	return o.WebhooksErrorResponse
}
