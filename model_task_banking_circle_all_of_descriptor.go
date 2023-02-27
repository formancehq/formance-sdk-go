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

// checks if the TaskBankingCircleAllOfDescriptor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaskBankingCircleAllOfDescriptor{}

// TaskBankingCircleAllOfDescriptor struct for TaskBankingCircleAllOfDescriptor
type TaskBankingCircleAllOfDescriptor struct {
	Name *string `json:"name,omitempty"`
	Key *string `json:"key,omitempty"`
}

// NewTaskBankingCircleAllOfDescriptor instantiates a new TaskBankingCircleAllOfDescriptor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskBankingCircleAllOfDescriptor() *TaskBankingCircleAllOfDescriptor {
	this := TaskBankingCircleAllOfDescriptor{}
	return &this
}

// NewTaskBankingCircleAllOfDescriptorWithDefaults instantiates a new TaskBankingCircleAllOfDescriptor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskBankingCircleAllOfDescriptorWithDefaults() *TaskBankingCircleAllOfDescriptor {
	this := TaskBankingCircleAllOfDescriptor{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TaskBankingCircleAllOfDescriptor) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskBankingCircleAllOfDescriptor) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TaskBankingCircleAllOfDescriptor) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TaskBankingCircleAllOfDescriptor) SetName(v string) {
	o.Name = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *TaskBankingCircleAllOfDescriptor) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskBankingCircleAllOfDescriptor) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *TaskBankingCircleAllOfDescriptor) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *TaskBankingCircleAllOfDescriptor) SetKey(v string) {
	o.Key = &v
}

func (o TaskBankingCircleAllOfDescriptor) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaskBankingCircleAllOfDescriptor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	return toSerialize, nil
}

type NullableTaskBankingCircleAllOfDescriptor struct {
	value *TaskBankingCircleAllOfDescriptor
	isSet bool
}

func (v NullableTaskBankingCircleAllOfDescriptor) Get() *TaskBankingCircleAllOfDescriptor {
	return v.value
}

func (v *NullableTaskBankingCircleAllOfDescriptor) Set(val *TaskBankingCircleAllOfDescriptor) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskBankingCircleAllOfDescriptor) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskBankingCircleAllOfDescriptor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskBankingCircleAllOfDescriptor(val *TaskBankingCircleAllOfDescriptor) *NullableTaskBankingCircleAllOfDescriptor {
	return &NullableTaskBankingCircleAllOfDescriptor{value: val, isSet: true}
}

func (v NullableTaskBankingCircleAllOfDescriptor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskBankingCircleAllOfDescriptor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


