// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/shared"
	"net/http"
)

type GetBalanceRequest struct {
	BalanceName string `pathParam:"style=simple,explode=false,name=balanceName"`
	ID          string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetBalanceRequest) GetBalanceName() string {
	if o == nil {
		return ""
	}
	return o.BalanceName
}

func (o *GetBalanceRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetBalanceResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Balance summary
	GetBalanceResponse *shared.GetBalanceResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetBalanceResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetBalanceResponse) GetGetBalanceResponse() *shared.GetBalanceResponse {
	if o == nil {
		return nil
	}
	return o.GetBalanceResponse
}

func (o *GetBalanceResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetBalanceResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
