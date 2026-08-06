# ReplayDeliveriesResult


## Fields

| Field                                     | Type                                      | Required                                  | Description                               |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| `CreatedAtTo`                             | [time.Time](https://pkg.go.dev/time#Time) | :heavy_check_mark:                        | N/A                                       |
| `Expedited`                               | `int64`                                   | :heavy_check_mark:                        | N/A                                       |
| `HasMore`                                 | `bool`                                    | :heavy_check_mark:                        | N/A                                       |
| `NextCursor`                              | `*string`                                 | :heavy_minus_sign:                        | N/A                                       |
| `Replayed`                                | `int64`                                   | :heavy_check_mark:                        | N/A                                       |
| `Skipped`                                 | `int64`                                   | :heavy_check_mark:                        | N/A                                       |