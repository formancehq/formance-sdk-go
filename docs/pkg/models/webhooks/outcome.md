# Outcome

How the attempt ended. A retryable failure is retried, a permanent failure is not

## Example Usage

```go
import (
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/webhooks"
)

value := webhooks.OutcomeSucceeded
```


## Values

| Name                      | Value                     |
| ------------------------- | ------------------------- |
| `OutcomeSucceeded`        | succeeded                 |
| `OutcomeRetryableFailure` | retryable_failure         |
| `OutcomePermanentFailure` | permanent_failure         |