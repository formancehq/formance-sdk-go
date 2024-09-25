// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/formancehq/formance-sdk-go/v3/pkg/utils"
)

type DummyPayConfig struct {
	Directory string `json:"directory"`
	// The frequency at which the connector will try to fetch new payment objects from the directory
	FilePollingPeriod            *string `default:"10s" json:"filePollingPeriod"`
	Name                         string  `json:"name"`
	NumberOfAccountsPreGenerated *int64  `json:"numberOfAccountsPreGenerated,omitempty"`
	NumberOfPaymentsPreGenerated *int64  `json:"numberOfPaymentsPreGenerated,omitempty"`
	PrefixFileToIngest           *string `json:"prefixFileToIngest,omitempty"`
}

func (d DummyPayConfig) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DummyPayConfig) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DummyPayConfig) GetDirectory() string {
	if o == nil {
		return ""
	}
	return o.Directory
}

func (o *DummyPayConfig) GetFilePollingPeriod() *string {
	if o == nil {
		return nil
	}
	return o.FilePollingPeriod
}

func (o *DummyPayConfig) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *DummyPayConfig) GetNumberOfAccountsPreGenerated() *int64 {
	if o == nil {
		return nil
	}
	return o.NumberOfAccountsPreGenerated
}

func (o *DummyPayConfig) GetNumberOfPaymentsPreGenerated() *int64 {
	if o == nil {
		return nil
	}
	return o.NumberOfPaymentsPreGenerated
}

func (o *DummyPayConfig) GetPrefixFileToIngest() *string {
	if o == nil {
		return nil
	}
	return o.PrefixFileToIngest
}
