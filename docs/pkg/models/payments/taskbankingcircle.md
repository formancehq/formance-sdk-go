# TaskBankingCircle


## Fields

| Field                                                             | Type                                                              | Required                                                          | Description                                                       |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `ConnectorID`                                                     | `string`                                                          | :heavy_check_mark:                                                | N/A                                                               |
| `CreatedAt`                                                       | [time.Time](https://pkg.go.dev/time#Time)                         | :heavy_check_mark:                                                | N/A                                                               |
| `Descriptor`                                                      | [payments.Descriptor](../../../pkg/models/payments/descriptor.md) | :heavy_check_mark:                                                | N/A                                                               |
| `Error`                                                           | `*string`                                                         | :heavy_minus_sign:                                                | N/A                                                               |
| `ID`                                                              | `string`                                                          | :heavy_check_mark:                                                | N/A                                                               |
| `State`                                                           | [*payments.State](../../../pkg/models/payments/state.md)          | :heavy_minus_sign:                                                | N/A                                                               |
| `Status`                                                          | [payments.TaskStatus](../../../pkg/models/payments/taskstatus.md) | :heavy_check_mark:                                                | N/A                                                               |
| `UpdatedAt`                                                       | [time.Time](https://pkg.go.dev/time#Time)                         | :heavy_check_mark:                                                | N/A                                                               |