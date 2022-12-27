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

// checks if the GetManyConfigs200ResponseCursorAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetManyConfigs200ResponseCursorAllOf{}

// GetManyConfigs200ResponseCursorAllOf struct for GetManyConfigs200ResponseCursorAllOf
type GetManyConfigs200ResponseCursorAllOf struct {
	Data []WebhooksConfig `json:"data"`
}

// NewGetManyConfigs200ResponseCursorAllOf instantiates a new GetManyConfigs200ResponseCursorAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetManyConfigs200ResponseCursorAllOf(data []WebhooksConfig) *GetManyConfigs200ResponseCursorAllOf {
	this := GetManyConfigs200ResponseCursorAllOf{}
	this.Data = data
	return &this
}

// NewGetManyConfigs200ResponseCursorAllOfWithDefaults instantiates a new GetManyConfigs200ResponseCursorAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetManyConfigs200ResponseCursorAllOfWithDefaults() *GetManyConfigs200ResponseCursorAllOf {
	this := GetManyConfigs200ResponseCursorAllOf{}
	return &this
}

// GetData returns the Data field value
func (o *GetManyConfigs200ResponseCursorAllOf) GetData() []WebhooksConfig {
	if o == nil {
		var ret []WebhooksConfig
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetManyConfigs200ResponseCursorAllOf) GetDataOk() ([]WebhooksConfig, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *GetManyConfigs200ResponseCursorAllOf) SetData(v []WebhooksConfig) {
	o.Data = v
}

func (o GetManyConfigs200ResponseCursorAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetManyConfigs200ResponseCursorAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableGetManyConfigs200ResponseCursorAllOf struct {
	value *GetManyConfigs200ResponseCursorAllOf
	isSet bool
}

func (v NullableGetManyConfigs200ResponseCursorAllOf) Get() *GetManyConfigs200ResponseCursorAllOf {
	return v.value
}

func (v *NullableGetManyConfigs200ResponseCursorAllOf) Set(val *GetManyConfigs200ResponseCursorAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableGetManyConfigs200ResponseCursorAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableGetManyConfigs200ResponseCursorAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetManyConfigs200ResponseCursorAllOf(val *GetManyConfigs200ResponseCursorAllOf) *NullableGetManyConfigs200ResponseCursorAllOf {
	return &NullableGetManyConfigs200ResponseCursorAllOf{value: val, isSet: true}
}

func (v NullableGetManyConfigs200ResponseCursorAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetManyConfigs200ResponseCursorAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


