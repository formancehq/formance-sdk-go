# Lifecycle

Last recorded service outcome. completed means successful bounded completion. unknown means the recorded outcome is unrecognized, not healthy.

## Example Usage

```go
import (
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/connectivity"
)

value := connectivity.LifecyclePending
```


## Values

| Name                 | Value                |
| -------------------- | -------------------- |
| `LifecyclePending`   | pending              |
| `LifecycleActive`    | active               |
| `LifecycleSuspended` | suspended            |
| `LifecycleCompleted` | completed            |
| `LifecycleFailed`    | failed               |
| `LifecycleUnknown`   | unknown              |