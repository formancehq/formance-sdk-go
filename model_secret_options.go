/*
Formance Stack API

Open, modular foundation for unique payments flows  # Introduction This API is documented in **OpenAPI format**.  # Authentication Formance Stack offers one forms of authentication:   - OAuth2 OAuth2 - an open protocol to allow secure authorization in a simple and standard method from web, mobile and desktop applications. <SecurityDefinitions /> 

API version: v1.0.0-rc.3
Contact: support@formance.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package formance

import (
	"encoding/json"
)

// checks if the SecretOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecretOptions{}

// SecretOptions struct for SecretOptions
type SecretOptions struct {
	Name string `json:"name"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewSecretOptions instantiates a new SecretOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecretOptions(name string) *SecretOptions {
	this := SecretOptions{}
	this.Name = name
	return &this
}

// NewSecretOptionsWithDefaults instantiates a new SecretOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecretOptionsWithDefaults() *SecretOptions {
	this := SecretOptions{}
	return &this
}

// GetName returns the Name field value
func (o *SecretOptions) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SecretOptions) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SecretOptions) SetName(v string) {
	o.Name = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *SecretOptions) GetMetadata() map[string]interface{} {
	if o == nil || isNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecretOptions) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *SecretOptions) HasMetadata() bool {
	if o != nil && !isNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *SecretOptions) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o SecretOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecretOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !isNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

type NullableSecretOptions struct {
	value *SecretOptions
	isSet bool
}

func (v NullableSecretOptions) Get() *SecretOptions {
	return v.value
}

func (v *NullableSecretOptions) Set(val *SecretOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableSecretOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableSecretOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecretOptions(val *SecretOptions) *NullableSecretOptions {
	return &NullableSecretOptions{value: val, isSet: true}
}

func (v NullableSecretOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecretOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


