# V1
(*Auth.V1*)

## Overview

### Available Operations

* [CreateClient](#createclient) - Create client
* [CreateSecret](#createsecret) - Add a secret to a client
* [DeleteClient](#deleteclient) - Delete client
* [DeleteSecret](#deletesecret) - Delete a secret from a client
* [GetOIDCWellKnowns](#getoidcwellknowns) - Retrieve OpenID connect well-knowns.
* [GetServerInfo](#getserverinfo) - Get server info
* [ListClients](#listclients) - List clients
* [ListUsers](#listusers) - List users
* [ReadClient](#readclient) - Read client
* [ReadUser](#readuser) - Read user
* [UpdateClient](#updateclient) - Update client

## CreateClient

Create client

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
    res, err := s.Auth.V1.CreateClient(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateClientResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [shared.CreateClientRequest](../../pkg/models/shared/createclientrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `opts`                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                 | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.CreateClientResponse](../../pkg/models/operations/createclientresponse.md), error**

### Errors

| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |


## CreateSecret

Add a secret to a client

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
    res, err := s.Auth.V1.CreateSecret(ctx, operations.CreateSecretRequest{
        ClientID: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateSecretResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.CreateSecretRequest](../../pkg/models/operations/createsecretrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                         | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.CreateSecretResponse](../../pkg/models/operations/createsecretresponse.md), error**

### Errors

| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |


## DeleteClient

Delete client

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
    res, err := s.Auth.V1.DeleteClient(ctx, operations.DeleteClientRequest{
        ClientID: "<value>",
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
| `request`                                                                            | [operations.DeleteClientRequest](../../pkg/models/operations/deleteclientrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                         | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.DeleteClientResponse](../../pkg/models/operations/deleteclientresponse.md), error**

### Errors

| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |


## DeleteSecret

Delete a secret from a client

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
    res, err := s.Auth.V1.DeleteSecret(ctx, operations.DeleteSecretRequest{
        ClientID: "<value>",
        SecretID: "<value>",
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
| `request`                                                                            | [operations.DeleteSecretRequest](../../pkg/models/operations/deletesecretrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                         | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.DeleteSecretResponse](../../pkg/models/operations/deletesecretresponse.md), error**

### Errors

| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |


## GetOIDCWellKnowns

Retrieve OpenID connect well-knowns.

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
    res, err := s.Auth.V1.GetOIDCWellKnowns(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
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

**[*operations.GetOIDCWellKnownsResponse](../../pkg/models/operations/getoidcwellknownsresponse.md), error**

### Errors

| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
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
    res, err := s.Auth.V1.GetServerInfo(ctx)
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

**[*operations.GetServerInfoResponse](../../pkg/models/operations/getserverinforesponse.md), error**

### Errors

| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |


## ListClients

List clients

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
    res, err := s.Auth.V1.ListClients(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.ListClientsResponse != nil {
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

**[*operations.ListClientsResponse](../../pkg/models/operations/listclientsresponse.md), error**

### Errors

| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |


## ListUsers

List users

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
    res, err := s.Auth.V1.ListUsers(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.ListUsersResponse != nil {
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

**[*operations.ListUsersResponse](../../pkg/models/operations/listusersresponse.md), error**

### Errors

| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |


## ReadClient

Read client

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
    res, err := s.Auth.V1.ReadClient(ctx, operations.ReadClientRequest{
        ClientID: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ReadClientResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.ReadClientRequest](../../pkg/models/operations/readclientrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                     | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.ReadClientResponse](../../pkg/models/operations/readclientresponse.md), error**

### Errors

| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |


## ReadUser

Read user

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
    res, err := s.Auth.V1.ReadUser(ctx, operations.ReadUserRequest{
        UserID: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ReadUserResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.ReadUserRequest](../../pkg/models/operations/readuserrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `opts`                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                 | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.ReadUserResponse](../../pkg/models/operations/readuserresponse.md), error**

### Errors

| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |


## UpdateClient

Update client

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
    res, err := s.Auth.V1.UpdateClient(ctx, operations.UpdateClientRequest{
        ClientID: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdateClientResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.UpdateClientRequest](../../pkg/models/operations/updateclientrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                         | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.UpdateClientResponse](../../pkg/models/operations/updateclientresponse.md), error**

### Errors

| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
