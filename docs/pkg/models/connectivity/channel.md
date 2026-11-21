# Channel

Maturity channel the service resolves to an exact release. Exactly
one of `version` and `channel`. The last applied release is reported in `status.appliedVersion`.


## Example Usage

```go
import (
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/connectivity"
)

value := connectivity.ChannelStable
```


## Values

| Name            | Value           |
| --------------- | --------------- |
| `ChannelStable` | stable          |
| `ChannelRc`     | rc              |
| `ChannelBeta`   | beta            |
| `ChannelAlpha`  | alpha           |