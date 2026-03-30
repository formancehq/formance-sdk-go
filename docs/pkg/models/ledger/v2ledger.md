# V2Ledger


## Fields

| Field                                      | Type                                       | Required                                   | Description                                | Example                                    |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| `V2Metadata`                               | map[string]`string`                        | :heavy_minus_sign:                         | N/A                                        | {<br/>"admin": "true"<br/>}                |
| `AddedAt`                                  | [time.Time](https://pkg.go.dev/time#Time)  | :heavy_check_mark:                         | N/A                                        |                                            |
| `Bucket`                                   | `string`                                   | :heavy_check_mark:                         | N/A                                        |                                            |
| `DeletedAt`                                | [*time.Time](https://pkg.go.dev/time#Time) | :heavy_minus_sign:                         | N/A                                        |                                            |
| `Features`                                 | map[string]`string`                        | :heavy_minus_sign:                         | N/A                                        |                                            |
| `ID`                                       | `*int64`                                   | :heavy_minus_sign:                         | N/A                                        |                                            |
| `Name`                                     | `string`                                   | :heavy_check_mark:                         | N/A                                        |                                            |