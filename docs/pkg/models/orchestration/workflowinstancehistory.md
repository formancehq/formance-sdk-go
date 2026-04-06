# WorkflowInstanceHistory


## Fields

| Field                                                             | Type                                                              | Required                                                          | Description                                                       |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `Stage`                                                           | [orchestration.Stage](../../../pkg/models/orchestration/stage.md) | :heavy_check_mark:                                                | N/A                                                               |
| `Error`                                                           | `*string`                                                         | :heavy_minus_sign:                                                | N/A                                                               |
| `Name`                                                            | `string`                                                          | :heavy_check_mark:                                                | N/A                                                               |
| `StartedAt`                                                       | [time.Time](https://pkg.go.dev/time#Time)                         | :heavy_check_mark:                                                | N/A                                                               |
| `Terminated`                                                      | `bool`                                                            | :heavy_check_mark:                                                | N/A                                                               |
| `TerminatedAt`                                                    | [*time.Time](https://pkg.go.dev/time#Time)                        | :heavy_minus_sign:                                                | N/A                                                               |