# Persistence

How account volumes are stored (NORMAL, EPHEMERAL, or TRANSIENT)

## Example Usage

```go
import (
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/ledgerv3"
)

value := ledgerv3.PersistenceNormal
```


## Values

| Name                   | Value                  |
| ---------------------- | ---------------------- |
| `PersistenceNormal`    | NORMAL                 |
| `PersistenceEphemeral` | EPHEMERAL              |
| `PersistenceTransient` | TRANSIENT              |