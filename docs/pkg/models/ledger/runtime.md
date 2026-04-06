# Runtime

The numscript runtime used to execute the script. Uses "machine" by default, unless the "--experimental-numscript-interpreter" feature flag is passed.

## Example Usage

```go
import (
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/ledger"
)

value := ledger.RuntimeExperimentalInterpreter
```


## Values

| Name                             | Value                            |
| -------------------------------- | -------------------------------- |
| `RuntimeExperimentalInterpreter` | experimental-interpreter         |
| `RuntimeMachine`                 | machine                          |