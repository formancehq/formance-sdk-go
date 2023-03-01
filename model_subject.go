/*
Formance Stack API

Open, modular foundation for unique payments flows  # Introduction This API is documented in **OpenAPI format**.  # Authentication Formance Stack offers one forms of authentication:   - OAuth2 OAuth2 - an open protocol to allow secure authorization in a simple and standard method from web, mobile and desktop applications. <SecurityDefinitions /> 

API version: v1.0.20230228
Contact: support@formance.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package formance

import (
	"encoding/json"
	"fmt"
)

// Subject struct for Subject
type Subject struct {
	LedgerAccountSubject *LedgerAccountSubject
	WalletSubject *WalletSubject
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *Subject) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = json.Unmarshal(data, &jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'ACCOUNT'
	if jsonDict["type"] == "ACCOUNT" {
		// try to unmarshal JSON data into LedgerAccountSubject
		err = json.Unmarshal(data, &dst.LedgerAccountSubject);
		if err == nil {
			jsonLedgerAccountSubject, _ := json.Marshal(dst.LedgerAccountSubject)
			if string(jsonLedgerAccountSubject) == "{}" { // empty struct
				dst.LedgerAccountSubject = nil
			} else {
				return nil // data stored in dst.LedgerAccountSubject, return on the first match
			}
		} else {
			dst.LedgerAccountSubject = nil
		}
	}

	// check if the discriminator value is 'LedgerAccountSubject'
	if jsonDict["type"] == "LedgerAccountSubject" {
		// try to unmarshal JSON data into LedgerAccountSubject
		err = json.Unmarshal(data, &dst.LedgerAccountSubject);
		if err == nil {
			jsonLedgerAccountSubject, _ := json.Marshal(dst.LedgerAccountSubject)
			if string(jsonLedgerAccountSubject) == "{}" { // empty struct
				dst.LedgerAccountSubject = nil
			} else {
				return nil // data stored in dst.LedgerAccountSubject, return on the first match
			}
		} else {
			dst.LedgerAccountSubject = nil
		}
	}

	// check if the discriminator value is 'WALLET'
	if jsonDict["type"] == "WALLET" {
		// try to unmarshal JSON data into WalletSubject
		err = json.Unmarshal(data, &dst.WalletSubject);
		if err == nil {
			jsonWalletSubject, _ := json.Marshal(dst.WalletSubject)
			if string(jsonWalletSubject) == "{}" { // empty struct
				dst.WalletSubject = nil
			} else {
				return nil // data stored in dst.WalletSubject, return on the first match
			}
		} else {
			dst.WalletSubject = nil
		}
	}

	// check if the discriminator value is 'WalletSubject'
	if jsonDict["type"] == "WalletSubject" {
		// try to unmarshal JSON data into WalletSubject
		err = json.Unmarshal(data, &dst.WalletSubject);
		if err == nil {
			jsonWalletSubject, _ := json.Marshal(dst.WalletSubject)
			if string(jsonWalletSubject) == "{}" { // empty struct
				dst.WalletSubject = nil
			} else {
				return nil // data stored in dst.WalletSubject, return on the first match
			}
		} else {
			dst.WalletSubject = nil
		}
	}

	// try to unmarshal JSON data into LedgerAccountSubject
	err = json.Unmarshal(data, &dst.LedgerAccountSubject);
	if err == nil {
		jsonLedgerAccountSubject, _ := json.Marshal(dst.LedgerAccountSubject)
		if string(jsonLedgerAccountSubject) == "{}" { // empty struct
			dst.LedgerAccountSubject = nil
		} else {
			return nil // data stored in dst.LedgerAccountSubject, return on the first match
		}
	} else {
		dst.LedgerAccountSubject = nil
	}

	// try to unmarshal JSON data into WalletSubject
	err = json.Unmarshal(data, &dst.WalletSubject);
	if err == nil {
		jsonWalletSubject, _ := json.Marshal(dst.WalletSubject)
		if string(jsonWalletSubject) == "{}" { // empty struct
			dst.WalletSubject = nil
		} else {
			return nil // data stored in dst.WalletSubject, return on the first match
		}
	} else {
		dst.WalletSubject = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(Subject)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *Subject) MarshalJSON() ([]byte, error) {
	if src.LedgerAccountSubject != nil {
		return json.Marshal(&src.LedgerAccountSubject)
	}

	if src.WalletSubject != nil {
		return json.Marshal(&src.WalletSubject)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableSubject struct {
	value *Subject
	isSet bool
}

func (v NullableSubject) Get() *Subject {
	return v.value
}

func (v *NullableSubject) Set(val *Subject) {
	v.value = val
	v.isSet = true
}

func (v NullableSubject) IsSet() bool {
	return v.isSet
}

func (v *NullableSubject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubject(val *Subject) *NullableSubject {
	return &NullableSubject{value: val, isSet: true}
}

func (v NullableSubject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


