# Ledger
(*Ledger*)

## Overview

### Available Operations

* [GetInfo](#getinfo) - Show server information
* [GetMetrics](#getmetrics) - Read in memory metrics

## GetInfo

Show server information

### Example Usage

<!-- UsageSnippet language="go" operationID="v2GetInfo" method="get" path="/api/ledger/_/info" -->
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

    res, err := s.Ledger.GetInfo(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.V2ConfigInfoResponse != nil {
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

**[*operations.V2GetInfoResponse](../../pkg/models/operations/v2getinforesponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.V2ErrorResponse | default                   | application/json          |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |

## GetMetrics

Read in memory metrics

### Example Usage

<!-- UsageSnippet language="go" operationID="getMetrics" method="get" path="/api/ledger/_/metrics" -->
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

    res, err := s.Ledger.GetMetrics(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
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

**[*operations.GetMetricsResponse](../../pkg/models/operations/getmetricsresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.V2ErrorResponse | default                   | application/json          |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |