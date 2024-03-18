// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/v2/v2/pkg/models/shared"
	"net/http"
)

type AddAccountToPoolRequest struct {
	AddAccountToPoolRequest shared.AddAccountToPoolRequest `request:"mediaType=application/json"`
	// The pool ID.
	PoolID string `pathParam:"style=simple,explode=false,name=poolId"`
}

func (o *AddAccountToPoolRequest) GetAddAccountToPoolRequest() shared.AddAccountToPoolRequest {
	if o == nil {
		return shared.AddAccountToPoolRequest{}
	}
	return o.AddAccountToPoolRequest
}

func (o *AddAccountToPoolRequest) GetPoolID() string {
	if o == nil {
		return ""
	}
	return o.PoolID
}

type AddAccountToPoolResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *AddAccountToPoolResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *AddAccountToPoolResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *AddAccountToPoolResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
