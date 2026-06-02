# V2BulkElementResult


## Supported Types

### V2BulkElementResultCreateTransaction

```go
v2BulkElementResult := shared.CreateV2BulkElementResultV2BulkElementResultCreateTransaction(ledger.V2BulkElementResultCreateTransaction{/* values here */})
```

### V2BulkElementResultAddMetadata

```go
v2BulkElementResult := shared.CreateV2BulkElementResultV2BulkElementResultAddMetadata(ledger.V2BulkElementResultAddMetadata{/* values here */})
```

### V2BulkElementResultRevertTransaction

```go
v2BulkElementResult := shared.CreateV2BulkElementResultV2BulkElementResultRevertTransaction(ledger.V2BulkElementResultRevertTransaction{/* values here */})
```

### V2BulkElementResultDeleteMetadata

```go
v2BulkElementResult := shared.CreateV2BulkElementResultV2BulkElementResultDeleteMetadata(ledger.V2BulkElementResultDeleteMetadata{/* values here */})
```

### V2BulkElementResultError

```go
v2BulkElementResult := shared.CreateV2BulkElementResultV2BulkElementResultError(ledger.V2BulkElementResultError{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch v2BulkElementResult.Type {
	case shared.V2BulkElementResultTypeV2BulkElementResultCreateTransaction:
		// v2BulkElementResult.V2BulkElementResultCreateTransaction is populated
	case shared.V2BulkElementResultTypeV2BulkElementResultAddMetadata:
		// v2BulkElementResult.V2BulkElementResultAddMetadata is populated
	case shared.V2BulkElementResultTypeV2BulkElementResultRevertTransaction:
		// v2BulkElementResult.V2BulkElementResultRevertTransaction is populated
	case shared.V2BulkElementResultTypeV2BulkElementResultDeleteMetadata:
		// v2BulkElementResult.V2BulkElementResultDeleteMetadata is populated
	case shared.V2BulkElementResultTypeV2BulkElementResultError:
		// v2BulkElementResult.V2BulkElementResultError is populated
}
```
