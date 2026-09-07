# TaskStatus

Where a task stands, from pending or active through to stopped, terminated or failed

## Example Usage

```go
import (
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/payments"
)

value := payments.TaskStatusPending
```


## Values

| Name                   | Value                  |
| ---------------------- | ---------------------- |
| `TaskStatusPending`    | PENDING                |
| `TaskStatusActive`     | ACTIVE                 |
| `TaskStatusStopped`    | STOPPED                |
| `TaskStatusTerminated` | TERMINATED             |
| `TaskStatusFailed`     | FAILED                 |