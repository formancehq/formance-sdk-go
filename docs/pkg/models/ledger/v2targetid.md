# V2TargetID


## Supported Types

### 

```go
v2TargetID := shared.CreateV2TargetIDStr(string{/* values here */})
```

### 

```go
v2TargetID := shared.CreateV2TargetIDBigint(*big.Int{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch v2TargetID.Type {
	case shared.V2TargetIDTypeStr:
		// v2TargetID.Str is populated
	case shared.V2TargetIDTypeBigint:
		// v2TargetID.Bigint is populated
}
```
