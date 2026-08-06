# DeliveryAttempt


## Fields

| Field                                                       | Type                                                        | Required                                                    | Description                                                 |
| ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- |
| `AttemptNumber`                                             | `int64`                                                     | :heavy_check_mark:                                          | N/A                                                         |
| `CreatedAt`                                                 | [time.Time](https://pkg.go.dev/time#Time)                   | :heavy_check_mark:                                          | N/A                                                         |
| `DeliveryID`                                                | `string`                                                    | :heavy_check_mark:                                          | N/A                                                         |
| `DurationMillis`                                            | `*int64`                                                    | :heavy_minus_sign:                                          | N/A                                                         |
| `Endpoint`                                                  | `string`                                                    | :heavy_check_mark:                                          | N/A                                                         |
| `Error`                                                     | `*string`                                                   | :heavy_minus_sign:                                          | N/A                                                         |
| `ID`                                                        | `string`                                                    | :heavy_check_mark:                                          | N/A                                                         |
| `Outcome`                                                   | [webhooks.Outcome](../../../pkg/models/webhooks/outcome.md) | :heavy_check_mark:                                          | N/A                                                         |
| `ReplayGeneration`                                          | `int64`                                                     | :heavy_check_mark:                                          | N/A                                                         |
| `ResponseExcerpt`                                           | `*string`                                                   | :heavy_minus_sign:                                          | N/A                                                         |
| `StatusCode`                                                | `int64`                                                     | :heavy_check_mark:                                          | N/A                                                         |