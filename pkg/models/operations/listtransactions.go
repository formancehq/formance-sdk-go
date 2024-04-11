// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/v2/pkg/models/shared"
	"github.com/formancehq/formance-sdk-go/v2/pkg/utils"
	"net/http"
	"time"
)

type ListTransactionsRequest struct {
	// Filter transactions with postings involving given account, either as source or destination (regular expression placed between ^ and $).
	Account *string `queryParam:"style=form,explode=true,name=account"`
	// Pagination cursor, will return transactions after given txid (in descending order).
	After *string `queryParam:"style=form,explode=true,name=after"`
	// Parameter used in pagination requests. Maximum page size is set to 1000.
	// Set to the value of next for the next page of results.
	// Set to the value of previous for the previous page of results.
	// No other parameters can be set when this parameter is set.
	//
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
	// Filter transactions with postings involving given account at destination (regular expression placed between ^ and $).
	Destination *string `queryParam:"style=form,explode=true,name=destination"`
	// Filter transactions that occurred before this timestamp.
	// The format is RFC3339 and is exclusive (for example, "2023-01-02T15:04:01Z" excludes the first second of 4th minute).
	//
	EndTime *time.Time `queryParam:"style=form,explode=true,name=endTime"`
	// Name of the ledger.
	Ledger string `pathParam:"style=simple,explode=false,name=ledger"`
	// Filter transactions by metadata key value pairs. Nested objects can be used as seen in the example below.
	Metadata map[string]interface{} `queryParam:"style=deepObject,explode=true,name=metadata"`
	// The maximum number of results to return per page.
	//
	PageSize *int64 `default:"15" queryParam:"style=form,explode=true,name=pageSize"`
	// Find transactions by reference field.
	Reference *string `queryParam:"style=form,explode=true,name=reference"`
	// Filter transactions with postings involving given account at source (regular expression placed between ^ and $).
	Source *string `queryParam:"style=form,explode=true,name=source"`
	// Filter transactions that occurred after this timestamp.
	// The format is RFC3339 and is inclusive (for example, "2023-01-02T15:04:01Z" includes the first second of 4th minute).
	//
	StartTime *time.Time `queryParam:"style=form,explode=true,name=startTime"`
}

func (l ListTransactionsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListTransactionsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListTransactionsRequest) GetAccount() *string {
	if o == nil {
		return nil
	}
	return o.Account
}

func (o *ListTransactionsRequest) GetAfter() *string {
	if o == nil {
		return nil
	}
	return o.After
}

func (o *ListTransactionsRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

func (o *ListTransactionsRequest) GetDestination() *string {
	if o == nil {
		return nil
	}
	return o.Destination
}

func (o *ListTransactionsRequest) GetEndTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.EndTime
}

func (o *ListTransactionsRequest) GetLedger() string {
	if o == nil {
		return ""
	}
	return o.Ledger
}

func (o *ListTransactionsRequest) GetMetadata() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *ListTransactionsRequest) GetPageSize() *int64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *ListTransactionsRequest) GetReference() *string {
	if o == nil {
		return nil
	}
	return o.Reference
}

func (o *ListTransactionsRequest) GetSource() *string {
	if o == nil {
		return nil
	}
	return o.Source
}

func (o *ListTransactionsRequest) GetStartTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartTime
}

type ListTransactionsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	TransactionsCursorResponse *shared.TransactionsCursorResponse
}

func (o *ListTransactionsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListTransactionsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListTransactionsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListTransactionsResponse) GetTransactionsCursorResponse() *shared.TransactionsCursorResponse {
	if o == nil {
		return nil
	}
	return o.TransactionsCursorResponse
}
