# V3PaymentInitiationStatusEnum

## Example Usage

```go
import (
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/payments"
)

value := payments.V3PaymentInitiationStatusEnumUnknown
```


## Values

| Name                                                  | Value                                                 |
| ----------------------------------------------------- | ----------------------------------------------------- |
| `V3PaymentInitiationStatusEnumUnknown`                | UNKNOWN                                               |
| `V3PaymentInitiationStatusEnumWaitingForValidation`   | WAITING_FOR_VALIDATION                                |
| `V3PaymentInitiationStatusEnumScheduledForProcessing` | SCHEDULED_FOR_PROCESSING                              |
| `V3PaymentInitiationStatusEnumProcessing`             | PROCESSING                                            |
| `V3PaymentInitiationStatusEnumProcessed`              | PROCESSED                                             |
| `V3PaymentInitiationStatusEnumFailed`                 | FAILED                                                |
| `V3PaymentInitiationStatusEnumRejected`               | REJECTED                                              |
| `V3PaymentInitiationStatusEnumReverseProcessing`      | REVERSE_PROCESSING                                    |
| `V3PaymentInitiationStatusEnumReverseFailed`          | REVERSE_FAILED                                        |
| `V3PaymentInitiationStatusEnumReversed`               | REVERSED                                              |