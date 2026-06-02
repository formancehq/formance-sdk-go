# V2BulkElement


## Supported Types

### V2BulkElementCreateTransaction

```go
v2BulkElement := shared.CreateV2BulkElementV2BulkElementCreateTransaction(ledger.V2BulkElementCreateTransaction{/* values here */})
```

### V2BulkElementAddMetadata

```go
v2BulkElement := shared.CreateV2BulkElementV2BulkElementAddMetadata(ledger.V2BulkElementAddMetadata{/* values here */})
```

### V2BulkElementRevertTransaction

```go
v2BulkElement := shared.CreateV2BulkElementV2BulkElementRevertTransaction(ledger.V2BulkElementRevertTransaction{/* values here */})
```

### V2BulkElementDeleteMetadata

```go
v2BulkElement := shared.CreateV2BulkElementV2BulkElementDeleteMetadata(ledger.V2BulkElementDeleteMetadata{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch v2BulkElement.Type {
	case shared.V2BulkElementTypeV2BulkElementCreateTransaction:
		// v2BulkElement.V2BulkElementCreateTransaction is populated
	case shared.V2BulkElementTypeV2BulkElementAddMetadata:
		// v2BulkElement.V2BulkElementAddMetadata is populated
	case shared.V2BulkElementTypeV2BulkElementRevertTransaction:
		// v2BulkElement.V2BulkElementRevertTransaction is populated
	case shared.V2BulkElementTypeV2BulkElementDeleteMetadata:
		// v2BulkElement.V2BulkElementDeleteMetadata is populated
}
```
