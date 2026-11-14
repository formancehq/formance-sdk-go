# Secret

A client secret. The full value is returned only at creation and cannot be retrieved again


## Fields

| Field                                              | Type                                               | Required                                           | Description                                        |
| -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- |
| `Clear`                                            | `string`                                           | :heavy_check_mark:                                 | N/A                                                |
| `ID`                                               | `string`                                           | :heavy_check_mark:                                 | N/A                                                |
| `LastDigits`                                       | `string`                                           | :heavy_check_mark:                                 | N/A                                                |
| `Metadata`                                         | map[string]`string`                                | :heavy_minus_sign:                                 | Arbitrary key/value pairs attached to the resource |
| `Name`                                             | `string`                                           | :heavy_check_mark:                                 | Human-readable name identifying the secret         |