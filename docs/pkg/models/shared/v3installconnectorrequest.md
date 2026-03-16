# V3InstallConnectorRequest


## Supported Types

### V3AdyenConfig

```go
v3InstallConnectorRequest := shared.CreateV3InstallConnectorRequestAdyen(shared.V3AdyenConfig{/* values here */})
```

### V3AtlarConfig

```go
v3InstallConnectorRequest := shared.CreateV3InstallConnectorRequestAtlar(shared.V3AtlarConfig{/* values here */})
```

### V3BankingcircleConfig

```go
v3InstallConnectorRequest := shared.CreateV3InstallConnectorRequestBankingcircle(shared.V3BankingcircleConfig{/* values here */})
```

### V3CoinbaseprimeConfig

```go
v3InstallConnectorRequest := shared.CreateV3InstallConnectorRequestCoinbaseprime(shared.V3CoinbaseprimeConfig{/* values here */})
```

### V3ColumnConfig

```go
v3InstallConnectorRequest := shared.CreateV3InstallConnectorRequestColumn(shared.V3ColumnConfig{/* values here */})
```

### V3CurrencycloudConfig

```go
v3InstallConnectorRequest := shared.CreateV3InstallConnectorRequestCurrencycloud(shared.V3CurrencycloudConfig{/* values here */})
```

### V3DummypayConfig

```go
v3InstallConnectorRequest := shared.CreateV3InstallConnectorRequestDummypay(shared.V3DummypayConfig{/* values here */})
```

### V3FireblocksConfig

```go
v3InstallConnectorRequest := shared.CreateV3InstallConnectorRequestFireblocks(shared.V3FireblocksConfig{/* values here */})
```

### V3GenericConfig

```go
v3InstallConnectorRequest := shared.CreateV3InstallConnectorRequestGeneric(shared.V3GenericConfig{/* values here */})
```

### V3IncreaseConfig

```go
v3InstallConnectorRequest := shared.CreateV3InstallConnectorRequestIncrease(shared.V3IncreaseConfig{/* values here */})
```

### V3MangopayConfig

```go
v3InstallConnectorRequest := shared.CreateV3InstallConnectorRequestMangopay(shared.V3MangopayConfig{/* values here */})
```

### V3ModulrConfig

```go
v3InstallConnectorRequest := shared.CreateV3InstallConnectorRequestModulr(shared.V3ModulrConfig{/* values here */})
```

### V3MoneycorpConfig

```go
v3InstallConnectorRequest := shared.CreateV3InstallConnectorRequestMoneycorp(shared.V3MoneycorpConfig{/* values here */})
```

### V3PlaidConfig

```go
v3InstallConnectorRequest := shared.CreateV3InstallConnectorRequestPlaid(shared.V3PlaidConfig{/* values here */})
```

### V3PowensConfig

```go
v3InstallConnectorRequest := shared.CreateV3InstallConnectorRequestPowens(shared.V3PowensConfig{/* values here */})
```

### V3QontoConfig

```go
v3InstallConnectorRequest := shared.CreateV3InstallConnectorRequestQonto(shared.V3QontoConfig{/* values here */})
```

### V3StripeConfig

```go
v3InstallConnectorRequest := shared.CreateV3InstallConnectorRequestStripe(shared.V3StripeConfig{/* values here */})
```

### V3TinkConfig

```go
v3InstallConnectorRequest := shared.CreateV3InstallConnectorRequestTink(shared.V3TinkConfig{/* values here */})
```

### V3WiseConfig

```go
v3InstallConnectorRequest := shared.CreateV3InstallConnectorRequestWise(shared.V3WiseConfig{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch v3InstallConnectorRequest.Type {
	case shared.V3InstallConnectorRequestTypeAdyen:
		// v3InstallConnectorRequest.V3AdyenConfig is populated
	case shared.V3InstallConnectorRequestTypeAtlar:
		// v3InstallConnectorRequest.V3AtlarConfig is populated
	case shared.V3InstallConnectorRequestTypeBankingcircle:
		// v3InstallConnectorRequest.V3BankingcircleConfig is populated
	case shared.V3InstallConnectorRequestTypeCoinbaseprime:
		// v3InstallConnectorRequest.V3CoinbaseprimeConfig is populated
	case shared.V3InstallConnectorRequestTypeColumn:
		// v3InstallConnectorRequest.V3ColumnConfig is populated
	case shared.V3InstallConnectorRequestTypeCurrencycloud:
		// v3InstallConnectorRequest.V3CurrencycloudConfig is populated
	case shared.V3InstallConnectorRequestTypeDummypay:
		// v3InstallConnectorRequest.V3DummypayConfig is populated
	case shared.V3InstallConnectorRequestTypeFireblocks:
		// v3InstallConnectorRequest.V3FireblocksConfig is populated
	case shared.V3InstallConnectorRequestTypeGeneric:
		// v3InstallConnectorRequest.V3GenericConfig is populated
	case shared.V3InstallConnectorRequestTypeIncrease:
		// v3InstallConnectorRequest.V3IncreaseConfig is populated
	case shared.V3InstallConnectorRequestTypeMangopay:
		// v3InstallConnectorRequest.V3MangopayConfig is populated
	case shared.V3InstallConnectorRequestTypeModulr:
		// v3InstallConnectorRequest.V3ModulrConfig is populated
	case shared.V3InstallConnectorRequestTypeMoneycorp:
		// v3InstallConnectorRequest.V3MoneycorpConfig is populated
	case shared.V3InstallConnectorRequestTypePlaid:
		// v3InstallConnectorRequest.V3PlaidConfig is populated
	case shared.V3InstallConnectorRequestTypePowens:
		// v3InstallConnectorRequest.V3PowensConfig is populated
	case shared.V3InstallConnectorRequestTypeQonto:
		// v3InstallConnectorRequest.V3QontoConfig is populated
	case shared.V3InstallConnectorRequestTypeStripe:
		// v3InstallConnectorRequest.V3StripeConfig is populated
	case shared.V3InstallConnectorRequestTypeTink:
		// v3InstallConnectorRequest.V3TinkConfig is populated
	case shared.V3InstallConnectorRequestTypeWise:
		// v3InstallConnectorRequest.V3WiseConfig is populated
}
```
