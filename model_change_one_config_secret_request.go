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

// checks if the ChangeOneConfigSecretRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChangeOneConfigSecretRequest{}

// ChangeOneConfigSecretRequest struct for ChangeOneConfigSecretRequest
type ChangeOneConfigSecretRequest struct {
	Secret string `json:"secret"`
}

// NewChangeOneConfigSecretRequest instantiates a new ChangeOneConfigSecretRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChangeOneConfigSecretRequest(secret string) *ChangeOneConfigSecretRequest {
	this := ChangeOneConfigSecretRequest{}
	this.Secret = secret
	return &this
}

// NewChangeOneConfigSecretRequestWithDefaults instantiates a new ChangeOneConfigSecretRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChangeOneConfigSecretRequestWithDefaults() *ChangeOneConfigSecretRequest {
	this := ChangeOneConfigSecretRequest{}
	return &this
}

// GetSecret returns the Secret field value
func (o *ChangeOneConfigSecretRequest) GetSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Secret
}

// GetSecretOk returns a tuple with the Secret field value
// and a boolean to check if the value has been set.
func (o *ChangeOneConfigSecretRequest) GetSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Secret, true
}

// SetSecret sets field value
func (o *ChangeOneConfigSecretRequest) SetSecret(v string) {
	o.Secret = v
}

func (o ChangeOneConfigSecretRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChangeOneConfigSecretRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["secret"] = o.Secret
	return toSerialize, nil
}

type NullableChangeOneConfigSecretRequest struct {
	value *ChangeOneConfigSecretRequest
	isSet bool
}

func (v NullableChangeOneConfigSecretRequest) Get() *ChangeOneConfigSecretRequest {
	return v.value
}

func (v *NullableChangeOneConfigSecretRequest) Set(val *ChangeOneConfigSecretRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableChangeOneConfigSecretRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableChangeOneConfigSecretRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChangeOneConfigSecretRequest(val *ChangeOneConfigSecretRequest) *NullableChangeOneConfigSecretRequest {
	return &NullableChangeOneConfigSecretRequest{value: val, isSet: true}
}

func (v NullableChangeOneConfigSecretRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChangeOneConfigSecretRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


