/*
Formance Stack API

Open, modular foundation for unique payments flows  # Introduction This API is documented in **OpenAPI format**.  # Authentication Formance Stack offers one forms of authentication:   - OAuth2 OAuth2 - an open protocol to allow secure authorization in a simple and standard method from web, mobile and desktop applications. <SecurityDefinitions /> 

API version: v0.2.4
Contact: support@formance.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package formance

import (
	"encoding/json"
)

// checks if the GetBalancesAggregated400Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetBalancesAggregated400Response{}

// GetBalancesAggregated400Response struct for GetBalancesAggregated400Response
type GetBalancesAggregated400Response struct {
	ErrorCode string `json:"error_code"`
	ErrorMessage *string `json:"error_message,omitempty"`
}

// NewGetBalancesAggregated400Response instantiates a new GetBalancesAggregated400Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBalancesAggregated400Response(errorCode string) *GetBalancesAggregated400Response {
	this := GetBalancesAggregated400Response{}
	this.ErrorCode = errorCode
	return &this
}

// NewGetBalancesAggregated400ResponseWithDefaults instantiates a new GetBalancesAggregated400Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBalancesAggregated400ResponseWithDefaults() *GetBalancesAggregated400Response {
	this := GetBalancesAggregated400Response{}
	return &this
}

// GetErrorCode returns the ErrorCode field value
func (o *GetBalancesAggregated400Response) GetErrorCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value
// and a boolean to check if the value has been set.
func (o *GetBalancesAggregated400Response) GetErrorCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorCode, true
}

// SetErrorCode sets field value
func (o *GetBalancesAggregated400Response) SetErrorCode(v string) {
	o.ErrorCode = v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *GetBalancesAggregated400Response) GetErrorMessage() string {
	if o == nil || isNil(o.ErrorMessage) {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalancesAggregated400Response) GetErrorMessageOk() (*string, bool) {
	if o == nil || isNil(o.ErrorMessage) {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *GetBalancesAggregated400Response) HasErrorMessage() bool {
	if o != nil && !isNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *GetBalancesAggregated400Response) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

func (o GetBalancesAggregated400Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetBalancesAggregated400Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["error_code"] = o.ErrorCode
	if !isNil(o.ErrorMessage) {
		toSerialize["error_message"] = o.ErrorMessage
	}
	return toSerialize, nil
}

type NullableGetBalancesAggregated400Response struct {
	value *GetBalancesAggregated400Response
	isSet bool
}

func (v NullableGetBalancesAggregated400Response) Get() *GetBalancesAggregated400Response {
	return v.value
}

func (v *NullableGetBalancesAggregated400Response) Set(val *GetBalancesAggregated400Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBalancesAggregated400Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBalancesAggregated400Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBalancesAggregated400Response(val *GetBalancesAggregated400Response) *NullableGetBalancesAggregated400Response {
	return &NullableGetBalancesAggregated400Response{value: val, isSet: true}
}

func (v NullableGetBalancesAggregated400Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBalancesAggregated400Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


