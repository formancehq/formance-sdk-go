# Type

Type of transfer initiation:
- TRANSFER: Internal to internal account transfer
- PAYOUT: Internal to external account payout


## Example Usage

```go
import (
	"github.com/formancehq/formance-sdk-go/v4/pkg/models/orchestration"
)

value := orchestration.TypeTransfer
```


## Values

| Name           | Value          |
| -------------- | -------------- |
| `TypeTransfer` | TRANSFER       |
| `TypePayout`   | PAYOUT         |