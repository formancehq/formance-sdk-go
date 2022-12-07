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

// checks if the Stats type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Stats{}

// Stats struct for Stats
type Stats struct {
	Accounts int32 `json:"accounts"`
	Transactions int32 `json:"transactions"`
}

// NewStats instantiates a new Stats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStats(accounts int32, transactions int32) *Stats {
	this := Stats{}
	this.Accounts = accounts
	this.Transactions = transactions
	return &this
}

// NewStatsWithDefaults instantiates a new Stats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatsWithDefaults() *Stats {
	this := Stats{}
	return &this
}

// GetAccounts returns the Accounts field value
func (o *Stats) GetAccounts() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value
// and a boolean to check if the value has been set.
func (o *Stats) GetAccountsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Accounts, true
}

// SetAccounts sets field value
func (o *Stats) SetAccounts(v int32) {
	o.Accounts = v
}

// GetTransactions returns the Transactions field value
func (o *Stats) GetTransactions() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Transactions
}

// GetTransactionsOk returns a tuple with the Transactions field value
// and a boolean to check if the value has been set.
func (o *Stats) GetTransactionsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Transactions, true
}

// SetTransactions sets field value
func (o *Stats) SetTransactions(v int32) {
	o.Transactions = v
}

func (o Stats) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Stats) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accounts"] = o.Accounts
	toSerialize["transactions"] = o.Transactions
	return toSerialize, nil
}

type NullableStats struct {
	value *Stats
	isSet bool
}

func (v NullableStats) Get() *Stats {
	return v.value
}

func (v *NullableStats) Set(val *Stats) {
	v.value = val
	v.isSet = true
}

func (v NullableStats) IsSet() bool {
	return v.isSet
}

func (v *NullableStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStats(val *Stats) *NullableStats {
	return &NullableStats{value: val, isSet: true}
}

func (v NullableStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


