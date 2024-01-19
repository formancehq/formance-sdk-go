// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/v2/pkg/models/sdkerrors"
	"net/http"
)

type V2CancelEventRequest struct {
	// The instance id
	InstanceID string `pathParam:"style=simple,explode=false,name=instanceID"`
}

func (o *V2CancelEventRequest) GetInstanceID() string {
	if o == nil {
		return ""
	}
	return o.InstanceID
}

type V2CancelEventResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// General error
	V2Error *sdkerrors.V2Error
}

func (o *V2CancelEventResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *V2CancelEventResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *V2CancelEventResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *V2CancelEventResponse) GetV2Error() *sdkerrors.V2Error {
	if o == nil {
		return nil
	}
	return o.V2Error
}
