# Ledger.V3

## Overview

### Available Operations

* [AddAccountType](#addaccounttype) - Add an account type
* [AggregateVolumes](#aggregatevolumes) - Aggregate volumes
* [AnalyzeAccounts](#analyzeaccounts) - Analyze accounts and suggest a Chart of Accounts
* [AnalyzeNumscript](#analyzenumscript) - Analyze Numscript source
* [AnalyzeTransactions](#analyzetransactions) - Analyze transaction flow patterns
* [BulkOperations](#bulkoperations) - Bulk operations
* [Clusterz](#clusterz) - Cluster-availability probe
* [CreateIndex](#createindex) - Create an index on a ledger
* [CreateLedger](#createledger) - Create a new ledger
* [CreatePreparedQuery](#createpreparedquery) - Create a prepared query
* [CreateTransaction](#createtransaction) - Create a new transaction
* [DeleteAccountMetadata](#deleteaccountmetadata) - Delete account metadata
* [DeleteLedger](#deleteledger) - Delete a ledger
* [DeleteLedgerMetadata](#deleteledgermetadata) - Delete ledger metadata
* [DeletePreparedQuery](#deletepreparedquery) - Delete a prepared query
* [DeleteTransactionMetadata](#deletetransactionmetadata) - Delete transaction metadata
* [DropIndex](#dropindex) - Drop an index from a ledger
* [ExecutePreparedQuery](#executepreparedquery) - Execute a prepared query
* [GetAccount](#getaccount) - Get an account
* [GetAccountType](#getaccounttype) - Get an account type
* [GetAuditEntry](#getauditentry) - Get an audit entry
* [GetBucketIndex](#getbucketindex) - Get a single bucket-scoped index registry entry
* [GetBucketIndexEntryStatus](#getbucketindexentrystatus) - Get the per-replica status of a single bucket-scoped index
* [GetEventsSinks](#geteventssinks) - List configured event sinks
* [GetIndex](#getindex) - Get a single index registry entry
* [GetIndexEntryStatus](#getindexentrystatus) - Get the per-replica status of a single index
* [GetIndexStatus](#getindexstatus) - Get aggregated index registry status
* [GetLedger](#getledger) - Get a ledger
* [GetLedgerStats](#getledgerstats) - Get ledger statistics
* [GetLog](#getlog) - Get a single system log by bucket-wide sequence
* [GetMetadataSchemaStatus](#getmetadataschemastatus) - Get metadata schema
* [GetNumscript](#getnumscript) - Get a numscript
* [GetNumscriptUsage](#getnumscriptusage) - Get template usage
* [GetTransaction](#gettransaction) - Get a transaction
* [Health](#health) - Health check
* [Info](#info) - Get server build information
* [InspectIndex](#inspectindex) - Inspect a metadata index
* [ListAccountTypes](#listaccounttypes) - List account types
* [ListAccounts](#listaccounts) - List accounts
* [ListAllLedgers](#listallledgers) - List all ledgers
* [ListAuditEntries](#listauditentries) - List audit entries
* [ListBucketIndexes](#listbucketindexes) - List bucket-wide indexes
* [ListLedgerIndexes](#listledgerindexes) - List indexes for a ledger
* [ListLedgerLogs](#listledgerlogs) - List logs for a ledger
* [ListLedgerTransactions](#listledgertransactions) - List transactions for a ledger
* [ListNumscriptVersions](#listnumscriptversions) - List numscript versions
* [ListNumscripts](#listnumscripts) - List all numscripts
* [ListPreparedQueries](#listpreparedqueries) - List prepared queries
* [ListSigningKeys](#listsigningkeys) - List registered Ed25519 signing keys
* [Livez](#livez) - Liveness probe
* [PromoteLedger](#promoteledger) - Promote a mirror ledger
* [Readyz](#readyz) - Readiness probe
* [RemoveAccountType](#removeaccounttype) - Remove an account type
* [RemoveEventsSink](#removeeventssink) - Remove an event sink
* [RemoveMetadataFieldType](#removemetadatafieldtype) - Remove metadata field type
* [RevertTransaction](#reverttransaction) - Revert a transaction
* [SaveAccountMetadata](#saveaccountmetadata) - Save account metadata
* [SaveLedgerMetadata](#saveledgermetadata) - Save ledger metadata
* [SaveNumscript](#savenumscript) - Save a numscript
* [SaveTransactionMetadata](#savetransactionmetadata) - Save transaction metadata
* [SetDefaultEnforcementMode](#setdefaultenforcementmode) - Set default enforcement mode
* [SetMetadataFieldType](#setmetadatafieldtype) - Set metadata field type
* [UpdatePreparedQuery](#updatepreparedquery) - Update a prepared query

## AddAccountType

Adds a new account type to a ledger's chart of accounts. Account types define
address patterns for account validation. Requires `ledger:WriteMetadata` scope.


### Example Usage

<!-- UsageSnippet language="go" operationID="v3AddAccountType" method="post" path="/api/ledger/v3/{ledgerName}/account-types" -->
```go
package main

import(
	"context"
	"github.com/formancehq/formance-sdk-go/v5"
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/ledgerv3"
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := v5.New()

    res, err := s.Ledger.V3.AddAccountType(ctx, operations.V3AddAccountTypeRequest{
        AddAccountTypeRequest: ledgerv3.AddAccountTypeRequest{
            Name: "user-checking",
            Pattern: "users:{id}:checking",
            Persistence: ledgerv3.AddAccountTypeRequestPersistenceEphemeral.ToPointer(),
            SegmentTypes: map[string]*ledgerv3.SegmentType{
                "id": &ledgerv3.SegmentType{
                    Type: ledgerv3.SegmentTypeTypeUUID,
                },
            },
        },
        LedgerName: "<value>",
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.V3AddAccountTypeRequest](../../pkg/models/operations/v3addaccounttyperequest.md)   | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `security`                                                                                     | [operations.V3AddAccountTypeSecurity](../../pkg/models/operations/v3addaccounttypesecurity.md) | :heavy_check_mark:                                                                             | The security requirements to use for the request.                                              |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.V3AddAccountTypeResponse](../../pkg/models/operations/v3addaccounttyperesponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## AggregateVolumes

Returns aggregated volumes (input/output/balance per asset) across all accounts in the ledger.
Supports filtering accounts with the generic `filter` parameter, grouping by account prefixes,
and max-precision merging.
Requires `ledger:ReadAccount` scope.


### Example Usage

<!-- UsageSnippet language="go" operationID="v3AggregateVolumes" method="get" path="/api/ledger/v3/{ledgerName}/volumes" -->
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

    res, err := s.Ledger.V3.AggregateVolumes(ctx, operations.V3AggregateVolumesRequest{
        FormanceBigintAsString: v5.Pointer("true"),
        LedgerName: "<value>",
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.AggregateVolumesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.V3AggregateVolumesRequest](../../pkg/models/operations/v3aggregatevolumesrequest.md)   | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `security`                                                                                         | [operations.V3AggregateVolumesSecurity](../../pkg/models/operations/v3aggregatevolumessecurity.md) | :heavy_check_mark:                                                                                 | The security requirements to use for the request.                                                  |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.V3AggregateVolumesResponse](../../pkg/models/operations/v3aggregatevolumesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## AnalyzeAccounts

Scans all accounts in a ledger, discovers address patterns, and returns a
suggested Chart of Accounts. Useful after a mirror import to understand
account structure. Requires `ledger:read` scope.


### Example Usage

<!-- UsageSnippet language="go" operationID="v3AnalyzeAccounts" method="get" path="/api/ledger/v3/{ledgerName}/analyze-accounts" -->
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

    res, err := s.Ledger.V3.AnalyzeAccounts(ctx, operations.V3AnalyzeAccountsRequest{
        LedgerName: "<value>",
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.AnalyzeAccountsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.V3AnalyzeAccountsRequest](../../pkg/models/operations/v3analyzeaccountsrequest.md)   | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `security`                                                                                       | [operations.V3AnalyzeAccountsSecurity](../../pkg/models/operations/v3analyzeaccountssecurity.md) | :heavy_check_mark:                                                                               | The security requirements to use for the request.                                                |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.V3AnalyzeAccountsResponse](../../pkg/models/operations/v3analyzeaccountsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## AnalyzeNumscript

Validates Numscript v1 and returns the external input schema without reading or changing ledger state. Recognized source failures are returned as diagnostics with an empty input list. Sources over 64 KiB or with more than 32 block-comment open markers receive a syntax diagnostic; internal failures return a server error.

### Example Usage

<!-- UsageSnippet language="go" operationID="v3AnalyzeNumscript" method="post" path="/api/ledger/v3/numscripts/analyze" -->
```go
package main

import(
	"context"
	"github.com/formancehq/formance-sdk-go/v5"
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/ledgerv3"
	"log"
)

func main() {
    ctx := context.Background()

    s := v5.New()

    res, err := s.Ledger.V3.AnalyzeNumscript(ctx, ledgerv3.AnalyzeNumscriptRequest{
        Source: "<value>",
        Version: "v1",
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.AnalyzeNumscriptResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [ledgerv3.AnalyzeNumscriptRequest](../../pkg/models/ledgerV3/analyzenumscriptrequest.md)           | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `security`                                                                                         | [operations.V3AnalyzeNumscriptSecurity](../../pkg/models/operations/v3analyzenumscriptsecurity.md) | :heavy_check_mark:                                                                                 | The security requirements to use for the request.                                                  |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.V3AnalyzeNumscriptResponse](../../pkg/models/operations/v3analyzenumscriptresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## AnalyzeTransactions

Scans all transactions in a ledger, discovers flow patterns by normalizing
account addresses, and returns statistics per flow type including temporal
distribution and volume metrics. Requires `ledger:read` scope.


### Example Usage

<!-- UsageSnippet language="go" operationID="v3AnalyzeTransactions" method="get" path="/api/ledger/v3/{ledgerName}/analyze-transactions" -->
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

    res, err := s.Ledger.V3.AnalyzeTransactions(ctx, operations.V3AnalyzeTransactionsRequest{
        FormanceBigintAsString: v5.Pointer("true"),
        LedgerName: "<value>",
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.AnalyzeTransactionsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.V3AnalyzeTransactionsRequest](../../pkg/models/operations/v3analyzetransactionsrequest.md)   | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `security`                                                                                               | [operations.V3AnalyzeTransactionsSecurity](../../pkg/models/operations/v3analyzetransactionssecurity.md) | :heavy_check_mark:                                                                                       | The security requirements to use for the request.                                                        |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.V3AnalyzeTransactionsResponse](../../pkg/models/operations/v3analyzetransactionsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## BulkOperations

Execute multiple operations (create transactions, add metadata,
revert transactions, delete metadata) in a single request. Requires
`ledger:write` scope.

With `atomic=true`, all elements target the URL's ledger and are
submitted as one batch. `Idempotency-Key` identifies that whole batch;
per-element `ik` values are ignored. With `atomic=false`
(default), each element is submitted separately under its own
`ik`; the header is ignored.

`continueOnFailure` controls sequential continuation and business-error
status rollup. It does not allow partial application of an atomic
batch. Non-atomic failures can follow already committed elements:
inspect every result and preserve the original identities when retrying.
A timeout or infrastructure failure does not prove that work was not
committed. This contract does not enable automatic SDK retries.

The response status is decided in three tiers:

1. **Request-level failures** (before per-element processing):
   malformed body → `400`, missing/insufficient scope → `401` or
   `403`, too many elements → `413`. Exceeding the byte limit while
   decoding the body currently returns `400`.
   Returned unconditionally, `continueOnFailure` does not apply.
2. **Infrastructure failures during processing** (transport
   timeout, cache-horizon exceeded, storage error): the overall
   status is `503` for retryable failures and `500` otherwise,
   regardless of `continueOnFailure`. Callers
   must inspect the element outcomes before deciding what to retry.
3. **Per-element errors during processing**: each element's error
   carries the HTTP status it would surface as a single request
   (via the same mapping `handleError` uses). The top-level
   status is the "worst" of them, with `continueOnFailure`
   applying only to per-element business failures (4xx):
   - Business failures (validation, not-found, conflict,
     precondition — 4xx, except authentication/permission): rolled up per
     `continueOnFailure`. `true` → `200`; `false` (default) →
     the highest 4xx among failing elements (typically `400`).
   - Authentication and permission failures: `401`/`403`, regardless
     of `continueOnFailure`, with all processing outcomes retained.
   - Resource exhaustion (`KindResourceExhausted`, `429`), including
     rate limits and permanent authoritative-sequence exhaustion:
     surface unconditionally as `429`.
   - Retryable infrastructure failures (leader loss,
     cache-horizon exceeded, `KindUnavailable`): surface as
     `503` with `Retry-After: 1`, regardless of
     `continueOnFailure`.
   - Non-retryable infrastructure failures (transport / storage
     errors, `KindInternal`, unclassified): surface as `500`.
   - Aborted elements (`context.Canceled` from the sequential
     runner after a prior failure): contribute nothing — the
     originating error drives the rollup.
Signed requests use `application/json` (including MIME parameters) with
`request` and standard-base64 `payload`. Either signing header, even if
empty, selects this wrapper; with neither header present, the ordinary REST body applies.
Runtime headers choose the representation from the documented anyOf.
The signed protobuf batch must
equal the route/body/Idempotency-Key-derived batch in full. Supply exactly
one nonempty Formance-Signing-Key-Id and Formance-Signature header; the
latter is a standard-base64 Ed25519 signature of 64 decoded bytes. Missing,
empty or repeated signing headers, invalid base64 and wrong signature
length return 400. Signing headers with a non-JSON or missing content type return 415.
Signature enforcement
returns 401 `UNAUTHENTICATED` when required but absent and 403
`PERMISSION_DENIED` for invalid, unknown/revoked-key or mismatched
signatures. Malformed signed envelopes return 400. The complete
wrapper counts toward the 4 MiB body limit. JWT scopes still apply.
Signing requires a nonempty array and `atomic=true` (otherwise 400).
The header key applies; per-element `ik` is ignored. Authentication
and permission failures remain 401/403 even with continueOnFailure;
processing errors preserve every element outcome.


### Example Usage

<!-- UsageSnippet language="go" operationID="v3BulkOperations" method="post" path="/api/ledger/v3/{ledgerName}/bulk" -->
```go
package main

import(
	"context"
	"github.com/formancehq/formance-sdk-go/v5"
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/ledgerv3"
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := v5.New()

    res, err := s.Ledger.V3.BulkOperations(ctx, operations.V3BulkOperationsRequest{
        FormanceBigintAsString: v5.Pointer("true"),
        RequestBody: operations.CreateV3BulkOperationsRequestBodyArrayOfBulkElement(
            []ledgerv3.BulkElement{
                ledgerv3.BulkElement{
                    Action: ledgerv3.ActionRevertTransaction,
                    Data: ledgerv3.CreateBulkElementDataRevertTransactionRequest(
                        ledgerv3.RevertTransactionRequest{
                            ID: 725647,
                        },
                    ),
                },
            },
        ),
        LedgerName: "<value>",
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.BulkResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.V3BulkOperationsRequest](../../pkg/models/operations/v3bulkoperationsrequest.md)   | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `security`                                                                                     | [operations.V3BulkOperationsSecurity](../../pkg/models/operations/v3bulkoperationssecurity.md) | :heavy_check_mark:                                                                             | The security requirements to use for the request.                                              |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.V3BulkOperationsResponse](../../pkg/models/operations/v3bulkoperationsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Clusterz

Returns 200 only when the node is part of a healthy cluster (Raft state
healthy and a leader elected). Disk and clock skew gate writes (see
admission), not cluster readiness. Returns 503 otherwise. Use this
for external monitoring or any client that needs to wait until the node
can actually serve cluster-dependent traffic. Kubernetes readiness
probes should target `/readyz` instead — this endpoint MUST NOT gate
StatefulSet ordering, otherwise the cluster cannot recover from a
full-kill cold start.


### Example Usage

<!-- UsageSnippet language="go" operationID="v3Clusterz" method="get" path="/api/ledger/clusterz" -->
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

    res, err := s.Ledger.V3.Clusterz(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.HealthResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `security`                                                                         | [operations.V3ClusterzSecurity](../../pkg/models/operations/v3clusterzsecurity.md) | :heavy_check_mark:                                                                 | The security requirements to use for the request.                                  |
| `opts`                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                       | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.V3ClusterzResponse](../../pkg/models/operations/v3clusterzresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateIndex

Registers a new index on the ledger. An existing canonical index on
this ledger is rejected with `INDEX_ALREADY_EXISTS` (HTTP 409), even
while building or retyping, without changing its definition or progress.
Replaying an identical batch with its retained `Idempotency-Key` returns
the original result. The indexbuilder processes the accepted creation.
The request body carries the target `IndexID` in
canonical form (see `DELETE /v3/{ledgerName}/indexes/{canonicalId}`
for the format). Requires `ledger:WriteLedger` scope.


### Example Usage

<!-- UsageSnippet language="go" operationID="v3CreateIndex" method="post" path="/api/ledger/v3/{ledgerName}/indexes" -->
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

    res, err := s.Ledger.V3.CreateIndex(ctx, operations.V3CreateIndexRequest{
        RequestBody: operations.V3CreateIndexRequestBody{
            ID: "metadata:TARGET_TYPE_ACCOUNT:color",
        },
        LedgerName: "<value>",
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateIndexResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.V3CreateIndexRequest](../../pkg/models/operations/v3createindexrequest.md)   | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `security`                                                                               | [operations.V3CreateIndexSecurity](../../pkg/models/operations/v3createindexsecurity.md) | :heavy_check_mark:                                                                       | The security requirements to use for the request.                                        |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.V3CreateIndexResponse](../../pkg/models/operations/v3createindexresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateLedger

Creates a new ledger with the specified name. Storage is configured at the server level.
Each ledger has its own Raft group for data consistency. Requires `ledger:write` scope.

The ledger name `_` is reserved: all system / non-ledger routes live
under the `/v3/_/…` segment (e.g. `/v3/_/audit-entries`), so a ledger
named `_` would be unreachable over REST. Creating it is rejected with a
validation error. The exact name `_system` is also reserved for system
events and rejected. Other underscore-prefixed names remain valid.


### Example Usage

<!-- UsageSnippet language="go" operationID="v3CreateLedger" method="post" path="/api/ledger/v3/{ledgerName}" -->
```go
package main

import(
	"context"
	"github.com/formancehq/formance-sdk-go/v5"
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/ledgerv3"
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := v5.New()

    res, err := s.Ledger.V3.CreateLedger(ctx, operations.V3CreateLedgerRequest{
        CreateLedgerRequest: ledgerv3.CreateLedgerRequest{
            AccountTypes: map[string]ledgerv3.AddAccountTypeRequest{
                "user-checking": ledgerv3.AddAccountTypeRequest{
                    Name: "user-checking",
                    Pattern: "users:{id}:checking",
                    Persistence: ledgerv3.AddAccountTypeRequestPersistenceEphemeral.ToPointer(),
                    SegmentTypes: map[string]*ledgerv3.SegmentType{
                        "id": &ledgerv3.SegmentType{
                            Type: ledgerv3.SegmentTypeTypeUUID,
                        },
                    },
                },
            },
            InitialSchema: []ledgerv3.MetadataFieldTypeCommand{
                ledgerv3.MetadataFieldTypeCommand{
                    Key: "color",
                    TargetType: ledgerv3.MetadataFieldTypeCommandTargetTypeAccount,
                    Type: ledgerv3.MetadataFieldTypeCommandTypeString,
                },
            },
        },
        LedgerName: "<value>",
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateLedgerResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.V3CreateLedgerRequest](../../pkg/models/operations/v3createledgerrequest.md)   | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `security`                                                                                 | [operations.V3CreateLedgerSecurity](../../pkg/models/operations/v3createledgersecurity.md) | :heavy_check_mark:                                                                         | The security requirements to use for the request.                                          |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.V3CreateLedgerResponse](../../pkg/models/operations/v3createledgerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreatePreparedQuery

Creates a new prepared query for a ledger. A prepared query defines a reusable
filter that can be executed later with parameter substitution. Requires `ledger:WriteQuery` scope.


### Example Usage

<!-- UsageSnippet language="go" operationID="v3CreatePreparedQuery" method="post" path="/api/ledger/v3/{ledgerName}/prepared-queries" -->
```go
package main

import(
	"context"
	"github.com/formancehq/formance-sdk-go/v5"
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/ledgerv3"
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := v5.New()

    res, err := s.Ledger.V3.CreatePreparedQuery(ctx, operations.V3CreatePreparedQueryRequest{
        CreatePreparedQueryRequest: ledgerv3.CreatePreparedQueryRequest{
            Name: "<value>",
            Target: ledgerv3.TargetAccounts,
        },
        LedgerName: "<value>",
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.V3CreatePreparedQueryRequest](../../pkg/models/operations/v3createpreparedqueryrequest.md)   | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `security`                                                                                               | [operations.V3CreatePreparedQuerySecurity](../../pkg/models/operations/v3createpreparedquerysecurity.md) | :heavy_check_mark:                                                                                       | The security requirements to use for the request.                                                        |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.V3CreatePreparedQueryResponse](../../pkg/models/operations/v3createpreparedqueryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateTransaction

Creates a new transaction in the specified ledger with the specified postings. Requires `ledger:write` scope.
Signed requests use `application/json` (including MIME parameters) with
`request` and standard-base64 `payload`. Either signing header, even if
empty, selects this wrapper; with neither header present, the ordinary REST body applies.
Runtime headers choose the representation from the documented anyOf.
The signed protobuf batch must
equal the route/body/Idempotency-Key-derived batch in full. Supply exactly
one nonempty Formance-Signing-Key-Id and Formance-Signature header; the
latter is a standard-base64 Ed25519 signature of 64 decoded bytes. Missing,
empty or repeated signing headers, invalid base64 and wrong signature
length return 400. Signing headers with a non-JSON or missing content type return 415.
Signature enforcement
returns 401 `UNAUTHENTICATED` when required but absent and 403
`PERMISSION_DENIED` for invalid, unknown/revoked-key or mismatched
signatures. Malformed signed envelopes return 400. The complete
wrapper counts toward the 4 MiB body limit. JWT scopes still apply.


### Example Usage

<!-- UsageSnippet language="go" operationID="v3CreateTransaction" method="post" path="/api/ledger/v3/{ledgerName}/transactions" -->
```go
package main

import(
	"context"
	"github.com/formancehq/formance-sdk-go/v5"
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/ledgerv3"
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := v5.New()

    res, err := s.Ledger.V3.CreateTransaction(ctx, operations.V3CreateTransactionRequest{
        FormanceBigintAsString: v5.Pointer("true"),
        RequestBody: operations.CreateV3CreateTransactionRequestBodyCreateTransactionRequest(
            ledgerv3.CreateTransactionRequest{},
        ),
        LedgerName: "<value>",
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateTransactionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.V3CreateTransactionRequest](../../pkg/models/operations/v3createtransactionrequest.md)   | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `security`                                                                                           | [operations.V3CreateTransactionSecurity](../../pkg/models/operations/v3createtransactionsecurity.md) | :heavy_check_mark:                                                                                   | The security requirements to use for the request.                                                    |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.V3CreateTransactionResponse](../../pkg/models/operations/v3createtransactionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteAccountMetadata

Deletes a metadata key for a specific account in the ledger. Requires `ledger:WriteMetadata` scope (included in the default `ledger:write` mapping).

### Example Usage

<!-- UsageSnippet language="go" operationID="v3DeleteAccountMetadata" method="delete" path="/api/ledger/v3/{ledgerName}/accounts/{address}/metadata/{key}" -->
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

    res, err := s.Ledger.V3.DeleteAccountMetadata(ctx, operations.V3DeleteAccountMetadataRequest{
        Address: "users:001",
        Key: "formance.com/reviewed",
        LedgerName: "<value>",
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.V3DeleteAccountMetadataRequest](../../pkg/models/operations/v3deleteaccountmetadatarequest.md)   | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `security`                                                                                                   | [operations.V3DeleteAccountMetadataSecurity](../../pkg/models/operations/v3deleteaccountmetadatasecurity.md) | :heavy_check_mark:                                                                                           | The security requirements to use for the request.                                                            |
| `opts`                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.V3DeleteAccountMetadataResponse](../../pkg/models/operations/v3deleteaccountmetadataresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteLedger

Deletes a ledger by its name. This operation removes the ledger and its associated Raft group from the cluster. Requires `ledger:write` scope.

### Example Usage

<!-- UsageSnippet language="go" operationID="v3DeleteLedger" method="delete" path="/api/ledger/v3/{ledgerName}" -->
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

    res, err := s.Ledger.V3.DeleteLedger(ctx, operations.V3DeleteLedgerRequest{
        LedgerName: "<value>",
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.V3DeleteLedgerRequest](../../pkg/models/operations/v3deleteledgerrequest.md)   | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `security`                                                                                 | [operations.V3DeleteLedgerSecurity](../../pkg/models/operations/v3deleteledgersecurity.md) | :heavy_check_mark:                                                                         | The security requirements to use for the request.                                          |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.V3DeleteLedgerResponse](../../pkg/models/operations/v3deleteledgerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteLedgerMetadata

Deletes a metadata key from the ledger. Requires `ledger:WriteMetadata` scope (included in the default `ledger:write` mapping).

### Example Usage

<!-- UsageSnippet language="go" operationID="v3DeleteLedgerMetadata" method="delete" path="/api/ledger/v3/{ledgerName}/metadata/{key}" -->
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

    res, err := s.Ledger.V3.DeleteLedgerMetadata(ctx, operations.V3DeleteLedgerMetadataRequest{
        Key: "formance.com/reviewed",
        LedgerName: "<value>",
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.V3DeleteLedgerMetadataRequest](../../pkg/models/operations/v3deleteledgermetadatarequest.md)   | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `security`                                                                                                 | [operations.V3DeleteLedgerMetadataSecurity](../../pkg/models/operations/v3deleteledgermetadatasecurity.md) | :heavy_check_mark:                                                                                         | The security requirements to use for the request.                                                          |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.V3DeleteLedgerMetadataResponse](../../pkg/models/operations/v3deleteledgermetadataresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeletePreparedQuery

Deletes a prepared query from a ledger. Requires `ledger:WriteQuery` scope.

### Example Usage

<!-- UsageSnippet language="go" operationID="v3DeletePreparedQuery" method="delete" path="/api/ledger/v3/{ledgerName}/prepared-queries/{queryName}" -->
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

    res, err := s.Ledger.V3.DeletePreparedQuery(ctx, operations.V3DeletePreparedQueryRequest{
        LedgerName: "<value>",
        QueryName: "<value>",
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.V3DeletePreparedQueryRequest](../../pkg/models/operations/v3deletepreparedqueryrequest.md)   | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `security`                                                                                               | [operations.V3DeletePreparedQuerySecurity](../../pkg/models/operations/v3deletepreparedquerysecurity.md) | :heavy_check_mark:                                                                                       | The security requirements to use for the request.                                                        |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.V3DeletePreparedQueryResponse](../../pkg/models/operations/v3deletepreparedqueryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteTransactionMetadata

Deletes a metadata key for a specific transaction in the ledger. Requires `ledger:WriteMetadata` scope (included in the default `ledger:write` mapping).
Signed requests use `application/json` (including MIME parameters) with
`request` and standard-base64 `payload`. Either signing header, even if
empty, selects this wrapper; with neither header present, the ordinary REST body applies.
Runtime headers choose the representation from the documented anyOf.
The signed protobuf batch must
equal the route/body/Idempotency-Key-derived batch in full. Supply exactly
one nonempty Formance-Signing-Key-Id and Formance-Signature header; the
latter is a standard-base64 Ed25519 signature of 64 decoded bytes. Missing,
empty or repeated signing headers, invalid base64 and wrong signature
length return 400. Signing headers with a non-JSON or missing content type return 415.
Signature enforcement
returns 401 `UNAUTHENTICATED` when required but absent and 403
`PERMISSION_DENIED` for invalid, unknown/revoked-key or mismatched
signatures. Malformed signed envelopes return 400. The complete
wrapper counts toward the 4 MiB body limit. JWT scopes still apply.


### Example Usage

<!-- UsageSnippet language="go" operationID="v3DeleteTransactionMetadata" method="delete" path="/api/ledger/v3/{ledgerName}/transactions/{transactionId}/metadata/{key}" -->
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

    res, err := s.Ledger.V3.DeleteTransactionMetadata(ctx, operations.V3DeleteTransactionMetadataRequest{
        Key: "formance.com/reviewed",
        LedgerName: "<value>",
        TransactionID: 831017,
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.V3DeleteTransactionMetadataRequest](../../pkg/models/operations/v3deletetransactionmetadatarequest.md)   | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `security`                                                                                                           | [operations.V3DeleteTransactionMetadataSecurity](../../pkg/models/operations/v3deletetransactionmetadatasecurity.md) | :heavy_check_mark:                                                                                                   | The security requirements to use for the request.                                                                    |
| `opts`                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |

### Response

**[*operations.V3DeleteTransactionMetadataResponse](../../pkg/models/operations/v3deletetransactionmetadataresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DropIndex

Removes an index from a ledger. The path parameter carries the target
`IndexID` in canonical form (e.g. `metadata:TARGET_TYPE_ACCOUNT:color`,
`tx_builtin:TX_BUILTIN_INDEX_TIMESTAMP`). Requires `ledger:WriteLedger` scope.

The drop is idempotent. An `IndexID` that has no registry entry on the
ledger (never created, or already dropped) also answers `204`, and each
committed drop appends a `DROP_INDEX` entry to the ledger log. A
mistyped canonical id is therefore not reported; read the index with
`GET /v3/{ledgerName}/indexes/{canonicalId}` first when the caller must
know that it exists. `404` means that the ledger does not exist.


### Example Usage

<!-- UsageSnippet language="go" operationID="v3DropIndex" method="delete" path="/api/ledger/v3/{ledgerName}/indexes/{canonicalId}" -->
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

    res, err := s.Ledger.V3.DropIndex(ctx, operations.V3DropIndexRequest{
        CanonicalID: "<id>",
        LedgerName: "<value>",
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.V3DropIndexRequest](../../pkg/models/operations/v3dropindexrequest.md)   | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `security`                                                                           | [operations.V3DropIndexSecurity](../../pkg/models/operations/v3dropindexsecurity.md) | :heavy_check_mark:                                                                   | The security requirements to use for the request.                                    |
| `opts`                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                         | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.V3DropIndexResponse](../../pkg/models/operations/v3dropindexresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ExecutePreparedQuery

Executes a prepared query with optional parameter substitution. Supports LIST and
AGGREGATE_VOLUMES modes. LIST uses cursor pagination with an effective page size
of 100 when omitted or zero, capped at 1000 before collecting and enriching rows.
AGGREGATE_VOLUMES ignores pageSize and aggregates all matching accounts. Also
accepts pageSize, cursor and reverse as query parameters. Supplied pageSize and
reverse values, and nonempty cursor values, override the body. Requires `ledger:ReadQuery` scope.


### Example Usage

<!-- UsageSnippet language="go" operationID="v3ExecutePreparedQuery" method="post" path="/api/ledger/v3/{ledgerName}/prepared-queries/{queryName}/execute" -->
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

    res, err := s.Ledger.V3.ExecutePreparedQuery(ctx, operations.V3ExecutePreparedQueryRequest{
        FormanceBigintAsString: v5.Pointer("true"),
        LedgerName: "<value>",
        QueryName: "<value>",
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ExecutePreparedQueryResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.V3ExecutePreparedQueryRequest](../../pkg/models/operations/v3executepreparedqueryrequest.md)   | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `security`                                                                                                 | [operations.V3ExecutePreparedQuerySecurity](../../pkg/models/operations/v3executepreparedquerysecurity.md) | :heavy_check_mark:                                                                                         | The security requirements to use for the request.                                                          |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.V3ExecutePreparedQueryResponse](../../pkg/models/operations/v3executepreparedqueryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetAccount

Retrieves account information including metadata for a specific account address. Requires `ledger:read` scope.

### Example Usage

<!-- UsageSnippet language="go" operationID="v3GetAccount_ledgerV3" method="get" path="/api/ledger/v3/{ledgerName}/accounts/{address}" -->
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

    res, err := s.Ledger.V3.GetAccount(ctx, operations.V3GetAccountLedgerV3Request{
        FormanceBigintAsString: v5.Pointer("true"),
        Address: "users:001",
        LedgerName: "<value>",
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.GetAccountResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.V3GetAccountLedgerV3Request](../../pkg/models/operations/v3getaccountledgerv3request.md)   | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `security`                                                                                             | [operations.V3GetAccountLedgerV3Security](../../pkg/models/operations/v3getaccountledgerv3security.md) | :heavy_check_mark:                                                                                     | The security requirements to use for the request.                                                      |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.V3GetAccountLedgerV3Response](../../pkg/models/operations/v3getaccountledgerv3response.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetAccountType

Retrieves a specific account type by name from a ledger. Requires `ledger:ReadAccount` scope.

### Example Usage

<!-- UsageSnippet language="go" operationID="v3GetAccountType" method="get" path="/api/ledger/v3/{ledgerName}/account-types/{typeName}" -->
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

    res, err := s.Ledger.V3.GetAccountType(ctx, operations.V3GetAccountTypeRequest{
        LedgerName: "<value>",
        TypeName: "<value>",
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.GetAccountTypeResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.V3GetAccountTypeRequest](../../pkg/models/operations/v3getaccounttyperequest.md)   | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `security`                                                                                     | [operations.V3GetAccountTypeSecurity](../../pkg/models/operations/v3getaccounttypesecurity.md) | :heavy_check_mark:                                                                             | The security requirements to use for the request.                                              |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.V3GetAccountTypeResponse](../../pkg/models/operations/v3getaccounttyperesponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetAuditEntry

Returns a single audit entry by its global sequence, with per-order
`items` populated (the list endpoint omits `items` for compactness).
Requires `ledger:ReadAudit` scope.


### Example Usage

<!-- UsageSnippet language="go" operationID="v3GetAuditEntry" method="get" path="/api/ledger/v3/_/audit-entries/{sequence}" -->
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

    res, err := s.Ledger.V3.GetAuditEntry(ctx, operations.V3GetAuditEntryRequest{
        Sequence: 98403,
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.V3GetAuditEntryRequest](../../pkg/models/operations/v3getauditentryrequest.md)   | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `security`                                                                                   | [operations.V3GetAuditEntrySecurity](../../pkg/models/operations/v3getauditentrysecurity.md) | :heavy_check_mark:                                                                           | The security requirements to use for the request.                                            |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.V3GetAuditEntryResponse](../../pkg/models/operations/v3getauditentryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetBucketIndex

Returns the `Index` registry entry for a bucket-scoped index
(`Ledger == ""`, e.g. audit indexes). Ledger-scoped lookups go
through `GET /v3/{ledgerName}/indexes/{canonicalId}` — the two
keyspaces are distinct so a bucket-scoped lookup with the same
canonical will NotFound if only the ledger-scoped entry exists.
Requires `ledger:ReadOps` scope.


### Example Usage

<!-- UsageSnippet language="go" operationID="v3GetBucketIndex" method="get" path="/api/ledger/v3/_/indexes/{canonicalId}" -->
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

    res, err := s.Ledger.V3.GetBucketIndex(ctx, operations.V3GetBucketIndexRequest{
        CanonicalID: "<id>",
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.V3GetBucketIndexRequest](../../pkg/models/operations/v3getbucketindexrequest.md)   | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `security`                                                                                     | [operations.V3GetBucketIndexSecurity](../../pkg/models/operations/v3getbucketindexsecurity.md) | :heavy_check_mark:                                                                             | The security requirements to use for the request.                                              |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.V3GetBucketIndexResponse](../../pkg/models/operations/v3getbucketindexresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetBucketIndexEntryStatus

Returns the `IndexEntry` (registry entry + backfill cursor +
per-replica IndexVersionState) for a bucket-scoped index.
Requires `ledger:ReadOps` scope.


### Example Usage

<!-- UsageSnippet language="go" operationID="v3GetBucketIndexEntryStatus" method="get" path="/api/ledger/v3/_/indexes/{canonicalId}/status" -->
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

    res, err := s.Ledger.V3.GetBucketIndexEntryStatus(ctx, operations.V3GetBucketIndexEntryStatusRequest{
        CanonicalID: "<id>",
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.V3GetBucketIndexEntryStatusRequest](../../pkg/models/operations/v3getbucketindexentrystatusrequest.md)   | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `security`                                                                                                           | [operations.V3GetBucketIndexEntryStatusSecurity](../../pkg/models/operations/v3getbucketindexentrystatussecurity.md) | :heavy_check_mark:                                                                                                   | The security requirements to use for the request.                                                                    |
| `opts`                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |

### Response

**[*operations.V3GetBucketIndexEntryStatusResponse](../../pkg/models/operations/v3getbucketindexentrystatusresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetEventsSinks

Returns every event sink configured on the cluster. Requires
`ledger:ReadOps` scope.


### Example Usage

<!-- UsageSnippet language="go" operationID="v3GetEventsSinks" method="get" path="/api/ledger/v3/_/events-sinks" -->
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

    res, err := s.Ledger.V3.GetEventsSinks(ctx, operations.V3GetEventsSinksRequest{}, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.V3GetEventsSinksRequest](../../pkg/models/operations/v3geteventssinksrequest.md)   | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `security`                                                                                     | [operations.V3GetEventsSinksSecurity](../../pkg/models/operations/v3geteventssinkssecurity.md) | :heavy_check_mark:                                                                             | The security requirements to use for the request.                                              |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.V3GetEventsSinksResponse](../../pkg/models/operations/v3geteventssinksresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetIndex

Returns the `Index` registry entry (id, ledger, createdAt,
forwardEncodingVersion) for a single index, serialized in
protobuf-JSON camelCase. Requires `ledger:ReadLedger` scope.


### Example Usage

<!-- UsageSnippet language="go" operationID="v3GetIndex" method="get" path="/api/ledger/v3/{ledgerName}/indexes/{canonicalId}" -->
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

    res, err := s.Ledger.V3.GetIndex(ctx, operations.V3GetIndexRequest{
        CanonicalID: "<id>",
        LedgerName: "<value>",
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.V3GetIndexRequest](../../pkg/models/operations/v3getindexrequest.md)   | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `security`                                                                         | [operations.V3GetIndexSecurity](../../pkg/models/operations/v3getindexsecurity.md) | :heavy_check_mark:                                                                 | The security requirements to use for the request.                                  |
| `opts`                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                       | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.V3GetIndexResponse](../../pkg/models/operations/v3getindexresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetIndexEntryStatus

Returns the `IndexEntry` for a single index — the registry entry
joined with the backfill cursor and the per-replica
IndexVersionState (current + pending version). Use this to poll
for backfill completion after a `POST /v3/{ledgerName}/indexes`.
Requires `ledger:ReadLedger` scope.


### Example Usage

<!-- UsageSnippet language="go" operationID="v3GetIndexEntryStatus" method="get" path="/api/ledger/v3/{ledgerName}/indexes/{canonicalId}/status" -->
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

    res, err := s.Ledger.V3.GetIndexEntryStatus(ctx, operations.V3GetIndexEntryStatusRequest{
        CanonicalID: "<id>",
        LedgerName: "<value>",
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.V3GetIndexEntryStatusRequest](../../pkg/models/operations/v3getindexentrystatusrequest.md)   | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `security`                                                                                               | [operations.V3GetIndexEntryStatusSecurity](../../pkg/models/operations/v3getindexentrystatussecurity.md) | :heavy_check_mark:                                                                                       | The security requirements to use for the request.                                                        |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.V3GetIndexEntryStatusResponse](../../pkg/models/operations/v3getindexentrystatusresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetIndexStatus

Returns the aggregated index registry status: per-index backfill cursor,
per-replica IndexVersionState (current/pending), and cluster-wide
counters (LastIndexedSequence, LastLogSequence, Lag, IndexFileSize).
The optional `ledger` query parameter narrows the per-index entries to
a single ledger; the aggregated counters always cover the bucket.
Requires `ledger:ReadOps` scope.


### Example Usage

<!-- UsageSnippet language="go" operationID="v3GetIndexStatus" method="get" path="/api/ledger/v3/_/indexes/status" -->
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

    res, err := s.Ledger.V3.GetIndexStatus(ctx, operations.V3GetIndexStatusRequest{}, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.V3GetIndexStatusRequest](../../pkg/models/operations/v3getindexstatusrequest.md)   | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `security`                                                                                     | [operations.V3GetIndexStatusSecurity](../../pkg/models/operations/v3getindexstatussecurity.md) | :heavy_check_mark:                                                                             | The security requirements to use for the request.                                              |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.V3GetIndexStatusResponse](../../pkg/models/operations/v3getindexstatusresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetLedger

Retrieves a ledger by its name. Requires `ledger:read` scope.

### Example Usage

<!-- UsageSnippet language="go" operationID="v3GetLedger" method="get" path="/api/ledger/v3/{ledgerName}" -->
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

    res, err := s.Ledger.V3.GetLedger(ctx, operations.V3GetLedgerRequest{
        LedgerName: "<value>",
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.GetLedgerResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.V3GetLedgerRequest](../../pkg/models/operations/v3getledgerrequest.md)   | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `security`                                                                           | [operations.V3GetLedgerSecurity](../../pkg/models/operations/v3getledgersecurity.md) | :heavy_check_mark:                                                                   | The security requirements to use for the request.                                    |
| `opts`                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                         | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.V3GetLedgerResponse](../../pkg/models/operations/v3getledgerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetLedgerStats

Returns aggregate usage statistics (transaction, volume, reference, posting, log, revert and Numscript-execution counts, plus ephemeral-evicted and transient-used volume tallies) for a ledger. Requires `ledger:read` scope. On live reads, `transactionCount` and `logCount` are FSM-backed. The seven usage-derived counters (`postingCount`, `revertCount`, `numscriptExecutionCount`, `referenceCount`, `ephemeralEvictedCount`, `transientUsedCount`, `volumeCount`) are materialised asynchronously in the serving replica's usagestore and may lag committed state even after the request's ReadIndex barrier. Note: for checkpoint-scoped (historical) reads, only `transactionCount` and `logCount` are checkpoint-consistent; the usage-derived counters (`postingCount`, `revertCount`, `numscriptExecutionCount`, `referenceCount`, `ephemeralEvictedCount`, `transientUsedCount`, `volumeCount`) are sourced from the live usagestore projection, which is not checkpoint-aware, and are returned as `0` at a checkpoint rather than a live value — treat `0` on a checkpoint read as "not available at this checkpoint".

### Example Usage

<!-- UsageSnippet language="go" operationID="v3GetLedgerStats" method="get" path="/api/ledger/v3/{ledgerName}/stats" -->
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

    res, err := s.Ledger.V3.GetLedgerStats(ctx, operations.V3GetLedgerStatsRequest{
        LedgerName: "<value>",
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.GetLedgerStatsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.V3GetLedgerStatsRequest](../../pkg/models/operations/v3getledgerstatsrequest.md)   | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `security`                                                                                     | [operations.V3GetLedgerStatsSecurity](../../pkg/models/operations/v3getledgerstatssecurity.md) | :heavy_check_mark:                                                                             | The security requirements to use for the request.                                              |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.V3GetLedgerStatsResponse](../../pkg/models/operations/v3getledgerstatsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetLog

Returns the system log entry with the given bucket-wide sequence number.
Requires `ledger:ReadOps` scope.


### Example Usage

<!-- UsageSnippet language="go" operationID="v3GetLog" method="get" path="/api/ledger/v3/_/logs/{sequence}" -->
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

    res, err := s.Ledger.V3.GetLog(ctx, operations.V3GetLogRequest{
        FormanceBigintAsString: v5.Pointer("true"),
        Sequence: 593519,
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.V3GetLogRequest](../../pkg/models/operations/v3getlogrequest.md)   | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `security`                                                                     | [operations.V3GetLogSecurity](../../pkg/models/operations/v3getlogsecurity.md) | :heavy_check_mark:                                                             | The security requirements to use for the request.                              |
| `opts`                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                   | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.V3GetLogResponse](../../pkg/models/operations/v3getlogresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetMetadataSchemaStatus

Returns the declared type of every metadata field on a ledger. Requires `ledger:read` scope.

### Example Usage

<!-- UsageSnippet language="go" operationID="v3GetMetadataSchemaStatus" method="get" path="/api/ledger/v3/{ledgerName}/metadata-schema" -->
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

    res, err := s.Ledger.V3.GetMetadataSchemaStatus(ctx, operations.V3GetMetadataSchemaStatusRequest{
        LedgerName: "<value>",
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.MetadataSchemaStatusResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.V3GetMetadataSchemaStatusRequest](../../pkg/models/operations/v3getmetadataschemastatusrequest.md)   | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `security`                                                                                                       | [operations.V3GetMetadataSchemaStatusSecurity](../../pkg/models/operations/v3getmetadataschemastatussecurity.md) | :heavy_check_mark:                                                                                               | The security requirements to use for the request.                                                                |
| `opts`                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.V3GetMetadataSchemaStatusResponse](../../pkg/models/operations/v3getmetadataschemastatusresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetNumscript

Retrieves a numscript by name from a ledger. Use the `version` query parameter to get a specific
semver version. If omitted or empty, the latest version is returned.


### Example Usage

<!-- UsageSnippet language="go" operationID="v3GetNumscript" method="get" path="/api/ledger/v3/{ledgerName}/numscripts/{name}" -->
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

    res, err := s.Ledger.V3.GetNumscript(ctx, operations.V3GetNumscriptRequest{
        LedgerName: "<value>",
        Name: "<value>",
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.GetNumscriptResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.V3GetNumscriptRequest](../../pkg/models/operations/v3getnumscriptrequest.md)   | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `security`                                                                                 | [operations.V3GetNumscriptSecurity](../../pkg/models/operations/v3getnumscriptsecurity.md) | :heavy_check_mark:                                                                         | The security requirements to use for the request.                                          |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.V3GetNumscriptResponse](../../pkg/models/operations/v3getnumscriptresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetNumscriptUsage

Returns the invocation counter and last-used timestamp for a Numscript template. Values are
materialised asynchronously by the `usagebuilder` subsystem from the FSM audit chain and may
lag the live FSM while the worker drains pending batches. A never-invoked template returns a
zero-valued response (not 404), so clients handle "never used" uniformly.


### Example Usage

<!-- UsageSnippet language="go" operationID="v3GetNumscriptUsage" method="get" path="/api/ledger/v3/{ledgerName}/numscripts/{name}/usage" -->
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

    res, err := s.Ledger.V3.GetNumscriptUsage(ctx, operations.V3GetNumscriptUsageRequest{
        LedgerName: "<value>",
        Name: "<value>",
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.GetTemplateUsageResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.V3GetNumscriptUsageRequest](../../pkg/models/operations/v3getnumscriptusagerequest.md)   | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `security`                                                                                           | [operations.V3GetNumscriptUsageSecurity](../../pkg/models/operations/v3getnumscriptusagesecurity.md) | :heavy_check_mark:                                                                                   | The security requirements to use for the request.                                                    |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.V3GetNumscriptUsageResponse](../../pkg/models/operations/v3getnumscriptusageresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetTransaction

Retrieves a transaction by its ID from the specified ledger. Requires `ledger:read` scope.

### Example Usage

<!-- UsageSnippet language="go" operationID="v3GetTransaction" method="get" path="/api/ledger/v3/{ledgerName}/transactions/{transactionId}" -->
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

    res, err := s.Ledger.V3.GetTransaction(ctx, operations.V3GetTransactionRequest{
        FormanceBigintAsString: v5.Pointer("true"),
        LedgerName: "<value>",
        TransactionID: 431503,
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.GetTransactionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.V3GetTransactionRequest](../../pkg/models/operations/v3gettransactionrequest.md)   | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `security`                                                                                     | [operations.V3GetTransactionSecurity](../../pkg/models/operations/v3gettransactionsecurity.md) | :heavy_check_mark:                                                                             | The security requirements to use for the request.                                              |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.V3GetTransactionResponse](../../pkg/models/operations/v3gettransactionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Health

Returns the health status of the node. Returns 503 if the node is not connected to the cluster.

### Example Usage

<!-- UsageSnippet language="go" operationID="v3Health" method="get" path="/api/ledger/health" -->
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

    res, err := s.Ledger.V3.Health(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.HealthResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `security`                                                                     | [operations.V3HealthSecurity](../../pkg/models/operations/v3healthsecurity.md) | :heavy_check_mark:                                                             | The security requirements to use for the request.                              |
| `opts`                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                   | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.V3HealthResponse](../../pkg/models/operations/v3healthresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Info

Returns the server build metadata (version, commit, build date, Go runtime version). Unauthenticated.

### Example Usage

<!-- UsageSnippet language="go" operationID="v3Info" method="get" path="/api/ledger/_info#ledgerV3" -->
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

    res, err := s.Ledger.V3.Info(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ServerInfo != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `security`                                                                 | [operations.V3InfoSecurity](../../pkg/models/operations/v3infosecurity.md) | :heavy_check_mark:                                                         | The security requirements to use for the request.                          |
| `opts`                                                                     | [][operations.Option](../../pkg/models/operations/option.md)               | :heavy_minus_sign:                                                         | The options for this request.                                              |

### Response

**[*operations.V3InfoResponse](../../pkg/models/operations/v3inforesponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## InspectIndex

Scans a metadata index and returns distinct values, facets (value + count),
or a summary (cardinality, min/max, existence counts). The index must be
in READY state. Only metadata indexes are inspectable (builtin indexes
return 400). Requires `ledger:read` scope.


### Example Usage

<!-- UsageSnippet language="go" operationID="v3InspectIndex" method="get" path="/api/ledger/v3/{ledgerName}/indexes/{canonicalId}/inspect" -->
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

    res, err := s.Ledger.V3.InspectIndex(ctx, operations.V3InspectIndexRequest{
        CanonicalID: "<id>",
        LedgerName: "<value>",
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        switch res.Object.Data.Type {
            case operations.V3InspectIndexDataTypeInspectSummaryResponse:
                // res.Object.Data.InspectSummaryResponse is populated
            case operations.V3InspectIndexDataTypeInspectDistinctValuesResponse:
                // res.Object.Data.InspectDistinctValuesResponse is populated
            case operations.V3InspectIndexDataTypeInspectFacetsResponse:
                // res.Object.Data.InspectFacetsResponse is populated
        }

    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.V3InspectIndexRequest](../../pkg/models/operations/v3inspectindexrequest.md)   | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `security`                                                                                 | [operations.V3InspectIndexSecurity](../../pkg/models/operations/v3inspectindexsecurity.md) | :heavy_check_mark:                                                                         | The security requirements to use for the request.                                          |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.V3InspectIndexResponse](../../pkg/models/operations/v3inspectindexresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListAccountTypes

Returns all account types defined on a ledger, sorted by name.
Requires `ledger:ReadAccount` scope.


### Example Usage

<!-- UsageSnippet language="go" operationID="v3ListAccountTypes" method="get" path="/api/ledger/v3/{ledgerName}/account-types" -->
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

    res, err := s.Ledger.V3.ListAccountTypes(ctx, operations.V3ListAccountTypesRequest{
        LedgerName: "<value>",
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ListAccountTypesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.V3ListAccountTypesRequest](../../pkg/models/operations/v3listaccounttypesrequest.md)   | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `security`                                                                                         | [operations.V3ListAccountTypesSecurity](../../pkg/models/operations/v3listaccounttypessecurity.md) | :heavy_check_mark:                                                                                 | The security requirements to use for the request.                                                  |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.V3ListAccountTypesResponse](../../pkg/models/operations/v3listaccounttypesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListAccounts

Lists accounts in a ledger in alphabetical order with cursor
pagination and an optional generic `filter` (e.g. address-prefix
selection via `filter=address ^= "users:"`). Requires `ledger:read`
scope.


### Example Usage

<!-- UsageSnippet language="go" operationID="v3ListAccounts_ledgerV3" method="get" path="/api/ledger/v3/{ledgerName}/accounts" -->
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

    res, err := s.Ledger.V3.ListAccounts(ctx, operations.V3ListAccountsLedgerV3Request{
        FormanceBigintAsString: v5.Pointer("true"),
        LedgerName: "<value>",
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.V3ListAccountsLedgerV3Request](../../pkg/models/operations/v3listaccountsledgerv3request.md)   | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `security`                                                                                                 | [operations.V3ListAccountsLedgerV3Security](../../pkg/models/operations/v3listaccountsledgerv3security.md) | :heavy_check_mark:                                                                                         | The security requirements to use for the request.                                                          |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.V3ListAccountsLedgerV3Response](../../pkg/models/operations/v3listaccountsledgerv3response.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListAllLedgers

Lists the ledgers in the cluster, paged by name (ascending unless
`reverse=true`). Requires `ledger:read` scope.


### Example Usage

<!-- UsageSnippet language="go" operationID="v3ListAllLedgers" method="get" path="/api/ledger/v3/" -->
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

    res, err := s.Ledger.V3.ListAllLedgers(ctx, operations.V3ListAllLedgersRequest{}, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ListAllLedgersResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.V3ListAllLedgersRequest](../../pkg/models/operations/v3listallledgersrequest.md)   | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `security`                                                                                     | [operations.V3ListAllLedgersSecurity](../../pkg/models/operations/v3listallledgerssecurity.md) | :heavy_check_mark:                                                                             | The security requirements to use for the request.                                              |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.V3ListAllLedgersResponse](../../pkg/models/operations/v3listallledgersresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListAuditEntries

Lists audit trail entries (one per proposal, success or failure) across
the whole bucket, paged by audit sequence (ascending unless
`reverse=true`). Audit reads are
not ledger-scoped in the path — a single proposal can touch several
ledgers — so ledger scope, outcome and other dimensions are expressed
through the `filter` query parameter.

This endpoint always performs a live read and does not expose pinned
checkpoint reads. Filters containing indexed audit conditions wait for
the audit projection to certify the request's fixed main-store horizon;
matching entries are loaded from that same main snapshot. Unfiltered
reads and sequence-only bounds scan the audit zone directly without
waiting for the audit projection.

Requires `ledger:ReadAudit` scope.


### Example Usage

<!-- UsageSnippet language="go" operationID="v3ListAuditEntries" method="get" path="/api/ledger/v3/_/audit-entries" -->
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

    res, err := s.Ledger.V3.ListAuditEntries(ctx, operations.V3ListAuditEntriesRequest{}, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.V3ListAuditEntriesRequest](../../pkg/models/operations/v3listauditentriesrequest.md)   | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `security`                                                                                         | [operations.V3ListAuditEntriesSecurity](../../pkg/models/operations/v3listauditentriessecurity.md) | :heavy_check_mark:                                                                                 | The security requirements to use for the request.                                                  |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.V3ListAuditEntriesResponse](../../pkg/models/operations/v3listauditentriesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListBucketIndexes

Streams the bucket-scoped index registry. The `scope` query parameter
selects the flavor: `all` (default, matches the ListIndexesRequest
proto default) streams every entry across ledgers, skipping orphan
entries whose owning ledger no longer exists; `bucket` surfaces only
bucket-scoped entries (empty ledger, e.g. audit indexes). Use
`GET /v3/{ledgerName}/indexes` for the per-ledger view. Requires
`ledger:ReadOps` scope.


### Example Usage

<!-- UsageSnippet language="go" operationID="v3ListBucketIndexes" method="get" path="/api/ledger/v3/_/indexes" -->
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

    res, err := s.Ledger.V3.ListBucketIndexes(ctx, operations.V3ListBucketIndexesRequest{}, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.V3ListBucketIndexesRequest](../../pkg/models/operations/v3listbucketindexesrequest.md)   | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `security`                                                                                           | [operations.V3ListBucketIndexesSecurity](../../pkg/models/operations/v3listbucketindexessecurity.md) | :heavy_check_mark:                                                                                   | The security requirements to use for the request.                                                    |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.V3ListBucketIndexesResponse](../../pkg/models/operations/v3listbucketindexesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListLedgerIndexes

Streams the indexes registered on a specific ledger. Requires
`ledger:ReadLedger` scope.


### Example Usage

<!-- UsageSnippet language="go" operationID="v3ListLedgerIndexes" method="get" path="/api/ledger/v3/{ledgerName}/indexes" -->
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

    res, err := s.Ledger.V3.ListLedgerIndexes(ctx, operations.V3ListLedgerIndexesRequest{
        LedgerName: "<value>",
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.V3ListLedgerIndexesRequest](../../pkg/models/operations/v3listledgerindexesrequest.md)   | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `security`                                                                                           | [operations.V3ListLedgerIndexesSecurity](../../pkg/models/operations/v3listledgerindexessecurity.md) | :heavy_check_mark:                                                                                   | The security requirements to use for the request.                                                    |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.V3ListLedgerIndexesResponse](../../pkg/models/operations/v3listledgerindexesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListLedgerLogs

Lists log entries for a specific ledger, paged by ledger-local log ID
(ascending unless `reverse=true`).
Requires the builtin `LOG` index to be enabled on the ledger via `CreateIndex`.
Requires `ledger:read` scope.


### Example Usage

<!-- UsageSnippet language="go" operationID="v3ListLedgerLogs" method="get" path="/api/ledger/v3/{ledgerName}/logs" -->
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

    res, err := s.Ledger.V3.ListLedgerLogs(ctx, operations.V3ListLedgerLogsRequest{
        FormanceBigintAsString: v5.Pointer("true"),
        LedgerName: "<value>",
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.V3ListLedgerLogsRequest](../../pkg/models/operations/v3listledgerlogsrequest.md)   | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `security`                                                                                     | [operations.V3ListLedgerLogsSecurity](../../pkg/models/operations/v3listledgerlogssecurity.md) | :heavy_check_mark:                                                                             | The security requirements to use for the request.                                              |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.V3ListLedgerLogsResponse](../../pkg/models/operations/v3listledgerlogsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListLedgerTransactions

Lists transactions for a specific ledger. The default order is
newest-first (descending transaction id); pass `reverse=true` to
get oldest-first (ascending). Supports cursor pagination (`pageSize`,
`cursor`), an optional date range (`startDate`/`endDate`),
and the generic `filter` parameter (e.g. reference selection via
`filter={"$match":{"reference":"..."}}`, URL-encoded). Requires
`ledger:ReadTransaction` scope.

This is a live read, linearizable by default. `X-Consistency: stale`
reads the receiving node's local state, which may lag the leader. Every
projection used by the filter is automatically aligned. Unlike the
gRPC API, it does not accept the `checkpointId` option to pin the read
to a specific historical snapshot — use gRPC for checkpoint-pinned
reads. (Tracked follow-up, same carve-out as the audit reads.)


### Example Usage

<!-- UsageSnippet language="go" operationID="v3ListLedgerTransactions" method="get" path="/api/ledger/v3/{ledgerName}/transactions" -->
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

    res, err := s.Ledger.V3.ListLedgerTransactions(ctx, operations.V3ListLedgerTransactionsRequest{
        FormanceBigintAsString: v5.Pointer("true"),
        LedgerName: "<value>",
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.V3ListLedgerTransactionsRequest](../../pkg/models/operations/v3listledgertransactionsrequest.md)   | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `security`                                                                                                     | [operations.V3ListLedgerTransactionsSecurity](../../pkg/models/operations/v3listledgertransactionssecurity.md) | :heavy_check_mark:                                                                                             | The security requirements to use for the request.                                                              |
| `opts`                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.V3ListLedgerTransactionsResponse](../../pkg/models/operations/v3listledgertransactionsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListNumscriptVersions

Returns the numscript's history: its current latest (greatest stored
semver) and every stored version, highest-first.


### Example Usage

<!-- UsageSnippet language="go" operationID="v3ListNumscriptVersions" method="get" path="/api/ledger/v3/{ledgerName}/numscripts/{name}/versions" -->
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

    res, err := s.Ledger.V3.ListNumscriptVersions(ctx, operations.V3ListNumscriptVersionsRequest{
        LedgerName: "<value>",
        Name: "<value>",
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ListNumscriptVersionsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.V3ListNumscriptVersionsRequest](../../pkg/models/operations/v3listnumscriptversionsrequest.md)   | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `security`                                                                                                   | [operations.V3ListNumscriptVersionsSecurity](../../pkg/models/operations/v3listnumscriptversionssecurity.md) | :heavy_check_mark:                                                                                           | The security requirements to use for the request.                                                            |
| `opts`                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.V3ListNumscriptVersionsResponse](../../pkg/models/operations/v3listnumscriptversionsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListNumscripts

Lists every numscript of a ledger at its greatest stored version, paged
by name (ascending unless `reverse=true`).


### Example Usage

<!-- UsageSnippet language="go" operationID="v3ListNumscripts" method="get" path="/api/ledger/v3/{ledgerName}/numscripts" -->
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

    res, err := s.Ledger.V3.ListNumscripts(ctx, operations.V3ListNumscriptsRequest{
        LedgerName: "<value>",
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ListNumscriptsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.V3ListNumscriptsRequest](../../pkg/models/operations/v3listnumscriptsrequest.md)   | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `security`                                                                                     | [operations.V3ListNumscriptsSecurity](../../pkg/models/operations/v3listnumscriptssecurity.md) | :heavy_check_mark:                                                                             | The security requirements to use for the request.                                              |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.V3ListNumscriptsResponse](../../pkg/models/operations/v3listnumscriptsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListPreparedQueries

Returns a list of all prepared queries for a ledger. Requires `ledger:ReadQuery` scope.

### Example Usage

<!-- UsageSnippet language="go" operationID="v3ListPreparedQueries" method="get" path="/api/ledger/v3/{ledgerName}/prepared-queries" -->
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

    res, err := s.Ledger.V3.ListPreparedQueries(ctx, operations.V3ListPreparedQueriesRequest{
        LedgerName: "<value>",
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ListPreparedQueriesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.V3ListPreparedQueriesRequest](../../pkg/models/operations/v3listpreparedqueriesrequest.md)   | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `security`                                                                                               | [operations.V3ListPreparedQueriesSecurity](../../pkg/models/operations/v3listpreparedqueriessecurity.md) | :heavy_check_mark:                                                                                       | The security requirements to use for the request.                                                        |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.V3ListPreparedQueriesResponse](../../pkg/models/operations/v3listpreparedqueriesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListSigningKeys

Lists the Ed25519 signing keys registered on the cluster, paged by key
id (ascending unless `reverse=true`). Requires `ledger:ReadOps` scope.

This route performs a live read, linearizable by default.
`X-Consistency: stale` reads the receiving node's local state, which
may lag the leader. Signing-key reads are live-only on both transports.


### Example Usage

<!-- UsageSnippet language="go" operationID="v3ListSigningKeys" method="get" path="/api/ledger/v3/_/signing-keys" -->
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

    res, err := s.Ledger.V3.ListSigningKeys(ctx, operations.V3ListSigningKeysRequest{}, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.V3ListSigningKeysRequest](../../pkg/models/operations/v3listsigningkeysrequest.md)   | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `security`                                                                                       | [operations.V3ListSigningKeysSecurity](../../pkg/models/operations/v3listsigningkeyssecurity.md) | :heavy_check_mark:                                                                               | The security requirements to use for the request.                                                |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.V3ListSigningKeysResponse](../../pkg/models/operations/v3listsigningkeysresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Livez

Returns 200 as long as the process is alive. Suitable for Kubernetes liveness probes.

### Example Usage

<!-- UsageSnippet language="go" operationID="v3Livez" method="get" path="/api/ledger/livez" -->
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

    res, err := s.Ledger.V3.Livez(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.HealthResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `security`                                                                   | [operations.V3LivezSecurity](../../pkg/models/operations/v3livezsecurity.md) | :heavy_check_mark:                                                           | The security requirements to use for the request.                            |
| `opts`                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                 | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.V3LivezResponse](../../pkg/models/operations/v3livezresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PromoteLedger

Promotes a mirror-mode ledger to normal mode. After promotion, the mirror sync
worker is stopped and the ledger accepts direct write operations. Requires `ledger:write` scope.


### Example Usage

<!-- UsageSnippet language="go" operationID="v3PromoteLedger" method="post" path="/api/ledger/v3/{ledgerName}/promote" -->
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

    res, err := s.Ledger.V3.PromoteLedger(ctx, operations.V3PromoteLedgerRequest{
        LedgerName: "<value>",
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.PromoteLedgerResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.V3PromoteLedgerRequest](../../pkg/models/operations/v3promoteledgerrequest.md)   | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `security`                                                                                   | [operations.V3PromoteLedgerSecurity](../../pkg/models/operations/v3promoteledgersecurity.md) | :heavy_check_mark:                                                                           | The security requirements to use for the request.                                            |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.V3PromoteLedgerResponse](../../pkg/models/operations/v3promoteledgerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Readyz

Returns 200 once the local Raft loop has started, regardless of whether
a leader has been elected. Suitable for Kubernetes readiness probes and
StatefulSet `OrderedReady` ordering: it is intentionally permissive so
peer pods can be brought up during a cold start where quorum cannot
form until they exist. Use `/clusterz` for the stricter
cluster-availability signal.


### Example Usage

<!-- UsageSnippet language="go" operationID="v3Readyz" method="get" path="/api/ledger/readyz" -->
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

    res, err := s.Ledger.V3.Readyz(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.HealthResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `security`                                                                     | [operations.V3ReadyzSecurity](../../pkg/models/operations/v3readyzsecurity.md) | :heavy_check_mark:                                                             | The security requirements to use for the request.                              |
| `opts`                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                   | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.V3ReadyzResponse](../../pkg/models/operations/v3readyzresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveAccountType

Removes an account type from a ledger's chart of accounts. Requires `ledger:WriteMetadata` scope.

### Example Usage

<!-- UsageSnippet language="go" operationID="v3RemoveAccountType" method="delete" path="/api/ledger/v3/{ledgerName}/account-types/{typeName}" -->
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

    res, err := s.Ledger.V3.RemoveAccountType(ctx, operations.V3RemoveAccountTypeRequest{
        LedgerName: "<value>",
        TypeName: "<value>",
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.V3RemoveAccountTypeRequest](../../pkg/models/operations/v3removeaccounttyperequest.md)   | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `security`                                                                                           | [operations.V3RemoveAccountTypeSecurity](../../pkg/models/operations/v3removeaccounttypesecurity.md) | :heavy_check_mark:                                                                                   | The security requirements to use for the request.                                                    |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.V3RemoveAccountTypeResponse](../../pkg/models/operations/v3removeaccounttyperesponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveEventsSink

Removes a named event sink from the cluster. Requires `ledger:WriteOps`
scope. A non-empty controllerId must match the sink's controllerId;
a mismatch returns `SINK_CONTROLLER_MISMATCH` (409). An omitted or empty
controllerId removes the sink without an ownership precondition.
A missing sink returns `SINK_NOT_FOUND` (404). Malformed query encoding
or repeated controllerId parameters return `INVALID_REQUEST` (400)
before Apply, instead of becoming unconditional removal.


### Example Usage

<!-- UsageSnippet language="go" operationID="v3RemoveEventsSink" method="delete" path="/api/ledger/v3/_/events-sinks/{sinkName}" -->
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

    res, err := s.Ledger.V3.RemoveEventsSink(ctx, operations.V3RemoveEventsSinkRequest{
        SinkName: "<value>",
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.V3RemoveEventsSinkRequest](../../pkg/models/operations/v3removeeventssinkrequest.md)   | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `security`                                                                                         | [operations.V3RemoveEventsSinkSecurity](../../pkg/models/operations/v3removeeventssinksecurity.md) | :heavy_check_mark:                                                                                 | The security requirements to use for the request.                                                  |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.V3RemoveEventsSinkResponse](../../pkg/models/operations/v3removeeventssinkresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveMetadataFieldType

Removes the type declaration for a metadata key on a ledger.
Existing metadata values remain unchanged and any attached index is dropped. Requires `ledger:WriteMetadata` scope (included in the default `ledger:write` mapping).


### Example Usage

<!-- UsageSnippet language="go" operationID="v3RemoveMetadataFieldType" method="delete" path="/api/ledger/v3/{ledgerName}/metadata-schema/{targetType}/{key}" -->
```go
package main

import(
	"context"
	"github.com/formancehq/formance-sdk-go/v5"
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/ledgerv3"
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := v5.New()

    res, err := s.Ledger.V3.RemoveMetadataFieldType(ctx, operations.V3RemoveMetadataFieldTypeRequest{
        Key: "formance.com/reviewed",
        LedgerName: "<value>",
        TargetType: ledgerv3.TargetTypeLedger,
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.V3RemoveMetadataFieldTypeRequest](../../pkg/models/operations/v3removemetadatafieldtyperequest.md)   | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `security`                                                                                                       | [operations.V3RemoveMetadataFieldTypeSecurity](../../pkg/models/operations/v3removemetadatafieldtypesecurity.md) | :heavy_check_mark:                                                                                               | The security requirements to use for the request.                                                                |
| `opts`                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.V3RemoveMetadataFieldTypeResponse](../../pkg/models/operations/v3removemetadatafieldtyperesponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RevertTransaction

Reverts a transaction by creating a reverse transaction that undoes its effects. Requires `ledger:write` scope.
The body is optional; an absent or empty body uses default options. Supplied
options and metadata are decoded for both known-length and HTTP/1.1 chunked
requests. Malformed or truncated JSON, extra values, trailing non-whitespace
and options of the wrong JSON type (for example `"atEffectiveDate": "true"`
or a non-object `metadata`) return 400 before applying the revert. A JSON
`null` option uses its default. The complete body, including trailing
whitespace, counts towards the 4 MiB limit; exceeding it returns 413.
Signed requests use `application/json` (including MIME parameters) with
`request` and standard-base64 `payload`. Either signing header, even if
empty, selects this wrapper; with neither header present, the ordinary REST body applies.
Runtime headers choose the representation from the documented anyOf.
The signed protobuf batch must
equal the route/body/Idempotency-Key-derived batch in full. Supply exactly
one nonempty Formance-Signing-Key-Id and Formance-Signature header; the
latter is a standard-base64 Ed25519 signature of 64 decoded bytes. Missing,
empty or repeated signing headers, invalid base64 and wrong signature
length return 400. Signing headers with a non-JSON or missing content type return 415.
Signature enforcement
returns 401 `UNAUTHENTICATED` when required but absent and 403
`PERMISSION_DENIED` for invalid, unknown/revoked-key or mismatched
signatures. Malformed signed envelopes return 400. The complete
wrapper counts toward the 4 MiB body limit. JWT scopes still apply.


### Example Usage

<!-- UsageSnippet language="go" operationID="v3RevertTransaction" method="post" path="/api/ledger/v3/{ledgerName}/transactions/{transactionId}/revert" -->
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

    res, err := s.Ledger.V3.RevertTransaction(ctx, operations.V3RevertTransactionRequest{
        FormanceBigintAsString: v5.Pointer("true"),
        LedgerName: "<value>",
        TransactionID: 541831,
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.RevertedTransactionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.V3RevertTransactionRequest](../../pkg/models/operations/v3reverttransactionrequest.md)   | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `security`                                                                                           | [operations.V3RevertTransactionSecurity](../../pkg/models/operations/v3reverttransactionsecurity.md) | :heavy_check_mark:                                                                                   | The security requirements to use for the request.                                                    |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.V3RevertTransactionResponse](../../pkg/models/operations/v3reverttransactionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## SaveAccountMetadata

Saves metadata for a specific account in the ledger. Requires `ledger:WriteMetadata` scope (included in the default `ledger:write` mapping).

### Example Usage

<!-- UsageSnippet language="go" operationID="v3SaveAccountMetadata" method="post" path="/api/ledger/v3/{ledgerName}/accounts/{address}/metadata" -->
```go
package main

import(
	"context"
	"github.com/formancehq/formance-sdk-go/v5"
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/ledgerv3"
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := v5.New()

    res, err := s.Ledger.V3.SaveAccountMetadata(ctx, operations.V3SaveAccountMetadataRequest{
        RequestBody: map[string]*ledgerv3.MetadataValue{
            "account_type": v5.Pointer(ledgerv3.CreateMetadataValueStr(
                "asset",
            )),
            "label": v5.Pointer(ledgerv3.CreateMetadataValueStr(
                "Bank Account",
            )),
            "active": v5.Pointer(ledgerv3.CreateMetadataValueBoolean(
                true,
            )),
            "balance_threshold": v5.Pointer(ledgerv3.CreateMetadataValueInteger(
                1000,
            )),
        },
        Address: "users:001",
        LedgerName: "<value>",
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.V3SaveAccountMetadataRequest](../../pkg/models/operations/v3saveaccountmetadatarequest.md)   | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `security`                                                                                               | [operations.V3SaveAccountMetadataSecurity](../../pkg/models/operations/v3saveaccountmetadatasecurity.md) | :heavy_check_mark:                                                                                       | The security requirements to use for the request.                                                        |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.V3SaveAccountMetadataResponse](../../pkg/models/operations/v3saveaccountmetadataresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## SaveLedgerMetadata

Saves metadata key-value pairs on the ledger itself. Requires `ledger:WriteMetadata`; `ledger:write` grants it under the default scope mapping. Custom scope mappings can differ.

### Example Usage

<!-- UsageSnippet language="go" operationID="v3SaveLedgerMetadata" method="post" path="/api/ledger/v3/{ledgerName}/metadata" -->
```go
package main

import(
	"context"
	"github.com/formancehq/formance-sdk-go/v5"
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/ledgerv3"
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := v5.New()

    res, err := s.Ledger.V3.SaveLedgerMetadata(ctx, operations.V3SaveLedgerMetadataRequest{
        RequestBody: map[string]*ledgerv3.MetadataValue{
            "environment": v5.Pointer(ledgerv3.CreateMetadataValueStr(
                "production",
            )),
            "team": v5.Pointer(ledgerv3.CreateMetadataValueStr(
                "payments",
            )),
            "region": v5.Pointer(ledgerv3.CreateMetadataValueStr(
                "eu-west-1",
            )),
        },
        LedgerName: "<value>",
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.V3SaveLedgerMetadataRequest](../../pkg/models/operations/v3saveledgermetadatarequest.md)   | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `security`                                                                                             | [operations.V3SaveLedgerMetadataSecurity](../../pkg/models/operations/v3saveledgermetadatasecurity.md) | :heavy_check_mark:                                                                                     | The security requirements to use for the request.                                                      |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.V3SaveLedgerMetadataResponse](../../pkg/models/operations/v3saveledgermetadataresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## SaveNumscript

Appends an immutable version to a ledger's numscript library. The `version`
field must be an explicit full semver (e.g. "1.0.0"); it creates a new
version and fails if that version already exists. The per-name latest is
always the greatest stored semver.


### Example Usage

<!-- UsageSnippet language="go" operationID="v3SaveNumscript" method="put" path="/api/ledger/v3/{ledgerName}/numscripts/{name}" -->
```go
package main

import(
	"context"
	"github.com/formancehq/formance-sdk-go/v5"
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/ledgerv3"
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := v5.New()

    res, err := s.Ledger.V3.SaveNumscript(ctx, operations.V3SaveNumscriptRequest{
        SaveNumscriptRequest: ledgerv3.SaveNumscriptRequest{
            Content: "<value>",
            Version: "1.0.0",
        },
        LedgerName: "<value>",
        Name: "<value>",
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.GetNumscriptResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.V3SaveNumscriptRequest](../../pkg/models/operations/v3savenumscriptrequest.md)   | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `security`                                                                                   | [operations.V3SaveNumscriptSecurity](../../pkg/models/operations/v3savenumscriptsecurity.md) | :heavy_check_mark:                                                                           | The security requirements to use for the request.                                            |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.V3SaveNumscriptResponse](../../pkg/models/operations/v3savenumscriptresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## SaveTransactionMetadata

Saves metadata for a specific transaction in the ledger. Requires `ledger:WriteMetadata` scope (included in the default `ledger:write` mapping).
Signed requests use `application/json` (including MIME parameters) with
`request` and standard-base64 `payload`. Either signing header, even if
empty, selects this wrapper; with neither header present, the ordinary REST body applies.
Runtime headers choose the representation from the documented anyOf.
The signed protobuf batch must
equal the route/body/Idempotency-Key-derived batch in full. Supply exactly
one nonempty Formance-Signing-Key-Id and Formance-Signature header; the
latter is a standard-base64 Ed25519 signature of 64 decoded bytes. Missing,
empty or repeated signing headers, invalid base64 and wrong signature
length return 400. Signing headers with a non-JSON or missing content type return 415.
Signature enforcement
returns 401 `UNAUTHENTICATED` when required but absent and 403
`PERMISSION_DENIED` for invalid, unknown/revoked-key or mismatched
signatures. Malformed signed envelopes return 400. The complete
wrapper counts toward the 4 MiB body limit. JWT scopes still apply.


### Example Usage

<!-- UsageSnippet language="go" operationID="v3SaveTransactionMetadata" method="post" path="/api/ledger/v3/{ledgerName}/transactions/{transactionId}/metadata" -->
```go
package main

import(
	"context"
	"github.com/formancehq/formance-sdk-go/v5"
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/ledgerv3"
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := v5.New()

    res, err := s.Ledger.V3.SaveTransactionMetadata(ctx, operations.V3SaveTransactionMetadataRequest{
        RequestBody: operations.CreateV3SaveTransactionMetadataRequestBodyMapOfMetadataValue(
            map[string]*ledgerv3.MetadataValue{
                "reason": v5.Pointer(ledgerv3.CreateMetadataValueStr(
                    "correction",
                )),
                "source": v5.Pointer(ledgerv3.CreateMetadataValueStr(
                    "support",
                )),
                "priority": v5.Pointer(ledgerv3.CreateMetadataValueInteger(
                    1,
                )),
            },
        ),
        LedgerName: "<value>",
        TransactionID: 736514,
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.V3SaveTransactionMetadataRequest](../../pkg/models/operations/v3savetransactionmetadatarequest.md)   | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `security`                                                                                                       | [operations.V3SaveTransactionMetadataSecurity](../../pkg/models/operations/v3savetransactionmetadatasecurity.md) | :heavy_check_mark:                                                                                               | The security requirements to use for the request.                                                                |
| `opts`                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.V3SaveTransactionMetadataResponse](../../pkg/models/operations/v3savetransactionmetadataresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## SetDefaultEnforcementMode

Sets the default enforcement mode for accounts that do not match any defined
account type. Must be STRICT or AUDIT. STRICT rejects an order that uses an
unmatched account with `400 ACCOUNT_NOT_MATCHING_TYPE`. AUDIT accepts the
order and does not report the unmatched account: the response carries no
warning. Requires `ledger:WriteMetadata` scope.


### Example Usage

<!-- UsageSnippet language="go" operationID="v3SetDefaultEnforcementMode" method="put" path="/api/ledger/v3/{ledgerName}/account-types/default-enforcement-mode" -->
```go
package main

import(
	"context"
	"github.com/formancehq/formance-sdk-go/v5"
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/ledgerv3"
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := v5.New()

    res, err := s.Ledger.V3.SetDefaultEnforcementMode(ctx, operations.V3SetDefaultEnforcementModeRequest{
        SetDefaultEnforcementModeRequest: ledgerv3.SetDefaultEnforcementModeRequest{
            EnforcementMode: ledgerv3.EnforcementModeAudit,
        },
        LedgerName: "<value>",
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.V3SetDefaultEnforcementModeRequest](../../pkg/models/operations/v3setdefaultenforcementmoderequest.md)   | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `security`                                                                                                           | [operations.V3SetDefaultEnforcementModeSecurity](../../pkg/models/operations/v3setdefaultenforcementmodesecurity.md) | :heavy_check_mark:                                                                                                   | The security requirements to use for the request.                                                                    |
| `opts`                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |

### Response

**[*operations.V3SetDefaultEnforcementModeResponse](../../pkg/models/operations/v3setdefaultenforcementmoderesponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## SetMetadataFieldType

Declares or changes the type used to encode a metadata key in an attached query index.
Metadata values are neither validated nor converted by this declaration, and existing primary values remain unchanged.
Any attached index is rewritten asynchronously. Requires `ledger:WriteMetadata` scope (included in the default `ledger:write` mapping).


### Example Usage

<!-- UsageSnippet language="go" operationID="v3SetMetadataFieldType" method="put" path="/api/ledger/v3/{ledgerName}/metadata-schema/{targetType}/{key}" -->
```go
package main

import(
	"context"
	"github.com/formancehq/formance-sdk-go/v5"
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/ledgerv3"
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := v5.New()

    res, err := s.Ledger.V3.SetMetadataFieldType(ctx, operations.V3SetMetadataFieldTypeRequest{
        SetMetadataTypeRequest: ledgerv3.SetMetadataTypeRequest{
            Type: ledgerv3.SetMetadataTypeRequestTypeBool,
        },
        Key: "formance.com/reviewed",
        LedgerName: "<value>",
        TargetType: ledgerv3.TargetTypeLedger,
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.V3SetMetadataFieldTypeRequest](../../pkg/models/operations/v3setmetadatafieldtyperequest.md)   | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `security`                                                                                                 | [operations.V3SetMetadataFieldTypeSecurity](../../pkg/models/operations/v3setmetadatafieldtypesecurity.md) | :heavy_check_mark:                                                                                         | The security requirements to use for the request.                                                          |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.V3SetMetadataFieldTypeResponse](../../pkg/models/operations/v3setmetadatafieldtyperesponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdatePreparedQuery

Updates the filter of an existing prepared query. Requires `ledger:WriteQuery` scope.


### Example Usage

<!-- UsageSnippet language="go" operationID="v3UpdatePreparedQuery" method="put" path="/api/ledger/v3/{ledgerName}/prepared-queries/{queryName}" -->
```go
package main

import(
	"context"
	"github.com/formancehq/formance-sdk-go/v5"
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/ledgerv3"
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := v5.New()

    res, err := s.Ledger.V3.UpdatePreparedQuery(ctx, operations.V3UpdatePreparedQueryRequest{
        UpdatePreparedQueryRequest: ledgerv3.UpdatePreparedQueryRequest{},
        LedgerName: "<value>",
        QueryName: "<value>",
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.V3UpdatePreparedQueryRequest](../../pkg/models/operations/v3updatepreparedqueryrequest.md)   | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `security`                                                                                               | [operations.V3UpdatePreparedQuerySecurity](../../pkg/models/operations/v3updatepreparedquerysecurity.md) | :heavy_check_mark:                                                                                       | The security requirements to use for the request.                                                        |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.V3UpdatePreparedQueryResponse](../../pkg/models/operations/v3updatepreparedqueryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |