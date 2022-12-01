/*
Formance Stack API

Open, modular foundation for unique payments flows  # Introduction This API is documented in **OpenAPI format**.  # Authentication Formance Stack offers one forms of authentication:   - OAuth2 OAuth2 - an open protocol to allow secure authorization in a simple and standard method from web, mobile and desktop applications. <SecurityDefinitions /> 

API version: v0.2.4
Contact: support@formance.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package formance

import (
	"encoding/json"
)

// checks if the ConnectorBaseInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConnectorBaseInfo{}

// ConnectorBaseInfo struct for ConnectorBaseInfo
type ConnectorBaseInfo struct {
	Provider *string `json:"provider,omitempty"`
	Disabled *bool `json:"disabled,omitempty"`
}

// NewConnectorBaseInfo instantiates a new ConnectorBaseInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorBaseInfo() *ConnectorBaseInfo {
	this := ConnectorBaseInfo{}
	return &this
}

// NewConnectorBaseInfoWithDefaults instantiates a new ConnectorBaseInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorBaseInfoWithDefaults() *ConnectorBaseInfo {
	this := ConnectorBaseInfo{}
	return &this
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *ConnectorBaseInfo) GetProvider() string {
	if o == nil || isNil(o.Provider) {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorBaseInfo) GetProviderOk() (*string, bool) {
	if o == nil || isNil(o.Provider) {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *ConnectorBaseInfo) HasProvider() bool {
	if o != nil && !isNil(o.Provider) {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *ConnectorBaseInfo) SetProvider(v string) {
	o.Provider = &v
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *ConnectorBaseInfo) GetDisabled() bool {
	if o == nil || isNil(o.Disabled) {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorBaseInfo) GetDisabledOk() (*bool, bool) {
	if o == nil || isNil(o.Disabled) {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *ConnectorBaseInfo) HasDisabled() bool {
	if o != nil && !isNil(o.Disabled) {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *ConnectorBaseInfo) SetDisabled(v bool) {
	o.Disabled = &v
}

func (o ConnectorBaseInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConnectorBaseInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Provider) {
		toSerialize["provider"] = o.Provider
	}
	if !isNil(o.Disabled) {
		toSerialize["disabled"] = o.Disabled
	}
	return toSerialize, nil
}

type NullableConnectorBaseInfo struct {
	value *ConnectorBaseInfo
	isSet bool
}

func (v NullableConnectorBaseInfo) Get() *ConnectorBaseInfo {
	return v.value
}

func (v *NullableConnectorBaseInfo) Set(val *ConnectorBaseInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorBaseInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorBaseInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorBaseInfo(val *ConnectorBaseInfo) *NullableConnectorBaseInfo {
	return &NullableConnectorBaseInfo{value: val, isSet: true}
}

func (v NullableConnectorBaseInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorBaseInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


