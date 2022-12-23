# DebitWalletRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | [**Monetary**](Monetary.md) |  | 
**Pending** | Pointer to **bool** | Set to true to create a pending hold. If false, the wallet will be debited immediately. | [optional] 

## Methods

### NewDebitWalletRequest

`func NewDebitWalletRequest(amount Monetary, ) *DebitWalletRequest`

NewDebitWalletRequest instantiates a new DebitWalletRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDebitWalletRequestWithDefaults

`func NewDebitWalletRequestWithDefaults() *DebitWalletRequest`

NewDebitWalletRequestWithDefaults instantiates a new DebitWalletRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *DebitWalletRequest) GetAmount() Monetary`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *DebitWalletRequest) GetAmountOk() (*Monetary, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *DebitWalletRequest) SetAmount(v Monetary)`

SetAmount sets Amount field to given value.


### GetPending

`func (o *DebitWalletRequest) GetPending() bool`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *DebitWalletRequest) GetPendingOk() (*bool, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *DebitWalletRequest) SetPending(v bool)`

SetPending sets Pending field to given value.

### HasPending

`func (o *DebitWalletRequest) HasPending() bool`

HasPending returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


