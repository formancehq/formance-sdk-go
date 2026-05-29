# Subject


## Supported Types

### LedgerAccountSubject

```go
subject := shared.CreateSubjectLedgerAccountSubject(orchestration.LedgerAccountSubject{/* values here */})
```

### WalletSubject

```go
subject := shared.CreateSubjectWalletSubject(orchestration.WalletSubject{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch subject.Type {
	case shared.SubjectTypeLedgerAccountSubject:
		// subject.LedgerAccountSubject is populated
	case shared.SubjectTypeWalletSubject:
		// subject.WalletSubject is populated
}
```
