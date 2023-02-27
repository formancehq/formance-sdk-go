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

// checks if the TaskModulrAllOfDescriptor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaskModulrAllOfDescriptor{}

// TaskModulrAllOfDescriptor struct for TaskModulrAllOfDescriptor
type TaskModulrAllOfDescriptor struct {
	Name *string `json:"name,omitempty"`
	Key *string `json:"key,omitempty"`
	AccountID *string `json:"accountID,omitempty"`
}

// NewTaskModulrAllOfDescriptor instantiates a new TaskModulrAllOfDescriptor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskModulrAllOfDescriptor() *TaskModulrAllOfDescriptor {
	this := TaskModulrAllOfDescriptor{}
	return &this
}

// NewTaskModulrAllOfDescriptorWithDefaults instantiates a new TaskModulrAllOfDescriptor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskModulrAllOfDescriptorWithDefaults() *TaskModulrAllOfDescriptor {
	this := TaskModulrAllOfDescriptor{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TaskModulrAllOfDescriptor) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskModulrAllOfDescriptor) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TaskModulrAllOfDescriptor) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TaskModulrAllOfDescriptor) SetName(v string) {
	o.Name = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *TaskModulrAllOfDescriptor) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskModulrAllOfDescriptor) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *TaskModulrAllOfDescriptor) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *TaskModulrAllOfDescriptor) SetKey(v string) {
	o.Key = &v
}

// GetAccountID returns the AccountID field value if set, zero value otherwise.
func (o *TaskModulrAllOfDescriptor) GetAccountID() string {
	if o == nil || IsNil(o.AccountID) {
		var ret string
		return ret
	}
	return *o.AccountID
}

// GetAccountIDOk returns a tuple with the AccountID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskModulrAllOfDescriptor) GetAccountIDOk() (*string, bool) {
	if o == nil || IsNil(o.AccountID) {
		return nil, false
	}
	return o.AccountID, true
}

// HasAccountID returns a boolean if a field has been set.
func (o *TaskModulrAllOfDescriptor) HasAccountID() bool {
	if o != nil && !IsNil(o.AccountID) {
		return true
	}

	return false
}

// SetAccountID gets a reference to the given string and assigns it to the AccountID field.
func (o *TaskModulrAllOfDescriptor) SetAccountID(v string) {
	o.AccountID = &v
}

func (o TaskModulrAllOfDescriptor) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaskModulrAllOfDescriptor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.AccountID) {
		toSerialize["accountID"] = o.AccountID
	}
	return toSerialize, nil
}

type NullableTaskModulrAllOfDescriptor struct {
	value *TaskModulrAllOfDescriptor
	isSet bool
}

func (v NullableTaskModulrAllOfDescriptor) Get() *TaskModulrAllOfDescriptor {
	return v.value
}

func (v *NullableTaskModulrAllOfDescriptor) Set(val *TaskModulrAllOfDescriptor) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskModulrAllOfDescriptor) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskModulrAllOfDescriptor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskModulrAllOfDescriptor(val *TaskModulrAllOfDescriptor) *NullableTaskModulrAllOfDescriptor {
	return &NullableTaskModulrAllOfDescriptor{value: val, isSet: true}
}

func (v NullableTaskModulrAllOfDescriptor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskModulrAllOfDescriptor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


