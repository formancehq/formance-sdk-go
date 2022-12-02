/*
Formance Stack API

Open, modular foundation for unique payments flows  # Introduction This API is documented in **OpenAPI format**.  # Authentication Formance Stack offers one forms of authentication:   - OAuth2 OAuth2 - an open protocol to allow secure authorization in a simple and standard method from web, mobile and desktop applications. <SecurityDefinitions /> 

API version: v0.2.5
Contact: support@formance.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package formance

import (
	"encoding/json"
)

// checks if the GetAccount200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAccount200Response{}

// GetAccount200Response struct for GetAccount200Response
type GetAccount200Response struct {
	Data AccountWithVolumesAndBalances `json:"data"`
}

// NewGetAccount200Response instantiates a new GetAccount200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAccount200Response(data AccountWithVolumesAndBalances) *GetAccount200Response {
	this := GetAccount200Response{}
	this.Data = data
	return &this
}

// NewGetAccount200ResponseWithDefaults instantiates a new GetAccount200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAccount200ResponseWithDefaults() *GetAccount200Response {
	this := GetAccount200Response{}
	return &this
}

// GetData returns the Data field value
func (o *GetAccount200Response) GetData() AccountWithVolumesAndBalances {
	if o == nil {
		var ret AccountWithVolumesAndBalances
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetAccount200Response) GetDataOk() (*AccountWithVolumesAndBalances, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetAccount200Response) SetData(v AccountWithVolumesAndBalances) {
	o.Data = v
}

func (o GetAccount200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAccount200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableGetAccount200Response struct {
	value *GetAccount200Response
	isSet bool
}

func (v NullableGetAccount200Response) Get() *GetAccount200Response {
	return v.value
}

func (v *NullableGetAccount200Response) Set(val *GetAccount200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAccount200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAccount200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAccount200Response(val *GetAccount200Response) *NullableGetAccount200Response {
	return &NullableGetAccount200Response{value: val, isSet: true}
}

func (v NullableGetAccount200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAccount200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


