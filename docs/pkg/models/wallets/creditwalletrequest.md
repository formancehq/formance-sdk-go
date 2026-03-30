# CreditWalletRequest


## Fields

| Field                                                       | Type                                                        | Required                                                    | Description                                                 |
| ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- |
| `Monetary`                                                  | [wallets.Monetary](../../../pkg/models/wallets/monetary.md) | :heavy_check_mark:                                          | N/A                                                         |
| `Balance`                                                   | `*string`                                                   | :heavy_minus_sign:                                          | The balance to credit                                       |
| `Metadata`                                                  | map[string]`string`                                         | :heavy_minus_sign:                                          | Metadata associated with the wallet.                        |
| `Reference`                                                 | `*string`                                                   | :heavy_minus_sign:                                          | N/A                                                         |
| `Sources`                                                   | [][wallets.Subject](../../../pkg/models/wallets/subject.md) | :heavy_minus_sign:                                          | N/A                                                         |
| `Timestamp`                                                 | [*time.Time](https://pkg.go.dev/time#Time)                  | :heavy_minus_sign:                                          | N/A                                                         |