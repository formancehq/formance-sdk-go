// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/shared"
	"net/http"
)

type V3UpdatePaymentMetadataRequest struct {
	V3UpdatePaymentMetadataRequest *shared.V3UpdatePaymentMetadataRequest `request:"mediaType=application/json"`
	// The payment ID
	PaymentID string `pathParam:"style=simple,explode=false,name=paymentID"`
}

func (o *V3UpdatePaymentMetadataRequest) GetV3UpdatePaymentMetadataRequest() *shared.V3UpdatePaymentMetadataRequest {
	if o == nil {
		return nil
	}
	return o.V3UpdatePaymentMetadataRequest
}

func (o *V3UpdatePaymentMetadataRequest) GetPaymentID() string {
	if o == nil {
		return ""
	}
	return o.PaymentID
}

type V3UpdatePaymentMetadataResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *V3UpdatePaymentMetadataResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *V3UpdatePaymentMetadataResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *V3UpdatePaymentMetadataResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
