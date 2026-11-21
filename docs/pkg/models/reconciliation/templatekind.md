# TemplateKind

Which built-in check a rule applies, determining the shape of its templateSpec

## Example Usage

```go
import (
	"github.com/formancehq/formance-sdk-go/v5/pkg/models/reconciliation"
)

value := reconciliation.TemplateKindLedgerVsPoolDrift
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `TemplateKindLedgerVsPoolDrift` | ledger_vs_pool_drift            |
| `TemplateKindLedgerInvariant`   | ledger_invariant                |
| `TemplateKindAccountThreshold`  | account_threshold               |
| `TemplateKindSourceParity`      | source_parity                   |