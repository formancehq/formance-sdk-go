// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/formancehq/formance-sdk-go/v2/pkg/utils"
	"math/big"
	"time"
)

type TransferInitiationRequestType string

const (
	TransferInitiationRequestTypeTransfer TransferInitiationRequestType = "TRANSFER"
	TransferInitiationRequestTypePayout   TransferInitiationRequestType = "PAYOUT"
)

func (e TransferInitiationRequestType) ToPointer() *TransferInitiationRequestType {
	return &e
}
func (e *TransferInitiationRequestType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "TRANSFER":
		fallthrough
	case "PAYOUT":
		*e = TransferInitiationRequestType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TransferInitiationRequestType: %v", v)
	}
}

type TransferInitiationRequest struct {
	Amount               *big.Int                      `json:"amount"`
	Asset                string                        `json:"asset"`
	ConnectorID          *string                       `json:"connectorID,omitempty"`
	Description          string                        `json:"description"`
	DestinationAccountID string                        `json:"destinationAccountID"`
	Metadata             map[string]string             `json:"metadata,omitempty"`
	Provider             *Connector                    `json:"provider,omitempty"`
	Reference            string                        `json:"reference"`
	ScheduledAt          time.Time                     `json:"scheduledAt"`
	SourceAccountID      string                        `json:"sourceAccountID"`
	Type                 TransferInitiationRequestType `json:"type"`
	Validated            bool                          `json:"validated"`
}

func (t TransferInitiationRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TransferInitiationRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TransferInitiationRequest) GetAmount() *big.Int {
	if o == nil {
		return big.NewInt(0)
	}
	return o.Amount
}

func (o *TransferInitiationRequest) GetAsset() string {
	if o == nil {
		return ""
	}
	return o.Asset
}

func (o *TransferInitiationRequest) GetConnectorID() *string {
	if o == nil {
		return nil
	}
	return o.ConnectorID
}

func (o *TransferInitiationRequest) GetDescription() string {
	if o == nil {
		return ""
	}
	return o.Description
}

func (o *TransferInitiationRequest) GetDestinationAccountID() string {
	if o == nil {
		return ""
	}
	return o.DestinationAccountID
}

func (o *TransferInitiationRequest) GetMetadata() map[string]string {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *TransferInitiationRequest) GetProvider() *Connector {
	if o == nil {
		return nil
	}
	return o.Provider
}

func (o *TransferInitiationRequest) GetReference() string {
	if o == nil {
		return ""
	}
	return o.Reference
}

func (o *TransferInitiationRequest) GetScheduledAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.ScheduledAt
}

func (o *TransferInitiationRequest) GetSourceAccountID() string {
	if o == nil {
		return ""
	}
	return o.SourceAccountID
}

func (o *TransferInitiationRequest) GetType() TransferInitiationRequestType {
	if o == nil {
		return TransferInitiationRequestType("")
	}
	return o.Type
}

func (o *TransferInitiationRequest) GetValidated() bool {
	if o == nil {
		return false
	}
	return o.Validated
}
