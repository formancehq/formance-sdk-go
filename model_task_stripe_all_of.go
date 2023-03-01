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

// checks if the TaskStripeAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaskStripeAllOf{}

// TaskStripeAllOf struct for TaskStripeAllOf
type TaskStripeAllOf struct {
	Descriptor TaskStripeAllOfDescriptor `json:"descriptor"`
}

// NewTaskStripeAllOf instantiates a new TaskStripeAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskStripeAllOf(descriptor TaskStripeAllOfDescriptor) *TaskStripeAllOf {
	this := TaskStripeAllOf{}
	this.Descriptor = descriptor
	return &this
}

// NewTaskStripeAllOfWithDefaults instantiates a new TaskStripeAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskStripeAllOfWithDefaults() *TaskStripeAllOf {
	this := TaskStripeAllOf{}
	return &this
}

// GetDescriptor returns the Descriptor field value
func (o *TaskStripeAllOf) GetDescriptor() TaskStripeAllOfDescriptor {
	if o == nil {
		var ret TaskStripeAllOfDescriptor
		return ret
	}

	return o.Descriptor
}

// GetDescriptorOk returns a tuple with the Descriptor field value
// and a boolean to check if the value has been set.
func (o *TaskStripeAllOf) GetDescriptorOk() (*TaskStripeAllOfDescriptor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Descriptor, true
}

// SetDescriptor sets field value
func (o *TaskStripeAllOf) SetDescriptor(v TaskStripeAllOfDescriptor) {
	o.Descriptor = v
}

func (o TaskStripeAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaskStripeAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["descriptor"] = o.Descriptor
	return toSerialize, nil
}

type NullableTaskStripeAllOf struct {
	value *TaskStripeAllOf
	isSet bool
}

func (v NullableTaskStripeAllOf) Get() *TaskStripeAllOf {
	return v.value
}

func (v *NullableTaskStripeAllOf) Set(val *TaskStripeAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskStripeAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskStripeAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskStripeAllOf(val *TaskStripeAllOf) *NullableTaskStripeAllOf {
	return &NullableTaskStripeAllOf{value: val, isSet: true}
}

func (v NullableTaskStripeAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskStripeAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

