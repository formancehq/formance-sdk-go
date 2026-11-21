# PostingAmount

Unsigned amount from 0 through 2^256-1. Requests accept JSON integer
tokens or canonical decimal strings independently of the response
header; the server enforces the bound for both forms. Responses use
numbers by default and strings with Formance-Bigint-As-String enabled,
as do all response Uint256 values regardless of business meaning.



## Supported Types

### 

```go
postingAmount := shared.CreatePostingAmountBigint(*big.Int{/* values here */})
```

### 

```go
postingAmount := shared.CreatePostingAmountStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch postingAmount.Type {
	case shared.PostingAmountTypeBigint:
		// postingAmount.Bigint is populated
	case shared.PostingAmountTypeStr:
		// postingAmount.Str is populated
}
```
