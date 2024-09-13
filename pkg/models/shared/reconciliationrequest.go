// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/formancehq/formance-sdk-go/v2/pkg/utils"
	"time"
)

type ReconciliationRequest struct {
	ReconciledAtLedger   time.Time `json:"reconciledAtLedger"`
	ReconciledAtPayments time.Time `json:"reconciledAtPayments"`
}

func (r ReconciliationRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *ReconciliationRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ReconciliationRequest) GetReconciledAtLedger() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.ReconciledAtLedger
}

func (o *ReconciliationRequest) GetReconciledAtPayments() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.ReconciledAtPayments
}
