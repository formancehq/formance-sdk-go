// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ModulrConfig struct {
	APIKey    string  `json:"apiKey"`
	APISecret string  `json:"apiSecret"`
	Endpoint  *string `json:"endpoint,omitempty"`
	// The frequency at which the connector will try to fetch new BalanceTransaction objects from Stripe API.
	//
	PollingPeriod *string `json:"pollingPeriod,omitempty"`
}
