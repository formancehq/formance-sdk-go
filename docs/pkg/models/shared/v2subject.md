# V2Subject


## Supported Types

### V2LedgerAccountSubject

```go
v2Subject := shared.CreateV2SubjectAccount(shared.V2LedgerAccountSubject{/* values here */})
```

### V2WalletSubject

```go
v2Subject := shared.CreateV2SubjectWallet(shared.V2WalletSubject{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch v2Subject.Type {
	case shared.V2SubjectTypeAccount:
		// v2Subject.V2LedgerAccountSubject is populated
	case shared.V2SubjectTypeWallet:
		// v2Subject.V2WalletSubject is populated
}
```
