# Runtime

The numscript runtime used to execute the script. Uses "machine" by default, unless the "--experimental-numscript-interpreter" feature flag is passed.

## Example Usage

```go
import (
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/shared"
)

value := shared.RuntimeExperimentalInterpreter
```


## Values

| Name                             | Value                            |
| -------------------------------- | -------------------------------- |
| `RuntimeExperimentalInterpreter` | experimental-interpreter         |
| `RuntimeMachine`                 | machine                          |