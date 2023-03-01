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
)

// checks if the PaymentsCursor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentsCursor{}

// PaymentsCursor struct for PaymentsCursor
type PaymentsCursor struct {
	Cursor PaymentsCursorCursor `json:"cursor"`
}

// NewPaymentsCursor instantiates a new PaymentsCursor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentsCursor(cursor PaymentsCursorCursor) *PaymentsCursor {
	this := PaymentsCursor{}
	this.Cursor = cursor
	return &this
}

// NewPaymentsCursorWithDefaults instantiates a new PaymentsCursor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentsCursorWithDefaults() *PaymentsCursor {
	this := PaymentsCursor{}
	return &this
}

// GetCursor returns the Cursor field value
func (o *PaymentsCursor) GetCursor() PaymentsCursorCursor {
	if o == nil {
		var ret PaymentsCursorCursor
		return ret
	}

	return o.Cursor
}

// GetCursorOk returns a tuple with the Cursor field value
// and a boolean to check if the value has been set.
func (o *PaymentsCursor) GetCursorOk() (*PaymentsCursorCursor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cursor, true
}

// SetCursor sets field value
func (o *PaymentsCursor) SetCursor(v PaymentsCursorCursor) {
	o.Cursor = v
}

func (o PaymentsCursor) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentsCursor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cursor"] = o.Cursor
	return toSerialize, nil
}

type NullablePaymentsCursor struct {
	value *PaymentsCursor
	isSet bool
}

func (v NullablePaymentsCursor) Get() *PaymentsCursor {
	return v.value
}

func (v *NullablePaymentsCursor) Set(val *PaymentsCursor) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentsCursor) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentsCursor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentsCursor(val *PaymentsCursor) *NullablePaymentsCursor {
	return &NullablePaymentsCursor{value: val, isSet: true}
}

func (v NullablePaymentsCursor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentsCursor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

