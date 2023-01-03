/*
Formance Stack API

Open, modular foundation for unique payments flows  # Introduction This API is documented in **OpenAPI format**.  # Authentication Formance Stack offers one forms of authentication:   - OAuth2 OAuth2 - an open protocol to allow secure authorization in a simple and standard method from web, mobile and desktop applications. <SecurityDefinitions /> 

API version: v1.0.0-rc.4
Contact: support@formance.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package formance

import (
	"encoding/json"
)

// checks if the ConfigChangeSecret type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigChangeSecret{}

// ConfigChangeSecret struct for ConfigChangeSecret
type ConfigChangeSecret struct {
	Secret *string `json:"secret,omitempty"`
}

// NewConfigChangeSecret instantiates a new ConfigChangeSecret object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigChangeSecret() *ConfigChangeSecret {
	this := ConfigChangeSecret{}
	return &this
}

// NewConfigChangeSecretWithDefaults instantiates a new ConfigChangeSecret object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigChangeSecretWithDefaults() *ConfigChangeSecret {
	this := ConfigChangeSecret{}
	return &this
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *ConfigChangeSecret) GetSecret() string {
	if o == nil || isNil(o.Secret) {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigChangeSecret) GetSecretOk() (*string, bool) {
	if o == nil || isNil(o.Secret) {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *ConfigChangeSecret) HasSecret() bool {
	if o != nil && !isNil(o.Secret) {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *ConfigChangeSecret) SetSecret(v string) {
	o.Secret = &v
}

func (o ConfigChangeSecret) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigChangeSecret) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Secret) {
		toSerialize["secret"] = o.Secret
	}
	return toSerialize, nil
}

type NullableConfigChangeSecret struct {
	value *ConfigChangeSecret
	isSet bool
}

func (v NullableConfigChangeSecret) Get() *ConfigChangeSecret {
	return v.value
}

func (v *NullableConfigChangeSecret) Set(val *ConfigChangeSecret) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigChangeSecret) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigChangeSecret) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigChangeSecret(val *ConfigChangeSecret) *NullableConfigChangeSecret {
	return &NullableConfigChangeSecret{value: val, isSet: true}
}

func (v NullableConfigChangeSecret) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigChangeSecret) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


