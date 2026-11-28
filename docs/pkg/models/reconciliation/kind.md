# Kind

How the alert was resolved, whether automatically, by booking a correcting transaction, or by business acceptance

## Example Usage

```go
import (
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/reconciliation"
)

value := reconciliation.KindAuto
```


## Values

| Name                     | Value                    |
| ------------------------ | ------------------------ |
| `KindAuto`               | auto                     |
| `KindFixedByBooking`     | fixed_by_booking         |
| `KindAcceptedByBusiness` | accepted_by_business     |