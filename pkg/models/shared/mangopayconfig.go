// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/formancehq/formance-sdk-go/v3/pkg/utils"
)

type MangoPayConfig struct {
	APIKey   string `json:"apiKey"`
	ClientID string `json:"clientID"`
	Endpoint string `json:"endpoint"`
	Name     string `json:"name"`
	// The frequency at which the connector will try to fetch new BalanceTransaction objects from MangoPay API.
	//
	PollingPeriod *string `default:"120s" json:"pollingPeriod"`
}

func (m MangoPayConfig) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(m, "", false)
}

func (m *MangoPayConfig) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &m, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *MangoPayConfig) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

func (o *MangoPayConfig) GetClientID() string {
	if o == nil {
		return ""
	}
	return o.ClientID
}

func (o *MangoPayConfig) GetEndpoint() string {
	if o == nil {
		return ""
	}
	return o.Endpoint
}

func (o *MangoPayConfig) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *MangoPayConfig) GetPollingPeriod() *string {
	if o == nil {
		return nil
	}
	return o.PollingPeriod
}
