// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/shared"
	"net/http"
)

type V3GetBankAccountRequest struct {
	// The bank account ID
	BankAccountID string `pathParam:"style=simple,explode=false,name=bankAccountID"`
}

func (o *V3GetBankAccountRequest) GetBankAccountID() string {
	if o == nil {
		return ""
	}
	return o.BankAccountID
}

type V3GetBankAccountResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	V3GetBankAccountResponse *shared.V3GetBankAccountResponse
}

func (o *V3GetBankAccountResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *V3GetBankAccountResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *V3GetBankAccountResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *V3GetBankAccountResponse) GetV3GetBankAccountResponse() *shared.V3GetBankAccountResponse {
	if o == nil {
		return nil
	}
	return o.V3GetBankAccountResponse
}
