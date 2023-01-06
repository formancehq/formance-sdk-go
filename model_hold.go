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

// Hold struct for Hold
type Hold struct {
	// The unique ID of the hold.
	Id string `json:"id"`
	// The ID of the wallet the hold is associated with.
	WalletID string `json:"walletID"`
	// Metadata associated with the hold.
	Metadata map[string]interface{} `json:"metadata"`
	Description string `json:"description"`
	Destination *Subject `json:"destination,omitempty"`
}

// NewHold instantiates a new Hold object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHold(id string, walletID string, metadata map[string]interface{}, description string) *Hold {
	this := Hold{}
	this.Id = id
	this.WalletID = walletID
	this.Metadata = metadata
	this.Description = description
	return &this
}

// NewHoldWithDefaults instantiates a new Hold object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHoldWithDefaults() *Hold {
	this := Hold{}
	return &this
}

// GetId returns the Id field value
func (o *Hold) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Hold) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Hold) SetId(v string) {
	o.Id = v
}

// GetWalletID returns the WalletID field value
func (o *Hold) GetWalletID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WalletID
}

// GetWalletIDOk returns a tuple with the WalletID field value
// and a boolean to check if the value has been set.
func (o *Hold) GetWalletIDOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.WalletID, true
}

// SetWalletID sets field value
func (o *Hold) SetWalletID(v string) {
	o.WalletID = v
}

// GetMetadata returns the Metadata field value
func (o *Hold) GetMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *Hold) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil {
    return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// SetMetadata sets field value
func (o *Hold) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetDescription returns the Description field value
func (o *Hold) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *Hold) GetDescriptionOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *Hold) SetDescription(v string) {
	o.Description = v
}

// GetDestination returns the Destination field value if set, zero value otherwise.
func (o *Hold) GetDestination() Subject {
	if o == nil || isNil(o.Destination) {
		var ret Subject
		return ret
	}
	return *o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Hold) GetDestinationOk() (*Subject, bool) {
	if o == nil || isNil(o.Destination) {
    return nil, false
	}
	return o.Destination, true
}

// HasDestination returns a boolean if a field has been set.
func (o *Hold) HasDestination() bool {
	if o != nil && !isNil(o.Destination) {
		return true
	}

	return false
}

// SetDestination gets a reference to the given Subject and assigns it to the Destination field.
func (o *Hold) SetDestination(v Subject) {
	o.Destination = &v
}

func (o Hold) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["walletID"] = o.WalletID
	}
	if true {
		toSerialize["metadata"] = o.Metadata
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.Destination) {
		toSerialize["destination"] = o.Destination
	}
	return json.Marshal(toSerialize)
}

type NullableHold struct {
	value *Hold
	isSet bool
}

func (v NullableHold) Get() *Hold {
	return v.value
}

func (v *NullableHold) Set(val *Hold) {
	v.value = val
	v.isSet = true
}

func (v NullableHold) IsSet() bool {
	return v.isSet
}

func (v *NullableHold) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHold(val *Hold) *NullableHold {
	return &NullableHold{value: val, isSet: true}
}

func (v NullableHold) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHold) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


