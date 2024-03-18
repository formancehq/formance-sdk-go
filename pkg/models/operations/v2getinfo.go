// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/v2/v2/pkg/models/sdkerrors"
	"github.com/formancehq/formance-sdk-go/v2/v2/pkg/models/shared"
	"net/http"
)

type V2GetInfoResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	V2ConfigInfoResponse *shared.V2ConfigInfoResponse
	// Error
	V2ErrorResponse *sdkerrors.V2ErrorResponse
}

func (o *V2GetInfoResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *V2GetInfoResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *V2GetInfoResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *V2GetInfoResponse) GetV2ConfigInfoResponse() *shared.V2ConfigInfoResponse {
	if o == nil {
		return nil
	}
	return o.V2ConfigInfoResponse
}

func (o *V2GetInfoResponse) GetV2ErrorResponse() *sdkerrors.V2ErrorResponse {
	if o == nil {
		return nil
	}
	return o.V2ErrorResponse
}
