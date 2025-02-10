// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/shared"
	"net/http"
)

type V3ListConnectorConfigsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	V3ConnectorConfigsResponse *shared.V3ConnectorConfigsResponse
}

func (o *V3ListConnectorConfigsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *V3ListConnectorConfigsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *V3ListConnectorConfigsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *V3ListConnectorConfigsResponse) GetV3ConnectorConfigsResponse() *shared.V3ConnectorConfigsResponse {
	if o == nil {
		return nil
	}
	return o.V3ConnectorConfigsResponse
}
