# UnsignedMonetaryAmount

Arbitrary-precision unsigned BigUint value encoded as a JSON number by
default, or a canonical decimal string when Formance-Bigint-As-String
is enabled. The same rule applies to every response BigUint value,
including nested fields, regardless of business meaning.
JavaScript consumers should enable the header to avoid number rounding.



## Supported Types

### 

```go
unsignedMonetaryAmount := shared.CreateUnsignedMonetaryAmountBigint(*big.Int{/* values here */})
```

### 

```go
unsignedMonetaryAmount := shared.CreateUnsignedMonetaryAmountStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch unsignedMonetaryAmount.Type {
	case shared.UnsignedMonetaryAmountTypeBigint:
		// unsignedMonetaryAmount.Bigint is populated
	case shared.UnsignedMonetaryAmountTypeStr:
		// unsignedMonetaryAmount.Str is populated
}
```
