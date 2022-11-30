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

// checks if the ListConnectorsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListConnectorsResponse{}

// ListConnectorsResponse struct for ListConnectorsResponse
type ListConnectorsResponse struct {
	Data []ConnectorBaseInfo `json:"data"`
}

// NewListConnectorsResponse instantiates a new ListConnectorsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListConnectorsResponse(data []ConnectorBaseInfo) *ListConnectorsResponse {
	this := ListConnectorsResponse{}
	this.Data = data
	return &this
}

// NewListConnectorsResponseWithDefaults instantiates a new ListConnectorsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListConnectorsResponseWithDefaults() *ListConnectorsResponse {
	this := ListConnectorsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *ListConnectorsResponse) GetData() []ConnectorBaseInfo {
	if o == nil {
		var ret []ConnectorBaseInfo
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ListConnectorsResponse) GetDataOk() ([]ConnectorBaseInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ListConnectorsResponse) SetData(v []ConnectorBaseInfo) {
	o.Data = v
}

func (o ListConnectorsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListConnectorsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableListConnectorsResponse struct {
	value *ListConnectorsResponse
	isSet bool
}

func (v NullableListConnectorsResponse) Get() *ListConnectorsResponse {
	return v.value
}

func (v *NullableListConnectorsResponse) Set(val *ListConnectorsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListConnectorsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListConnectorsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListConnectorsResponse(val *ListConnectorsResponse) *NullableListConnectorsResponse {
	return &NullableListConnectorsResponse{value: val, isSet: true}
}

func (v NullableListConnectorsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListConnectorsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


