// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/shared"
	"net/http"
)

type V3ApprovePaymentInitiationRequest struct {
	// The payment initiation ID
	PaymentInitiationID string `pathParam:"style=simple,explode=false,name=paymentInitiationID"`
}

func (o *V3ApprovePaymentInitiationRequest) GetPaymentInitiationID() string {
	if o == nil {
		return ""
	}
	return o.PaymentInitiationID
}

type V3ApprovePaymentInitiationResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Accepted
	V3ApprovePaymentInitiationResponse *shared.V3ApprovePaymentInitiationResponse
}

func (o *V3ApprovePaymentInitiationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *V3ApprovePaymentInitiationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *V3ApprovePaymentInitiationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *V3ApprovePaymentInitiationResponse) GetV3ApprovePaymentInitiationResponse() *shared.V3ApprovePaymentInitiationResponse {
	if o == nil {
		return nil
	}
	return o.V3ApprovePaymentInitiationResponse
}
