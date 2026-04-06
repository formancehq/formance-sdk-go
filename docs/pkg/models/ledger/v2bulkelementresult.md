# V2BulkElementResult


## Supported Types

### SchemasV2BaseBulkElementResult

```go
v2BulkElementResult := shared.CreateV2BulkElementResultSchemasV2BaseBulkElementResult(ledger.SchemasV2BaseBulkElementResult{/* values here */})
```

### V2BaseBulkElementResult

```go
v2BulkElementResult := shared.CreateV2BulkElementResultV2BaseBulkElementResult(ledger.V2BaseBulkElementResult{/* values here */})
```

### SchemasV2BulkElementResultRevertTransactionV2BaseBulkElementResult

```go
v2BulkElementResult := shared.CreateV2BulkElementResultSchemasV2BulkElementResultRevertTransactionV2BaseBulkElementResult(ledger.SchemasV2BulkElementResultRevertTransactionV2BaseBulkElementResult{/* values here */})
```

### SchemasV2BulkElementResultDeleteMetadataV2BaseBulkElementResult

```go
v2BulkElementResult := shared.CreateV2BulkElementResultSchemasV2BulkElementResultDeleteMetadataV2BaseBulkElementResult(ledger.SchemasV2BulkElementResultDeleteMetadataV2BaseBulkElementResult{/* values here */})
```

### SchemasV2BulkElementResultErrorV2BaseBulkElementResult

```go
v2BulkElementResult := shared.CreateV2BulkElementResultSchemasV2BulkElementResultErrorV2BaseBulkElementResult(ledger.SchemasV2BulkElementResultErrorV2BaseBulkElementResult{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch v2BulkElementResult.Type {
	case shared.V2BulkElementResultTypeSchemasV2BaseBulkElementResult:
		// v2BulkElementResult.SchemasV2BaseBulkElementResult is populated
	case shared.V2BulkElementResultTypeV2BaseBulkElementResult:
		// v2BulkElementResult.V2BaseBulkElementResult is populated
	case shared.V2BulkElementResultTypeSchemasV2BulkElementResultRevertTransactionV2BaseBulkElementResult:
		// v2BulkElementResult.SchemasV2BulkElementResultRevertTransactionV2BaseBulkElementResult is populated
	case shared.V2BulkElementResultTypeSchemasV2BulkElementResultDeleteMetadataV2BaseBulkElementResult:
		// v2BulkElementResult.SchemasV2BulkElementResultDeleteMetadataV2BaseBulkElementResult is populated
	case shared.V2BulkElementResultTypeSchemasV2BulkElementResultErrorV2BaseBulkElementResult:
		// v2BulkElementResult.SchemasV2BulkElementResultErrorV2BaseBulkElementResult is populated
}
```
