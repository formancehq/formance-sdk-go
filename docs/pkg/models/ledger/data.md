# Data

The payload of the log entry. Structure depends on the log type:
- NEW_TRANSACTION: V2LogDataNewTransaction
- SET_METADATA: V2LogDataSetMetadata
- REVERTED_TRANSACTION: V2LogDataRevertedTransaction
- DELETE_METADATA: V2LogDataDeleteMetadata
- INSERTED_SCHEMA: V2LogDataInsertedSchema



## Supported Types

### V2LogDataNewTransaction

```go
data := shared.CreateDataV2LogDataNewTransaction(ledger.V2LogDataNewTransaction{/* values here */})
```

### V2LogDataSetMetadata

```go
data := shared.CreateDataV2LogDataSetMetadata(ledger.V2LogDataSetMetadata{/* values here */})
```

### V2LogDataRevertedTransaction

```go
data := shared.CreateDataV2LogDataRevertedTransaction(ledger.V2LogDataRevertedTransaction{/* values here */})
```

### V2LogDataDeleteMetadata

```go
data := shared.CreateDataV2LogDataDeleteMetadata(ledger.V2LogDataDeleteMetadata{/* values here */})
```

### V2LogDataInsertedSchema

```go
data := shared.CreateDataV2LogDataInsertedSchema(ledger.V2LogDataInsertedSchema{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch data.Type {
	case shared.DataTypeV2LogDataNewTransaction:
		// data.V2LogDataNewTransaction is populated
	case shared.DataTypeV2LogDataSetMetadata:
		// data.V2LogDataSetMetadata is populated
	case shared.DataTypeV2LogDataRevertedTransaction:
		// data.V2LogDataRevertedTransaction is populated
	case shared.DataTypeV2LogDataDeleteMetadata:
		// data.V2LogDataDeleteMetadata is populated
	case shared.DataTypeV2LogDataInsertedSchema:
		// data.V2LogDataInsertedSchema is populated
}
```
