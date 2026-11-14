# DeliveryStatus

Current state of a delivery in its retry lifecycle

## Example Usage

```go
import (
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/webhooks"
)

value := webhooks.DeliveryStatusPending
```


## Values

| Name                       | Value                      |
| -------------------------- | -------------------------- |
| `DeliveryStatusPending`    | pending                    |
| `DeliveryStatusDelivering` | delivering                 |
| `DeliveryStatusSucceeded`  | succeeded                  |
| `DeliveryStatusFailed`     | failed                     |
| `DeliveryStatusCancelled`  | cancelled                  |