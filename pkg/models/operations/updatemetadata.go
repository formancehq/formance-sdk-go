// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/pkg/models/shared"
	"net/http"
)

type UpdateMetadataRequest struct {
	PaymentMetadata *shared.PaymentMetadata `request:"mediaType=application/json"`
	// The payment ID.
	PaymentID string `pathParam:"style=simple,explode=false,name=paymentId"`
}

func (o *UpdateMetadataRequest) GetPaymentMetadata() *shared.PaymentMetadata {
	if o == nil {
		return nil
	}
	return o.PaymentMetadata
}

func (o *UpdateMetadataRequest) GetPaymentID() string {
	if o == nil {
		return ""
	}
	return o.PaymentID
}

type UpdateMetadataResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *UpdateMetadataResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateMetadataResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateMetadataResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
