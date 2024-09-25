// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/shared"
	"github.com/formancehq/formance-sdk-go/v3/pkg/utils"
	"net/http"
)

type GetHoldsRequest struct {
	// Parameter used in pagination requests.
	// Set to the value of next for the next page of results.
	// Set to the value of previous for the previous page of results.
	// No other parameters can be set when the pagination token is set.
	//
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
	// Filter holds by metadata key value pairs. Nested objects can be used as seen in the example below.
	Metadata map[string]string `queryParam:"style=deepObject,explode=true,name=metadata"`
	// The maximum number of results to return per page
	PageSize *int64 `default:"15" queryParam:"style=form,explode=true,name=pageSize"`
	// The wallet to filter on
	WalletID *string `queryParam:"style=form,explode=true,name=walletID"`
}

func (g GetHoldsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetHoldsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetHoldsRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

func (o *GetHoldsRequest) GetMetadata() map[string]string {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *GetHoldsRequest) GetPageSize() *int64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *GetHoldsRequest) GetWalletID() *string {
	if o == nil {
		return nil
	}
	return o.WalletID
}

type GetHoldsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Holds
	GetHoldsResponse *shared.GetHoldsResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetHoldsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetHoldsResponse) GetGetHoldsResponse() *shared.GetHoldsResponse {
	if o == nil {
		return nil
	}
	return o.GetHoldsResponse
}

func (o *GetHoldsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetHoldsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
