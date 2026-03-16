# V2BulkElement


## Supported Types

### V2BulkElementAddMetadata

```go
v2BulkElement := shared.CreateV2BulkElementAddMetadata(shared.V2BulkElementAddMetadata{/* values here */})
```

### V2BulkElementCreateTransaction

```go
v2BulkElement := shared.CreateV2BulkElementCreateTransaction(shared.V2BulkElementCreateTransaction{/* values here */})
```

### V2BulkElementDeleteMetadata

```go
v2BulkElement := shared.CreateV2BulkElementDeleteMetadata(shared.V2BulkElementDeleteMetadata{/* values here */})
```

### V2BulkElementRevertTransaction

```go
v2BulkElement := shared.CreateV2BulkElementRevertTransaction(shared.V2BulkElementRevertTransaction{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch v2BulkElement.Type {
	case shared.V2BulkElementTypeAddMetadata:
		// v2BulkElement.V2BulkElementAddMetadata is populated
	case shared.V2BulkElementTypeCreateTransaction:
		// v2BulkElement.V2BulkElementCreateTransaction is populated
	case shared.V2BulkElementTypeDeleteMetadata:
		// v2BulkElement.V2BulkElementDeleteMetadata is populated
	case shared.V2BulkElementTypeRevertTransaction:
		// v2BulkElement.V2BulkElementRevertTransaction is populated
}
```
