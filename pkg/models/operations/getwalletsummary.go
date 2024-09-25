// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/shared"
	"net/http"
)

type GetWalletSummaryRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetWalletSummaryRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetWalletSummaryResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Wallet summary
	GetWalletSummaryResponse *shared.GetWalletSummaryResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetWalletSummaryResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetWalletSummaryResponse) GetGetWalletSummaryResponse() *shared.GetWalletSummaryResponse {
	if o == nil {
		return nil
	}
	return o.GetWalletSummaryResponse
}

func (o *GetWalletSummaryResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetWalletSummaryResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
