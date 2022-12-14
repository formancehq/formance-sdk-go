/*
Formance Stack API

Open, modular foundation for unique payments flows  # Introduction This API is documented in **OpenAPI format**.  # Authentication Formance Stack offers one forms of authentication:   - OAuth2 OAuth2 - an open protocol to allow secure authorization in a simple and standard method from web, mobile and desktop applications. <SecurityDefinitions /> 

API version: v1.0.0-rc.1
Contact: support@formance.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package formance

import (
	"encoding/json"
)

// checks if the TaskDescriptorWiseAllOfDescriptor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaskDescriptorWiseAllOfDescriptor{}

// TaskDescriptorWiseAllOfDescriptor struct for TaskDescriptorWiseAllOfDescriptor
type TaskDescriptorWiseAllOfDescriptor struct {
	Name *string `json:"name,omitempty"`
	Key *string `json:"key,omitempty"`
	ProfileID *int32 `json:"profileID,omitempty"`
}

// NewTaskDescriptorWiseAllOfDescriptor instantiates a new TaskDescriptorWiseAllOfDescriptor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskDescriptorWiseAllOfDescriptor() *TaskDescriptorWiseAllOfDescriptor {
	this := TaskDescriptorWiseAllOfDescriptor{}
	return &this
}

// NewTaskDescriptorWiseAllOfDescriptorWithDefaults instantiates a new TaskDescriptorWiseAllOfDescriptor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskDescriptorWiseAllOfDescriptorWithDefaults() *TaskDescriptorWiseAllOfDescriptor {
	this := TaskDescriptorWiseAllOfDescriptor{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TaskDescriptorWiseAllOfDescriptor) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskDescriptorWiseAllOfDescriptor) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TaskDescriptorWiseAllOfDescriptor) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TaskDescriptorWiseAllOfDescriptor) SetName(v string) {
	o.Name = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *TaskDescriptorWiseAllOfDescriptor) GetKey() string {
	if o == nil || isNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskDescriptorWiseAllOfDescriptor) GetKeyOk() (*string, bool) {
	if o == nil || isNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *TaskDescriptorWiseAllOfDescriptor) HasKey() bool {
	if o != nil && !isNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *TaskDescriptorWiseAllOfDescriptor) SetKey(v string) {
	o.Key = &v
}

// GetProfileID returns the ProfileID field value if set, zero value otherwise.
func (o *TaskDescriptorWiseAllOfDescriptor) GetProfileID() int32 {
	if o == nil || isNil(o.ProfileID) {
		var ret int32
		return ret
	}
	return *o.ProfileID
}

// GetProfileIDOk returns a tuple with the ProfileID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskDescriptorWiseAllOfDescriptor) GetProfileIDOk() (*int32, bool) {
	if o == nil || isNil(o.ProfileID) {
		return nil, false
	}
	return o.ProfileID, true
}

// HasProfileID returns a boolean if a field has been set.
func (o *TaskDescriptorWiseAllOfDescriptor) HasProfileID() bool {
	if o != nil && !isNil(o.ProfileID) {
		return true
	}

	return false
}

// SetProfileID gets a reference to the given int32 and assigns it to the ProfileID field.
func (o *TaskDescriptorWiseAllOfDescriptor) SetProfileID(v int32) {
	o.ProfileID = &v
}

func (o TaskDescriptorWiseAllOfDescriptor) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaskDescriptorWiseAllOfDescriptor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !isNil(o.ProfileID) {
		toSerialize["profileID"] = o.ProfileID
	}
	return toSerialize, nil
}

type NullableTaskDescriptorWiseAllOfDescriptor struct {
	value *TaskDescriptorWiseAllOfDescriptor
	isSet bool
}

func (v NullableTaskDescriptorWiseAllOfDescriptor) Get() *TaskDescriptorWiseAllOfDescriptor {
	return v.value
}

func (v *NullableTaskDescriptorWiseAllOfDescriptor) Set(val *TaskDescriptorWiseAllOfDescriptor) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskDescriptorWiseAllOfDescriptor) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskDescriptorWiseAllOfDescriptor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskDescriptorWiseAllOfDescriptor(val *TaskDescriptorWiseAllOfDescriptor) *NullableTaskDescriptorWiseAllOfDescriptor {
	return &NullableTaskDescriptorWiseAllOfDescriptor{value: val, isSet: true}
}

func (v NullableTaskDescriptorWiseAllOfDescriptor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskDescriptorWiseAllOfDescriptor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


