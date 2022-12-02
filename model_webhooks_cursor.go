/*
Formance Stack API

Open, modular foundation for unique payments flows  # Introduction This API is documented in **OpenAPI format**.  # Authentication Formance Stack offers one forms of authentication:   - OAuth2 OAuth2 - an open protocol to allow secure authorization in a simple and standard method from web, mobile and desktop applications. <SecurityDefinitions /> 

API version: v0.2.6
Contact: support@formance.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package formance

import (
	"encoding/json"
)

// checks if the WebhooksCursor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhooksCursor{}

// WebhooksCursor struct for WebhooksCursor
type WebhooksCursor struct {
	HasMore *bool `json:"has_more,omitempty"`
}

// NewWebhooksCursor instantiates a new WebhooksCursor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhooksCursor() *WebhooksCursor {
	this := WebhooksCursor{}
	return &this
}

// NewWebhooksCursorWithDefaults instantiates a new WebhooksCursor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhooksCursorWithDefaults() *WebhooksCursor {
	this := WebhooksCursor{}
	return &this
}

// GetHasMore returns the HasMore field value if set, zero value otherwise.
func (o *WebhooksCursor) GetHasMore() bool {
	if o == nil || isNil(o.HasMore) {
		var ret bool
		return ret
	}
	return *o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhooksCursor) GetHasMoreOk() (*bool, bool) {
	if o == nil || isNil(o.HasMore) {
		return nil, false
	}
	return o.HasMore, true
}

// HasHasMore returns a boolean if a field has been set.
func (o *WebhooksCursor) HasHasMore() bool {
	if o != nil && !isNil(o.HasMore) {
		return true
	}

	return false
}

// SetHasMore gets a reference to the given bool and assigns it to the HasMore field.
func (o *WebhooksCursor) SetHasMore(v bool) {
	o.HasMore = &v
}

func (o WebhooksCursor) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhooksCursor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.HasMore) {
		toSerialize["has_more"] = o.HasMore
	}
	return toSerialize, nil
}

type NullableWebhooksCursor struct {
	value *WebhooksCursor
	isSet bool
}

func (v NullableWebhooksCursor) Get() *WebhooksCursor {
	return v.value
}

func (v *NullableWebhooksCursor) Set(val *WebhooksCursor) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhooksCursor) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhooksCursor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhooksCursor(val *WebhooksCursor) *NullableWebhooksCursor {
	return &NullableWebhooksCursor{value: val, isSet: true}
}

func (v NullableWebhooksCursor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhooksCursor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


