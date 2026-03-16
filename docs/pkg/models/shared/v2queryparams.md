# V2QueryParams


## Supported Types

### QueryTemplateAccountParams

```go
v2QueryParams := shared.CreateV2QueryParamsQueryTemplateAccountParams(shared.QueryTemplateAccountParams{/* values here */})
```

### QueryTemplateTransactionParams

```go
v2QueryParams := shared.CreateV2QueryParamsQueryTemplateTransactionParams(shared.QueryTemplateTransactionParams{/* values here */})
```

### QueryTemplateLogParams

```go
v2QueryParams := shared.CreateV2QueryParamsQueryTemplateLogParams(shared.QueryTemplateLogParams{/* values here */})
```

### QueryTemplateVolumeParams

```go
v2QueryParams := shared.CreateV2QueryParamsQueryTemplateVolumeParams(shared.QueryTemplateVolumeParams{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch v2QueryParams.Type {
	case shared.V2QueryParamsTypeQueryTemplateAccountParams:
		// v2QueryParams.QueryTemplateAccountParams is populated
	case shared.V2QueryParamsTypeQueryTemplateTransactionParams:
		// v2QueryParams.QueryTemplateTransactionParams is populated
	case shared.V2QueryParamsTypeQueryTemplateLogParams:
		// v2QueryParams.QueryTemplateLogParams is populated
	case shared.V2QueryParamsTypeQueryTemplateVolumeParams:
		// v2QueryParams.QueryTemplateVolumeParams is populated
}
```
