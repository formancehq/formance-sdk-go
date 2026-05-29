# V2Subject


## Supported Types

### V2LedgerAccountSubject

```go
v2Subject := shared.CreateV2SubjectV2LedgerAccountSubject(orchestration.V2LedgerAccountSubject{/* values here */})
```

### V2WalletSubject

```go
v2Subject := shared.CreateV2SubjectV2WalletSubject(orchestration.V2WalletSubject{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch v2Subject.Type {
	case shared.V2SubjectTypeV2LedgerAccountSubject:
		// v2Subject.V2LedgerAccountSubject is populated
	case shared.V2SubjectTypeV2WalletSubject:
		// v2Subject.V2WalletSubject is populated
}
```
