# AccountRequest


## Fields

| Field                                                               | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `AccountName`                                                       | `*string`                                                           | :heavy_minus_sign:                                                  | N/A                                                                 |
| `ConnectorID`                                                       | `string`                                                            | :heavy_check_mark:                                                  | N/A                                                                 |
| `CreatedAt`                                                         | [time.Time](https://pkg.go.dev/time#Time)                           | :heavy_check_mark:                                                  | N/A                                                                 |
| `DefaultAsset`                                                      | `*string`                                                           | :heavy_minus_sign:                                                  | N/A                                                                 |
| `Metadata`                                                          | map[string]`string`                                                 | :heavy_minus_sign:                                                  | N/A                                                                 |
| `Reference`                                                         | `string`                                                            | :heavy_check_mark:                                                  | N/A                                                                 |
| `Type`                                                              | [payments.AccountType](../../../pkg/models/payments/accounttype.md) | :heavy_check_mark:                                                  | N/A                                                                 |