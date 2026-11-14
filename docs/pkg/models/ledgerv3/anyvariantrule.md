# AnyVariantRule

Rule scoped to every variant. `log` is the full MirrorLogEntry — use `has(log.<variant>)` guards to read variant-specific fields.



## Fields

| Field                                                                           | Type                                                                            | Required                                                                        | Description                                                                     |
| ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- |
| `Actions`                                                                       | [][ledgerv3.AnyVariantAction](../../../pkg/models/ledgerv3/anyvariantaction.md) | :heavy_minus_sign:                                                              | N/A                                                                             |
| `Match`                                                                         | `*string`                                                                       | :heavy_minus_sign:                                                              | N/A                                                                             |