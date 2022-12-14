/*
Formance Stack API

Open, modular foundation for unique payments flows  # Introduction This API is documented in **OpenAPI format**.  # Authentication Formance Stack offers one forms of authentication:   - OAuth2 OAuth2 - an open protocol to allow secure authorization in a simple and standard method from web, mobile and desktop applications. <SecurityDefinitions /> 

API version: v1.0.0-rc.1
Contact: support@formance.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package formance

import (
	"encoding/json"
)

// checks if the ResponseCursor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseCursor{}

// ResponseCursor struct for ResponseCursor
type ResponseCursor struct {
	PageSize *float32 `json:"pageSize,omitempty"`
	HasMore *bool `json:"hasMore,omitempty"`
	Total *ResponseCursorTotal `json:"total,omitempty"`
	Next *string `json:"next,omitempty"`
	Previous *string `json:"previous,omitempty"`
	Data []interface{} `json:"data,omitempty"`
}

// NewResponseCursor instantiates a new ResponseCursor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseCursor() *ResponseCursor {
	this := ResponseCursor{}
	return &this
}

// NewResponseCursorWithDefaults instantiates a new ResponseCursor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseCursorWithDefaults() *ResponseCursor {
	this := ResponseCursor{}
	return &this
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *ResponseCursor) GetPageSize() float32 {
	if o == nil || isNil(o.PageSize) {
		var ret float32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCursor) GetPageSizeOk() (*float32, bool) {
	if o == nil || isNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *ResponseCursor) HasPageSize() bool {
	if o != nil && !isNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given float32 and assigns it to the PageSize field.
func (o *ResponseCursor) SetPageSize(v float32) {
	o.PageSize = &v
}

// GetHasMore returns the HasMore field value if set, zero value otherwise.
func (o *ResponseCursor) GetHasMore() bool {
	if o == nil || isNil(o.HasMore) {
		var ret bool
		return ret
	}
	return *o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCursor) GetHasMoreOk() (*bool, bool) {
	if o == nil || isNil(o.HasMore) {
		return nil, false
	}
	return o.HasMore, true
}

// HasHasMore returns a boolean if a field has been set.
func (o *ResponseCursor) HasHasMore() bool {
	if o != nil && !isNil(o.HasMore) {
		return true
	}

	return false
}

// SetHasMore gets a reference to the given bool and assigns it to the HasMore field.
func (o *ResponseCursor) SetHasMore(v bool) {
	o.HasMore = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *ResponseCursor) GetTotal() ResponseCursorTotal {
	if o == nil || isNil(o.Total) {
		var ret ResponseCursorTotal
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCursor) GetTotalOk() (*ResponseCursorTotal, bool) {
	if o == nil || isNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *ResponseCursor) HasTotal() bool {
	if o != nil && !isNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given ResponseCursorTotal and assigns it to the Total field.
func (o *ResponseCursor) SetTotal(v ResponseCursorTotal) {
	o.Total = &v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *ResponseCursor) GetNext() string {
	if o == nil || isNil(o.Next) {
		var ret string
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCursor) GetNextOk() (*string, bool) {
	if o == nil || isNil(o.Next) {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *ResponseCursor) HasNext() bool {
	if o != nil && !isNil(o.Next) {
		return true
	}

	return false
}

// SetNext gets a reference to the given string and assigns it to the Next field.
func (o *ResponseCursor) SetNext(v string) {
	o.Next = &v
}

// GetPrevious returns the Previous field value if set, zero value otherwise.
func (o *ResponseCursor) GetPrevious() string {
	if o == nil || isNil(o.Previous) {
		var ret string
		return ret
	}
	return *o.Previous
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCursor) GetPreviousOk() (*string, bool) {
	if o == nil || isNil(o.Previous) {
		return nil, false
	}
	return o.Previous, true
}

// HasPrevious returns a boolean if a field has been set.
func (o *ResponseCursor) HasPrevious() bool {
	if o != nil && !isNil(o.Previous) {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given string and assigns it to the Previous field.
func (o *ResponseCursor) SetPrevious(v string) {
	o.Previous = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ResponseCursor) GetData() []interface{} {
	if o == nil || isNil(o.Data) {
		var ret []interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCursor) GetDataOk() ([]interface{}, bool) {
	if o == nil || isNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ResponseCursor) HasData() bool {
	if o != nil && !isNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []interface{} and assigns it to the Data field.
func (o *ResponseCursor) SetData(v []interface{}) {
	o.Data = v
}

func (o ResponseCursor) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseCursor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.PageSize) {
		toSerialize["pageSize"] = o.PageSize
	}
	if !isNil(o.HasMore) {
		toSerialize["hasMore"] = o.HasMore
	}
	if !isNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !isNil(o.Next) {
		toSerialize["next"] = o.Next
	}
	if !isNil(o.Previous) {
		toSerialize["previous"] = o.Previous
	}
	if !isNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableResponseCursor struct {
	value *ResponseCursor
	isSet bool
}

func (v NullableResponseCursor) Get() *ResponseCursor {
	return v.value
}

func (v *NullableResponseCursor) Set(val *ResponseCursor) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseCursor) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseCursor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseCursor(val *ResponseCursor) *NullableResponseCursor {
	return &NullableResponseCursor{value: val, isSet: true}
}

func (v NullableResponseCursor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseCursor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


