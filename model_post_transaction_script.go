/*
Formance Stack API

Open, modular foundation for unique payments flows  # Introduction This API is documented in **OpenAPI format**.  # Authentication Formance Stack offers one forms of authentication:   - OAuth2 OAuth2 - an open protocol to allow secure authorization in a simple and standard method from web, mobile and desktop applications. <SecurityDefinitions /> 

API version: v0.2.3
Contact: support@formance.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package formance

import (
	"encoding/json"
)

// checks if the PostTransactionScript type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostTransactionScript{}

// PostTransactionScript struct for PostTransactionScript
type PostTransactionScript struct {
	Plain string `json:"plain"`
	Vars map[string]interface{} `json:"vars,omitempty"`
}

// NewPostTransactionScript instantiates a new PostTransactionScript object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostTransactionScript(plain string) *PostTransactionScript {
	this := PostTransactionScript{}
	this.Plain = plain
	return &this
}

// NewPostTransactionScriptWithDefaults instantiates a new PostTransactionScript object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostTransactionScriptWithDefaults() *PostTransactionScript {
	this := PostTransactionScript{}
	return &this
}

// GetPlain returns the Plain field value
func (o *PostTransactionScript) GetPlain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Plain
}

// GetPlainOk returns a tuple with the Plain field value
// and a boolean to check if the value has been set.
func (o *PostTransactionScript) GetPlainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Plain, true
}

// SetPlain sets field value
func (o *PostTransactionScript) SetPlain(v string) {
	o.Plain = v
}

// GetVars returns the Vars field value if set, zero value otherwise.
func (o *PostTransactionScript) GetVars() map[string]interface{} {
	if o == nil || isNil(o.Vars) {
		var ret map[string]interface{}
		return ret
	}
	return o.Vars
}

// GetVarsOk returns a tuple with the Vars field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostTransactionScript) GetVarsOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Vars) {
		return map[string]interface{}{}, false
	}
	return o.Vars, true
}

// HasVars returns a boolean if a field has been set.
func (o *PostTransactionScript) HasVars() bool {
	if o != nil && !isNil(o.Vars) {
		return true
	}

	return false
}

// SetVars gets a reference to the given map[string]interface{} and assigns it to the Vars field.
func (o *PostTransactionScript) SetVars(v map[string]interface{}) {
	o.Vars = v
}

func (o PostTransactionScript) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostTransactionScript) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["plain"] = o.Plain
	if !isNil(o.Vars) {
		toSerialize["vars"] = o.Vars
	}
	return toSerialize, nil
}

type NullablePostTransactionScript struct {
	value *PostTransactionScript
	isSet bool
}

func (v NullablePostTransactionScript) Get() *PostTransactionScript {
	return v.value
}

func (v *NullablePostTransactionScript) Set(val *PostTransactionScript) {
	v.value = val
	v.isSet = true
}

func (v NullablePostTransactionScript) IsSet() bool {
	return v.isSet
}

func (v *NullablePostTransactionScript) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostTransactionScript(val *PostTransactionScript) *NullablePostTransactionScript {
	return &NullablePostTransactionScript{value: val, isSet: true}
}

func (v NullablePostTransactionScript) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostTransactionScript) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


