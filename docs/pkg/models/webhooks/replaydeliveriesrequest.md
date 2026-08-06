# ReplayDeliveriesRequest


## Fields

| Field                                                           | Type                                                            | Required                                                        | Description                                                     |
| --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- |
| `ConfigIds`                                                     | []`string`                                                      | :heavy_minus_sign:                                              | N/A                                                             |
| `CreatedAtFrom`                                                 | [time.Time](https://pkg.go.dev/time#Time)                       | :heavy_check_mark:                                              | N/A                                                             |
| `CreatedAtTo`                                                   | [*time.Time](https://pkg.go.dev/time#Time)                      | :heavy_minus_sign:                                              | N/A                                                             |
| `Cursor`                                                        | `*string`                                                       | :heavy_minus_sign:                                              | N/A                                                             |
| `PageSize`                                                      | `*int64`                                                        | :heavy_minus_sign:                                              | N/A                                                             |
| `Statuses`                                                      | [][webhooks.Statuses](../../../pkg/models/webhooks/statuses.md) | :heavy_minus_sign:                                              | N/A                                                             |