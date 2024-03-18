// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/v2/v2/pkg/models/shared"
	"net/http"
)

type DebitWalletRequest struct {
	DebitWalletRequest *shared.DebitWalletRequest `request:"mediaType=application/json"`
	ID                 string                     `pathParam:"style=simple,explode=false,name=id"`
}

func (o *DebitWalletRequest) GetDebitWalletRequest() *shared.DebitWalletRequest {
	if o == nil {
		return nil
	}
	return o.DebitWalletRequest
}

func (o *DebitWalletRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type DebitWalletResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Wallet successfully debited as a pending hold
	DebitWalletResponse *shared.DebitWalletResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Error
	WalletsErrorResponse *shared.WalletsErrorResponse
}

func (o *DebitWalletResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DebitWalletResponse) GetDebitWalletResponse() *shared.DebitWalletResponse {
	if o == nil {
		return nil
	}
	return o.DebitWalletResponse
}

func (o *DebitWalletResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DebitWalletResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DebitWalletResponse) GetWalletsErrorResponse() *shared.WalletsErrorResponse {
	if o == nil {
		return nil
	}
	return o.WalletsErrorResponse
}
