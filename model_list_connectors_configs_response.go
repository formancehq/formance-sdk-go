/*
Formance Stack API

Open, modular foundation for unique payments flows  # Introduction This API is documented in **OpenAPI format**.  # Authentication Formance Stack offers one forms of authentication:   - OAuth2 OAuth2 - an open protocol to allow secure authorization in a simple and standard method from web, mobile and desktop applications. <SecurityDefinitions /> 

API version: v1.0.0-rc.3
Contact: support@formance.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package formance

import (
	"encoding/json"
)

// checks if the ListConnectorsConfigsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListConnectorsConfigsResponse{}

// ListConnectorsConfigsResponse struct for ListConnectorsConfigsResponse
type ListConnectorsConfigsResponse struct {
	Connector *ListConnectorsConfigsResponseConnector `json:"connector,omitempty"`
}

// NewListConnectorsConfigsResponse instantiates a new ListConnectorsConfigsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListConnectorsConfigsResponse() *ListConnectorsConfigsResponse {
	this := ListConnectorsConfigsResponse{}
	return &this
}

// NewListConnectorsConfigsResponseWithDefaults instantiates a new ListConnectorsConfigsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListConnectorsConfigsResponseWithDefaults() *ListConnectorsConfigsResponse {
	this := ListConnectorsConfigsResponse{}
	return &this
}

// GetConnector returns the Connector field value if set, zero value otherwise.
func (o *ListConnectorsConfigsResponse) GetConnector() ListConnectorsConfigsResponseConnector {
	if o == nil || isNil(o.Connector) {
		var ret ListConnectorsConfigsResponseConnector
		return ret
	}
	return *o.Connector
}

// GetConnectorOk returns a tuple with the Connector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConnectorsConfigsResponse) GetConnectorOk() (*ListConnectorsConfigsResponseConnector, bool) {
	if o == nil || isNil(o.Connector) {
		return nil, false
	}
	return o.Connector, true
}

// HasConnector returns a boolean if a field has been set.
func (o *ListConnectorsConfigsResponse) HasConnector() bool {
	if o != nil && !isNil(o.Connector) {
		return true
	}

	return false
}

// SetConnector gets a reference to the given ListConnectorsConfigsResponseConnector and assigns it to the Connector field.
func (o *ListConnectorsConfigsResponse) SetConnector(v ListConnectorsConfigsResponseConnector) {
	o.Connector = &v
}

func (o ListConnectorsConfigsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListConnectorsConfigsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Connector) {
		toSerialize["connector"] = o.Connector
	}
	return toSerialize, nil
}

type NullableListConnectorsConfigsResponse struct {
	value *ListConnectorsConfigsResponse
	isSet bool
}

func (v NullableListConnectorsConfigsResponse) Get() *ListConnectorsConfigsResponse {
	return v.value
}

func (v *NullableListConnectorsConfigsResponse) Set(val *ListConnectorsConfigsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListConnectorsConfigsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListConnectorsConfigsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListConnectorsConfigsResponse(val *ListConnectorsConfigsResponse) *NullableListConnectorsConfigsResponse {
	return &NullableListConnectorsConfigsResponse{value: val, isSet: true}
}

func (v NullableListConnectorsConfigsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListConnectorsConfigsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


