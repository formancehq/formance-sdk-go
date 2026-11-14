# V3BulkOperationsRequestBody


## Supported Types

### 

```go
v3BulkOperationsRequestBody := operations.CreateV3BulkOperationsRequestBodyArrayOfBulkElement([]ledgerv3.BulkElement{/* values here */})
```

### SignedBulkRequest

```go
v3BulkOperationsRequestBody := operations.CreateV3BulkOperationsRequestBodySignedBulkRequest(ledgerv3.SignedBulkRequest{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch v3BulkOperationsRequestBody.Type {
	case operations.V3BulkOperationsRequestBodyTypeArrayOfBulkElement:
		// v3BulkOperationsRequestBody.ArrayOfBulkElement is populated
	case operations.V3BulkOperationsRequestBodyTypeSignedBulkRequest:
		// v3BulkOperationsRequestBody.SignedBulkRequest is populated
}
```
