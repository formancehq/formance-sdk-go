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

// checks if the TasksCursor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TasksCursor{}

// TasksCursor struct for TasksCursor
type TasksCursor struct {
	Cursor TasksCursorCursor `json:"cursor"`
}

// NewTasksCursor instantiates a new TasksCursor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTasksCursor(cursor TasksCursorCursor) *TasksCursor {
	this := TasksCursor{}
	this.Cursor = cursor
	return &this
}

// NewTasksCursorWithDefaults instantiates a new TasksCursor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTasksCursorWithDefaults() *TasksCursor {
	this := TasksCursor{}
	return &this
}

// GetCursor returns the Cursor field value
func (o *TasksCursor) GetCursor() TasksCursorCursor {
	if o == nil {
		var ret TasksCursorCursor
		return ret
	}

	return o.Cursor
}

// GetCursorOk returns a tuple with the Cursor field value
// and a boolean to check if the value has been set.
func (o *TasksCursor) GetCursorOk() (*TasksCursorCursor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cursor, true
}

// SetCursor sets field value
func (o *TasksCursor) SetCursor(v TasksCursorCursor) {
	o.Cursor = v
}

func (o TasksCursor) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TasksCursor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cursor"] = o.Cursor
	return toSerialize, nil
}

type NullableTasksCursor struct {
	value *TasksCursor
	isSet bool
}

func (v NullableTasksCursor) Get() *TasksCursor {
	return v.value
}

func (v *NullableTasksCursor) Set(val *TasksCursor) {
	v.value = val
	v.isSet = true
}

func (v NullableTasksCursor) IsSet() bool {
	return v.isSet
}

func (v *NullableTasksCursor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTasksCursor(val *TasksCursor) *NullableTasksCursor {
	return &NullableTasksCursor{value: val, isSet: true}
}

func (v NullableTasksCursor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTasksCursor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


