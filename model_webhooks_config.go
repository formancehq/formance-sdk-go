/*
Formance Stack API

Open, modular foundation for unique payments flows  # Introduction This API is documented in **OpenAPI format**.  # Authentication Formance Stack offers one forms of authentication:   - OAuth2 OAuth2 - an open protocol to allow secure authorization in a simple and standard method from web, mobile and desktop applications. <SecurityDefinitions /> 

API version: v0.2.7
Contact: support@formance.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package formance

import (
	"encoding/json"
	"time"
)

// checks if the WebhooksConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhooksConfig{}

// WebhooksConfig struct for WebhooksConfig
type WebhooksConfig struct {
	Endpoint *string `json:"endpoint,omitempty"`
	Secret *string `json:"secret,omitempty"`
	EventTypes []string `json:"eventTypes,omitempty"`
	Active *bool `json:"active,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	ModifiedAt *time.Time `json:"modifiedAt,omitempty"`
}

// NewWebhooksConfig instantiates a new WebhooksConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhooksConfig() *WebhooksConfig {
	this := WebhooksConfig{}
	return &this
}

// NewWebhooksConfigWithDefaults instantiates a new WebhooksConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhooksConfigWithDefaults() *WebhooksConfig {
	this := WebhooksConfig{}
	return &this
}

// GetEndpoint returns the Endpoint field value if set, zero value otherwise.
func (o *WebhooksConfig) GetEndpoint() string {
	if o == nil || isNil(o.Endpoint) {
		var ret string
		return ret
	}
	return *o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhooksConfig) GetEndpointOk() (*string, bool) {
	if o == nil || isNil(o.Endpoint) {
		return nil, false
	}
	return o.Endpoint, true
}

// HasEndpoint returns a boolean if a field has been set.
func (o *WebhooksConfig) HasEndpoint() bool {
	if o != nil && !isNil(o.Endpoint) {
		return true
	}

	return false
}

// SetEndpoint gets a reference to the given string and assigns it to the Endpoint field.
func (o *WebhooksConfig) SetEndpoint(v string) {
	o.Endpoint = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *WebhooksConfig) GetSecret() string {
	if o == nil || isNil(o.Secret) {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhooksConfig) GetSecretOk() (*string, bool) {
	if o == nil || isNil(o.Secret) {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *WebhooksConfig) HasSecret() bool {
	if o != nil && !isNil(o.Secret) {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *WebhooksConfig) SetSecret(v string) {
	o.Secret = &v
}

// GetEventTypes returns the EventTypes field value if set, zero value otherwise.
func (o *WebhooksConfig) GetEventTypes() []string {
	if o == nil || isNil(o.EventTypes) {
		var ret []string
		return ret
	}
	return o.EventTypes
}

// GetEventTypesOk returns a tuple with the EventTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhooksConfig) GetEventTypesOk() ([]string, bool) {
	if o == nil || isNil(o.EventTypes) {
		return nil, false
	}
	return o.EventTypes, true
}

// HasEventTypes returns a boolean if a field has been set.
func (o *WebhooksConfig) HasEventTypes() bool {
	if o != nil && !isNil(o.EventTypes) {
		return true
	}

	return false
}

// SetEventTypes gets a reference to the given []string and assigns it to the EventTypes field.
func (o *WebhooksConfig) SetEventTypes(v []string) {
	o.EventTypes = v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *WebhooksConfig) GetActive() bool {
	if o == nil || isNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhooksConfig) GetActiveOk() (*bool, bool) {
	if o == nil || isNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *WebhooksConfig) HasActive() bool {
	if o != nil && !isNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *WebhooksConfig) SetActive(v bool) {
	o.Active = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *WebhooksConfig) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhooksConfig) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *WebhooksConfig) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *WebhooksConfig) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *WebhooksConfig) GetModifiedAt() time.Time {
	if o == nil || isNil(o.ModifiedAt) {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhooksConfig) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.ModifiedAt) {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *WebhooksConfig) HasModifiedAt() bool {
	if o != nil && !isNil(o.ModifiedAt) {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.
func (o *WebhooksConfig) SetModifiedAt(v time.Time) {
	o.ModifiedAt = &v
}

func (o WebhooksConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhooksConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Endpoint) {
		toSerialize["endpoint"] = o.Endpoint
	}
	if !isNil(o.Secret) {
		toSerialize["secret"] = o.Secret
	}
	if !isNil(o.EventTypes) {
		toSerialize["eventTypes"] = o.EventTypes
	}
	if !isNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !isNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !isNil(o.ModifiedAt) {
		toSerialize["modifiedAt"] = o.ModifiedAt
	}
	return toSerialize, nil
}

type NullableWebhooksConfig struct {
	value *WebhooksConfig
	isSet bool
}

func (v NullableWebhooksConfig) Get() *WebhooksConfig {
	return v.value
}

func (v *NullableWebhooksConfig) Set(val *WebhooksConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhooksConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhooksConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhooksConfig(val *WebhooksConfig) *NullableWebhooksConfig {
	return &NullableWebhooksConfig{value: val, isSet: true}
}

func (v NullableWebhooksConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhooksConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


