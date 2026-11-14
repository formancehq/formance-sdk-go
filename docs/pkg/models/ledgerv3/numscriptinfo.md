# NumscriptInfo


## Fields

| Field                                       | Type                                        | Required                                    | Description                                 | Example                                     |
| ------------------------------------------- | ------------------------------------------- | ------------------------------------------- | ------------------------------------------- | ------------------------------------------- |
| `Content`                                   | `*string`                                   | :heavy_minus_sign:                          | The numscript source code                   |                                             |
| `CreatedAt`                                 | [*time.Time](https://pkg.go.dev/time#Time)  | :heavy_minus_sign:                          | Creation/update timestamp (ISO 8601 format) |                                             |
| `Ledger`                                    | `*string`                                   | :heavy_minus_sign:                          | Ledger name this numscript belongs to       |                                             |
| `Name`                                      | `string`                                    | :heavy_check_mark:                          | Name of the numscript                       |                                             |
| `Version`                                   | `*string`                                   | :heavy_minus_sign:                          | Semver version string (e.g. "1.0.0")        | 1.0.0                                       |