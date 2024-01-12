// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/formancehq/formance-sdk-go/pkg/utils"
	"math/big"
	"time"
)

type TransferInitiationType string

const (
	TransferInitiationTypeTransfer TransferInitiationType = "TRANSFER"
	TransferInitiationTypePayout   TransferInitiationType = "PAYOUT"
)

func (e TransferInitiationType) ToPointer() *TransferInitiationType {
	return &e
}

func (e *TransferInitiationType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "TRANSFER":
		fallthrough
	case "PAYOUT":
		*e = TransferInitiationType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TransferInitiationType: %v", v)
	}
}

type TransferInitiation struct {
	Amount               *big.Int                       `json:"amount"`
	Asset                string                         `json:"asset"`
	ConnectorID          string                         `json:"connectorID"`
	CreatedAt            time.Time                      `json:"createdAt"`
	Description          string                         `json:"description"`
	DestinationAccountID string                         `json:"destinationAccountID"`
	Error                string                         `json:"error"`
	ID                   string                         `json:"id"`
	Metadata             map[string]string              `json:"metadata,omitempty"`
	Reference            string                         `json:"reference"`
	RelatedAdjustments   []TransferInitiationAdjusments `json:"relatedAdjustments,omitempty"`
	RelatedPayments      []TransferInitiationPayments   `json:"relatedPayments,omitempty"`
	ScheduledAt          time.Time                      `json:"scheduledAt"`
	SourceAccountID      string                         `json:"sourceAccountID"`
	Status               TransferInitiationStatus       `json:"status"`
	Type                 TransferInitiationType         `json:"type"`
}

func (t TransferInitiation) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TransferInitiation) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TransferInitiation) GetAmount() *big.Int {
	if o == nil {
		return big.NewInt(0)
	}
	return o.Amount
}

func (o *TransferInitiation) GetAsset() string {
	if o == nil {
		return ""
	}
	return o.Asset
}

func (o *TransferInitiation) GetConnectorID() string {
	if o == nil {
		return ""
	}
	return o.ConnectorID
}

func (o *TransferInitiation) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *TransferInitiation) GetDescription() string {
	if o == nil {
		return ""
	}
	return o.Description
}

func (o *TransferInitiation) GetDestinationAccountID() string {
	if o == nil {
		return ""
	}
	return o.DestinationAccountID
}

func (o *TransferInitiation) GetError() string {
	if o == nil {
		return ""
	}
	return o.Error
}

func (o *TransferInitiation) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *TransferInitiation) GetMetadata() map[string]string {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *TransferInitiation) GetReference() string {
	if o == nil {
		return ""
	}
	return o.Reference
}

func (o *TransferInitiation) GetRelatedAdjustments() []TransferInitiationAdjusments {
	if o == nil {
		return nil
	}
	return o.RelatedAdjustments
}

func (o *TransferInitiation) GetRelatedPayments() []TransferInitiationPayments {
	if o == nil {
		return nil
	}
	return o.RelatedPayments
}

func (o *TransferInitiation) GetScheduledAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.ScheduledAt
}

func (o *TransferInitiation) GetSourceAccountID() string {
	if o == nil {
		return ""
	}
	return o.SourceAccountID
}

func (o *TransferInitiation) GetStatus() TransferInitiationStatus {
	if o == nil {
		return TransferInitiationStatus("")
	}
	return o.Status
}

func (o *TransferInitiation) GetType() TransferInitiationType {
	if o == nil {
		return TransferInitiationType("")
	}
	return o.Type
}
