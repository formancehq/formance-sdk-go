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

// checks if the CreateTransactions400Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTransactions400Response{}

// CreateTransactions400Response struct for CreateTransactions400Response
type CreateTransactions400Response struct {
	ErrorCode string `json:"error_code"`
	ErrorMessage *string `json:"error_message,omitempty"`
}

// NewCreateTransactions400Response instantiates a new CreateTransactions400Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTransactions400Response(errorCode string) *CreateTransactions400Response {
	this := CreateTransactions400Response{}
	this.ErrorCode = errorCode
	return &this
}

// NewCreateTransactions400ResponseWithDefaults instantiates a new CreateTransactions400Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTransactions400ResponseWithDefaults() *CreateTransactions400Response {
	this := CreateTransactions400Response{}
	return &this
}

// GetErrorCode returns the ErrorCode field value
func (o *CreateTransactions400Response) GetErrorCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value
// and a boolean to check if the value has been set.
func (o *CreateTransactions400Response) GetErrorCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorCode, true
}

// SetErrorCode sets field value
func (o *CreateTransactions400Response) SetErrorCode(v string) {
	o.ErrorCode = v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *CreateTransactions400Response) GetErrorMessage() string {
	if o == nil || isNil(o.ErrorMessage) {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTransactions400Response) GetErrorMessageOk() (*string, bool) {
	if o == nil || isNil(o.ErrorMessage) {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *CreateTransactions400Response) HasErrorMessage() bool {
	if o != nil && !isNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *CreateTransactions400Response) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

func (o CreateTransactions400Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateTransactions400Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["error_code"] = o.ErrorCode
	if !isNil(o.ErrorMessage) {
		toSerialize["error_message"] = o.ErrorMessage
	}
	return toSerialize, nil
}

type NullableCreateTransactions400Response struct {
	value *CreateTransactions400Response
	isSet bool
}

func (v NullableCreateTransactions400Response) Get() *CreateTransactions400Response {
	return v.value
}

func (v *NullableCreateTransactions400Response) Set(val *CreateTransactions400Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTransactions400Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTransactions400Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTransactions400Response(val *CreateTransactions400Response) *NullableCreateTransactions400Response {
	return &NullableCreateTransactions400Response{value: val, isSet: true}
}

func (v NullableCreateTransactions400Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTransactions400Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


