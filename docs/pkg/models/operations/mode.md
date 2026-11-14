# Mode

Inspection mode. The deprecated distinct-values alias for distinctValues remains accepted. An explicitly empty or unsupported value returns 400.

## Example Usage

```go
import (
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/operations"
)

value := operations.ModeSummary
```


## Values

| Name                 | Value                |
| -------------------- | -------------------- |
| `ModeSummary`        | summary              |
| `ModeDistinctValues` | distinctValues       |
| `ModeFacets`         | facets               |