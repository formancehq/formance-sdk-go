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

// checks if the ListTransactions200ResponseCursorAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListTransactions200ResponseCursorAllOf{}

// ListTransactions200ResponseCursorAllOf struct for ListTransactions200ResponseCursorAllOf
type ListTransactions200ResponseCursorAllOf struct {
	Data []Transaction `json:"data"`
}

// NewListTransactions200ResponseCursorAllOf instantiates a new ListTransactions200ResponseCursorAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListTransactions200ResponseCursorAllOf(data []Transaction) *ListTransactions200ResponseCursorAllOf {
	this := ListTransactions200ResponseCursorAllOf{}
	this.Data = data
	return &this
}

// NewListTransactions200ResponseCursorAllOfWithDefaults instantiates a new ListTransactions200ResponseCursorAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListTransactions200ResponseCursorAllOfWithDefaults() *ListTransactions200ResponseCursorAllOf {
	this := ListTransactions200ResponseCursorAllOf{}
	return &this
}

// GetData returns the Data field value
func (o *ListTransactions200ResponseCursorAllOf) GetData() []Transaction {
	if o == nil {
		var ret []Transaction
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ListTransactions200ResponseCursorAllOf) GetDataOk() ([]Transaction, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ListTransactions200ResponseCursorAllOf) SetData(v []Transaction) {
	o.Data = v
}

func (o ListTransactions200ResponseCursorAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListTransactions200ResponseCursorAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableListTransactions200ResponseCursorAllOf struct {
	value *ListTransactions200ResponseCursorAllOf
	isSet bool
}

func (v NullableListTransactions200ResponseCursorAllOf) Get() *ListTransactions200ResponseCursorAllOf {
	return v.value
}

func (v *NullableListTransactions200ResponseCursorAllOf) Set(val *ListTransactions200ResponseCursorAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableListTransactions200ResponseCursorAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableListTransactions200ResponseCursorAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListTransactions200ResponseCursorAllOf(val *ListTransactions200ResponseCursorAllOf) *NullableListTransactions200ResponseCursorAllOf {
	return &NullableListTransactions200ResponseCursorAllOf{value: val, isSet: true}
}

func (v NullableListTransactions200ResponseCursorAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListTransactions200ResponseCursorAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


