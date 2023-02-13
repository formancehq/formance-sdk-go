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

// checks if the SecretAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecretAllOf{}

// SecretAllOf struct for SecretAllOf
type SecretAllOf struct {
	Id string `json:"id"`
	LastDigits string `json:"lastDigits"`
	Clear string `json:"clear"`
}

// NewSecretAllOf instantiates a new SecretAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecretAllOf(id string, lastDigits string, clear string) *SecretAllOf {
	this := SecretAllOf{}
	this.Id = id
	this.LastDigits = lastDigits
	this.Clear = clear
	return &this
}

// NewSecretAllOfWithDefaults instantiates a new SecretAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecretAllOfWithDefaults() *SecretAllOf {
	this := SecretAllOf{}
	return &this
}

// GetId returns the Id field value
func (o *SecretAllOf) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SecretAllOf) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SecretAllOf) SetId(v string) {
	o.Id = v
}

// GetLastDigits returns the LastDigits field value
func (o *SecretAllOf) GetLastDigits() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastDigits
}

// GetLastDigitsOk returns a tuple with the LastDigits field value
// and a boolean to check if the value has been set.
func (o *SecretAllOf) GetLastDigitsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastDigits, true
}

// SetLastDigits sets field value
func (o *SecretAllOf) SetLastDigits(v string) {
	o.LastDigits = v
}

// GetClear returns the Clear field value
func (o *SecretAllOf) GetClear() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Clear
}

// GetClearOk returns a tuple with the Clear field value
// and a boolean to check if the value has been set.
func (o *SecretAllOf) GetClearOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Clear, true
}

// SetClear sets field value
func (o *SecretAllOf) SetClear(v string) {
	o.Clear = v
}

func (o SecretAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecretAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["lastDigits"] = o.LastDigits
	toSerialize["clear"] = o.Clear
	return toSerialize, nil
}

type NullableSecretAllOf struct {
	value *SecretAllOf
	isSet bool
}

func (v NullableSecretAllOf) Get() *SecretAllOf {
	return v.value
}

func (v *NullableSecretAllOf) Set(val *SecretAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSecretAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSecretAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecretAllOf(val *SecretAllOf) *NullableSecretAllOf {
	return &NullableSecretAllOf{value: val, isSet: true}
}

func (v NullableSecretAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecretAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


