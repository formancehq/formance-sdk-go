# V2LogDataSetMetadataTargetID


## Supported Types

### 

```go
v2LogDataSetMetadataTargetID := shared.CreateV2LogDataSetMetadataTargetIDStr(string{/* values here */})
```

### 

```go
v2LogDataSetMetadataTargetID := shared.CreateV2LogDataSetMetadataTargetIDBigint(*big.Int{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch v2LogDataSetMetadataTargetID.Type {
	case shared.V2LogDataSetMetadataTargetIDTypeStr:
		// v2LogDataSetMetadataTargetID.Str is populated
	case shared.V2LogDataSetMetadataTargetIDTypeBigint:
		// v2LogDataSetMetadataTargetID.Bigint is populated
}
```
