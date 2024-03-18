// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/formancehq/formance-sdk-go/v2/v2/pkg/utils"
	"math/big"
)

// V2ActivityStripeTransferMetadata - A set of key/value pairs that you can attach to a transfer object.
// It can be useful for storing additional information about the transfer in a structured format.
type V2ActivityStripeTransferMetadata struct {
}

type V2ActivityStripeTransfer struct {
	Amount      *big.Int `json:"amount,omitempty"`
	Asset       *string  `json:"asset,omitempty"`
	ConnectorID *string  `json:"connectorID,omitempty"`
	Destination *string  `json:"destination,omitempty"`
	// A set of key/value pairs that you can attach to a transfer object.
	// It can be useful for storing additional information about the transfer in a structured format.
	//
	Metadata          *V2ActivityStripeTransferMetadata `json:"metadata,omitempty"`
	WaitingValidation *bool                             `default:"false" json:"waitingValidation"`
}

func (v V2ActivityStripeTransfer) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(v, "", false)
}

func (v *V2ActivityStripeTransfer) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &v, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *V2ActivityStripeTransfer) GetAmount() *big.Int {
	if o == nil {
		return nil
	}
	return o.Amount
}

func (o *V2ActivityStripeTransfer) GetAsset() *string {
	if o == nil {
		return nil
	}
	return o.Asset
}

func (o *V2ActivityStripeTransfer) GetConnectorID() *string {
	if o == nil {
		return nil
	}
	return o.ConnectorID
}

func (o *V2ActivityStripeTransfer) GetDestination() *string {
	if o == nil {
		return nil
	}
	return o.Destination
}

func (o *V2ActivityStripeTransfer) GetMetadata() *V2ActivityStripeTransferMetadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *V2ActivityStripeTransfer) GetWaitingValidation() *bool {
	if o == nil {
		return nil
	}
	return o.WaitingValidation
}
