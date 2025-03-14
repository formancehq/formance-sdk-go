// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/shared"
	"github.com/formancehq/formance-sdk-go/v3/pkg/utils"
	"net/http"
)

type V3InitiatePaymentRequest struct {
	V3InitiatePaymentRequest *shared.V3InitiatePaymentRequest `request:"mediaType=application/json"`
	// If set to true, the request will not have to be validated. This is useful if we want to directly forward the request to the PSP.
	//
	NoValidation *bool `default:"false" queryParam:"style=form,explode=true,name=noValidation"`
}

func (v V3InitiatePaymentRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(v, "", false)
}

func (v *V3InitiatePaymentRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &v, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *V3InitiatePaymentRequest) GetV3InitiatePaymentRequest() *shared.V3InitiatePaymentRequest {
	if o == nil {
		return nil
	}
	return o.V3InitiatePaymentRequest
}

func (o *V3InitiatePaymentRequest) GetNoValidation() *bool {
	if o == nil {
		return nil
	}
	return o.NoValidation
}

type V3InitiatePaymentResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Accepted
	V3InitiatePaymentResponse *shared.V3InitiatePaymentResponse
}

func (o *V3InitiatePaymentResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *V3InitiatePaymentResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *V3InitiatePaymentResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *V3InitiatePaymentResponse) GetV3InitiatePaymentResponse() *shared.V3InitiatePaymentResponse {
	if o == nil {
		return nil
	}
	return o.V3InitiatePaymentResponse
}
