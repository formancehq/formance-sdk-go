# ConnectorConfig


## Supported Types

### AdyenConfig

```go
connectorConfig := shared.CreateConnectorConfigAdyen(shared.AdyenConfig{/* values here */})
```

### AtlarConfig

```go
connectorConfig := shared.CreateConnectorConfigAtlar(shared.AtlarConfig{/* values here */})
```

### BankingCircleConfig

```go
connectorConfig := shared.CreateConnectorConfigBankingcircle(shared.BankingCircleConfig{/* values here */})
```

### CurrencyCloudConfig

```go
connectorConfig := shared.CreateConnectorConfigCurrencycloud(shared.CurrencyCloudConfig{/* values here */})
```

### DummyPayConfig

```go
connectorConfig := shared.CreateConnectorConfigDummypay(shared.DummyPayConfig{/* values here */})
```

### GenericConfig

```go
connectorConfig := shared.CreateConnectorConfigGeneric(shared.GenericConfig{/* values here */})
```

### MangoPayConfig

```go
connectorConfig := shared.CreateConnectorConfigMangopay(shared.MangoPayConfig{/* values here */})
```

### ModulrConfig

```go
connectorConfig := shared.CreateConnectorConfigModulr(shared.ModulrConfig{/* values here */})
```

### MoneycorpConfig

```go
connectorConfig := shared.CreateConnectorConfigMoneycorp(shared.MoneycorpConfig{/* values here */})
```

### StripeConfig

```go
connectorConfig := shared.CreateConnectorConfigStripe(shared.StripeConfig{/* values here */})
```

### WiseConfig

```go
connectorConfig := shared.CreateConnectorConfigWise(shared.WiseConfig{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch connectorConfig.Type {
	case shared.ConnectorConfigTypeAdyen:
		// connectorConfig.AdyenConfig is populated
	case shared.ConnectorConfigTypeAtlar:
		// connectorConfig.AtlarConfig is populated
	case shared.ConnectorConfigTypeBankingcircle:
		// connectorConfig.BankingCircleConfig is populated
	case shared.ConnectorConfigTypeCurrencycloud:
		// connectorConfig.CurrencyCloudConfig is populated
	case shared.ConnectorConfigTypeDummypay:
		// connectorConfig.DummyPayConfig is populated
	case shared.ConnectorConfigTypeGeneric:
		// connectorConfig.GenericConfig is populated
	case shared.ConnectorConfigTypeMangopay:
		// connectorConfig.MangoPayConfig is populated
	case shared.ConnectorConfigTypeModulr:
		// connectorConfig.ModulrConfig is populated
	case shared.ConnectorConfigTypeMoneycorp:
		// connectorConfig.MoneycorpConfig is populated
	case shared.ConnectorConfigTypeStripe:
		// connectorConfig.StripeConfig is populated
	case shared.ConnectorConfigTypeWise:
		// connectorConfig.WiseConfig is populated
}
```
