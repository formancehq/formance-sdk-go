/*
Formance Stack API

Open, modular foundation for unique payments flows  # Introduction This API is documented in **OpenAPI format**.  # Authentication Formance Stack offers one forms of authentication:   - OAuth2 OAuth2 - an open protocol to allow secure authorization in a simple and standard method from web, mobile and desktop applications. <SecurityDefinitions /> 

API version: v1.0.0
Contact: support@formance.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package formance

import (
	"encoding/json"
)

// checks if the ListScopesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListScopesResponse{}

// ListScopesResponse struct for ListScopesResponse
type ListScopesResponse struct {
	Data []Scope `json:"data,omitempty"`
}

// NewListScopesResponse instantiates a new ListScopesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListScopesResponse() *ListScopesResponse {
	this := ListScopesResponse{}
	return &this
}

// NewListScopesResponseWithDefaults instantiates a new ListScopesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListScopesResponseWithDefaults() *ListScopesResponse {
	this := ListScopesResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListScopesResponse) GetData() []Scope {
	if o == nil || isNil(o.Data) {
		var ret []Scope
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListScopesResponse) GetDataOk() ([]Scope, bool) {
	if o == nil || isNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListScopesResponse) HasData() bool {
	if o != nil && !isNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []Scope and assigns it to the Data field.
func (o *ListScopesResponse) SetData(v []Scope) {
	o.Data = v
}

func (o ListScopesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListScopesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableListScopesResponse struct {
	value *ListScopesResponse
	isSet bool
}

func (v NullableListScopesResponse) Get() *ListScopesResponse {
	return v.value
}

func (v *NullableListScopesResponse) Set(val *ListScopesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListScopesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListScopesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListScopesResponse(val *ListScopesResponse) *NullableListScopesResponse {
	return &NullableListScopesResponse{value: val, isSet: true}
}

func (v NullableListScopesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListScopesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


