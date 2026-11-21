# Action

The type of operation to perform. The REST bulk endpoint accepts
these four actions only; account-type actions
(`ADD_ACCOUNT_TYPE` / `REMOVE_ACCOUNT_TYPE`) are gRPC-only and
are rejected here.


## Example Usage

```go
import (
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/ledgerv3"
)

value := ledgerv3.ActionCreateTransaction
```


## Values

| Name                      | Value                     |
| ------------------------- | ------------------------- |
| `ActionCreateTransaction` | CREATE_TRANSACTION        |
| `ActionAddMetadata`       | ADD_METADATA              |
| `ActionRevertTransaction` | REVERT_TRANSACTION        |
| `ActionDeleteMetadata`    | DELETE_METADATA           |