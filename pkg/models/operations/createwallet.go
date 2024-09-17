// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/shared"
	"net/http"
)

type CreateWalletRequest struct {
	CreateWalletRequest *shared.CreateWalletRequest `request:"mediaType=application/json"`
	// Use an idempotency key
	IdempotencyKey *string `header:"style=simple,explode=false,name=Idempotency-Key"`
}

func (o *CreateWalletRequest) GetCreateWalletRequest() *shared.CreateWalletRequest {
	if o == nil {
		return nil
	}
	return o.CreateWalletRequest
}

func (o *CreateWalletRequest) GetIdempotencyKey() *string {
	if o == nil {
		return nil
	}
	return o.IdempotencyKey
}

type CreateWalletResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Wallet created
	CreateWalletResponse *shared.CreateWalletResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateWalletResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateWalletResponse) GetCreateWalletResponse() *shared.CreateWalletResponse {
	if o == nil {
		return nil
	}
	return o.CreateWalletResponse
}

func (o *CreateWalletResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateWalletResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
