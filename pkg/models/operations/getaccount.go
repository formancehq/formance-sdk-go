// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/v2/pkg/models/shared"
	"net/http"
)

type GetAccountRequest struct {
	// Exact address of the account. It must match the following regular expressions pattern:
	// ```
	// ^\w+(:\w+)*$
	// ```
	//
	Address string `pathParam:"style=simple,explode=false,name=address"`
	// Name of the ledger.
	Ledger string `pathParam:"style=simple,explode=false,name=ledger"`
}

func (o *GetAccountRequest) GetAddress() string {
	if o == nil {
		return ""
	}
	return o.Address
}

func (o *GetAccountRequest) GetLedger() string {
	if o == nil {
		return ""
	}
	return o.Ledger
}

type GetAccountResponse struct {
	// OK
	AccountResponse *shared.AccountResponse
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetAccountResponse) GetAccountResponse() *shared.AccountResponse {
	if o == nil {
		return nil
	}
	return o.AccountResponse
}

func (o *GetAccountResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAccountResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAccountResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
