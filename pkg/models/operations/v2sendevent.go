// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"net/http"
)

type V2SendEventRequestBody struct {
	Name string `json:"name"`
}

func (o *V2SendEventRequestBody) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

type V2SendEventRequest struct {
	RequestBody *V2SendEventRequestBody `request:"mediaType=application/json"`
	// The instance id
	InstanceID string `pathParam:"style=simple,explode=false,name=instanceID"`
}

func (o *V2SendEventRequest) GetRequestBody() *V2SendEventRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *V2SendEventRequest) GetInstanceID() string {
	if o == nil {
		return ""
	}
	return o.InstanceID
}

type V2SendEventResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *V2SendEventResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *V2SendEventResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *V2SendEventResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
