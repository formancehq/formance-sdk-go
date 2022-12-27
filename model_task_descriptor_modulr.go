/*
Formance Stack API

Open, modular foundation for unique payments flows  # Introduction This API is documented in **OpenAPI format**.  # Authentication Formance Stack offers one forms of authentication:   - OAuth2 OAuth2 - an open protocol to allow secure authorization in a simple and standard method from web, mobile and desktop applications. <SecurityDefinitions /> 

API version: v1.0.0-rc.3
Contact: support@formance.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package formance

import (
	"encoding/json"
	"time"
)

// checks if the TaskDescriptorModulr type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaskDescriptorModulr{}

// TaskDescriptorModulr struct for TaskDescriptorModulr
type TaskDescriptorModulr struct {
	// The connector code
	Provider *string `json:"provider,omitempty"`
	// The date when the task was created
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// The task status
	Status *string `json:"status,omitempty"`
	// The error message if the task failed
	Error *string `json:"error,omitempty"`
	// The task state
	State map[string]interface{} `json:"state,omitempty"`
	Descriptor *TaskDescriptorModulrDescriptor `json:"descriptor,omitempty"`
}

// NewTaskDescriptorModulr instantiates a new TaskDescriptorModulr object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskDescriptorModulr() *TaskDescriptorModulr {
	this := TaskDescriptorModulr{}
	return &this
}

// NewTaskDescriptorModulrWithDefaults instantiates a new TaskDescriptorModulr object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskDescriptorModulrWithDefaults() *TaskDescriptorModulr {
	this := TaskDescriptorModulr{}
	return &this
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *TaskDescriptorModulr) GetProvider() string {
	if o == nil || isNil(o.Provider) {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskDescriptorModulr) GetProviderOk() (*string, bool) {
	if o == nil || isNil(o.Provider) {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *TaskDescriptorModulr) HasProvider() bool {
	if o != nil && !isNil(o.Provider) {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *TaskDescriptorModulr) SetProvider(v string) {
	o.Provider = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *TaskDescriptorModulr) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskDescriptorModulr) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *TaskDescriptorModulr) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *TaskDescriptorModulr) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *TaskDescriptorModulr) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskDescriptorModulr) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *TaskDescriptorModulr) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *TaskDescriptorModulr) SetStatus(v string) {
	o.Status = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *TaskDescriptorModulr) GetError() string {
	if o == nil || isNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskDescriptorModulr) GetErrorOk() (*string, bool) {
	if o == nil || isNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *TaskDescriptorModulr) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *TaskDescriptorModulr) SetError(v string) {
	o.Error = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *TaskDescriptorModulr) GetState() map[string]interface{} {
	if o == nil || isNil(o.State) {
		var ret map[string]interface{}
		return ret
	}
	return o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskDescriptorModulr) GetStateOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.State) {
		return map[string]interface{}{}, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *TaskDescriptorModulr) HasState() bool {
	if o != nil && !isNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given map[string]interface{} and assigns it to the State field.
func (o *TaskDescriptorModulr) SetState(v map[string]interface{}) {
	o.State = v
}

// GetDescriptor returns the Descriptor field value if set, zero value otherwise.
func (o *TaskDescriptorModulr) GetDescriptor() TaskDescriptorModulrDescriptor {
	if o == nil || isNil(o.Descriptor) {
		var ret TaskDescriptorModulrDescriptor
		return ret
	}
	return *o.Descriptor
}

// GetDescriptorOk returns a tuple with the Descriptor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskDescriptorModulr) GetDescriptorOk() (*TaskDescriptorModulrDescriptor, bool) {
	if o == nil || isNil(o.Descriptor) {
		return nil, false
	}
	return o.Descriptor, true
}

// HasDescriptor returns a boolean if a field has been set.
func (o *TaskDescriptorModulr) HasDescriptor() bool {
	if o != nil && !isNil(o.Descriptor) {
		return true
	}

	return false
}

// SetDescriptor gets a reference to the given TaskDescriptorModulrDescriptor and assigns it to the Descriptor field.
func (o *TaskDescriptorModulr) SetDescriptor(v TaskDescriptorModulrDescriptor) {
	o.Descriptor = &v
}

func (o TaskDescriptorModulr) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaskDescriptorModulr) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Provider) {
		toSerialize["provider"] = o.Provider
	}
	if !isNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !isNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !isNil(o.Descriptor) {
		toSerialize["descriptor"] = o.Descriptor
	}
	return toSerialize, nil
}

type NullableTaskDescriptorModulr struct {
	value *TaskDescriptorModulr
	isSet bool
}

func (v NullableTaskDescriptorModulr) Get() *TaskDescriptorModulr {
	return v.value
}

func (v *NullableTaskDescriptorModulr) Set(val *TaskDescriptorModulr) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskDescriptorModulr) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskDescriptorModulr) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskDescriptorModulr(val *TaskDescriptorModulr) *NullableTaskDescriptorModulr {
	return &NullableTaskDescriptorModulr{value: val, isSet: true}
}

func (v NullableTaskDescriptorModulr) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskDescriptorModulr) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


