# BulkElementData

Operation-specific data


## Supported Types

### CreateTransactionRequest

```go
bulkElementData := shared.CreateBulkElementDataCreateTransactionRequest(ledgerv3.CreateTransactionRequest{/* values here */})
```

### AddMetadataRequest

```go
bulkElementData := shared.CreateBulkElementDataAddMetadataRequest(ledgerv3.AddMetadataRequest{/* values here */})
```

### RevertTransactionRequest

```go
bulkElementData := shared.CreateBulkElementDataRevertTransactionRequest(ledgerv3.RevertTransactionRequest{/* values here */})
```

### DeleteMetadataRequest

```go
bulkElementData := shared.CreateBulkElementDataDeleteMetadataRequest(ledgerv3.DeleteMetadataRequest{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch bulkElementData.Type {
	case shared.BulkElementDataTypeCreateTransactionRequest:
		// bulkElementData.CreateTransactionRequest is populated
	case shared.BulkElementDataTypeAddMetadataRequest:
		// bulkElementData.AddMetadataRequest is populated
	case shared.BulkElementDataTypeRevertTransactionRequest:
		// bulkElementData.RevertTransactionRequest is populated
	case shared.BulkElementDataTypeDeleteMetadataRequest:
		// bulkElementData.DeleteMetadataRequest is populated
}
```
