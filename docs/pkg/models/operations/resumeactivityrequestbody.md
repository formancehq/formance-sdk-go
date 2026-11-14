# ResumeActivityRequestBody


## Fields

| Field                                                                                          | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `PausedAt`                                                                                     | [time.Time](https://pkg.go.dev/time#Time)                                                      | :heavy_check_mark:                                                                             | Nonzero RFC3339 timestamp from the paused activity projection identifying the pause generation |
| `TemporalRunID`                                                                                | `string`                                                                                       | :heavy_check_mark:                                                                             | Expected Temporal run id from the paused activity projection                                   |