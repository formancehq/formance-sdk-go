# Mode

Ledger mode. MIRROR creates a read-only ledger that syncs from an external v2 source. Values are case-sensitive; any other value is rejected with 400 INVALID_REQUEST.

## Example Usage

```go
import (
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/ledgerv3"
)

value := ledgerv3.ModeNormal
```


## Values

| Name         | Value        |
| ------------ | ------------ |
| `ModeNormal` | NORMAL       |
| `ModeMirror` | MIRROR       |