/*
Formance Stack API

Open, modular foundation for unique payments flows  # Introduction This API is documented in **OpenAPI format**.  # Authentication Formance Stack offers one forms of authentication:   - OAuth2 OAuth2 - an open protocol to allow secure authorization in a simple and standard method from web, mobile and desktop applications. <SecurityDefinitions /> 

API version: v1.0.0-beta.4
Contact: support@formance.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package formance

import (
	"encoding/json"
)

// checks if the Posting type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Posting{}

// Posting struct for Posting
type Posting struct {
	Amount int32 `json:"amount"`
	Asset string `json:"asset"`
	Destination string `json:"destination"`
	Source string `json:"source"`
}

// NewPosting instantiates a new Posting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPosting(amount int32, asset string, destination string, source string) *Posting {
	this := Posting{}
	this.Amount = amount
	this.Asset = asset
	this.Destination = destination
	this.Source = source
	return &this
}

// NewPostingWithDefaults instantiates a new Posting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostingWithDefaults() *Posting {
	this := Posting{}
	return &this
}

// GetAmount returns the Amount field value
func (o *Posting) GetAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *Posting) GetAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *Posting) SetAmount(v int32) {
	o.Amount = v
}

// GetAsset returns the Asset field value
func (o *Posting) GetAsset() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Asset
}

// GetAssetOk returns a tuple with the Asset field value
// and a boolean to check if the value has been set.
func (o *Posting) GetAssetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Asset, true
}

// SetAsset sets field value
func (o *Posting) SetAsset(v string) {
	o.Asset = v
}

// GetDestination returns the Destination field value
func (o *Posting) GetDestination() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value
// and a boolean to check if the value has been set.
func (o *Posting) GetDestinationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Destination, true
}

// SetDestination sets field value
func (o *Posting) SetDestination(v string) {
	o.Destination = v
}

// GetSource returns the Source field value
func (o *Posting) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *Posting) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *Posting) SetSource(v string) {
	o.Source = v
}

func (o Posting) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Posting) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	toSerialize["asset"] = o.Asset
	toSerialize["destination"] = o.Destination
	toSerialize["source"] = o.Source
	return toSerialize, nil
}

type NullablePosting struct {
	value *Posting
	isSet bool
}

func (v NullablePosting) Get() *Posting {
	return v.value
}

func (v *NullablePosting) Set(val *Posting) {
	v.value = val
	v.isSet = true
}

func (v NullablePosting) IsSet() bool {
	return v.isSet
}

func (v *NullablePosting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePosting(val *Posting) *NullablePosting {
	return &NullablePosting{value: val, isSet: true}
}

func (v NullablePosting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePosting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


