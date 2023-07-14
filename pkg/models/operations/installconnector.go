// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/pkg/models/shared"
	"net/http"
)

type InstallConnectorRequest struct {
	RequestBody interface{} `request:"mediaType=application/json"`
	// The name of the connector.
	Connector shared.Connector `pathParam:"style=simple,explode=false,name=connector"`
}

type InstallConnectorResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}