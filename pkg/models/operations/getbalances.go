// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/v2/pkg/models/sdkerrors"
	"github.com/formancehq/formance-sdk-go/v2/pkg/models/shared"
	"net/http"
)

type GetBalancesRequest struct {
	// Filter balances involving given account, either as source or destination.
	Address *string `queryParam:"style=form,explode=true,name=address"`
	// Pagination cursor, will return accounts after given address, in descending order.
	After *string `queryParam:"style=form,explode=true,name=after"`
	// Parameter used in pagination requests. Maximum page size is set to 15.
	// Set to the value of next for the next page of results.
	// Set to the value of previous for the previous page of results.
	// No other parameters can be set when this parameter is set.
	//
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
	// Name of the ledger.
	Ledger string `pathParam:"style=simple,explode=false,name=ledger"`
}

func (o *GetBalancesRequest) GetAddress() *string {
	if o == nil {
		return nil
	}
	return o.Address
}

func (o *GetBalancesRequest) GetAfter() *string {
	if o == nil {
		return nil
	}
	return o.After
}

func (o *GetBalancesRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

func (o *GetBalancesRequest) GetLedger() string {
	if o == nil {
		return ""
	}
	return o.Ledger
}

type GetBalancesResponse struct {
	// OK
	BalancesCursorResponse *shared.BalancesCursorResponse
	// HTTP response content type for this operation
	ContentType string
	// Error
	ErrorResponse *sdkerrors.ErrorResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetBalancesResponse) GetBalancesCursorResponse() *shared.BalancesCursorResponse {
	if o == nil {
		return nil
	}
	return o.BalancesCursorResponse
}

func (o *GetBalancesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetBalancesResponse) GetErrorResponse() *sdkerrors.ErrorResponse {
	if o == nil {
		return nil
	}
	return o.ErrorResponse
}

func (o *GetBalancesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetBalancesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
