// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/shared"
	"net/http"
)

type V3ListPaymentServiceUsersRequest struct {
	RequestBody map[string]any `request:"mediaType=application/json"`
	// Parameter used in pagination requests. Set to the value of next for the next page of results. Set to the value of previous for the previous page of results. No other parameters can be set when this parameter is set.
	//
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
	// The number of items to return
	PageSize *int64 `queryParam:"style=form,explode=true,name=pageSize"`
}

func (o *V3ListPaymentServiceUsersRequest) GetRequestBody() map[string]any {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *V3ListPaymentServiceUsersRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

func (o *V3ListPaymentServiceUsersRequest) GetPageSize() *int64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

type V3ListPaymentServiceUsersResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	V3PaymentServiceUsersCursorResponse *shared.V3PaymentServiceUsersCursorResponse
}

func (o *V3ListPaymentServiceUsersResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *V3ListPaymentServiceUsersResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *V3ListPaymentServiceUsersResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *V3ListPaymentServiceUsersResponse) GetV3PaymentServiceUsersCursorResponse() *shared.V3PaymentServiceUsersCursorResponse {
	if o == nil {
		return nil
	}
	return o.V3PaymentServiceUsersCursorResponse
}
