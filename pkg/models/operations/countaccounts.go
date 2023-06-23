// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/pkg/models/shared"
	"net/http"
)

// CountAccountsMetadata - Filter accounts by metadata key value pairs. The filter can be used like this metadata[key]=value1&metadata[a.nested.key]=value2
type CountAccountsMetadata struct {
}

type CountAccountsRequest struct {
	// Filter accounts by address pattern (regular expression placed between ^ and $).
	Address *string `queryParam:"style=form,explode=true,name=address"`
	// Name of the ledger.
	Ledger string `pathParam:"style=simple,explode=false,name=ledger"`
	// Filter accounts by metadata key value pairs. The filter can be used like this metadata[key]=value1&metadata[a.nested.key]=value2
	Metadata *CountAccountsMetadata `queryParam:"style=deepObject,explode=true,name=metadata"`
}

type CountAccountsResponse struct {
	ContentType string
	// Error
	ErrorResponse *shared.ErrorResponse
	Headers       map[string][]string
	StatusCode    int
	RawResponse   *http.Response
}
