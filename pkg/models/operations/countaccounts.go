// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"net/http"
)

type CountAccountsRequest struct {
	// Filter accounts by address pattern (regular expression placed between ^ and $).
	Address *string `queryParam:"style=form,explode=true,name=address"`
	// Name of the ledger.
	Ledger string `pathParam:"style=simple,explode=false,name=ledger"`
	// Filter accounts by metadata key value pairs. The filter can be used like this metadata[key]=value1&metadata[a.nested.key]=value2
	Metadata map[string]any `queryParam:"style=deepObject,explode=true,name=metadata"`
}

func (o *CountAccountsRequest) GetAddress() *string {
	if o == nil {
		return nil
	}
	return o.Address
}

func (o *CountAccountsRequest) GetLedger() string {
	if o == nil {
		return ""
	}
	return o.Ledger
}

func (o *CountAccountsRequest) GetMetadata() map[string]any {
	if o == nil {
		return nil
	}
	return o.Metadata
}

type CountAccountsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CountAccountsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CountAccountsResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}

func (o *CountAccountsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CountAccountsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
