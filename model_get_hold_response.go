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

// GetHoldResponse struct for GetHoldResponse
type GetHoldResponse struct {
	Data ExpandedDebitHold `json:"data"`
}

// NewGetHoldResponse instantiates a new GetHoldResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetHoldResponse(data ExpandedDebitHold) *GetHoldResponse {
	this := GetHoldResponse{}
	this.Data = data
	return &this
}

// NewGetHoldResponseWithDefaults instantiates a new GetHoldResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetHoldResponseWithDefaults() *GetHoldResponse {
	this := GetHoldResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GetHoldResponse) GetData() ExpandedDebitHold {
	if o == nil {
		var ret ExpandedDebitHold
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetHoldResponse) GetDataOk() (*ExpandedDebitHold, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetHoldResponse) SetData(v ExpandedDebitHold) {
	o.Data = v
}

func (o GetHoldResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGetHoldResponse struct {
	value *GetHoldResponse
	isSet bool
}

func (v NullableGetHoldResponse) Get() *GetHoldResponse {
	return v.value
}

func (v *NullableGetHoldResponse) Set(val *GetHoldResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetHoldResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetHoldResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetHoldResponse(val *GetHoldResponse) *NullableGetHoldResponse {
	return &NullableGetHoldResponse{value: val, isSet: true}
}

func (v NullableGetHoldResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetHoldResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


