# ClientSecret

Empty when unset; otherwise a presence marker only. Reusable secret material is never returned

## Example Usage

```go
import (
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/ledgerv3"
)

value := ledgerv3.ClientSecretUnknown
```


## Values

| Name                   | Value                  |
| ---------------------- | ---------------------- |
| `ClientSecretUnknown`  |                        |
| `ClientSecretRedacted` | [redacted]             |