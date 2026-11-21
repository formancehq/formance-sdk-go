# PreparedQueryTarget

Target entity type. `LOGS`, `ACCOUNTS` and `TRANSACTIONS` are all usable prepared-query targets (create, list/get and execute).

## Example Usage

```go
import (
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/ledgerv3"
)

value := ledgerv3.PreparedQueryTargetAccounts
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `PreparedQueryTargetAccounts`     | ACCOUNTS                          |
| `PreparedQueryTargetTransactions` | TRANSACTIONS                      |
| `PreparedQueryTargetLogs`         | LOGS                              |