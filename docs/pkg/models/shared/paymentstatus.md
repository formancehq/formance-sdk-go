# PaymentStatus

## Example Usage

```go
import (
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/shared"
)

value := shared.PaymentStatusPending
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `PaymentStatusPending`         | PENDING                        |
| `PaymentStatusSucceeded`       | SUCCEEDED                      |
| `PaymentStatusCancelled`       | CANCELLED                      |
| `PaymentStatusFailed`          | FAILED                         |
| `PaymentStatusExpired`         | EXPIRED                        |
| `PaymentStatusRefunded`        | REFUNDED                       |
| `PaymentStatusRefundedFailure` | REFUNDED_FAILURE               |
| `PaymentStatusDispute`         | DISPUTE                        |
| `PaymentStatusDisputeWon`      | DISPUTE_WON                    |
| `PaymentStatusDisputeLost`     | DISPUTE_LOST                   |
| `PaymentStatusOther`           | OTHER                          |