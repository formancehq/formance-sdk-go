# ConnectorConfig


## Supported Types

### StripeConfig

```go
connectorConfig := shared.CreateConnectorConfigStripeConfig(payments.StripeConfig{/* values here */})
```

### DummyPayConfig

```go
connectorConfig := shared.CreateConnectorConfigDummyPayConfig(payments.DummyPayConfig{/* values here */})
```

### WiseConfig

```go
connectorConfig := shared.CreateConnectorConfigWiseConfig(payments.WiseConfig{/* values here */})
```

### ModulrConfig

```go
connectorConfig := shared.CreateConnectorConfigModulrConfig(payments.ModulrConfig{/* values here */})
```

### CurrencyCloudConfig

```go
connectorConfig := shared.CreateConnectorConfigCurrencyCloudConfig(payments.CurrencyCloudConfig{/* values here */})
```

### BankingCircleConfig

```go
connectorConfig := shared.CreateConnectorConfigBankingCircleConfig(payments.BankingCircleConfig{/* values here */})
```

### MangoPayConfig

```go
connectorConfig := shared.CreateConnectorConfigMangoPayConfig(payments.MangoPayConfig{/* values here */})
```

### MoneycorpConfig

```go
connectorConfig := shared.CreateConnectorConfigMoneycorpConfig(payments.MoneycorpConfig{/* values here */})
```

### AtlarConfig

```go
connectorConfig := shared.CreateConnectorConfigAtlarConfig(payments.AtlarConfig{/* values here */})
```

### AdyenConfig

```go
connectorConfig := shared.CreateConnectorConfigAdyenConfig(payments.AdyenConfig{/* values here */})
```

### GenericConfig

```go
connectorConfig := shared.CreateConnectorConfigGenericConfig(payments.GenericConfig{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch connectorConfig.Type {
	case shared.ConnectorConfigTypeStripeConfig:
		// connectorConfig.StripeConfig is populated
	case shared.ConnectorConfigTypeDummyPayConfig:
		// connectorConfig.DummyPayConfig is populated
	case shared.ConnectorConfigTypeWiseConfig:
		// connectorConfig.WiseConfig is populated
	case shared.ConnectorConfigTypeModulrConfig:
		// connectorConfig.ModulrConfig is populated
	case shared.ConnectorConfigTypeCurrencyCloudConfig:
		// connectorConfig.CurrencyCloudConfig is populated
	case shared.ConnectorConfigTypeBankingCircleConfig:
		// connectorConfig.BankingCircleConfig is populated
	case shared.ConnectorConfigTypeMangoPayConfig:
		// connectorConfig.MangoPayConfig is populated
	case shared.ConnectorConfigTypeMoneycorpConfig:
		// connectorConfig.MoneycorpConfig is populated
	case shared.ConnectorConfigTypeAtlarConfig:
		// connectorConfig.AtlarConfig is populated
	case shared.ConnectorConfigTypeAdyenConfig:
		// connectorConfig.AdyenConfig is populated
	case shared.ConnectorConfigTypeGenericConfig:
		// connectorConfig.GenericConfig is populated
}
```
