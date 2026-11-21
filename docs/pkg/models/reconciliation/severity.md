# Severity

How serious an alert is, used to route and prioritise notifications

## Example Usage

```go
import (
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/reconciliation"
)

value := reconciliation.SeverityInfo
```


## Values

| Name               | Value              |
| ------------------ | ------------------ |
| `SeverityInfo`     | info               |
| `SeverityLow`      | low                |
| `SeverityMedium`   | medium             |
| `SeverityHigh`     | high               |
| `SeverityCritical` | critical           |