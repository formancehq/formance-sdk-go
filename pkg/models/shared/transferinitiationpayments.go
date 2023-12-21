// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/formancehq/formance-sdk-go/pkg/utils"
	"time"
)

type TransferInitiationPayments struct {
	CreatedAt time.Time     `json:"createdAt"`
	Error     string        `json:"error"`
	PaymentID string        `json:"paymentID"`
	Status    PaymentStatus `json:"status"`
}

func (t TransferInitiationPayments) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TransferInitiationPayments) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TransferInitiationPayments) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *TransferInitiationPayments) GetError() string {
	if o == nil {
		return ""
	}
	return o.Error
}

func (o *TransferInitiationPayments) GetPaymentID() string {
	if o == nil {
		return ""
	}
	return o.PaymentID
}

func (o *TransferInitiationPayments) GetStatus() PaymentStatus {
	if o == nil {
		return PaymentStatus("")
	}
	return o.Status
}
