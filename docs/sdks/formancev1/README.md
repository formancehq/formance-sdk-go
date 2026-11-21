# Connectivity.V1

## Overview

### Available Operations

* [ConnectivityInfo](#connectivityinfo) - Build information
* [CreateConnection](#createconnection) - Create a Connection
* [DeleteConnection](#deleteconnection) - Delete a Connection
* [GetConnection](#getconnection) - Get a Connection
* [GetConnector](#getconnector) - Get a Connector (lean connector metadata)
* [GetConnectorFacets](#getconnectorfacets) - Facet-value distribution over the published catalogue
* [GetConnectorVersion](#getconnectorversion) - Get one Connector version (including its configSchema and additionalMetadata)
* [GetQueryCapabilities](#getquerycapabilities) - The filterable keys, operators and value enums per resource
* [Healthcheck](#healthcheck) - Health probe
* [ListConnections](#listconnections) - List connector instances for this tenant
* [ListConnectorVersions](#listconnectorversions) - List the installable versions of a Connector
* [ListConnectors](#listconnectors) - List published Connectors with their catalogue identity
* [PatchConnection](#patchconnection) - Merge-patch a connector instance configuration
* [ReplaceConnection](#replaceconnection) - Replace a connector instance configuration

## ConnectivityInfo

Version probe used by the stack gateway to report this module in the
`/versions` aggregation. Response shape follows the platform-wide
ServiceInfo contract.


### Example Usage

<!-- UsageSnippet language="go" operationID="connectivityInfo" method="get" path="/api/connectivity/_info" -->
```go
package main

import(
	"context"
	"github.com/formancehq/formance-sdk-go/v5"
	"log"
)

func main() {
    ctx := context.Background()

    s := v5.New()

    res, err := s.Connectivity.V1.ConnectivityInfo(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.Info != nil {
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

**[*operations.ConnectivityInfoResponse](../../pkg/models/operations/connectivityinforesponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| connectivity.Error | 4XX                | application/json   |
| connectivity.Error | 5XX                | application/json   |

## CreateConnection

Creates a new `Connection` in the caller's environment.
`configuration.connector`, `configuration.ledger` and exactly one of `configuration.version` or
`configuration.channel` are required. `env` and `files` are string key/value maps;
properties marked `format: password` by the selected connector schema
are stored confidentially and omitted from read responses.
Storage metadata and configuration references are rejected on input.
A field outside the tenant contract,
including unsupported instance fields such as `image`, `additionalEnv`,
`additionalFiles` or `replicas`, is rejected with
`400 unsupported_field`.


### Example Usage: channel

<!-- UsageSnippet language="go" operationID="createConnection" method="post" path="/api/connectivity/connections" example="channel" -->
```go
package main

import(
	"context"
	"github.com/formancehq/formance-sdk-go/v5"
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/connectivity"
	"log"
)

func main() {
    ctx := context.Background()

    s := v5.New()

    res, err := s.Connectivity.V1.CreateConnection(ctx, connectivity.ConnectionCreate{
        Configuration: connectivity.Configuration{
            Channel: connectivity.ConnectionCreateChannelStable.ToPointer(),
            Connector: "stripe",
            Env: map[string]string{
                "STRIPE_API_KEY": "<your-api-key>",
            },
            Ledger: "main",
            PollInterval: v5.Pointer("30s"),
        },
        Name: "stripe-main",
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Connection != nil {
        // handle response
    }
}
```
### Example Usage: version

<!-- UsageSnippet language="go" operationID="createConnection" method="post" path="/api/connectivity/connections" example="version" -->
```go
package main

import(
	"context"
	"github.com/formancehq/formance-sdk-go/v5"
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/connectivity"
	"log"
)

func main() {
    ctx := context.Background()

    s := v5.New()

    res, err := s.Connectivity.V1.CreateConnection(ctx, connectivity.ConnectionCreate{
        Configuration: connectivity.Configuration{
            Connector: "bitcoin",
            Env: map[string]string{
                "RPC_URL": "http://bitcoind:8332",
            },
            Files: map[string]string{
                "/etc/bitcoin/ca.pem": "<PEM certificate content>",
            },
            InitialCursor: v5.Pointer("AAAAAAAAA+g="),
            Ledger: "btc",
            Version: v5.Pointer("v1.4.2"),
        },
        Name: "bitcoin-main",
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Connection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [connectivity.ConnectionCreate](../../pkg/models/connectivity/connectioncreate.md)             | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `security`                                                                                     | [operations.CreateConnectionSecurity](../../pkg/models/operations/createconnectionsecurity.md) | :heavy_check_mark:                                                                             | The security requirements to use for the request.                                              |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.CreateConnectionResponse](../../pkg/models/operations/createconnectionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| connectivity.Error | 4XX                | application/json   |
| connectivity.Error | 5XX                | application/json   |

## DeleteConnection

Delete a Connection

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteConnection" method="delete" path="/api/connectivity/connections/{name}" -->
```go
package main

import(
	"context"
	"github.com/formancehq/formance-sdk-go/v5"
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := v5.New()

    res, err := s.Connectivity.V1.DeleteConnection(ctx, operations.DeleteConnectionRequest{
        Name: "<value>",
    }, nil)
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
| `request`                                                                                      | [operations.DeleteConnectionRequest](../../pkg/models/operations/deleteconnectionrequest.md)   | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `security`                                                                                     | [operations.DeleteConnectionSecurity](../../pkg/models/operations/deleteconnectionsecurity.md) | :heavy_check_mark:                                                                             | The security requirements to use for the request.                                              |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.DeleteConnectionResponse](../../pkg/models/operations/deleteconnectionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| connectivity.Error | 4XX                | application/json   |
| connectivity.Error | 5XX                | application/json   |

## GetConnection

Get a Connection

### Example Usage

<!-- UsageSnippet language="go" operationID="getConnection" method="get" path="/api/connectivity/connections/{name}" -->
```go
package main

import(
	"context"
	"github.com/formancehq/formance-sdk-go/v5"
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := v5.New()

    res, err := s.Connectivity.V1.GetConnection(ctx, operations.GetConnectionRequest{
        Name: "<value>",
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Connection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetConnectionRequest](../../pkg/models/operations/getconnectionrequest.md)   | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `security`                                                                               | [operations.GetConnectionSecurity](../../pkg/models/operations/getconnectionsecurity.md) | :heavy_check_mark:                                                                       | The security requirements to use for the request.                                        |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.GetConnectionResponse](../../pkg/models/operations/getconnectionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| connectivity.Error | 4XX                | application/json   |
| connectivity.Error | 5XX                | application/json   |

## GetConnector

Get a Connector (lean connector metadata)

### Example Usage

<!-- UsageSnippet language="go" operationID="getConnector" method="get" path="/api/connectivity/connectors/{name}" -->
```go
package main

import(
	"context"
	"github.com/formancehq/formance-sdk-go/v5"
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := v5.New()

    res, err := s.Connectivity.V1.GetConnector(ctx, operations.GetConnectorRequest{
        Name: "<value>",
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Connector != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetConnectorRequest](../../pkg/models/operations/getconnectorrequest.md)   | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `security`                                                                             | [operations.GetConnectorSecurity](../../pkg/models/operations/getconnectorsecurity.md) | :heavy_check_mark:                                                                     | The security requirements to use for the request.                                      |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.GetConnectorResponse](../../pkg/models/operations/getconnectorresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| connectivity.Error | 4XX                | application/json   |
| connectivity.Error | 5XX                | application/json   |

## GetConnectorFacets

Counts distinct published `Connector`s per `facet:value` tag (split on
the first `:`; flat tags carry no facet and are excluded), narrowed by
the same `query` filter as the list. `_facets` can never collide with
a connector: names are DNS labels, which cannot contain `_`.


### Example Usage

<!-- UsageSnippet language="go" operationID="getConnectorFacets" method="get" path="/api/connectivity/connectors/_facets" -->
```go
package main

import(
	"context"
	"github.com/formancehq/formance-sdk-go/v5"
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := v5.New()

    res, err := s.Connectivity.V1.GetConnectorFacets(ctx, operations.GetConnectorFacetsRequest{}, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.FacetDistribution != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.GetConnectorFacetsRequest](../../pkg/models/operations/getconnectorfacetsrequest.md)   | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `security`                                                                                         | [operations.GetConnectorFacetsSecurity](../../pkg/models/operations/getconnectorfacetssecurity.md) | :heavy_check_mark:                                                                                 | The security requirements to use for the request.                                                  |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.GetConnectorFacetsResponse](../../pkg/models/operations/getconnectorfacetsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| connectivity.Error | 4XX                | application/json   |
| connectivity.Error | 5XX                | application/json   |

## GetConnectorVersion

Get one Connector version (including its configSchema and additionalMetadata)

### Example Usage

<!-- UsageSnippet language="go" operationID="getConnectorVersion" method="get" path="/api/connectivity/connectors/{name}/versions/{version}" -->
```go
package main

import(
	"context"
	"github.com/formancehq/formance-sdk-go/v5"
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := v5.New()

    res, err := s.Connectivity.V1.GetConnectorVersion(ctx, operations.GetConnectorVersionRequest{
        Name: "<value>",
        Version: "<value>",
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ConnectorVersion != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.GetConnectorVersionRequest](../../pkg/models/operations/getconnectorversionrequest.md)   | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `security`                                                                                           | [operations.GetConnectorVersionSecurity](../../pkg/models/operations/getconnectorversionsecurity.md) | :heavy_check_mark:                                                                                   | The security requirements to use for the request.                                                    |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.GetConnectorVersionResponse](../../pkg/models/operations/getconnectorversionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| connectivity.Error | 4XX                | application/json   |
| connectivity.Error | 5XX                | application/json   |

## GetQueryCapabilities

Serialized from the same allowlist the server enforces on `?query=`,
so a UI can build its filter bar dynamically and never drift from
what the API accepts.


### Example Usage

<!-- UsageSnippet language="go" operationID="getQueryCapabilities" method="get" path="/api/connectivity/_query/capabilities" -->
```go
package main

import(
	"context"
	"github.com/formancehq/formance-sdk-go/v5"
	"log"
)

func main() {
    ctx := context.Background()

    s := v5.New()

    res, err := s.Connectivity.V1.GetQueryCapabilities(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.QueryCapabilities != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `security`                                                                                             | [operations.GetQueryCapabilitiesSecurity](../../pkg/models/operations/getquerycapabilitiessecurity.md) | :heavy_check_mark:                                                                                     | The security requirements to use for the request.                                                      |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.GetQueryCapabilitiesResponse](../../pkg/models/operations/getquerycapabilitiesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| connectivity.Error | 4XX                | application/json   |
| connectivity.Error | 5XX                | application/json   |

## Healthcheck

Single health endpoint, probed by the stack gateway for the `/versions`
aggregation and by service readiness/liveness probes. Ready
means the service can reach its storage dependency.


### Example Usage

<!-- UsageSnippet language="go" operationID="healthcheck" method="get" path="/api/connectivity/_healthcheck" -->
```go
package main

import(
	"context"
	"github.com/formancehq/formance-sdk-go/v5"
	"log"
)

func main() {
    ctx := context.Background()

    s := v5.New()

    res, err := s.Connectivity.V1.Healthcheck(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.Health != nil {
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

**[*operations.HealthcheckResponse](../../pkg/models/operations/healthcheckresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| connectivity.HealthError | 500                      | application/json         |
| connectivity.Error       | 4XX                      | application/json         |
| connectivity.Error       | 5XX                      | application/json         |

## ListConnections

List connector instances for this tenant

### Example Usage

<!-- UsageSnippet language="go" operationID="listConnections" method="get" path="/api/connectivity/connections" -->
```go
package main

import(
	"context"
	"github.com/formancehq/formance-sdk-go/v5"
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := v5.New()

    res, err := s.Connectivity.V1.ListConnections(ctx, operations.ListConnectionsRequest{}, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ConnectionCursor != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListConnectionsRequest](../../pkg/models/operations/listconnectionsrequest.md)   | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `security`                                                                                   | [operations.ListConnectionsSecurity](../../pkg/models/operations/listconnectionssecurity.md) | :heavy_check_mark:                                                                           | The security requirements to use for the request.                                            |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.ListConnectionsResponse](../../pkg/models/operations/listconnectionsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| connectivity.Error | 4XX                | application/json   |
| connectivity.Error | 5XX                | application/json   |

## ListConnectorVersions

Returns the `ConnectorVersion`s published for this Connector, sorted
ascending by semantic version. Use `configuration.version` to pin a Connection,
or `configuration.channel` to track a maturity channel.


### Example Usage

<!-- UsageSnippet language="go" operationID="listConnectorVersions" method="get" path="/api/connectivity/connectors/{name}/versions" -->
```go
package main

import(
	"context"
	"github.com/formancehq/formance-sdk-go/v5"
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := v5.New()

    res, err := s.Connectivity.V1.ListConnectorVersions(ctx, operations.ListConnectorVersionsRequest{
        Name: "<value>",
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ConnectorVersionCursor != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.ListConnectorVersionsRequest](../../pkg/models/operations/listconnectorversionsrequest.md)   | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `security`                                                                                               | [operations.ListConnectorVersionsSecurity](../../pkg/models/operations/listconnectorversionssecurity.md) | :heavy_check_mark:                                                                                       | The security requirements to use for the request.                                                        |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.ListConnectorVersionsResponse](../../pkg/models/operations/listconnectorversionsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| connectivity.Error | 4XX                | application/json   |
| connectivity.Error | 5XX                | application/json   |

## ListConnectors

Returns the published `Connector`s, each carrying the identity of its
newest Validated version (branding, tagline, tags, `latestVersion`) so
a catalogue UI renders and filters from this single request.


### Example Usage

<!-- UsageSnippet language="go" operationID="listConnectors" method="get" path="/api/connectivity/connectors" -->
```go
package main

import(
	"context"
	"github.com/formancehq/formance-sdk-go/v5"
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := v5.New()

    res, err := s.Connectivity.V1.ListConnectors(ctx, operations.ListConnectorsRequest{}, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ConnectorCursor != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ListConnectorsRequest](../../pkg/models/operations/listconnectorsrequest.md)   | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `security`                                                                                 | [operations.ListConnectorsSecurity](../../pkg/models/operations/listconnectorssecurity.md) | :heavy_check_mark:                                                                         | The security requirements to use for the request.                                          |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.ListConnectorsResponse](../../pkg/models/operations/listconnectorsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| connectivity.Error | 4XX                | application/json   |
| connectivity.Error | 5XX                | application/json   |

## PatchConnection

JSON merge patch (RFC 7396) over tenant `configuration` fields only. Other fields
are rejected with `400 unsupported_field`. `env` and `files` are string
maps: omitted entries are preserved, including secret-backed values;
`null` deletes a field or map entry. A supplied literal replaces the
stored source. `connector` and `ledger` are not PATCH fields: any presence, including
unchanged values and null, returns `400 unsupported_field`. The merged
result must retain exactly one of `version` and `channel`, so
switching selectors sends `null` for the other one.


### Example Usage: rotatePassword

<!-- UsageSnippet language="go" operationID="patchConnection" method="patch" path="/api/connectivity/connections/{name}" example="rotatePassword" -->
```go
package main

import(
	"context"
	"github.com/formancehq/formance-sdk-go/v5"
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/connectivity"
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := v5.New()

    res, err := s.Connectivity.V1.PatchConnection(ctx, operations.PatchConnectionRequest{
        ConnectionPatch: connectivity.ConnectionPatch{
            Configuration: &connectivity.ConnectionPatchConfiguration{
                Env: map[string]*string{
                    "STRIPE_API_KEY": v5.Pointer("<new-api-key>"),
                },
            },
        },
        Name: "<value>",
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Connection != nil {
        // handle response
    }
}
```
### Example Usage: suspend

<!-- UsageSnippet language="go" operationID="patchConnection" method="patch" path="/api/connectivity/connections/{name}" example="suspend" -->
```go
package main

import(
	"context"
	"github.com/formancehq/formance-sdk-go/v5"
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/connectivity"
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := v5.New()

    res, err := s.Connectivity.V1.PatchConnection(ctx, operations.PatchConnectionRequest{
        ConnectionPatch: connectivity.ConnectionPatch{
            Configuration: &connectivity.ConnectionPatchConfiguration{
                Suspend: v5.Pointer(true),
            },
        },
        Name: "<value>",
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Connection != nil {
        // handle response
    }
}
```
### Example Usage: switchToChannel

<!-- UsageSnippet language="go" operationID="patchConnection" method="patch" path="/api/connectivity/connections/{name}" example="switchToChannel" -->
```go
package main

import(
	"context"
	"github.com/formancehq/formance-sdk-go/v5"
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/connectivity"
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := v5.New()

    res, err := s.Connectivity.V1.PatchConnection(ctx, operations.PatchConnectionRequest{
        ConnectionPatch: connectivity.ConnectionPatch{
            Configuration: &connectivity.ConnectionPatchConfiguration{
                Channel: connectivity.ConnectionPatchChannelStable.ToPointer(),
                Version: nil,
            },
        },
        Name: "<value>",
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Connection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.PatchConnectionRequest](../../pkg/models/operations/patchconnectionrequest.md)   | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `security`                                                                                   | [operations.PatchConnectionSecurity](../../pkg/models/operations/patchconnectionsecurity.md) | :heavy_check_mark:                                                                           | The security requirements to use for the request.                                            |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.PatchConnectionResponse](../../pkg/models/operations/patchconnectionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| connectivity.Error | 4XX                | application/json   |
| connectivity.Error | 5XX                | application/json   |

## ReplaceConnection

Full-replace semantics on mutable tenant `configuration`: omitted mutable
fields are cleared. The API does not write fields outside its tenant scope.
`status` is preserved. `connector` and `ledger` are immutable once set:
omission preserves them, first assignment and unchanged values are accepted,
and changes or null return `422 connector_immutable` or `422 ledger_immutable`. Read responses omit reference-backed entries,
so PUT must supply the complete configuration, including required secrets.
Use PATCH to preserve omitted configuration. Fields outside the tenant contract are rejected with
`400 unsupported_field`.


### Example Usage

<!-- UsageSnippet language="go" operationID="replaceConnection" method="put" path="/api/connectivity/connections/{name}" -->
```go
package main

import(
	"context"
	"github.com/formancehq/formance-sdk-go/v5"
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/connectivity"
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := v5.New()

    res, err := s.Connectivity.V1.ReplaceConnection(ctx, operations.ReplaceConnectionRequest{
        ConnectionUpdate: connectivity.ConnectionUpdate{
            Configuration: connectivity.ConnectionConfiguration{},
        },
        Name: "<value>",
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Connection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.ReplaceConnectionRequest](../../pkg/models/operations/replaceconnectionrequest.md)   | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `security`                                                                                       | [operations.ReplaceConnectionSecurity](../../pkg/models/operations/replaceconnectionsecurity.md) | :heavy_check_mark:                                                                               | The security requirements to use for the request.                                                |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.ReplaceConnectionResponse](../../pkg/models/operations/replaceconnectionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| connectivity.Error | 4XX                | application/json   |
| connectivity.Error | 5XX                | application/json   |