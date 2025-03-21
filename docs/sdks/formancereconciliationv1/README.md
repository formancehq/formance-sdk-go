# FormanceReconciliationV1
(*Reconciliation.V1*)

## Overview

### Available Operations

* [CreatePolicy](#createpolicy) - Create a policy
* [DeletePolicy](#deletepolicy) - Delete a policy
* [GetPolicy](#getpolicy) - Get a policy
* [GetReconciliation](#getreconciliation) - Get a reconciliation
* [ListPolicies](#listpolicies) - List policies
* [ListReconciliations](#listreconciliations) - List reconciliations
* [Reconcile](#reconcile) - Reconcile using a policy
* [ReconciliationgetServerInfo](#reconciliationgetserverinfo) - Get server info

## CreatePolicy

Create a policy

### Example Usage

```go
package main

import(
	"context"
	formancesdkgo "github.com/formancehq/formance-sdk-go/v3"
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/shared"
	"log"
)

func main() {
    ctx := context.Background()

    s := formancesdkgo.New(
        formancesdkgo.WithSecurity(shared.Security{
            ClientID: formancesdkgo.String("<YOUR_CLIENT_ID_HERE>"),
            ClientSecret: formancesdkgo.String("<YOUR_CLIENT_SECRET_HERE>"),
        }),
    )

    res, err := s.Reconciliation.V1.CreatePolicy(ctx, shared.PolicyRequest{
        LedgerName: "default",
        LedgerQuery: map[string]any{
            "key": "<value>",
        },
        Name: "XXX",
        PaymentsPoolID: "XXX",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PolicyResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                        | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `ctx`                                                            | [context.Context](https://pkg.go.dev/context#Context)            | :heavy_check_mark:                                               | The context to use for the request.                              |
| `request`                                                        | [shared.PolicyRequest](../../pkg/models/shared/policyrequest.md) | :heavy_check_mark:                                               | The request object to use for the request.                       |
| `opts`                                                           | [][operations.Option](../../pkg/models/operations/option.md)     | :heavy_minus_sign:                                               | The options for this request.                                    |

### Response

**[*operations.CreatePolicyResponse](../../pkg/models/operations/createpolicyresponse.md), error**

### Errors

| Error Type                            | Status Code                           | Content Type                          |
| ------------------------------------- | ------------------------------------- | ------------------------------------- |
| sdkerrors.ReconciliationErrorResponse | default                               | application/json                      |
| sdkerrors.SDKError                    | 4XX, 5XX                              | \*/\*                                 |

## DeletePolicy

Delete a policy by its id.

### Example Usage

```go
package main

import(
	"context"
	formancesdkgo "github.com/formancehq/formance-sdk-go/v3"
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/shared"
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := formancesdkgo.New(
        formancesdkgo.WithSecurity(shared.Security{
            ClientID: formancesdkgo.String("<YOUR_CLIENT_ID_HERE>"),
            ClientSecret: formancesdkgo.String("<YOUR_CLIENT_SECRET_HERE>"),
        }),
    )

    res, err := s.Reconciliation.V1.DeletePolicy(ctx, operations.DeletePolicyRequest{
        PolicyID: "XXX",
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
| `request`                                                                            | [operations.DeletePolicyRequest](../../pkg/models/operations/deletepolicyrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                         | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.DeletePolicyResponse](../../pkg/models/operations/deletepolicyresponse.md), error**

### Errors

| Error Type                            | Status Code                           | Content Type                          |
| ------------------------------------- | ------------------------------------- | ------------------------------------- |
| sdkerrors.ReconciliationErrorResponse | default                               | application/json                      |
| sdkerrors.SDKError                    | 4XX, 5XX                              | \*/\*                                 |

## GetPolicy

Get a policy

### Example Usage

```go
package main

import(
	"context"
	formancesdkgo "github.com/formancehq/formance-sdk-go/v3"
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/shared"
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := formancesdkgo.New(
        formancesdkgo.WithSecurity(shared.Security{
            ClientID: formancesdkgo.String("<YOUR_CLIENT_ID_HERE>"),
            ClientSecret: formancesdkgo.String("<YOUR_CLIENT_SECRET_HERE>"),
        }),
    )

    res, err := s.Reconciliation.V1.GetPolicy(ctx, operations.GetPolicyRequest{
        PolicyID: "XXX",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PolicyResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.GetPolicyRequest](../../pkg/models/operations/getpolicyrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                   | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.GetPolicyResponse](../../pkg/models/operations/getpolicyresponse.md), error**

### Errors

| Error Type                            | Status Code                           | Content Type                          |
| ------------------------------------- | ------------------------------------- | ------------------------------------- |
| sdkerrors.ReconciliationErrorResponse | default                               | application/json                      |
| sdkerrors.SDKError                    | 4XX, 5XX                              | \*/\*                                 |

## GetReconciliation

Get a reconciliation

### Example Usage

```go
package main

import(
	"context"
	formancesdkgo "github.com/formancehq/formance-sdk-go/v3"
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/shared"
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := formancesdkgo.New(
        formancesdkgo.WithSecurity(shared.Security{
            ClientID: formancesdkgo.String("<YOUR_CLIENT_ID_HERE>"),
            ClientSecret: formancesdkgo.String("<YOUR_CLIENT_SECRET_HERE>"),
        }),
    )

    res, err := s.Reconciliation.V1.GetReconciliation(ctx, operations.GetReconciliationRequest{
        ReconciliationID: "XXX",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ReconciliationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.GetReconciliationRequest](../../pkg/models/operations/getreconciliationrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.GetReconciliationResponse](../../pkg/models/operations/getreconciliationresponse.md), error**

### Errors

| Error Type                            | Status Code                           | Content Type                          |
| ------------------------------------- | ------------------------------------- | ------------------------------------- |
| sdkerrors.ReconciliationErrorResponse | default                               | application/json                      |
| sdkerrors.SDKError                    | 4XX, 5XX                              | \*/\*                                 |

## ListPolicies

List policies

### Example Usage

```go
package main

import(
	"context"
	formancesdkgo "github.com/formancehq/formance-sdk-go/v3"
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/shared"
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := formancesdkgo.New(
        formancesdkgo.WithSecurity(shared.Security{
            ClientID: formancesdkgo.String("<YOUR_CLIENT_ID_HERE>"),
            ClientSecret: formancesdkgo.String("<YOUR_CLIENT_SECRET_HERE>"),
        }),
    )

    res, err := s.Reconciliation.V1.ListPolicies(ctx, operations.ListPoliciesRequest{
        Cursor: formancesdkgo.String("aHR0cHM6Ly9nLnBhZ2UvTmVrby1SYW1lbj9zaGFyZQ=="),
        PageSize: formancesdkgo.Int64(100),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PoliciesCursorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.ListPoliciesRequest](../../pkg/models/operations/listpoliciesrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                         | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.ListPoliciesResponse](../../pkg/models/operations/listpoliciesresponse.md), error**

### Errors

| Error Type                            | Status Code                           | Content Type                          |
| ------------------------------------- | ------------------------------------- | ------------------------------------- |
| sdkerrors.ReconciliationErrorResponse | default                               | application/json                      |
| sdkerrors.SDKError                    | 4XX, 5XX                              | \*/\*                                 |

## ListReconciliations

List reconciliations

### Example Usage

```go
package main

import(
	"context"
	formancesdkgo "github.com/formancehq/formance-sdk-go/v3"
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/shared"
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := formancesdkgo.New(
        formancesdkgo.WithSecurity(shared.Security{
            ClientID: formancesdkgo.String("<YOUR_CLIENT_ID_HERE>"),
            ClientSecret: formancesdkgo.String("<YOUR_CLIENT_SECRET_HERE>"),
        }),
    )

    res, err := s.Reconciliation.V1.ListReconciliations(ctx, operations.ListReconciliationsRequest{
        Cursor: formancesdkgo.String("aHR0cHM6Ly9nLnBhZ2UvTmVrby1SYW1lbj9zaGFyZQ=="),
        PageSize: formancesdkgo.Int64(100),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ReconciliationsCursorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.ListReconciliationsRequest](../../pkg/models/operations/listreconciliationsrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.ListReconciliationsResponse](../../pkg/models/operations/listreconciliationsresponse.md), error**

### Errors

| Error Type                            | Status Code                           | Content Type                          |
| ------------------------------------- | ------------------------------------- | ------------------------------------- |
| sdkerrors.ReconciliationErrorResponse | default                               | application/json                      |
| sdkerrors.SDKError                    | 4XX, 5XX                              | \*/\*                                 |

## Reconcile

Reconcile using a policy

### Example Usage

```go
package main

import(
	"context"
	formancesdkgo "github.com/formancehq/formance-sdk-go/v3"
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/shared"
	"github.com/formancehq/formance-sdk-go/v3/pkg/types"
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := formancesdkgo.New(
        formancesdkgo.WithSecurity(shared.Security{
            ClientID: formancesdkgo.String("<YOUR_CLIENT_ID_HERE>"),
            ClientSecret: formancesdkgo.String("<YOUR_CLIENT_SECRET_HERE>"),
        }),
    )

    res, err := s.Reconciliation.V1.Reconcile(ctx, operations.ReconcileRequest{
        ReconciliationRequest: shared.ReconciliationRequest{
            ReconciledAtLedger: types.MustTimeFromString("2021-01-01T00:00:00.000Z"),
            ReconciledAtPayments: types.MustTimeFromString("2021-01-01T00:00:00.000Z"),
        },
        PolicyID: "XXX",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ReconciliationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.ReconcileRequest](../../pkg/models/operations/reconcilerequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                   | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.ReconcileResponse](../../pkg/models/operations/reconcileresponse.md), error**

### Errors

| Error Type                            | Status Code                           | Content Type                          |
| ------------------------------------- | ------------------------------------- | ------------------------------------- |
| sdkerrors.ReconciliationErrorResponse | default                               | application/json                      |
| sdkerrors.SDKError                    | 4XX, 5XX                              | \*/\*                                 |

## ReconciliationgetServerInfo

Get server info

### Example Usage

```go
package main

import(
	"context"
	formancesdkgo "github.com/formancehq/formance-sdk-go/v3"
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/shared"
	"log"
)

func main() {
    ctx := context.Background()

    s := formancesdkgo.New(
        formancesdkgo.WithSecurity(shared.Security{
            ClientID: formancesdkgo.String("<YOUR_CLIENT_ID_HERE>"),
            ClientSecret: formancesdkgo.String("<YOUR_CLIENT_SECRET_HERE>"),
        }),
    )

    res, err := s.Reconciliation.V1.ReconciliationgetServerInfo(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.ServerInfo != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                    | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `ctx`                                                        | [context.Context](https://pkg.go.dev/context#Context)        | :heavy_check_mark:                                           | The context to use for the request.                          |
| `opts`                                                       | [][operations.Option](../../pkg/models/operations/option.md) | :heavy_minus_sign:                                           | The options for this request.                                |

### Response

**[*operations.ReconciliationgetServerInfoResponse](../../pkg/models/operations/reconciliationgetserverinforesponse.md), error**

### Errors

| Error Type                            | Status Code                           | Content Type                          |
| ------------------------------------- | ------------------------------------- | ------------------------------------- |
| sdkerrors.ReconciliationErrorResponse | default                               | application/json                      |
| sdkerrors.SDKError                    | 4XX, 5XX                              | \*/\*                                 |