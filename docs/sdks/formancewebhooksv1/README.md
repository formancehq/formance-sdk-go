# Webhooks.V1

## Overview

### Available Operations

* [ActivateConfig](#activateconfig) - Activate one config
* [ChangeConfigSecret](#changeconfigsecret) - Change the signing secret of a config
* [DeactivateConfig](#deactivateconfig) - Deactivate one config
* [DeleteConfig](#deleteconfig) - Delete one config
* [GetDeliveries](#getdeliveries) - List webhook deliveries
* [GetDelivery](#getdelivery) - Get a webhook delivery
* [GetDeliveryAttempts](#getdeliveryattempts) - List attempts for a webhook delivery
* [GetManyConfigs](#getmanyconfigs) - Get many configs
* [InsertConfig](#insertconfig) - Insert a new config
* [ReplayDeliveries](#replaydeliveries) - Replay a page of failed or pending deliveries
* [ReplayDelivery](#replaydelivery) - Replay one failed or pending delivery
* [TestConfig](#testconfig) - Test one config
* [UpdateConfig](#updateconfig) - Update one config

## ActivateConfig

Activate a webhooks config by ID, to start receiving webhooks to its endpoint.

### Example Usage

<!-- UsageSnippet language="go" operationID="activateConfig" method="put" path="/api/webhooks/configs/{id}/activate" -->
```go
package main

import(
	"context"
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/shared"
	"github.com/formancehq/formance-sdk-go/v5"
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := v5.New(
        v5.WithSecurity(shared.Security{
            ClientID: v5.Pointer("<YOUR_CLIENT_ID_HERE>"),
            ClientSecret: v5.Pointer("<YOUR_CLIENT_SECRET_HERE>"),
        }),
    )

    res, err := s.Webhooks.V1.ActivateConfig(ctx, operations.ActivateConfigRequest{
        ID: "4997257d-dfb6-445b-929c-cbe2ab182818",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ConfigResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.ActivateConfigRequest](../../pkg/models/operations/activateconfigrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.ActivateConfigResponse](../../pkg/models/operations/activateconfigresponse.md), error**

### Errors

| Error Type             | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| webhooks.ErrorResponse | default                | application/json       |
| sdkerrors.SDKError     | 4XX, 5XX               | \*/\*                  |

## ChangeConfigSecret

Change the signing secret of the endpoint of a webhooks config.

If not passed or empty, a secret is automatically generated.
The format is a random string of bytes of size 24, base64 encoded. (larger size after encoding)


### Example Usage

<!-- UsageSnippet language="go" operationID="changeConfigSecret" method="put" path="/api/webhooks/configs/{id}/secret/change" -->
```go
package main

import(
	"context"
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/shared"
	"github.com/formancehq/formance-sdk-go/v5"
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/webhooks"
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := v5.New(
        v5.WithSecurity(shared.Security{
            ClientID: v5.Pointer("<YOUR_CLIENT_ID_HERE>"),
            ClientSecret: v5.Pointer("<YOUR_CLIENT_SECRET_HERE>"),
        }),
    )

    res, err := s.Webhooks.V1.ChangeConfigSecret(ctx, operations.ChangeConfigSecretRequest{
        ConfigChangeSecret: &webhooks.ConfigChangeSecret{
            Secret: "V0bivxRWveaoz08afqjU6Ko/jwO0Cb+3",
        },
        ID: "4997257d-dfb6-445b-929c-cbe2ab182818",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ConfigResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.ChangeConfigSecretRequest](../../pkg/models/operations/changeconfigsecretrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.ChangeConfigSecretResponse](../../pkg/models/operations/changeconfigsecretresponse.md), error**

### Errors

| Error Type             | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| webhooks.ErrorResponse | default                | application/json       |
| sdkerrors.SDKError     | 4XX, 5XX               | \*/\*                  |

## DeactivateConfig

Deactivate a webhooks config by ID, to stop receiving webhooks to its endpoint.

### Example Usage

<!-- UsageSnippet language="go" operationID="deactivateConfig" method="put" path="/api/webhooks/configs/{id}/deactivate" -->
```go
package main

import(
	"context"
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/shared"
	"github.com/formancehq/formance-sdk-go/v5"
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := v5.New(
        v5.WithSecurity(shared.Security{
            ClientID: v5.Pointer("<YOUR_CLIENT_ID_HERE>"),
            ClientSecret: v5.Pointer("<YOUR_CLIENT_SECRET_HERE>"),
        }),
    )

    res, err := s.Webhooks.V1.DeactivateConfig(ctx, operations.DeactivateConfigRequest{
        ID: "4997257d-dfb6-445b-929c-cbe2ab182818",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ConfigResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.DeactivateConfigRequest](../../pkg/models/operations/deactivateconfigrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.DeactivateConfigResponse](../../pkg/models/operations/deactivateconfigresponse.md), error**

### Errors

| Error Type             | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| webhooks.ErrorResponse | default                | application/json       |
| sdkerrors.SDKError     | 4XX, 5XX               | \*/\*                  |

## DeleteConfig

Delete a webhooks config by ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteConfig" method="delete" path="/api/webhooks/configs/{id}" -->
```go
package main

import(
	"context"
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/shared"
	"github.com/formancehq/formance-sdk-go/v5"
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := v5.New(
        v5.WithSecurity(shared.Security{
            ClientID: v5.Pointer("<YOUR_CLIENT_ID_HERE>"),
            ClientSecret: v5.Pointer("<YOUR_CLIENT_SECRET_HERE>"),
        }),
    )

    res, err := s.Webhooks.V1.DeleteConfig(ctx, operations.DeleteConfigRequest{
        ID: "4997257d-dfb6-445b-929c-cbe2ab182818",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.DeleteConfigRequest](../../pkg/models/operations/deleteconfigrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                         | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.DeleteConfigResponse](../../pkg/models/operations/deleteconfigresponse.md), error**

### Errors

| Error Type             | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| webhooks.ErrorResponse | default                | application/json       |
| sdkerrors.SDKError     | 4XX, 5XX               | \*/\*                  |

## GetDeliveries

List webhook deliveries

### Example Usage

<!-- UsageSnippet language="go" operationID="getDeliveries" method="get" path="/api/webhooks/deliveries" -->
```go
package main

import(
	"context"
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/shared"
	"github.com/formancehq/formance-sdk-go/v5"
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := v5.New(
        v5.WithSecurity(shared.Security{
            ClientID: v5.Pointer("<YOUR_CLIENT_ID_HERE>"),
            ClientSecret: v5.Pointer("<YOUR_CLIENT_SECRET_HERE>"),
        }),
    )

    res, err := s.Webhooks.V1.GetDeliveries(ctx, operations.GetDeliveriesRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.DeliveriesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetDeliveriesRequest](../../pkg/models/operations/getdeliveriesrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.GetDeliveriesResponse](../../pkg/models/operations/getdeliveriesresponse.md), error**

### Errors

| Error Type             | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| webhooks.ErrorResponse | default                | application/json       |
| sdkerrors.SDKError     | 4XX, 5XX               | \*/\*                  |

## GetDelivery

Get a webhook delivery

### Example Usage

<!-- UsageSnippet language="go" operationID="getDelivery" method="get" path="/api/webhooks/deliveries/{id}" -->
```go
package main

import(
	"context"
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/shared"
	"github.com/formancehq/formance-sdk-go/v5"
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := v5.New(
        v5.WithSecurity(shared.Security{
            ClientID: v5.Pointer("<YOUR_CLIENT_ID_HERE>"),
            ClientSecret: v5.Pointer("<YOUR_CLIENT_SECRET_HERE>"),
        }),
    )

    res, err := s.Webhooks.V1.GetDelivery(ctx, operations.GetDeliveryRequest{
        ID: "01e5cac6-75f1-4720-81ca-5563ce22d2e0",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DeliveryResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetDeliveryRequest](../../pkg/models/operations/getdeliveryrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                       | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.GetDeliveryResponse](../../pkg/models/operations/getdeliveryresponse.md), error**

### Errors

| Error Type             | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| webhooks.ErrorResponse | default                | application/json       |
| sdkerrors.SDKError     | 4XX, 5XX               | \*/\*                  |

## GetDeliveryAttempts

List attempts for a webhook delivery

### Example Usage

<!-- UsageSnippet language="go" operationID="getDeliveryAttempts" method="get" path="/api/webhooks/deliveries/{id}/attempts" -->
```go
package main

import(
	"context"
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/shared"
	"github.com/formancehq/formance-sdk-go/v5"
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := v5.New(
        v5.WithSecurity(shared.Security{
            ClientID: v5.Pointer("<YOUR_CLIENT_ID_HERE>"),
            ClientSecret: v5.Pointer("<YOUR_CLIENT_SECRET_HERE>"),
        }),
    )

    res, err := s.Webhooks.V1.GetDeliveryAttempts(ctx, operations.GetDeliveryAttemptsRequest{
        ID: "967e7a38-b11b-4809-92cf-6789e24dbe13",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DeliveryAttemptsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.GetDeliveryAttemptsRequest](../../pkg/models/operations/getdeliveryattemptsrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.GetDeliveryAttemptsResponse](../../pkg/models/operations/getdeliveryattemptsresponse.md), error**

### Errors

| Error Type             | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| webhooks.ErrorResponse | default                | application/json       |
| sdkerrors.SDKError     | 4XX, 5XX               | \*/\*                  |

## GetManyConfigs

Sorted by updated date descending

### Example Usage

<!-- UsageSnippet language="go" operationID="getManyConfigs" method="get" path="/api/webhooks/configs" -->
```go
package main

import(
	"context"
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/shared"
	"github.com/formancehq/formance-sdk-go/v5"
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := v5.New(
        v5.WithSecurity(shared.Security{
            ClientID: v5.Pointer("<YOUR_CLIENT_ID_HERE>"),
            ClientSecret: v5.Pointer("<YOUR_CLIENT_SECRET_HERE>"),
        }),
    )

    res, err := s.Webhooks.V1.GetManyConfigs(ctx, operations.GetManyConfigsRequest{
        Endpoint: v5.Pointer("https://example.com"),
        ID: v5.Pointer("4997257d-dfb6-445b-929c-cbe2ab182818"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ConfigsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetManyConfigsRequest](../../pkg/models/operations/getmanyconfigsrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.GetManyConfigsResponse](../../pkg/models/operations/getmanyconfigsresponse.md), error**

### Errors

| Error Type             | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| webhooks.ErrorResponse | default                | application/json       |
| sdkerrors.SDKError     | 4XX, 5XX               | \*/\*                  |

## InsertConfig

Insert a new webhooks config.

The endpoint should be a valid https URL and be unique.

The secret is the endpoint's verification secret.
If not passed or empty, a secret is automatically generated.
The format is a random string of bytes of size 24, base64 encoded. (larger size after encoding)

All eventTypes are converted to lower-case when inserted.


### Example Usage

<!-- UsageSnippet language="go" operationID="insertConfig" method="post" path="/api/webhooks/configs" -->
```go
package main

import(
	"context"
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/shared"
	"github.com/formancehq/formance-sdk-go/v5"
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/webhooks"
	"log"
)

func main() {
    ctx := context.Background()

    s := v5.New(
        v5.WithSecurity(shared.Security{
            ClientID: v5.Pointer("<YOUR_CLIENT_ID_HERE>"),
            ClientSecret: v5.Pointer("<YOUR_CLIENT_SECRET_HERE>"),
        }),
    )

    res, err := s.Webhooks.V1.InsertConfig(ctx, webhooks.ConfigUser{
        Endpoint: "https://example.com",
        EventTypes: []string{
            "TYPE1",
            "TYPE2",
        },
        Secret: v5.Pointer("V0bivxRWveaoz08afqjU6Ko/jwO0Cb+3"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ConfigResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                      | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `ctx`                                                          | [context.Context](https://pkg.go.dev/context#Context)          | :heavy_check_mark:                                             | The context to use for the request.                            |
| `request`                                                      | [webhooks.ConfigUser](../../pkg/models/webhooks/configuser.md) | :heavy_check_mark:                                             | The request object to use for the request.                     |
| `opts`                                                         | [][operations.Option](../../pkg/models/operations/option.md)   | :heavy_minus_sign:                                             | The options for this request.                                  |

### Response

**[*operations.InsertConfigResponse](../../pkg/models/operations/insertconfigresponse.md), error**

### Errors

| Error Type             | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| webhooks.ErrorResponse | default                | application/json       |
| sdkerrors.SDKError     | 4XX, 5XX               | \*/\*                  |

## ReplayDeliveries

Replay a page of failed or pending deliveries

### Example Usage

<!-- UsageSnippet language="go" operationID="replayDeliveries" method="post" path="/api/webhooks/deliveries/replay" -->
```go
package main

import(
	"context"
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/shared"
	"github.com/formancehq/formance-sdk-go/v5"
	"github.com/formancehq/formance-sdk-go/v5/pkg/types"
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/webhooks"
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := v5.New(
        v5.WithSecurity(shared.Security{
            ClientID: v5.Pointer("<YOUR_CLIENT_ID_HERE>"),
            ClientSecret: v5.Pointer("<YOUR_CLIENT_SECRET_HERE>"),
        }),
    )

    res, err := s.Webhooks.V1.ReplayDeliveries(ctx, operations.ReplayDeliveriesRequest{
        IdempotencyKey: "<value>",
        ReplayDeliveriesRequest: webhooks.ReplayDeliveriesRequest{
            CreatedAtFrom: types.MustTimeFromString("2026-10-16T11:02:44.647Z"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ReplayDeliveriesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ReplayDeliveriesRequest](../../pkg/models/operations/replaydeliveriesrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.ReplayDeliveriesResponse](../../pkg/models/operations/replaydeliveriesresponse.md), error**

### Errors

| Error Type             | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| webhooks.ErrorResponse | default                | application/json       |
| sdkerrors.SDKError     | 4XX, 5XX               | \*/\*                  |

## ReplayDelivery

Replay one failed or pending delivery

### Example Usage

<!-- UsageSnippet language="go" operationID="replayDelivery" method="post" path="/api/webhooks/deliveries/{id}/replay" -->
```go
package main

import(
	"context"
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/shared"
	"github.com/formancehq/formance-sdk-go/v5"
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := v5.New(
        v5.WithSecurity(shared.Security{
            ClientID: v5.Pointer("<YOUR_CLIENT_ID_HERE>"),
            ClientSecret: v5.Pointer("<YOUR_CLIENT_SECRET_HERE>"),
        }),
    )

    res, err := s.Webhooks.V1.ReplayDelivery(ctx, operations.ReplayDeliveryRequest{
        IdempotencyKey: "<value>",
        ID: "06a0d0bb-48de-45f0-b12f-6458a3a41bbe",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DeliveryResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.ReplayDeliveryRequest](../../pkg/models/operations/replaydeliveryrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.ReplayDeliveryResponse](../../pkg/models/operations/replaydeliveryresponse.md), error**

### Errors

| Error Type             | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| webhooks.ErrorResponse | default                | application/json       |
| sdkerrors.SDKError     | 4XX, 5XX               | \*/\*                  |

## TestConfig

Test a config by sending a webhook to its endpoint.

### Example Usage

<!-- UsageSnippet language="go" operationID="testConfig" method="get" path="/api/webhooks/configs/{id}/test" -->
```go
package main

import(
	"context"
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/shared"
	"github.com/formancehq/formance-sdk-go/v5"
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := v5.New(
        v5.WithSecurity(shared.Security{
            ClientID: v5.Pointer("<YOUR_CLIENT_ID_HERE>"),
            ClientSecret: v5.Pointer("<YOUR_CLIENT_SECRET_HERE>"),
        }),
    )

    res, err := s.Webhooks.V1.TestConfig(ctx, operations.TestConfigRequest{
        ID: "4997257d-dfb6-445b-929c-cbe2ab182818",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AttemptResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.TestConfigRequest](../../pkg/models/operations/testconfigrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                     | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.TestConfigResponse](../../pkg/models/operations/testconfigresponse.md), error**

### Errors

| Error Type             | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| webhooks.ErrorResponse | default                | application/json       |
| sdkerrors.SDKError     | 4XX, 5XX               | \*/\*                  |

## UpdateConfig

Update a webhooks config by ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateConfig" method="put" path="/api/webhooks/configs/{id}" -->
```go
package main

import(
	"context"
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/shared"
	"github.com/formancehq/formance-sdk-go/v5"
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/webhooks"
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := v5.New(
        v5.WithSecurity(shared.Security{
            ClientID: v5.Pointer("<YOUR_CLIENT_ID_HERE>"),
            ClientSecret: v5.Pointer("<YOUR_CLIENT_SECRET_HERE>"),
        }),
    )

    res, err := s.Webhooks.V1.UpdateConfig(ctx, operations.UpdateConfigRequest{
        ConfigUser: webhooks.ConfigUser{
            Endpoint: "https://example.com",
            EventTypes: []string{
                "TYPE1",
                "TYPE2",
            },
            Secret: v5.Pointer("V0bivxRWveaoz08afqjU6Ko/jwO0Cb+3"),
        },
        ID: "4997257d-dfb6-445b-929c-cbe2ab182818",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.UpdateConfigRequest](../../pkg/models/operations/updateconfigrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                         | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.UpdateConfigResponse](../../pkg/models/operations/updateconfigresponse.md), error**

### Errors

| Error Type             | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| webhooks.ErrorResponse | default                | application/json       |
| sdkerrors.SDKError     | 4XX, 5XX               | \*/\*                  |