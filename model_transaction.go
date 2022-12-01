/*
Formance Stack API

Open, modular foundation for unique payments flows  # Introduction This API is documented in **OpenAPI format**.  # Authentication Formance Stack offers one forms of authentication:   - OAuth2 OAuth2 - an open protocol to allow secure authorization in a simple and standard method from web, mobile and desktop applications. <SecurityDefinitions /> 

API version: v0.2.4
Contact: support@formance.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package formance

import (
	"encoding/json"
	"time"
)

// checks if the Transaction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Transaction{}

// Transaction struct for Transaction
type Transaction struct {
	Timestamp time.Time `json:"timestamp"`
	Postings []Posting `json:"postings"`
	Reference *string `json:"reference,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Txid int32 `json:"txid"`
	PreCommitVolumes *map[string]map[string]Volume `json:"preCommitVolumes,omitempty"`
	PostCommitVolumes *map[string]map[string]Volume `json:"postCommitVolumes,omitempty"`
}

// NewTransaction instantiates a new Transaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransaction(timestamp time.Time, postings []Posting, txid int32) *Transaction {
	this := Transaction{}
	this.Timestamp = timestamp
	this.Postings = postings
	this.Txid = txid
	return &this
}

// NewTransactionWithDefaults instantiates a new Transaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionWithDefaults() *Transaction {
	this := Transaction{}
	return &this
}

// GetTimestamp returns the Timestamp field value
func (o *Transaction) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *Transaction) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

// GetPostings returns the Postings field value
func (o *Transaction) GetPostings() []Posting {
	if o == nil {
		var ret []Posting
		return ret
	}

	return o.Postings
}

// GetPostingsOk returns a tuple with the Postings field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetPostingsOk() ([]Posting, bool) {
	if o == nil {
		return nil, false
	}
	return o.Postings, true
}

// SetPostings sets field value
func (o *Transaction) SetPostings(v []Posting) {
	o.Postings = v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *Transaction) GetReference() string {
	if o == nil || isNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetReferenceOk() (*string, bool) {
	if o == nil || isNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *Transaction) HasReference() bool {
	if o != nil && !isNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *Transaction) SetReference(v string) {
	o.Reference = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Transaction) GetMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Transaction) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Transaction) HasMetadata() bool {
	if o != nil && isNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *Transaction) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetTxid returns the Txid field value
func (o *Transaction) GetTxid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Txid
}

// GetTxidOk returns a tuple with the Txid field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetTxidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Txid, true
}

// SetTxid sets field value
func (o *Transaction) SetTxid(v int32) {
	o.Txid = v
}

// GetPreCommitVolumes returns the PreCommitVolumes field value if set, zero value otherwise.
func (o *Transaction) GetPreCommitVolumes() map[string]map[string]Volume {
	if o == nil || isNil(o.PreCommitVolumes) {
		var ret map[string]map[string]Volume
		return ret
	}
	return *o.PreCommitVolumes
}

// GetPreCommitVolumesOk returns a tuple with the PreCommitVolumes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetPreCommitVolumesOk() (*map[string]map[string]Volume, bool) {
	if o == nil || isNil(o.PreCommitVolumes) {
		return nil, false
	}
	return o.PreCommitVolumes, true
}

// HasPreCommitVolumes returns a boolean if a field has been set.
func (o *Transaction) HasPreCommitVolumes() bool {
	if o != nil && !isNil(o.PreCommitVolumes) {
		return true
	}

	return false
}

// SetPreCommitVolumes gets a reference to the given map[string]map[string]Volume and assigns it to the PreCommitVolumes field.
func (o *Transaction) SetPreCommitVolumes(v map[string]map[string]Volume) {
	o.PreCommitVolumes = &v
}

// GetPostCommitVolumes returns the PostCommitVolumes field value if set, zero value otherwise.
func (o *Transaction) GetPostCommitVolumes() map[string]map[string]Volume {
	if o == nil || isNil(o.PostCommitVolumes) {
		var ret map[string]map[string]Volume
		return ret
	}
	return *o.PostCommitVolumes
}

// GetPostCommitVolumesOk returns a tuple with the PostCommitVolumes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetPostCommitVolumesOk() (*map[string]map[string]Volume, bool) {
	if o == nil || isNil(o.PostCommitVolumes) {
		return nil, false
	}
	return o.PostCommitVolumes, true
}

// HasPostCommitVolumes returns a boolean if a field has been set.
func (o *Transaction) HasPostCommitVolumes() bool {
	if o != nil && !isNil(o.PostCommitVolumes) {
		return true
	}

	return false
}

// SetPostCommitVolumes gets a reference to the given map[string]map[string]Volume and assigns it to the PostCommitVolumes field.
func (o *Transaction) SetPostCommitVolumes(v map[string]map[string]Volume) {
	o.PostCommitVolumes = &v
}

func (o Transaction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Transaction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["timestamp"] = o.Timestamp
	toSerialize["postings"] = o.Postings
	if !isNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	toSerialize["txid"] = o.Txid
	if !isNil(o.PreCommitVolumes) {
		toSerialize["preCommitVolumes"] = o.PreCommitVolumes
	}
	if !isNil(o.PostCommitVolumes) {
		toSerialize["postCommitVolumes"] = o.PostCommitVolumes
	}
	return toSerialize, nil
}

type NullableTransaction struct {
	value *Transaction
	isSet bool
}

func (v NullableTransaction) Get() *Transaction {
	return v.value
}

func (v *NullableTransaction) Set(val *Transaction) {
	v.value = val
	v.isSet = true
}

func (v NullableTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransaction(val *Transaction) *NullableTransaction {
	return &NullableTransaction{value: val, isSet: true}
}

func (v NullableTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


