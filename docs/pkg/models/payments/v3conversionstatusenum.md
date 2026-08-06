# V3ConversionStatusEnum

Lifecycle of a conversion.
`PENDING` — accepted by the PSP, not yet settled.
`COMPLETED` — settled, terminal.
`FAILED` — rejected or reverted, terminal. See `error`.


## Example Usage

```go
import (
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/payments"
)

value := payments.V3ConversionStatusEnumUnknown
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `V3ConversionStatusEnumUnknown`   | UNKNOWN                           |
| `V3ConversionStatusEnumPending`   | PENDING                           |
| `V3ConversionStatusEnumCompleted` | COMPLETED                         |
| `V3ConversionStatusEnumFailed`    | FAILED                            |