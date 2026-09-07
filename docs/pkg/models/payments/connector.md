# Connector

The payment provider behind a connector

## Example Usage

```go
import (
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/payments"
)

value := payments.ConnectorStripe
```


## Values

| Name                     | Value                    |
| ------------------------ | ------------------------ |
| `ConnectorStripe`        | STRIPE                   |
| `ConnectorDummyPay`      | DUMMY-PAY                |
| `ConnectorWise`          | WISE                     |
| `ConnectorModulr`        | MODULR                   |
| `ConnectorCurrencyCloud` | CURRENCY-CLOUD           |
| `ConnectorBankingCircle` | BANKING-CIRCLE           |
| `ConnectorMangopay`      | MANGOPAY                 |
| `ConnectorMoneycorp`     | MONEYCORP                |
| `ConnectorAtlar`         | ATLAR                    |
| `ConnectorAdyen`         | ADYEN                    |
| `ConnectorGeneric`       | GENERIC                  |