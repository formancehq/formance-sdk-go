/*
Formance Stack API

Open, modular foundation for unique payments flows  # Introduction This API is documented in **OpenAPI format**.  # Authentication Formance Stack offers one forms of authentication:   - OAuth2 OAuth2 - an open protocol to allow secure authorization in a simple and standard method from web, mobile and desktop applications. <SecurityDefinitions /> 

API version: v0.2.2
Contact: support@formance.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package formance

import (
	"encoding/json"
)

// checks if the ClientAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClientAllOf{}

// ClientAllOf struct for ClientAllOf
type ClientAllOf struct {
	Id string `json:"id"`
	Scopes []string `json:"scopes,omitempty"`
	Secrets []ClientSecret `json:"secrets,omitempty"`
}

// NewClientAllOf instantiates a new ClientAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientAllOf(id string) *ClientAllOf {
	this := ClientAllOf{}
	this.Id = id
	return &this
}

// NewClientAllOfWithDefaults instantiates a new ClientAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientAllOfWithDefaults() *ClientAllOf {
	this := ClientAllOf{}
	return &this
}

// GetId returns the Id field value
func (o *ClientAllOf) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ClientAllOf) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ClientAllOf) SetId(v string) {
	o.Id = v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *ClientAllOf) GetScopes() []string {
	if o == nil || isNil(o.Scopes) {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientAllOf) GetScopesOk() ([]string, bool) {
	if o == nil || isNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *ClientAllOf) HasScopes() bool {
	if o != nil && !isNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *ClientAllOf) SetScopes(v []string) {
	o.Scopes = v
}

// GetSecrets returns the Secrets field value if set, zero value otherwise.
func (o *ClientAllOf) GetSecrets() []ClientSecret {
	if o == nil || isNil(o.Secrets) {
		var ret []ClientSecret
		return ret
	}
	return o.Secrets
}

// GetSecretsOk returns a tuple with the Secrets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientAllOf) GetSecretsOk() ([]ClientSecret, bool) {
	if o == nil || isNil(o.Secrets) {
		return nil, false
	}
	return o.Secrets, true
}

// HasSecrets returns a boolean if a field has been set.
func (o *ClientAllOf) HasSecrets() bool {
	if o != nil && !isNil(o.Secrets) {
		return true
	}

	return false
}

// SetSecrets gets a reference to the given []ClientSecret and assigns it to the Secrets field.
func (o *ClientAllOf) SetSecrets(v []ClientSecret) {
	o.Secrets = v
}

func (o ClientAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !isNil(o.Scopes) {
		toSerialize["scopes"] = o.Scopes
	}
	if !isNil(o.Secrets) {
		toSerialize["secrets"] = o.Secrets
	}
	return toSerialize, nil
}

type NullableClientAllOf struct {
	value *ClientAllOf
	isSet bool
}

func (v NullableClientAllOf) Get() *ClientAllOf {
	return v.value
}

func (v *NullableClientAllOf) Set(val *ClientAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableClientAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableClientAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientAllOf(val *ClientAllOf) *NullableClientAllOf {
	return &NullableClientAllOf{value: val, isSet: true}
}

func (v NullableClientAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


