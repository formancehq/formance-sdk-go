# V2WalletWithBalances

A wallet together with the funds it currently holds


## Fields

| Field                                                                   | Type                                                                    | Required                                                                | Description                                                             |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `Balances`                                                              | [orchestration.Balances](../../../pkg/models/orchestration/balances.md) | :heavy_check_mark:                                                      | Aggregated funds held by the wallet, keyed by asset                     |
| `CreatedAt`                                                             | [time.Time](https://pkg.go.dev/time#Time)                               | :heavy_check_mark:                                                      | When the wallet was created                                             |
| `ID`                                                                    | `string`                                                                | :heavy_check_mark:                                                      | The unique ID of the wallet.                                            |
| `Ledger`                                                                | `string`                                                                | :heavy_check_mark:                                                      | Name of the ledger backing this wallet                                  |
| `Metadata`                                                              | map[string]`string`                                                     | :heavy_check_mark:                                                      | Metadata associated with the wallet.                                    |
| `Name`                                                                  | `string`                                                                | :heavy_check_mark:                                                      | Human-readable name of the wallet                                       |