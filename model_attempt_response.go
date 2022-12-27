/*
Formance Stack API

Open, modular foundation for unique payments flows  # Introduction This API is documented in **OpenAPI format**.  # Authentication Formance Stack offers one forms of authentication:   - OAuth2 OAuth2 - an open protocol to allow secure authorization in a simple and standard method from web, mobile and desktop applications. <SecurityDefinitions /> 

API version: v1.0.0-rc.2
Contact: support@formance.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package formance

import (
	"encoding/json"
)

// checks if the AttemptResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AttemptResponse{}

// AttemptResponse struct for AttemptResponse
type AttemptResponse struct {
	Data *Attempt `json:"data,omitempty"`
}

// NewAttemptResponse instantiates a new AttemptResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttemptResponse() *AttemptResponse {
	this := AttemptResponse{}
	return &this
}

// NewAttemptResponseWithDefaults instantiates a new AttemptResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttemptResponseWithDefaults() *AttemptResponse {
	this := AttemptResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AttemptResponse) GetData() Attempt {
	if o == nil || isNil(o.Data) {
		var ret Attempt
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttemptResponse) GetDataOk() (*Attempt, bool) {
	if o == nil || isNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AttemptResponse) HasData() bool {
	if o != nil && !isNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given Attempt and assigns it to the Data field.
func (o *AttemptResponse) SetData(v Attempt) {
	o.Data = &v
}

func (o AttemptResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AttemptResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableAttemptResponse struct {
	value *AttemptResponse
	isSet bool
}

func (v NullableAttemptResponse) Get() *AttemptResponse {
	return v.value
}

func (v *NullableAttemptResponse) Set(val *AttemptResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAttemptResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAttemptResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttemptResponse(val *AttemptResponse) *NullableAttemptResponse {
	return &NullableAttemptResponse{value: val, isSet: true}
}

func (v NullableAttemptResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttemptResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


