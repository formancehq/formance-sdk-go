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

// checks if the TasksCursorCursorAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TasksCursorCursorAllOf{}

// TasksCursorCursorAllOf struct for TasksCursorCursorAllOf
type TasksCursorCursorAllOf struct {
	Data []TasksCursorCursorAllOfDataInner `json:"data"`
}

// NewTasksCursorCursorAllOf instantiates a new TasksCursorCursorAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTasksCursorCursorAllOf(data []TasksCursorCursorAllOfDataInner) *TasksCursorCursorAllOf {
	this := TasksCursorCursorAllOf{}
	this.Data = data
	return &this
}

// NewTasksCursorCursorAllOfWithDefaults instantiates a new TasksCursorCursorAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTasksCursorCursorAllOfWithDefaults() *TasksCursorCursorAllOf {
	this := TasksCursorCursorAllOf{}
	return &this
}

// GetData returns the Data field value
func (o *TasksCursorCursorAllOf) GetData() []TasksCursorCursorAllOfDataInner {
	if o == nil {
		var ret []TasksCursorCursorAllOfDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *TasksCursorCursorAllOf) GetDataOk() ([]TasksCursorCursorAllOfDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *TasksCursorCursorAllOf) SetData(v []TasksCursorCursorAllOfDataInner) {
	o.Data = v
}

func (o TasksCursorCursorAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TasksCursorCursorAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableTasksCursorCursorAllOf struct {
	value *TasksCursorCursorAllOf
	isSet bool
}

func (v NullableTasksCursorCursorAllOf) Get() *TasksCursorCursorAllOf {
	return v.value
}

func (v *NullableTasksCursorCursorAllOf) Set(val *TasksCursorCursorAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTasksCursorCursorAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTasksCursorCursorAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTasksCursorCursorAllOf(val *TasksCursorCursorAllOf) *NullableTasksCursorCursorAllOf {
	return &NullableTasksCursorCursorAllOf{value: val, isSet: true}
}

func (v NullableTasksCursorCursorAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTasksCursorCursorAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


