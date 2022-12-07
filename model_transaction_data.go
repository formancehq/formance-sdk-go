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
	"time"
)

// checks if the TransactionData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionData{}

// TransactionData struct for TransactionData
type TransactionData struct {
	Timestamp *time.Time `json:"timestamp,omitempty"`
	Postings []Posting `json:"postings"`
	Reference *string `json:"reference,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewTransactionData instantiates a new TransactionData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionData(postings []Posting) *TransactionData {
	this := TransactionData{}
	this.Postings = postings
	return &this
}

// NewTransactionDataWithDefaults instantiates a new TransactionData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionDataWithDefaults() *TransactionData {
	this := TransactionData{}
	return &this
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *TransactionData) GetTimestamp() time.Time {
	if o == nil || isNil(o.Timestamp) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionData) GetTimestampOk() (*time.Time, bool) {
	if o == nil || isNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *TransactionData) HasTimestamp() bool {
	if o != nil && !isNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *TransactionData) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetPostings returns the Postings field value
func (o *TransactionData) GetPostings() []Posting {
	if o == nil {
		var ret []Posting
		return ret
	}

	return o.Postings
}

// GetPostingsOk returns a tuple with the Postings field value
// and a boolean to check if the value has been set.
func (o *TransactionData) GetPostingsOk() ([]Posting, bool) {
	if o == nil {
		return nil, false
	}
	return o.Postings, true
}

// SetPostings sets field value
func (o *TransactionData) SetPostings(v []Posting) {
	o.Postings = v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *TransactionData) GetReference() string {
	if o == nil || isNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionData) GetReferenceOk() (*string, bool) {
	if o == nil || isNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *TransactionData) HasReference() bool {
	if o != nil && !isNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *TransactionData) SetReference(v string) {
	o.Reference = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransactionData) GetMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionData) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *TransactionData) HasMetadata() bool {
	if o != nil && isNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *TransactionData) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o TransactionData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	toSerialize["postings"] = o.Postings
	if !isNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

type NullableTransactionData struct {
	value *TransactionData
	isSet bool
}

func (v NullableTransactionData) Get() *TransactionData {
	return v.value
}

func (v *NullableTransactionData) Set(val *TransactionData) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionData) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionData(val *TransactionData) *NullableTransactionData {
	return &NullableTransactionData{value: val, isSet: true}
}

func (v NullableTransactionData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


