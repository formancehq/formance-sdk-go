/*
Formance Stack API

Open, modular foundation for unique payments flows  # Introduction This API is documented in **OpenAPI format**.  # Authentication Formance Stack offers one forms of authentication:   - OAuth2 OAuth2 - an open protocol to allow secure authorization in a simple and standard method from web, mobile and desktop applications. <SecurityDefinitions /> 

API version: SDK_VERSION
Contact: support@formance.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package formance

import (
	"encoding/json"
)

// checks if the AddMetadataToAccount409Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddMetadataToAccount409Response{}

// AddMetadataToAccount409Response struct for AddMetadataToAccount409Response
type AddMetadataToAccount409Response struct {
	ErrorCode string `json:"error_code"`
	ErrorMessage *string `json:"error_message,omitempty"`
}

// NewAddMetadataToAccount409Response instantiates a new AddMetadataToAccount409Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddMetadataToAccount409Response(errorCode string) *AddMetadataToAccount409Response {
	this := AddMetadataToAccount409Response{}
	this.ErrorCode = errorCode
	return &this
}

// NewAddMetadataToAccount409ResponseWithDefaults instantiates a new AddMetadataToAccount409Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddMetadataToAccount409ResponseWithDefaults() *AddMetadataToAccount409Response {
	this := AddMetadataToAccount409Response{}
	return &this
}

// GetErrorCode returns the ErrorCode field value
func (o *AddMetadataToAccount409Response) GetErrorCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value
// and a boolean to check if the value has been set.
func (o *AddMetadataToAccount409Response) GetErrorCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorCode, true
}

// SetErrorCode sets field value
func (o *AddMetadataToAccount409Response) SetErrorCode(v string) {
	o.ErrorCode = v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *AddMetadataToAccount409Response) GetErrorMessage() string {
	if o == nil || isNil(o.ErrorMessage) {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddMetadataToAccount409Response) GetErrorMessageOk() (*string, bool) {
	if o == nil || isNil(o.ErrorMessage) {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *AddMetadataToAccount409Response) HasErrorMessage() bool {
	if o != nil && !isNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *AddMetadataToAccount409Response) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

func (o AddMetadataToAccount409Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddMetadataToAccount409Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["error_code"] = o.ErrorCode
	if !isNil(o.ErrorMessage) {
		toSerialize["error_message"] = o.ErrorMessage
	}
	return toSerialize, nil
}

type NullableAddMetadataToAccount409Response struct {
	value *AddMetadataToAccount409Response
	isSet bool
}

func (v NullableAddMetadataToAccount409Response) Get() *AddMetadataToAccount409Response {
	return v.value
}

func (v *NullableAddMetadataToAccount409Response) Set(val *AddMetadataToAccount409Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAddMetadataToAccount409Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAddMetadataToAccount409Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddMetadataToAccount409Response(val *AddMetadataToAccount409Response) *NullableAddMetadataToAccount409Response {
	return &NullableAddMetadataToAccount409Response{value: val, isSet: true}
}

func (v NullableAddMetadataToAccount409Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddMetadataToAccount409Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


