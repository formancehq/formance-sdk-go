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

// checks if the ListTransactions200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListTransactions200Response{}

// ListTransactions200Response struct for ListTransactions200Response
type ListTransactions200Response struct {
	Cursor ListTransactions200ResponseCursor `json:"cursor"`
}

// NewListTransactions200Response instantiates a new ListTransactions200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListTransactions200Response(cursor ListTransactions200ResponseCursor) *ListTransactions200Response {
	this := ListTransactions200Response{}
	this.Cursor = cursor
	return &this
}

// NewListTransactions200ResponseWithDefaults instantiates a new ListTransactions200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListTransactions200ResponseWithDefaults() *ListTransactions200Response {
	this := ListTransactions200Response{}
	return &this
}

// GetCursor returns the Cursor field value
func (o *ListTransactions200Response) GetCursor() ListTransactions200ResponseCursor {
	if o == nil {
		var ret ListTransactions200ResponseCursor
		return ret
	}

	return o.Cursor
}

// GetCursorOk returns a tuple with the Cursor field value
// and a boolean to check if the value has been set.
func (o *ListTransactions200Response) GetCursorOk() (*ListTransactions200ResponseCursor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cursor, true
}

// SetCursor sets field value
func (o *ListTransactions200Response) SetCursor(v ListTransactions200ResponseCursor) {
	o.Cursor = v
}

func (o ListTransactions200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListTransactions200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cursor"] = o.Cursor
	return toSerialize, nil
}

type NullableListTransactions200Response struct {
	value *ListTransactions200Response
	isSet bool
}

func (v NullableListTransactions200Response) Get() *ListTransactions200Response {
	return v.value
}

func (v *NullableListTransactions200Response) Set(val *ListTransactions200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListTransactions200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListTransactions200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListTransactions200Response(val *ListTransactions200Response) *NullableListTransactions200Response {
	return &NullableListTransactions200Response{value: val, isSet: true}
}

func (v NullableListTransactions200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListTransactions200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


