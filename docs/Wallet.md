# Wallet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique ID of the wallet. | 
**Metadata** | **map[string]interface{}** | Metadata associated with the wallet. | 
**Balances** | **map[string]int32** | The balances of the wallet. | 
**Name** | **string** |  | 

## Methods

### NewWallet

`func NewWallet(id string, metadata map[string]interface{}, balances map[string]int32, name string, ) *Wallet`

NewWallet instantiates a new Wallet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWalletWithDefaults

`func NewWalletWithDefaults() *Wallet`

NewWalletWithDefaults instantiates a new Wallet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Wallet) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Wallet) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Wallet) SetId(v string)`

SetId sets Id field to given value.


### GetMetadata

`func (o *Wallet) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Wallet) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Wallet) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### GetBalances

`func (o *Wallet) GetBalances() map[string]int32`

GetBalances returns the Balances field if non-nil, zero value otherwise.

### GetBalancesOk

`func (o *Wallet) GetBalancesOk() (*map[string]int32, bool)`

GetBalancesOk returns a tuple with the Balances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalances

`func (o *Wallet) SetBalances(v map[string]int32)`

SetBalances sets Balances field to given value.


### GetName

`func (o *Wallet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Wallet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Wallet) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


