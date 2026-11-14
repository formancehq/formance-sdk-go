# SegmentTypeType

regex: custom regular expression (requires `regex`).
uuid: RFC 4122 UUID format.
uint64: decimal uint64.
bytes: hex-encoded bytes (even length).


## Example Usage

```go
import (
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/ledgerv3"
)

value := ledgerv3.SegmentTypeTypeRegex
```


## Values

| Name                    | Value                   |
| ----------------------- | ----------------------- |
| `SegmentTypeTypeRegex`  | regex                   |
| `SegmentTypeTypeUUID`   | uuid                    |
| `SegmentTypeTypeUint64` | uint64                  |
| `SegmentTypeTypeBytes`  | bytes                   |