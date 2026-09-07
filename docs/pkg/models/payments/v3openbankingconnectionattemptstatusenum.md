# V3OpenBankingConnectionAttemptStatusEnum

Where a link attempt stands, from pending through to completed on success or exited when the user abandoned the flow or the provider reported an error

## Example Usage

```go
import (
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/payments"
)

value := payments.V3OpenBankingConnectionAttemptStatusEnumPending
```


## Values

| Name                                                | Value                                               |
| --------------------------------------------------- | --------------------------------------------------- |
| `V3OpenBankingConnectionAttemptStatusEnumPending`   | pending                                             |
| `V3OpenBankingConnectionAttemptStatusEnumCompleted` | completed                                           |
| `V3OpenBankingConnectionAttemptStatusEnumExited`    | exited                                              |