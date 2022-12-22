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

// checks if the Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Response{}

// Response struct for Response
type Response struct {
	// The payload
	Data map[string]interface{} `json:"data,omitempty"`
	Cursor *ResponseCursor `json:"cursor,omitempty"`
	// The kind of the object, either \"TRANSACTION\" or \"META\"
	Kind *string `json:"kind,omitempty"`
	// The ledger
	Ledger *string `json:"ledger,omitempty"`
}

// NewResponse instantiates a new Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponse() *Response {
	this := Response{}
	return &this
}

// NewResponseWithDefaults instantiates a new Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseWithDefaults() *Response {
	this := Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *Response) GetData() map[string]interface{} {
	if o == nil || isNil(o.Data) {
		var ret map[string]interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Response) GetDataOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Data) {
		return map[string]interface{}{}, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *Response) HasData() bool {
	if o != nil && !isNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]interface{} and assigns it to the Data field.
func (o *Response) SetData(v map[string]interface{}) {
	o.Data = v
}

// GetCursor returns the Cursor field value if set, zero value otherwise.
func (o *Response) GetCursor() ResponseCursor {
	if o == nil || isNil(o.Cursor) {
		var ret ResponseCursor
		return ret
	}
	return *o.Cursor
}

// GetCursorOk returns a tuple with the Cursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Response) GetCursorOk() (*ResponseCursor, bool) {
	if o == nil || isNil(o.Cursor) {
		return nil, false
	}
	return o.Cursor, true
}

// HasCursor returns a boolean if a field has been set.
func (o *Response) HasCursor() bool {
	if o != nil && !isNil(o.Cursor) {
		return true
	}

	return false
}

// SetCursor gets a reference to the given ResponseCursor and assigns it to the Cursor field.
func (o *Response) SetCursor(v ResponseCursor) {
	o.Cursor = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *Response) GetKind() string {
	if o == nil || isNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Response) GetKindOk() (*string, bool) {
	if o == nil || isNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *Response) HasKind() bool {
	if o != nil && !isNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *Response) SetKind(v string) {
	o.Kind = &v
}

// GetLedger returns the Ledger field value if set, zero value otherwise.
func (o *Response) GetLedger() string {
	if o == nil || isNil(o.Ledger) {
		var ret string
		return ret
	}
	return *o.Ledger
}

// GetLedgerOk returns a tuple with the Ledger field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Response) GetLedgerOk() (*string, bool) {
	if o == nil || isNil(o.Ledger) {
		return nil, false
	}
	return o.Ledger, true
}

// HasLedger returns a boolean if a field has been set.
func (o *Response) HasLedger() bool {
	if o != nil && !isNil(o.Ledger) {
		return true
	}

	return false
}

// SetLedger gets a reference to the given string and assigns it to the Ledger field.
func (o *Response) SetLedger(v string) {
	o.Ledger = &v
}

func (o Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !isNil(o.Cursor) {
		toSerialize["cursor"] = o.Cursor
	}
	if !isNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	if !isNil(o.Ledger) {
		toSerialize["ledger"] = o.Ledger
	}
	return toSerialize, nil
}

type NullableResponse struct {
	value *Response
	isSet bool
}

func (v NullableResponse) Get() *Response {
	return v.value
}

func (v *NullableResponse) Set(val *Response) {
	v.value = val
	v.isSet = true
}

func (v NullableResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponse(val *Response) *NullableResponse {
	return &NullableResponse{value: val, isSet: true}
}

func (v NullableResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


