# SignedMonetaryAmount

Arbitrary-precision SignedBigInt value encoded as a JSON number by
default, or a canonical decimal string when Formance-Bigint-As-String
is enabled. The same rule applies to every response SignedBigInt value,
including nested fields, regardless of business meaning.



## Supported Types

### 

```go
signedMonetaryAmount := shared.CreateSignedMonetaryAmountBigint(*big.Int{/* values here */})
```

### 

```go
signedMonetaryAmount := shared.CreateSignedMonetaryAmountStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch signedMonetaryAmount.Type {
	case shared.SignedMonetaryAmountTypeBigint:
		// signedMonetaryAmount.Bigint is populated
	case shared.SignedMonetaryAmountTypeStr:
		// signedMonetaryAmount.Str is populated
}
```
