# TemporalStats


## Fields

| Field                                                               | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `FirstSeen`                                                         | [*time.Time](https://pkg.go.dev/time#Time)                          | :heavy_minus_sign:                                                  | N/A                                                                 |
| `LastSeen`                                                          | [*time.Time](https://pkg.go.dev/time#Time)                          | :heavy_minus_sign:                                                  | N/A                                                                 |
| `PeakHours`                                                         | [][ledgerv3.HourBucket](../../../pkg/models/ledgerv3/hourbucket.md) | :heavy_minus_sign:                                                  | N/A                                                                 |
| `TransactionsPerDay`                                                | `*float64`                                                          | :heavy_minus_sign:                                                  | N/A                                                                 |