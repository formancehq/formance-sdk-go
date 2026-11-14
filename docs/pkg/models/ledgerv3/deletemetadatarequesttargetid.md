# DeleteMetadataRequestTargetID

ID of the target (account address or transaction ID). Required.



## Supported Types

### 

```go
deleteMetadataRequestTargetID := shared.CreateDeleteMetadataRequestTargetIDStr(string{/* values here */})
```

### 

```go
deleteMetadataRequestTargetID := shared.CreateDeleteMetadataRequestTargetIDInteger(int64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch deleteMetadataRequestTargetID.Type {
	case shared.DeleteMetadataRequestTargetIDTypeStr:
		// deleteMetadataRequestTargetID.Str is populated
	case shared.DeleteMetadataRequestTargetIDTypeInteger:
		// deleteMetadataRequestTargetID.Integer is populated
}
```
