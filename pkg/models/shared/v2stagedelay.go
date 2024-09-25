// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/formancehq/formance-sdk-go/v3/pkg/utils"
	"time"
)

type V2StageDelay struct {
	Duration *string    `json:"duration,omitempty"`
	Until    *time.Time `json:"until,omitempty"`
}

func (v V2StageDelay) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(v, "", false)
}

func (v *V2StageDelay) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &v, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *V2StageDelay) GetDuration() *string {
	if o == nil {
		return nil
	}
	return o.Duration
}

func (o *V2StageDelay) GetUntil() *time.Time {
	if o == nil {
		return nil
	}
	return o.Until
}
