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

// checks if the CreateBalanceResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateBalanceResponse{}

// CreateBalanceResponse struct for CreateBalanceResponse
type CreateBalanceResponse struct {
	Data Balance `json:"data"`
}

// NewCreateBalanceResponse instantiates a new CreateBalanceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateBalanceResponse(data Balance) *CreateBalanceResponse {
	this := CreateBalanceResponse{}
	this.Data = data
	return &this
}

// NewCreateBalanceResponseWithDefaults instantiates a new CreateBalanceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateBalanceResponseWithDefaults() *CreateBalanceResponse {
	this := CreateBalanceResponse{}
	return &this
}

// GetData returns the Data field value
func (o *CreateBalanceResponse) GetData() Balance {
	if o == nil {
		var ret Balance
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CreateBalanceResponse) GetDataOk() (*Balance, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CreateBalanceResponse) SetData(v Balance) {
	o.Data = v
}

func (o CreateBalanceResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateBalanceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableCreateBalanceResponse struct {
	value *CreateBalanceResponse
	isSet bool
}

func (v NullableCreateBalanceResponse) Get() *CreateBalanceResponse {
	return v.value
}

func (v *NullableCreateBalanceResponse) Set(val *CreateBalanceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateBalanceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateBalanceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateBalanceResponse(val *CreateBalanceResponse) *NullableCreateBalanceResponse {
	return &NullableCreateBalanceResponse{value: val, isSet: true}
}

func (v NullableCreateBalanceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateBalanceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


