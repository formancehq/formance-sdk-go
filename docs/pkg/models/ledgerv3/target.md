# Target

Target entity type. `LOGS` queries return log entries; the log-only filter fields (`logId`, `date`, `ledger`) are valid only when the target is `LOGS`.

## Example Usage

```go
import (
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/ledgerv3"
)

value := ledgerv3.TargetAccounts
```


## Values

| Name                 | Value                |
| -------------------- | -------------------- |
| `TargetAccounts`     | ACCOUNTS             |
| `TargetTransactions` | TRANSACTIONS         |
| `TargetLogs`         | LOGS                 |