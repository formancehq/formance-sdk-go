/*
Formance Stack API

Open, modular foundation for unique payments flows  # Introduction This API is documented in **OpenAPI format**.  # Authentication Formance Stack offers one forms of authentication:   - OAuth2 OAuth2 - an open protocol to allow secure authorization in a simple and standard method from web, mobile and desktop applications. <SecurityDefinitions /> 

API version: v1.0.0-rc.5
Contact: support@formance.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package formance

import (
	"encoding/json"
)

// checks if the Query type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Query{}

// Query struct for Query
type Query struct {
	Ledgers []string `json:"ledgers,omitempty"`
	NextToken *string `json:"nextToken,omitempty"`
	Size *int32 `json:"size,omitempty"`
	Terms []string `json:"terms,omitempty"`
	Target *string `json:"target,omitempty"`
}

// NewQuery instantiates a new Query object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuery() *Query {
	this := Query{}
	return &this
}

// NewQueryWithDefaults instantiates a new Query object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryWithDefaults() *Query {
	this := Query{}
	return &this
}

// GetLedgers returns the Ledgers field value if set, zero value otherwise.
func (o *Query) GetLedgers() []string {
	if o == nil || isNil(o.Ledgers) {
		var ret []string
		return ret
	}
	return o.Ledgers
}

// GetLedgersOk returns a tuple with the Ledgers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Query) GetLedgersOk() ([]string, bool) {
	if o == nil || isNil(o.Ledgers) {
		return nil, false
	}
	return o.Ledgers, true
}

// HasLedgers returns a boolean if a field has been set.
func (o *Query) HasLedgers() bool {
	if o != nil && !isNil(o.Ledgers) {
		return true
	}

	return false
}

// SetLedgers gets a reference to the given []string and assigns it to the Ledgers field.
func (o *Query) SetLedgers(v []string) {
	o.Ledgers = v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *Query) GetNextToken() string {
	if o == nil || isNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Query) GetNextTokenOk() (*string, bool) {
	if o == nil || isNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *Query) HasNextToken() bool {
	if o != nil && !isNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *Query) SetNextToken(v string) {
	o.NextToken = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *Query) GetSize() int32 {
	if o == nil || isNil(o.Size) {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Query) GetSizeOk() (*int32, bool) {
	if o == nil || isNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *Query) HasSize() bool {
	if o != nil && !isNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *Query) SetSize(v int32) {
	o.Size = &v
}

// GetTerms returns the Terms field value if set, zero value otherwise.
func (o *Query) GetTerms() []string {
	if o == nil || isNil(o.Terms) {
		var ret []string
		return ret
	}
	return o.Terms
}

// GetTermsOk returns a tuple with the Terms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Query) GetTermsOk() ([]string, bool) {
	if o == nil || isNil(o.Terms) {
		return nil, false
	}
	return o.Terms, true
}

// HasTerms returns a boolean if a field has been set.
func (o *Query) HasTerms() bool {
	if o != nil && !isNil(o.Terms) {
		return true
	}

	return false
}

// SetTerms gets a reference to the given []string and assigns it to the Terms field.
func (o *Query) SetTerms(v []string) {
	o.Terms = v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *Query) GetTarget() string {
	if o == nil || isNil(o.Target) {
		var ret string
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Query) GetTargetOk() (*string, bool) {
	if o == nil || isNil(o.Target) {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *Query) HasTarget() bool {
	if o != nil && !isNil(o.Target) {
		return true
	}

	return false
}

// SetTarget gets a reference to the given string and assigns it to the Target field.
func (o *Query) SetTarget(v string) {
	o.Target = &v
}

func (o Query) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Query) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Ledgers) {
		toSerialize["ledgers"] = o.Ledgers
	}
	if !isNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	if !isNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !isNil(o.Terms) {
		toSerialize["terms"] = o.Terms
	}
	if !isNil(o.Target) {
		toSerialize["target"] = o.Target
	}
	return toSerialize, nil
}

type NullableQuery struct {
	value *Query
	isSet bool
}

func (v NullableQuery) Get() *Query {
	return v.value
}

func (v *NullableQuery) Set(val *Query) {
	v.value = val
	v.isSet = true
}

func (v NullableQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuery(val *Query) *NullableQuery {
	return &NullableQuery{value: val, isSet: true}
}

func (v NullableQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


