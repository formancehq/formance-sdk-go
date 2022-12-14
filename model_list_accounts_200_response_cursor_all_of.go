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

// checks if the ListAccounts200ResponseCursorAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListAccounts200ResponseCursorAllOf{}

// ListAccounts200ResponseCursorAllOf struct for ListAccounts200ResponseCursorAllOf
type ListAccounts200ResponseCursorAllOf struct {
	Data []Account `json:"data"`
}

// NewListAccounts200ResponseCursorAllOf instantiates a new ListAccounts200ResponseCursorAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListAccounts200ResponseCursorAllOf(data []Account) *ListAccounts200ResponseCursorAllOf {
	this := ListAccounts200ResponseCursorAllOf{}
	this.Data = data
	return &this
}

// NewListAccounts200ResponseCursorAllOfWithDefaults instantiates a new ListAccounts200ResponseCursorAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListAccounts200ResponseCursorAllOfWithDefaults() *ListAccounts200ResponseCursorAllOf {
	this := ListAccounts200ResponseCursorAllOf{}
	return &this
}

// GetData returns the Data field value
func (o *ListAccounts200ResponseCursorAllOf) GetData() []Account {
	if o == nil {
		var ret []Account
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ListAccounts200ResponseCursorAllOf) GetDataOk() ([]Account, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ListAccounts200ResponseCursorAllOf) SetData(v []Account) {
	o.Data = v
}

func (o ListAccounts200ResponseCursorAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListAccounts200ResponseCursorAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableListAccounts200ResponseCursorAllOf struct {
	value *ListAccounts200ResponseCursorAllOf
	isSet bool
}

func (v NullableListAccounts200ResponseCursorAllOf) Get() *ListAccounts200ResponseCursorAllOf {
	return v.value
}

func (v *NullableListAccounts200ResponseCursorAllOf) Set(val *ListAccounts200ResponseCursorAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableListAccounts200ResponseCursorAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableListAccounts200ResponseCursorAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListAccounts200ResponseCursorAllOf(val *ListAccounts200ResponseCursorAllOf) *NullableListAccounts200ResponseCursorAllOf {
	return &NullableListAccounts200ResponseCursorAllOf{value: val, isSet: true}
}

func (v NullableListAccounts200ResponseCursorAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListAccounts200ResponseCursorAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


