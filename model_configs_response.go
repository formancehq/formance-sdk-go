/*
Formance Stack API

Open, modular foundation for unique payments flows  # Introduction This API is documented in **OpenAPI format**.  # Authentication Formance Stack offers one forms of authentication:   - OAuth2 OAuth2 - an open protocol to allow secure authorization in a simple and standard method from web, mobile and desktop applications. <SecurityDefinitions /> 

API version: v1.0.0-rc.5
Contact: support@formance.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package formance

import (
	"encoding/json"
)

// checks if the ConfigsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigsResponse{}

// ConfigsResponse struct for ConfigsResponse
type ConfigsResponse struct {
	Cursor WebhooksCursor `json:"cursor"`
}

// NewConfigsResponse instantiates a new ConfigsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigsResponse(cursor WebhooksCursor) *ConfigsResponse {
	this := ConfigsResponse{}
	this.Cursor = cursor
	return &this
}

// NewConfigsResponseWithDefaults instantiates a new ConfigsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigsResponseWithDefaults() *ConfigsResponse {
	this := ConfigsResponse{}
	return &this
}

// GetCursor returns the Cursor field value
func (o *ConfigsResponse) GetCursor() WebhooksCursor {
	if o == nil {
		var ret WebhooksCursor
		return ret
	}

	return o.Cursor
}

// GetCursorOk returns a tuple with the Cursor field value
// and a boolean to check if the value has been set.
func (o *ConfigsResponse) GetCursorOk() (*WebhooksCursor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cursor, true
}

// SetCursor sets field value
func (o *ConfigsResponse) SetCursor(v WebhooksCursor) {
	o.Cursor = v
}

func (o ConfigsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cursor"] = o.Cursor
	return toSerialize, nil
}

type NullableConfigsResponse struct {
	value *ConfigsResponse
	isSet bool
}

func (v NullableConfigsResponse) Get() *ConfigsResponse {
	return v.value
}

func (v *NullableConfigsResponse) Set(val *ConfigsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigsResponse(val *ConfigsResponse) *NullableConfigsResponse {
	return &NullableConfigsResponse{value: val, isSet: true}
}

func (v NullableConfigsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


