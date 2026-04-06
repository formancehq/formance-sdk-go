# V2WorkflowInstanceHistory


## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `V2Stage`                                                             | [orchestration.V2Stage](../../../pkg/models/orchestration/v2stage.md) | :heavy_check_mark:                                                    | N/A                                                                   |
| `Error`                                                               | `*string`                                                             | :heavy_minus_sign:                                                    | N/A                                                                   |
| `Name`                                                                | `string`                                                              | :heavy_check_mark:                                                    | N/A                                                                   |
| `StartedAt`                                                           | [time.Time](https://pkg.go.dev/time#Time)                             | :heavy_check_mark:                                                    | N/A                                                                   |
| `Terminated`                                                          | `bool`                                                                | :heavy_check_mark:                                                    | N/A                                                                   |
| `TerminatedAt`                                                        | [*time.Time](https://pkg.go.dev/time#Time)                            | :heavy_minus_sign:                                                    | N/A                                                                   |