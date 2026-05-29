# TargetID


## Supported Types

### 

```go
targetID := shared.CreateTargetIDStr(string{/* values here */})
```

### 

```go
targetID := shared.CreateTargetIDBigint(*big.Int{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch targetID.Type {
	case shared.TargetIDTypeStr:
		// targetID.Str is populated
	case shared.TargetIDTypeBigint:
		// targetID.Bigint is populated
}
```
