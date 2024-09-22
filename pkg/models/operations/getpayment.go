// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/v2/pkg/models/shared"
	"net/http"
)

type GetPaymentRequest struct {
	// The payment ID.
	PaymentID string `pathParam:"style=simple,explode=false,name=paymentId"`
}

func (o *GetPaymentRequest) GetPaymentID() string {
	if o == nil {
		return ""
	}
	return o.PaymentID
}

type GetPaymentResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	PaymentResponse *shared.PaymentResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetPaymentResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetPaymentResponse) GetPaymentResponse() *shared.PaymentResponse {
	if o == nil {
		return nil
	}
	return o.PaymentResponse
}

func (o *GetPaymentResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetPaymentResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
