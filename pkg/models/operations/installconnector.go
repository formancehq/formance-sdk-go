// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/pkg/models/shared"
	"net/http"
)

type InstallConnectorRequest struct {
	ConnectorConfig shared.ConnectorConfig `request:"mediaType=application/json"`
	// The name of the connector.
	Connector shared.Connector `pathParam:"style=simple,explode=false,name=connector"`
}

func (o *InstallConnectorRequest) GetConnectorConfig() shared.ConnectorConfig {
	if o == nil {
		return shared.ConnectorConfig{}
	}
	return o.ConnectorConfig
}

func (o *InstallConnectorRequest) GetConnector() shared.Connector {
	if o == nil {
		return shared.Connector("")
	}
	return o.Connector
}

type InstallConnectorResponse struct {
	// OK
	ConnectorResponse *shared.ConnectorResponse
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *InstallConnectorResponse) GetConnectorResponse() *shared.ConnectorResponse {
	if o == nil {
		return nil
	}
	return o.ConnectorResponse
}

func (o *InstallConnectorResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *InstallConnectorResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *InstallConnectorResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
