// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/formancehq/formance-sdk-go/v3/pkg/utils"
	"math/big"
	"time"
)

type OrchestrationPaymentRaw struct {
}

type OrchestrationPaymentScheme string

const (
	OrchestrationPaymentSchemeVisa       OrchestrationPaymentScheme = "visa"
	OrchestrationPaymentSchemeMastercard OrchestrationPaymentScheme = "mastercard"
	OrchestrationPaymentSchemeAmex       OrchestrationPaymentScheme = "amex"
	OrchestrationPaymentSchemeDiners     OrchestrationPaymentScheme = "diners"
	OrchestrationPaymentSchemeDiscover   OrchestrationPaymentScheme = "discover"
	OrchestrationPaymentSchemeJcb        OrchestrationPaymentScheme = "jcb"
	OrchestrationPaymentSchemeUnionpay   OrchestrationPaymentScheme = "unionpay"
	OrchestrationPaymentSchemeSepaDebit  OrchestrationPaymentScheme = "sepa debit"
	OrchestrationPaymentSchemeSepaCredit OrchestrationPaymentScheme = "sepa credit"
	OrchestrationPaymentSchemeSepa       OrchestrationPaymentScheme = "sepa"
	OrchestrationPaymentSchemeApplePay   OrchestrationPaymentScheme = "apple pay"
	OrchestrationPaymentSchemeGooglePay  OrchestrationPaymentScheme = "google pay"
	OrchestrationPaymentSchemeA2a        OrchestrationPaymentScheme = "a2a"
	OrchestrationPaymentSchemeAchDebit   OrchestrationPaymentScheme = "ach debit"
	OrchestrationPaymentSchemeAch        OrchestrationPaymentScheme = "ach"
	OrchestrationPaymentSchemeRtp        OrchestrationPaymentScheme = "rtp"
	OrchestrationPaymentSchemeUnknown    OrchestrationPaymentScheme = "unknown"
	OrchestrationPaymentSchemeOther      OrchestrationPaymentScheme = "other"
)

func (e OrchestrationPaymentScheme) ToPointer() *OrchestrationPaymentScheme {
	return &e
}
func (e *OrchestrationPaymentScheme) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "visa":
		fallthrough
	case "mastercard":
		fallthrough
	case "amex":
		fallthrough
	case "diners":
		fallthrough
	case "discover":
		fallthrough
	case "jcb":
		fallthrough
	case "unionpay":
		fallthrough
	case "sepa debit":
		fallthrough
	case "sepa credit":
		fallthrough
	case "sepa":
		fallthrough
	case "apple pay":
		fallthrough
	case "google pay":
		fallthrough
	case "a2a":
		fallthrough
	case "ach debit":
		fallthrough
	case "ach":
		fallthrough
	case "rtp":
		fallthrough
	case "unknown":
		fallthrough
	case "other":
		*e = OrchestrationPaymentScheme(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OrchestrationPaymentScheme: %v", v)
	}
}

type OrchestrationPaymentType string

const (
	OrchestrationPaymentTypePayIn    OrchestrationPaymentType = "PAY-IN"
	OrchestrationPaymentTypePayout   OrchestrationPaymentType = "PAYOUT"
	OrchestrationPaymentTypeTransfer OrchestrationPaymentType = "TRANSFER"
	OrchestrationPaymentTypeOther    OrchestrationPaymentType = "OTHER"
)

func (e OrchestrationPaymentType) ToPointer() *OrchestrationPaymentType {
	return &e
}
func (e *OrchestrationPaymentType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "PAY-IN":
		fallthrough
	case "PAYOUT":
		fallthrough
	case "TRANSFER":
		fallthrough
	case "OTHER":
		*e = OrchestrationPaymentType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OrchestrationPaymentType: %v", v)
	}
}

type OrchestrationPayment struct {
	Adjustments          []OrchestrationPaymentAdjustment `json:"adjustments"`
	Asset                string                           `json:"asset"`
	ConnectorID          string                           `json:"connectorID"`
	CreatedAt            time.Time                        `json:"createdAt"`
	DestinationAccountID string                           `json:"destinationAccountID"`
	ID                   string                           `json:"id"`
	InitialAmount        *big.Int                         `json:"initialAmount"`
	Metadata             *OrchestrationPaymentMetadata    `json:"metadata"`
	Provider             *OrchestrationConnector          `json:"provider,omitempty"`
	Raw                  *OrchestrationPaymentRaw         `json:"raw"`
	Reference            string                           `json:"reference"`
	Scheme               OrchestrationPaymentScheme       `json:"scheme"`
	SourceAccountID      string                           `json:"sourceAccountID"`
	Status               OrchestrationPaymentStatus       `json:"status"`
	Type                 OrchestrationPaymentType         `json:"type"`
}

func (o OrchestrationPayment) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(o, "", false)
}

func (o *OrchestrationPayment) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &o, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *OrchestrationPayment) GetAdjustments() []OrchestrationPaymentAdjustment {
	if o == nil {
		return []OrchestrationPaymentAdjustment{}
	}
	return o.Adjustments
}

func (o *OrchestrationPayment) GetAsset() string {
	if o == nil {
		return ""
	}
	return o.Asset
}

func (o *OrchestrationPayment) GetConnectorID() string {
	if o == nil {
		return ""
	}
	return o.ConnectorID
}

func (o *OrchestrationPayment) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *OrchestrationPayment) GetDestinationAccountID() string {
	if o == nil {
		return ""
	}
	return o.DestinationAccountID
}

func (o *OrchestrationPayment) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *OrchestrationPayment) GetInitialAmount() *big.Int {
	if o == nil {
		return big.NewInt(0)
	}
	return o.InitialAmount
}

func (o *OrchestrationPayment) GetMetadata() *OrchestrationPaymentMetadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *OrchestrationPayment) GetProvider() *OrchestrationConnector {
	if o == nil {
		return nil
	}
	return o.Provider
}

func (o *OrchestrationPayment) GetRaw() *OrchestrationPaymentRaw {
	if o == nil {
		return nil
	}
	return o.Raw
}

func (o *OrchestrationPayment) GetReference() string {
	if o == nil {
		return ""
	}
	return o.Reference
}

func (o *OrchestrationPayment) GetScheme() OrchestrationPaymentScheme {
	if o == nil {
		return OrchestrationPaymentScheme("")
	}
	return o.Scheme
}

func (o *OrchestrationPayment) GetSourceAccountID() string {
	if o == nil {
		return ""
	}
	return o.SourceAccountID
}

func (o *OrchestrationPayment) GetStatus() OrchestrationPaymentStatus {
	if o == nil {
		return OrchestrationPaymentStatus("")
	}
	return o.Status
}

func (o *OrchestrationPayment) GetType() OrchestrationPaymentType {
	if o == nil {
		return OrchestrationPaymentType("")
	}
	return o.Type
}
