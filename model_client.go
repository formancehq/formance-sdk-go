/*
Formance Stack API

Open, modular foundation for unique payments flows  # Introduction This API is documented in **OpenAPI format**.  # Authentication Formance Stack offers one forms of authentication:   - OAuth2 OAuth2 - an open protocol to allow secure authorization in a simple and standard method from web, mobile and desktop applications. <SecurityDefinitions /> 

API version: v0.2.5
Contact: support@formance.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package formance

import (
	"encoding/json"
)

// checks if the Client type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Client{}

// Client struct for Client
type Client struct {
	Public *bool `json:"public,omitempty"`
	RedirectUris []string `json:"redirectUris,omitempty"`
	Description *string `json:"description,omitempty"`
	Name string `json:"name"`
	Trusted *bool `json:"trusted,omitempty"`
	PostLogoutRedirectUris []string `json:"postLogoutRedirectUris,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Id string `json:"id"`
	Scopes []string `json:"scopes,omitempty"`
	Secrets []ClientSecret `json:"secrets,omitempty"`
}

// NewClient instantiates a new Client object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClient(name string, id string) *Client {
	this := Client{}
	this.Name = name
	this.Id = id
	return &this
}

// NewClientWithDefaults instantiates a new Client object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientWithDefaults() *Client {
	this := Client{}
	return &this
}

// GetPublic returns the Public field value if set, zero value otherwise.
func (o *Client) GetPublic() bool {
	if o == nil || isNil(o.Public) {
		var ret bool
		return ret
	}
	return *o.Public
}

// GetPublicOk returns a tuple with the Public field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetPublicOk() (*bool, bool) {
	if o == nil || isNil(o.Public) {
		return nil, false
	}
	return o.Public, true
}

// HasPublic returns a boolean if a field has been set.
func (o *Client) HasPublic() bool {
	if o != nil && !isNil(o.Public) {
		return true
	}

	return false
}

// SetPublic gets a reference to the given bool and assigns it to the Public field.
func (o *Client) SetPublic(v bool) {
	o.Public = &v
}

// GetRedirectUris returns the RedirectUris field value if set, zero value otherwise.
func (o *Client) GetRedirectUris() []string {
	if o == nil || isNil(o.RedirectUris) {
		var ret []string
		return ret
	}
	return o.RedirectUris
}

// GetRedirectUrisOk returns a tuple with the RedirectUris field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetRedirectUrisOk() ([]string, bool) {
	if o == nil || isNil(o.RedirectUris) {
		return nil, false
	}
	return o.RedirectUris, true
}

// HasRedirectUris returns a boolean if a field has been set.
func (o *Client) HasRedirectUris() bool {
	if o != nil && !isNil(o.RedirectUris) {
		return true
	}

	return false
}

// SetRedirectUris gets a reference to the given []string and assigns it to the RedirectUris field.
func (o *Client) SetRedirectUris(v []string) {
	o.RedirectUris = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Client) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Client) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Client) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value
func (o *Client) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Client) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Client) SetName(v string) {
	o.Name = v
}

// GetTrusted returns the Trusted field value if set, zero value otherwise.
func (o *Client) GetTrusted() bool {
	if o == nil || isNil(o.Trusted) {
		var ret bool
		return ret
	}
	return *o.Trusted
}

// GetTrustedOk returns a tuple with the Trusted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetTrustedOk() (*bool, bool) {
	if o == nil || isNil(o.Trusted) {
		return nil, false
	}
	return o.Trusted, true
}

// HasTrusted returns a boolean if a field has been set.
func (o *Client) HasTrusted() bool {
	if o != nil && !isNil(o.Trusted) {
		return true
	}

	return false
}

// SetTrusted gets a reference to the given bool and assigns it to the Trusted field.
func (o *Client) SetTrusted(v bool) {
	o.Trusted = &v
}

// GetPostLogoutRedirectUris returns the PostLogoutRedirectUris field value if set, zero value otherwise.
func (o *Client) GetPostLogoutRedirectUris() []string {
	if o == nil || isNil(o.PostLogoutRedirectUris) {
		var ret []string
		return ret
	}
	return o.PostLogoutRedirectUris
}

// GetPostLogoutRedirectUrisOk returns a tuple with the PostLogoutRedirectUris field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetPostLogoutRedirectUrisOk() ([]string, bool) {
	if o == nil || isNil(o.PostLogoutRedirectUris) {
		return nil, false
	}
	return o.PostLogoutRedirectUris, true
}

// HasPostLogoutRedirectUris returns a boolean if a field has been set.
func (o *Client) HasPostLogoutRedirectUris() bool {
	if o != nil && !isNil(o.PostLogoutRedirectUris) {
		return true
	}

	return false
}

// SetPostLogoutRedirectUris gets a reference to the given []string and assigns it to the PostLogoutRedirectUris field.
func (o *Client) SetPostLogoutRedirectUris(v []string) {
	o.PostLogoutRedirectUris = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *Client) GetMetadata() map[string]interface{} {
	if o == nil || isNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Client) HasMetadata() bool {
	if o != nil && !isNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *Client) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetId returns the Id field value
func (o *Client) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Client) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Client) SetId(v string) {
	o.Id = v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *Client) GetScopes() []string {
	if o == nil || isNil(o.Scopes) {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetScopesOk() ([]string, bool) {
	if o == nil || isNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *Client) HasScopes() bool {
	if o != nil && !isNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *Client) SetScopes(v []string) {
	o.Scopes = v
}

// GetSecrets returns the Secrets field value if set, zero value otherwise.
func (o *Client) GetSecrets() []ClientSecret {
	if o == nil || isNil(o.Secrets) {
		var ret []ClientSecret
		return ret
	}
	return o.Secrets
}

// GetSecretsOk returns a tuple with the Secrets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetSecretsOk() ([]ClientSecret, bool) {
	if o == nil || isNil(o.Secrets) {
		return nil, false
	}
	return o.Secrets, true
}

// HasSecrets returns a boolean if a field has been set.
func (o *Client) HasSecrets() bool {
	if o != nil && !isNil(o.Secrets) {
		return true
	}

	return false
}

// SetSecrets gets a reference to the given []ClientSecret and assigns it to the Secrets field.
func (o *Client) SetSecrets(v []ClientSecret) {
	o.Secrets = v
}

func (o Client) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Client) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Public) {
		toSerialize["public"] = o.Public
	}
	if !isNil(o.RedirectUris) {
		toSerialize["redirectUris"] = o.RedirectUris
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["name"] = o.Name
	if !isNil(o.Trusted) {
		toSerialize["trusted"] = o.Trusted
	}
	if !isNil(o.PostLogoutRedirectUris) {
		toSerialize["postLogoutRedirectUris"] = o.PostLogoutRedirectUris
	}
	if !isNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	toSerialize["id"] = o.Id
	if !isNil(o.Scopes) {
		toSerialize["scopes"] = o.Scopes
	}
	if !isNil(o.Secrets) {
		toSerialize["secrets"] = o.Secrets
	}
	return toSerialize, nil
}

type NullableClient struct {
	value *Client
	isSet bool
}

func (v NullableClient) Get() *Client {
	return v.value
}

func (v *NullableClient) Set(val *Client) {
	v.value = val
	v.isSet = true
}

func (v NullableClient) IsSet() bool {
	return v.isSet
}

func (v *NullableClient) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClient(val *Client) *NullableClient {
	return &NullableClient{value: val, isSet: true}
}

func (v NullableClient) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClient) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


