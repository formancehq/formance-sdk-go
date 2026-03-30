# TransferInitiationStatus

## Example Usage

```go
import (
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/payments"
)

value := payments.TransferInitiationStatusWaitingForValidation
```


## Values

| Name                                           | Value                                          |
| ---------------------------------------------- | ---------------------------------------------- |
| `TransferInitiationStatusWaitingForValidation` | WAITING_FOR_VALIDATION                         |
| `TransferInitiationStatusProcessing`           | PROCESSING                                     |
| `TransferInitiationStatusProcessed`            | PROCESSED                                      |
| `TransferInitiationStatusFailed`               | FAILED                                         |
| `TransferInitiationStatusRejected`             | REJECTED                                       |
| `TransferInitiationStatusValidated`            | VALIDATED                                      |
| `TransferInitiationStatusAskRetried`           | ASK_RETRIED                                    |
| `TransferInitiationStatusAskReversed`          | ASK_REVERSED                                   |
| `TransferInitiationStatusReverseProcessing`    | REVERSE_PROCESSING                             |
| `TransferInitiationStatusReverseFailed`        | REVERSE_FAILED                                 |
| `TransferInitiationStatusPartiallyReversed`    | PARTIALLY_REVERSED                             |
| `TransferInitiationStatusReversed`             | REVERSED                                       |