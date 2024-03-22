// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/v2/pkg/models/sdkerrors"
	"github.com/formancehq/formance-sdk-go/v2/pkg/models/shared"
	"net/http"
)

type GetWalletRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetWalletRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetWalletResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Wallet
	GetWalletResponse *shared.GetWalletResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Error
	WalletsErrorResponse *sdkerrors.WalletsErrorResponse
}

func (o *GetWalletResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetWalletResponse) GetGetWalletResponse() *shared.GetWalletResponse {
	if o == nil {
		return nil
	}
	return o.GetWalletResponse
}

func (o *GetWalletResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetWalletResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetWalletResponse) GetWalletsErrorResponse() *sdkerrors.WalletsErrorResponse {
	if o == nil {
		return nil
	}
	return o.WalletsErrorResponse
}
