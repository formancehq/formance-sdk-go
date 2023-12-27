// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/pkg/models/shared"
	"net/http"
)

type V2AddMetadataToAccountRequest struct {
	// Use an idempotency key
	IdempotencyKey *string `header:"style=simple,explode=false,name=Idempotency-Key"`
	// metadata
	RequestBody map[string]string `request:"mediaType=application/json"`
	// Exact address of the account. It must match the following regular expressions pattern:
	// ```
	// ^\w+(:\w+)*$
	// ```
	//
	Address string `pathParam:"style=simple,explode=false,name=address"`
	// Set the dry run mode. Dry run mode doesn't add the logs to the database or publish a message to the message broker.
	DryRun *bool `queryParam:"style=form,explode=true,name=dryRun"`
	// Name of the ledger.
	Ledger string `pathParam:"style=simple,explode=false,name=ledger"`
}

func (o *V2AddMetadataToAccountRequest) GetIdempotencyKey() *string {
	if o == nil {
		return nil
	}
	return o.IdempotencyKey
}

func (o *V2AddMetadataToAccountRequest) GetRequestBody() map[string]string {
	if o == nil {
		return map[string]string{}
	}
	return o.RequestBody
}

func (o *V2AddMetadataToAccountRequest) GetAddress() string {
	if o == nil {
		return ""
	}
	return o.Address
}

func (o *V2AddMetadataToAccountRequest) GetDryRun() *bool {
	if o == nil {
		return nil
	}
	return o.DryRun
}

func (o *V2AddMetadataToAccountRequest) GetLedger() string {
	if o == nil {
		return ""
	}
	return o.Ledger
}

type V2AddMetadataToAccountResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Error
	V2ErrorResponse *shared.V2ErrorResponse
}

func (o *V2AddMetadataToAccountResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *V2AddMetadataToAccountResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *V2AddMetadataToAccountResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *V2AddMetadataToAccountResponse) GetV2ErrorResponse() *shared.V2ErrorResponse {
	if o == nil {
		return nil
	}
	return o.V2ErrorResponse
}
