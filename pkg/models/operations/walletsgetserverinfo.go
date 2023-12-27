// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/pkg/models/shared"
	"net/http"
)

type WalletsgetServerInfoResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Server information
	ServerInfo *shared.ServerInfo
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Error
	WalletsErrorResponse *shared.WalletsErrorResponse
}

func (o *WalletsgetServerInfoResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *WalletsgetServerInfoResponse) GetServerInfo() *shared.ServerInfo {
	if o == nil {
		return nil
	}
	return o.ServerInfo
}

func (o *WalletsgetServerInfoResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *WalletsgetServerInfoResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *WalletsgetServerInfoResponse) GetWalletsErrorResponse() *shared.WalletsErrorResponse {
	if o == nil {
		return nil
	}
	return o.WalletsErrorResponse
}
