// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/v2/pkg/models/shared"
	"net/http"
)

type V2ReadStatsRequest struct {
	// name of the ledger
	Ledger string `pathParam:"style=simple,explode=false,name=ledger"`
}

func (o *V2ReadStatsRequest) GetLedger() string {
	if o == nil {
		return ""
	}
	return o.Ledger
}

type V2ReadStatsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	V2StatsResponse *shared.V2StatsResponse
}

func (o *V2ReadStatsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *V2ReadStatsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *V2ReadStatsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *V2ReadStatsResponse) GetV2StatsResponse() *shared.V2StatsResponse {
	if o == nil {
		return nil
	}
	return o.V2StatsResponse
}
