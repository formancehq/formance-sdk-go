// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/v2/pkg/models/shared"
	"github.com/formancehq/formance-sdk-go/v2/pkg/utils"
	"net/http"
	"time"
)

type V2GetBalancesAggregatedRequest struct {
	RequestBody map[string]interface{} `request:"mediaType=application/json"`
	// Name of the ledger.
	Ledger string     `pathParam:"style=simple,explode=false,name=ledger"`
	Pit    *time.Time `queryParam:"style=form,explode=true,name=pit"`
	// Use insertion date instead of effective date
	UseInsertionDate *bool `queryParam:"style=form,explode=true,name=useInsertionDate"`
}

func (v V2GetBalancesAggregatedRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(v, "", false)
}

func (v *V2GetBalancesAggregatedRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &v, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *V2GetBalancesAggregatedRequest) GetRequestBody() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *V2GetBalancesAggregatedRequest) GetLedger() string {
	if o == nil {
		return ""
	}
	return o.Ledger
}

func (o *V2GetBalancesAggregatedRequest) GetPit() *time.Time {
	if o == nil {
		return nil
	}
	return o.Pit
}

func (o *V2GetBalancesAggregatedRequest) GetUseInsertionDate() *bool {
	if o == nil {
		return nil
	}
	return o.UseInsertionDate
}

type V2GetBalancesAggregatedResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	V2AggregateBalancesResponse *shared.V2AggregateBalancesResponse
}

func (o *V2GetBalancesAggregatedResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *V2GetBalancesAggregatedResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *V2GetBalancesAggregatedResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *V2GetBalancesAggregatedResponse) GetV2AggregateBalancesResponse() *shared.V2AggregateBalancesResponse {
	if o == nil {
		return nil
	}
	return o.V2AggregateBalancesResponse
}
