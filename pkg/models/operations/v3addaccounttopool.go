// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"net/http"
)

type V3AddAccountToPoolRequest struct {
	// The account ID
	AccountID string `pathParam:"style=simple,explode=false,name=accountID"`
	// The pool ID
	PoolID string `pathParam:"style=simple,explode=false,name=poolID"`
}

func (o *V3AddAccountToPoolRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *V3AddAccountToPoolRequest) GetPoolID() string {
	if o == nil {
		return ""
	}
	return o.PoolID
}

type V3AddAccountToPoolResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *V3AddAccountToPoolResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *V3AddAccountToPoolResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *V3AddAccountToPoolResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
