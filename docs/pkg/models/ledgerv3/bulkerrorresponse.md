# BulkErrorResponse

Bulk processing returns per-element outcomes in data, including any
preceding successes and subsequent aborted elements. Failures before
processing and panic recovery can return only errorCode/errorMessage.
These object shapes overlap: use a non-exclusive union and preserve
the bulk fields when decoding a processing failure.



## Supported Types

### BulkResponse

```go
bulkErrorResponse := shared.CreateBulkErrorResponseBulkResponse(ledgerv3.BulkResponse{/* values here */})
```

### ErrorResponse

```go
bulkErrorResponse := shared.CreateBulkErrorResponseErrorResponse(ledgerv3.ErrorResponse{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch bulkErrorResponse.Type {
	case shared.BulkErrorResponseTypeBulkResponse:
		// bulkErrorResponse.BulkResponse is populated
	case shared.BulkErrorResponseTypeErrorResponse:
		// bulkErrorResponse.ErrorResponse is populated
}
```
