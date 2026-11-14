# MirrorSourceRead

Detached credential-safe mirror configuration returned by ledger reads


## Supported Types

### One

```go
mirrorSourceRead := shared.CreateMirrorSourceReadOne(ledgerv3.One{/* values here */})
```

### Two

```go
mirrorSourceRead := shared.CreateMirrorSourceReadTwo(ledgerv3.Two{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch mirrorSourceRead.Type {
	case shared.MirrorSourceReadTypeOne:
		// mirrorSourceRead.One is populated
	case shared.MirrorSourceReadTypeTwo:
		// mirrorSourceRead.Two is populated
}
```
