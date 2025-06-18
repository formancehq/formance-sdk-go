# V3
(*Payments.V3*)

## Overview

### Available Operations

* [AddAccountToPool](#addaccounttopool) - Add an account to a pool
* [AddBankAccountToPaymentServiceUser](#addbankaccounttopaymentserviceuser) - Add a bank account to a payment service user
* [ApprovePaymentInitiation](#approvepaymentinitiation) - Approve a payment initiation
* [CreateAccount](#createaccount) - Create a formance account object. This object will not be forwarded to the connector. It is only used for internal purposes.

* [CreateBankAccount](#createbankaccount) - Create a formance bank account object. This object will not be forwarded to the connector until you called the forwardBankAccount method.

* [CreatePayment](#createpayment) - Create a formance payment object. This object will not be forwarded to the connector. It is only used for internal purposes.

* [CreatePaymentServiceUser](#createpaymentserviceuser) - Create a formance payment service user object
* [CreatePool](#createpool) - Create a formance pool object
* [DeletePaymentInitiation](#deletepaymentinitiation) - Delete a payment initiation by ID
* [DeletePool](#deletepool) - Delete a pool by ID
* [ForwardBankAccount](#forwardbankaccount) - Forward a Bank Account to a PSP for creation
* [ForwardPaymentServiceUserBankAccount](#forwardpaymentserviceuserbankaccount) - Forward a payment service user's bank account to a connector
* [GetAccount](#getaccount) - Get an account by ID
* [GetAccountBalances](#getaccountbalances) - Get account balances
* [GetBankAccount](#getbankaccount) - Get a Bank Account by ID
* [GetConnectorConfig](#getconnectorconfig) - Get a connector configuration by ID
* [GetConnectorSchedule](#getconnectorschedule) - Get a connector schedule by ID
* [GetPayment](#getpayment) - Get a payment by ID
* [GetPaymentInitiation](#getpaymentinitiation) - Get a payment initiation by ID
* [GetPaymentServiceUser](#getpaymentserviceuser) - Get a payment service user by ID
* [GetPool](#getpool) - Get a pool by ID
* [GetPoolBalances](#getpoolbalances) - Get historical pool balances from a particular point in time
* [GetPoolBalancesLatest](#getpoolbalanceslatest) - Get latest pool balances
* [GetTask](#gettask) - Get a task and its result by ID
* [InitiatePayment](#initiatepayment) - Initiate a payment
* [InstallConnector](#installconnector) - Install a connector
* [ListAccounts](#listaccounts) - List all accounts
* [ListBankAccounts](#listbankaccounts) - List all bank accounts
* [ListConnectorConfigs](#listconnectorconfigs) - List all connector configurations
* [ListConnectorScheduleInstances](#listconnectorscheduleinstances) - List all connector schedule instances
* [ListConnectorSchedules](#listconnectorschedules) - List all connector schedules
* [ListConnectors](#listconnectors) - List all connectors
* [ListPaymentInitiationAdjustments](#listpaymentinitiationadjustments) - List all payment initiation adjustments
* [ListPaymentInitiationRelatedPayments](#listpaymentinitiationrelatedpayments) - List all payments related to a payment initiation
* [ListPaymentInitiations](#listpaymentinitiations) - List all payment initiations
* [ListPaymentServiceUsers](#listpaymentserviceusers) - List all payment service users
* [ListPayments](#listpayments) - List all payments
* [ListPools](#listpools) - List all pools
* [RejectPaymentInitiation](#rejectpaymentinitiation) - Reject a payment initiation
* [RemoveAccountFromPool](#removeaccountfrompool) - Remove an account from a pool
* [ResetConnector](#resetconnector) - Reset a connector. Be aware that this will delete all data and stop all existing tasks like payment initiations and bank account creations.
* [RetryPaymentInitiation](#retrypaymentinitiation) - Retry a payment initiation
* [ReversePaymentInitiation](#reversepaymentinitiation) - Reverse a payment initiation
* [UninstallConnector](#uninstallconnector) - Uninstall a connector
* [UpdateBankAccountMetadata](#updatebankaccountmetadata) - Update a bank account's metadata
* [UpdatePaymentMetadata](#updatepaymentmetadata) - Update a payment's metadata
* [V3UpdateConnectorConfig](#v3updateconnectorconfig) - Update the config of a connector

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

    res, err := s.Payments.V3.AddAccountToPool(ctx, operations.V3AddAccountToPoolRequest{
        AccountID: "<id>",
        PoolID: "<id>",
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
| `request`                                                                                        | [operations.V3AddAccountToPoolRequest](../../pkg/models/operations/v3addaccounttopoolrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.V3AddAccountToPoolResponse](../../pkg/models/operations/v3addaccounttopoolresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.V3ErrorResponse | default                   | application/json          |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |

## AddBankAccountToPaymentServiceUser

Add a bank account to a payment service user

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

    res, err := s.Payments.V3.AddBankAccountToPaymentServiceUser(ctx, operations.V3AddBankAccountToPaymentServiceUserRequest{
        BankAccountID: "<id>",
        PaymentServiceUserID: "<id>",
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

| Parameter                                                                                                                            | Type                                                                                                                                 | Required                                                                                                                             | Description                                                                                                                          |
| ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                | :heavy_check_mark:                                                                                                                   | The context to use for the request.                                                                                                  |
| `request`                                                                                                                            | [operations.V3AddBankAccountToPaymentServiceUserRequest](../../pkg/models/operations/v3addbankaccounttopaymentserviceuserrequest.md) | :heavy_check_mark:                                                                                                                   | The request object to use for the request.                                                                                           |
| `opts`                                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                                         | :heavy_minus_sign:                                                                                                                   | The options for this request.                                                                                                        |

### Response

**[*operations.V3AddBankAccountToPaymentServiceUserResponse](../../pkg/models/operations/v3addbankaccounttopaymentserviceuserresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.V3ErrorResponse | default                   | application/json          |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |

## ApprovePaymentInitiation

Approve a payment initiation

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

    res, err := s.Payments.V3.ApprovePaymentInitiation(ctx, operations.V3ApprovePaymentInitiationRequest{
        PaymentInitiationID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.V3ApprovePaymentInitiationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.V3ApprovePaymentInitiationRequest](../../pkg/models/operations/v3approvepaymentinitiationrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.V3ApprovePaymentInitiationResponse](../../pkg/models/operations/v3approvepaymentinitiationresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.V3ErrorResponse | default                   | application/json          |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |

## CreateAccount

Create a formance account object. This object will not be forwarded to the connector. It is only used for internal purposes.


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

    res, err := s.Payments.V3.CreateAccount(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.V3CreateAccountResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [shared.V3CreateAccountRequest](../../pkg/models/shared/v3createaccountrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                       | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.V3CreateAccountResponse](../../pkg/models/operations/v3createaccountresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.V3ErrorResponse | default                   | application/json          |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |

## CreateBankAccount

Create a formance bank account object. This object will not be forwarded to the connector until you called the forwardBankAccount method.


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

    res, err := s.Payments.V3.CreateBankAccount(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.V3CreateBankAccountResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [shared.V3CreateBankAccountRequest](../../pkg/models/shared/v3createbankaccountrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.V3CreateBankAccountResponse](../../pkg/models/operations/v3createbankaccountresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.V3ErrorResponse | default                   | application/json          |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |

## CreatePayment

Create a formance payment object. This object will not be forwarded to the connector. It is only used for internal purposes.


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

    res, err := s.Payments.V3.CreatePayment(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.V3CreatePaymentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [shared.V3CreatePaymentRequest](../../pkg/models/shared/v3createpaymentrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                       | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.V3CreatePaymentResponse](../../pkg/models/operations/v3createpaymentresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.V3ErrorResponse | default                   | application/json          |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |

## CreatePaymentServiceUser

Create a formance payment service user object

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

    res, err := s.Payments.V3.CreatePaymentServiceUser(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.V3CreatePaymentServiceUserResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [shared.V3CreatePaymentServiceUserRequest](../../pkg/models/shared/v3createpaymentserviceuserrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.V3CreatePaymentServiceUserResponse](../../pkg/models/operations/v3createpaymentserviceuserresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.V3ErrorResponse | default                   | application/json          |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |

## CreatePool

Create a formance pool object

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

    res, err := s.Payments.V3.CreatePool(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.V3CreatePoolResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [shared.V3CreatePoolRequest](../../pkg/models/shared/v3createpoolrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `opts`                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                 | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.V3CreatePoolResponse](../../pkg/models/operations/v3createpoolresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.V3ErrorResponse | default                   | application/json          |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |

## DeletePaymentInitiation

Delete a payment initiation by ID

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

    res, err := s.Payments.V3.DeletePaymentInitiation(ctx, operations.V3DeletePaymentInitiationRequest{
        PaymentInitiationID: "<id>",
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
| `request`                                                                                                      | [operations.V3DeletePaymentInitiationRequest](../../pkg/models/operations/v3deletepaymentinitiationrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.V3DeletePaymentInitiationResponse](../../pkg/models/operations/v3deletepaymentinitiationresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.V3ErrorResponse | default                   | application/json          |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |

## DeletePool

Delete a pool by ID

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

    res, err := s.Payments.V3.DeletePool(ctx, operations.V3DeletePoolRequest{
        PoolID: "<id>",
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
| `request`                                                                            | [operations.V3DeletePoolRequest](../../pkg/models/operations/v3deletepoolrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                         | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.V3DeletePoolResponse](../../pkg/models/operations/v3deletepoolresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.V3ErrorResponse | default                   | application/json          |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |

## ForwardBankAccount

Forward a Bank Account to a PSP for creation

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

    res, err := s.Payments.V3.ForwardBankAccount(ctx, operations.V3ForwardBankAccountRequest{
        BankAccountID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.V3ForwardBankAccountResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.V3ForwardBankAccountRequest](../../pkg/models/operations/v3forwardbankaccountrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.V3ForwardBankAccountResponse](../../pkg/models/operations/v3forwardbankaccountresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.V3ErrorResponse | default                   | application/json          |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |

## ForwardPaymentServiceUserBankAccount

Forward a payment service user's bank account to a connector

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

    res, err := s.Payments.V3.ForwardPaymentServiceUserBankAccount(ctx, operations.V3ForwardPaymentServiceUserBankAccountRequest{
        BankAccountID: "<id>",
        PaymentServiceUserID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.V3ForwardPaymentServiceUserBankAccountResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                | Type                                                                                                                                     | Required                                                                                                                                 | Description                                                                                                                              |
| ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                    | :heavy_check_mark:                                                                                                                       | The context to use for the request.                                                                                                      |
| `request`                                                                                                                                | [operations.V3ForwardPaymentServiceUserBankAccountRequest](../../pkg/models/operations/v3forwardpaymentserviceuserbankaccountrequest.md) | :heavy_check_mark:                                                                                                                       | The request object to use for the request.                                                                                               |
| `opts`                                                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                                                             | :heavy_minus_sign:                                                                                                                       | The options for this request.                                                                                                            |

### Response

**[*operations.V3ForwardPaymentServiceUserBankAccountResponse](../../pkg/models/operations/v3forwardpaymentserviceuserbankaccountresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.V3ErrorResponse | default                   | application/json          |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |

## GetAccount

Get an account by ID

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

    res, err := s.Payments.V3.GetAccount(ctx, operations.V3GetAccountRequest{
        AccountID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.V3GetAccountResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.V3GetAccountRequest](../../pkg/models/operations/v3getaccountrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                         | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.V3GetAccountResponse](../../pkg/models/operations/v3getaccountresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.V3ErrorResponse | default                   | application/json          |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |

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

    res, err := s.Payments.V3.GetAccountBalances(ctx, operations.V3GetAccountBalancesRequest{
        AccountID: "<id>",
        Cursor: formancesdkgo.String("aHR0cHM6Ly9nLnBhZ2UvTmVrby1SYW1lbj9zaGFyZQ=="),
        PageSize: formancesdkgo.Int64(100),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.V3BalancesCursorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.V3GetAccountBalancesRequest](../../pkg/models/operations/v3getaccountbalancesrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.V3GetAccountBalancesResponse](../../pkg/models/operations/v3getaccountbalancesresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.V3ErrorResponse | default                   | application/json          |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |

## GetBankAccount

Get a Bank Account by ID

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

    res, err := s.Payments.V3.GetBankAccount(ctx, operations.V3GetBankAccountRequest{
        BankAccountID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.V3GetBankAccountResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.V3GetBankAccountRequest](../../pkg/models/operations/v3getbankaccountrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.V3GetBankAccountResponse](../../pkg/models/operations/v3getbankaccountresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.V3ErrorResponse | default                   | application/json          |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |

## GetConnectorConfig

Get a connector configuration by ID

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

    res, err := s.Payments.V3.GetConnectorConfig(ctx, operations.V3GetConnectorConfigRequest{
        ConnectorID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.V3GetConnectorConfigResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.V3GetConnectorConfigRequest](../../pkg/models/operations/v3getconnectorconfigrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.V3GetConnectorConfigResponse](../../pkg/models/operations/v3getconnectorconfigresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.V3ErrorResponse | default                   | application/json          |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |

## GetConnectorSchedule

Get a connector schedule by ID

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

    res, err := s.Payments.V3.GetConnectorSchedule(ctx, operations.V3GetConnectorScheduleRequest{
        ConnectorID: "<id>",
        ScheduleID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.V3ConnectorScheduleResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.V3GetConnectorScheduleRequest](../../pkg/models/operations/v3getconnectorschedulerequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.V3GetConnectorScheduleResponse](../../pkg/models/operations/v3getconnectorscheduleresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.V3ErrorResponse | default                   | application/json          |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |

## GetPayment

Get a payment by ID

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

    res, err := s.Payments.V3.GetPayment(ctx, operations.V3GetPaymentRequest{
        PaymentID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.V3GetPaymentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.V3GetPaymentRequest](../../pkg/models/operations/v3getpaymentrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                         | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.V3GetPaymentResponse](../../pkg/models/operations/v3getpaymentresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.V3ErrorResponse | default                   | application/json          |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |

## GetPaymentInitiation

Get a payment initiation by ID

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

    res, err := s.Payments.V3.GetPaymentInitiation(ctx, operations.V3GetPaymentInitiationRequest{
        PaymentInitiationID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.V3GetPaymentInitiationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.V3GetPaymentInitiationRequest](../../pkg/models/operations/v3getpaymentinitiationrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.V3GetPaymentInitiationResponse](../../pkg/models/operations/v3getpaymentinitiationresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.V3ErrorResponse | default                   | application/json          |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |

## GetPaymentServiceUser

Get a payment service user by ID

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

    res, err := s.Payments.V3.GetPaymentServiceUser(ctx, operations.V3GetPaymentServiceUserRequest{
        PaymentServiceUserID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.V3GetPaymentServiceUserResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.V3GetPaymentServiceUserRequest](../../pkg/models/operations/v3getpaymentserviceuserrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.V3GetPaymentServiceUserResponse](../../pkg/models/operations/v3getpaymentserviceuserresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.V3ErrorResponse | default                   | application/json          |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |

## GetPool

Get a pool by ID

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

    res, err := s.Payments.V3.GetPool(ctx, operations.V3GetPoolRequest{
        PoolID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.V3GetPoolResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.V3GetPoolRequest](../../pkg/models/operations/v3getpoolrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                   | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.V3GetPoolResponse](../../pkg/models/operations/v3getpoolresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.V3ErrorResponse | default                   | application/json          |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |

## GetPoolBalances

Get historical pool balances from a particular point in time

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

    res, err := s.Payments.V3.GetPoolBalances(ctx, operations.V3GetPoolBalancesRequest{
        PoolID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.V3PoolBalancesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.V3GetPoolBalancesRequest](../../pkg/models/operations/v3getpoolbalancesrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.V3GetPoolBalancesResponse](../../pkg/models/operations/v3getpoolbalancesresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.V3ErrorResponse | default                   | application/json          |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |

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

    res, err := s.Payments.V3.GetPoolBalancesLatest(ctx, operations.V3GetPoolBalancesLatestRequest{
        PoolID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.V3PoolBalancesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.V3GetPoolBalancesLatestRequest](../../pkg/models/operations/v3getpoolbalanceslatestrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.V3GetPoolBalancesLatestResponse](../../pkg/models/operations/v3getpoolbalanceslatestresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.V3ErrorResponse | default                   | application/json          |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |

## GetTask

Get a task and its result by ID

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

    res, err := s.Payments.V3.GetTask(ctx, operations.V3GetTaskRequest{
        TaskID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.V3GetTaskResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.V3GetTaskRequest](../../pkg/models/operations/v3gettaskrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                   | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.V3GetTaskResponse](../../pkg/models/operations/v3gettaskresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.V3ErrorResponse | default                   | application/json          |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |

## InitiatePayment

Initiate a payment

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

    res, err := s.Payments.V3.InitiatePayment(ctx, operations.V3InitiatePaymentRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.V3InitiatePaymentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.V3InitiatePaymentRequest](../../pkg/models/operations/v3initiatepaymentrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.V3InitiatePaymentResponse](../../pkg/models/operations/v3initiatepaymentresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.V3ErrorResponse | default                   | application/json          |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |

## InstallConnector

Install a connector

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

    res, err := s.Payments.V3.InstallConnector(ctx, operations.V3InstallConnectorRequest{
        Connector: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.V3InstallConnectorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.V3InstallConnectorRequest](../../pkg/models/operations/v3installconnectorrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.V3InstallConnectorResponse](../../pkg/models/operations/v3installconnectorresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.V3ErrorResponse | default                   | application/json          |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |

## ListAccounts

List all accounts

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

    res, err := s.Payments.V3.ListAccounts(ctx, operations.V3ListAccountsRequest{
        Cursor: formancesdkgo.String("aHR0cHM6Ly9nLnBhZ2UvTmVrby1SYW1lbj9zaGFyZQ=="),
        PageSize: formancesdkgo.Int64(100),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.V3AccountsCursorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.V3ListAccountsRequest](../../pkg/models/operations/v3listaccountsrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.V3ListAccountsResponse](../../pkg/models/operations/v3listaccountsresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.V3ErrorResponse | default                   | application/json          |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |

## ListBankAccounts

List all bank accounts

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

    res, err := s.Payments.V3.ListBankAccounts(ctx, operations.V3ListBankAccountsRequest{
        Cursor: formancesdkgo.String("aHR0cHM6Ly9nLnBhZ2UvTmVrby1SYW1lbj9zaGFyZQ=="),
        PageSize: formancesdkgo.Int64(100),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.V3BankAccountsCursorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.V3ListBankAccountsRequest](../../pkg/models/operations/v3listbankaccountsrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.V3ListBankAccountsResponse](../../pkg/models/operations/v3listbankaccountsresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.V3ErrorResponse | default                   | application/json          |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |

## ListConnectorConfigs

List all connector configurations

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

    res, err := s.Payments.V3.ListConnectorConfigs(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.V3ConnectorConfigsResponse != nil {
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

**[*operations.V3ListConnectorConfigsResponse](../../pkg/models/operations/v3listconnectorconfigsresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.V3ErrorResponse | default                   | application/json          |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |

## ListConnectorScheduleInstances

List all connector schedule instances

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

    res, err := s.Payments.V3.ListConnectorScheduleInstances(ctx, operations.V3ListConnectorScheduleInstancesRequest{
        ConnectorID: "<id>",
        Cursor: formancesdkgo.String("aHR0cHM6Ly9nLnBhZ2UvTmVrby1SYW1lbj9zaGFyZQ=="),
        PageSize: formancesdkgo.Int64(100),
        ScheduleID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.V3ConnectorScheduleInstancesCursorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `request`                                                                                                                    | [operations.V3ListConnectorScheduleInstancesRequest](../../pkg/models/operations/v3listconnectorscheduleinstancesrequest.md) | :heavy_check_mark:                                                                                                           | The request object to use for the request.                                                                                   |
| `opts`                                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                           | The options for this request.                                                                                                |

### Response

**[*operations.V3ListConnectorScheduleInstancesResponse](../../pkg/models/operations/v3listconnectorscheduleinstancesresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.V3ErrorResponse | default                   | application/json          |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |

## ListConnectorSchedules

List all connector schedules

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

    res, err := s.Payments.V3.ListConnectorSchedules(ctx, operations.V3ListConnectorSchedulesRequest{
        ConnectorID: "<id>",
        Cursor: formancesdkgo.String("aHR0cHM6Ly9nLnBhZ2UvTmVrby1SYW1lbj9zaGFyZQ=="),
        PageSize: formancesdkgo.Int64(100),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.V3ConnectorSchedulesCursorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.V3ListConnectorSchedulesRequest](../../pkg/models/operations/v3listconnectorschedulesrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.V3ListConnectorSchedulesResponse](../../pkg/models/operations/v3listconnectorschedulesresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.V3ErrorResponse | default                   | application/json          |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |

## ListConnectors

List all connectors

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

    res, err := s.Payments.V3.ListConnectors(ctx, operations.V3ListConnectorsRequest{
        Cursor: formancesdkgo.String("aHR0cHM6Ly9nLnBhZ2UvTmVrby1SYW1lbj9zaGFyZQ=="),
        PageSize: formancesdkgo.Int64(100),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.V3ConnectorsCursorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.V3ListConnectorsRequest](../../pkg/models/operations/v3listconnectorsrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.V3ListConnectorsResponse](../../pkg/models/operations/v3listconnectorsresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.V3ErrorResponse | default                   | application/json          |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |

## ListPaymentInitiationAdjustments

List all payment initiation adjustments

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

    res, err := s.Payments.V3.ListPaymentInitiationAdjustments(ctx, operations.V3ListPaymentInitiationAdjustmentsRequest{
        Cursor: formancesdkgo.String("aHR0cHM6Ly9nLnBhZ2UvTmVrby1SYW1lbj9zaGFyZQ=="),
        PageSize: formancesdkgo.Int64(100),
        PaymentInitiationID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.V3PaymentInitiationAdjustmentsCursorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |
| `request`                                                                                                                        | [operations.V3ListPaymentInitiationAdjustmentsRequest](../../pkg/models/operations/v3listpaymentinitiationadjustmentsrequest.md) | :heavy_check_mark:                                                                                                               | The request object to use for the request.                                                                                       |
| `opts`                                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                                     | :heavy_minus_sign:                                                                                                               | The options for this request.                                                                                                    |

### Response

**[*operations.V3ListPaymentInitiationAdjustmentsResponse](../../pkg/models/operations/v3listpaymentinitiationadjustmentsresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.V3ErrorResponse | default                   | application/json          |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |

## ListPaymentInitiationRelatedPayments

List all payments related to a payment initiation

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

    res, err := s.Payments.V3.ListPaymentInitiationRelatedPayments(ctx, operations.V3ListPaymentInitiationRelatedPaymentsRequest{
        Cursor: formancesdkgo.String("aHR0cHM6Ly9nLnBhZ2UvTmVrby1SYW1lbj9zaGFyZQ=="),
        PageSize: formancesdkgo.Int64(100),
        PaymentInitiationID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.V3PaymentInitiationRelatedPaymentsCursorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                | Type                                                                                                                                     | Required                                                                                                                                 | Description                                                                                                                              |
| ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                    | :heavy_check_mark:                                                                                                                       | The context to use for the request.                                                                                                      |
| `request`                                                                                                                                | [operations.V3ListPaymentInitiationRelatedPaymentsRequest](../../pkg/models/operations/v3listpaymentinitiationrelatedpaymentsrequest.md) | :heavy_check_mark:                                                                                                                       | The request object to use for the request.                                                                                               |
| `opts`                                                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                                                             | :heavy_minus_sign:                                                                                                                       | The options for this request.                                                                                                            |

### Response

**[*operations.V3ListPaymentInitiationRelatedPaymentsResponse](../../pkg/models/operations/v3listpaymentinitiationrelatedpaymentsresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.V3ErrorResponse | default                   | application/json          |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |

## ListPaymentInitiations

List all payment initiations

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

    res, err := s.Payments.V3.ListPaymentInitiations(ctx, operations.V3ListPaymentInitiationsRequest{
        Cursor: formancesdkgo.String("aHR0cHM6Ly9nLnBhZ2UvTmVrby1SYW1lbj9zaGFyZQ=="),
        PageSize: formancesdkgo.Int64(100),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.V3PaymentInitiationsCursorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.V3ListPaymentInitiationsRequest](../../pkg/models/operations/v3listpaymentinitiationsrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.V3ListPaymentInitiationsResponse](../../pkg/models/operations/v3listpaymentinitiationsresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.V3ErrorResponse | default                   | application/json          |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |

## ListPaymentServiceUsers

List all payment service users

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

    res, err := s.Payments.V3.ListPaymentServiceUsers(ctx, operations.V3ListPaymentServiceUsersRequest{
        Cursor: formancesdkgo.String("aHR0cHM6Ly9nLnBhZ2UvTmVrby1SYW1lbj9zaGFyZQ=="),
        PageSize: formancesdkgo.Int64(100),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.V3PaymentServiceUsersCursorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.V3ListPaymentServiceUsersRequest](../../pkg/models/operations/v3listpaymentserviceusersrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.V3ListPaymentServiceUsersResponse](../../pkg/models/operations/v3listpaymentserviceusersresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.V3ErrorResponse | default                   | application/json          |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |

## ListPayments

List all payments

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

    res, err := s.Payments.V3.ListPayments(ctx, operations.V3ListPaymentsRequest{
        Cursor: formancesdkgo.String("aHR0cHM6Ly9nLnBhZ2UvTmVrby1SYW1lbj9zaGFyZQ=="),
        PageSize: formancesdkgo.Int64(100),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.V3PaymentsCursorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.V3ListPaymentsRequest](../../pkg/models/operations/v3listpaymentsrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.V3ListPaymentsResponse](../../pkg/models/operations/v3listpaymentsresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.V3ErrorResponse | default                   | application/json          |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |

## ListPools

List all pools

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

    res, err := s.Payments.V3.ListPools(ctx, operations.V3ListPoolsRequest{
        Cursor: formancesdkgo.String("aHR0cHM6Ly9nLnBhZ2UvTmVrby1SYW1lbj9zaGFyZQ=="),
        PageSize: formancesdkgo.Int64(100),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.V3PoolsCursorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.V3ListPoolsRequest](../../pkg/models/operations/v3listpoolsrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                       | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.V3ListPoolsResponse](../../pkg/models/operations/v3listpoolsresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.V3ErrorResponse | default                   | application/json          |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |

## RejectPaymentInitiation

Reject a payment initiation

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

    res, err := s.Payments.V3.RejectPaymentInitiation(ctx, operations.V3RejectPaymentInitiationRequest{
        PaymentInitiationID: "<id>",
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
| `request`                                                                                                      | [operations.V3RejectPaymentInitiationRequest](../../pkg/models/operations/v3rejectpaymentinitiationrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.V3RejectPaymentInitiationResponse](../../pkg/models/operations/v3rejectpaymentinitiationresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.V3ErrorResponse | default                   | application/json          |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |

## RemoveAccountFromPool

Remove an account from a pool

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

    res, err := s.Payments.V3.RemoveAccountFromPool(ctx, operations.V3RemoveAccountFromPoolRequest{
        AccountID: "<id>",
        PoolID: "<id>",
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
| `request`                                                                                                  | [operations.V3RemoveAccountFromPoolRequest](../../pkg/models/operations/v3removeaccountfrompoolrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.V3RemoveAccountFromPoolResponse](../../pkg/models/operations/v3removeaccountfrompoolresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.V3ErrorResponse | default                   | application/json          |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |

## ResetConnector

Reset a connector. Be aware that this will delete all data and stop all existing tasks like payment initiations and bank account creations.

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

    res, err := s.Payments.V3.ResetConnector(ctx, operations.V3ResetConnectorRequest{
        ConnectorID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.V3ResetConnectorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.V3ResetConnectorRequest](../../pkg/models/operations/v3resetconnectorrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.V3ResetConnectorResponse](../../pkg/models/operations/v3resetconnectorresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.V3ErrorResponse | default                   | application/json          |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |

## RetryPaymentInitiation

Retry a payment initiation

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

    res, err := s.Payments.V3.RetryPaymentInitiation(ctx, operations.V3RetryPaymentInitiationRequest{
        PaymentInitiationID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.V3RetryPaymentInitiationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.V3RetryPaymentInitiationRequest](../../pkg/models/operations/v3retrypaymentinitiationrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.V3RetryPaymentInitiationResponse](../../pkg/models/operations/v3retrypaymentinitiationresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.V3ErrorResponse | default                   | application/json          |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |

## ReversePaymentInitiation

Reverse a payment initiation

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

    res, err := s.Payments.V3.ReversePaymentInitiation(ctx, operations.V3ReversePaymentInitiationRequest{
        PaymentInitiationID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.V3ReversePaymentInitiationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.V3ReversePaymentInitiationRequest](../../pkg/models/operations/v3reversepaymentinitiationrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.V3ReversePaymentInitiationResponse](../../pkg/models/operations/v3reversepaymentinitiationresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.V3ErrorResponse | default                   | application/json          |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |

## UninstallConnector

Uninstall a connector

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

    res, err := s.Payments.V3.UninstallConnector(ctx, operations.V3UninstallConnectorRequest{
        ConnectorID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.V3UninstallConnectorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.V3UninstallConnectorRequest](../../pkg/models/operations/v3uninstallconnectorrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.V3UninstallConnectorResponse](../../pkg/models/operations/v3uninstallconnectorresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.V3ErrorResponse | default                   | application/json          |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |

## UpdateBankAccountMetadata

Update a bank account's metadata

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

    res, err := s.Payments.V3.UpdateBankAccountMetadata(ctx, operations.V3UpdateBankAccountMetadataRequest{
        BankAccountID: "<id>",
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

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.V3UpdateBankAccountMetadataRequest](../../pkg/models/operations/v3updatebankaccountmetadatarequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.V3UpdateBankAccountMetadataResponse](../../pkg/models/operations/v3updatebankaccountmetadataresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.V3ErrorResponse | default                   | application/json          |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |

## UpdatePaymentMetadata

Update a payment's metadata

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

    res, err := s.Payments.V3.UpdatePaymentMetadata(ctx, operations.V3UpdatePaymentMetadataRequest{
        PaymentID: "<id>",
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
| `request`                                                                                                  | [operations.V3UpdatePaymentMetadataRequest](../../pkg/models/operations/v3updatepaymentmetadatarequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.V3UpdatePaymentMetadataResponse](../../pkg/models/operations/v3updatepaymentmetadataresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.V3ErrorResponse | default                   | application/json          |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |

## V3UpdateConnectorConfig

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

    res, err := s.Payments.V3.V3UpdateConnectorConfig(ctx, operations.V3UpdateConnectorConfigRequest{
        ConnectorID: "<id>",
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
| `request`                                                                                                  | [operations.V3UpdateConnectorConfigRequest](../../pkg/models/operations/v3updateconnectorconfigrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.V3UpdateConnectorConfigResponse](../../pkg/models/operations/v3updateconnectorconfigresponse.md), error**

### Errors

| Error Type                      | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.PaymentsErrorResponse | default                         | application/json                |
| sdkerrors.SDKError              | 4XX, 5XX                        | \*/\*                           |