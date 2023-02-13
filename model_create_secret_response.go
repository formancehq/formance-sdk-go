/*
Formance Stack API

Open, modular foundation for unique payments flows  # Introduction This API is documented in **OpenAPI format**.  # Authentication Formance Stack offers one forms of authentication:   - OAuth2 OAuth2 - an open protocol to allow secure authorization in a simple and standard method from web, mobile and desktop applications. <SecurityDefinitions /> 

API version: v1.0.0-rc.5
Contact: support@formance.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package formance

import (
	"encoding/json"
)

// checks if the CreateSecretResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSecretResponse{}

// CreateSecretResponse struct for CreateSecretResponse
type CreateSecretResponse struct {
	Data *Secret `json:"data,omitempty"`
}

// NewCreateSecretResponse instantiates a new CreateSecretResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSecretResponse() *CreateSecretResponse {
	this := CreateSecretResponse{}
	return &this
}

// NewCreateSecretResponseWithDefaults instantiates a new CreateSecretResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSecretResponseWithDefaults() *CreateSecretResponse {
	this := CreateSecretResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CreateSecretResponse) GetData() Secret {
	if o == nil || isNil(o.Data) {
		var ret Secret
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSecretResponse) GetDataOk() (*Secret, bool) {
	if o == nil || isNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CreateSecretResponse) HasData() bool {
	if o != nil && !isNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given Secret and assigns it to the Data field.
func (o *CreateSecretResponse) SetData(v Secret) {
	o.Data = &v
}

func (o CreateSecretResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateSecretResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableCreateSecretResponse struct {
	value *CreateSecretResponse
	isSet bool
}

func (v NullableCreateSecretResponse) Get() *CreateSecretResponse {
	return v.value
}

func (v *NullableCreateSecretResponse) Set(val *CreateSecretResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSecretResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSecretResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSecretResponse(val *CreateSecretResponse) *NullableCreateSecretResponse {
	return &NullableCreateSecretResponse{value: val, isSet: true}
}

func (v NullableCreateSecretResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSecretResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


