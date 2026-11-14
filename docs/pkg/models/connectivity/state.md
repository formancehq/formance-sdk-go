# State

Last recorded ingestion state; it can be stale. paused can represent requested suspension before ingestion exists. finished is successful bounded completion.

## Example Usage

```go
import (
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/connectivity"
)

value := connectivity.StateRunning
```


## Values

| Name            | Value           |
| --------------- | --------------- |
| `StateRunning`  | running         |
| `StateError`    | error           |
| `StateFailed`   | failed          |
| `StateStopped`  | stopped         |
| `StatePaused`   | paused          |
| `StateFinished` | finished        |
| `StateUnknown`  | unknown         |