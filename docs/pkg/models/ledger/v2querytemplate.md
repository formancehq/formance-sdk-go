# V2QueryTemplate

A query template. `resource` selects which parameters apply.


## Supported Types

### V2AccountQueryTemplate

```go
v2QueryTemplate := shared.CreateV2QueryTemplateAccounts(ledger.V2AccountQueryTemplate{/* values here */})
```

### V2LogQueryTemplate

```go
v2QueryTemplate := shared.CreateV2QueryTemplateLogs(ledger.V2LogQueryTemplate{/* values here */})
```

### V2TransactionQueryTemplate

```go
v2QueryTemplate := shared.CreateV2QueryTemplateTransactions(ledger.V2TransactionQueryTemplate{/* values here */})
```

### V2VolumeQueryTemplate

```go
v2QueryTemplate := shared.CreateV2QueryTemplateVolumes(ledger.V2VolumeQueryTemplate{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch v2QueryTemplate.Type {
	case shared.V2QueryTemplateTypeAccounts:
		// v2QueryTemplate.V2AccountQueryTemplate is populated
	case shared.V2QueryTemplateTypeLogs:
		// v2QueryTemplate.V2LogQueryTemplate is populated
	case shared.V2QueryTemplateTypeTransactions:
		// v2QueryTemplate.V2TransactionQueryTemplate is populated
	case shared.V2QueryTemplateTypeVolumes:
		// v2QueryTemplate.V2VolumeQueryTemplate is populated
}
```
