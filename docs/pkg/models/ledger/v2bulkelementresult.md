# V2BulkElementResult


## Supported Types

### V2BulkElementResultAddMetadata

```go
v2BulkElementResult := shared.CreateV2BulkElementResultAddMetadata(ledger.V2BulkElementResultAddMetadata{/* values here */})
```

### V2BulkElementResultCreateTransaction

```go
v2BulkElementResult := shared.CreateV2BulkElementResultCreateTransaction(ledger.V2BulkElementResultCreateTransaction{/* values here */})
```

### V2BulkElementResultDeleteMetadata

```go
v2BulkElementResult := shared.CreateV2BulkElementResultDeleteMetadata(ledger.V2BulkElementResultDeleteMetadata{/* values here */})
```

### V2BulkElementResultError

```go
v2BulkElementResult := shared.CreateV2BulkElementResultError(ledger.V2BulkElementResultError{/* values here */})
```

### V2BulkElementResultRevertTransaction

```go
v2BulkElementResult := shared.CreateV2BulkElementResultRevertTransaction(ledger.V2BulkElementResultRevertTransaction{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch v2BulkElementResult.Type {
	case shared.V2BulkElementResultTypeAddMetadata:
		// v2BulkElementResult.V2BulkElementResultAddMetadata is populated
	case shared.V2BulkElementResultTypeCreateTransaction:
		// v2BulkElementResult.V2BulkElementResultCreateTransaction is populated
	case shared.V2BulkElementResultTypeDeleteMetadata:
		// v2BulkElementResult.V2BulkElementResultDeleteMetadata is populated
	case shared.V2BulkElementResultTypeError:
		// v2BulkElementResult.V2BulkElementResultError is populated
	case shared.V2BulkElementResultTypeRevertTransaction:
		// v2BulkElementResult.V2BulkElementResultRevertTransaction is populated
}
```
