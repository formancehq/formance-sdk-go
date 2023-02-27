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

// checks if the AccountsCursor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountsCursor{}

// AccountsCursor struct for AccountsCursor
type AccountsCursor struct {
	Cursor AccountsCursorCursor `json:"cursor"`
}

// NewAccountsCursor instantiates a new AccountsCursor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountsCursor(cursor AccountsCursorCursor) *AccountsCursor {
	this := AccountsCursor{}
	this.Cursor = cursor
	return &this
}

// NewAccountsCursorWithDefaults instantiates a new AccountsCursor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountsCursorWithDefaults() *AccountsCursor {
	this := AccountsCursor{}
	return &this
}

// GetCursor returns the Cursor field value
func (o *AccountsCursor) GetCursor() AccountsCursorCursor {
	if o == nil {
		var ret AccountsCursorCursor
		return ret
	}

	return o.Cursor
}

// GetCursorOk returns a tuple with the Cursor field value
// and a boolean to check if the value has been set.
func (o *AccountsCursor) GetCursorOk() (*AccountsCursorCursor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cursor, true
}

// SetCursor sets field value
func (o *AccountsCursor) SetCursor(v AccountsCursorCursor) {
	o.Cursor = v
}

func (o AccountsCursor) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountsCursor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cursor"] = o.Cursor
	return toSerialize, nil
}

type NullableAccountsCursor struct {
	value *AccountsCursor
	isSet bool
}

func (v NullableAccountsCursor) Get() *AccountsCursor {
	return v.value
}

func (v *NullableAccountsCursor) Set(val *AccountsCursor) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountsCursor) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountsCursor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountsCursor(val *AccountsCursor) *NullableAccountsCursor {
	return &NullableAccountsCursor{value: val, isSet: true}
}

func (v NullableAccountsCursor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountsCursor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


