// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/v2/v2/pkg/models/sdkerrors"
	"github.com/formancehq/formance-sdk-go/v2/v2/pkg/models/shared"
	"github.com/formancehq/formance-sdk-go/v2/v2/pkg/utils"
	"net/http"
	"time"
)

type V2ListTransactionsRequest struct {
	RequestBody map[string]interface{} `request:"mediaType=application/json"`
	// Parameter used in pagination requests. Maximum page size is set to 15.
	// Set to the value of next for the next page of results.
	// Set to the value of previous for the previous page of results.
	// No other parameters can be set when this parameter is set.
	//
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
	Expand *string `queryParam:"style=form,explode=true,name=expand"`
	// Name of the ledger.
	Ledger string `pathParam:"style=simple,explode=false,name=ledger"`
	// The maximum number of results to return per page.
	//
	PageSize *int64     `queryParam:"style=form,explode=true,name=pageSize"`
	Pit      *time.Time `queryParam:"style=form,explode=true,name=pit"`
}

func (v V2ListTransactionsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(v, "", false)
}

func (v *V2ListTransactionsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &v, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *V2ListTransactionsRequest) GetRequestBody() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *V2ListTransactionsRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

func (o *V2ListTransactionsRequest) GetExpand() *string {
	if o == nil {
		return nil
	}
	return o.Expand
}

func (o *V2ListTransactionsRequest) GetLedger() string {
	if o == nil {
		return ""
	}
	return o.Ledger
}

func (o *V2ListTransactionsRequest) GetPageSize() *int64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *V2ListTransactionsRequest) GetPit() *time.Time {
	if o == nil {
		return nil
	}
	return o.Pit
}

type V2ListTransactionsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Error
	V2ErrorResponse *sdkerrors.V2ErrorResponse
	// OK
	V2TransactionsCursorResponse *shared.V2TransactionsCursorResponse
}

func (o *V2ListTransactionsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *V2ListTransactionsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *V2ListTransactionsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *V2ListTransactionsResponse) GetV2ErrorResponse() *sdkerrors.V2ErrorResponse {
	if o == nil {
		return nil
	}
	return o.V2ErrorResponse
}

func (o *V2ListTransactionsResponse) GetV2TransactionsCursorResponse() *shared.V2TransactionsCursorResponse {
	if o == nil {
		return nil
	}
	return o.V2TransactionsCursorResponse
}
