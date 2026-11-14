# Subject

The counterparty of a wallet movement, either a ledger account or another wallet


## Supported Types

### LedgerAccountSubject

```go
subject := shared.CreateSubjectAccount(wallets.LedgerAccountSubject{/* values here */})
```

### WalletSubject

```go
subject := shared.CreateSubjectWallet(wallets.WalletSubject{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch subject.Type {
	case shared.SubjectTypeAccount:
		// subject.LedgerAccountSubject is populated
	case shared.SubjectTypeWallet:
		// subject.WalletSubject is populated
}
```
