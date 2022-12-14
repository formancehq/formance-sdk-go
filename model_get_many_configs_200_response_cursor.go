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

// checks if the GetManyConfigs200ResponseCursor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetManyConfigs200ResponseCursor{}

// GetManyConfigs200ResponseCursor struct for GetManyConfigs200ResponseCursor
type GetManyConfigs200ResponseCursor struct {
	HasMore *bool `json:"has_more,omitempty"`
	Data []WebhooksConfig `json:"data"`
}

// NewGetManyConfigs200ResponseCursor instantiates a new GetManyConfigs200ResponseCursor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetManyConfigs200ResponseCursor(data []WebhooksConfig) *GetManyConfigs200ResponseCursor {
	this := GetManyConfigs200ResponseCursor{}
	this.Data = data
	return &this
}

// NewGetManyConfigs200ResponseCursorWithDefaults instantiates a new GetManyConfigs200ResponseCursor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetManyConfigs200ResponseCursorWithDefaults() *GetManyConfigs200ResponseCursor {
	this := GetManyConfigs200ResponseCursor{}
	return &this
}

// GetHasMore returns the HasMore field value if set, zero value otherwise.
func (o *GetManyConfigs200ResponseCursor) GetHasMore() bool {
	if o == nil || isNil(o.HasMore) {
		var ret bool
		return ret
	}
	return *o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetManyConfigs200ResponseCursor) GetHasMoreOk() (*bool, bool) {
	if o == nil || isNil(o.HasMore) {
		return nil, false
	}
	return o.HasMore, true
}

// HasHasMore returns a boolean if a field has been set.
func (o *GetManyConfigs200ResponseCursor) HasHasMore() bool {
	if o != nil && !isNil(o.HasMore) {
		return true
	}

	return false
}

// SetHasMore gets a reference to the given bool and assigns it to the HasMore field.
func (o *GetManyConfigs200ResponseCursor) SetHasMore(v bool) {
	o.HasMore = &v
}

// GetData returns the Data field value
func (o *GetManyConfigs200ResponseCursor) GetData() []WebhooksConfig {
	if o == nil {
		var ret []WebhooksConfig
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetManyConfigs200ResponseCursor) GetDataOk() ([]WebhooksConfig, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *GetManyConfigs200ResponseCursor) SetData(v []WebhooksConfig) {
	o.Data = v
}

func (o GetManyConfigs200ResponseCursor) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetManyConfigs200ResponseCursor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.HasMore) {
		toSerialize["has_more"] = o.HasMore
	}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableGetManyConfigs200ResponseCursor struct {
	value *GetManyConfigs200ResponseCursor
	isSet bool
}

func (v NullableGetManyConfigs200ResponseCursor) Get() *GetManyConfigs200ResponseCursor {
	return v.value
}

func (v *NullableGetManyConfigs200ResponseCursor) Set(val *GetManyConfigs200ResponseCursor) {
	v.value = val
	v.isSet = true
}

func (v NullableGetManyConfigs200ResponseCursor) IsSet() bool {
	return v.isSet
}

func (v *NullableGetManyConfigs200ResponseCursor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetManyConfigs200ResponseCursor(val *GetManyConfigs200ResponseCursor) *NullableGetManyConfigs200ResponseCursor {
	return &NullableGetManyConfigs200ResponseCursor{value: val, isSet: true}
}

func (v NullableGetManyConfigs200ResponseCursor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetManyConfigs200ResponseCursor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


