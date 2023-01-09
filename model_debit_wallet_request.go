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

// DebitWalletRequest struct for DebitWalletRequest
type DebitWalletRequest struct {
	Amount Monetary `json:"amount"`
	// Set to true to create a pending hold. If false, the wallet will be debited immediately.
	Pending *bool `json:"pending,omitempty"`
	// Metadata associated with the wallet.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Description *string `json:"description,omitempty"`
	Destination *Subject `json:"destination,omitempty"`
	Balances []string `json:"balances,omitempty"`
}

// NewDebitWalletRequest instantiates a new DebitWalletRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDebitWalletRequest(amount Monetary) *DebitWalletRequest {
	this := DebitWalletRequest{}
	this.Amount = amount
	return &this
}

// NewDebitWalletRequestWithDefaults instantiates a new DebitWalletRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDebitWalletRequestWithDefaults() *DebitWalletRequest {
	this := DebitWalletRequest{}
	return &this
}

// GetAmount returns the Amount field value
func (o *DebitWalletRequest) GetAmount() Monetary {
	if o == nil {
		var ret Monetary
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *DebitWalletRequest) GetAmountOk() (*Monetary, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *DebitWalletRequest) SetAmount(v Monetary) {
	o.Amount = v
}

// GetPending returns the Pending field value if set, zero value otherwise.
func (o *DebitWalletRequest) GetPending() bool {
	if o == nil || isNil(o.Pending) {
		var ret bool
		return ret
	}
	return *o.Pending
}

// GetPendingOk returns a tuple with the Pending field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DebitWalletRequest) GetPendingOk() (*bool, bool) {
	if o == nil || isNil(o.Pending) {
    return nil, false
	}
	return o.Pending, true
}

// HasPending returns a boolean if a field has been set.
func (o *DebitWalletRequest) HasPending() bool {
	if o != nil && !isNil(o.Pending) {
		return true
	}

	return false
}

// SetPending gets a reference to the given bool and assigns it to the Pending field.
func (o *DebitWalletRequest) SetPending(v bool) {
	o.Pending = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *DebitWalletRequest) GetMetadata() map[string]interface{} {
	if o == nil || isNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DebitWalletRequest) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Metadata) {
    return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *DebitWalletRequest) HasMetadata() bool {
	if o != nil && !isNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *DebitWalletRequest) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DebitWalletRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DebitWalletRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DebitWalletRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DebitWalletRequest) SetDescription(v string) {
	o.Description = &v
}

// GetDestination returns the Destination field value if set, zero value otherwise.
func (o *DebitWalletRequest) GetDestination() Subject {
	if o == nil || isNil(o.Destination) {
		var ret Subject
		return ret
	}
	return *o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DebitWalletRequest) GetDestinationOk() (*Subject, bool) {
	if o == nil || isNil(o.Destination) {
    return nil, false
	}
	return o.Destination, true
}

// HasDestination returns a boolean if a field has been set.
func (o *DebitWalletRequest) HasDestination() bool {
	if o != nil && !isNil(o.Destination) {
		return true
	}

	return false
}

// SetDestination gets a reference to the given Subject and assigns it to the Destination field.
func (o *DebitWalletRequest) SetDestination(v Subject) {
	o.Destination = &v
}

// GetBalances returns the Balances field value if set, zero value otherwise.
func (o *DebitWalletRequest) GetBalances() []string {
	if o == nil || isNil(o.Balances) {
		var ret []string
		return ret
	}
	return o.Balances
}

// GetBalancesOk returns a tuple with the Balances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DebitWalletRequest) GetBalancesOk() ([]string, bool) {
	if o == nil || isNil(o.Balances) {
    return nil, false
	}
	return o.Balances, true
}

// HasBalances returns a boolean if a field has been set.
func (o *DebitWalletRequest) HasBalances() bool {
	if o != nil && !isNil(o.Balances) {
		return true
	}

	return false
}

// SetBalances gets a reference to the given []string and assigns it to the Balances field.
func (o *DebitWalletRequest) SetBalances(v []string) {
	o.Balances = v
}

func (o DebitWalletRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if !isNil(o.Pending) {
		toSerialize["pending"] = o.Pending
	}
	if !isNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.Destination) {
		toSerialize["destination"] = o.Destination
	}
	if !isNil(o.Balances) {
		toSerialize["balances"] = o.Balances
	}
	return json.Marshal(toSerialize)
}

type NullableDebitWalletRequest struct {
	value *DebitWalletRequest
	isSet bool
}

func (v NullableDebitWalletRequest) Get() *DebitWalletRequest {
	return v.value
}

func (v *NullableDebitWalletRequest) Set(val *DebitWalletRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDebitWalletRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDebitWalletRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDebitWalletRequest(val *DebitWalletRequest) *NullableDebitWalletRequest {
	return &NullableDebitWalletRequest{value: val, isSet: true}
}

func (v NullableDebitWalletRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDebitWalletRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


