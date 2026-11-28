# V3RevertTransactionRequestBody


## Supported Types

### RevertTransactionRequestBody

```go
v3RevertTransactionRequestBody := operations.CreateV3RevertTransactionRequestBodyRevertTransactionRequestBody(ledgerv3.RevertTransactionRequestBody{/* values here */})
```

### SignedRevertTransactionRequest

```go
v3RevertTransactionRequestBody := operations.CreateV3RevertTransactionRequestBodySignedRevertTransactionRequest(ledgerv3.SignedRevertTransactionRequest{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch v3RevertTransactionRequestBody.Type {
	case operations.V3RevertTransactionRequestBodyTypeRevertTransactionRequestBody:
		// v3RevertTransactionRequestBody.RevertTransactionRequestBody is populated
	case operations.V3RevertTransactionRequestBodyTypeSignedRevertTransactionRequest:
		// v3RevertTransactionRequestBody.SignedRevertTransactionRequest is populated
}
```
