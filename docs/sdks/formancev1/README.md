# FormanceV1
(*Ledger.V1*)

## Overview

### Available Operations

* [CreateTransactions](#createtransactions) - Create a new batch of transactions to a ledger
* [AddMetadataOnTransaction](#addmetadataontransaction) - Set the metadata of a transaction by its ID
* [AddMetadataToAccount](#addmetadatatoaccount) - Add metadata to an account
* [CountAccounts](#countaccounts) - Count the accounts from a ledger
* [CountTransactions](#counttransactions) - Count the transactions from a ledger
* [CreateTransaction](#createtransaction) - Create a new transaction to a ledger
* [GetAccount](#getaccount) - Get account by its address
* [GetBalances](#getbalances) - Get the balances from a ledger's account
* [GetBalancesAggregated](#getbalancesaggregated) - Get the aggregated balances from selected accounts
* [GetInfo](#getinfo) - Show server information
* [GetLedgerInfo](#getledgerinfo) - Get information about a ledger
* [GetMapping](#getmapping) - Get the mapping of a ledger
* [GetTransaction](#gettransaction) - Get transaction from a ledger by its ID
* [ListAccounts](#listaccounts) - List accounts from a ledger
* [ListLogs](#listlogs) - List the logs from a ledger
* [ListTransactions](#listtransactions) - List transactions from a ledger
* [ReadStats](#readstats) - Get statistics from a ledger
* [RevertTransaction](#reverttransaction) - Revert a ledger transaction by its ID
* [~~RunScript~~](#runscript) - Execute a Numscript :warning: **Deprecated**
* [UpdateMapping](#updatemapping) - Update the mapping of a ledger

## CreateTransactions

Create a new batch of transactions to a ledger

### Example Usage

```go
package main

import(
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/shared"
	formancesdkgo "github.com/formancehq/formance-sdk-go/v3"
	"context"
	"math/big"
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/operations"
	"log"
)

func main() {
    s := formancesdkgo.New(
        formancesdkgo.WithSecurity(shared.Security{
            ClientID: formancesdkgo.String("<YOUR_CLIENT_ID_HERE>"),
            ClientSecret: formancesdkgo.String("<YOUR_CLIENT_SECRET_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.Ledger.V1.CreateTransactions(ctx, operations.CreateTransactionsRequest{
        Transactions: shared.Transactions{
            Transactions: []shared.TransactionData{
                shared.TransactionData{
                    Postings: []shared.Posting{
                        shared.Posting{
                            Amount: big.NewInt(100),
                            Asset: "COIN",
                            Destination: "users:002",
                            Source: "users:001",
                        },
                        shared.Posting{
                            Amount: big.NewInt(100),
                            Asset: "COIN",
                            Destination: "users:002",
                            Source: "users:001",
                        },
                        shared.Posting{
                            Amount: big.NewInt(100),
                            Asset: "COIN",
                            Destination: "users:002",
                            Source: "users:001",
                        },
                    },
                    Reference: formancesdkgo.String("ref:001"),
                },
            },
        },
        Ledger: "ledger001",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TransactionsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.CreateTransactionsRequest](../../pkg/models/operations/createtransactionsrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.CreateTransactionsResponse](../../pkg/models/operations/createtransactionsresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.ErrorResponse | default                 | application/json        |
| sdkerrors.SDKError      | 4XX, 5XX                | \*/\*                   |

## AddMetadataOnTransaction

Set the metadata of a transaction by its ID

### Example Usage

```go
package main

import(
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/shared"
	formancesdkgo "github.com/formancehq/formance-sdk-go/v3"
	"context"
	"math/big"
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/operations"
	"log"
)

func main() {
    s := formancesdkgo.New(
        formancesdkgo.WithSecurity(shared.Security{
            ClientID: formancesdkgo.String("<YOUR_CLIENT_ID_HERE>"),
            ClientSecret: formancesdkgo.String("<YOUR_CLIENT_SECRET_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.Ledger.V1.AddMetadataOnTransaction(ctx, operations.AddMetadataOnTransactionRequest{
        Ledger: "ledger001",
        Txid: big.NewInt(1234),
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

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.AddMetadataOnTransactionRequest](../../pkg/models/operations/addmetadataontransactionrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.AddMetadataOnTransactionResponse](../../pkg/models/operations/addmetadataontransactionresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.ErrorResponse | default                 | application/json        |
| sdkerrors.SDKError      | 4XX, 5XX                | \*/\*                   |

## AddMetadataToAccount

Add metadata to an account

### Example Usage

```go
package main

import(
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/shared"
	formancesdkgo "github.com/formancehq/formance-sdk-go/v3"
	"context"
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/operations"
	"log"
)

func main() {
    s := formancesdkgo.New(
        formancesdkgo.WithSecurity(shared.Security{
            ClientID: formancesdkgo.String("<YOUR_CLIENT_ID_HERE>"),
            ClientSecret: formancesdkgo.String("<YOUR_CLIENT_SECRET_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.Ledger.V1.AddMetadataToAccount(ctx, operations.AddMetadataToAccountRequest{
        RequestBody: map[string]any{

        },
        Address: "users:001",
        Ledger: "ledger001",
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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.AddMetadataToAccountRequest](../../pkg/models/operations/addmetadatatoaccountrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.AddMetadataToAccountResponse](../../pkg/models/operations/addmetadatatoaccountresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.ErrorResponse | default                 | application/json        |
| sdkerrors.SDKError      | 4XX, 5XX                | \*/\*                   |

## CountAccounts

Count the accounts from a ledger

### Example Usage

```go
package main

import(
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/shared"
	formancesdkgo "github.com/formancehq/formance-sdk-go/v3"
	"context"
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/operations"
	"log"
)

func main() {
    s := formancesdkgo.New(
        formancesdkgo.WithSecurity(shared.Security{
            ClientID: formancesdkgo.String("<YOUR_CLIENT_ID_HERE>"),
            ClientSecret: formancesdkgo.String("<YOUR_CLIENT_SECRET_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.Ledger.V1.CountAccounts(ctx, operations.CountAccountsRequest{
        Address: formancesdkgo.String("users:.+"),
        Ledger: "ledger001",
        Metadata: map[string]any{
            "0": "m",
            "1": "e",
            "2": "t",
            "3": "a",
            "4": "d",
            "5": "a",
            "6": "t",
            "7": "a",
            "8": "[",
            "9": "k",
            "10": "e",
            "11": "y",
            "12": "]",
            "13": "=",
            "14": "v",
            "15": "a",
            "16": "l",
            "17": "u",
            "18": "e",
            "19": "1",
            "20": "&",
            "21": "m",
            "22": "e",
            "23": "t",
            "24": "a",
            "25": "d",
            "26": "a",
            "27": "t",
            "28": "a",
            "29": "[",
            "30": "a",
            "31": ".",
            "32": "n",
            "33": "e",
            "34": "s",
            "35": "t",
            "36": "e",
            "37": "d",
            "38": ".",
            "39": "k",
            "40": "e",
            "41": "y",
            "42": "]",
            "43": "=",
            "44": "v",
            "45": "a",
            "46": "l",
            "47": "u",
            "48": "e",
            "49": "2",
        },
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
| `request`                                                                              | [operations.CountAccountsRequest](../../pkg/models/operations/countaccountsrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.CountAccountsResponse](../../pkg/models/operations/countaccountsresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.ErrorResponse | default                 | application/json        |
| sdkerrors.SDKError      | 4XX, 5XX                | \*/\*                   |

## CountTransactions

Count the transactions from a ledger

### Example Usage

```go
package main

import(
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/shared"
	formancesdkgo "github.com/formancehq/formance-sdk-go/v3"
	"context"
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/operations"
	"log"
)

func main() {
    s := formancesdkgo.New(
        formancesdkgo.WithSecurity(shared.Security{
            ClientID: formancesdkgo.String("<YOUR_CLIENT_ID_HERE>"),
            ClientSecret: formancesdkgo.String("<YOUR_CLIENT_SECRET_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.Ledger.V1.CountTransactions(ctx, operations.CountTransactionsRequest{
        Account: formancesdkgo.String("users:001"),
        Destination: formancesdkgo.String("users:001"),
        Ledger: "ledger001",
        Metadata: &operations.Metadata{},
        Reference: formancesdkgo.String("ref:001"),
        Source: formancesdkgo.String("users:001"),
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.CountTransactionsRequest](../../pkg/models/operations/counttransactionsrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.CountTransactionsResponse](../../pkg/models/operations/counttransactionsresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.ErrorResponse | default                 | application/json        |
| sdkerrors.SDKError      | 4XX, 5XX                | \*/\*                   |

## CreateTransaction

Create a new transaction to a ledger

### Example Usage

```go
package main

import(
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/shared"
	formancesdkgo "github.com/formancehq/formance-sdk-go/v3"
	"context"
	"math/big"
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/operations"
	"log"
)

func main() {
    s := formancesdkgo.New(
        formancesdkgo.WithSecurity(shared.Security{
            ClientID: formancesdkgo.String("<YOUR_CLIENT_ID_HERE>"),
            ClientSecret: formancesdkgo.String("<YOUR_CLIENT_SECRET_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.Ledger.V1.CreateTransaction(ctx, operations.CreateTransactionRequest{
        PostTransaction: shared.PostTransaction{
            Postings: []shared.Posting{
                shared.Posting{
                    Amount: big.NewInt(100),
                    Asset: "COIN",
                    Destination: "users:002",
                    Source: "users:001",
                },
                shared.Posting{
                    Amount: big.NewInt(100),
                    Asset: "COIN",
                    Destination: "users:002",
                    Source: "users:001",
                },
                shared.Posting{
                    Amount: big.NewInt(100),
                    Asset: "COIN",
                    Destination: "users:002",
                    Source: "users:001",
                },
            },
            Reference: formancesdkgo.String("ref:001"),
            Script: &shared.PostTransactionScript{
                Plain: "vars {\n" +
                "account $user\n" +
                "}\n" +
                "send [COIN 10] (\n" +
                "	source = @world\n" +
                "	destination = $user\n" +
                ")\n" +
                "",
                Vars: map[string]any{
                    "user": "users:042",
                },
            },
        },
        Ledger: "ledger001",
        Preview: formancesdkgo.Bool(true),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TransactionsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.CreateTransactionRequest](../../pkg/models/operations/createtransactionrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.CreateTransactionResponse](../../pkg/models/operations/createtransactionresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.ErrorResponse | default                 | application/json        |
| sdkerrors.SDKError      | 4XX, 5XX                | \*/\*                   |

## GetAccount

Get account by its address

### Example Usage

```go
package main

import(
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/shared"
	formancesdkgo "github.com/formancehq/formance-sdk-go/v3"
	"context"
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/operations"
	"log"
)

func main() {
    s := formancesdkgo.New(
        formancesdkgo.WithSecurity(shared.Security{
            ClientID: formancesdkgo.String("<YOUR_CLIENT_ID_HERE>"),
            ClientSecret: formancesdkgo.String("<YOUR_CLIENT_SECRET_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.Ledger.V1.GetAccount(ctx, operations.GetAccountRequest{
        Address: "users:001",
        Ledger: "ledger001",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.GetAccountRequest](../../pkg/models/operations/getaccountrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                     | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.GetAccountResponse](../../pkg/models/operations/getaccountresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.ErrorResponse | default                 | application/json        |
| sdkerrors.SDKError      | 4XX, 5XX                | \*/\*                   |

## GetBalances

Get the balances from a ledger's account

### Example Usage

```go
package main

import(
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/shared"
	formancesdkgo "github.com/formancehq/formance-sdk-go/v3"
	"context"
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/operations"
	"log"
)

func main() {
    s := formancesdkgo.New(
        formancesdkgo.WithSecurity(shared.Security{
            ClientID: formancesdkgo.String("<YOUR_CLIENT_ID_HERE>"),
            ClientSecret: formancesdkgo.String("<YOUR_CLIENT_SECRET_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.Ledger.V1.GetBalances(ctx, operations.GetBalancesRequest{
        Address: formancesdkgo.String("users:001"),
        After: formancesdkgo.String("users:003"),
        Cursor: formancesdkgo.String("aHR0cHM6Ly9nLnBhZ2UvTmVrby1SYW1lbj9zaGFyZQ=="),
        Ledger: "ledger001",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.BalancesCursorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetBalancesRequest](../../pkg/models/operations/getbalancesrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                       | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.GetBalancesResponse](../../pkg/models/operations/getbalancesresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.ErrorResponse | default                 | application/json        |
| sdkerrors.SDKError      | 4XX, 5XX                | \*/\*                   |

## GetBalancesAggregated

Get the aggregated balances from selected accounts

### Example Usage

```go
package main

import(
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/shared"
	formancesdkgo "github.com/formancehq/formance-sdk-go/v3"
	"context"
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/operations"
	"log"
)

func main() {
    s := formancesdkgo.New(
        formancesdkgo.WithSecurity(shared.Security{
            ClientID: formancesdkgo.String("<YOUR_CLIENT_ID_HERE>"),
            ClientSecret: formancesdkgo.String("<YOUR_CLIENT_SECRET_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.Ledger.V1.GetBalancesAggregated(ctx, operations.GetBalancesAggregatedRequest{
        Address: formancesdkgo.String("users:001"),
        Ledger: "ledger001",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AggregateBalancesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.GetBalancesAggregatedRequest](../../pkg/models/operations/getbalancesaggregatedrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.GetBalancesAggregatedResponse](../../pkg/models/operations/getbalancesaggregatedresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.ErrorResponse | default                 | application/json        |
| sdkerrors.SDKError      | 4XX, 5XX                | \*/\*                   |

## GetInfo

Show server information

### Example Usage

```go
package main

import(
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/shared"
	formancesdkgo "github.com/formancehq/formance-sdk-go/v3"
	"context"
	"log"
)

func main() {
    s := formancesdkgo.New(
        formancesdkgo.WithSecurity(shared.Security{
            ClientID: formancesdkgo.String("<YOUR_CLIENT_ID_HERE>"),
            ClientSecret: formancesdkgo.String("<YOUR_CLIENT_SECRET_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.Ledger.V1.GetInfo(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.ConfigInfoResponse != nil {
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

**[*operations.GetInfoResponse](../../pkg/models/operations/getinforesponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.ErrorResponse | default                 | application/json        |
| sdkerrors.SDKError      | 4XX, 5XX                | \*/\*                   |

## GetLedgerInfo

Get information about a ledger

### Example Usage

```go
package main

import(
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/shared"
	formancesdkgo "github.com/formancehq/formance-sdk-go/v3"
	"context"
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/operations"
	"log"
)

func main() {
    s := formancesdkgo.New(
        formancesdkgo.WithSecurity(shared.Security{
            ClientID: formancesdkgo.String("<YOUR_CLIENT_ID_HERE>"),
            ClientSecret: formancesdkgo.String("<YOUR_CLIENT_SECRET_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.Ledger.V1.GetLedgerInfo(ctx, operations.GetLedgerInfoRequest{
        Ledger: "ledger001",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LedgerInfoResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetLedgerInfoRequest](../../pkg/models/operations/getledgerinforequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.GetLedgerInfoResponse](../../pkg/models/operations/getledgerinforesponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.ErrorResponse | default                 | application/json        |
| sdkerrors.SDKError      | 4XX, 5XX                | \*/\*                   |

## GetMapping

Get the mapping of a ledger

### Example Usage

```go
package main

import(
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/shared"
	formancesdkgo "github.com/formancehq/formance-sdk-go/v3"
	"context"
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/operations"
	"log"
)

func main() {
    s := formancesdkgo.New(
        formancesdkgo.WithSecurity(shared.Security{
            ClientID: formancesdkgo.String("<YOUR_CLIENT_ID_HERE>"),
            ClientSecret: formancesdkgo.String("<YOUR_CLIENT_SECRET_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.Ledger.V1.GetMapping(ctx, operations.GetMappingRequest{
        Ledger: "ledger001",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MappingResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.GetMappingRequest](../../pkg/models/operations/getmappingrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                     | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.GetMappingResponse](../../pkg/models/operations/getmappingresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.ErrorResponse | default                 | application/json        |
| sdkerrors.SDKError      | 4XX, 5XX                | \*/\*                   |

## GetTransaction

Get transaction from a ledger by its ID

### Example Usage

```go
package main

import(
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/shared"
	formancesdkgo "github.com/formancehq/formance-sdk-go/v3"
	"context"
	"math/big"
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/operations"
	"log"
)

func main() {
    s := formancesdkgo.New(
        formancesdkgo.WithSecurity(shared.Security{
            ClientID: formancesdkgo.String("<YOUR_CLIENT_ID_HERE>"),
            ClientSecret: formancesdkgo.String("<YOUR_CLIENT_SECRET_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.Ledger.V1.GetTransaction(ctx, operations.GetTransactionRequest{
        Ledger: "ledger001",
        Txid: big.NewInt(1234),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TransactionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetTransactionRequest](../../pkg/models/operations/gettransactionrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.GetTransactionResponse](../../pkg/models/operations/gettransactionresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.ErrorResponse | default                 | application/json        |
| sdkerrors.SDKError      | 4XX, 5XX                | \*/\*                   |

## ListAccounts

List accounts from a ledger, sorted by address in descending order.

### Example Usage

```go
package main

import(
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/shared"
	formancesdkgo "github.com/formancehq/formance-sdk-go/v3"
	"context"
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/operations"
	"log"
)

func main() {
    s := formancesdkgo.New(
        formancesdkgo.WithSecurity(shared.Security{
            ClientID: formancesdkgo.String("<YOUR_CLIENT_ID_HERE>"),
            ClientSecret: formancesdkgo.String("<YOUR_CLIENT_SECRET_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.Ledger.V1.ListAccounts(ctx, operations.ListAccountsRequest{
        Address: formancesdkgo.String("users:.+"),
        After: formancesdkgo.String("users:003"),
        Balance: formancesdkgo.Int64(2400),
        Cursor: formancesdkgo.String("aHR0cHM6Ly9nLnBhZ2UvTmVrby1SYW1lbj9zaGFyZQ=="),
        Ledger: "ledger001",
        Metadata: map[string]any{
            "0": "m",
            "1": "e",
            "2": "t",
            "3": "a",
            "4": "d",
            "5": "a",
            "6": "t",
            "7": "a",
            "8": "[",
            "9": "k",
            "10": "e",
            "11": "y",
            "12": "]",
            "13": "=",
            "14": "v",
            "15": "a",
            "16": "l",
            "17": "u",
            "18": "e",
            "19": "1",
            "20": "&",
            "21": "m",
            "22": "e",
            "23": "t",
            "24": "a",
            "25": "d",
            "26": "a",
            "27": "t",
            "28": "a",
            "29": "[",
            "30": "a",
            "31": ".",
            "32": "n",
            "33": "e",
            "34": "s",
            "35": "t",
            "36": "e",
            "37": "d",
            "38": ".",
            "39": "k",
            "40": "e",
            "41": "y",
            "42": "]",
            "43": "=",
            "44": "v",
            "45": "a",
            "46": "l",
            "47": "u",
            "48": "e",
            "49": "2",
        },
        PageSize: formancesdkgo.Int64(100),
        PaginationToken: formancesdkgo.String("aHR0cHM6Ly9nLnBhZ2UvTmVrby1SYW1lbj9zaGFyZQ=="),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountsCursorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.ListAccountsRequest](../../pkg/models/operations/listaccountsrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                         | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.ListAccountsResponse](../../pkg/models/operations/listaccountsresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.ErrorResponse | default                 | application/json        |
| sdkerrors.SDKError      | 4XX, 5XX                | \*/\*                   |

## ListLogs

List the logs from a ledger, sorted by ID in descending order.

### Example Usage

```go
package main

import(
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/shared"
	formancesdkgo "github.com/formancehq/formance-sdk-go/v3"
	"context"
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/operations"
	"log"
)

func main() {
    s := formancesdkgo.New(
        formancesdkgo.WithSecurity(shared.Security{
            ClientID: formancesdkgo.String("<YOUR_CLIENT_ID_HERE>"),
            ClientSecret: formancesdkgo.String("<YOUR_CLIENT_SECRET_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.Ledger.V1.ListLogs(ctx, operations.ListLogsRequest{
        After: formancesdkgo.String("1234"),
        Cursor: formancesdkgo.String("aHR0cHM6Ly9nLnBhZ2UvTmVrby1SYW1lbj9zaGFyZQ=="),
        Ledger: "ledger001",
        PageSize: formancesdkgo.Int64(100),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LogsCursorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.ListLogsRequest](../../pkg/models/operations/listlogsrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `opts`                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                 | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.ListLogsResponse](../../pkg/models/operations/listlogsresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.ErrorResponse | default                 | application/json        |
| sdkerrors.SDKError      | 4XX, 5XX                | \*/\*                   |

## ListTransactions

List transactions from a ledger, sorted by txid in descending order.

### Example Usage

```go
package main

import(
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/shared"
	formancesdkgo "github.com/formancehq/formance-sdk-go/v3"
	"context"
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/operations"
	"log"
)

func main() {
    s := formancesdkgo.New(
        formancesdkgo.WithSecurity(shared.Security{
            ClientID: formancesdkgo.String("<YOUR_CLIENT_ID_HERE>"),
            ClientSecret: formancesdkgo.String("<YOUR_CLIENT_SECRET_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.Ledger.V1.ListTransactions(ctx, operations.ListTransactionsRequest{
        Account: formancesdkgo.String("users:001"),
        After: formancesdkgo.String("1234"),
        Cursor: formancesdkgo.String("aHR0cHM6Ly9nLnBhZ2UvTmVrby1SYW1lbj9zaGFyZQ=="),
        Destination: formancesdkgo.String("users:001"),
        Ledger: "ledger001",
        PageSize: formancesdkgo.Int64(100),
        Reference: formancesdkgo.String("ref:001"),
        Source: formancesdkgo.String("users:001"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TransactionsCursorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListTransactionsRequest](../../pkg/models/operations/listtransactionsrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.ListTransactionsResponse](../../pkg/models/operations/listtransactionsresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.ErrorResponse | default                 | application/json        |
| sdkerrors.SDKError      | 4XX, 5XX                | \*/\*                   |

## ReadStats

Get statistics from a ledger. (aggregate metrics on accounts and transactions)


### Example Usage

```go
package main

import(
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/shared"
	formancesdkgo "github.com/formancehq/formance-sdk-go/v3"
	"context"
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/operations"
	"log"
)

func main() {
    s := formancesdkgo.New(
        formancesdkgo.WithSecurity(shared.Security{
            ClientID: formancesdkgo.String("<YOUR_CLIENT_ID_HERE>"),
            ClientSecret: formancesdkgo.String("<YOUR_CLIENT_SECRET_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.Ledger.V1.ReadStats(ctx, operations.ReadStatsRequest{
        Ledger: "ledger001",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.StatsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.ReadStatsRequest](../../pkg/models/operations/readstatsrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                   | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.ReadStatsResponse](../../pkg/models/operations/readstatsresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.ErrorResponse | default                 | application/json        |
| sdkerrors.SDKError      | 4XX, 5XX                | \*/\*                   |

## RevertTransaction

Revert a ledger transaction by its ID

### Example Usage

```go
package main

import(
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/shared"
	formancesdkgo "github.com/formancehq/formance-sdk-go/v3"
	"context"
	"math/big"
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/operations"
	"log"
)

func main() {
    s := formancesdkgo.New(
        formancesdkgo.WithSecurity(shared.Security{
            ClientID: formancesdkgo.String("<YOUR_CLIENT_ID_HERE>"),
            ClientSecret: formancesdkgo.String("<YOUR_CLIENT_SECRET_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.Ledger.V1.RevertTransaction(ctx, operations.RevertTransactionRequest{
        Ledger: "ledger001",
        Txid: big.NewInt(1234),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TransactionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.RevertTransactionRequest](../../pkg/models/operations/reverttransactionrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.RevertTransactionResponse](../../pkg/models/operations/reverttransactionresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.ErrorResponse | default                 | application/json        |
| sdkerrors.SDKError      | 4XX, 5XX                | \*/\*                   |

## ~~RunScript~~

This route is deprecated, and has been merged into `POST /{ledger}/transactions`.


> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

```go
package main

import(
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/shared"
	formancesdkgo "github.com/formancehq/formance-sdk-go/v3"
	"context"
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/operations"
	"log"
)

func main() {
    s := formancesdkgo.New(
        formancesdkgo.WithSecurity(shared.Security{
            ClientID: formancesdkgo.String("<YOUR_CLIENT_ID_HERE>"),
            ClientSecret: formancesdkgo.String("<YOUR_CLIENT_SECRET_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.Ledger.V1.RunScript(ctx, operations.RunScriptRequest{
        Script: shared.Script{
            Plain: "vars {\n" +
            "account $user\n" +
            "}\n" +
            "send [COIN 10] (\n" +
            "	source = @world\n" +
            "	destination = $user\n" +
            ")\n" +
            "",
            Reference: formancesdkgo.String("order_1234"),
            Vars: map[string]any{
                "user": "users:042",
            },
        },
        Ledger: "ledger001",
        Preview: formancesdkgo.Bool(true),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ScriptResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.RunScriptRequest](../../pkg/models/operations/runscriptrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                   | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.RunScriptResponse](../../pkg/models/operations/runscriptresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateMapping

Update the mapping of a ledger

### Example Usage

```go
package main

import(
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/shared"
	formancesdkgo "github.com/formancehq/formance-sdk-go/v3"
	"context"
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/operations"
	"log"
)

func main() {
    s := formancesdkgo.New(
        formancesdkgo.WithSecurity(shared.Security{
            ClientID: formancesdkgo.String("<YOUR_CLIENT_ID_HERE>"),
            ClientSecret: formancesdkgo.String("<YOUR_CLIENT_SECRET_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.Ledger.V1.UpdateMapping(ctx, operations.UpdateMappingRequest{
        Mapping: &shared.Mapping{
            Contracts: []shared.Contract{
                shared.Contract{
                    Account: formancesdkgo.String("users:001"),
                    Expr: shared.Expr{},
                },
                shared.Contract{
                    Account: formancesdkgo.String("users:001"),
                    Expr: shared.Expr{},
                },
            },
        },
        Ledger: "ledger001",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MappingResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.UpdateMappingRequest](../../pkg/models/operations/updatemappingrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.UpdateMappingResponse](../../pkg/models/operations/updatemappingresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.ErrorResponse | default                 | application/json        |
| sdkerrors.SDKError      | 4XX, 5XX                | \*/\*                   |