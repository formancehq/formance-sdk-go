/*
Formance Stack API

Open, modular foundation for unique payments flows  # Introduction This API is documented in **OpenAPI format**.  # Authentication Formance Stack offers one forms of authentication:   - OAuth2 OAuth2 - an open protocol to allow secure authorization in a simple and standard method from web, mobile and desktop applications. <SecurityDefinitions /> 

API version: v1.0.0
Contact: support@formance.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package formance

import (
	"encoding/json"
)

// checks if the CurrencyCloudConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CurrencyCloudConfig{}

// CurrencyCloudConfig struct for CurrencyCloudConfig
type CurrencyCloudConfig struct {
	ApiKey string `json:"apiKey"`
	// Username of the API Key holder
	LoginID string `json:"loginID"`
	// The frequency at which the connector will fetch transactions
	PollingPeriod *string `json:"pollingPeriod,omitempty"`
	// The endpoint to use for the API. Defaults to https://devapi.currencycloud.com
	Endpoint *string `json:"endpoint,omitempty"`
}

// NewCurrencyCloudConfig instantiates a new CurrencyCloudConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCurrencyCloudConfig(apiKey string, loginID string) *CurrencyCloudConfig {
	this := CurrencyCloudConfig{}
	this.ApiKey = apiKey
	this.LoginID = loginID
	return &this
}

// NewCurrencyCloudConfigWithDefaults instantiates a new CurrencyCloudConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCurrencyCloudConfigWithDefaults() *CurrencyCloudConfig {
	this := CurrencyCloudConfig{}
	return &this
}

// GetApiKey returns the ApiKey field value
func (o *CurrencyCloudConfig) GetApiKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value
// and a boolean to check if the value has been set.
func (o *CurrencyCloudConfig) GetApiKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiKey, true
}

// SetApiKey sets field value
func (o *CurrencyCloudConfig) SetApiKey(v string) {
	o.ApiKey = v
}

// GetLoginID returns the LoginID field value
func (o *CurrencyCloudConfig) GetLoginID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LoginID
}

// GetLoginIDOk returns a tuple with the LoginID field value
// and a boolean to check if the value has been set.
func (o *CurrencyCloudConfig) GetLoginIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LoginID, true
}

// SetLoginID sets field value
func (o *CurrencyCloudConfig) SetLoginID(v string) {
	o.LoginID = v
}

// GetPollingPeriod returns the PollingPeriod field value if set, zero value otherwise.
func (o *CurrencyCloudConfig) GetPollingPeriod() string {
	if o == nil || isNil(o.PollingPeriod) {
		var ret string
		return ret
	}
	return *o.PollingPeriod
}

// GetPollingPeriodOk returns a tuple with the PollingPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrencyCloudConfig) GetPollingPeriodOk() (*string, bool) {
	if o == nil || isNil(o.PollingPeriod) {
		return nil, false
	}
	return o.PollingPeriod, true
}

// HasPollingPeriod returns a boolean if a field has been set.
func (o *CurrencyCloudConfig) HasPollingPeriod() bool {
	if o != nil && !isNil(o.PollingPeriod) {
		return true
	}

	return false
}

// SetPollingPeriod gets a reference to the given string and assigns it to the PollingPeriod field.
func (o *CurrencyCloudConfig) SetPollingPeriod(v string) {
	o.PollingPeriod = &v
}

// GetEndpoint returns the Endpoint field value if set, zero value otherwise.
func (o *CurrencyCloudConfig) GetEndpoint() string {
	if o == nil || isNil(o.Endpoint) {
		var ret string
		return ret
	}
	return *o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrencyCloudConfig) GetEndpointOk() (*string, bool) {
	if o == nil || isNil(o.Endpoint) {
		return nil, false
	}
	return o.Endpoint, true
}

// HasEndpoint returns a boolean if a field has been set.
func (o *CurrencyCloudConfig) HasEndpoint() bool {
	if o != nil && !isNil(o.Endpoint) {
		return true
	}

	return false
}

// SetEndpoint gets a reference to the given string and assigns it to the Endpoint field.
func (o *CurrencyCloudConfig) SetEndpoint(v string) {
	o.Endpoint = &v
}

func (o CurrencyCloudConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CurrencyCloudConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["apiKey"] = o.ApiKey
	toSerialize["loginID"] = o.LoginID
	if !isNil(o.PollingPeriod) {
		toSerialize["pollingPeriod"] = o.PollingPeriod
	}
	if !isNil(o.Endpoint) {
		toSerialize["endpoint"] = o.Endpoint
	}
	return toSerialize, nil
}

type NullableCurrencyCloudConfig struct {
	value *CurrencyCloudConfig
	isSet bool
}

func (v NullableCurrencyCloudConfig) Get() *CurrencyCloudConfig {
	return v.value
}

func (v *NullableCurrencyCloudConfig) Set(val *CurrencyCloudConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableCurrencyCloudConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableCurrencyCloudConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCurrencyCloudConfig(val *CurrencyCloudConfig) *NullableCurrencyCloudConfig {
	return &NullableCurrencyCloudConfig{value: val, isSet: true}
}

func (v NullableCurrencyCloudConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCurrencyCloudConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


