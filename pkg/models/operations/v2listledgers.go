// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/shared"
	"net/http"
)

type V2ListLedgersRequest struct {
	// Parameter used in pagination requests. Maximum page size is set to 15.
	// Set to the value of next for the next page of results.
	// Set to the value of previous for the previous page of results.
	// No other parameters can be set when this parameter is set.
	//
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
	// The maximum number of results to return per page.
	//
	PageSize *int64 `queryParam:"style=form,explode=true,name=pageSize"`
}

func (o *V2ListLedgersRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

func (o *V2ListLedgersRequest) GetPageSize() *int64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

type V2ListLedgersResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	V2LedgerListResponse *shared.V2LedgerListResponse
}

func (o *V2ListLedgersResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *V2ListLedgersResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *V2ListLedgersResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *V2ListLedgersResponse) GetV2LedgerListResponse() *shared.V2LedgerListResponse {
	if o == nil {
		return nil
	}
	return o.V2LedgerListResponse
}
