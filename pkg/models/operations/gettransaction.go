// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/pkg/models/sdkerrors"
	"github.com/formancehq/formance-sdk-go/pkg/models/shared"
	"github.com/formancehq/formance-sdk-go/pkg/utils"
	"math/big"
	"net/http"
)

type GetTransactionRequest struct {
	// Name of the ledger.
	Ledger string `pathParam:"style=simple,explode=false,name=ledger"`
	// Transaction ID.
	Txid *big.Int `pathParam:"style=simple,explode=false,name=txid"`
}

func (g GetTransactionRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetTransactionRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetTransactionRequest) GetLedger() string {
	if o == nil {
		return ""
	}
	return o.Ledger
}

func (o *GetTransactionRequest) GetTxid() *big.Int {
	if o == nil {
		return big.NewInt(0)
	}
	return o.Txid
}

type GetTransactionResponse struct {
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

func (o *GetTransactionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetTransactionResponse) GetErrorResponse() *sdkerrors.ErrorResponse {
	if o == nil {
		return nil
	}
	return o.ErrorResponse
}

func (o *GetTransactionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetTransactionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetTransactionResponse) GetTransactionResponse() *shared.TransactionResponse {
	if o == nil {
		return nil
	}
	return o.TransactionResponse
}
