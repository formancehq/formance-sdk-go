# V3ConnectorConfig


## Supported Types

### V3AdyenConfig

```go
v3ConnectorConfig := shared.CreateV3ConnectorConfigV3AdyenConfig(payments.V3AdyenConfig{/* values here */})
```

### V3AtlarConfig

```go
v3ConnectorConfig := shared.CreateV3ConnectorConfigV3AtlarConfig(payments.V3AtlarConfig{/* values here */})
```

### V3BankingcircleConfig

```go
v3ConnectorConfig := shared.CreateV3ConnectorConfigV3BankingcircleConfig(payments.V3BankingcircleConfig{/* values here */})
```

### V3ColumnConfig

```go
v3ConnectorConfig := shared.CreateV3ConnectorConfigV3ColumnConfig(payments.V3ColumnConfig{/* values here */})
```

### V3CurrencycloudConfig

```go
v3ConnectorConfig := shared.CreateV3ConnectorConfigV3CurrencycloudConfig(payments.V3CurrencycloudConfig{/* values here */})
```

### V3DummypayConfig

```go
v3ConnectorConfig := shared.CreateV3ConnectorConfigV3DummypayConfig(payments.V3DummypayConfig{/* values here */})
```

### V3GenericConfig

```go
v3ConnectorConfig := shared.CreateV3ConnectorConfigV3GenericConfig(payments.V3GenericConfig{/* values here */})
```

### V3IncreaseConfig

```go
v3ConnectorConfig := shared.CreateV3ConnectorConfigV3IncreaseConfig(payments.V3IncreaseConfig{/* values here */})
```

### V3MangopayConfig

```go
v3ConnectorConfig := shared.CreateV3ConnectorConfigV3MangopayConfig(payments.V3MangopayConfig{/* values here */})
```

### V3ModulrConfig

```go
v3ConnectorConfig := shared.CreateV3ConnectorConfigV3ModulrConfig(payments.V3ModulrConfig{/* values here */})
```

### V3MoneycorpConfig

```go
v3ConnectorConfig := shared.CreateV3ConnectorConfigV3MoneycorpConfig(payments.V3MoneycorpConfig{/* values here */})
```

### V3PlaidConfig

```go
v3ConnectorConfig := shared.CreateV3ConnectorConfigV3PlaidConfig(payments.V3PlaidConfig{/* values here */})
```

### V3PowensConfig

```go
v3ConnectorConfig := shared.CreateV3ConnectorConfigV3PowensConfig(payments.V3PowensConfig{/* values here */})
```

### V3QontoConfig

```go
v3ConnectorConfig := shared.CreateV3ConnectorConfigV3QontoConfig(payments.V3QontoConfig{/* values here */})
```

### V3StripeConfig

```go
v3ConnectorConfig := shared.CreateV3ConnectorConfigV3StripeConfig(payments.V3StripeConfig{/* values here */})
```

### V3TinkConfig

```go
v3ConnectorConfig := shared.CreateV3ConnectorConfigV3TinkConfig(payments.V3TinkConfig{/* values here */})
```

### V3WiseConfig

```go
v3ConnectorConfig := shared.CreateV3ConnectorConfigV3WiseConfig(payments.V3WiseConfig{/* values here */})
```

### V3CoinbaseprimeConfig

```go
v3ConnectorConfig := shared.CreateV3ConnectorConfigV3CoinbaseprimeConfig(payments.V3CoinbaseprimeConfig{/* values here */})
```

### V3FireblocksConfig

```go
v3ConnectorConfig := shared.CreateV3ConnectorConfigV3FireblocksConfig(payments.V3FireblocksConfig{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch v3ConnectorConfig.Type {
	case shared.V3ConnectorConfigTypeV3AdyenConfig:
		// v3ConnectorConfig.V3AdyenConfig is populated
	case shared.V3ConnectorConfigTypeV3AtlarConfig:
		// v3ConnectorConfig.V3AtlarConfig is populated
	case shared.V3ConnectorConfigTypeV3BankingcircleConfig:
		// v3ConnectorConfig.V3BankingcircleConfig is populated
	case shared.V3ConnectorConfigTypeV3ColumnConfig:
		// v3ConnectorConfig.V3ColumnConfig is populated
	case shared.V3ConnectorConfigTypeV3CurrencycloudConfig:
		// v3ConnectorConfig.V3CurrencycloudConfig is populated
	case shared.V3ConnectorConfigTypeV3DummypayConfig:
		// v3ConnectorConfig.V3DummypayConfig is populated
	case shared.V3ConnectorConfigTypeV3GenericConfig:
		// v3ConnectorConfig.V3GenericConfig is populated
	case shared.V3ConnectorConfigTypeV3IncreaseConfig:
		// v3ConnectorConfig.V3IncreaseConfig is populated
	case shared.V3ConnectorConfigTypeV3MangopayConfig:
		// v3ConnectorConfig.V3MangopayConfig is populated
	case shared.V3ConnectorConfigTypeV3ModulrConfig:
		// v3ConnectorConfig.V3ModulrConfig is populated
	case shared.V3ConnectorConfigTypeV3MoneycorpConfig:
		// v3ConnectorConfig.V3MoneycorpConfig is populated
	case shared.V3ConnectorConfigTypeV3PlaidConfig:
		// v3ConnectorConfig.V3PlaidConfig is populated
	case shared.V3ConnectorConfigTypeV3PowensConfig:
		// v3ConnectorConfig.V3PowensConfig is populated
	case shared.V3ConnectorConfigTypeV3QontoConfig:
		// v3ConnectorConfig.V3QontoConfig is populated
	case shared.V3ConnectorConfigTypeV3StripeConfig:
		// v3ConnectorConfig.V3StripeConfig is populated
	case shared.V3ConnectorConfigTypeV3TinkConfig:
		// v3ConnectorConfig.V3TinkConfig is populated
	case shared.V3ConnectorConfigTypeV3WiseConfig:
		// v3ConnectorConfig.V3WiseConfig is populated
	case shared.V3ConnectorConfigTypeV3CoinbaseprimeConfig:
		// v3ConnectorConfig.V3CoinbaseprimeConfig is populated
	case shared.V3ConnectorConfigTypeV3FireblocksConfig:
		// v3ConnectorConfig.V3FireblocksConfig is populated
}
```
