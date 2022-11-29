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
	"time"
)

// checks if the StripeTask type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StripeTask{}

// StripeTask struct for StripeTask
type StripeTask struct {
	// The id of the oldest BalanceTransaction fetched from stripe for this account
	OldestId *string `json:"oldestId,omitempty"`
	// The creation date of the oldest BalanceTransaction fetched from stripe for this account
	OldestDate *time.Time `json:"oldestDate,omitempty"`
	// The id of the more recent BalanceTransaction fetched from stripe for this account
	MoreRecentId *string `json:"moreRecentId,omitempty"`
	// The creation date of the more recent BalanceTransaction fetched from stripe for this account
	MoreRecentDate *time.Time `json:"moreRecentDate,omitempty"`
	NoMoreHistory *bool `json:"noMoreHistory,omitempty"`
}

// NewStripeTask instantiates a new StripeTask object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStripeTask() *StripeTask {
	this := StripeTask{}
	return &this
}

// NewStripeTaskWithDefaults instantiates a new StripeTask object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStripeTaskWithDefaults() *StripeTask {
	this := StripeTask{}
	return &this
}

// GetOldestId returns the OldestId field value if set, zero value otherwise.
func (o *StripeTask) GetOldestId() string {
	if o == nil || isNil(o.OldestId) {
		var ret string
		return ret
	}
	return *o.OldestId
}

// GetOldestIdOk returns a tuple with the OldestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StripeTask) GetOldestIdOk() (*string, bool) {
	if o == nil || isNil(o.OldestId) {
		return nil, false
	}
	return o.OldestId, true
}

// HasOldestId returns a boolean if a field has been set.
func (o *StripeTask) HasOldestId() bool {
	if o != nil && !isNil(o.OldestId) {
		return true
	}

	return false
}

// SetOldestId gets a reference to the given string and assigns it to the OldestId field.
func (o *StripeTask) SetOldestId(v string) {
	o.OldestId = &v
}

// GetOldestDate returns the OldestDate field value if set, zero value otherwise.
func (o *StripeTask) GetOldestDate() time.Time {
	if o == nil || isNil(o.OldestDate) {
		var ret time.Time
		return ret
	}
	return *o.OldestDate
}

// GetOldestDateOk returns a tuple with the OldestDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StripeTask) GetOldestDateOk() (*time.Time, bool) {
	if o == nil || isNil(o.OldestDate) {
		return nil, false
	}
	return o.OldestDate, true
}

// HasOldestDate returns a boolean if a field has been set.
func (o *StripeTask) HasOldestDate() bool {
	if o != nil && !isNil(o.OldestDate) {
		return true
	}

	return false
}

// SetOldestDate gets a reference to the given time.Time and assigns it to the OldestDate field.
func (o *StripeTask) SetOldestDate(v time.Time) {
	o.OldestDate = &v
}

// GetMoreRecentId returns the MoreRecentId field value if set, zero value otherwise.
func (o *StripeTask) GetMoreRecentId() string {
	if o == nil || isNil(o.MoreRecentId) {
		var ret string
		return ret
	}
	return *o.MoreRecentId
}

// GetMoreRecentIdOk returns a tuple with the MoreRecentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StripeTask) GetMoreRecentIdOk() (*string, bool) {
	if o == nil || isNil(o.MoreRecentId) {
		return nil, false
	}
	return o.MoreRecentId, true
}

// HasMoreRecentId returns a boolean if a field has been set.
func (o *StripeTask) HasMoreRecentId() bool {
	if o != nil && !isNil(o.MoreRecentId) {
		return true
	}

	return false
}

// SetMoreRecentId gets a reference to the given string and assigns it to the MoreRecentId field.
func (o *StripeTask) SetMoreRecentId(v string) {
	o.MoreRecentId = &v
}

// GetMoreRecentDate returns the MoreRecentDate field value if set, zero value otherwise.
func (o *StripeTask) GetMoreRecentDate() time.Time {
	if o == nil || isNil(o.MoreRecentDate) {
		var ret time.Time
		return ret
	}
	return *o.MoreRecentDate
}

// GetMoreRecentDateOk returns a tuple with the MoreRecentDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StripeTask) GetMoreRecentDateOk() (*time.Time, bool) {
	if o == nil || isNil(o.MoreRecentDate) {
		return nil, false
	}
	return o.MoreRecentDate, true
}

// HasMoreRecentDate returns a boolean if a field has been set.
func (o *StripeTask) HasMoreRecentDate() bool {
	if o != nil && !isNil(o.MoreRecentDate) {
		return true
	}

	return false
}

// SetMoreRecentDate gets a reference to the given time.Time and assigns it to the MoreRecentDate field.
func (o *StripeTask) SetMoreRecentDate(v time.Time) {
	o.MoreRecentDate = &v
}

// GetNoMoreHistory returns the NoMoreHistory field value if set, zero value otherwise.
func (o *StripeTask) GetNoMoreHistory() bool {
	if o == nil || isNil(o.NoMoreHistory) {
		var ret bool
		return ret
	}
	return *o.NoMoreHistory
}

// GetNoMoreHistoryOk returns a tuple with the NoMoreHistory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StripeTask) GetNoMoreHistoryOk() (*bool, bool) {
	if o == nil || isNil(o.NoMoreHistory) {
		return nil, false
	}
	return o.NoMoreHistory, true
}

// HasNoMoreHistory returns a boolean if a field has been set.
func (o *StripeTask) HasNoMoreHistory() bool {
	if o != nil && !isNil(o.NoMoreHistory) {
		return true
	}

	return false
}

// SetNoMoreHistory gets a reference to the given bool and assigns it to the NoMoreHistory field.
func (o *StripeTask) SetNoMoreHistory(v bool) {
	o.NoMoreHistory = &v
}

func (o StripeTask) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StripeTask) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.OldestId) {
		toSerialize["oldestId"] = o.OldestId
	}
	if !isNil(o.OldestDate) {
		toSerialize["oldestDate"] = o.OldestDate
	}
	if !isNil(o.MoreRecentId) {
		toSerialize["moreRecentId"] = o.MoreRecentId
	}
	if !isNil(o.MoreRecentDate) {
		toSerialize["moreRecentDate"] = o.MoreRecentDate
	}
	if !isNil(o.NoMoreHistory) {
		toSerialize["noMoreHistory"] = o.NoMoreHistory
	}
	return toSerialize, nil
}

type NullableStripeTask struct {
	value *StripeTask
	isSet bool
}

func (v NullableStripeTask) Get() *StripeTask {
	return v.value
}

func (v *NullableStripeTask) Set(val *StripeTask) {
	v.value = val
	v.isSet = true
}

func (v NullableStripeTask) IsSet() bool {
	return v.isSet
}

func (v *NullableStripeTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStripeTask(val *StripeTask) *NullableStripeTask {
	return &NullableStripeTask{value: val, isSet: true}
}

func (v NullableStripeTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStripeTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


