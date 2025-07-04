# FormancePaymentsV1
(*Payments.V1*)

## Overview

### Available Operations

* [AddAccountToPool](#addaccounttopool) - Add an account to a pool
* [ConnectorsTransfer](#connectorstransfer) - Transfer funds between Connector accounts
* [CreateAccount](#createaccount) - Create an account
* [CreateBankAccount](#createbankaccount) - Create a BankAccount in Payments and on the PSP
* [CreatePayment](#createpayment) - Create a payment
* [CreatePool](#createpool) - Create a Pool
* [CreateTransferInitiation](#createtransferinitiation) - Create a TransferInitiation
* [DeletePool](#deletepool) - Delete a Pool
* [DeleteTransferInitiation](#deletetransferinitiation) - Delete a transfer initiation
* [ForwardBankAccount](#forwardbankaccount) - Forward a bank account to a connector
* [GetAccountBalances](#getaccountbalances) - Get account balances
* [GetBankAccount](#getbankaccount) - Get a bank account created by user on Formance
* [~~GetConnectorTask~~](#getconnectortask) - Read a specific task of the connector :warning: **Deprecated**
* [GetConnectorTaskV1](#getconnectortaskv1) - Read a specific task of the connector
* [GetPayment](#getpayment) - Get a payment
* [GetPool](#getpool) - Get a Pool
* [GetPoolBalances](#getpoolbalances) - Get historical pool balances at a particular point in time
* [GetPoolBalancesLatest](#getpoolbalanceslatest) - Get latest pool balances
* [GetTransferInitiation](#gettransferinitiation) - Get a transfer initiation
* [InstallConnector](#installconnector) - Install a connector
* [ListAllConnectors](#listallconnectors) - List all installed connectors
* [ListBankAccounts](#listbankaccounts) - List bank accounts created by user on Formance
* [ListConfigsAvailableConnectors](#listconfigsavailableconnectors) - List the configs of each available connector
* [~~ListConnectorTasks~~](#listconnectortasks) - List tasks from a connector :warning: **Deprecated**
* [ListConnectorTasksV1](#listconnectortasksv1) - List tasks from a connector
* [ListPayments](#listpayments) - List payments
* [ListPools](#listpools) - List Pools
* [ListTransferInitiations](#listtransferinitiations) - List Transfer Initiations
* [PaymentsgetAccount](#paymentsgetaccount) - Get an account
* [PaymentsgetServerInfo](#paymentsgetserverinfo) - Get server info
* [PaymentslistAccounts](#paymentslistaccounts) - List accounts
* [~~ReadConnectorConfig~~](#readconnectorconfig) - Read the config of a connector :warning: **Deprecated**
* [ReadConnectorConfigV1](#readconnectorconfigv1) - Read the config of a connector
* [RemoveAccountFromPool](#removeaccountfrompool) - Remove an account from a pool
* [~~ResetConnector~~](#resetconnector) - Reset a connector :warning: **Deprecated**
* [ResetConnectorV1](#resetconnectorv1) - Reset a connector
* [RetryTransferInitiation](#retrytransferinitiation) - Retry a failed transfer initiation
* [ReverseTransferInitiation](#reversetransferinitiation) - Reverse a transfer initiation
* [~~UninstallConnector~~](#uninstallconnector) - Uninstall a connector :warning: **Deprecated**
* [UninstallConnectorV1](#uninstallconnectorv1) - Uninstall a connector
* [UpdateBankAccountMetadata](#updatebankaccountmetadata) - Update metadata of a bank account
* [UpdateConnectorConfigV1](#updateconnectorconfigv1) - Update the config of a connector
* [UpdateMetadata](#updatemetadata) - Update metadata
* [UpdateTransferInitiationStatus](#updatetransferinitiationstatus) - Update the status of a transfer initiation

## AddAccountToPool

Add an account to a pool

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

    res, err := s.Payments.V1.AddAccountToPool(ctx, operations.AddAccountToPoolRequest{
        AddAccountToPoolRequest: shared.AddAccountToPoolRequest{
            AccountID: "<id>",
        },
        PoolID: "XXX",
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
| `request`                                                                                    | [operations.AddAccountToPoolRequest](../../pkg/models/operations/addaccounttopoolrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.AddAccountToPoolResponse](../../pkg/models/operations/addaccounttopoolresponse.md), error**

### Errors

| Error Type                      | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.PaymentsErrorResponse | default                         | application/json                |
| sdkerrors.SDKError              | 4XX, 5XX                        | \*/\*                           |

## ConnectorsTransfer

Execute a transfer between two accounts.

### Example Usage

```go
package main

import(
	"context"
	formancesdkgo "github.com/formancehq/formance-sdk-go/v3"
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/shared"
	"math/big"
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

    res, err := s.Payments.V1.ConnectorsTransfer(ctx, operations.ConnectorsTransferRequest{
        TransferRequest: shared.TransferRequest{
            Amount: big.NewInt(100),
            Asset: "USD",
            Destination: "acct_1Gqj58KZcSIg2N2q",
            Source: formancesdkgo.String("acct_1Gqj58KZcSIg2N2q"),
        },
        Connector: shared.ConnectorGeneric,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TransferResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.ConnectorsTransferRequest](../../pkg/models/operations/connectorstransferrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.ConnectorsTransferResponse](../../pkg/models/operations/connectorstransferresponse.md), error**

### Errors

| Error Type                      | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.PaymentsErrorResponse | default                         | application/json                |
| sdkerrors.SDKError              | 4XX, 5XX                        | \*/\*                           |

## CreateAccount

Create an account

### Example Usage

```go
package main

import(
	"context"
	formancesdkgo "github.com/formancehq/formance-sdk-go/v3"
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/shared"
	"github.com/formancehq/formance-sdk-go/v3/pkg/types"
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

    res, err := s.Payments.V1.CreateAccount(ctx, shared.AccountRequest{
        ConnectorID: "<id>",
        CreatedAt: types.MustTimeFromString("2025-07-27T08:57:17.388Z"),
        Reference: "<value>",
        Type: shared.AccountTypeUnknown,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PaymentsAccountResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |
| `request`                                                          | [shared.AccountRequest](../../pkg/models/shared/accountrequest.md) | :heavy_check_mark:                                                 | The request object to use for the request.                         |
| `opts`                                                             | [][operations.Option](../../pkg/models/operations/option.md)       | :heavy_minus_sign:                                                 | The options for this request.                                      |

### Response

**[*operations.CreateAccountResponse](../../pkg/models/operations/createaccountresponse.md), error**

### Errors

| Error Type                      | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.PaymentsErrorResponse | default                         | application/json                |
| sdkerrors.SDKError              | 4XX, 5XX                        | \*/\*                           |

## CreateBankAccount

Create a bank account in Payments and on the PSP.

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

    res, err := s.Payments.V1.CreateBankAccount(ctx, shared.BankAccountRequest{
        Country: "GB",
        Name: "My account",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.BankAccountResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [shared.BankAccountRequest](../../pkg/models/shared/bankaccountrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |
| `opts`                                                                     | [][operations.Option](../../pkg/models/operations/option.md)               | :heavy_minus_sign:                                                         | The options for this request.                                              |

### Response

**[*operations.CreateBankAccountResponse](../../pkg/models/operations/createbankaccountresponse.md), error**

### Errors

| Error Type                      | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.PaymentsErrorResponse | default                         | application/json                |
| sdkerrors.SDKError              | 4XX, 5XX                        | \*/\*                           |

## CreatePayment

Create a payment

### Example Usage

```go
package main

import(
	"context"
	formancesdkgo "github.com/formancehq/formance-sdk-go/v3"
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/shared"
	"math/big"
	"github.com/formancehq/formance-sdk-go/v3/pkg/types"
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

    res, err := s.Payments.V1.CreatePayment(ctx, shared.PaymentRequest{
        Amount: big.NewInt(100),
        Asset: "USD",
        ConnectorID: "<id>",
        CreatedAt: types.MustTimeFromString("2025-08-26T06:29:11.777Z"),
        Reference: "<value>",
        Scheme: shared.PaymentSchemeRtp,
        Status: shared.PaymentStatusRefundedFailure,
        Type: shared.PaymentTypePayout,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PaymentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |
| `request`                                                          | [shared.PaymentRequest](../../pkg/models/shared/paymentrequest.md) | :heavy_check_mark:                                                 | The request object to use for the request.                         |
| `opts`                                                             | [][operations.Option](../../pkg/models/operations/option.md)       | :heavy_minus_sign:                                                 | The options for this request.                                      |

### Response

**[*operations.CreatePaymentResponse](../../pkg/models/operations/createpaymentresponse.md), error**

### Errors

| Error Type                      | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.PaymentsErrorResponse | default                         | application/json                |
| sdkerrors.SDKError              | 4XX, 5XX                        | \*/\*                           |

## CreatePool

Create a Pool

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

    res, err := s.Payments.V1.CreatePool(ctx, shared.PoolRequest{
        AccountIDs: []string{
            "<value 1>",
            "<value 2>",
        },
        Name: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PoolResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                    | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `ctx`                                                        | [context.Context](https://pkg.go.dev/context#Context)        | :heavy_check_mark:                                           | The context to use for the request.                          |
| `request`                                                    | [shared.PoolRequest](../../pkg/models/shared/poolrequest.md) | :heavy_check_mark:                                           | The request object to use for the request.                   |
| `opts`                                                       | [][operations.Option](../../pkg/models/operations/option.md) | :heavy_minus_sign:                                           | The options for this request.                                |

### Response

**[*operations.CreatePoolResponse](../../pkg/models/operations/createpoolresponse.md), error**

### Errors

| Error Type                      | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.PaymentsErrorResponse | default                         | application/json                |
| sdkerrors.SDKError              | 4XX, 5XX                        | \*/\*                           |

## CreateTransferInitiation

Create a transfer initiation

### Example Usage

```go
package main

import(
	"context"
	formancesdkgo "github.com/formancehq/formance-sdk-go/v3"
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/shared"
	"math/big"
	"github.com/formancehq/formance-sdk-go/v3/pkg/types"
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

    res, err := s.Payments.V1.CreateTransferInitiation(ctx, shared.TransferInitiationRequest{
        Amount: big.NewInt(83093),
        Asset: "USD",
        Description: "flowery yum keenly operating knavishly commemorate recent apropos",
        DestinationAccountID: "<id>",
        Reference: "XXX",
        ScheduledAt: types.MustTimeFromString("2025-07-09T05:18:01.065Z"),
        SourceAccountID: "<id>",
        Type: shared.TransferInitiationRequestTypeTransfer,
        Validated: false,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TransferInitiationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [shared.TransferInitiationRequest](../../pkg/models/shared/transferinitiationrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.CreateTransferInitiationResponse](../../pkg/models/operations/createtransferinitiationresponse.md), error**

### Errors

| Error Type                      | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.PaymentsErrorResponse | default                         | application/json                |
| sdkerrors.SDKError              | 4XX, 5XX                        | \*/\*                           |

## DeletePool

Delete a pool by its id.

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

    res, err := s.Payments.V1.DeletePool(ctx, operations.DeletePoolRequest{
        PoolID: "XXX",
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
| `request`                                                                        | [operations.DeletePoolRequest](../../pkg/models/operations/deletepoolrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                     | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.DeletePoolResponse](../../pkg/models/operations/deletepoolresponse.md), error**

### Errors

| Error Type                      | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.PaymentsErrorResponse | default                         | application/json                |
| sdkerrors.SDKError              | 4XX, 5XX                        | \*/\*                           |

## DeleteTransferInitiation

Delete a transfer initiation by its id.

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

    res, err := s.Payments.V1.DeleteTransferInitiation(ctx, operations.DeleteTransferInitiationRequest{
        TransferID: "XXX",
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
| `request`                                                                                                    | [operations.DeleteTransferInitiationRequest](../../pkg/models/operations/deletetransferinitiationrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.DeleteTransferInitiationResponse](../../pkg/models/operations/deletetransferinitiationresponse.md), error**

### Errors

| Error Type                      | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.PaymentsErrorResponse | default                         | application/json                |
| sdkerrors.SDKError              | 4XX, 5XX                        | \*/\*                           |

## ForwardBankAccount

Forward a bank account to a connector

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

    res, err := s.Payments.V1.ForwardBankAccount(ctx, operations.ForwardBankAccountRequest{
        ForwardBankAccountRequest: shared.ForwardBankAccountRequest{
            ConnectorID: "<id>",
        },
        BankAccountID: "XXX",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.BankAccountResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.ForwardBankAccountRequest](../../pkg/models/operations/forwardbankaccountrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.ForwardBankAccountResponse](../../pkg/models/operations/forwardbankaccountresponse.md), error**

### Errors

| Error Type                      | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.PaymentsErrorResponse | default                         | application/json                |
| sdkerrors.SDKError              | 4XX, 5XX                        | \*/\*                           |

## GetAccountBalances

Get account balances

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

    res, err := s.Payments.V1.GetAccountBalances(ctx, operations.GetAccountBalancesRequest{
        AccountID: "XXX",
        Cursor: formancesdkgo.String("aHR0cHM6Ly9nLnBhZ2UvTmVrby1SYW1lbj9zaGFyZQ=="),
        PageSize: formancesdkgo.Int64(100),
        Sort: []string{
            "date:asc",
            "status:desc",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.BalancesCursor != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.GetAccountBalancesRequest](../../pkg/models/operations/getaccountbalancesrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.GetAccountBalancesResponse](../../pkg/models/operations/getaccountbalancesresponse.md), error**

### Errors

| Error Type                      | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.PaymentsErrorResponse | default                         | application/json                |
| sdkerrors.SDKError              | 4XX, 5XX                        | \*/\*                           |

## GetBankAccount

Get a bank account created by user on Formance

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

    res, err := s.Payments.V1.GetBankAccount(ctx, operations.GetBankAccountRequest{
        BankAccountID: "XXX",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.BankAccountResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetBankAccountRequest](../../pkg/models/operations/getbankaccountrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.GetBankAccountResponse](../../pkg/models/operations/getbankaccountresponse.md), error**

### Errors

| Error Type                      | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.PaymentsErrorResponse | default                         | application/json                |
| sdkerrors.SDKError              | 4XX, 5XX                        | \*/\*                           |

## ~~GetConnectorTask~~

Get a specific task associated to the connector.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

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

    res, err := s.Payments.V1.GetConnectorTask(ctx, operations.GetConnectorTaskRequest{
        Connector: shared.ConnectorMoneycorp,
        TaskID: "task1",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TaskResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetConnectorTaskRequest](../../pkg/models/operations/getconnectortaskrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.GetConnectorTaskResponse](../../pkg/models/operations/getconnectortaskresponse.md), error**

### Errors

| Error Type                      | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.PaymentsErrorResponse | default                         | application/json                |
| sdkerrors.SDKError              | 4XX, 5XX                        | \*/\*                           |

## GetConnectorTaskV1

Get a specific task associated to the connector.

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

    res, err := s.Payments.V1.GetConnectorTaskV1(ctx, operations.GetConnectorTaskV1Request{
        Connector: shared.ConnectorModulr,
        ConnectorID: "XXX",
        TaskID: "task1",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TaskResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.GetConnectorTaskV1Request](../../pkg/models/operations/getconnectortaskv1request.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.GetConnectorTaskV1Response](../../pkg/models/operations/getconnectortaskv1response.md), error**

### Errors

| Error Type                      | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.PaymentsErrorResponse | default                         | application/json                |
| sdkerrors.SDKError              | 4XX, 5XX                        | \*/\*                           |

## GetPayment

Get a payment

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

    res, err := s.Payments.V1.GetPayment(ctx, operations.GetPaymentRequest{
        PaymentID: "XXX",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PaymentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.GetPaymentRequest](../../pkg/models/operations/getpaymentrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                     | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.GetPaymentResponse](../../pkg/models/operations/getpaymentresponse.md), error**

### Errors

| Error Type                      | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.PaymentsErrorResponse | default                         | application/json                |
| sdkerrors.SDKError              | 4XX, 5XX                        | \*/\*                           |

## GetPool

Get a Pool

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

    res, err := s.Payments.V1.GetPool(ctx, operations.GetPoolRequest{
        PoolID: "XXX",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PoolResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.GetPoolRequest](../../pkg/models/operations/getpoolrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |
| `opts`                                                                     | [][operations.Option](../../pkg/models/operations/option.md)               | :heavy_minus_sign:                                                         | The options for this request.                                              |

### Response

**[*operations.GetPoolResponse](../../pkg/models/operations/getpoolresponse.md), error**

### Errors

| Error Type                      | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.PaymentsErrorResponse | default                         | application/json                |
| sdkerrors.SDKError              | 4XX, 5XX                        | \*/\*                           |

## GetPoolBalances

Get historical pool balances at a particular point in time

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

    res, err := s.Payments.V1.GetPoolBalances(ctx, operations.GetPoolBalancesRequest{
        At: types.MustTimeFromString("2024-11-27T10:59:51.663Z"),
        PoolID: "XXX",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PoolBalancesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetPoolBalancesRequest](../../pkg/models/operations/getpoolbalancesrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.GetPoolBalancesResponse](../../pkg/models/operations/getpoolbalancesresponse.md), error**

### Errors

| Error Type                      | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.PaymentsErrorResponse | default                         | application/json                |
| sdkerrors.SDKError              | 4XX, 5XX                        | \*/\*                           |

## GetPoolBalancesLatest

Get latest pool balances

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

    res, err := s.Payments.V1.GetPoolBalancesLatest(ctx, operations.GetPoolBalancesLatestRequest{
        PoolID: "XXX",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PoolBalancesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.GetPoolBalancesLatestRequest](../../pkg/models/operations/getpoolbalanceslatestrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.GetPoolBalancesLatestResponse](../../pkg/models/operations/getpoolbalanceslatestresponse.md), error**

### Errors

| Error Type                      | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.PaymentsErrorResponse | default                         | application/json                |
| sdkerrors.SDKError              | 4XX, 5XX                        | \*/\*                           |

## GetTransferInitiation

Get a transfer initiation

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

    res, err := s.Payments.V1.GetTransferInitiation(ctx, operations.GetTransferInitiationRequest{
        TransferID: "XXX",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TransferInitiationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.GetTransferInitiationRequest](../../pkg/models/operations/gettransferinitiationrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.GetTransferInitiationResponse](../../pkg/models/operations/gettransferinitiationresponse.md), error**

### Errors

| Error Type                      | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.PaymentsErrorResponse | default                         | application/json                |
| sdkerrors.SDKError              | 4XX, 5XX                        | \*/\*                           |

## InstallConnector

Install a connector by its name and config.

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

    res, err := s.Payments.V1.InstallConnector(ctx, operations.InstallConnectorRequest{
        ConnectorConfig: shared.CreateConnectorConfigCurrencycloud(
            shared.CurrencyCloudConfig{
                APIKey: "XXX",
                LoginID: "XXX",
                Name: "My CurrencyCloud Account",
                PollingPeriod: formancesdkgo.String("60s"),
            },
        ),
        Connector: shared.ConnectorMangopay,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ConnectorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.InstallConnectorRequest](../../pkg/models/operations/installconnectorrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.InstallConnectorResponse](../../pkg/models/operations/installconnectorresponse.md), error**

### Errors

| Error Type                      | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.PaymentsErrorResponse | default                         | application/json                |
| sdkerrors.SDKError              | 4XX, 5XX                        | \*/\*                           |

## ListAllConnectors

List all installed connectors.

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

    res, err := s.Payments.V1.ListAllConnectors(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.ConnectorsResponse != nil {
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

**[*operations.ListAllConnectorsResponse](../../pkg/models/operations/listallconnectorsresponse.md), error**

### Errors

| Error Type                      | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.PaymentsErrorResponse | default                         | application/json                |
| sdkerrors.SDKError              | 4XX, 5XX                        | \*/\*                           |

## ListBankAccounts

List all bank accounts created by user on Formance.

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

    res, err := s.Payments.V1.ListBankAccounts(ctx, operations.ListBankAccountsRequest{
        Cursor: formancesdkgo.String("aHR0cHM6Ly9nLnBhZ2UvTmVrby1SYW1lbj9zaGFyZQ=="),
        PageSize: formancesdkgo.Int64(100),
        Sort: []string{
            "date:asc",
            "status:desc",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.BankAccountsCursor != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListBankAccountsRequest](../../pkg/models/operations/listbankaccountsrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.ListBankAccountsResponse](../../pkg/models/operations/listbankaccountsresponse.md), error**

### Errors

| Error Type                      | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.PaymentsErrorResponse | default                         | application/json                |
| sdkerrors.SDKError              | 4XX, 5XX                        | \*/\*                           |

## ListConfigsAvailableConnectors

List the configs of each available connector.

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

    res, err := s.Payments.V1.ListConfigsAvailableConnectors(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.ConnectorsConfigsResponse != nil {
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

**[*operations.ListConfigsAvailableConnectorsResponse](../../pkg/models/operations/listconfigsavailableconnectorsresponse.md), error**

### Errors

| Error Type                      | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.PaymentsErrorResponse | default                         | application/json                |
| sdkerrors.SDKError              | 4XX, 5XX                        | \*/\*                           |

## ~~ListConnectorTasks~~

List all tasks associated with this connector.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

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

    res, err := s.Payments.V1.ListConnectorTasks(ctx, operations.ListConnectorTasksRequest{
        Connector: shared.ConnectorModulr,
        Cursor: formancesdkgo.String("aHR0cHM6Ly9nLnBhZ2UvTmVrby1SYW1lbj9zaGFyZQ=="),
        PageSize: formancesdkgo.Int64(100),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TasksCursor != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.ListConnectorTasksRequest](../../pkg/models/operations/listconnectortasksrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.ListConnectorTasksResponse](../../pkg/models/operations/listconnectortasksresponse.md), error**

### Errors

| Error Type                      | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.PaymentsErrorResponse | default                         | application/json                |
| sdkerrors.SDKError              | 4XX, 5XX                        | \*/\*                           |

## ListConnectorTasksV1

List all tasks associated with this connector.

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

    res, err := s.Payments.V1.ListConnectorTasksV1(ctx, operations.ListConnectorTasksV1Request{
        Connector: shared.ConnectorWise,
        ConnectorID: "XXX",
        Cursor: formancesdkgo.String("aHR0cHM6Ly9nLnBhZ2UvTmVrby1SYW1lbj9zaGFyZQ=="),
        PageSize: formancesdkgo.Int64(100),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TasksCursor != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.ListConnectorTasksV1Request](../../pkg/models/operations/listconnectortasksv1request.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.ListConnectorTasksV1Response](../../pkg/models/operations/listconnectortasksv1response.md), error**

### Errors

| Error Type                      | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.PaymentsErrorResponse | default                         | application/json                |
| sdkerrors.SDKError              | 4XX, 5XX                        | \*/\*                           |

## ListPayments

List payments

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

    res, err := s.Payments.V1.ListPayments(ctx, operations.ListPaymentsRequest{
        Cursor: formancesdkgo.String("aHR0cHM6Ly9nLnBhZ2UvTmVrby1SYW1lbj9zaGFyZQ=="),
        PageSize: formancesdkgo.Int64(100),
        Sort: []string{
            "date:asc",
            "status:desc",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PaymentsCursor != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.ListPaymentsRequest](../../pkg/models/operations/listpaymentsrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                         | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.ListPaymentsResponse](../../pkg/models/operations/listpaymentsresponse.md), error**

### Errors

| Error Type                      | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.PaymentsErrorResponse | default                         | application/json                |
| sdkerrors.SDKError              | 4XX, 5XX                        | \*/\*                           |

## ListPools

List Pools

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

    res, err := s.Payments.V1.ListPools(ctx, operations.ListPoolsRequest{
        Cursor: formancesdkgo.String("aHR0cHM6Ly9nLnBhZ2UvTmVrby1SYW1lbj9zaGFyZQ=="),
        PageSize: formancesdkgo.Int64(100),
        Sort: []string{
            "date:asc",
            "status:desc",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PoolsCursor != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.ListPoolsRequest](../../pkg/models/operations/listpoolsrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                   | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.ListPoolsResponse](../../pkg/models/operations/listpoolsresponse.md), error**

### Errors

| Error Type                      | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.PaymentsErrorResponse | default                         | application/json                |
| sdkerrors.SDKError              | 4XX, 5XX                        | \*/\*                           |

## ListTransferInitiations

List Transfer Initiations

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

    res, err := s.Payments.V1.ListTransferInitiations(ctx, operations.ListTransferInitiationsRequest{
        Cursor: formancesdkgo.String("aHR0cHM6Ly9nLnBhZ2UvTmVrby1SYW1lbj9zaGFyZQ=="),
        PageSize: formancesdkgo.Int64(100),
        Sort: []string{
            "date:asc",
            "status:desc",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TransferInitiationsCursor != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.ListTransferInitiationsRequest](../../pkg/models/operations/listtransferinitiationsrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.ListTransferInitiationsResponse](../../pkg/models/operations/listtransferinitiationsresponse.md), error**

### Errors

| Error Type                      | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.PaymentsErrorResponse | default                         | application/json                |
| sdkerrors.SDKError              | 4XX, 5XX                        | \*/\*                           |

## PaymentsgetAccount

Get an account

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

    res, err := s.Payments.V1.PaymentsgetAccount(ctx, operations.PaymentsgetAccountRequest{
        AccountID: "XXX",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PaymentsAccountResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.PaymentsgetAccountRequest](../../pkg/models/operations/paymentsgetaccountrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.PaymentsgetAccountResponse](../../pkg/models/operations/paymentsgetaccountresponse.md), error**

### Errors

| Error Type                      | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.PaymentsErrorResponse | default                         | application/json                |
| sdkerrors.SDKError              | 4XX, 5XX                        | \*/\*                           |

## PaymentsgetServerInfo

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

    res, err := s.Payments.V1.PaymentsgetServerInfo(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.PaymentsServerInfo != nil {
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

**[*operations.PaymentsgetServerInfoResponse](../../pkg/models/operations/paymentsgetserverinforesponse.md), error**

### Errors

| Error Type                      | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.PaymentsErrorResponse | default                         | application/json                |
| sdkerrors.SDKError              | 4XX, 5XX                        | \*/\*                           |

## PaymentslistAccounts

List accounts

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

    res, err := s.Payments.V1.PaymentslistAccounts(ctx, operations.PaymentslistAccountsRequest{
        Cursor: formancesdkgo.String("aHR0cHM6Ly9nLnBhZ2UvTmVrby1SYW1lbj9zaGFyZQ=="),
        PageSize: formancesdkgo.Int64(100),
        Sort: []string{
            "date:asc",
            "status:desc",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountsCursor != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.PaymentslistAccountsRequest](../../pkg/models/operations/paymentslistaccountsrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.PaymentslistAccountsResponse](../../pkg/models/operations/paymentslistaccountsresponse.md), error**

### Errors

| Error Type                      | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.PaymentsErrorResponse | default                         | application/json                |
| sdkerrors.SDKError              | 4XX, 5XX                        | \*/\*                           |

## ~~ReadConnectorConfig~~

Read connector config

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

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

    res, err := s.Payments.V1.ReadConnectorConfig(ctx, operations.ReadConnectorConfigRequest{
        Connector: shared.ConnectorModulr,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ConnectorConfigResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.ReadConnectorConfigRequest](../../pkg/models/operations/readconnectorconfigrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.ReadConnectorConfigResponse](../../pkg/models/operations/readconnectorconfigresponse.md), error**

### Errors

| Error Type                      | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.PaymentsErrorResponse | default                         | application/json                |
| sdkerrors.SDKError              | 4XX, 5XX                        | \*/\*                           |

## ReadConnectorConfigV1

Read connector config

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

    res, err := s.Payments.V1.ReadConnectorConfigV1(ctx, operations.ReadConnectorConfigV1Request{
        Connector: shared.ConnectorMangopay,
        ConnectorID: "XXX",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ConnectorConfigResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.ReadConnectorConfigV1Request](../../pkg/models/operations/readconnectorconfigv1request.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.ReadConnectorConfigV1Response](../../pkg/models/operations/readconnectorconfigv1response.md), error**

### Errors

| Error Type                      | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.PaymentsErrorResponse | default                         | application/json                |
| sdkerrors.SDKError              | 4XX, 5XX                        | \*/\*                           |

## RemoveAccountFromPool

Remove an account from a pool by its id.

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

    res, err := s.Payments.V1.RemoveAccountFromPool(ctx, operations.RemoveAccountFromPoolRequest{
        AccountID: "XXX",
        PoolID: "XXX",
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

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.RemoveAccountFromPoolRequest](../../pkg/models/operations/removeaccountfrompoolrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.RemoveAccountFromPoolResponse](../../pkg/models/operations/removeaccountfrompoolresponse.md), error**

### Errors

| Error Type                      | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.PaymentsErrorResponse | default                         | application/json                |
| sdkerrors.SDKError              | 4XX, 5XX                        | \*/\*                           |

## ~~ResetConnector~~

Reset a connector by its name.
It will remove the connector and ALL PAYMENTS generated with it.


> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

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

    res, err := s.Payments.V1.ResetConnector(ctx, operations.ResetConnectorRequest{
        Connector: shared.ConnectorWise,
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
| `request`                                                                                | [operations.ResetConnectorRequest](../../pkg/models/operations/resetconnectorrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.ResetConnectorResponse](../../pkg/models/operations/resetconnectorresponse.md), error**

### Errors

| Error Type                      | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.PaymentsErrorResponse | default                         | application/json                |
| sdkerrors.SDKError              | 4XX, 5XX                        | \*/\*                           |

## ResetConnectorV1

Reset a connector by its name.
It will remove the connector and ALL PAYMENTS generated with it.


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

    res, err := s.Payments.V1.ResetConnectorV1(ctx, operations.ResetConnectorV1Request{
        Connector: shared.ConnectorWise,
        ConnectorID: "XXX",
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
| `request`                                                                                    | [operations.ResetConnectorV1Request](../../pkg/models/operations/resetconnectorv1request.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.ResetConnectorV1Response](../../pkg/models/operations/resetconnectorv1response.md), error**

### Errors

| Error Type                      | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.PaymentsErrorResponse | default                         | application/json                |
| sdkerrors.SDKError              | 4XX, 5XX                        | \*/\*                           |

## RetryTransferInitiation

Retry a failed transfer initiation

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

    res, err := s.Payments.V1.RetryTransferInitiation(ctx, operations.RetryTransferInitiationRequest{
        TransferID: "XXX",
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

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.RetryTransferInitiationRequest](../../pkg/models/operations/retrytransferinitiationrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.RetryTransferInitiationResponse](../../pkg/models/operations/retrytransferinitiationresponse.md), error**

### Errors

| Error Type                      | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.PaymentsErrorResponse | default                         | application/json                |
| sdkerrors.SDKError              | 4XX, 5XX                        | \*/\*                           |

## ReverseTransferInitiation

Reverse transfer initiation

### Example Usage

```go
package main

import(
	"context"
	formancesdkgo "github.com/formancehq/formance-sdk-go/v3"
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/shared"
	"math/big"
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

    res, err := s.Payments.V1.ReverseTransferInitiation(ctx, operations.ReverseTransferInitiationRequest{
        ReverseTransferInitiationRequest: shared.ReverseTransferInitiationRequest{
            Amount: big.NewInt(978875),
            Asset: "USD",
            Description: "whenever phooey a unlike tremendously whoever after when tight",
            Metadata: map[string]string{
                "key": "<value>",
            },
            Reference: "XXX",
        },
        TransferID: "XXX",
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

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.ReverseTransferInitiationRequest](../../pkg/models/operations/reversetransferinitiationrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.ReverseTransferInitiationResponse](../../pkg/models/operations/reversetransferinitiationresponse.md), error**

### Errors

| Error Type                      | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.PaymentsErrorResponse | default                         | application/json                |
| sdkerrors.SDKError              | 4XX, 5XX                        | \*/\*                           |

## ~~UninstallConnector~~

Uninstall a connector by its name.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

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

    res, err := s.Payments.V1.UninstallConnector(ctx, operations.UninstallConnectorRequest{
        Connector: shared.ConnectorGeneric,
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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.UninstallConnectorRequest](../../pkg/models/operations/uninstallconnectorrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.UninstallConnectorResponse](../../pkg/models/operations/uninstallconnectorresponse.md), error**

### Errors

| Error Type                      | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.PaymentsErrorResponse | default                         | application/json                |
| sdkerrors.SDKError              | 4XX, 5XX                        | \*/\*                           |

## UninstallConnectorV1

Uninstall a connector by its name.

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

    res, err := s.Payments.V1.UninstallConnectorV1(ctx, operations.UninstallConnectorV1Request{
        Connector: shared.ConnectorBankingCircle,
        ConnectorID: "XXX",
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
| `request`                                                                                            | [operations.UninstallConnectorV1Request](../../pkg/models/operations/uninstallconnectorv1request.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.UninstallConnectorV1Response](../../pkg/models/operations/uninstallconnectorv1response.md), error**

### Errors

| Error Type                      | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.PaymentsErrorResponse | default                         | application/json                |
| sdkerrors.SDKError              | 4XX, 5XX                        | \*/\*                           |

## UpdateBankAccountMetadata

Update metadata of a bank account

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

    res, err := s.Payments.V1.UpdateBankAccountMetadata(ctx, operations.UpdateBankAccountMetadataRequest{
        UpdateBankAccountMetadataRequest: shared.UpdateBankAccountMetadataRequest{
            Metadata: map[string]string{
                "key": "<value>",
                "key1": "<value>",
                "key2": "<value>",
            },
        },
        BankAccountID: "XXX",
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

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.UpdateBankAccountMetadataRequest](../../pkg/models/operations/updatebankaccountmetadatarequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.UpdateBankAccountMetadataResponse](../../pkg/models/operations/updatebankaccountmetadataresponse.md), error**

### Errors

| Error Type                      | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.PaymentsErrorResponse | default                         | application/json                |
| sdkerrors.SDKError              | 4XX, 5XX                        | \*/\*                           |

## UpdateConnectorConfigV1

Update connector config

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

    res, err := s.Payments.V1.UpdateConnectorConfigV1(ctx, operations.UpdateConnectorConfigV1Request{
        ConnectorConfig: shared.CreateConnectorConfigModulr(
            shared.ModulrConfig{
                APIKey: "XXX",
                APISecret: "XXX",
                Name: "My Modulr Account",
                PollingPeriod: formancesdkgo.String("60s"),
            },
        ),
        Connector: shared.ConnectorMangopay,
        ConnectorID: "XXX",
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

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.UpdateConnectorConfigV1Request](../../pkg/models/operations/updateconnectorconfigv1request.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.UpdateConnectorConfigV1Response](../../pkg/models/operations/updateconnectorconfigv1response.md), error**

### Errors

| Error Type                      | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.PaymentsErrorResponse | default                         | application/json                |
| sdkerrors.SDKError              | 4XX, 5XX                        | \*/\*                           |

## UpdateMetadata

Update metadata

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

    res, err := s.Payments.V1.UpdateMetadata(ctx, operations.UpdateMetadataRequest{
        RequestBody: map[string]string{
            "key": "<value>",
        },
        PaymentID: "XXX",
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
| `request`                                                                                | [operations.UpdateMetadataRequest](../../pkg/models/operations/updatemetadatarequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.UpdateMetadataResponse](../../pkg/models/operations/updatemetadataresponse.md), error**

### Errors

| Error Type                      | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.PaymentsErrorResponse | default                         | application/json                |
| sdkerrors.SDKError              | 4XX, 5XX                        | \*/\*                           |

## UpdateTransferInitiationStatus

Update a transfer initiation status

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

    res, err := s.Payments.V1.UpdateTransferInitiationStatus(ctx, operations.UpdateTransferInitiationStatusRequest{
        UpdateTransferInitiationStatusRequest: shared.UpdateTransferInitiationStatusRequest{
            Status: shared.StatusValidated,
        },
        TransferID: "XXX",
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

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.UpdateTransferInitiationStatusRequest](../../pkg/models/operations/updatetransferinitiationstatusrequest.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |
| `opts`                                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                                             | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.UpdateTransferInitiationStatusResponse](../../pkg/models/operations/updatetransferinitiationstatusresponse.md), error**

### Errors

| Error Type                      | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.PaymentsErrorResponse | default                         | application/json                |
| sdkerrors.SDKError              | 4XX, 5XX                        | \*/\*                           |