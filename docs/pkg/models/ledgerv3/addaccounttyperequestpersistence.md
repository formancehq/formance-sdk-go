# AddAccountTypeRequestPersistence

How account volumes are stored.
NORMAL: fully persisted (default).
EPHEMERAL: persisted, purged when zero balance after commit.
TRANSIENT: never persisted, must be zero at end of batch.


## Example Usage

```go
import (
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/ledgerv3"
)

value := ledgerv3.AddAccountTypeRequestPersistenceNormal
```


## Values

| Name                                        | Value                                       |
| ------------------------------------------- | ------------------------------------------- |
| `AddAccountTypeRequestPersistenceNormal`    | NORMAL                                      |
| `AddAccountTypeRequestPersistenceEphemeral` | EPHEMERAL                                   |
| `AddAccountTypeRequestPersistenceTransient` | TRANSIENT                                   |