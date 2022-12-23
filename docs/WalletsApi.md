# \WalletsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConfirmHold**](WalletsApi.md#ConfirmHold) | **Post** /api/wallets/wallets/{id}/holds/{hold_id}/confirm | Confirm a hold
[**CreateWallet**](WalletsApi.md#CreateWallet) | **Post** /api/wallets/wallets | Create a new wallet
[**CreditWallet**](WalletsApi.md#CreditWallet) | **Post** /api/wallets/wallets/{id}/credit | Credit a wallet
[**DebitWallet**](WalletsApi.md#DebitWallet) | **Post** /api/wallets/wallets/{id}/debit | Debit a wallet
[**GetHolds**](WalletsApi.md#GetHolds) | **Get** /api/wallets/wallets/{id}/holds | Get all holds for a wallet
[**GetWallet**](WalletsApi.md#GetWallet) | **Get** /api/wallets/wallets/{id} | Get a wallet
[**GetWallets**](WalletsApi.md#GetWallets) | **Get** /api/wallets/wallets | Get all wallets
[**UpdateWallet**](WalletsApi.md#UpdateWallet) | **Patch** /api/wallets/wallets/{id} | Update a wallet
[**VoidHold**](WalletsApi.md#VoidHold) | **Post** /api/wallets/wallets/{id}/holds/{hold_id}/void | Cancel a hold



## ConfirmHold

> ConfirmHold(ctx, id, holdId).Execute()

Confirm a hold

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    client "./openapi"
)

func main() {
    id := "id_example" // string | 
    holdId := "holdId_example" // string | 

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.WalletsApi.ConfirmHold(context.Background(), id, holdId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletsApi.ConfirmHold``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**holdId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConfirmHoldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWallet

> CreateWalletResponse CreateWallet(ctx).CreateWalletRequest(createWalletRequest).Execute()

Create a new wallet

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    client "./openapi"
)

func main() {
    createWalletRequest := *client.NewCreateWalletRequest("Name_example") // CreateWalletRequest |  (optional)

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.WalletsApi.CreateWallet(context.Background()).CreateWalletRequest(createWalletRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletsApi.CreateWallet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWallet`: CreateWalletResponse
    fmt.Fprintf(os.Stdout, "Response from `WalletsApi.CreateWallet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWalletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createWalletRequest** | [**CreateWalletRequest**](CreateWalletRequest.md) |  | 

### Return type

[**CreateWalletResponse**](CreateWalletResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreditWallet

> CreditWallet(ctx, id).CreditWalletRequest(creditWalletRequest).Execute()

Credit a wallet

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    client "./openapi"
)

func main() {
    id := "id_example" // string | 
    creditWalletRequest := *client.NewCreditWalletRequest(*client.NewMonetary("Asset_example", float32(123))) // CreditWalletRequest |  (optional)

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.WalletsApi.CreditWallet(context.Background(), id).CreditWalletRequest(creditWalletRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletsApi.CreditWallet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreditWalletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **creditWalletRequest** | [**CreditWalletRequest**](CreditWalletRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DebitWallet

> DebitWalletResponse DebitWallet(ctx, id).DebitWalletRequest(debitWalletRequest).Execute()

Debit a wallet

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    client "./openapi"
)

func main() {
    id := "id_example" // string | 
    debitWalletRequest := *client.NewDebitWalletRequest(*client.NewMonetary("Asset_example", float32(123))) // DebitWalletRequest |  (optional)

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.WalletsApi.DebitWallet(context.Background(), id).DebitWalletRequest(debitWalletRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletsApi.DebitWallet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DebitWallet`: DebitWalletResponse
    fmt.Fprintf(os.Stdout, "Response from `WalletsApi.DebitWallet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDebitWalletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **debitWalletRequest** | [**DebitWalletRequest**](DebitWalletRequest.md) |  | 

### Return type

[**DebitWalletResponse**](DebitWalletResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHolds

> GetHoldsResponse GetHolds(ctx, id).Execute()

Get all holds for a wallet

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    client "./openapi"
)

func main() {
    id := "id_example" // string | 

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.WalletsApi.GetHolds(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletsApi.GetHolds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHolds`: GetHoldsResponse
    fmt.Fprintf(os.Stdout, "Response from `WalletsApi.GetHolds`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHoldsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetHoldsResponse**](GetHoldsResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWallet

> GetWalletResponse GetWallet(ctx, id).Execute()

Get a wallet

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    client "./openapi"
)

func main() {
    id := "id_example" // string | 

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.WalletsApi.GetWallet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletsApi.GetWallet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWallet`: GetWalletResponse
    fmt.Fprintf(os.Stdout, "Response from `WalletsApi.GetWallet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWalletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetWalletResponse**](GetWalletResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWallets

> GetWalletsResponse GetWallets(ctx).PageSize(pageSize).After(after).PaginationToken(paginationToken).Execute()

Get all wallets

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    client "./openapi"
)

func main() {
    pageSize := int32(100) // int32 | The maximum number of results to return per page (optional) (default to 15)
    after := "users:003" // string | Pagination cursor, will return accounts after given address, in descending order. (optional)
    paginationToken := "aHR0cHM6Ly9nLnBhZ2UvTmVrby1SYW1lbj9zaGFyZQ==" // string | Parameter used in pagination requests. Maximum page size is set to 15. Set to the value of next for the next page of results. Set to the value of previous for the previous page of results. No other parameters can be set when the pagination token is set.  (optional)

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.WalletsApi.GetWallets(context.Background()).PageSize(pageSize).After(after).PaginationToken(paginationToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletsApi.GetWallets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWallets`: GetWalletsResponse
    fmt.Fprintf(os.Stdout, "Response from `WalletsApi.GetWallets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWalletsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | The maximum number of results to return per page | [default to 15]
 **after** | **string** | Pagination cursor, will return accounts after given address, in descending order. | 
 **paginationToken** | **string** | Parameter used in pagination requests. Maximum page size is set to 15. Set to the value of next for the next page of results. Set to the value of previous for the previous page of results. No other parameters can be set when the pagination token is set.  | 

### Return type

[**GetWalletsResponse**](GetWalletsResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWallet

> UpdateWallet(ctx, id).UpdateWalletRequest(updateWalletRequest).Execute()

Update a wallet

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    client "./openapi"
)

func main() {
    id := "id_example" // string | 
    updateWalletRequest := *client.NewUpdateWalletRequest() // UpdateWalletRequest |  (optional)

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.WalletsApi.UpdateWallet(context.Background(), id).UpdateWalletRequest(updateWalletRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletsApi.UpdateWallet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWalletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateWalletRequest** | [**UpdateWalletRequest**](UpdateWalletRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VoidHold

> VoidHold(ctx, id, holdId).Execute()

Cancel a hold

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    client "./openapi"
)

func main() {
    id := "id_example" // string | 
    holdId := "holdId_example" // string | 

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.WalletsApi.VoidHold(context.Background(), id, holdId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletsApi.VoidHold``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**holdId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVoidHoldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

