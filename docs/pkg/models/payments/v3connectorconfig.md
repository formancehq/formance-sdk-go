# V3ConnectorConfig


## Supported Types

### V3AdyenConfig

```go
v3ConnectorConfig := shared.CreateV3ConnectorConfigAdyen(payments.V3AdyenConfig{/* values here */})
```

### V3AtlarConfig

```go
v3ConnectorConfig := shared.CreateV3ConnectorConfigAtlar(payments.V3AtlarConfig{/* values here */})
```

### V3BankingbridgeConfig

```go
v3ConnectorConfig := shared.CreateV3ConnectorConfigBankingbridge(payments.V3BankingbridgeConfig{/* values here */})
```

### V3BankingcircleConfig

```go
v3ConnectorConfig := shared.CreateV3ConnectorConfigBankingcircle(payments.V3BankingcircleConfig{/* values here */})
```

### V3BitstampConfig

```go
v3ConnectorConfig := shared.CreateV3ConnectorConfigBitstamp(payments.V3BitstampConfig{/* values here */})
```

### V3CoinbaseprimeConfig

```go
v3ConnectorConfig := shared.CreateV3ConnectorConfigCoinbaseprime(payments.V3CoinbaseprimeConfig{/* values here */})
```

### V3ColumnConfig

```go
v3ConnectorConfig := shared.CreateV3ConnectorConfigColumn(payments.V3ColumnConfig{/* values here */})
```

### V3CurrencycloudConfig

```go
v3ConnectorConfig := shared.CreateV3ConnectorConfigCurrencycloud(payments.V3CurrencycloudConfig{/* values here */})
```

### V3DummypayConfig

```go
v3ConnectorConfig := shared.CreateV3ConnectorConfigDummypay(payments.V3DummypayConfig{/* values here */})
```

### V3FireblocksConfig

```go
v3ConnectorConfig := shared.CreateV3ConnectorConfigFireblocks(payments.V3FireblocksConfig{/* values here */})
```

### V3GenericConfig

```go
v3ConnectorConfig := shared.CreateV3ConnectorConfigGeneric(payments.V3GenericConfig{/* values here */})
```

### V3IncreaseConfig

```go
v3ConnectorConfig := shared.CreateV3ConnectorConfigIncrease(payments.V3IncreaseConfig{/* values here */})
```

### V3KrakenproConfig

```go
v3ConnectorConfig := shared.CreateV3ConnectorConfigKrakenpro(payments.V3KrakenproConfig{/* values here */})
```

### V3MangopayConfig

```go
v3ConnectorConfig := shared.CreateV3ConnectorConfigMangopay(payments.V3MangopayConfig{/* values here */})
```

### V3ModulrConfig

```go
v3ConnectorConfig := shared.CreateV3ConnectorConfigModulr(payments.V3ModulrConfig{/* values here */})
```

### V3MoneycorpConfig

```go
v3ConnectorConfig := shared.CreateV3ConnectorConfigMoneycorp(payments.V3MoneycorpConfig{/* values here */})
```

### V3PlaidConfig

```go
v3ConnectorConfig := shared.CreateV3ConnectorConfigPlaid(payments.V3PlaidConfig{/* values here */})
```

### V3PowensConfig

```go
v3ConnectorConfig := shared.CreateV3ConnectorConfigPowens(payments.V3PowensConfig{/* values here */})
```

### V3QontoConfig

```go
v3ConnectorConfig := shared.CreateV3ConnectorConfigQonto(payments.V3QontoConfig{/* values here */})
```

### V3RoutableConfig

```go
v3ConnectorConfig := shared.CreateV3ConnectorConfigRoutable(payments.V3RoutableConfig{/* values here */})
```

### V3StripeConfig

```go
v3ConnectorConfig := shared.CreateV3ConnectorConfigStripe(payments.V3StripeConfig{/* values here */})
```

### V3TinkConfig

```go
v3ConnectorConfig := shared.CreateV3ConnectorConfigTink(payments.V3TinkConfig{/* values here */})
```

### V3WiseConfig

```go
v3ConnectorConfig := shared.CreateV3ConnectorConfigWise(payments.V3WiseConfig{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch v3ConnectorConfig.Type {
	case shared.V3ConnectorConfigTypeAdyen:
		// v3ConnectorConfig.V3AdyenConfig is populated
	case shared.V3ConnectorConfigTypeAtlar:
		// v3ConnectorConfig.V3AtlarConfig is populated
	case shared.V3ConnectorConfigTypeBankingbridge:
		// v3ConnectorConfig.V3BankingbridgeConfig is populated
	case shared.V3ConnectorConfigTypeBankingcircle:
		// v3ConnectorConfig.V3BankingcircleConfig is populated
	case shared.V3ConnectorConfigTypeBitstamp:
		// v3ConnectorConfig.V3BitstampConfig is populated
	case shared.V3ConnectorConfigTypeCoinbaseprime:
		// v3ConnectorConfig.V3CoinbaseprimeConfig is populated
	case shared.V3ConnectorConfigTypeColumn:
		// v3ConnectorConfig.V3ColumnConfig is populated
	case shared.V3ConnectorConfigTypeCurrencycloud:
		// v3ConnectorConfig.V3CurrencycloudConfig is populated
	case shared.V3ConnectorConfigTypeDummypay:
		// v3ConnectorConfig.V3DummypayConfig is populated
	case shared.V3ConnectorConfigTypeFireblocks:
		// v3ConnectorConfig.V3FireblocksConfig is populated
	case shared.V3ConnectorConfigTypeGeneric:
		// v3ConnectorConfig.V3GenericConfig is populated
	case shared.V3ConnectorConfigTypeIncrease:
		// v3ConnectorConfig.V3IncreaseConfig is populated
	case shared.V3ConnectorConfigTypeKrakenpro:
		// v3ConnectorConfig.V3KrakenproConfig is populated
	case shared.V3ConnectorConfigTypeMangopay:
		// v3ConnectorConfig.V3MangopayConfig is populated
	case shared.V3ConnectorConfigTypeModulr:
		// v3ConnectorConfig.V3ModulrConfig is populated
	case shared.V3ConnectorConfigTypeMoneycorp:
		// v3ConnectorConfig.V3MoneycorpConfig is populated
	case shared.V3ConnectorConfigTypePlaid:
		// v3ConnectorConfig.V3PlaidConfig is populated
	case shared.V3ConnectorConfigTypePowens:
		// v3ConnectorConfig.V3PowensConfig is populated
	case shared.V3ConnectorConfigTypeQonto:
		// v3ConnectorConfig.V3QontoConfig is populated
	case shared.V3ConnectorConfigTypeRoutable:
		// v3ConnectorConfig.V3RoutableConfig is populated
	case shared.V3ConnectorConfigTypeStripe:
		// v3ConnectorConfig.V3StripeConfig is populated
	case shared.V3ConnectorConfigTypeTink:
		// v3ConnectorConfig.V3TinkConfig is populated
	case shared.V3ConnectorConfigTypeWise:
		// v3ConnectorConfig.V3WiseConfig is populated
}
```
