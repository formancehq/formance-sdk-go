# V2BulkElementResult


## Supported Types

### Schemas

```go
v2BulkElementResult := shared.CreateV2BulkElementResultAddMetadata(shared.Schemas{/* values here */})
```

### V2BulkElementResultCreateTransactionSchemas

```go
v2BulkElementResult := shared.CreateV2BulkElementResultCreateTransaction(shared.V2BulkElementResultCreateTransactionSchemas{/* values here */})
```

### V2BulkElementResultDeleteMetadataSchemas

```go
v2BulkElementResult := shared.CreateV2BulkElementResultDeleteMetadata(shared.V2BulkElementResultDeleteMetadataSchemas{/* values here */})
```

### V2BulkElementResultErrorSchemas

```go
v2BulkElementResult := shared.CreateV2BulkElementResultError(shared.V2BulkElementResultErrorSchemas{/* values here */})
```

### V2BulkElementResultRevertTransactionSchemas

```go
v2BulkElementResult := shared.CreateV2BulkElementResultRevertTransaction(shared.V2BulkElementResultRevertTransactionSchemas{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch v2BulkElementResult.Type {
	case shared.V2BulkElementResultTypeAddMetadata:
		// v2BulkElementResult.Schemas is populated
	case shared.V2BulkElementResultTypeCreateTransaction:
		// v2BulkElementResult.V2BulkElementResultCreateTransactionSchemas is populated
	case shared.V2BulkElementResultTypeDeleteMetadata:
		// v2BulkElementResult.V2BulkElementResultDeleteMetadataSchemas is populated
	case shared.V2BulkElementResultTypeError:
		// v2BulkElementResult.V2BulkElementResultErrorSchemas is populated
	case shared.V2BulkElementResultTypeRevertTransaction:
		// v2BulkElementResult.V2BulkElementResultRevertTransactionSchemas is populated
}
```
