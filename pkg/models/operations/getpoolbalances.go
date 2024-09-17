// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/shared"
	"github.com/formancehq/formance-sdk-go/v3/pkg/utils"
	"net/http"
	"time"
)

type GetPoolBalancesRequest struct {
	// Filter balances by date.
	//
	At time.Time `queryParam:"style=form,explode=true,name=at"`
	// The pool ID.
	PoolID string `pathParam:"style=simple,explode=false,name=poolId"`
}

func (g GetPoolBalancesRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetPoolBalancesRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetPoolBalancesRequest) GetAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.At
}

func (o *GetPoolBalancesRequest) GetPoolID() string {
	if o == nil {
		return ""
	}
	return o.PoolID
}

type GetPoolBalancesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	PoolBalancesResponse *shared.PoolBalancesResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetPoolBalancesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetPoolBalancesResponse) GetPoolBalancesResponse() *shared.PoolBalancesResponse {
	if o == nil {
		return nil
	}
	return o.PoolBalancesResponse
}

func (o *GetPoolBalancesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetPoolBalancesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
