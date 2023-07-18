# Ledger

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
	"context"
	"log"
	"github.com/formancehq/formance-sdk-go"
	"github.com/formancehq/formance-sdk-go/pkg/models/operations"
	"github.com/formancehq/formance-sdk-go/pkg/models/shared"
	"github.com/formancehq/formance-sdk-go/pkg/types"
)

func main() {
    s := formance.New(
        formance.WithSecurity(shared.Security{
            Authorization: "Bearer YOUR_ACCESS_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Ledger.CreateTransactions(ctx, operations.CreateTransactionsRequest{
        Transactions: shared.Transactions{
            Transactions: []shared.TransactionData{
                shared.TransactionData{
                    Metadata: map[string]interface{}{
                        "possimus": "aut",
                        "quasi": "error",
                        "temporibus": "laborum",
                    },
                    Postings: []shared.Posting{
                        shared.Posting{
                            Amount: 100,
                            Asset: "COIN",
                            Destination: "users:002",
                            Source: "users:001",
                        },
                    },
                    Reference: formance.String("ref:001"),
                    Timestamp: types.MustTimeFromString("2020-01-27T18:38:42.890Z"),
                },
                shared.TransactionData{
                    Metadata: map[string]interface{}{
                        "nihil": "praesentium",
                        "voluptatibus": "ipsa",
                        "omnis": "voluptate",
                        "cum": "perferendis",
                    },
                    Postings: []shared.Posting{
                        shared.Posting{
                            Amount: 100,
                            Asset: "COIN",
                            Destination: "users:002",
                            Source: "users:001",
                        },
                    },
                    Reference: formance.String("ref:001"),
                    Timestamp: types.MustTimeFromString("2022-09-19T18:36:39.009Z"),
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

## AddMetadataOnTransaction

Set the metadata of a transaction by its ID

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/formancehq/formance-sdk-go"
	"github.com/formancehq/formance-sdk-go/pkg/models/operations"
)

func main() {
    s := formance.New(
        formance.WithSecurity(shared.Security{
            Authorization: "Bearer YOUR_ACCESS_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Ledger.AddMetadataOnTransaction(ctx, operations.AddMetadataOnTransactionRequest{
        RequestBody: map[string]interface{}{
            "dicta": "corporis",
            "dolore": "iusto",
            "dicta": "harum",
            "enim": "accusamus",
        },
        Ledger: "ledger001",
        Txid: 1234,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

## AddMetadataToAccount

Add metadata to an account

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/formancehq/formance-sdk-go"
	"github.com/formancehq/formance-sdk-go/pkg/models/operations"
)

func main() {
    s := formance.New(
        formance.WithSecurity(shared.Security{
            Authorization: "Bearer YOUR_ACCESS_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Ledger.AddMetadataToAccount(ctx, operations.AddMetadataToAccountRequest{
        RequestBody: map[string]interface{}{
            "repudiandae": "quae",
            "ipsum": "quidem",
        },
        Address: "users:001",
        Ledger: "ledger001",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

## CountAccounts

Count the accounts from a ledger

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/formancehq/formance-sdk-go"
	"github.com/formancehq/formance-sdk-go/pkg/models/operations"
)

func main() {
    s := formance.New(
        formance.WithSecurity(shared.Security{
            Authorization: "Bearer YOUR_ACCESS_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Ledger.CountAccounts(ctx, operations.CountAccountsRequest{
        Address: formance.String("users:.+"),
        Ledger: "ledger001",
        Metadata: map[string]interface{}{
            "excepturi": "pariatur",
            "modi": "praesentium",
            "rem": "voluptates",
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

## CountTransactions

Count the transactions from a ledger

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/formancehq/formance-sdk-go"
	"github.com/formancehq/formance-sdk-go/pkg/models/operations"
	"github.com/formancehq/formance-sdk-go/pkg/types"
)

func main() {
    s := formance.New(
        formance.WithSecurity(shared.Security{
            Authorization: "Bearer YOUR_ACCESS_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Ledger.CountTransactions(ctx, operations.CountTransactionsRequest{
        Account: formance.String("users:001"),
        Destination: formance.String("users:001"),
        EndTime: types.MustTimeFromString("2022-01-29T18:39:33.469Z"),
        Ledger: "ledger001",
        Metadata: map[string]interface{}{
            "veritatis": "itaque",
            "incidunt": "enim",
            "consequatur": "est",
        },
        Reference: formance.String("ref:001"),
        Source: formance.String("users:001"),
        StartTime: types.MustTimeFromString("2022-08-09T16:21:07.003Z"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

## CreateTransaction

Create a new transaction to a ledger

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/formancehq/formance-sdk-go"
	"github.com/formancehq/formance-sdk-go/pkg/models/operations"
	"github.com/formancehq/formance-sdk-go/pkg/models/shared"
	"github.com/formancehq/formance-sdk-go/pkg/types"
)

func main() {
    s := formance.New(
        formance.WithSecurity(shared.Security{
            Authorization: "Bearer YOUR_ACCESS_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Ledger.CreateTransaction(ctx, operations.CreateTransactionRequest{
        PostTransaction: shared.PostTransaction{
            Metadata: map[string]interface{}{
                "distinctio": "quibusdam",
                "labore": "modi",
                "qui": "aliquid",
            },
            Postings: []shared.Posting{
                shared.Posting{
                    Amount: 100,
                    Asset: "COIN",
                    Destination: "users:002",
                    Source: "users:001",
                },
                shared.Posting{
                    Amount: 100,
                    Asset: "COIN",
                    Destination: "users:002",
                    Source: "users:001",
                },
                shared.Posting{
                    Amount: 100,
                    Asset: "COIN",
                    Destination: "users:002",
                    Source: "users:001",
                },
            },
            Reference: formance.String("ref:001"),
            Script: &shared.PostTransactionScript{
                Plain: "vars {
            account $user
            }
            send [COIN 10] (
            	source = @world
            	destination = $user
            )
            ",
                Vars: map[string]interface{}{
                    "perferendis": "magni",
                    "assumenda": "ipsam",
                    "alias": "fugit",
                },
            },
            Timestamp: types.MustTimeFromString("2021-11-11T04:17:07.569Z"),
        },
        Ledger: "ledger001",
        Preview: formance.Bool(true),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TransactionsResponse != nil {
        // handle response
    }
}
```

## GetAccount

Get account by its address

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/formancehq/formance-sdk-go"
	"github.com/formancehq/formance-sdk-go/pkg/models/operations"
)

func main() {
    s := formance.New(
        formance.WithSecurity(shared.Security{
            Authorization: "Bearer YOUR_ACCESS_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Ledger.GetAccount(ctx, operations.GetAccountRequest{
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

## GetBalances

Get the balances from a ledger's account

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/formancehq/formance-sdk-go"
	"github.com/formancehq/formance-sdk-go/pkg/models/operations"
)

func main() {
    s := formance.New(
        formance.WithSecurity(shared.Security{
            Authorization: "Bearer YOUR_ACCESS_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Ledger.GetBalances(ctx, operations.GetBalancesRequest{
        Address: formance.String("users:001"),
        After: formance.String("users:003"),
        Cursor: formance.String("aHR0cHM6Ly9nLnBhZ2UvTmVrby1SYW1lbj9zaGFyZQ=="),
        Ledger: "ledger001",
        PaginationToken: formance.String("aHR0cHM6Ly9nLnBhZ2UvTmVrby1SYW1lbj9zaGFyZQ=="),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.BalancesCursorResponse != nil {
        // handle response
    }
}
```

## GetBalancesAggregated

Get the aggregated balances from selected accounts

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/formancehq/formance-sdk-go"
	"github.com/formancehq/formance-sdk-go/pkg/models/operations"
)

func main() {
    s := formance.New(
        formance.WithSecurity(shared.Security{
            Authorization: "Bearer YOUR_ACCESS_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Ledger.GetBalancesAggregated(ctx, operations.GetBalancesAggregatedRequest{
        Address: formance.String("users:001"),
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

## GetInfo

Show server information

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/formancehq/formance-sdk-go"
)

func main() {
    s := formance.New(
        formance.WithSecurity(shared.Security{
            Authorization: "Bearer YOUR_ACCESS_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Ledger.GetInfo(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.ConfigInfoResponse != nil {
        // handle response
    }
}
```

## GetLedgerInfo

Get information about a ledger

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/formancehq/formance-sdk-go"
	"github.com/formancehq/formance-sdk-go/pkg/models/operations"
)

func main() {
    s := formance.New(
        formance.WithSecurity(shared.Security{
            Authorization: "Bearer YOUR_ACCESS_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Ledger.GetLedgerInfo(ctx, operations.GetLedgerInfoRequest{
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

## GetMapping

Get the mapping of a ledger

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/formancehq/formance-sdk-go"
	"github.com/formancehq/formance-sdk-go/pkg/models/operations"
)

func main() {
    s := formance.New(
        formance.WithSecurity(shared.Security{
            Authorization: "Bearer YOUR_ACCESS_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Ledger.GetMapping(ctx, operations.GetMappingRequest{
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

## GetTransaction

Get transaction from a ledger by its ID

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/formancehq/formance-sdk-go"
	"github.com/formancehq/formance-sdk-go/pkg/models/operations"
)

func main() {
    s := formance.New(
        formance.WithSecurity(shared.Security{
            Authorization: "Bearer YOUR_ACCESS_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Ledger.GetTransaction(ctx, operations.GetTransactionRequest{
        Ledger: "ledger001",
        Txid: 1234,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TransactionResponse != nil {
        // handle response
    }
}
```

## ListAccounts

List accounts from a ledger, sorted by address in descending order.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/formancehq/formance-sdk-go"
	"github.com/formancehq/formance-sdk-go/pkg/models/operations"
	"math/big"
)

func main() {
    s := formance.New(
        formance.WithSecurity(shared.Security{
            Authorization: "Bearer YOUR_ACCESS_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Ledger.ListAccounts(ctx, operations.ListAccountsRequest{
        Address: formance.String("users:.+"),
        After: formance.String("users:003"),
        Balance: big.NewInt(2400),
        BalanceOperator: operations.ListAccountsBalanceOperatorGte.ToPointer(),
        Cursor: formance.String("aHR0cHM6Ly9nLnBhZ2UvTmVrby1SYW1lbj9zaGFyZQ=="),
        Ledger: "ledger001",
        Metadata: map[string]interface{}{
            "facilis": "tempore",
            "labore": "delectus",
        },
        PageSize: formance.Int64(433288),
        PaginationToken: formance.String("aHR0cHM6Ly9nLnBhZ2UvTmVrby1SYW1lbj9zaGFyZQ=="),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AccountsCursorResponse != nil {
        // handle response
    }
}
```

## ListLogs

List the logs from a ledger, sorted by ID in descending order.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/formancehq/formance-sdk-go"
	"github.com/formancehq/formance-sdk-go/pkg/models/operations"
	"github.com/formancehq/formance-sdk-go/pkg/types"
)

func main() {
    s := formance.New(
        formance.WithSecurity(shared.Security{
            Authorization: "Bearer YOUR_ACCESS_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Ledger.ListLogs(ctx, operations.ListLogsRequest{
        After: formance.String("1234"),
        Cursor: formance.String("aHR0cHM6Ly9nLnBhZ2UvTmVrby1SYW1lbj9zaGFyZQ=="),
        EndTime: types.MustTimeFromString("2022-03-31T00:30:19.135Z"),
        Ledger: "ledger001",
        PageSize: formance.Int64(576157),
        PaginationToken: formance.String("aHR0cHM6Ly9nLnBhZ2UvTmVrby1SYW1lbj9zaGFyZQ=="),
        StartTime: types.MustTimeFromString("2022-05-29T21:42:45.399Z"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LogsCursorResponse != nil {
        // handle response
    }
}
```

## ListTransactions

List transactions from a ledger, sorted by txid in descending order.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/formancehq/formance-sdk-go"
	"github.com/formancehq/formance-sdk-go/pkg/models/operations"
	"github.com/formancehq/formance-sdk-go/pkg/types"
)

func main() {
    s := formance.New(
        formance.WithSecurity(shared.Security{
            Authorization: "Bearer YOUR_ACCESS_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Ledger.ListTransactions(ctx, operations.ListTransactionsRequest{
        Account: formance.String("users:001"),
        After: formance.String("1234"),
        Cursor: formance.String("aHR0cHM6Ly9nLnBhZ2UvTmVrby1SYW1lbj9zaGFyZQ=="),
        Destination: formance.String("users:001"),
        EndTime: types.MustTimeFromString("2021-04-14T09:13:11.675Z"),
        Ledger: "ledger001",
        Metadata: map[string]interface{}{
            "dolor": "debitis",
            "a": "dolorum",
            "in": "in",
        },
        PageSize: formance.Int64(846409),
        PaginationToken: formance.String("aHR0cHM6Ly9nLnBhZ2UvTmVrby1SYW1lbj9zaGFyZQ=="),
        Reference: formance.String("ref:001"),
        Source: formance.String("users:001"),
        StartTime: types.MustTimeFromString("2020-11-26T01:41:04.216Z"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TransactionsCursorResponse != nil {
        // handle response
    }
}
```

## ReadStats

Get statistics from a ledger. (aggregate metrics on accounts and transactions)


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/formancehq/formance-sdk-go"
	"github.com/formancehq/formance-sdk-go/pkg/models/operations"
)

func main() {
    s := formance.New(
        formance.WithSecurity(shared.Security{
            Authorization: "Bearer YOUR_ACCESS_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Ledger.ReadStats(ctx, operations.ReadStatsRequest{
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

## RevertTransaction

Revert a ledger transaction by its ID

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/formancehq/formance-sdk-go"
	"github.com/formancehq/formance-sdk-go/pkg/models/operations"
)

func main() {
    s := formance.New(
        formance.WithSecurity(shared.Security{
            Authorization: "Bearer YOUR_ACCESS_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Ledger.RevertTransaction(ctx, operations.RevertTransactionRequest{
        Ledger: "ledger001",
        Txid: 1234,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TransactionResponse != nil {
        // handle response
    }
}
```

## ~~RunScript~~

This route is deprecated, and has been merged into `POST /{ledger}/transactions`.


> :warning: **DEPRECATED**: this method will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/formancehq/formance-sdk-go"
	"github.com/formancehq/formance-sdk-go/pkg/models/operations"
	"github.com/formancehq/formance-sdk-go/pkg/models/shared"
)

func main() {
    s := formance.New(
        formance.WithSecurity(shared.Security{
            Authorization: "Bearer YOUR_ACCESS_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Ledger.RunScript(ctx, operations.RunScriptRequest{
        Script: shared.Script{
            Metadata: map[string]interface{}{
                "magnam": "cumque",
            },
            Plain: "vars {
        account $user
        }
        send [COIN 10] (
        	source = @world
        	destination = $user
        )
        ",
            Reference: formance.String("order_1234"),
            Vars: map[string]interface{}{
                "ea": "aliquid",
                "laborum": "accusamus",
                "non": "occaecati",
                "enim": "accusamus",
            },
        },
        Ledger: "ledger001",
        Preview: formance.Bool(true),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ScriptResponse != nil {
        // handle response
    }
}
```

## UpdateMapping

Update the mapping of a ledger

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/formancehq/formance-sdk-go"
	"github.com/formancehq/formance-sdk-go/pkg/models/operations"
	"github.com/formancehq/formance-sdk-go/pkg/models/shared"
)

func main() {
    s := formance.New(
        formance.WithSecurity(shared.Security{
            Authorization: "Bearer YOUR_ACCESS_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Ledger.UpdateMapping(ctx, operations.UpdateMappingRequest{
        Mapping: shared.Mapping{
            Contracts: []shared.Contract{
                shared.Contract{
                    Account: formance.String("users:001"),
                    Expr: map[string]interface{}{
                        "provident": "nam",
                        "id": "blanditiis",
                        "deleniti": "sapiente",
                    },
                },
                shared.Contract{
                    Account: formance.String("users:001"),
                    Expr: map[string]interface{}{
                        "deserunt": "nisi",
                    },
                },
                shared.Contract{
                    Account: formance.String("users:001"),
                    Expr: map[string]interface{}{
                        "natus": "omnis",
                        "molestiae": "perferendis",
                    },
                },
                shared.Contract{
                    Account: formance.String("users:001"),
                    Expr: map[string]interface{}{
                        "magnam": "distinctio",
                        "id": "labore",
                    },
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
