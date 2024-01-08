// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/pkg/models/sdkerrors"
	"github.com/formancehq/formance-sdk-go/pkg/models/shared"
	"github.com/formancehq/formance-sdk-go/pkg/utils"
	"math/big"
	"net/http"
)

type RevertTransactionRequest struct {
	// Allow to disable balances checks
	DisableChecks *bool `queryParam:"style=form,explode=true,name=disableChecks"`
	// Name of the ledger.
	Ledger string `pathParam:"style=simple,explode=false,name=ledger"`
	// Transaction ID.
	Txid *big.Int `pathParam:"style=simple,explode=false,name=txid"`
}

func (r RevertTransactionRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RevertTransactionRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *RevertTransactionRequest) GetDisableChecks() *bool {
	if o == nil {
		return nil
	}
	return o.DisableChecks
}

func (o *RevertTransactionRequest) GetLedger() string {
	if o == nil {
		return ""
	}
	return o.Ledger
}

func (o *RevertTransactionRequest) GetTxid() *big.Int {
	if o == nil {
		return big.NewInt(0)
	}
	return o.Txid
}

type RevertTransactionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Error
	ErrorResponse *sdkerrors.ErrorResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	TransactionResponse *shared.TransactionResponse
}

func (o *RevertTransactionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *RevertTransactionResponse) GetErrorResponse() *sdkerrors.ErrorResponse {
	if o == nil {
		return nil
	}
	return o.ErrorResponse
}

func (o *RevertTransactionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *RevertTransactionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *RevertTransactionResponse) GetTransactionResponse() *shared.TransactionResponse {
	if o == nil {
		return nil
	}
	return o.TransactionResponse
}
