// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/shared"
	"github.com/formancehq/formance-sdk-go/v3/pkg/utils"
	"net/http"
	"time"
)

type V3GetAccountBalancesRequest struct {
	// The account ID
	AccountID string `pathParam:"style=simple,explode=false,name=accountID"`
	// The asset to filter by
	Asset *string `queryParam:"style=form,explode=true,name=asset"`
	// Parameter used in pagination requests. Set to the value of next for the next page of results. Set to the value of previous for the previous page of results. No other parameters can be set when this parameter is set.
	//
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
	// The start of the time range to filter by
	FromTimestamp *time.Time `queryParam:"style=form,explode=true,name=fromTimestamp"`
	// The number of items to return
	PageSize *int64 `queryParam:"style=form,explode=true,name=pageSize"`
	// The end of the time range to filter by
	ToTimestamp *time.Time `queryParam:"style=form,explode=true,name=toTimestamp"`
}

func (v V3GetAccountBalancesRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(v, "", false)
}

func (v *V3GetAccountBalancesRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &v, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *V3GetAccountBalancesRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *V3GetAccountBalancesRequest) GetAsset() *string {
	if o == nil {
		return nil
	}
	return o.Asset
}

func (o *V3GetAccountBalancesRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

func (o *V3GetAccountBalancesRequest) GetFromTimestamp() *time.Time {
	if o == nil {
		return nil
	}
	return o.FromTimestamp
}

func (o *V3GetAccountBalancesRequest) GetPageSize() *int64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *V3GetAccountBalancesRequest) GetToTimestamp() *time.Time {
	if o == nil {
		return nil
	}
	return o.ToTimestamp
}

type V3GetAccountBalancesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	V3BalancesCursorResponse *shared.V3BalancesCursorResponse
}

func (o *V3GetAccountBalancesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *V3GetAccountBalancesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *V3GetAccountBalancesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *V3GetAccountBalancesResponse) GetV3BalancesCursorResponse() *shared.V3BalancesCursorResponse {
	if o == nil {
		return nil
	}
	return o.V3BalancesCursorResponse
}
