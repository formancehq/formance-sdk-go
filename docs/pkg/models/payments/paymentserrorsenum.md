# PaymentsErrorsEnum

Machine-readable error code identifying the failure

## Example Usage

```go
import (
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/payments"
)

value := payments.PaymentsErrorsEnumInternal
```


## Values

| Name                                                | Value                                               |
| --------------------------------------------------- | --------------------------------------------------- |
| `PaymentsErrorsEnumInternal`                        | INTERNAL                                            |
| `PaymentsErrorsEnumValidation`                      | VALIDATION                                          |
| `PaymentsErrorsEnumInvalidID`                       | INVALID_ID                                          |
| `PaymentsErrorsEnumMissingOrInvalidBody`            | MISSING_OR_INVALID_BODY                             |
| `PaymentsErrorsEnumConflict`                        | CONFLICT                                            |
| `PaymentsErrorsEnumConnectorCapabilityNotSupported` | CONNECTOR_CAPABILITY_NOT_SUPPORTED                  |
| `PaymentsErrorsEnumNotFound`                        | NOT_FOUND                                           |