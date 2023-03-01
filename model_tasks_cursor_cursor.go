/*
Formance Stack API

Open, modular foundation for unique payments flows  # Introduction This API is documented in **OpenAPI format**.  # Authentication Formance Stack offers one forms of authentication:   - OAuth2 OAuth2 - an open protocol to allow secure authorization in a simple and standard method from web, mobile and desktop applications. <SecurityDefinitions /> 

API version: v1.0.20230301
Contact: support@formance.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package formance

import (
	"encoding/json"
)

// checks if the TasksCursorCursor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TasksCursorCursor{}

// TasksCursorCursor struct for TasksCursorCursor
type TasksCursorCursor struct {
	PageSize int64 `json:"pageSize"`
	HasMore bool `json:"hasMore"`
	Previous *string `json:"previous,omitempty"`
	Next *string `json:"next,omitempty"`
	Data []TasksCursorCursorAllOfDataInner `json:"data"`
}

// NewTasksCursorCursor instantiates a new TasksCursorCursor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTasksCursorCursor(pageSize int64, hasMore bool, data []TasksCursorCursorAllOfDataInner) *TasksCursorCursor {
	this := TasksCursorCursor{}
	this.PageSize = pageSize
	this.HasMore = hasMore
	this.Data = data
	return &this
}

// NewTasksCursorCursorWithDefaults instantiates a new TasksCursorCursor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTasksCursorCursorWithDefaults() *TasksCursorCursor {
	this := TasksCursorCursor{}
	return &this
}

// GetPageSize returns the PageSize field value
func (o *TasksCursorCursor) GetPageSize() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value
// and a boolean to check if the value has been set.
func (o *TasksCursorCursor) GetPageSizeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageSize, true
}

// SetPageSize sets field value
func (o *TasksCursorCursor) SetPageSize(v int64) {
	o.PageSize = v
}

// GetHasMore returns the HasMore field value
func (o *TasksCursorCursor) GetHasMore() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value
// and a boolean to check if the value has been set.
func (o *TasksCursorCursor) GetHasMoreOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasMore, true
}

// SetHasMore sets field value
func (o *TasksCursorCursor) SetHasMore(v bool) {
	o.HasMore = v
}

// GetPrevious returns the Previous field value if set, zero value otherwise.
func (o *TasksCursorCursor) GetPrevious() string {
	if o == nil || IsNil(o.Previous) {
		var ret string
		return ret
	}
	return *o.Previous
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TasksCursorCursor) GetPreviousOk() (*string, bool) {
	if o == nil || IsNil(o.Previous) {
		return nil, false
	}
	return o.Previous, true
}

// HasPrevious returns a boolean if a field has been set.
func (o *TasksCursorCursor) HasPrevious() bool {
	if o != nil && !IsNil(o.Previous) {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given string and assigns it to the Previous field.
func (o *TasksCursorCursor) SetPrevious(v string) {
	o.Previous = &v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *TasksCursorCursor) GetNext() string {
	if o == nil || IsNil(o.Next) {
		var ret string
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TasksCursorCursor) GetNextOk() (*string, bool) {
	if o == nil || IsNil(o.Next) {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *TasksCursorCursor) HasNext() bool {
	if o != nil && !IsNil(o.Next) {
		return true
	}

	return false
}

// SetNext gets a reference to the given string and assigns it to the Next field.
func (o *TasksCursorCursor) SetNext(v string) {
	o.Next = &v
}

// GetData returns the Data field value
func (o *TasksCursorCursor) GetData() []TasksCursorCursorAllOfDataInner {
	if o == nil {
		var ret []TasksCursorCursorAllOfDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *TasksCursorCursor) GetDataOk() ([]TasksCursorCursorAllOfDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *TasksCursorCursor) SetData(v []TasksCursorCursorAllOfDataInner) {
	o.Data = v
}

func (o TasksCursorCursor) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TasksCursorCursor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pageSize"] = o.PageSize
	toSerialize["hasMore"] = o.HasMore
	if !IsNil(o.Previous) {
		toSerialize["previous"] = o.Previous
	}
	if !IsNil(o.Next) {
		toSerialize["next"] = o.Next
	}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableTasksCursorCursor struct {
	value *TasksCursorCursor
	isSet bool
}

func (v NullableTasksCursorCursor) Get() *TasksCursorCursor {
	return v.value
}

func (v *NullableTasksCursorCursor) Set(val *TasksCursorCursor) {
	v.value = val
	v.isSet = true
}

func (v NullableTasksCursorCursor) IsSet() bool {
	return v.isSet
}

func (v *NullableTasksCursorCursor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTasksCursorCursor(val *TasksCursorCursor) *NullableTasksCursorCursor {
	return &NullableTasksCursorCursor{value: val, isSet: true}
}

func (v NullableTasksCursorCursor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTasksCursorCursor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


