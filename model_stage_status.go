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

// checks if the StageStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StageStatus{}

// StageStatus struct for StageStatus
type StageStatus struct {
	Stage float32 `json:"stage"`
	InstanceID string `json:"instanceID"`
	StartedAt time.Time `json:"startedAt"`
	TerminatedAt *time.Time `json:"terminatedAt,omitempty"`
	Error *string `json:"error,omitempty"`
}

// NewStageStatus instantiates a new StageStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStageStatus(stage float32, instanceID string, startedAt time.Time) *StageStatus {
	this := StageStatus{}
	this.Stage = stage
	this.InstanceID = instanceID
	this.StartedAt = startedAt
	return &this
}

// NewStageStatusWithDefaults instantiates a new StageStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStageStatusWithDefaults() *StageStatus {
	this := StageStatus{}
	return &this
}

// GetStage returns the Stage field value
func (o *StageStatus) GetStage() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Stage
}

// GetStageOk returns a tuple with the Stage field value
// and a boolean to check if the value has been set.
func (o *StageStatus) GetStageOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Stage, true
}

// SetStage sets field value
func (o *StageStatus) SetStage(v float32) {
	o.Stage = v
}

// GetInstanceID returns the InstanceID field value
func (o *StageStatus) GetInstanceID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstanceID
}

// GetInstanceIDOk returns a tuple with the InstanceID field value
// and a boolean to check if the value has been set.
func (o *StageStatus) GetInstanceIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceID, true
}

// SetInstanceID sets field value
func (o *StageStatus) SetInstanceID(v string) {
	o.InstanceID = v
}

// GetStartedAt returns the StartedAt field value
func (o *StageStatus) GetStartedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value
// and a boolean to check if the value has been set.
func (o *StageStatus) GetStartedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartedAt, true
}

// SetStartedAt sets field value
func (o *StageStatus) SetStartedAt(v time.Time) {
	o.StartedAt = v
}

// GetTerminatedAt returns the TerminatedAt field value if set, zero value otherwise.
func (o *StageStatus) GetTerminatedAt() time.Time {
	if o == nil || IsNil(o.TerminatedAt) {
		var ret time.Time
		return ret
	}
	return *o.TerminatedAt
}

// GetTerminatedAtOk returns a tuple with the TerminatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StageStatus) GetTerminatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.TerminatedAt) {
		return nil, false
	}
	return o.TerminatedAt, true
}

// HasTerminatedAt returns a boolean if a field has been set.
func (o *StageStatus) HasTerminatedAt() bool {
	if o != nil && !IsNil(o.TerminatedAt) {
		return true
	}

	return false
}

// SetTerminatedAt gets a reference to the given time.Time and assigns it to the TerminatedAt field.
func (o *StageStatus) SetTerminatedAt(v time.Time) {
	o.TerminatedAt = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *StageStatus) GetError() string {
	if o == nil || IsNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StageStatus) GetErrorOk() (*string, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *StageStatus) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *StageStatus) SetError(v string) {
	o.Error = &v
}

func (o StageStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StageStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["stage"] = o.Stage
	toSerialize["instanceID"] = o.InstanceID
	toSerialize["startedAt"] = o.StartedAt
	if !IsNil(o.TerminatedAt) {
		toSerialize["terminatedAt"] = o.TerminatedAt
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

type NullableStageStatus struct {
	value *StageStatus
	isSet bool
}

func (v NullableStageStatus) Get() *StageStatus {
	return v.value
}

func (v *NullableStageStatus) Set(val *StageStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableStageStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableStageStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStageStatus(val *StageStatus) *NullableStageStatus {
	return &NullableStageStatus{value: val, isSet: true}
}

func (v NullableStageStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStageStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

