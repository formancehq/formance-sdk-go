// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/shared"
	"github.com/formancehq/formance-sdk-go/v3/pkg/utils"
	"net/http"
	"time"
)

type V2ListAccountsRequest struct {
	RequestBody map[string]any `request:"mediaType=application/json"`
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

func (v V2ListAccountsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(v, "", false)
}

func (v *V2ListAccountsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &v, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *V2ListAccountsRequest) GetRequestBody() map[string]any {
	if o == nil {
		return map[string]any{}
	}
	return o.RequestBody
}

func (o *V2ListAccountsRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

func (o *V2ListAccountsRequest) GetExpand() *string {
	if o == nil {
		return nil
	}
	return o.Expand
}

func (o *V2ListAccountsRequest) GetLedger() string {
	if o == nil {
		return ""
	}
	return o.Ledger
}

func (o *V2ListAccountsRequest) GetPageSize() *int64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *V2ListAccountsRequest) GetPit() *time.Time {
	if o == nil {
		return nil
	}
	return o.Pit
}

type V2ListAccountsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	V2AccountsCursorResponse *shared.V2AccountsCursorResponse
}

func (o *V2ListAccountsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *V2ListAccountsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *V2ListAccountsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *V2ListAccountsResponse) GetV2AccountsCursorResponse() *shared.V2AccountsCursorResponse {
	if o == nil {
		return nil
	}
	return o.V2AccountsCursorResponse
}
