# V3Instance


## Fields

| Field                                                   | Type                                                    | Required                                                | Description                                             |
| ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- |
| `ConnectorID`                                           | `string`                                                | :heavy_check_mark:                                      | Identifier of the connector this run belongs to         |
| `CreatedAt`                                             | [time.Time](https://pkg.go.dev/time#Time)               | :heavy_check_mark:                                      | When the run started                                    |
| `Error`                                                 | `*string`                                               | :heavy_minus_sign:                                      | Why the run failed, absent when it succeeded            |
| `ID`                                                    | `string`                                                | :heavy_check_mark:                                      | Unique identifier of the run                            |
| `ScheduleID`                                            | `string`                                                | :heavy_check_mark:                                      | Identifier of the schedule that started this run        |
| `Terminated`                                            | `bool`                                                  | :heavy_check_mark:                                      | Whether the run has finished, successfully or not       |
| `TerminatedAt`                                          | [*time.Time](https://pkg.go.dev/time#Time)              | :heavy_minus_sign:                                      | When the run finished, absent while it is still running |
| `UpdatedAt`                                             | [time.Time](https://pkg.go.dev/time#Time)               | :heavy_check_mark:                                      | When the run was last updated                           |