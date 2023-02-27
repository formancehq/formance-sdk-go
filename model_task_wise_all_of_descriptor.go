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

// checks if the TaskWiseAllOfDescriptor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaskWiseAllOfDescriptor{}

// TaskWiseAllOfDescriptor struct for TaskWiseAllOfDescriptor
type TaskWiseAllOfDescriptor struct {
	Name *string `json:"name,omitempty"`
	Key *string `json:"key,omitempty"`
	ProfileID *int64 `json:"profileID,omitempty"`
}

// NewTaskWiseAllOfDescriptor instantiates a new TaskWiseAllOfDescriptor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskWiseAllOfDescriptor() *TaskWiseAllOfDescriptor {
	this := TaskWiseAllOfDescriptor{}
	return &this
}

// NewTaskWiseAllOfDescriptorWithDefaults instantiates a new TaskWiseAllOfDescriptor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskWiseAllOfDescriptorWithDefaults() *TaskWiseAllOfDescriptor {
	this := TaskWiseAllOfDescriptor{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TaskWiseAllOfDescriptor) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskWiseAllOfDescriptor) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TaskWiseAllOfDescriptor) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TaskWiseAllOfDescriptor) SetName(v string) {
	o.Name = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *TaskWiseAllOfDescriptor) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskWiseAllOfDescriptor) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *TaskWiseAllOfDescriptor) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *TaskWiseAllOfDescriptor) SetKey(v string) {
	o.Key = &v
}

// GetProfileID returns the ProfileID field value if set, zero value otherwise.
func (o *TaskWiseAllOfDescriptor) GetProfileID() int64 {
	if o == nil || IsNil(o.ProfileID) {
		var ret int64
		return ret
	}
	return *o.ProfileID
}

// GetProfileIDOk returns a tuple with the ProfileID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskWiseAllOfDescriptor) GetProfileIDOk() (*int64, bool) {
	if o == nil || IsNil(o.ProfileID) {
		return nil, false
	}
	return o.ProfileID, true
}

// HasProfileID returns a boolean if a field has been set.
func (o *TaskWiseAllOfDescriptor) HasProfileID() bool {
	if o != nil && !IsNil(o.ProfileID) {
		return true
	}

	return false
}

// SetProfileID gets a reference to the given int64 and assigns it to the ProfileID field.
func (o *TaskWiseAllOfDescriptor) SetProfileID(v int64) {
	o.ProfileID = &v
}

func (o TaskWiseAllOfDescriptor) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaskWiseAllOfDescriptor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.ProfileID) {
		toSerialize["profileID"] = o.ProfileID
	}
	return toSerialize, nil
}

type NullableTaskWiseAllOfDescriptor struct {
	value *TaskWiseAllOfDescriptor
	isSet bool
}

func (v NullableTaskWiseAllOfDescriptor) Get() *TaskWiseAllOfDescriptor {
	return v.value
}

func (v *NullableTaskWiseAllOfDescriptor) Set(val *TaskWiseAllOfDescriptor) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskWiseAllOfDescriptor) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskWiseAllOfDescriptor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskWiseAllOfDescriptor(val *TaskWiseAllOfDescriptor) *NullableTaskWiseAllOfDescriptor {
	return &NullableTaskWiseAllOfDescriptor{value: val, isSet: true}
}

func (v NullableTaskWiseAllOfDescriptor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskWiseAllOfDescriptor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


