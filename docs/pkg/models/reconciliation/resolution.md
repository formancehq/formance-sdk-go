# Resolution


## Fields

| Field                                                             | Type                                                              | Required                                                          | Description                                                       |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `At`                                                              | [time.Time](https://pkg.go.dev/time#Time)                         | :heavy_check_mark:                                                | N/A                                                               |
| `By`                                                              | `string`                                                          | :heavy_check_mark:                                                | N/A                                                               |
| `EvidenceSnapshot`                                                | map[string]`any`                                                  | :heavy_minus_sign:                                                | N/A                                                               |
| `Kind`                                                            | [reconciliation.Kind](../../../pkg/models/reconciliation/kind.md) | :heavy_check_mark:                                                | N/A                                                               |
| `Note`                                                            | `*string`                                                         | :heavy_minus_sign:                                                | N/A                                                               |
| `TransactionRefs`                                                 | []`string`                                                        | :heavy_minus_sign:                                                | N/A                                                               |