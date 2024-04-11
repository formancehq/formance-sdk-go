// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/v2/pkg/models/shared"
	"github.com/formancehq/formance-sdk-go/v2/pkg/utils"
	"net/http"
	"time"
)

type ListLogsRequest struct {
	// Pagination cursor, will return the logs after a given ID. (in descending order).
	After *string `queryParam:"style=form,explode=true,name=after"`
	// Parameter used in pagination requests. Maximum page size is set to 1000.
	// Set to the value of next for the next page of results.
	// Set to the value of previous for the previous page of results.
	// No other parameters can be set when this parameter is set.
	//
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
	// Filter transactions that occurred before this timestamp.
	// The format is RFC3339 and is exclusive (for example, "2023-01-02T15:04:01Z" excludes the first second of 4th minute).
	//
	EndTime *time.Time `queryParam:"style=form,explode=true,name=endTime"`
	// Name of the ledger.
	Ledger string `pathParam:"style=simple,explode=false,name=ledger"`
	// The maximum number of results to return per page.
	//
	PageSize *int64 `default:"15" queryParam:"style=form,explode=true,name=pageSize"`
	// Filter transactions that occurred after this timestamp.
	// The format is RFC3339 and is inclusive (for example, "2023-01-02T15:04:01Z" includes the first second of 4th minute).
	//
	StartTime *time.Time `queryParam:"style=form,explode=true,name=startTime"`
}

func (l ListLogsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListLogsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListLogsRequest) GetAfter() *string {
	if o == nil {
		return nil
	}
	return o.After
}

func (o *ListLogsRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

func (o *ListLogsRequest) GetEndTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.EndTime
}

func (o *ListLogsRequest) GetLedger() string {
	if o == nil {
		return ""
	}
	return o.Ledger
}

func (o *ListLogsRequest) GetPageSize() *int64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *ListLogsRequest) GetStartTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartTime
}

type ListLogsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	LogsCursorResponse *shared.LogsCursorResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ListLogsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListLogsResponse) GetLogsCursorResponse() *shared.LogsCursorResponse {
	if o == nil {
		return nil
	}
	return o.LogsCursorResponse
}

func (o *ListLogsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListLogsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
