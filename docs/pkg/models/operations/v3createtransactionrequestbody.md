# V3CreateTransactionRequestBody


## Supported Types

### CreateTransactionRequest

```go
v3CreateTransactionRequestBody := operations.CreateV3CreateTransactionRequestBodyCreateTransactionRequest(ledgerv3.CreateTransactionRequest{/* values here */})
```

### SignedCreateTransactionRequest

```go
v3CreateTransactionRequestBody := operations.CreateV3CreateTransactionRequestBodySignedCreateTransactionRequest(ledgerv3.SignedCreateTransactionRequest{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch v3CreateTransactionRequestBody.Type {
	case operations.V3CreateTransactionRequestBodyTypeCreateTransactionRequest:
		// v3CreateTransactionRequestBody.CreateTransactionRequest is populated
	case operations.V3CreateTransactionRequestBodyTypeSignedCreateTransactionRequest:
		// v3CreateTransactionRequestBody.SignedCreateTransactionRequest is populated
}
```
