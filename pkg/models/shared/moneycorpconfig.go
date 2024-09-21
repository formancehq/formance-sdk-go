// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/formancehq/formance-sdk-go/v2/pkg/utils"
)

type MoneycorpConfig struct {
	APIKey   string `json:"apiKey"`
	ClientID string `json:"clientID"`
	Endpoint string `json:"endpoint"`
	Name     string `json:"name"`
	// The frequency at which the connector will try to fetch new BalanceTransaction objects from MoneyCorp API.
	//
	PollingPeriod *string `default:"120s" json:"pollingPeriod"`
}

func (m MoneycorpConfig) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(m, "", false)
}

func (m *MoneycorpConfig) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &m, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *MoneycorpConfig) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

func (o *MoneycorpConfig) GetClientID() string {
	if o == nil {
		return ""
	}
	return o.ClientID
}

func (o *MoneycorpConfig) GetEndpoint() string {
	if o == nil {
		return ""
	}
	return o.Endpoint
}

func (o *MoneycorpConfig) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *MoneycorpConfig) GetPollingPeriod() *string {
	if o == nil {
		return nil
	}
	return o.PollingPeriod
}
