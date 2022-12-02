/*
Formance Stack API

Open, modular foundation for unique payments flows  # Introduction This API is documented in **OpenAPI format**.  # Authentication Formance Stack offers one forms of authentication:   - OAuth2 OAuth2 - an open protocol to allow secure authorization in a simple and standard method from web, mobile and desktop applications. <SecurityDefinitions /> 

API version: v0.2.7
Contact: support@formance.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package formance

import (
	"encoding/json"
	"fmt"
)

// ConnectorConfig - struct for ConnectorConfig
type ConnectorConfig struct {
	BankingCircleConfig *BankingCircleConfig
	CurrencyCloudConfig *CurrencyCloudConfig
	DummyPayConfig *DummyPayConfig
	ModulrConfig *ModulrConfig
	StripeConfig *StripeConfig
	WiseConfig *WiseConfig
}

// BankingCircleConfigAsConnectorConfig is a convenience function that returns BankingCircleConfig wrapped in ConnectorConfig
func BankingCircleConfigAsConnectorConfig(v *BankingCircleConfig) ConnectorConfig {
	return ConnectorConfig{
		BankingCircleConfig: v,
	}
}

// CurrencyCloudConfigAsConnectorConfig is a convenience function that returns CurrencyCloudConfig wrapped in ConnectorConfig
func CurrencyCloudConfigAsConnectorConfig(v *CurrencyCloudConfig) ConnectorConfig {
	return ConnectorConfig{
		CurrencyCloudConfig: v,
	}
}

// DummyPayConfigAsConnectorConfig is a convenience function that returns DummyPayConfig wrapped in ConnectorConfig
func DummyPayConfigAsConnectorConfig(v *DummyPayConfig) ConnectorConfig {
	return ConnectorConfig{
		DummyPayConfig: v,
	}
}

// ModulrConfigAsConnectorConfig is a convenience function that returns ModulrConfig wrapped in ConnectorConfig
func ModulrConfigAsConnectorConfig(v *ModulrConfig) ConnectorConfig {
	return ConnectorConfig{
		ModulrConfig: v,
	}
}

// StripeConfigAsConnectorConfig is a convenience function that returns StripeConfig wrapped in ConnectorConfig
func StripeConfigAsConnectorConfig(v *StripeConfig) ConnectorConfig {
	return ConnectorConfig{
		StripeConfig: v,
	}
}

// WiseConfigAsConnectorConfig is a convenience function that returns WiseConfig wrapped in ConnectorConfig
func WiseConfigAsConnectorConfig(v *WiseConfig) ConnectorConfig {
	return ConnectorConfig{
		WiseConfig: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ConnectorConfig) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into BankingCircleConfig
	err = newStrictDecoder(data).Decode(&dst.BankingCircleConfig)
	if err == nil {
		jsonBankingCircleConfig, _ := json.Marshal(dst.BankingCircleConfig)
		if string(jsonBankingCircleConfig) == "{}" { // empty struct
			dst.BankingCircleConfig = nil
		} else {
			match++
		}
	} else {
		dst.BankingCircleConfig = nil
	}

	// try to unmarshal data into CurrencyCloudConfig
	err = newStrictDecoder(data).Decode(&dst.CurrencyCloudConfig)
	if err == nil {
		jsonCurrencyCloudConfig, _ := json.Marshal(dst.CurrencyCloudConfig)
		if string(jsonCurrencyCloudConfig) == "{}" { // empty struct
			dst.CurrencyCloudConfig = nil
		} else {
			match++
		}
	} else {
		dst.CurrencyCloudConfig = nil
	}

	// try to unmarshal data into DummyPayConfig
	err = newStrictDecoder(data).Decode(&dst.DummyPayConfig)
	if err == nil {
		jsonDummyPayConfig, _ := json.Marshal(dst.DummyPayConfig)
		if string(jsonDummyPayConfig) == "{}" { // empty struct
			dst.DummyPayConfig = nil
		} else {
			match++
		}
	} else {
		dst.DummyPayConfig = nil
	}

	// try to unmarshal data into ModulrConfig
	err = newStrictDecoder(data).Decode(&dst.ModulrConfig)
	if err == nil {
		jsonModulrConfig, _ := json.Marshal(dst.ModulrConfig)
		if string(jsonModulrConfig) == "{}" { // empty struct
			dst.ModulrConfig = nil
		} else {
			match++
		}
	} else {
		dst.ModulrConfig = nil
	}

	// try to unmarshal data into StripeConfig
	err = newStrictDecoder(data).Decode(&dst.StripeConfig)
	if err == nil {
		jsonStripeConfig, _ := json.Marshal(dst.StripeConfig)
		if string(jsonStripeConfig) == "{}" { // empty struct
			dst.StripeConfig = nil
		} else {
			match++
		}
	} else {
		dst.StripeConfig = nil
	}

	// try to unmarshal data into WiseConfig
	err = newStrictDecoder(data).Decode(&dst.WiseConfig)
	if err == nil {
		jsonWiseConfig, _ := json.Marshal(dst.WiseConfig)
		if string(jsonWiseConfig) == "{}" { // empty struct
			dst.WiseConfig = nil
		} else {
			match++
		}
	} else {
		dst.WiseConfig = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.BankingCircleConfig = nil
		dst.CurrencyCloudConfig = nil
		dst.DummyPayConfig = nil
		dst.ModulrConfig = nil
		dst.StripeConfig = nil
		dst.WiseConfig = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ConnectorConfig)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ConnectorConfig)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ConnectorConfig) MarshalJSON() ([]byte, error) {
	if src.BankingCircleConfig != nil {
		return json.Marshal(&src.BankingCircleConfig)
	}

	if src.CurrencyCloudConfig != nil {
		return json.Marshal(&src.CurrencyCloudConfig)
	}

	if src.DummyPayConfig != nil {
		return json.Marshal(&src.DummyPayConfig)
	}

	if src.ModulrConfig != nil {
		return json.Marshal(&src.ModulrConfig)
	}

	if src.StripeConfig != nil {
		return json.Marshal(&src.StripeConfig)
	}

	if src.WiseConfig != nil {
		return json.Marshal(&src.WiseConfig)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ConnectorConfig) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.BankingCircleConfig != nil {
		return obj.BankingCircleConfig
	}

	if obj.CurrencyCloudConfig != nil {
		return obj.CurrencyCloudConfig
	}

	if obj.DummyPayConfig != nil {
		return obj.DummyPayConfig
	}

	if obj.ModulrConfig != nil {
		return obj.ModulrConfig
	}

	if obj.StripeConfig != nil {
		return obj.StripeConfig
	}

	if obj.WiseConfig != nil {
		return obj.WiseConfig
	}

	// all schemas are nil
	return nil
}

type NullableConnectorConfig struct {
	value *ConnectorConfig
	isSet bool
}

func (v NullableConnectorConfig) Get() *ConnectorConfig {
	return v.value
}

func (v *NullableConnectorConfig) Set(val *ConnectorConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorConfig(val *ConnectorConfig) *NullableConnectorConfig {
	return &NullableConnectorConfig{value: val, isSet: true}
}

func (v NullableConnectorConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


