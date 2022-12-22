/*
Formance Stack API

Open, modular foundation for unique payments flows  # Introduction This API is documented in **OpenAPI format**.  # Authentication Formance Stack offers one forms of authentication:   - OAuth2 OAuth2 - an open protocol to allow secure authorization in a simple and standard method from web, mobile and desktop applications. <SecurityDefinitions /> 

API version: v1.0.0-rc.1
Contact: support@formance.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package formance

import (
	"encoding/json"
)

// checks if the Scope type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Scope{}

// Scope struct for Scope
type Scope struct {
	Label string `json:"label"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Id string `json:"id"`
	Transient []string `json:"transient,omitempty"`
}

// NewScope instantiates a new Scope object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScope(label string, id string) *Scope {
	this := Scope{}
	this.Label = label
	this.Id = id
	return &this
}

// NewScopeWithDefaults instantiates a new Scope object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScopeWithDefaults() *Scope {
	this := Scope{}
	return &this
}

// GetLabel returns the Label field value
func (o *Scope) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *Scope) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *Scope) SetLabel(v string) {
	o.Label = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *Scope) GetMetadata() map[string]interface{} {
	if o == nil || isNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Scope) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Scope) HasMetadata() bool {
	if o != nil && !isNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *Scope) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetId returns the Id field value
func (o *Scope) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Scope) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Scope) SetId(v string) {
	o.Id = v
}

// GetTransient returns the Transient field value if set, zero value otherwise.
func (o *Scope) GetTransient() []string {
	if o == nil || isNil(o.Transient) {
		var ret []string
		return ret
	}
	return o.Transient
}

// GetTransientOk returns a tuple with the Transient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Scope) GetTransientOk() ([]string, bool) {
	if o == nil || isNil(o.Transient) {
		return nil, false
	}
	return o.Transient, true
}

// HasTransient returns a boolean if a field has been set.
func (o *Scope) HasTransient() bool {
	if o != nil && !isNil(o.Transient) {
		return true
	}

	return false
}

// SetTransient gets a reference to the given []string and assigns it to the Transient field.
func (o *Scope) SetTransient(v []string) {
	o.Transient = v
}

func (o Scope) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Scope) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["label"] = o.Label
	if !isNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	toSerialize["id"] = o.Id
	if !isNil(o.Transient) {
		toSerialize["transient"] = o.Transient
	}
	return toSerialize, nil
}

type NullableScope struct {
	value *Scope
	isSet bool
}

func (v NullableScope) Get() *Scope {
	return v.value
}

func (v *NullableScope) Set(val *Scope) {
	v.value = val
	v.isSet = true
}

func (v NullableScope) IsSet() bool {
	return v.isSet
}

func (v *NullableScope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScope(val *Scope) *NullableScope {
	return &NullableScope{value: val, isSet: true}
}

func (v NullableScope) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


