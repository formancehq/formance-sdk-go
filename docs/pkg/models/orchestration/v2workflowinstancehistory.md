# V2WorkflowInstanceHistory


## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `Error`                                                               | `*string`                                                             | :heavy_minus_sign:                                                    | Why the stage failed, absent when it succeeded                        |
| `Input`                                                               | [orchestration.V2Stage](../../../pkg/models/orchestration/v2stage.md) | :heavy_check_mark:                                                    | One step of a workflow, whose shape depends on the stage type         |
| `Name`                                                                | `string`                                                              | :heavy_check_mark:                                                    | Name of the stage this history entry records                          |
| `StartedAt`                                                           | [time.Time](https://pkg.go.dev/time#Time)                             | :heavy_check_mark:                                                    | When the stage started                                                |
| `Terminated`                                                          | `bool`                                                                | :heavy_check_mark:                                                    | Whether the stage has finished                                        |
| `TerminatedAt`                                                        | [*time.Time](https://pkg.go.dev/time#Time)                            | :heavy_minus_sign:                                                    | When the stage finished, absent while it is still running             |