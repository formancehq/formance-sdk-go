/*
Formance Stack API

Open, modular foundation for unique payments flows  # Introduction This API is documented in **OpenAPI format**.  # Authentication Formance Stack offers one forms of authentication:   - OAuth2 OAuth2 - an open protocol to allow secure authorization in a simple and standard method from web, mobile and desktop applications. <SecurityDefinitions /> 

API version: v1.0.20230228
Contact: support@formance.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package formance

import (
	"encoding/json"
)

// checks if the PaymentResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentResponse{}

// PaymentResponse struct for PaymentResponse
type PaymentResponse struct {
	Data Payment `json:"data"`
}

// NewPaymentResponse instantiates a new PaymentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentResponse(data Payment) *PaymentResponse {
	this := PaymentResponse{}
	this.Data = data
	return &this
}

// NewPaymentResponseWithDefaults instantiates a new PaymentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentResponseWithDefaults() *PaymentResponse {
	this := PaymentResponse{}
	return &this
}

// GetData returns the Data field value
func (o *PaymentResponse) GetData() Payment {
	if o == nil {
		var ret Payment
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PaymentResponse) GetDataOk() (*Payment, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *PaymentResponse) SetData(v Payment) {
	o.Data = v
}

func (o PaymentResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullablePaymentResponse struct {
	value *PaymentResponse
	isSet bool
}

func (v NullablePaymentResponse) Get() *PaymentResponse {
	return v.value
}

func (v *NullablePaymentResponse) Set(val *PaymentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentResponse(val *PaymentResponse) *NullablePaymentResponse {
	return &NullablePaymentResponse{value: val, isSet: true}
}

func (v NullablePaymentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

