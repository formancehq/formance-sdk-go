/*
Formance Stack API

Open, modular foundation for unique payments flows  # Introduction This API is documented in **OpenAPI format**.  # Authentication Formance Stack offers one forms of authentication:   - OAuth2 OAuth2 - an open protocol to allow secure authorization in a simple and standard method from web, mobile and desktop applications. <SecurityDefinitions /> 

API version: v1.0.0-beta.4
Contact: support@formance.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package formance

import (
	"encoding/json"
)

// checks if the GetManyConfigs200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetManyConfigs200Response{}

// GetManyConfigs200Response struct for GetManyConfigs200Response
type GetManyConfigs200Response struct {
	Cursor GetManyConfigs200ResponseCursor `json:"cursor"`
}

// NewGetManyConfigs200Response instantiates a new GetManyConfigs200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetManyConfigs200Response(cursor GetManyConfigs200ResponseCursor) *GetManyConfigs200Response {
	this := GetManyConfigs200Response{}
	this.Cursor = cursor
	return &this
}

// NewGetManyConfigs200ResponseWithDefaults instantiates a new GetManyConfigs200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetManyConfigs200ResponseWithDefaults() *GetManyConfigs200Response {
	this := GetManyConfigs200Response{}
	return &this
}

// GetCursor returns the Cursor field value
func (o *GetManyConfigs200Response) GetCursor() GetManyConfigs200ResponseCursor {
	if o == nil {
		var ret GetManyConfigs200ResponseCursor
		return ret
	}

	return o.Cursor
}

// GetCursorOk returns a tuple with the Cursor field value
// and a boolean to check if the value has been set.
func (o *GetManyConfigs200Response) GetCursorOk() (*GetManyConfigs200ResponseCursor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cursor, true
}

// SetCursor sets field value
func (o *GetManyConfigs200Response) SetCursor(v GetManyConfigs200ResponseCursor) {
	o.Cursor = v
}

func (o GetManyConfigs200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetManyConfigs200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cursor"] = o.Cursor
	return toSerialize, nil
}

type NullableGetManyConfigs200Response struct {
	value *GetManyConfigs200Response
	isSet bool
}

func (v NullableGetManyConfigs200Response) Get() *GetManyConfigs200Response {
	return v.value
}

func (v *NullableGetManyConfigs200Response) Set(val *GetManyConfigs200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetManyConfigs200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetManyConfigs200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetManyConfigs200Response(val *GetManyConfigs200Response) *NullableGetManyConfigs200Response {
	return &NullableGetManyConfigs200Response{value: val, isSet: true}
}

func (v NullableGetManyConfigs200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetManyConfigs200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


