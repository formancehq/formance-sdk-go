// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/shared"
	"github.com/formancehq/formance-sdk-go/v3/pkg/utils"
	"net/http"
	"time"
)

type V2GetVolumesWithBalancesRequest struct {
	RequestBody map[string]any `request:"mediaType=application/json"`
	// Parameter used in pagination requests. Maximum page size is set to 15.
	// Set to the value of next for the next page of results.
	// Set to the value of previous for the previous page of results.
	// No other parameters can be set when this parameter is set.
	//
	Cursor  *string    `queryParam:"style=form,explode=true,name=cursor"`
	EndTime *time.Time `queryParam:"style=form,explode=true,name=endTime"`
	// Group volumes and balance by the level of the segment of the address
	GroupBy *int64 `queryParam:"style=form,explode=true,name=groupBy"`
	// Use insertion date instead of effective date
	InsertionDate *bool `queryParam:"style=form,explode=true,name=insertionDate"`
	// Name of the ledger.
	Ledger string `pathParam:"style=simple,explode=false,name=ledger"`
	// The maximum number of results to return per page.
	//
	PageSize  *int64     `queryParam:"style=form,explode=true,name=pageSize"`
	StartTime *time.Time `queryParam:"style=form,explode=true,name=startTime"`
}

func (v V2GetVolumesWithBalancesRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(v, "", false)
}

func (v *V2GetVolumesWithBalancesRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &v, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *V2GetVolumesWithBalancesRequest) GetRequestBody() map[string]any {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *V2GetVolumesWithBalancesRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

func (o *V2GetVolumesWithBalancesRequest) GetEndTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.EndTime
}

func (o *V2GetVolumesWithBalancesRequest) GetGroupBy() *int64 {
	if o == nil {
		return nil
	}
	return o.GroupBy
}

func (o *V2GetVolumesWithBalancesRequest) GetInsertionDate() *bool {
	if o == nil {
		return nil
	}
	return o.InsertionDate
}

func (o *V2GetVolumesWithBalancesRequest) GetLedger() string {
	if o == nil {
		return ""
	}
	return o.Ledger
}

func (o *V2GetVolumesWithBalancesRequest) GetPageSize() *int64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *V2GetVolumesWithBalancesRequest) GetStartTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartTime
}

type V2GetVolumesWithBalancesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	V2VolumesWithBalanceCursorResponse *shared.V2VolumesWithBalanceCursorResponse
}

func (o *V2GetVolumesWithBalancesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *V2GetVolumesWithBalancesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *V2GetVolumesWithBalancesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *V2GetVolumesWithBalancesResponse) GetV2VolumesWithBalanceCursorResponse() *shared.V2VolumesWithBalanceCursorResponse {
	if o == nil {
		return nil
	}
	return o.V2VolumesWithBalanceCursorResponse
}
