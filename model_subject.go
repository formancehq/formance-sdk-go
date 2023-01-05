/*
Formance Stack API

Open, modular foundation for unique payments flows  # Introduction This API is documented in **OpenAPI format**.  # Authentication Formance Stack offers one forms of authentication:   - OAuth2 OAuth2 - an open protocol to allow secure authorization in a simple and standard method from web, mobile and desktop applications. <SecurityDefinitions /> 

API version: develop
Contact: support@formance.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package formance

import (
	"encoding/json"
	"fmt"
)

// Subject - struct for Subject
type Subject struct {
	LedgerAccountSubject *LedgerAccountSubject
	WalletSubject *WalletSubject
}

// LedgerAccountSubjectAsSubject is a convenience function that returns LedgerAccountSubject wrapped in Subject
func LedgerAccountSubjectAsSubject(v *LedgerAccountSubject) Subject {
	return Subject{
		LedgerAccountSubject: v,
	}
}

// WalletSubjectAsSubject is a convenience function that returns WalletSubject wrapped in Subject
func WalletSubjectAsSubject(v *WalletSubject) Subject {
	return Subject{
		WalletSubject: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *Subject) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into LedgerAccountSubject
	err = newStrictDecoder(data).Decode(&dst.LedgerAccountSubject)
	if err == nil {
		jsonLedgerAccountSubject, _ := json.Marshal(dst.LedgerAccountSubject)
		if string(jsonLedgerAccountSubject) == "{}" { // empty struct
			dst.LedgerAccountSubject = nil
		} else {
			match++
		}
	} else {
		dst.LedgerAccountSubject = nil
	}

	// try to unmarshal data into WalletSubject
	err = newStrictDecoder(data).Decode(&dst.WalletSubject)
	if err == nil {
		jsonWalletSubject, _ := json.Marshal(dst.WalletSubject)
		if string(jsonWalletSubject) == "{}" { // empty struct
			dst.WalletSubject = nil
		} else {
			match++
		}
	} else {
		dst.WalletSubject = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.LedgerAccountSubject = nil
		dst.WalletSubject = nil

		return fmt.Errorf("data matches more than one schema in oneOf(Subject)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(Subject)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src Subject) MarshalJSON() ([]byte, error) {
	if src.LedgerAccountSubject != nil {
		return json.Marshal(&src.LedgerAccountSubject)
	}

	if src.WalletSubject != nil {
		return json.Marshal(&src.WalletSubject)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *Subject) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.LedgerAccountSubject != nil {
		return obj.LedgerAccountSubject
	}

	if obj.WalletSubject != nil {
		return obj.WalletSubject
	}

	// all schemas are nil
	return nil
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

