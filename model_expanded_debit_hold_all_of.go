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
)

// ExpandedDebitHoldAllOf struct for ExpandedDebitHoldAllOf
type ExpandedDebitHoldAllOf struct {
	// Remaining amount on hold
	Remaining int64 `json:"remaining"`
	// Original amount on hold
	OriginalAmount int64 `json:"originalAmount"`
}

// NewExpandedDebitHoldAllOf instantiates a new ExpandedDebitHoldAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExpandedDebitHoldAllOf(remaining int64, originalAmount int64) *ExpandedDebitHoldAllOf {
	this := ExpandedDebitHoldAllOf{}
	this.Remaining = remaining
	this.OriginalAmount = originalAmount
	return &this
}

// NewExpandedDebitHoldAllOfWithDefaults instantiates a new ExpandedDebitHoldAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExpandedDebitHoldAllOfWithDefaults() *ExpandedDebitHoldAllOf {
	this := ExpandedDebitHoldAllOf{}
	return &this
}

// GetRemaining returns the Remaining field value
func (o *ExpandedDebitHoldAllOf) GetRemaining() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Remaining
}

// GetRemainingOk returns a tuple with the Remaining field value
// and a boolean to check if the value has been set.
func (o *ExpandedDebitHoldAllOf) GetRemainingOk() (*int64, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Remaining, true
}

// SetRemaining sets field value
func (o *ExpandedDebitHoldAllOf) SetRemaining(v int64) {
	o.Remaining = v
}

// GetOriginalAmount returns the OriginalAmount field value
func (o *ExpandedDebitHoldAllOf) GetOriginalAmount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.OriginalAmount
}

// GetOriginalAmountOk returns a tuple with the OriginalAmount field value
// and a boolean to check if the value has been set.
func (o *ExpandedDebitHoldAllOf) GetOriginalAmountOk() (*int64, bool) {
	if o == nil {
    return nil, false
	}
	return &o.OriginalAmount, true
}

// SetOriginalAmount sets field value
func (o *ExpandedDebitHoldAllOf) SetOriginalAmount(v int64) {
	o.OriginalAmount = v
}

func (o ExpandedDebitHoldAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["remaining"] = o.Remaining
	}
	if true {
		toSerialize["originalAmount"] = o.OriginalAmount
	}
	return json.Marshal(toSerialize)
}

type NullableExpandedDebitHoldAllOf struct {
	value *ExpandedDebitHoldAllOf
	isSet bool
}

func (v NullableExpandedDebitHoldAllOf) Get() *ExpandedDebitHoldAllOf {
	return v.value
}

func (v *NullableExpandedDebitHoldAllOf) Set(val *ExpandedDebitHoldAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableExpandedDebitHoldAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableExpandedDebitHoldAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExpandedDebitHoldAllOf(val *ExpandedDebitHoldAllOf) *NullableExpandedDebitHoldAllOf {
	return &NullableExpandedDebitHoldAllOf{value: val, isSet: true}
}

func (v NullableExpandedDebitHoldAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExpandedDebitHoldAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


