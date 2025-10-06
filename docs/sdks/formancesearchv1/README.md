# FormanceSearchV1
(*Search.V1*)

## Overview

### Available Operations

* [~~Search~~](#search) - search.v1 :warning: **Deprecated**
* [~~SearchgetServerInfo~~](#searchgetserverinfo) - Get server info :warning: **Deprecated**

## ~~Search~~

Elasticsearch.v1 query engine

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

<!-- UsageSnippet language="go" operationID="search" method="post" path="/api/search/" -->
```go
package main

import(
	"context"
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/shared"
	"github.com/formancehq/formance-sdk-go/v3"
	"log"
)

func main() {
    ctx := context.Background()

    s := v3.New(
        v3.WithSecurity(shared.Security{
            ClientID: v3.Pointer("<YOUR_CLIENT_ID_HERE>"),
            ClientSecret: v3.Pointer("<YOUR_CLIENT_SECRET_HERE>"),
        }),
    )

    res, err := s.Search.V1.Search(ctx, shared.Query{
        After: []string{
            "users:002",
        },
        Cursor: v3.Pointer("YXVsdCBhbmQgYSBtYXhpbXVtIG1heF9yZXN1bHRzLol="),
        Ledgers: []string{
            "quickstart",
        },
        Policy: v3.Pointer("OR"),
        Raw: &shared.QueryRaw{},
        Sort: v3.Pointer("id:asc"),
        Terms: []string{
            "destination=central_bank1",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Response != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                    | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `ctx`                                                        | [context.Context](https://pkg.go.dev/context#Context)        | :heavy_check_mark:                                           | The context to use for the request.                          |
| `request`                                                    | [shared.Query](../../pkg/models/shared/query.md)             | :heavy_check_mark:                                           | The request object to use for the request.                   |
| `opts`                                                       | [][operations.Option](../../pkg/models/operations/option.md) | :heavy_minus_sign:                                           | The options for this request.                                |

### Response

**[*operations.SearchResponse](../../pkg/models/operations/searchresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ~~SearchgetServerInfo~~

Get server info

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

<!-- UsageSnippet language="go" operationID="searchgetServerInfo" method="get" path="/api/search/_info" -->
```go
package main

import(
	"context"
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/shared"
	"github.com/formancehq/formance-sdk-go/v3"
	"log"
)

func main() {
    ctx := context.Background()

    s := v3.New(
        v3.WithSecurity(shared.Security{
            ClientID: v3.Pointer("<YOUR_CLIENT_ID_HERE>"),
            ClientSecret: v3.Pointer("<YOUR_CLIENT_SECRET_HERE>"),
        }),
    )

    res, err := s.Search.V1.SearchgetServerInfo(ctx)
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

**[*operations.SearchgetServerInfoResponse](../../pkg/models/operations/searchgetserverinforesponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |