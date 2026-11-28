# Type

What happened, whether an evaluation outcome or a manual transition

## Example Usage

```go
import (
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/reconciliation"
)

value := reconciliation.TypeFail
```


## Values

| Name           | Value          |
| -------------- | -------------- |
| `TypeFail`     | fail           |
| `TypePass`     | pass           |
| `TypeAck`      | ack            |
| `TypeResolve`  | resolve        |
| `TypeAccept`   | accept         |
| `TypeSnooze`   | snooze         |
| `TypeUnsnooze` | unsnooze       |