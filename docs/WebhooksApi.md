# \WebhooksApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateOneConfig**](WebhooksApi.md#ActivateOneConfig) | **Put** /api/webhooks/configs/{id}/activate | Activate one config
[**ChangeOneConfigSecret**](WebhooksApi.md#ChangeOneConfigSecret) | **Put** /api/webhooks/configs/{id}/secret/change | Change the signing secret of a config
[**DeactivateOneConfig**](WebhooksApi.md#DeactivateOneConfig) | **Put** /api/webhooks/configs/{id}/deactivate | Deactivate one config
[**DeleteOneConfig**](WebhooksApi.md#DeleteOneConfig) | **Delete** /api/webhooks/configs/{id} | Delete one config
[**GetManyConfigs**](WebhooksApi.md#GetManyConfigs) | **Get** /api/webhooks/configs | Get many configs
[**InsertOneConfig**](WebhooksApi.md#InsertOneConfig) | **Post** /api/webhooks/configs | Insert a new config 
[**TestOneConfig**](WebhooksApi.md#TestOneConfig) | **Get** /api/webhooks/configs/{id}/test | Test one config



## ActivateOneConfig

> ConfigActivatedResponse ActivateOneConfig(ctx, id).Execute()

Activate one config

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
    id := "4997257d-dfb6-445b-929c-cbe2ab182818" // string | Config ID

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksApi.ActivateOneConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.ActivateOneConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActivateOneConfig`: ConfigActivatedResponse
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.ActivateOneConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Config ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateOneConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConfigActivatedResponse**](ConfigActivatedResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChangeOneConfigSecret

> ConfigActivatedResponse ChangeOneConfigSecret(ctx, id).ChangeOneConfigSecretRequest(changeOneConfigSecretRequest).Execute()

Change the signing secret of a config



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
    id := "4997257d-dfb6-445b-929c-cbe2ab182818" // string | Config ID
    changeOneConfigSecretRequest := *client.NewChangeOneConfigSecretRequest("V0bivxRWveaoz08afqjU6Ko/jwO0Cb+3") // ChangeOneConfigSecretRequest |  (optional)

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksApi.ChangeOneConfigSecret(context.Background(), id).ChangeOneConfigSecretRequest(changeOneConfigSecretRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.ChangeOneConfigSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChangeOneConfigSecret`: ConfigActivatedResponse
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.ChangeOneConfigSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Config ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeOneConfigSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **changeOneConfigSecretRequest** | [**ChangeOneConfigSecretRequest**](ChangeOneConfigSecretRequest.md) |  | 

### Return type

[**ConfigActivatedResponse**](ConfigActivatedResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeactivateOneConfig

> ConfigDeactivatedResponse DeactivateOneConfig(ctx, id).Execute()

Deactivate one config

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
    id := "4997257d-dfb6-445b-929c-cbe2ab182818" // string | Config ID

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksApi.DeactivateOneConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.DeactivateOneConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeactivateOneConfig`: ConfigDeactivatedResponse
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.DeactivateOneConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Config ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeactivateOneConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConfigDeactivatedResponse**](ConfigDeactivatedResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOneConfig

> DeleteOneConfig(ctx, id).Execute()

Delete one config

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
    id := "4997257d-dfb6-445b-929c-cbe2ab182818" // string | Config ID

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksApi.DeleteOneConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.DeleteOneConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Config ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOneConfigRequest struct via the builder pattern


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


## GetManyConfigs

> GetManyConfigs200Response GetManyConfigs(ctx).Id(id).Endpoint(endpoint).Execute()

Get many configs



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
    id := "4997257d-dfb6-445b-929c-cbe2ab182818" // string | Optional filter by Config ID (optional)
    endpoint := "https://example.com" // string | Optional filter by endpoint URL (optional)

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksApi.GetManyConfigs(context.Background()).Id(id).Endpoint(endpoint).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.GetManyConfigs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetManyConfigs`: GetManyConfigs200Response
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.GetManyConfigs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetManyConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Optional filter by Config ID | 
 **endpoint** | **string** | Optional filter by endpoint URL | 

### Return type

[**GetManyConfigs200Response**](GetManyConfigs200Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InsertOneConfig

> ConfigActivatedResponse InsertOneConfig(ctx).ConfigUser(configUser).Execute()

Insert a new config 



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
    configUser := *client.NewConfigUser() // ConfigUser | 

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksApi.InsertOneConfig(context.Background()).ConfigUser(configUser).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.InsertOneConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InsertOneConfig`: ConfigActivatedResponse
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.InsertOneConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInsertOneConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **configUser** | [**ConfigUser**](ConfigUser.md) |  | 

### Return type

[**ConfigActivatedResponse**](ConfigActivatedResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestOneConfig

> AttemptResponse TestOneConfig(ctx, id).Execute()

Test one config



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
    id := "4997257d-dfb6-445b-929c-cbe2ab182818" // string | Config ID

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksApi.TestOneConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.TestOneConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestOneConfig`: AttemptResponse
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.TestOneConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Config ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestOneConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AttemptResponse**](AttemptResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

