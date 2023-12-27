// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/pkg/models/shared"
	"net/http"
)

type TestConfigRequest struct {
	// Config ID
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *TestConfigRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type TestConfigResponse struct {
	// OK
	AttemptResponse *shared.AttemptResponse
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Error
	WebhooksErrorResponse *shared.WebhooksErrorResponse
}

func (o *TestConfigResponse) GetAttemptResponse() *shared.AttemptResponse {
	if o == nil {
		return nil
	}
	return o.AttemptResponse
}

func (o *TestConfigResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *TestConfigResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *TestConfigResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *TestConfigResponse) GetWebhooksErrorResponse() *shared.WebhooksErrorResponse {
	if o == nil {
		return nil
	}
	return o.WebhooksErrorResponse
}
