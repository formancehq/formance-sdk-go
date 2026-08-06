# Evidence

For PASS/FAIL: the per-fingerprint roster. Each PASS entry carries a compact `proof`: the observed balances that make the check hold, with self-describing keys — both sides where two are compared, e.g. {"ledger":"523500","pool":"-523500"} (drift), {"left":…,"right":…} (source_parity), {"positive":…,"negative":…} (ledger_invariant), or {"balance":…} (account_threshold). These reuse the FAIL balance keys, so under tolerance the residual (left−right) is visible without recomputation. Each FAIL entry carries the full `evidence` breakdown (balances, drift, tolerance, compiled CEL). Bounded by asset count (per-asset granularity; per-account is parked). An empty array means no outcomes were produced (no assets matched). For ERROR: an object describing the engine failure.


## Supported Types

### 

```go
evidence := shared.CreateEvidenceArrayOf1([]reconciliation.One{/* values here */})
```

### 

```go
evidence := shared.CreateEvidenceMapOfAny(map[string]any{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch evidence.Type {
	case shared.EvidenceTypeArrayOf1:
		// evidence.ArrayOf1 is populated
	case shared.EvidenceTypeMapOfAny:
		// evidence.MapOfAny is populated
}
```
