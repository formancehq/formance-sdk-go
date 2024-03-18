// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/v2/v2/pkg/models/shared"
	"net/http"
)

type ReconcileRequest struct {
	ReconciliationRequest shared.ReconciliationRequest `request:"mediaType=application/json"`
	// The policy ID.
	PolicyID string `pathParam:"style=simple,explode=false,name=policyID"`
}

func (o *ReconcileRequest) GetReconciliationRequest() shared.ReconciliationRequest {
	if o == nil {
		return shared.ReconciliationRequest{}
	}
	return o.ReconciliationRequest
}

func (o *ReconcileRequest) GetPolicyID() string {
	if o == nil {
		return ""
	}
	return o.PolicyID
}

type ReconcileResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	ReconciliationResponse *shared.ReconciliationResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Error response
	ReconciliationErrorResponse *shared.ReconciliationErrorResponse
}

func (o *ReconcileResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ReconcileResponse) GetReconciliationResponse() *shared.ReconciliationResponse {
	if o == nil {
		return nil
	}
	return o.ReconciliationResponse
}

func (o *ReconcileResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ReconcileResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ReconcileResponse) GetReconciliationErrorResponse() *shared.ReconciliationErrorResponse {
	if o == nil {
		return nil
	}
	return o.ReconciliationErrorResponse
}
