// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/formancehq/formance-sdk-go/v2/v2/pkg/utils"
	"math/big"
	"time"
)

type Raw struct {
}

type Payment struct {
	Adjustments          []PaymentAdjustment `json:"adjustments"`
	Amount               *big.Int            `json:"amount"`
	Asset                string              `json:"asset"`
	ConnectorID          string              `json:"connectorID"`
	CreatedAt            time.Time           `json:"createdAt"`
	DestinationAccountID string              `json:"destinationAccountID"`
	ID                   string              `json:"id"`
	InitialAmount        *big.Int            `json:"initialAmount"`
	Metadata             map[string]string   `json:"metadata"`
	Provider             *Connector          `json:"provider,omitempty"`
	Raw                  *Raw                `json:"raw"`
	Reference            string              `json:"reference"`
	Scheme               PaymentScheme       `json:"scheme"`
	SourceAccountID      string              `json:"sourceAccountID"`
	Status               PaymentStatus       `json:"status"`
	Type                 PaymentType         `json:"type"`
}

func (p Payment) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *Payment) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Payment) GetAdjustments() []PaymentAdjustment {
	if o == nil {
		return []PaymentAdjustment{}
	}
	return o.Adjustments
}

func (o *Payment) GetAmount() *big.Int {
	if o == nil {
		return big.NewInt(0)
	}
	return o.Amount
}

func (o *Payment) GetAsset() string {
	if o == nil {
		return ""
	}
	return o.Asset
}

func (o *Payment) GetConnectorID() string {
	if o == nil {
		return ""
	}
	return o.ConnectorID
}

func (o *Payment) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *Payment) GetDestinationAccountID() string {
	if o == nil {
		return ""
	}
	return o.DestinationAccountID
}

func (o *Payment) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *Payment) GetInitialAmount() *big.Int {
	if o == nil {
		return big.NewInt(0)
	}
	return o.InitialAmount
}

func (o *Payment) GetMetadata() map[string]string {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *Payment) GetProvider() *Connector {
	if o == nil {
		return nil
	}
	return o.Provider
}

func (o *Payment) GetRaw() *Raw {
	if o == nil {
		return nil
	}
	return o.Raw
}

func (o *Payment) GetReference() string {
	if o == nil {
		return ""
	}
	return o.Reference
}

func (o *Payment) GetScheme() PaymentScheme {
	if o == nil {
		return PaymentScheme("")
	}
	return o.Scheme
}

func (o *Payment) GetSourceAccountID() string {
	if o == nil {
		return ""
	}
	return o.SourceAccountID
}

func (o *Payment) GetStatus() PaymentStatus {
	if o == nil {
		return PaymentStatus("")
	}
	return o.Status
}

func (o *Payment) GetType() PaymentType {
	if o == nil {
		return PaymentType("")
	}
	return o.Type
}
