# FormanceV2
(*Orchestration.V2*)

## Overview

### Available Operations

* [CancelEvent](#cancelevent) - Cancel a running workflow
* [CreateTrigger](#createtrigger) - Create trigger
* [CreateWorkflow](#createworkflow) - Create workflow
* [DeleteTrigger](#deletetrigger) - Delete trigger
* [DeleteWorkflow](#deleteworkflow) - Delete a flow by id
* [GetInstance](#getinstance) - Get a workflow instance by id
* [GetInstanceHistory](#getinstancehistory) - Get a workflow instance history by id
* [GetInstanceStageHistory](#getinstancestagehistory) - Get a workflow instance stage history
* [GetServerInfo](#getserverinfo) - Get server info
* [GetWorkflow](#getworkflow) - Get a flow by id
* [ListInstances](#listinstances) - List instances of a workflow
* [ListTriggers](#listtriggers) - List triggers
* [ListTriggersOccurrences](#listtriggersoccurrences) - List triggers occurrences
* [ListWorkflows](#listworkflows) - List registered workflows
* [ReadTrigger](#readtrigger) - Read trigger
* [RunWorkflow](#runworkflow) - Run workflow
* [SendEvent](#sendevent) - Send an event to a running workflow
* [TestTrigger](#testtrigger) - Test trigger

## CancelEvent

Cancel a running workflow

### Example Usage

```go
package main

import(
	"github.com/formancehq/formance-sdk-go/v2/pkg/models/shared"
	formancesdkgo "github.com/formancehq/formance-sdk-go/v2"
	"context"
	"github.com/formancehq/formance-sdk-go/v2/pkg/models/operations"
	"log"
)

func main() {
    s := formancesdkgo.New(
        formancesdkgo.WithSecurity(shared.Security{
            Authorization: "<YOUR_AUTHORIZATION_HERE>",
        }),
    )

    ctx := context.Background()
    res, err := s.Orchestration.V2.CancelEvent(ctx, operations.V2CancelEventRequest{
        InstanceID: "xxx",
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.V2CancelEventRequest](../../pkg/models/operations/v2canceleventrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.V2CancelEventResponse](../../pkg/models/operations/v2canceleventresponse.md), error**

### Errors

| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.V2Error  | default            | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |


## CreateTrigger

Create trigger

### Example Usage

```go
package main

import(
	"github.com/formancehq/formance-sdk-go/v2/pkg/models/shared"
	formancesdkgo "github.com/formancehq/formance-sdk-go/v2"
	"context"
	"log"
)

func main() {
    s := formancesdkgo.New(
        formancesdkgo.WithSecurity(shared.Security{
            Authorization: "<YOUR_AUTHORIZATION_HERE>",
        }),
    )

    ctx := context.Background()
    res, err := s.Orchestration.V2.CreateTrigger(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.V2CreateTriggerResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                        | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `ctx`                                                            | [context.Context](https://pkg.go.dev/context#Context)            | :heavy_check_mark:                                               | The context to use for the request.                              |
| `request`                                                        | [shared.V2TriggerData](../../pkg/models/shared/v2triggerdata.md) | :heavy_check_mark:                                               | The request object to use for the request.                       |
| `opts`                                                           | [][operations.Option](../../pkg/models/operations/option.md)     | :heavy_minus_sign:                                               | The options for this request.                                    |

### Response

**[*operations.V2CreateTriggerResponse](../../pkg/models/operations/v2createtriggerresponse.md), error**

### Errors

| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.V2Error  | default            | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |


## CreateWorkflow

Create a workflow

### Example Usage

```go
package main

import(
	"github.com/formancehq/formance-sdk-go/v2/pkg/models/shared"
	formancesdkgo "github.com/formancehq/formance-sdk-go/v2"
	"context"
	"log"
)

func main() {
    s := formancesdkgo.New(
        formancesdkgo.WithSecurity(shared.Security{
            Authorization: "<YOUR_AUTHORIZATION_HERE>",
        }),
    )

    ctx := context.Background()
    res, err := s.Orchestration.V2.CreateWorkflow(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.V2CreateWorkflowResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [shared.V2CreateWorkflowRequest](../../pkg/models/shared/v2createworkflowrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                         | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.V2CreateWorkflowResponse](../../pkg/models/operations/v2createworkflowresponse.md), error**

### Errors

| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.V2Error  | default            | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |


## DeleteTrigger

Read trigger

### Example Usage

```go
package main

import(
	"github.com/formancehq/formance-sdk-go/v2/pkg/models/shared"
	formancesdkgo "github.com/formancehq/formance-sdk-go/v2"
	"context"
	"github.com/formancehq/formance-sdk-go/v2/pkg/models/operations"
	"log"
)

func main() {
    s := formancesdkgo.New(
        formancesdkgo.WithSecurity(shared.Security{
            Authorization: "<YOUR_AUTHORIZATION_HERE>",
        }),
    )

    ctx := context.Background()
    res, err := s.Orchestration.V2.DeleteTrigger(ctx, operations.V2DeleteTriggerRequest{
        TriggerID: "<value>",
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.V2DeleteTriggerRequest](../../pkg/models/operations/v2deletetriggerrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.V2DeleteTriggerResponse](../../pkg/models/operations/v2deletetriggerresponse.md), error**

### Errors

| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.V2Error  | default            | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |


## DeleteWorkflow

Delete a flow by id

### Example Usage

```go
package main

import(
	"github.com/formancehq/formance-sdk-go/v2/pkg/models/shared"
	formancesdkgo "github.com/formancehq/formance-sdk-go/v2"
	"context"
	"github.com/formancehq/formance-sdk-go/v2/pkg/models/operations"
	"log"
)

func main() {
    s := formancesdkgo.New(
        formancesdkgo.WithSecurity(shared.Security{
            Authorization: "<YOUR_AUTHORIZATION_HERE>",
        }),
    )

    ctx := context.Background()
    res, err := s.Orchestration.V2.DeleteWorkflow(ctx, operations.V2DeleteWorkflowRequest{
        FlowID: "xxx",
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.V2DeleteWorkflowRequest](../../pkg/models/operations/v2deleteworkflowrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.V2DeleteWorkflowResponse](../../pkg/models/operations/v2deleteworkflowresponse.md), error**

### Errors

| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.V2Error  | default            | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |


## GetInstance

Get a workflow instance by id

### Example Usage

```go
package main

import(
	"github.com/formancehq/formance-sdk-go/v2/pkg/models/shared"
	formancesdkgo "github.com/formancehq/formance-sdk-go/v2"
	"context"
	"github.com/formancehq/formance-sdk-go/v2/pkg/models/operations"
	"log"
)

func main() {
    s := formancesdkgo.New(
        formancesdkgo.WithSecurity(shared.Security{
            Authorization: "<YOUR_AUTHORIZATION_HERE>",
        }),
    )

    ctx := context.Background()
    res, err := s.Orchestration.V2.GetInstance(ctx, operations.V2GetInstanceRequest{
        InstanceID: "xxx",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.V2GetWorkflowInstanceResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.V2GetInstanceRequest](../../pkg/models/operations/v2getinstancerequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.V2GetInstanceResponse](../../pkg/models/operations/v2getinstanceresponse.md), error**

### Errors

| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.V2Error  | default            | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |


## GetInstanceHistory

Get a workflow instance history by id

### Example Usage

```go
package main

import(
	"github.com/formancehq/formance-sdk-go/v2/pkg/models/shared"
	formancesdkgo "github.com/formancehq/formance-sdk-go/v2"
	"context"
	"github.com/formancehq/formance-sdk-go/v2/pkg/models/operations"
	"log"
)

func main() {
    s := formancesdkgo.New(
        formancesdkgo.WithSecurity(shared.Security{
            Authorization: "<YOUR_AUTHORIZATION_HERE>",
        }),
    )

    ctx := context.Background()
    res, err := s.Orchestration.V2.GetInstanceHistory(ctx, operations.V2GetInstanceHistoryRequest{
        InstanceID: "xxx",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.V2GetWorkflowInstanceHistoryResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.V2GetInstanceHistoryRequest](../../pkg/models/operations/v2getinstancehistoryrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.V2GetInstanceHistoryResponse](../../pkg/models/operations/v2getinstancehistoryresponse.md), error**

### Errors

| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.V2Error  | default            | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |


## GetInstanceStageHistory

Get a workflow instance stage history

### Example Usage

```go
package main

import(
	"github.com/formancehq/formance-sdk-go/v2/pkg/models/shared"
	formancesdkgo "github.com/formancehq/formance-sdk-go/v2"
	"context"
	"github.com/formancehq/formance-sdk-go/v2/pkg/models/operations"
	"log"
)

func main() {
    s := formancesdkgo.New(
        formancesdkgo.WithSecurity(shared.Security{
            Authorization: "<YOUR_AUTHORIZATION_HERE>",
        }),
    )

    ctx := context.Background()
    res, err := s.Orchestration.V2.GetInstanceStageHistory(ctx, operations.V2GetInstanceStageHistoryRequest{
        InstanceID: "xxx",
        Number: 0,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.V2GetWorkflowInstanceHistoryStageResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.V2GetInstanceStageHistoryRequest](../../pkg/models/operations/v2getinstancestagehistoryrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.V2GetInstanceStageHistoryResponse](../../pkg/models/operations/v2getinstancestagehistoryresponse.md), error**

### Errors

| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.V2Error  | default            | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |


## GetServerInfo

Get server info

### Example Usage

```go
package main

import(
	"github.com/formancehq/formance-sdk-go/v2/pkg/models/shared"
	formancesdkgo "github.com/formancehq/formance-sdk-go/v2"
	"context"
	"log"
)

func main() {
    s := formancesdkgo.New(
        formancesdkgo.WithSecurity(shared.Security{
            Authorization: "<YOUR_AUTHORIZATION_HERE>",
        }),
    )

    ctx := context.Background()
    res, err := s.Orchestration.V2.GetServerInfo(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.V2ServerInfo != nil {
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

**[*operations.V2GetServerInfoResponse](../../pkg/models/operations/v2getserverinforesponse.md), error**

### Errors

| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.V2Error  | default            | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |


## GetWorkflow

Get a flow by id

### Example Usage

```go
package main

import(
	"github.com/formancehq/formance-sdk-go/v2/pkg/models/shared"
	formancesdkgo "github.com/formancehq/formance-sdk-go/v2"
	"context"
	"github.com/formancehq/formance-sdk-go/v2/pkg/models/operations"
	"log"
)

func main() {
    s := formancesdkgo.New(
        formancesdkgo.WithSecurity(shared.Security{
            Authorization: "<YOUR_AUTHORIZATION_HERE>",
        }),
    )

    ctx := context.Background()
    res, err := s.Orchestration.V2.GetWorkflow(ctx, operations.V2GetWorkflowRequest{
        FlowID: "xxx",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.V2GetWorkflowResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.V2GetWorkflowRequest](../../pkg/models/operations/v2getworkflowrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.V2GetWorkflowResponse](../../pkg/models/operations/v2getworkflowresponse.md), error**

### Errors

| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.V2Error  | default            | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |


## ListInstances

List instances of a workflow

### Example Usage

```go
package main

import(
	"github.com/formancehq/formance-sdk-go/v2/pkg/models/shared"
	formancesdkgo "github.com/formancehq/formance-sdk-go/v2"
	"context"
	"github.com/formancehq/formance-sdk-go/v2/pkg/models/operations"
	"log"
)

func main() {
    s := formancesdkgo.New(
        formancesdkgo.WithSecurity(shared.Security{
            Authorization: "<YOUR_AUTHORIZATION_HERE>",
        }),
    )

    ctx := context.Background()
    res, err := s.Orchestration.V2.ListInstances(ctx, operations.V2ListInstancesRequest{
        Cursor: formancesdkgo.String("aHR0cHM6Ly9nLnBhZ2UvTmVrby1SYW1lbj9zaGFyZQ=="),
        PageSize: formancesdkgo.Int64(100),
        Running: formancesdkgo.Bool(true),
        WorkflowID: formancesdkgo.String("xxx"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.V2ListRunsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.V2ListInstancesRequest](../../pkg/models/operations/v2listinstancesrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.V2ListInstancesResponse](../../pkg/models/operations/v2listinstancesresponse.md), error**

### Errors

| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.V2Error  | default            | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |


## ListTriggers

List triggers

### Example Usage

```go
package main

import(
	"github.com/formancehq/formance-sdk-go/v2/pkg/models/shared"
	formancesdkgo "github.com/formancehq/formance-sdk-go/v2"
	"context"
	"github.com/formancehq/formance-sdk-go/v2/pkg/models/operations"
	"log"
)

func main() {
    s := formancesdkgo.New(
        formancesdkgo.WithSecurity(shared.Security{
            Authorization: "<YOUR_AUTHORIZATION_HERE>",
        }),
    )

    ctx := context.Background()
    res, err := s.Orchestration.V2.ListTriggers(ctx, operations.V2ListTriggersRequest{
        Cursor: formancesdkgo.String("aHR0cHM6Ly9nLnBhZ2UvTmVrby1SYW1lbj9zaGFyZQ=="),
        PageSize: formancesdkgo.Int64(100),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.V2ListTriggersResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.V2ListTriggersRequest](../../pkg/models/operations/v2listtriggersrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.V2ListTriggersResponse](../../pkg/models/operations/v2listtriggersresponse.md), error**

### Errors

| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.V2Error  | default            | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |


## ListTriggersOccurrences

List triggers occurrences

### Example Usage

```go
package main

import(
	"github.com/formancehq/formance-sdk-go/v2/pkg/models/shared"
	formancesdkgo "github.com/formancehq/formance-sdk-go/v2"
	"context"
	"github.com/formancehq/formance-sdk-go/v2/pkg/models/operations"
	"log"
)

func main() {
    s := formancesdkgo.New(
        formancesdkgo.WithSecurity(shared.Security{
            Authorization: "<YOUR_AUTHORIZATION_HERE>",
        }),
    )

    ctx := context.Background()
    res, err := s.Orchestration.V2.ListTriggersOccurrences(ctx, operations.V2ListTriggersOccurrencesRequest{
        Cursor: formancesdkgo.String("aHR0cHM6Ly9nLnBhZ2UvTmVrby1SYW1lbj9zaGFyZQ=="),
        PageSize: formancesdkgo.Int64(100),
        TriggerID: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.V2ListTriggersOccurrencesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.V2ListTriggersOccurrencesRequest](../../pkg/models/operations/v2listtriggersoccurrencesrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.V2ListTriggersOccurrencesResponse](../../pkg/models/operations/v2listtriggersoccurrencesresponse.md), error**

### Errors

| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.V2Error  | default            | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |


## ListWorkflows

List registered workflows

### Example Usage

```go
package main

import(
	"github.com/formancehq/formance-sdk-go/v2/pkg/models/shared"
	formancesdkgo "github.com/formancehq/formance-sdk-go/v2"
	"context"
	"github.com/formancehq/formance-sdk-go/v2/pkg/models/operations"
	"log"
)

func main() {
    s := formancesdkgo.New(
        formancesdkgo.WithSecurity(shared.Security{
            Authorization: "<YOUR_AUTHORIZATION_HERE>",
        }),
    )

    ctx := context.Background()
    res, err := s.Orchestration.V2.ListWorkflows(ctx, operations.V2ListWorkflowsRequest{
        Cursor: formancesdkgo.String("aHR0cHM6Ly9nLnBhZ2UvTmVrby1SYW1lbj9zaGFyZQ=="),
        PageSize: formancesdkgo.Int64(100),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.V2ListWorkflowsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.V2ListWorkflowsRequest](../../pkg/models/operations/v2listworkflowsrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.V2ListWorkflowsResponse](../../pkg/models/operations/v2listworkflowsresponse.md), error**

### Errors

| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.V2Error  | default            | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |


## ReadTrigger

Read trigger

### Example Usage

```go
package main

import(
	"github.com/formancehq/formance-sdk-go/v2/pkg/models/shared"
	formancesdkgo "github.com/formancehq/formance-sdk-go/v2"
	"context"
	"github.com/formancehq/formance-sdk-go/v2/pkg/models/operations"
	"log"
)

func main() {
    s := formancesdkgo.New(
        formancesdkgo.WithSecurity(shared.Security{
            Authorization: "<YOUR_AUTHORIZATION_HERE>",
        }),
    )

    ctx := context.Background()
    res, err := s.Orchestration.V2.ReadTrigger(ctx, operations.V2ReadTriggerRequest{
        TriggerID: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.V2ReadTriggerResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.V2ReadTriggerRequest](../../pkg/models/operations/v2readtriggerrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.V2ReadTriggerResponse](../../pkg/models/operations/v2readtriggerresponse.md), error**

### Errors

| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.V2Error  | default            | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |


## RunWorkflow

Run workflow

### Example Usage

```go
package main

import(
	"github.com/formancehq/formance-sdk-go/v2/pkg/models/shared"
	formancesdkgo "github.com/formancehq/formance-sdk-go/v2"
	"context"
	"github.com/formancehq/formance-sdk-go/v2/pkg/models/operations"
	"log"
)

func main() {
    s := formancesdkgo.New(
        formancesdkgo.WithSecurity(shared.Security{
            Authorization: "<YOUR_AUTHORIZATION_HERE>",
        }),
    )

    ctx := context.Background()
    res, err := s.Orchestration.V2.RunWorkflow(ctx, operations.V2RunWorkflowRequest{
        WorkflowID: "xxx",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.V2RunWorkflowResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.V2RunWorkflowRequest](../../pkg/models/operations/v2runworkflowrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.V2RunWorkflowResponse](../../pkg/models/operations/v2runworkflowresponse.md), error**

### Errors

| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.V2Error  | default            | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |


## SendEvent

Send an event to a running workflow

### Example Usage

```go
package main

import(
	"github.com/formancehq/formance-sdk-go/v2/pkg/models/shared"
	formancesdkgo "github.com/formancehq/formance-sdk-go/v2"
	"context"
	"github.com/formancehq/formance-sdk-go/v2/pkg/models/operations"
	"log"
)

func main() {
    s := formancesdkgo.New(
        formancesdkgo.WithSecurity(shared.Security{
            Authorization: "<YOUR_AUTHORIZATION_HERE>",
        }),
    )

    ctx := context.Background()
    res, err := s.Orchestration.V2.SendEvent(ctx, operations.V2SendEventRequest{
        InstanceID: "xxx",
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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.V2SendEventRequest](../../pkg/models/operations/v2sendeventrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                       | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.V2SendEventResponse](../../pkg/models/operations/v2sendeventresponse.md), error**

### Errors

| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.V2Error  | default            | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |


## TestTrigger

Test trigger

### Example Usage

```go
package main

import(
	"github.com/formancehq/formance-sdk-go/v2/pkg/models/shared"
	formancesdkgo "github.com/formancehq/formance-sdk-go/v2"
	"context"
	"github.com/formancehq/formance-sdk-go/v2/pkg/models/operations"
	"log"
)

func main() {
    s := formancesdkgo.New(
        formancesdkgo.WithSecurity(shared.Security{
            Authorization: "<YOUR_AUTHORIZATION_HERE>",
        }),
    )

    ctx := context.Background()
    res, err := s.Orchestration.V2.TestTrigger(ctx, operations.TestTriggerRequest{
        TriggerID: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.V2TestTriggerResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.TestTriggerRequest](../../pkg/models/operations/testtriggerrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                       | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.TestTriggerResponse](../../pkg/models/operations/testtriggerresponse.md), error**

### Errors

| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.V2Error  | default            | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
