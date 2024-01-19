// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/v2/pkg/models/sdkerrors"
	"github.com/formancehq/formance-sdk-go/v2/pkg/utils"
	"math/big"
	"net/http"
)

type V2AddMetadataOnTransactionRequest struct {
	// Use an idempotency key
	IdempotencyKey *string `header:"style=simple,explode=false,name=Idempotency-Key"`
	// metadata
	RequestBody map[string]string `request:"mediaType=application/json"`
	// Set the dryRun mode. Dry run mode doesn't add the logs to the database or publish a message to the message broker.
	DryRun *bool `queryParam:"style=form,explode=true,name=dryRun"`
	// Transaction ID.
	ID *big.Int `pathParam:"style=simple,explode=false,name=id"`
	// Name of the ledger.
	Ledger string `pathParam:"style=simple,explode=false,name=ledger"`
}

func (v V2AddMetadataOnTransactionRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(v, "", false)
}

func (v *V2AddMetadataOnTransactionRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &v, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *V2AddMetadataOnTransactionRequest) GetIdempotencyKey() *string {
	if o == nil {
		return nil
	}
	return o.IdempotencyKey
}

func (o *V2AddMetadataOnTransactionRequest) GetRequestBody() map[string]string {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *V2AddMetadataOnTransactionRequest) GetDryRun() *bool {
	if o == nil {
		return nil
	}
	return o.DryRun
}

func (o *V2AddMetadataOnTransactionRequest) GetID() *big.Int {
	if o == nil {
		return big.NewInt(0)
	}
	return o.ID
}

func (o *V2AddMetadataOnTransactionRequest) GetLedger() string {
	if o == nil {
		return ""
	}
	return o.Ledger
}

type V2AddMetadataOnTransactionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Error
	V2ErrorResponse *sdkerrors.V2ErrorResponse
}

func (o *V2AddMetadataOnTransactionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *V2AddMetadataOnTransactionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *V2AddMetadataOnTransactionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *V2AddMetadataOnTransactionResponse) GetV2ErrorResponse() *sdkerrors.V2ErrorResponse {
	if o == nil {
		return nil
	}
	return o.V2ErrorResponse
}
