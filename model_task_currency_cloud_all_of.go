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

// checks if the TaskCurrencyCloudAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaskCurrencyCloudAllOf{}

// TaskCurrencyCloudAllOf struct for TaskCurrencyCloudAllOf
type TaskCurrencyCloudAllOf struct {
	Descriptor TaskCurrencyCloudAllOfDescriptor `json:"descriptor"`
}

// NewTaskCurrencyCloudAllOf instantiates a new TaskCurrencyCloudAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskCurrencyCloudAllOf(descriptor TaskCurrencyCloudAllOfDescriptor) *TaskCurrencyCloudAllOf {
	this := TaskCurrencyCloudAllOf{}
	this.Descriptor = descriptor
	return &this
}

// NewTaskCurrencyCloudAllOfWithDefaults instantiates a new TaskCurrencyCloudAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskCurrencyCloudAllOfWithDefaults() *TaskCurrencyCloudAllOf {
	this := TaskCurrencyCloudAllOf{}
	return &this
}

// GetDescriptor returns the Descriptor field value
func (o *TaskCurrencyCloudAllOf) GetDescriptor() TaskCurrencyCloudAllOfDescriptor {
	if o == nil {
		var ret TaskCurrencyCloudAllOfDescriptor
		return ret
	}

	return o.Descriptor
}

// GetDescriptorOk returns a tuple with the Descriptor field value
// and a boolean to check if the value has been set.
func (o *TaskCurrencyCloudAllOf) GetDescriptorOk() (*TaskCurrencyCloudAllOfDescriptor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Descriptor, true
}

// SetDescriptor sets field value
func (o *TaskCurrencyCloudAllOf) SetDescriptor(v TaskCurrencyCloudAllOfDescriptor) {
	o.Descriptor = v
}

func (o TaskCurrencyCloudAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaskCurrencyCloudAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["descriptor"] = o.Descriptor
	return toSerialize, nil
}

type NullableTaskCurrencyCloudAllOf struct {
	value *TaskCurrencyCloudAllOf
	isSet bool
}

func (v NullableTaskCurrencyCloudAllOf) Get() *TaskCurrencyCloudAllOf {
	return v.value
}

func (v *NullableTaskCurrencyCloudAllOf) Set(val *TaskCurrencyCloudAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskCurrencyCloudAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskCurrencyCloudAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskCurrencyCloudAllOf(val *TaskCurrencyCloudAllOf) *NullableTaskCurrencyCloudAllOf {
	return &NullableTaskCurrencyCloudAllOf{value: val, isSet: true}
}

func (v NullableTaskCurrencyCloudAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskCurrencyCloudAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


