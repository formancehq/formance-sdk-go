# CreditWalletRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | [**Monetary**](Monetary.md) |  | 

## Methods

### NewCreditWalletRequest

`func NewCreditWalletRequest(amount Monetary, ) *CreditWalletRequest`

NewCreditWalletRequest instantiates a new CreditWalletRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditWalletRequestWithDefaults

`func NewCreditWalletRequestWithDefaults() *CreditWalletRequest`

NewCreditWalletRequestWithDefaults instantiates a new CreditWalletRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *CreditWalletRequest) GetAmount() Monetary`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CreditWalletRequest) GetAmountOk() (*Monetary, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CreditWalletRequest) SetAmount(v Monetary)`

SetAmount sets Amount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


