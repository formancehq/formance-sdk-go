# Type

## Example Usage

```go
import (
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/ledgerv3"
)

value := ledgerv3.TypeNewTransaction
```


## Values

| Name                                | Value                               |
| ----------------------------------- | ----------------------------------- |
| `TypeNewTransaction`                | NEW_TRANSACTION                     |
| `TypeRevertedTransaction`           | REVERTED_TRANSACTION                |
| `TypeSetMetadata`                   | SET_METADATA                        |
| `TypeDeleteMetadata`                | DELETE_METADATA                     |
| `TypeSetMetadataFieldType`          | SET_METADATA_FIELD_TYPE             |
| `TypeRemovedMetadataFieldType`      | REMOVED_METADATA_FIELD_TYPE         |
| `TypeOrderSkipped`                  | ORDER_SKIPPED                       |
| `TypeFillGap`                       | FILL_GAP                            |
| `TypeCreateIndex`                   | CREATE_INDEX                        |
| `TypeDropIndex`                     | DROP_INDEX                          |
| `TypeAddedAccountType`              | ADDED_ACCOUNT_TYPE                  |
| `TypeRemovedAccountType`            | REMOVED_ACCOUNT_TYPE                |
| `TypeUpdatedDefaultEnforcementMode` | UPDATED_DEFAULT_ENFORCEMENT_MODE    |