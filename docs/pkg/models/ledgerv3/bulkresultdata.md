# BulkResultData

Operation result data. Shape depends on the originating action:
  - `CREATE_TRANSACTION` success: a `TransactionResponse` — the
    created Transaction object.
  - Any skippable action skip (when the item opted into
    `skippableReasons` and the FSM matched one): an
    `OrderSkippedResponse` with `skipped: true`, `reason`, and
    reason-specific `context`. Emitted for every REST-supported
    skippable action — CREATE_TRANSACTION, REVERT_TRANSACTION,
    DELETE_METADATA — not just create.
  - Other successful non-create actions: `null`.
Branch on `OrderSkippedResponse.skipped` to distinguish a skip
from a successful transaction.



## Supported Types

### SchemasTransactionResponse

```go
bulkResultData := shared.CreateBulkResultDataSchemasTransactionResponse(ledgerv3.SchemasTransactionResponse{/* values here */})
```

### OrderSkippedResponse

```go
bulkResultData := shared.CreateBulkResultDataOrderSkippedResponse(ledgerv3.OrderSkippedResponse{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch bulkResultData.Type {
	case shared.BulkResultDataTypeSchemasTransactionResponse:
		// bulkResultData.SchemasTransactionResponse is populated
	case shared.BulkResultDataTypeOrderSkippedResponse:
		// bulkResultData.OrderSkippedResponse is populated
}
```
