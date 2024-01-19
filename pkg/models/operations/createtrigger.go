// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/v2/pkg/models/shared"
	"net/http"
)

type CreateTriggerResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Created trigger
	CreateTriggerResponse *shared.CreateTriggerResponse
	// General error
	Error *shared.Error
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateTriggerResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateTriggerResponse) GetCreateTriggerResponse() *shared.CreateTriggerResponse {
	if o == nil {
		return nil
	}
	return o.CreateTriggerResponse
}

func (o *CreateTriggerResponse) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *CreateTriggerResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateTriggerResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
