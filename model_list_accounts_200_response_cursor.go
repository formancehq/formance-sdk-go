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

// checks if the ListAccounts200ResponseCursor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListAccounts200ResponseCursor{}

// ListAccounts200ResponseCursor struct for ListAccounts200ResponseCursor
type ListAccounts200ResponseCursor struct {
	PageSize int32 `json:"page_size"`
	HasMore *bool `json:"has_more,omitempty"`
	Previous *string `json:"previous,omitempty"`
	Next *string `json:"next,omitempty"`
	Data []Account `json:"data"`
}

// NewListAccounts200ResponseCursor instantiates a new ListAccounts200ResponseCursor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListAccounts200ResponseCursor(pageSize int32, data []Account) *ListAccounts200ResponseCursor {
	this := ListAccounts200ResponseCursor{}
	this.PageSize = pageSize
	this.Data = data
	return &this
}

// NewListAccounts200ResponseCursorWithDefaults instantiates a new ListAccounts200ResponseCursor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListAccounts200ResponseCursorWithDefaults() *ListAccounts200ResponseCursor {
	this := ListAccounts200ResponseCursor{}
	return &this
}

// GetPageSize returns the PageSize field value
func (o *ListAccounts200ResponseCursor) GetPageSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value
// and a boolean to check if the value has been set.
func (o *ListAccounts200ResponseCursor) GetPageSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageSize, true
}

// SetPageSize sets field value
func (o *ListAccounts200ResponseCursor) SetPageSize(v int32) {
	o.PageSize = v
}

// GetHasMore returns the HasMore field value if set, zero value otherwise.
func (o *ListAccounts200ResponseCursor) GetHasMore() bool {
	if o == nil || isNil(o.HasMore) {
		var ret bool
		return ret
	}
	return *o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAccounts200ResponseCursor) GetHasMoreOk() (*bool, bool) {
	if o == nil || isNil(o.HasMore) {
		return nil, false
	}
	return o.HasMore, true
}

// HasHasMore returns a boolean if a field has been set.
func (o *ListAccounts200ResponseCursor) HasHasMore() bool {
	if o != nil && !isNil(o.HasMore) {
		return true
	}

	return false
}

// SetHasMore gets a reference to the given bool and assigns it to the HasMore field.
func (o *ListAccounts200ResponseCursor) SetHasMore(v bool) {
	o.HasMore = &v
}

// GetPrevious returns the Previous field value if set, zero value otherwise.
func (o *ListAccounts200ResponseCursor) GetPrevious() string {
	if o == nil || isNil(o.Previous) {
		var ret string
		return ret
	}
	return *o.Previous
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAccounts200ResponseCursor) GetPreviousOk() (*string, bool) {
	if o == nil || isNil(o.Previous) {
		return nil, false
	}
	return o.Previous, true
}

// HasPrevious returns a boolean if a field has been set.
func (o *ListAccounts200ResponseCursor) HasPrevious() bool {
	if o != nil && !isNil(o.Previous) {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given string and assigns it to the Previous field.
func (o *ListAccounts200ResponseCursor) SetPrevious(v string) {
	o.Previous = &v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *ListAccounts200ResponseCursor) GetNext() string {
	if o == nil || isNil(o.Next) {
		var ret string
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAccounts200ResponseCursor) GetNextOk() (*string, bool) {
	if o == nil || isNil(o.Next) {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *ListAccounts200ResponseCursor) HasNext() bool {
	if o != nil && !isNil(o.Next) {
		return true
	}

	return false
}

// SetNext gets a reference to the given string and assigns it to the Next field.
func (o *ListAccounts200ResponseCursor) SetNext(v string) {
	o.Next = &v
}

// GetData returns the Data field value
func (o *ListAccounts200ResponseCursor) GetData() []Account {
	if o == nil {
		var ret []Account
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ListAccounts200ResponseCursor) GetDataOk() ([]Account, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ListAccounts200ResponseCursor) SetData(v []Account) {
	o.Data = v
}

func (o ListAccounts200ResponseCursor) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListAccounts200ResponseCursor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["page_size"] = o.PageSize
	if !isNil(o.HasMore) {
		toSerialize["has_more"] = o.HasMore
	}
	if !isNil(o.Previous) {
		toSerialize["previous"] = o.Previous
	}
	if !isNil(o.Next) {
		toSerialize["next"] = o.Next
	}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableListAccounts200ResponseCursor struct {
	value *ListAccounts200ResponseCursor
	isSet bool
}

func (v NullableListAccounts200ResponseCursor) Get() *ListAccounts200ResponseCursor {
	return v.value
}

func (v *NullableListAccounts200ResponseCursor) Set(val *ListAccounts200ResponseCursor) {
	v.value = val
	v.isSet = true
}

func (v NullableListAccounts200ResponseCursor) IsSet() bool {
	return v.isSet
}

func (v *NullableListAccounts200ResponseCursor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListAccounts200ResponseCursor(val *ListAccounts200ResponseCursor) *NullableListAccounts200ResponseCursor {
	return &NullableListAccounts200ResponseCursor{value: val, isSet: true}
}

func (v NullableListAccounts200ResponseCursor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListAccounts200ResponseCursor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


