// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/shared"
	"net/http"
)

type V3ReversePaymentInitiationRequest struct {
	V3ReversePaymentInitiationRequest *shared.V3ReversePaymentInitiationRequest `request:"mediaType=application/json"`
	// The payment initiation ID
	PaymentInitiationID string `pathParam:"style=simple,explode=false,name=paymentInitiationID"`
}

func (o *V3ReversePaymentInitiationRequest) GetV3ReversePaymentInitiationRequest() *shared.V3ReversePaymentInitiationRequest {
	if o == nil {
		return nil
	}
	return o.V3ReversePaymentInitiationRequest
}

func (o *V3ReversePaymentInitiationRequest) GetPaymentInitiationID() string {
	if o == nil {
		return ""
	}
	return o.PaymentInitiationID
}

type V3ReversePaymentInitiationResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Accepted
	V3ReversePaymentInitiationResponse *shared.V3ReversePaymentInitiationResponse
}

func (o *V3ReversePaymentInitiationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *V3ReversePaymentInitiationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *V3ReversePaymentInitiationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *V3ReversePaymentInitiationResponse) GetV3ReversePaymentInitiationResponse() *shared.V3ReversePaymentInitiationResponse {
	if o == nil {
		return nil
	}
	return o.V3ReversePaymentInitiationResponse
}
