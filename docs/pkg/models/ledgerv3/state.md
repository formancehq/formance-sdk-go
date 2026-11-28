# State

Current sync state (SYNCING = catching up, FOLLOWING = up to date)

## Example Usage

```go
import (
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/ledgerv3"
)

value := ledgerv3.StateSyncing
```


## Values

| Name             | Value            |
| ---------------- | ---------------- |
| `StateSyncing`   | SYNCING          |
| `StateFollowing` | FOLLOWING        |