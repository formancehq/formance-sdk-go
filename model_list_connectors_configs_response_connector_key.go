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

// checks if the ListConnectorsConfigsResponseConnectorKey type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListConnectorsConfigsResponseConnectorKey{}

// ListConnectorsConfigsResponseConnectorKey struct for ListConnectorsConfigsResponseConnectorKey
type ListConnectorsConfigsResponseConnectorKey struct {
	Datatype *string `json:"datatype,omitempty"`
	Required *bool `json:"required,omitempty"`
}

// NewListConnectorsConfigsResponseConnectorKey instantiates a new ListConnectorsConfigsResponseConnectorKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListConnectorsConfigsResponseConnectorKey() *ListConnectorsConfigsResponseConnectorKey {
	this := ListConnectorsConfigsResponseConnectorKey{}
	return &this
}

// NewListConnectorsConfigsResponseConnectorKeyWithDefaults instantiates a new ListConnectorsConfigsResponseConnectorKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListConnectorsConfigsResponseConnectorKeyWithDefaults() *ListConnectorsConfigsResponseConnectorKey {
	this := ListConnectorsConfigsResponseConnectorKey{}
	return &this
}

// GetDatatype returns the Datatype field value if set, zero value otherwise.
func (o *ListConnectorsConfigsResponseConnectorKey) GetDatatype() string {
	if o == nil || isNil(o.Datatype) {
		var ret string
		return ret
	}
	return *o.Datatype
}

// GetDatatypeOk returns a tuple with the Datatype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConnectorsConfigsResponseConnectorKey) GetDatatypeOk() (*string, bool) {
	if o == nil || isNil(o.Datatype) {
		return nil, false
	}
	return o.Datatype, true
}

// HasDatatype returns a boolean if a field has been set.
func (o *ListConnectorsConfigsResponseConnectorKey) HasDatatype() bool {
	if o != nil && !isNil(o.Datatype) {
		return true
	}

	return false
}

// SetDatatype gets a reference to the given string and assigns it to the Datatype field.
func (o *ListConnectorsConfigsResponseConnectorKey) SetDatatype(v string) {
	o.Datatype = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *ListConnectorsConfigsResponseConnectorKey) GetRequired() bool {
	if o == nil || isNil(o.Required) {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConnectorsConfigsResponseConnectorKey) GetRequiredOk() (*bool, bool) {
	if o == nil || isNil(o.Required) {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *ListConnectorsConfigsResponseConnectorKey) HasRequired() bool {
	if o != nil && !isNil(o.Required) {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *ListConnectorsConfigsResponseConnectorKey) SetRequired(v bool) {
	o.Required = &v
}

func (o ListConnectorsConfigsResponseConnectorKey) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListConnectorsConfigsResponseConnectorKey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Datatype) {
		toSerialize["datatype"] = o.Datatype
	}
	if !isNil(o.Required) {
		toSerialize["required"] = o.Required
	}
	return toSerialize, nil
}

type NullableListConnectorsConfigsResponseConnectorKey struct {
	value *ListConnectorsConfigsResponseConnectorKey
	isSet bool
}

func (v NullableListConnectorsConfigsResponseConnectorKey) Get() *ListConnectorsConfigsResponseConnectorKey {
	return v.value
}

func (v *NullableListConnectorsConfigsResponseConnectorKey) Set(val *ListConnectorsConfigsResponseConnectorKey) {
	v.value = val
	v.isSet = true
}

func (v NullableListConnectorsConfigsResponseConnectorKey) IsSet() bool {
	return v.isSet
}

func (v *NullableListConnectorsConfigsResponseConnectorKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListConnectorsConfigsResponseConnectorKey(val *ListConnectorsConfigsResponseConnectorKey) *NullableListConnectorsConfigsResponseConnectorKey {
	return &NullableListConnectorsConfigsResponseConnectorKey{value: val, isSet: true}
}

func (v NullableListConnectorsConfigsResponseConnectorKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListConnectorsConfigsResponseConnectorKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


