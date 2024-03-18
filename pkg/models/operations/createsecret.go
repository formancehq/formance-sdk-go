// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/v2/v2/pkg/models/shared"
	"net/http"
)

type CreateSecretRequest struct {
	CreateSecretRequest *shared.CreateSecretRequest `request:"mediaType=application/json"`
	// Client ID
	ClientID string `pathParam:"style=simple,explode=false,name=clientId"`
}

func (o *CreateSecretRequest) GetCreateSecretRequest() *shared.CreateSecretRequest {
	if o == nil {
		return nil
	}
	return o.CreateSecretRequest
}

func (o *CreateSecretRequest) GetClientID() string {
	if o == nil {
		return ""
	}
	return o.ClientID
}

type CreateSecretResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Created secret
	CreateSecretResponse *shared.CreateSecretResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateSecretResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateSecretResponse) GetCreateSecretResponse() *shared.CreateSecretResponse {
	if o == nil {
		return nil
	}
	return o.CreateSecretResponse
}

func (o *CreateSecretResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateSecretResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
