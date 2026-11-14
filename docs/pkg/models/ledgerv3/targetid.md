# TargetID

ID of the target (account address or transaction ID). Required.



## Supported Types

### 

```go
targetID := shared.CreateTargetIDStr(string{/* values here */})
```

### 

```go
targetID := shared.CreateTargetIDInteger(int64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch targetID.Type {
	case shared.TargetIDTypeStr:
		// targetID.Str is populated
	case shared.TargetIDTypeInteger:
		// targetID.Integer is populated
}
```
