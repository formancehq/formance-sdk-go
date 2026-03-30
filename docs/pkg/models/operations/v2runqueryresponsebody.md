# V2RunQueryResponseBody

OK


## Supported Types

### V2TransactionsCursorResponse

```go
v2RunQueryResponseBody := operations.CreateV2RunQueryResponseBodyV2TransactionsCursorResponse(ledger.V2TransactionsCursorResponse{/* values here */})
```

### V2AccountsCursorResponse

```go
v2RunQueryResponseBody := operations.CreateV2RunQueryResponseBodyV2AccountsCursorResponse(ledger.V2AccountsCursorResponse{/* values here */})
```

### V2LogsCursorResponse

```go
v2RunQueryResponseBody := operations.CreateV2RunQueryResponseBodyV2LogsCursorResponse(ledger.V2LogsCursorResponse{/* values here */})
```

### V2VolumesWithBalanceCursorResponse

```go
v2RunQueryResponseBody := operations.CreateV2RunQueryResponseBodyV2VolumesWithBalanceCursorResponse(ledger.V2VolumesWithBalanceCursorResponse{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch v2RunQueryResponseBody.Type {
	case operations.V2RunQueryResponseBodyTypeV2TransactionsCursorResponse:
		// v2RunQueryResponseBody.V2TransactionsCursorResponse is populated
	case operations.V2RunQueryResponseBodyTypeV2AccountsCursorResponse:
		// v2RunQueryResponseBody.V2AccountsCursorResponse is populated
	case operations.V2RunQueryResponseBodyTypeV2LogsCursorResponse:
		// v2RunQueryResponseBody.V2LogsCursorResponse is populated
	case operations.V2RunQueryResponseBodyTypeV2VolumesWithBalanceCursorResponse:
		// v2RunQueryResponseBody.V2VolumesWithBalanceCursorResponse is populated
}
```
