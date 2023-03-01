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

// checks if the GetWorkflowResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetWorkflowResponse{}

// GetWorkflowResponse struct for GetWorkflowResponse
type GetWorkflowResponse struct {
	Data Workflow `json:"data"`
}

// NewGetWorkflowResponse instantiates a new GetWorkflowResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetWorkflowResponse(data Workflow) *GetWorkflowResponse {
	this := GetWorkflowResponse{}
	this.Data = data
	return &this
}

// NewGetWorkflowResponseWithDefaults instantiates a new GetWorkflowResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetWorkflowResponseWithDefaults() *GetWorkflowResponse {
	this := GetWorkflowResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GetWorkflowResponse) GetData() Workflow {
	if o == nil {
		var ret Workflow
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetWorkflowResponse) GetDataOk() (*Workflow, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetWorkflowResponse) SetData(v Workflow) {
	o.Data = v
}

func (o GetWorkflowResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetWorkflowResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableGetWorkflowResponse struct {
	value *GetWorkflowResponse
	isSet bool
}

func (v NullableGetWorkflowResponse) Get() *GetWorkflowResponse {
	return v.value
}

func (v *NullableGetWorkflowResponse) Set(val *GetWorkflowResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetWorkflowResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetWorkflowResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetWorkflowResponse(val *GetWorkflowResponse) *NullableGetWorkflowResponse {
	return &NullableGetWorkflowResponse{value: val, isSet: true}
}

func (v NullableGetWorkflowResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetWorkflowResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


