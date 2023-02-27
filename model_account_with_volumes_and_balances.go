/*
Formance Stack API

Open, modular foundation for unique payments flows  # Introduction This API is documented in **OpenAPI format**.  # Authentication Formance Stack offers one forms of authentication:   - OAuth2 OAuth2 - an open protocol to allow secure authorization in a simple and standard method from web, mobile and desktop applications. <SecurityDefinitions /> 

API version: develop
Contact: support@formance.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package formance

import (
	"encoding/json"
)

// checks if the AccountWithVolumesAndBalances type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountWithVolumesAndBalances{}

// AccountWithVolumesAndBalances struct for AccountWithVolumesAndBalances
type AccountWithVolumesAndBalances struct {
	Address string `json:"address"`
	Type *string `json:"type,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Volumes *map[string]map[string]int64 `json:"volumes,omitempty"`
	Balances *map[string]int64 `json:"balances,omitempty"`
}

// NewAccountWithVolumesAndBalances instantiates a new AccountWithVolumesAndBalances object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountWithVolumesAndBalances(address string) *AccountWithVolumesAndBalances {
	this := AccountWithVolumesAndBalances{}
	this.Address = address
	return &this
}

// NewAccountWithVolumesAndBalancesWithDefaults instantiates a new AccountWithVolumesAndBalances object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountWithVolumesAndBalancesWithDefaults() *AccountWithVolumesAndBalances {
	this := AccountWithVolumesAndBalances{}
	return &this
}

// GetAddress returns the Address field value
func (o *AccountWithVolumesAndBalances) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *AccountWithVolumesAndBalances) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *AccountWithVolumesAndBalances) SetAddress(v string) {
	o.Address = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AccountWithVolumesAndBalances) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountWithVolumesAndBalances) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AccountWithVolumesAndBalances) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AccountWithVolumesAndBalances) SetType(v string) {
	o.Type = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *AccountWithVolumesAndBalances) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountWithVolumesAndBalances) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *AccountWithVolumesAndBalances) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *AccountWithVolumesAndBalances) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetVolumes returns the Volumes field value if set, zero value otherwise.
func (o *AccountWithVolumesAndBalances) GetVolumes() map[string]map[string]int64 {
	if o == nil || IsNil(o.Volumes) {
		var ret map[string]map[string]int64
		return ret
	}
	return *o.Volumes
}

// GetVolumesOk returns a tuple with the Volumes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountWithVolumesAndBalances) GetVolumesOk() (*map[string]map[string]int64, bool) {
	if o == nil || IsNil(o.Volumes) {
		return nil, false
	}
	return o.Volumes, true
}

// HasVolumes returns a boolean if a field has been set.
func (o *AccountWithVolumesAndBalances) HasVolumes() bool {
	if o != nil && !IsNil(o.Volumes) {
		return true
	}

	return false
}

// SetVolumes gets a reference to the given map[string]map[string]int64 and assigns it to the Volumes field.
func (o *AccountWithVolumesAndBalances) SetVolumes(v map[string]map[string]int64) {
	o.Volumes = &v
}

// GetBalances returns the Balances field value if set, zero value otherwise.
func (o *AccountWithVolumesAndBalances) GetBalances() map[string]int64 {
	if o == nil || IsNil(o.Balances) {
		var ret map[string]int64
		return ret
	}
	return *o.Balances
}

// GetBalancesOk returns a tuple with the Balances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountWithVolumesAndBalances) GetBalancesOk() (*map[string]int64, bool) {
	if o == nil || IsNil(o.Balances) {
		return nil, false
	}
	return o.Balances, true
}

// HasBalances returns a boolean if a field has been set.
func (o *AccountWithVolumesAndBalances) HasBalances() bool {
	if o != nil && !IsNil(o.Balances) {
		return true
	}

	return false
}

// SetBalances gets a reference to the given map[string]int64 and assigns it to the Balances field.
func (o *AccountWithVolumesAndBalances) SetBalances(v map[string]int64) {
	o.Balances = &v
}

func (o AccountWithVolumesAndBalances) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountWithVolumesAndBalances) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["address"] = o.Address
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.Volumes) {
		toSerialize["volumes"] = o.Volumes
	}
	if !IsNil(o.Balances) {
		toSerialize["balances"] = o.Balances
	}
	return toSerialize, nil
}

type NullableAccountWithVolumesAndBalances struct {
	value *AccountWithVolumesAndBalances
	isSet bool
}

func (v NullableAccountWithVolumesAndBalances) Get() *AccountWithVolumesAndBalances {
	return v.value
}

func (v *NullableAccountWithVolumesAndBalances) Set(val *AccountWithVolumesAndBalances) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountWithVolumesAndBalances) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountWithVolumesAndBalances) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountWithVolumesAndBalances(val *AccountWithVolumesAndBalances) *NullableAccountWithVolumesAndBalances {
	return &NullableAccountWithVolumesAndBalances{value: val, isSet: true}
}

func (v NullableAccountWithVolumesAndBalances) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountWithVolumesAndBalances) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


