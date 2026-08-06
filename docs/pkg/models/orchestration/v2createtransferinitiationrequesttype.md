# V2CreateTransferInitiationRequestType

Type of transfer initiation:
- TRANSFER: Internal to internal account transfer
- PAYOUT: Internal to external account payout


## Example Usage

```go
import (
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/orchestration"
)

value := orchestration.V2CreateTransferInitiationRequestTypeTransfer
```


## Values

| Name                                            | Value                                           |
| ----------------------------------------------- | ----------------------------------------------- |
| `V2CreateTransferInitiationRequestTypeTransfer` | TRANSFER                                        |
| `V2CreateTransferInitiationRequestTypePayout`   | PAYOUT                                          |