# AnyVariantRuleRead

Rule scoped to every variant. `log` is the full MirrorLogEntry — use `has(log.<variant>)` guards to read variant-specific fields.



## Fields

| Field                                                                                   | Type                                                                                    | Required                                                                                | Description                                                                             |
| --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- |
| `Actions`                                                                               | [][ledgerv3.AnyVariantActionRead](../../../pkg/models/ledgerv3/anyvariantactionread.md) | :heavy_minus_sign:                                                                      | N/A                                                                                     |
| `Match`                                                                                 | `*string`                                                                               | :heavy_minus_sign:                                                                      | N/A                                                                                     |