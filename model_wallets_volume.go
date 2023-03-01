/*
Formance Stack API

Open, modular foundation for unique payments flows  # Introduction This API is documented in **OpenAPI format**.  # Authentication Formance Stack offers one forms of authentication:   - OAuth2 OAuth2 - an open protocol to allow secure authorization in a simple and standard method from web, mobile and desktop applications. <SecurityDefinitions /> 

API version: v1.0.20230301
Contact: support@formance.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package formance

import (
	"encoding/json"
)

// checks if the WalletsVolume type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WalletsVolume{}

// WalletsVolume struct for WalletsVolume
type WalletsVolume struct {
	Input int32 `json:"input"`
	Output int32 `json:"output"`
	Balance int32 `json:"balance"`
}

// NewWalletsVolume instantiates a new WalletsVolume object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWalletsVolume(input int32, output int32, balance int32) *WalletsVolume {
	this := WalletsVolume{}
	this.Input = input
	this.Output = output
	this.Balance = balance
	return &this
}

// NewWalletsVolumeWithDefaults instantiates a new WalletsVolume object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWalletsVolumeWithDefaults() *WalletsVolume {
	this := WalletsVolume{}
	return &this
}

// GetInput returns the Input field value
func (o *WalletsVolume) GetInput() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Input
}

// GetInputOk returns a tuple with the Input field value
// and a boolean to check if the value has been set.
func (o *WalletsVolume) GetInputOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Input, true
}

// SetInput sets field value
func (o *WalletsVolume) SetInput(v int32) {
	o.Input = v
}

// GetOutput returns the Output field value
func (o *WalletsVolume) GetOutput() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Output
}

// GetOutputOk returns a tuple with the Output field value
// and a boolean to check if the value has been set.
func (o *WalletsVolume) GetOutputOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Output, true
}

// SetOutput sets field value
func (o *WalletsVolume) SetOutput(v int32) {
	o.Output = v
}

// GetBalance returns the Balance field value
func (o *WalletsVolume) GetBalance() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value
// and a boolean to check if the value has been set.
func (o *WalletsVolume) GetBalanceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Balance, true
}

// SetBalance sets field value
func (o *WalletsVolume) SetBalance(v int32) {
	o.Balance = v
}

func (o WalletsVolume) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WalletsVolume) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["input"] = o.Input
	toSerialize["output"] = o.Output
	toSerialize["balance"] = o.Balance
	return toSerialize, nil
}

type NullableWalletsVolume struct {
	value *WalletsVolume
	isSet bool
}

func (v NullableWalletsVolume) Get() *WalletsVolume {
	return v.value
}

func (v *NullableWalletsVolume) Set(val *WalletsVolume) {
	v.value = val
	v.isSet = true
}

func (v NullableWalletsVolume) IsSet() bool {
	return v.isSet
}

func (v *NullableWalletsVolume) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWalletsVolume(val *WalletsVolume) *NullableWalletsVolume {
	return &NullableWalletsVolume{value: val, isSet: true}
}

func (v NullableWalletsVolume) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWalletsVolume) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


