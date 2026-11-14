# ScheduleKind

Whether the rule runs only when triggered or on a recurring cron schedule

## Example Usage

```go
import (
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/reconciliation"
)

value := reconciliation.ScheduleKindOnDemand
```


## Values

| Name                   | Value                  |
| ---------------------- | ---------------------- |
| `ScheduleKindOnDemand` | on_demand              |
| `ScheduleKindCron`     | cron                   |