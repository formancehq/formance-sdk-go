# Cadence

Reconciliation rhythm. Scopes each failing fingerprint into a period so a
March break and an April break are distinct, independently-closable cases.
`continuous` (default) is a single unbounded period (live monitoring).


## Example Usage

```go
import (
	"github.com/formancehq/formance-sdk-go/v4/pkg/models/reconciliation"
)

value := reconciliation.CadenceContinuous
```


## Values

| Name                | Value               |
| ------------------- | ------------------- |
| `CadenceContinuous` | continuous          |
| `CadenceDaily`      | daily               |
| `CadenceWeekly`     | weekly              |
| `CadenceMonthly`    | monthly             |