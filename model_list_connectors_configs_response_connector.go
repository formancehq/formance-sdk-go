/*
Formance Stack API

Open, modular foundation for unique payments flows  # Introduction This API is documented in **OpenAPI format**.  # Authentication Formance Stack offers one forms of authentication:   - OAuth2 OAuth2 - an open protocol to allow secure authorization in a simple and standard method from web, mobile and desktop applications. <SecurityDefinitions /> 

API version: v0.2.8
Contact: support@formance.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package formance

import (
	"encoding/json"
)

// checks if the ListConnectorsConfigsResponseConnector type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListConnectorsConfigsResponseConnector{}

// ListConnectorsConfigsResponseConnector struct for ListConnectorsConfigsResponseConnector
type ListConnectorsConfigsResponseConnector struct {
	Key *ListConnectorsConfigsResponseConnectorKey `json:"key,omitempty"`
}

// NewListConnectorsConfigsResponseConnector instantiates a new ListConnectorsConfigsResponseConnector object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListConnectorsConfigsResponseConnector() *ListConnectorsConfigsResponseConnector {
	this := ListConnectorsConfigsResponseConnector{}
	return &this
}

// NewListConnectorsConfigsResponseConnectorWithDefaults instantiates a new ListConnectorsConfigsResponseConnector object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListConnectorsConfigsResponseConnectorWithDefaults() *ListConnectorsConfigsResponseConnector {
	this := ListConnectorsConfigsResponseConnector{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *ListConnectorsConfigsResponseConnector) GetKey() ListConnectorsConfigsResponseConnectorKey {
	if o == nil || isNil(o.Key) {
		var ret ListConnectorsConfigsResponseConnectorKey
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConnectorsConfigsResponseConnector) GetKeyOk() (*ListConnectorsConfigsResponseConnectorKey, bool) {
	if o == nil || isNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *ListConnectorsConfigsResponseConnector) HasKey() bool {
	if o != nil && !isNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given ListConnectorsConfigsResponseConnectorKey and assigns it to the Key field.
func (o *ListConnectorsConfigsResponseConnector) SetKey(v ListConnectorsConfigsResponseConnectorKey) {
	o.Key = &v
}

func (o ListConnectorsConfigsResponseConnector) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListConnectorsConfigsResponseConnector) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	return toSerialize, nil
}

type NullableListConnectorsConfigsResponseConnector struct {
	value *ListConnectorsConfigsResponseConnector
	isSet bool
}

func (v NullableListConnectorsConfigsResponseConnector) Get() *ListConnectorsConfigsResponseConnector {
	return v.value
}

func (v *NullableListConnectorsConfigsResponseConnector) Set(val *ListConnectorsConfigsResponseConnector) {
	v.value = val
	v.isSet = true
}

func (v NullableListConnectorsConfigsResponseConnector) IsSet() bool {
	return v.isSet
}

func (v *NullableListConnectorsConfigsResponseConnector) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListConnectorsConfigsResponseConnector(val *ListConnectorsConfigsResponseConnector) *NullableListConnectorsConfigsResponseConnector {
	return &NullableListConnectorsConfigsResponseConnector{value: val, isSet: true}
}

func (v NullableListConnectorsConfigsResponseConnector) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListConnectorsConfigsResponseConnector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


