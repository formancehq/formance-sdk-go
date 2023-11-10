// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/formancehq/formance-sdk-go/pkg/utils"
	"math/big"
)

// StripeTransferRequestMetadata - A set of key/value pairs that you can attach to a transfer object.
// It can be useful for storing additional information about the transfer in a structured format.
type StripeTransferRequestMetadata struct {
}

type StripeTransferRequest struct {
	Amount      *big.Int `json:"amount,omitempty"`
	Asset       *string  `json:"asset,omitempty"`
	Destination *string  `json:"destination,omitempty"`
	// A set of key/value pairs that you can attach to a transfer object.
	// It can be useful for storing additional information about the transfer in a structured format.
	//
	Metadata *StripeTransferRequestMetadata `json:"metadata,omitempty"`
}

func (s StripeTransferRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *StripeTransferRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *StripeTransferRequest) GetAmount() *big.Int {
	if o == nil {
		return nil
	}
	return o.Amount
}

func (o *StripeTransferRequest) GetAsset() *string {
	if o == nil {
		return nil
	}
	return o.Asset
}

func (o *StripeTransferRequest) GetDestination() *string {
	if o == nil {
		return nil
	}
	return o.Destination
}

func (o *StripeTransferRequest) GetMetadata() *StripeTransferRequestMetadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}
