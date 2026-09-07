# AccountType

Whether an account is internal to the provider or belongs to an external party

## Example Usage

```go
import (
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/payments"
)

value := payments.AccountTypeUnknown
```


## Values

| Name                  | Value                 |
| --------------------- | --------------------- |
| `AccountTypeUnknown`  | UNKNOWN               |
| `AccountTypeInternal` | INTERNAL              |
| `AccountTypeExternal` | EXTERNAL              |