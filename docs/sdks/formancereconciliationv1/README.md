# Reconciliation.V1

## Overview

### Available Operations

* [AcceptAlert](#acceptalert) - Accept an alert (accepted_by_business)
* [AckAlert](#ackalert) - Acknowledge an alert
* [CreatePolicy](#createpolicy) - Create a policy
* [CreateRule](#createrule) - Create a rule
* [DeletePolicy](#deletepolicy) - Delete a policy
* [DeleteRule](#deleterule) - Delete a rule (cascades to evaluations + alerts + alert events)
* [EvaluateRule](#evaluaterule) - Evaluate a rule now
* [GetAlert](#getalert) - Get an alert
* [GetEvaluation](#getevaluation) - Get an evaluation
* [GetPolicy](#getpolicy) - Get a policy
* [GetReconciliation](#getreconciliation) - Get a reconciliation
* [GetRule](#getrule) - Get a rule
* [GetServerInfoReconciliation](#getserverinforeconciliation) - Get server info
* [ListAlertEvents](#listalertevents) - List alert events (append-only timeline)
* [ListAlerts](#listalerts) - List alerts
* [ListEvaluations](#listevaluations) - List evaluations
* [ListPolicies](#listpolicies) - List policies
* [ListReconciliations](#listreconciliations) - List reconciliations
* [ListRules](#listrules) - List rules
* [PatchRule](#patchrule) - Patch a rule (partial update)
* [Reconcile](#reconcile) - Reconcile using a policy
* [ResolveAlert](#resolvealert) - Resolve an alert (fixed_by_booking)
* [SnoozeAlert](#snoozealert) - Snooze an alert's notifications until a future instant
* [UnsnoozeAlert](#unsnoozealert) - Lift a snooze early

## AcceptAlert

Accept an alert (accepted_by_business)

### Example Usage

<!-- UsageSnippet language="go" operationID="acceptAlert" method="post" path="/api/reconciliation/alerts/{alertID}/accept" -->
```go
package main

import(
	"context"
	"github.com/formancehq/formance-sdk-go/v4/pkg/models/shared"
	"github.com/formancehq/formance-sdk-go/v4"
	"github.com/formancehq/formance-sdk-go/v4/pkg/models/reconciliation"
	"github.com/formancehq/formance-sdk-go/v4/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := v4.New(
        v4.WithSecurity(shared.Security{
            ClientID: v4.Pointer("<YOUR_CLIENT_ID_HERE>"),
            ClientSecret: v4.Pointer("<YOUR_CLIENT_SECRET_HERE>"),
        }),
    )

    res, err := s.Reconciliation.V1.AcceptAlert(ctx, operations.AcceptAlertRequest{
        AcceptAlertRequest: reconciliation.AcceptAlertRequest{
            By: "<value>",
            Note: "<value>",
        },
        AlertID: "5550ef95-072d-4bbb-9d3b-6a9dd307b2bd",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AlertResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.AcceptAlertRequest](../../pkg/models/operations/acceptalertrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                       | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.AcceptAlertResponse](../../pkg/models/operations/acceptalertresponse.md), error**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| reconciliation.ErrorResponse | default                      | application/json             |
| sdkerrors.SDKError           | 4XX, 5XX                     | \*/\*                        |

## AckAlert

Acknowledge an alert

### Example Usage

<!-- UsageSnippet language="go" operationID="ackAlert" method="post" path="/api/reconciliation/alerts/{alertID}/ack" -->
```go
package main

import(
	"context"
	"github.com/formancehq/formance-sdk-go/v4/pkg/models/shared"
	"github.com/formancehq/formance-sdk-go/v4"
	"github.com/formancehq/formance-sdk-go/v4/pkg/models/reconciliation"
	"github.com/formancehq/formance-sdk-go/v4/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := v4.New(
        v4.WithSecurity(shared.Security{
            ClientID: v4.Pointer("<YOUR_CLIENT_ID_HERE>"),
            ClientSecret: v4.Pointer("<YOUR_CLIENT_SECRET_HERE>"),
        }),
    )

    res, err := s.Reconciliation.V1.AckAlert(ctx, operations.AckAlertRequest{
        AckAlertRequest: reconciliation.AckAlertRequest{
            By: "ops@buildr.com",
        },
        AlertID: "5439ab64-6482-49fb-993f-3411bfe19fef",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AlertResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.AckAlertRequest](../../pkg/models/operations/ackalertrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `opts`                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                 | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.AckAlertResponse](../../pkg/models/operations/ackalertresponse.md), error**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| reconciliation.ErrorResponse | default                      | application/json             |
| sdkerrors.SDKError           | 4XX, 5XX                     | \*/\*                        |

## CreatePolicy

Create a policy

### Example Usage

<!-- UsageSnippet language="go" operationID="createPolicy" method="post" path="/api/reconciliation/policies" -->
```go
package main

import(
	"context"
	"github.com/formancehq/formance-sdk-go/v4/pkg/models/shared"
	"github.com/formancehq/formance-sdk-go/v4"
	"github.com/formancehq/formance-sdk-go/v4/pkg/models/reconciliation"
	"log"
)

func main() {
    ctx := context.Background()

    s := v4.New(
        v4.WithSecurity(shared.Security{
            ClientID: v4.Pointer("<YOUR_CLIENT_ID_HERE>"),
            ClientSecret: v4.Pointer("<YOUR_CLIENT_SECRET_HERE>"),
        }),
    )

    res, err := s.Reconciliation.V1.CreatePolicy(ctx, reconciliation.PolicyRequest{
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

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [reconciliation.PolicyRequest](../../pkg/models/reconciliation/policyrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                     | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.CreatePolicyResponse](../../pkg/models/operations/createpolicyresponse.md), error**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| reconciliation.ErrorResponse | default                      | application/json             |
| sdkerrors.SDKError           | 4XX, 5XX                     | \*/\*                        |

## CreateRule

Create a rule

### Example Usage

<!-- UsageSnippet language="go" operationID="createRule" method="post" path="/api/reconciliation/rules" -->
```go
package main

import(
	"context"
	"github.com/formancehq/formance-sdk-go/v4/pkg/models/shared"
	"github.com/formancehq/formance-sdk-go/v4"
	"github.com/formancehq/formance-sdk-go/v4/pkg/models/reconciliation"
	"log"
)

func main() {
    ctx := context.Background()

    s := v4.New(
        v4.WithSecurity(shared.Security{
            ClientID: v4.Pointer("<YOUR_CLIENT_ID_HERE>"),
            ClientSecret: v4.Pointer("<YOUR_CLIENT_SECRET_HERE>"),
        }),
    )

    res, err := s.Reconciliation.V1.CreateRule(ctx, reconciliation.RuleRequest{
        Name: "<value>",
        Schedule: &reconciliation.Schedule{
            Expr: v4.Pointer("*/15 * * * *"),
            Kind: reconciliation.ScheduleKindOnDemand,
            SafetyMargin: v4.Pointer("30s"),
            Tz: v4.Pointer("UTC"),
        },
        TemplateKind: reconciliation.TemplateKindLedgerVsPoolDrift,
        TemplateSpec: map[string]any{
            "key": "<value>",
            "key1": "<value>",
            "key2": "<value>",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RuleResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [reconciliation.RuleRequest](../../pkg/models/reconciliation/rulerequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `opts`                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                 | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.CreateRuleResponse](../../pkg/models/operations/createruleresponse.md), error**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| reconciliation.ErrorResponse | default                      | application/json             |
| sdkerrors.SDKError           | 4XX, 5XX                     | \*/\*                        |

## DeletePolicy

Delete a policy by its id.

### Example Usage

<!-- UsageSnippet language="go" operationID="deletePolicy" method="delete" path="/api/reconciliation/policies/{policyID}" -->
```go
package main

import(
	"context"
	"github.com/formancehq/formance-sdk-go/v4/pkg/models/shared"
	"github.com/formancehq/formance-sdk-go/v4"
	"github.com/formancehq/formance-sdk-go/v4/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := v4.New(
        v4.WithSecurity(shared.Security{
            ClientID: v4.Pointer("<YOUR_CLIENT_ID_HERE>"),
            ClientSecret: v4.Pointer("<YOUR_CLIENT_SECRET_HERE>"),
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

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| reconciliation.ErrorResponse | default                      | application/json             |
| sdkerrors.SDKError           | 4XX, 5XX                     | \*/\*                        |

## DeleteRule

Delete a rule (cascades to evaluations + alerts + alert events)

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteRule" method="delete" path="/api/reconciliation/rules/{ruleID}" -->
```go
package main

import(
	"context"
	"github.com/formancehq/formance-sdk-go/v4/pkg/models/shared"
	"github.com/formancehq/formance-sdk-go/v4"
	"github.com/formancehq/formance-sdk-go/v4/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := v4.New(
        v4.WithSecurity(shared.Security{
            ClientID: v4.Pointer("<YOUR_CLIENT_ID_HERE>"),
            ClientSecret: v4.Pointer("<YOUR_CLIENT_SECRET_HERE>"),
        }),
    )

    res, err := s.Reconciliation.V1.DeleteRule(ctx, operations.DeleteRuleRequest{
        RuleID: "3254b217-2184-4bf4-bbc8-b529fa29bd7c",
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

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.DeleteRuleRequest](../../pkg/models/operations/deleterulerequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                     | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.DeleteRuleResponse](../../pkg/models/operations/deleteruleresponse.md), error**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| reconciliation.ErrorResponse | default                      | application/json             |
| sdkerrors.SDKError           | 4XX, 5XX                     | \*/\*                        |

## EvaluateRule

Evaluate a rule now

### Example Usage

<!-- UsageSnippet language="go" operationID="evaluateRule" method="post" path="/api/reconciliation/rules/{ruleID}/evaluate" -->
```go
package main

import(
	"context"
	"github.com/formancehq/formance-sdk-go/v4/pkg/models/shared"
	"github.com/formancehq/formance-sdk-go/v4"
	"github.com/formancehq/formance-sdk-go/v4/pkg/types"
	"time"
	"github.com/formancehq/formance-sdk-go/v4/pkg/models/reconciliation"
	"github.com/formancehq/formance-sdk-go/v4/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := v4.New(
        v4.WithSecurity(shared.Security{
            ClientID: v4.Pointer("<YOUR_CLIENT_ID_HERE>"),
            ClientSecret: v4.Pointer("<YOUR_CLIENT_SECRET_HERE>"),
        }),
    )

    res, err := s.Reconciliation.V1.EvaluateRule(ctx, operations.EvaluateRuleRequest{
        EvaluateRuleRequest: &reconciliation.EvaluateRuleRequest{
            SafetyMargin: v4.Pointer("30s"),
            SourcePITs: map[string]time.Time{
                "ledger:main#0": types.MustTimeFromString("2026-06-30T23:59:59Z"),
                "pool:acct#0": types.MustTimeFromString("2026-06-30T23:00:00Z"),
            },
        },
        RuleID: "e9d27cb2-b7fc-4383-b319-936c01a66703",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.EvaluationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.EvaluateRuleRequest](../../pkg/models/operations/evaluaterulerequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                         | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.EvaluateRuleResponse](../../pkg/models/operations/evaluateruleresponse.md), error**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| reconciliation.ErrorResponse | 409                          | application/json             |
| reconciliation.ErrorResponse | default                      | application/json             |
| sdkerrors.SDKError           | 4XX, 5XX                     | \*/\*                        |

## GetAlert

Get an alert

### Example Usage

<!-- UsageSnippet language="go" operationID="getAlert" method="get" path="/api/reconciliation/alerts/{alertID}" -->
```go
package main

import(
	"context"
	"github.com/formancehq/formance-sdk-go/v4/pkg/models/shared"
	"github.com/formancehq/formance-sdk-go/v4"
	"github.com/formancehq/formance-sdk-go/v4/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := v4.New(
        v4.WithSecurity(shared.Security{
            ClientID: v4.Pointer("<YOUR_CLIENT_ID_HERE>"),
            ClientSecret: v4.Pointer("<YOUR_CLIENT_SECRET_HERE>"),
        }),
    )

    res, err := s.Reconciliation.V1.GetAlert(ctx, operations.GetAlertRequest{
        AlertID: "c7c54af9-81a4-4208-844b-4f25f89cf8a1",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AlertResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.GetAlertRequest](../../pkg/models/operations/getalertrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `opts`                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                 | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.GetAlertResponse](../../pkg/models/operations/getalertresponse.md), error**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| reconciliation.ErrorResponse | default                      | application/json             |
| sdkerrors.SDKError           | 4XX, 5XX                     | \*/\*                        |

## GetEvaluation

Get an evaluation

### Example Usage

<!-- UsageSnippet language="go" operationID="getEvaluation" method="get" path="/api/reconciliation/evaluations/{evaluationID}" -->
```go
package main

import(
	"context"
	"github.com/formancehq/formance-sdk-go/v4/pkg/models/shared"
	"github.com/formancehq/formance-sdk-go/v4"
	"github.com/formancehq/formance-sdk-go/v4/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := v4.New(
        v4.WithSecurity(shared.Security{
            ClientID: v4.Pointer("<YOUR_CLIENT_ID_HERE>"),
            ClientSecret: v4.Pointer("<YOUR_CLIENT_SECRET_HERE>"),
        }),
    )

    res, err := s.Reconciliation.V1.GetEvaluation(ctx, operations.GetEvaluationRequest{
        EvaluationID: "121717d3-a7d1-444d-9d11-6ea2dc0d3db5",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.EvaluationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetEvaluationRequest](../../pkg/models/operations/getevaluationrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.GetEvaluationResponse](../../pkg/models/operations/getevaluationresponse.md), error**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| reconciliation.ErrorResponse | default                      | application/json             |
| sdkerrors.SDKError           | 4XX, 5XX                     | \*/\*                        |

## GetPolicy

Get a policy

### Example Usage

<!-- UsageSnippet language="go" operationID="getPolicy" method="get" path="/api/reconciliation/policies/{policyID}" -->
```go
package main

import(
	"context"
	"github.com/formancehq/formance-sdk-go/v4/pkg/models/shared"
	"github.com/formancehq/formance-sdk-go/v4"
	"github.com/formancehq/formance-sdk-go/v4/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := v4.New(
        v4.WithSecurity(shared.Security{
            ClientID: v4.Pointer("<YOUR_CLIENT_ID_HERE>"),
            ClientSecret: v4.Pointer("<YOUR_CLIENT_SECRET_HERE>"),
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

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| reconciliation.ErrorResponse | default                      | application/json             |
| sdkerrors.SDKError           | 4XX, 5XX                     | \*/\*                        |

## GetReconciliation

Get a reconciliation

### Example Usage

<!-- UsageSnippet language="go" operationID="getReconciliation" method="get" path="/api/reconciliation/reconciliations/{reconciliationID}" -->
```go
package main

import(
	"context"
	"github.com/formancehq/formance-sdk-go/v4/pkg/models/shared"
	"github.com/formancehq/formance-sdk-go/v4"
	"github.com/formancehq/formance-sdk-go/v4/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := v4.New(
        v4.WithSecurity(shared.Security{
            ClientID: v4.Pointer("<YOUR_CLIENT_ID_HERE>"),
            ClientSecret: v4.Pointer("<YOUR_CLIENT_SECRET_HERE>"),
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

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| reconciliation.ErrorResponse | default                      | application/json             |
| sdkerrors.SDKError           | 4XX, 5XX                     | \*/\*                        |

## GetRule

Get a rule

### Example Usage

<!-- UsageSnippet language="go" operationID="getRule" method="get" path="/api/reconciliation/rules/{ruleID}" -->
```go
package main

import(
	"context"
	"github.com/formancehq/formance-sdk-go/v4/pkg/models/shared"
	"github.com/formancehq/formance-sdk-go/v4"
	"github.com/formancehq/formance-sdk-go/v4/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := v4.New(
        v4.WithSecurity(shared.Security{
            ClientID: v4.Pointer("<YOUR_CLIENT_ID_HERE>"),
            ClientSecret: v4.Pointer("<YOUR_CLIENT_SECRET_HERE>"),
        }),
    )

    res, err := s.Reconciliation.V1.GetRule(ctx, operations.GetRuleRequest{
        RuleID: "fd71d712-041d-4271-b7c5-c9adac177f52",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RuleResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.GetRuleRequest](../../pkg/models/operations/getrulerequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |
| `opts`                                                                     | [][operations.Option](../../pkg/models/operations/option.md)               | :heavy_minus_sign:                                                         | The options for this request.                                              |

### Response

**[*operations.GetRuleResponse](../../pkg/models/operations/getruleresponse.md), error**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| reconciliation.ErrorResponse | default                      | application/json             |
| sdkerrors.SDKError           | 4XX, 5XX                     | \*/\*                        |

## GetServerInfoReconciliation

Get server info

### Example Usage

<!-- UsageSnippet language="go" operationID="getServerInfo_reconciliation" method="get" path="/api/reconciliation/_info" -->
```go
package main

import(
	"context"
	"github.com/formancehq/formance-sdk-go/v4/pkg/models/shared"
	"github.com/formancehq/formance-sdk-go/v4"
	"log"
)

func main() {
    ctx := context.Background()

    s := v4.New(
        v4.WithSecurity(shared.Security{
            ClientID: v4.Pointer("<YOUR_CLIENT_ID_HERE>"),
            ClientSecret: v4.Pointer("<YOUR_CLIENT_SECRET_HERE>"),
        }),
    )

    res, err := s.Reconciliation.V1.GetServerInfoReconciliation(ctx)
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

**[*operations.GetServerInfoReconciliationResponse](../../pkg/models/operations/getserverinforeconciliationresponse.md), error**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| reconciliation.ErrorResponse | default                      | application/json             |
| sdkerrors.SDKError           | 4XX, 5XX                     | \*/\*                        |

## ListAlertEvents

Returns a page of the events recorded for this alert — every evaluation
that touched it plus every manual transition. The list is append-only;
events are never modified or deleted. Ordered most-recent-first and
cursor-paginated: a long-lived alert's timeline is unbounded (one row per
failing evaluation), so callers must page through it.


### Example Usage

<!-- UsageSnippet language="go" operationID="listAlertEvents" method="get" path="/api/reconciliation/alerts/{alertID}/events" -->
```go
package main

import(
	"context"
	"github.com/formancehq/formance-sdk-go/v4/pkg/models/shared"
	"github.com/formancehq/formance-sdk-go/v4"
	"github.com/formancehq/formance-sdk-go/v4/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := v4.New(
        v4.WithSecurity(shared.Security{
            ClientID: v4.Pointer("<YOUR_CLIENT_ID_HERE>"),
            ClientSecret: v4.Pointer("<YOUR_CLIENT_SECRET_HERE>"),
        }),
    )

    res, err := s.Reconciliation.V1.ListAlertEvents(ctx, operations.ListAlertEventsRequest{
        AlertID: "259536e6-acd5-4e38-9154-10e46ea2bc63",
        Cursor: v4.Pointer("aHR0cHM6Ly9nLnBhZ2UvTmVrby1SYW1lbj9zaGFyZQ=="),
        PageSize: v4.Pointer[int64](100),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AlertEventsCursorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ListAlertEventsRequest](../../pkg/models/operations/listalerteventsrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.ListAlertEventsResponse](../../pkg/models/operations/listalerteventsresponse.md), error**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| reconciliation.ErrorResponse | default                      | application/json             |
| sdkerrors.SDKError           | 4XX, 5XX                     | \*/\*                        |

## ListAlerts

List alerts

### Example Usage

<!-- UsageSnippet language="go" operationID="listAlerts" method="get" path="/api/reconciliation/alerts" -->
```go
package main

import(
	"context"
	"github.com/formancehq/formance-sdk-go/v4/pkg/models/shared"
	"github.com/formancehq/formance-sdk-go/v4"
	"github.com/formancehq/formance-sdk-go/v4/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := v4.New(
        v4.WithSecurity(shared.Security{
            ClientID: v4.Pointer("<YOUR_CLIENT_ID_HERE>"),
            ClientSecret: v4.Pointer("<YOUR_CLIENT_SECRET_HERE>"),
        }),
    )

    res, err := s.Reconciliation.V1.ListAlerts(ctx, operations.ListAlertsRequest{
        Cursor: v4.Pointer("aHR0cHM6Ly9nLnBhZ2UvTmVrby1SYW1lbj9zaGFyZQ=="),
        PageSize: v4.Pointer[int64](100),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AlertsCursorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.ListAlertsRequest](../../pkg/models/operations/listalertsrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                     | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.ListAlertsResponse](../../pkg/models/operations/listalertsresponse.md), error**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| reconciliation.ErrorResponse | default                      | application/json             |
| sdkerrors.SDKError           | 4XX, 5XX                     | \*/\*                        |

## ListEvaluations

List evaluations

### Example Usage

<!-- UsageSnippet language="go" operationID="listEvaluations" method="get" path="/api/reconciliation/evaluations" -->
```go
package main

import(
	"context"
	"github.com/formancehq/formance-sdk-go/v4/pkg/models/shared"
	"github.com/formancehq/formance-sdk-go/v4"
	"github.com/formancehq/formance-sdk-go/v4/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := v4.New(
        v4.WithSecurity(shared.Security{
            ClientID: v4.Pointer("<YOUR_CLIENT_ID_HERE>"),
            ClientSecret: v4.Pointer("<YOUR_CLIENT_SECRET_HERE>"),
        }),
    )

    res, err := s.Reconciliation.V1.ListEvaluations(ctx, operations.ListEvaluationsRequest{
        Cursor: v4.Pointer("aHR0cHM6Ly9nLnBhZ2UvTmVrby1SYW1lbj9zaGFyZQ=="),
        PageSize: v4.Pointer[int64](100),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.EvaluationsCursorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ListEvaluationsRequest](../../pkg/models/operations/listevaluationsrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.ListEvaluationsResponse](../../pkg/models/operations/listevaluationsresponse.md), error**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| reconciliation.ErrorResponse | default                      | application/json             |
| sdkerrors.SDKError           | 4XX, 5XX                     | \*/\*                        |

## ListPolicies

List policies

### Example Usage

<!-- UsageSnippet language="go" operationID="listPolicies" method="get" path="/api/reconciliation/policies" -->
```go
package main

import(
	"context"
	"github.com/formancehq/formance-sdk-go/v4/pkg/models/shared"
	"github.com/formancehq/formance-sdk-go/v4"
	"github.com/formancehq/formance-sdk-go/v4/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := v4.New(
        v4.WithSecurity(shared.Security{
            ClientID: v4.Pointer("<YOUR_CLIENT_ID_HERE>"),
            ClientSecret: v4.Pointer("<YOUR_CLIENT_SECRET_HERE>"),
        }),
    )

    res, err := s.Reconciliation.V1.ListPolicies(ctx, operations.ListPoliciesRequest{
        Cursor: v4.Pointer("aHR0cHM6Ly9nLnBhZ2UvTmVrby1SYW1lbj9zaGFyZQ=="),
        PageSize: v4.Pointer[int64](100),
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

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| reconciliation.ErrorResponse | default                      | application/json             |
| sdkerrors.SDKError           | 4XX, 5XX                     | \*/\*                        |

## ListReconciliations

List reconciliations

### Example Usage

<!-- UsageSnippet language="go" operationID="listReconciliations" method="get" path="/api/reconciliation/reconciliations" -->
```go
package main

import(
	"context"
	"github.com/formancehq/formance-sdk-go/v4/pkg/models/shared"
	"github.com/formancehq/formance-sdk-go/v4"
	"github.com/formancehq/formance-sdk-go/v4/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := v4.New(
        v4.WithSecurity(shared.Security{
            ClientID: v4.Pointer("<YOUR_CLIENT_ID_HERE>"),
            ClientSecret: v4.Pointer("<YOUR_CLIENT_SECRET_HERE>"),
        }),
    )

    res, err := s.Reconciliation.V1.ListReconciliations(ctx, operations.ListReconciliationsRequest{
        Cursor: v4.Pointer("aHR0cHM6Ly9nLnBhZ2UvTmVrby1SYW1lbj9zaGFyZQ=="),
        PageSize: v4.Pointer[int64](100),
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

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| reconciliation.ErrorResponse | default                      | application/json             |
| sdkerrors.SDKError           | 4XX, 5XX                     | \*/\*                        |

## ListRules

List rules

### Example Usage

<!-- UsageSnippet language="go" operationID="listRules" method="get" path="/api/reconciliation/rules" -->
```go
package main

import(
	"context"
	"github.com/formancehq/formance-sdk-go/v4/pkg/models/shared"
	"github.com/formancehq/formance-sdk-go/v4"
	"github.com/formancehq/formance-sdk-go/v4/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := v4.New(
        v4.WithSecurity(shared.Security{
            ClientID: v4.Pointer("<YOUR_CLIENT_ID_HERE>"),
            ClientSecret: v4.Pointer("<YOUR_CLIENT_SECRET_HERE>"),
        }),
    )

    res, err := s.Reconciliation.V1.ListRules(ctx, operations.ListRulesRequest{
        Cursor: v4.Pointer("aHR0cHM6Ly9nLnBhZ2UvTmVrby1SYW1lbj9zaGFyZQ=="),
        PageSize: v4.Pointer[int64](100),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RulesCursorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.ListRulesRequest](../../pkg/models/operations/listrulesrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                   | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.ListRulesResponse](../../pkg/models/operations/listrulesresponse.md), error**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| reconciliation.ErrorResponse | default                      | application/json             |
| sdkerrors.SDKError           | 4XX, 5XX                     | \*/\*                        |

## PatchRule

Patch a rule (partial update)

### Example Usage

<!-- UsageSnippet language="go" operationID="patchRule" method="patch" path="/api/reconciliation/rules/{ruleID}" -->
```go
package main

import(
	"context"
	"github.com/formancehq/formance-sdk-go/v4/pkg/models/shared"
	"github.com/formancehq/formance-sdk-go/v4"
	"github.com/formancehq/formance-sdk-go/v4/pkg/models/reconciliation"
	"github.com/formancehq/formance-sdk-go/v4/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := v4.New(
        v4.WithSecurity(shared.Security{
            ClientID: v4.Pointer("<YOUR_CLIENT_ID_HERE>"),
            ClientSecret: v4.Pointer("<YOUR_CLIENT_SECRET_HERE>"),
        }),
    )

    res, err := s.Reconciliation.V1.PatchRule(ctx, operations.PatchRuleRequest{
        RulePatchRequest: reconciliation.RulePatchRequest{
            Schedule: &reconciliation.Schedule{
                Expr: v4.Pointer("*/15 * * * *"),
                Kind: reconciliation.ScheduleKindCron,
                SafetyMargin: v4.Pointer("30s"),
                Tz: v4.Pointer("UTC"),
            },
        },
        RuleID: "0b4aa7b1-cc5d-4700-91ec-4983510fef86",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RuleResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.PatchRuleRequest](../../pkg/models/operations/patchrulerequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                   | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.PatchRuleResponse](../../pkg/models/operations/patchruleresponse.md), error**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| reconciliation.ErrorResponse | default                      | application/json             |
| sdkerrors.SDKError           | 4XX, 5XX                     | \*/\*                        |

## Reconcile

Reconcile using a policy

### Example Usage

<!-- UsageSnippet language="go" operationID="reconcile" method="post" path="/api/reconciliation/policies/{policyID}/reconciliation" -->
```go
package main

import(
	"context"
	"github.com/formancehq/formance-sdk-go/v4/pkg/models/shared"
	"github.com/formancehq/formance-sdk-go/v4"
	"github.com/formancehq/formance-sdk-go/v4/pkg/types"
	"github.com/formancehq/formance-sdk-go/v4/pkg/models/reconciliation"
	"github.com/formancehq/formance-sdk-go/v4/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := v4.New(
        v4.WithSecurity(shared.Security{
            ClientID: v4.Pointer("<YOUR_CLIENT_ID_HERE>"),
            ClientSecret: v4.Pointer("<YOUR_CLIENT_SECRET_HERE>"),
        }),
    )

    res, err := s.Reconciliation.V1.Reconcile(ctx, operations.ReconcileRequest{
        ReconciliationRequest: reconciliation.ReconciliationRequest{
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

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| reconciliation.ErrorResponse | default                      | application/json             |
| sdkerrors.SDKError           | 4XX, 5XX                     | \*/\*                        |

## ResolveAlert

Resolve an alert (fixed_by_booking)

### Example Usage

<!-- UsageSnippet language="go" operationID="resolveAlert" method="post" path="/api/reconciliation/alerts/{alertID}/resolve" -->
```go
package main

import(
	"context"
	"github.com/formancehq/formance-sdk-go/v4/pkg/models/shared"
	"github.com/formancehq/formance-sdk-go/v4"
	"github.com/formancehq/formance-sdk-go/v4/pkg/models/reconciliation"
	"github.com/formancehq/formance-sdk-go/v4/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := v4.New(
        v4.WithSecurity(shared.Security{
            ClientID: v4.Pointer("<YOUR_CLIENT_ID_HERE>"),
            ClientSecret: v4.Pointer("<YOUR_CLIENT_SECRET_HERE>"),
        }),
    )

    res, err := s.Reconciliation.V1.ResolveAlert(ctx, operations.ResolveAlertRequest{
        ResolveAlertRequest: reconciliation.ResolveAlertRequest{
            By: "<value>",
        },
        AlertID: "53527ec3-b39f-4eee-ac1d-6e2bad87f240",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AlertResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.ResolveAlertRequest](../../pkg/models/operations/resolvealertrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                         | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.ResolveAlertResponse](../../pkg/models/operations/resolvealertresponse.md), error**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| reconciliation.ErrorResponse | default                      | application/json             |
| sdkerrors.SDKError           | 4XX, 5XX                     | \*/\*                        |

## SnoozeAlert

Mutes the alert's webhook notifications until `until`. The alert keeps
failing, keeps its status, and keeps counting against period-green —
only its notifications are suppressed, even if the discrepancy changes.
The first failing evaluation at or after `until` clears the snooze and
notifies once. Re-snoozing overwrites the window. Rejects RESOLVED
alerts and a non-future `until`.


### Example Usage

<!-- UsageSnippet language="go" operationID="snoozeAlert" method="post" path="/api/reconciliation/alerts/{alertID}/snooze" -->
```go
package main

import(
	"context"
	"github.com/formancehq/formance-sdk-go/v4/pkg/models/shared"
	"github.com/formancehq/formance-sdk-go/v4"
	"github.com/formancehq/formance-sdk-go/v4/pkg/types"
	"github.com/formancehq/formance-sdk-go/v4/pkg/models/reconciliation"
	"github.com/formancehq/formance-sdk-go/v4/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := v4.New(
        v4.WithSecurity(shared.Security{
            ClientID: v4.Pointer("<YOUR_CLIENT_ID_HERE>"),
            ClientSecret: v4.Pointer("<YOUR_CLIENT_SECRET_HERE>"),
        }),
    )

    res, err := s.Reconciliation.V1.SnoozeAlert(ctx, operations.SnoozeAlertRequest{
        SnoozeAlertRequest: reconciliation.SnoozeAlertRequest{
            By: "ops@buildr.com",
            Until: types.MustTimeFromString("2026-07-17T12:27:27.142Z"),
        },
        AlertID: "96529a25-9005-499e-a0ec-daa0ae32f4cb",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AlertResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.SnoozeAlertRequest](../../pkg/models/operations/snoozealertrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                       | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.SnoozeAlertResponse](../../pkg/models/operations/snoozealertresponse.md), error**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| reconciliation.ErrorResponse | default                      | application/json             |
| sdkerrors.SDKError           | 4XX, 5XX                     | \*/\*                        |

## UnsnoozeAlert

Clears an active snooze before its window elapses. Idempotent —
unsnoozing an alert that is not snoozed returns it unchanged.


### Example Usage

<!-- UsageSnippet language="go" operationID="unsnoozeAlert" method="post" path="/api/reconciliation/alerts/{alertID}/unsnooze" -->
```go
package main

import(
	"context"
	"github.com/formancehq/formance-sdk-go/v4/pkg/models/shared"
	"github.com/formancehq/formance-sdk-go/v4"
	"github.com/formancehq/formance-sdk-go/v4/pkg/models/reconciliation"
	"github.com/formancehq/formance-sdk-go/v4/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := v4.New(
        v4.WithSecurity(shared.Security{
            ClientID: v4.Pointer("<YOUR_CLIENT_ID_HERE>"),
            ClientSecret: v4.Pointer("<YOUR_CLIENT_SECRET_HERE>"),
        }),
    )

    res, err := s.Reconciliation.V1.UnsnoozeAlert(ctx, operations.UnsnoozeAlertRequest{
        UnsnoozeAlertRequest: reconciliation.UnsnoozeAlertRequest{
            By: "ops@buildr.com",
        },
        AlertID: "a1f12fdd-d9de-483a-b3c6-41ec79a76231",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AlertResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.UnsnoozeAlertRequest](../../pkg/models/operations/unsnoozealertrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.UnsnoozeAlertResponse](../../pkg/models/operations/unsnoozealertresponse.md), error**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| reconciliation.ErrorResponse | default                      | application/json             |
| sdkerrors.SDKError           | 4XX, 5XX                     | \*/\*                        |