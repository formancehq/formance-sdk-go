# ConnectionCreateChannel

Maturity channel the service resolves to an exact release. Exactly
one of `version` and `channel`. The last applied release is reported in `status.appliedVersion`.


## Example Usage

```go
import (
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/connectivity"
)

value := connectivity.ConnectionCreateChannelStable
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `ConnectionCreateChannelStable` | stable                          |
| `ConnectionCreateChannelRc`     | rc                              |
| `ConnectionCreateChannelBeta`   | beta                            |
| `ConnectionCreateChannelAlpha`  | alpha                           |