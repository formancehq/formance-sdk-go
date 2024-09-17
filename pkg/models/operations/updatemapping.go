// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/shared"
	"net/http"
)

type UpdateMappingRequest struct {
	Mapping *shared.Mapping `request:"mediaType=application/json"`
	// Name of the ledger.
	Ledger string `pathParam:"style=simple,explode=false,name=ledger"`
}

func (o *UpdateMappingRequest) GetMapping() *shared.Mapping {
	if o == nil {
		return nil
	}
	return o.Mapping
}

func (o *UpdateMappingRequest) GetLedger() string {
	if o == nil {
		return ""
	}
	return o.Ledger
}

type UpdateMappingResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	MappingResponse *shared.MappingResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *UpdateMappingResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateMappingResponse) GetMappingResponse() *shared.MappingResponse {
	if o == nil {
		return nil
	}
	return o.MappingResponse
}

func (o *UpdateMappingResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateMappingResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
