# Hold

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique ID of the hold. | 
**WalletId** | **string** | The ID of the wallet the hold is associated with. | 
**Metadata** | **map[string]interface{}** | Metadata associated with the hold. | 

## Methods

### NewHold

`func NewHold(id string, walletId string, metadata map[string]interface{}, ) *Hold`

NewHold instantiates a new Hold object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHoldWithDefaults

`func NewHoldWithDefaults() *Hold`

NewHoldWithDefaults instantiates a new Hold object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Hold) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Hold) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Hold) SetId(v string)`

SetId sets Id field to given value.


### GetWalletId

`func (o *Hold) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *Hold) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *Hold) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetMetadata

`func (o *Hold) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Hold) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Hold) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


