// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/v2/pkg/models/shared"
	"net/http"
)

type ReadConnectorConfigV1Request struct {
	// The name of the connector.
	Connector shared.Connector `pathParam:"style=simple,explode=false,name=connector"`
	// The connector ID.
	ConnectorID string `pathParam:"style=simple,explode=false,name=connectorId"`
}

func (o *ReadConnectorConfigV1Request) GetConnector() shared.Connector {
	if o == nil {
		return shared.Connector("")
	}
	return o.Connector
}

func (o *ReadConnectorConfigV1Request) GetConnectorID() string {
	if o == nil {
		return ""
	}
	return o.ConnectorID
}

type ReadConnectorConfigV1Response struct {
	// OK
	ConnectorConfigResponse *shared.ConnectorConfigResponse
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ReadConnectorConfigV1Response) GetConnectorConfigResponse() *shared.ConnectorConfigResponse {
	if o == nil {
		return nil
	}
	return o.ConnectorConfigResponse
}

func (o *ReadConnectorConfigV1Response) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ReadConnectorConfigV1Response) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ReadConnectorConfigV1Response) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
