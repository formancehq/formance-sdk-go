# FormanceOrchestrationV1
(*Orchestration.V1*)

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
* [GetWorkflow](#getworkflow) - Get a flow by id
* [ListInstances](#listinstances) - List instances of a workflow
* [ListTriggers](#listtriggers) - List triggers
* [ListTriggersOccurrences](#listtriggersoccurrences) - List triggers occurrences
* [ListWorkflows](#listworkflows) - List registered workflows
* [OrchestrationgetServerInfo](#orchestrationgetserverinfo) - Get server info
* [ReadTrigger](#readtrigger) - Read trigger
* [RunWorkflow](#runworkflow) - Run workflow
* [SendEvent](#sendevent) - Send an event to a running workflow

## CancelEvent

Cancel a running workflow

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

    res, err := s.Orchestration.V1.CancelEvent(ctx, operations.CancelEventRequest{
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
| `request`                                                                          | [operations.CancelEventRequest](../../pkg/models/operations/canceleventrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                       | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.CancelEventResponse](../../pkg/models/operations/canceleventresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | default            | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateTrigger

Create trigger

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

    res, err := s.Orchestration.V1.CreateTrigger(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateTriggerResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                    | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `ctx`                                                        | [context.Context](https://pkg.go.dev/context#Context)        | :heavy_check_mark:                                           | The context to use for the request.                          |
| `request`                                                    | [shared.TriggerData](../../pkg/models/shared/triggerdata.md) | :heavy_check_mark:                                           | The request object to use for the request.                   |
| `opts`                                                       | [][operations.Option](../../pkg/models/operations/option.md) | :heavy_minus_sign:                                           | The options for this request.                                |

### Response

**[*operations.CreateTriggerResponse](../../pkg/models/operations/createtriggerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | default            | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateWorkflow

Create a workflow

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

    res, err := s.Orchestration.V1.CreateWorkflow(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateWorkflowResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [shared.CreateWorkflowRequest](../../pkg/models/shared/createworkflowrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                     | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.CreateWorkflowResponse](../../pkg/models/operations/createworkflowresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | default            | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteTrigger

Read trigger

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

    res, err := s.Orchestration.V1.DeleteTrigger(ctx, operations.DeleteTriggerRequest{
        TriggerID: "<id>",
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
| `request`                                                                              | [operations.DeleteTriggerRequest](../../pkg/models/operations/deletetriggerrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.DeleteTriggerResponse](../../pkg/models/operations/deletetriggerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | default            | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteWorkflow

Delete a flow by id

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

    res, err := s.Orchestration.V1.DeleteWorkflow(ctx, operations.DeleteWorkflowRequest{
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.DeleteWorkflowRequest](../../pkg/models/operations/deleteworkflowrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.DeleteWorkflowResponse](../../pkg/models/operations/deleteworkflowresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | default            | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetInstance

Get a workflow instance by id

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

    res, err := s.Orchestration.V1.GetInstance(ctx, operations.GetInstanceRequest{
        InstanceID: "xxx",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetWorkflowInstanceResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetInstanceRequest](../../pkg/models/operations/getinstancerequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                       | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.GetInstanceResponse](../../pkg/models/operations/getinstanceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | default            | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetInstanceHistory

Get a workflow instance history by id

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

    res, err := s.Orchestration.V1.GetInstanceHistory(ctx, operations.GetInstanceHistoryRequest{
        InstanceID: "xxx",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetWorkflowInstanceHistoryResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.GetInstanceHistoryRequest](../../pkg/models/operations/getinstancehistoryrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.GetInstanceHistoryResponse](../../pkg/models/operations/getinstancehistoryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | default            | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetInstanceStageHistory

Get a workflow instance stage history

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

    res, err := s.Orchestration.V1.GetInstanceStageHistory(ctx, operations.GetInstanceStageHistoryRequest{
        InstanceID: "xxx",
        Number: 0,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetWorkflowInstanceHistoryStageResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.GetInstanceStageHistoryRequest](../../pkg/models/operations/getinstancestagehistoryrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.GetInstanceStageHistoryResponse](../../pkg/models/operations/getinstancestagehistoryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | default            | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetWorkflow

Get a flow by id

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

    res, err := s.Orchestration.V1.GetWorkflow(ctx, operations.GetWorkflowRequest{
        FlowID: "xxx",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetWorkflowResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetWorkflowRequest](../../pkg/models/operations/getworkflowrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                       | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.GetWorkflowResponse](../../pkg/models/operations/getworkflowresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | default            | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListInstances

List instances of a workflow

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

    res, err := s.Orchestration.V1.ListInstances(ctx, operations.ListInstancesRequest{
        Running: formancesdkgo.Bool(true),
        WorkflowID: formancesdkgo.String("xxx"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListRunsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.ListInstancesRequest](../../pkg/models/operations/listinstancesrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.ListInstancesResponse](../../pkg/models/operations/listinstancesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | default            | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListTriggers

List triggers

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

    res, err := s.Orchestration.V1.ListTriggers(ctx, operations.ListTriggersRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.ListTriggersResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.ListTriggersRequest](../../pkg/models/operations/listtriggersrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                         | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.ListTriggersResponse](../../pkg/models/operations/listtriggersresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | default            | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListTriggersOccurrences

List triggers occurrences

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

    res, err := s.Orchestration.V1.ListTriggersOccurrences(ctx, operations.ListTriggersOccurrencesRequest{
        TriggerID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListTriggersOccurrencesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.ListTriggersOccurrencesRequest](../../pkg/models/operations/listtriggersoccurrencesrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.ListTriggersOccurrencesResponse](../../pkg/models/operations/listtriggersoccurrencesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | default            | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListWorkflows

List registered workflows

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

    res, err := s.Orchestration.V1.ListWorkflows(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.ListWorkflowsResponse != nil {
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

**[*operations.ListWorkflowsResponse](../../pkg/models/operations/listworkflowsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | default            | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## OrchestrationgetServerInfo

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

    res, err := s.Orchestration.V1.OrchestrationgetServerInfo(ctx)
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

**[*operations.OrchestrationgetServerInfoResponse](../../pkg/models/operations/orchestrationgetserverinforesponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | default            | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ReadTrigger

Read trigger

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

    res, err := s.Orchestration.V1.ReadTrigger(ctx, operations.ReadTriggerRequest{
        TriggerID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ReadTriggerResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.ReadTriggerRequest](../../pkg/models/operations/readtriggerrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                       | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.ReadTriggerResponse](../../pkg/models/operations/readtriggerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | default            | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RunWorkflow

Run workflow

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

    res, err := s.Orchestration.V1.RunWorkflow(ctx, operations.RunWorkflowRequest{
        WorkflowID: "xxx",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RunWorkflowResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.RunWorkflowRequest](../../pkg/models/operations/runworkflowrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                       | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.RunWorkflowResponse](../../pkg/models/operations/runworkflowresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | default            | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## SendEvent

Send an event to a running workflow

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

    res, err := s.Orchestration.V1.SendEvent(ctx, operations.SendEventRequest{
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

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.SendEventRequest](../../pkg/models/operations/sendeventrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                   | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.SendEventResponse](../../pkg/models/operations/sendeventresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | default            | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |