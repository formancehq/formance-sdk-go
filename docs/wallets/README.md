# Wallets

### Available Operations

* [ConfirmHold](#confirmhold) - Confirm a hold
* [CreateBalance](#createbalance) - Create a balance
* [CreateWallet](#createwallet) - Create a new wallet
* [CreditWallet](#creditwallet) - Credit a wallet
* [DebitWallet](#debitwallet) - Debit a wallet
* [GetBalance](#getbalance) - Get detailed balance
* [GetHold](#gethold) - Get a hold
* [GetHolds](#getholds) - Get all holds for a wallet
* [GetTransactions](#gettransactions)
* [GetWallet](#getwallet) - Get a wallet
* [GetWalletSummary](#getwalletsummary) - Get wallet summary
* [ListBalances](#listbalances) - List balances of a wallet
* [ListWallets](#listwallets) - List all wallets
* [UpdateWallet](#updatewallet) - Update a wallet
* [VoidHold](#voidhold) - Cancel a hold
* [WalletsgetServerInfo](#walletsgetserverinfo) - Get server info

## ConfirmHold

Confirm a hold

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/formancehq/formance-sdk-go"
	"github.com/formancehq/formance-sdk-go/pkg/models/operations"
	"github.com/formancehq/formance-sdk-go/pkg/models/shared"
	"math/big"
)

func main() {
    s := formance.New(
        formance.WithSecurity(shared.Security{
            Authorization: "Bearer YOUR_ACCESS_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Wallets.ConfirmHold(ctx, operations.ConfirmHoldRequest{
        ConfirmHoldRequest: &shared.ConfirmHoldRequest{
            Amount: big.NewInt(100),
            Final: formance.Bool(true),
        },
        HoldID: "quibusdam",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

## CreateBalance

Create a balance

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
	"math/big"
)

func main() {
    s := formance.New(
        formance.WithSecurity(shared.Security{
            Authorization: "Bearer YOUR_ACCESS_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Wallets.CreateBalance(ctx, operations.CreateBalanceRequest{
        CreateBalanceRequest: &shared.CreateBalanceRequest{
            ExpiresAt: types.MustTimeFromString("2022-02-04T19:17:08.641Z"),
            Name: "Edward Crooks",
            Priority: big.NewInt(166847),
        },
        ID: "1cddc692-601f-4b57-ab0d-5f0d30c5fbb2",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateBalanceResponse != nil {
        // handle response
    }
}
```

## CreateWallet

Create a new wallet

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/formancehq/formance-sdk-go"
	"github.com/formancehq/formance-sdk-go/pkg/models/shared"
)

func main() {
    s := formance.New(
        formance.WithSecurity(shared.Security{
            Authorization: "Bearer YOUR_ACCESS_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Wallets.CreateWallet(ctx, shared.CreateWalletRequest{
        Metadata: map[string]string{
            "totam": "dignissimos",
            "eaque": "quis",
        },
        Name: "Ruby Auer",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateWalletResponse != nil {
        // handle response
    }
}
```

## CreditWallet

Credit a wallet

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/formancehq/formance-sdk-go"
	"github.com/formancehq/formance-sdk-go/pkg/models/operations"
	"github.com/formancehq/formance-sdk-go/pkg/models/shared"
	"math/big"
)

func main() {
    s := formance.New(
        formance.WithSecurity(shared.Security{
            Authorization: "Bearer YOUR_ACCESS_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Wallets.CreditWallet(ctx, operations.CreditWalletRequest{
        CreditWalletRequest: &shared.CreditWalletRequest{
            Amount: shared.Monetary{
                Amount: big.NewInt(463451),
                Asset: "dolor",
            },
            Balance: formance.String("vero"),
            Metadata: map[string]string{
                "hic": "recusandae",
                "omnis": "facilis",
            },
            Reference: formance.String("perspiciatis"),
            Sources: []shared.Subject{
                shared.Subject{},
            },
        },
        ID: "c28909b3-fe49-4a8d-9cbf-48633323f9b7",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

## DebitWallet

Debit a wallet

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/formancehq/formance-sdk-go"
	"github.com/formancehq/formance-sdk-go/pkg/models/operations"
	"github.com/formancehq/formance-sdk-go/pkg/models/shared"
	"math/big"
)

func main() {
    s := formance.New(
        formance.WithSecurity(shared.Security{
            Authorization: "Bearer YOUR_ACCESS_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Wallets.DebitWallet(ctx, operations.DebitWalletRequest{
        DebitWalletRequest: &shared.DebitWalletRequest{
            Amount: shared.Monetary{
                Amount: big.NewInt(490459),
                Asset: "reiciendis",
            },
            Balances: []string{
                "dolorum",
            },
            Description: formance.String("numquam"),
            Destination: &shared.Subject{},
            Metadata: map[string]string{
                "ipsa": "ipsa",
            },
            Pending: formance.Bool(false),
        },
        ID: "674ebf69-280d-41ba-b7a8-9ebf737ae420",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DebitWalletResponse != nil {
        // handle response
    }
}
```

## GetBalance

Get detailed balance

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
    res, err := s.Wallets.GetBalance(ctx, operations.GetBalanceRequest{
        BalanceName: "amet",
        ID: "ce5e6a95-d8a0-4d44-ace2-af7a73cf3be4",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetBalanceResponse != nil {
        // handle response
    }
}
```

## GetHold

Get a hold

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
    res, err := s.Wallets.GetHold(ctx, operations.GetHoldRequest{
        HoldID: "enim",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetHoldResponse != nil {
        // handle response
    }
}
```

## GetHolds

Get all holds for a wallet

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
    res, err := s.Wallets.GetHolds(ctx, operations.GetHoldsRequest{
        Cursor: formance.String("aHR0cHM6Ly9nLnBhZ2UvTmVrby1SYW1lbj9zaGFyZQ=="),
        Metadata: map[string]string{
            "sapiente": "totam",
        },
        PageSize: formance.Int64(471752),
        WalletID: formance.String("sit"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetHoldsResponse != nil {
        // handle response
    }
}
```

## GetTransactions

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
    res, err := s.Wallets.GetTransactions(ctx, operations.GetTransactionsRequest{
        Cursor: formance.String("aHR0cHM6Ly9nLnBhZ2UvTmVrby1SYW1lbj9zaGFyZQ=="),
        PageSize: formance.Int64(711584),
        WalletID: formance.String("neque"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetTransactionsResponse != nil {
        // handle response
    }
}
```

## GetWallet

Get a wallet

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
    res, err := s.Wallets.GetWallet(ctx, operations.GetWalletRequest{
        ID: "26b5a734-29cd-4b1a-8422-bb679d232271",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetWalletResponse != nil {
        // handle response
    }
}
```

## GetWalletSummary

Get wallet summary

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
    res, err := s.Wallets.GetWalletSummary(ctx, operations.GetWalletSummaryRequest{
        ID: "5bf0cbb1-e31b-48b9-8f34-43a1108e0adc",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetWalletSummaryResponse != nil {
        // handle response
    }
}
```

## ListBalances

List balances of a wallet

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
    res, err := s.Wallets.ListBalances(ctx, operations.ListBalancesRequest{
        ID: "f4b92187-9fce-4953-b73e-f7fbc7abd74d",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ListBalancesResponse != nil {
        // handle response
    }
}
```

## ListWallets

List all wallets

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
    res, err := s.Wallets.ListWallets(ctx, operations.ListWalletsRequest{
        Cursor: formance.String("aHR0cHM6Ly9nLnBhZ2UvTmVrby1SYW1lbj9zaGFyZQ=="),
        Metadata: map[string]string{
            "sequi": "natus",
            "impedit": "aut",
            "voluptatibus": "exercitationem",
            "nulla": "fugit",
        },
        Name: formance.String("Elijah Wyman"),
        PageSize: formance.Int64(497391),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ListWalletsResponse != nil {
        // handle response
    }
}
```

## UpdateWallet

Update a wallet

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
    res, err := s.Wallets.UpdateWallet(ctx, operations.UpdateWalletRequest{
        RequestBody: &operations.UpdateWalletRequestBody{
            Metadata: map[string]string{
                "officia": "tempora",
            },
        },
        ID: "5626d436-813f-416d-9f5f-ce6c556146c3",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

## VoidHold

Cancel a hold

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
    res, err := s.Wallets.VoidHold(ctx, operations.VoidHoldRequest{
        HoldID: "recusandae",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

## WalletsgetServerInfo

Get server info

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
    res, err := s.Wallets.WalletsgetServerInfo(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.ServerInfo != nil {
        // handle response
    }
}
```
