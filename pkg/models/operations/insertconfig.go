// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/pkg/models/shared"
	"net/http"
)

type InsertConfigResponse struct {
	// Config created successfully.
	ConfigResponse *shared.ConfigResponse
	ContentType    string
	StatusCode     int
	RawResponse    *http.Response
	// Error
	WebhooksErrorResponse *shared.WebhooksErrorResponse
}