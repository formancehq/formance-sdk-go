# TaskDescriptorCurrencyCloud

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | Pointer to **string** | The connector code | [optional] 
**CreatedAt** | Pointer to **time.Time** | The date when the task was created | [optional] 
**Status** | Pointer to **string** | The task status | [optional] 
**Error** | Pointer to **string** | The error message if the task failed | [optional] 
**State** | Pointer to **map[string]interface{}** | The task state | [optional] 
**Descriptor** | Pointer to [**TaskDescriptorCurrencyCloudAllOfDescriptor**](TaskDescriptorCurrencyCloudAllOfDescriptor.md) |  | [optional] 

## Methods

### NewTaskDescriptorCurrencyCloud

`func NewTaskDescriptorCurrencyCloud() *TaskDescriptorCurrencyCloud`

NewTaskDescriptorCurrencyCloud instantiates a new TaskDescriptorCurrencyCloud object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskDescriptorCurrencyCloudWithDefaults

`func NewTaskDescriptorCurrencyCloudWithDefaults() *TaskDescriptorCurrencyCloud`

NewTaskDescriptorCurrencyCloudWithDefaults instantiates a new TaskDescriptorCurrencyCloud object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *TaskDescriptorCurrencyCloud) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *TaskDescriptorCurrencyCloud) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *TaskDescriptorCurrencyCloud) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *TaskDescriptorCurrencyCloud) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TaskDescriptorCurrencyCloud) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TaskDescriptorCurrencyCloud) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TaskDescriptorCurrencyCloud) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TaskDescriptorCurrencyCloud) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetStatus

`func (o *TaskDescriptorCurrencyCloud) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TaskDescriptorCurrencyCloud) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TaskDescriptorCurrencyCloud) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TaskDescriptorCurrencyCloud) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetError

`func (o *TaskDescriptorCurrencyCloud) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *TaskDescriptorCurrencyCloud) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *TaskDescriptorCurrencyCloud) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *TaskDescriptorCurrencyCloud) HasError() bool`

HasError returns a boolean if a field has been set.

### GetState

`func (o *TaskDescriptorCurrencyCloud) GetState() map[string]interface{}`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *TaskDescriptorCurrencyCloud) GetStateOk() (*map[string]interface{}, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *TaskDescriptorCurrencyCloud) SetState(v map[string]interface{})`

SetState sets State field to given value.

### HasState

`func (o *TaskDescriptorCurrencyCloud) HasState() bool`

HasState returns a boolean if a field has been set.

### GetDescriptor

`func (o *TaskDescriptorCurrencyCloud) GetDescriptor() TaskDescriptorCurrencyCloudAllOfDescriptor`

GetDescriptor returns the Descriptor field if non-nil, zero value otherwise.

### GetDescriptorOk

`func (o *TaskDescriptorCurrencyCloud) GetDescriptorOk() (*TaskDescriptorCurrencyCloudAllOfDescriptor, bool)`

GetDescriptorOk returns a tuple with the Descriptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptor

`func (o *TaskDescriptorCurrencyCloud) SetDescriptor(v TaskDescriptorCurrencyCloudAllOfDescriptor)`

SetDescriptor sets Descriptor field to given value.

### HasDescriptor

`func (o *TaskDescriptorCurrencyCloud) HasDescriptor() bool`

HasDescriptor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


