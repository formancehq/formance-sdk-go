/*
Formance Stack API

Open, modular foundation for unique payments flows  # Introduction This API is documented in **OpenAPI format**.  # Authentication Formance Stack offers one forms of authentication:   - OAuth2 OAuth2 - an open protocol to allow secure authorization in a simple and standard method from web, mobile and desktop applications. <SecurityDefinitions /> 

API version: develop
Contact: support@formance.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package formance

import (
	"encoding/json"
)

// checks if the ServerInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerInfo{}

// ServerInfo struct for ServerInfo
type ServerInfo struct {
	Version string `json:"version"`
}

// NewServerInfo instantiates a new ServerInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerInfo(version string) *ServerInfo {
	this := ServerInfo{}
	this.Version = version
	return &this
}

// NewServerInfoWithDefaults instantiates a new ServerInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerInfoWithDefaults() *ServerInfo {
	this := ServerInfo{}
	return &this
}

// GetVersion returns the Version field value
func (o *ServerInfo) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *ServerInfo) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *ServerInfo) SetVersion(v string) {
	o.Version = v
}

func (o ServerInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServerInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["version"] = o.Version
	return toSerialize, nil
}

type NullableServerInfo struct {
	value *ServerInfo
	isSet bool
}

func (v NullableServerInfo) Get() *ServerInfo {
	return v.value
}

func (v *NullableServerInfo) Set(val *ServerInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableServerInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableServerInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerInfo(val *ServerInfo) *NullableServerInfo {
	return &NullableServerInfo{value: val, isSet: true}
}

func (v NullableServerInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


