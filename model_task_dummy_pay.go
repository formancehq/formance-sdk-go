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
	"time"
)

// checks if the TaskDummyPay type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaskDummyPay{}

// TaskDummyPay struct for TaskDummyPay
type TaskDummyPay struct {
	Id string `json:"id"`
	ConnectorId string `json:"connectorId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Status PaymentStatus `json:"status"`
	State map[string]interface{} `json:"state"`
	Error *string `json:"error,omitempty"`
	Descriptor TaskDummyPayAllOfDescriptor `json:"descriptor"`
}

// NewTaskDummyPay instantiates a new TaskDummyPay object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskDummyPay(id string, connectorId string, createdAt time.Time, updatedAt time.Time, status PaymentStatus, state map[string]interface{}, descriptor TaskDummyPayAllOfDescriptor) *TaskDummyPay {
	this := TaskDummyPay{}
	this.Id = id
	this.ConnectorId = connectorId
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.Status = status
	this.State = state
	this.Descriptor = descriptor
	return &this
}

// NewTaskDummyPayWithDefaults instantiates a new TaskDummyPay object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskDummyPayWithDefaults() *TaskDummyPay {
	this := TaskDummyPay{}
	return &this
}

// GetId returns the Id field value
func (o *TaskDummyPay) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TaskDummyPay) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TaskDummyPay) SetId(v string) {
	o.Id = v
}

// GetConnectorId returns the ConnectorId field value
func (o *TaskDummyPay) GetConnectorId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectorId
}

// GetConnectorIdOk returns a tuple with the ConnectorId field value
// and a boolean to check if the value has been set.
func (o *TaskDummyPay) GetConnectorIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectorId, true
}

// SetConnectorId sets field value
func (o *TaskDummyPay) SetConnectorId(v string) {
	o.ConnectorId = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *TaskDummyPay) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *TaskDummyPay) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *TaskDummyPay) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *TaskDummyPay) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *TaskDummyPay) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *TaskDummyPay) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetStatus returns the Status field value
func (o *TaskDummyPay) GetStatus() PaymentStatus {
	if o == nil {
		var ret PaymentStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *TaskDummyPay) GetStatusOk() (*PaymentStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *TaskDummyPay) SetStatus(v PaymentStatus) {
	o.Status = v
}

// GetState returns the State field value
func (o *TaskDummyPay) GetState() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *TaskDummyPay) GetStateOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.State, true
}

// SetState sets field value
func (o *TaskDummyPay) SetState(v map[string]interface{}) {
	o.State = v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *TaskDummyPay) GetError() string {
	if o == nil || IsNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskDummyPay) GetErrorOk() (*string, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *TaskDummyPay) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *TaskDummyPay) SetError(v string) {
	o.Error = &v
}

// GetDescriptor returns the Descriptor field value
func (o *TaskDummyPay) GetDescriptor() TaskDummyPayAllOfDescriptor {
	if o == nil {
		var ret TaskDummyPayAllOfDescriptor
		return ret
	}

	return o.Descriptor
}

// GetDescriptorOk returns a tuple with the Descriptor field value
// and a boolean to check if the value has been set.
func (o *TaskDummyPay) GetDescriptorOk() (*TaskDummyPayAllOfDescriptor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Descriptor, true
}

// SetDescriptor sets field value
func (o *TaskDummyPay) SetDescriptor(v TaskDummyPayAllOfDescriptor) {
	o.Descriptor = v
}

func (o TaskDummyPay) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaskDummyPay) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["connectorId"] = o.ConnectorId
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["updatedAt"] = o.UpdatedAt
	toSerialize["status"] = o.Status
	toSerialize["state"] = o.State
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	toSerialize["descriptor"] = o.Descriptor
	return toSerialize, nil
}

type NullableTaskDummyPay struct {
	value *TaskDummyPay
	isSet bool
}

func (v NullableTaskDummyPay) Get() *TaskDummyPay {
	return v.value
}

func (v *NullableTaskDummyPay) Set(val *TaskDummyPay) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskDummyPay) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskDummyPay) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskDummyPay(val *TaskDummyPay) *NullableTaskDummyPay {
	return &NullableTaskDummyPay{value: val, isSet: true}
}

func (v NullableTaskDummyPay) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskDummyPay) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


