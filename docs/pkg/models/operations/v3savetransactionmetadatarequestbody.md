# V3SaveTransactionMetadataRequestBody


## Supported Types

### 

```go
v3SaveTransactionMetadataRequestBody := operations.CreateV3SaveTransactionMetadataRequestBodyMapOfMetadataValue(map[string]*ledgerv3.MetadataValue{/* values here */})
```

### SignedSaveTransactionMetadataRequest

```go
v3SaveTransactionMetadataRequestBody := operations.CreateV3SaveTransactionMetadataRequestBodySignedSaveTransactionMetadataRequest(ledgerv3.SignedSaveTransactionMetadataRequest{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch v3SaveTransactionMetadataRequestBody.Type {
	case operations.V3SaveTransactionMetadataRequestBodyTypeMapOfMetadataValue:
		// v3SaveTransactionMetadataRequestBody.MapOfMetadataValue is populated
	case operations.V3SaveTransactionMetadataRequestBodyTypeSignedSaveTransactionMetadataRequest:
		// v3SaveTransactionMetadataRequestBody.SignedSaveTransactionMetadataRequest is populated
}
```
