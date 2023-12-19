// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/formancehq/formance-sdk-go/pkg/utils"
	"math/big"
	"time"
)

type Balance struct {
	ExpiresAt *time.Time `json:"expiresAt,omitempty"`
	Name      string     `json:"name"`
	Priority  *big.Int   `json:"priority,omitempty"`
}

func (b Balance) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(b, "", false)
}

func (b *Balance) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &b, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Balance) GetExpiresAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.ExpiresAt
}

func (o *Balance) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Balance) GetPriority() *big.Int {
	if o == nil {
		return nil
	}
	return o.Priority
}
