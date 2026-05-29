# WalletSummary


## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `AvailableFunds`                                                      | map[string][*big.Int](https://pkg.go.dev/math/big#Int)                | :heavy_check_mark:                                                    | N/A                                                                   |
| `Balances`                                                            | [][wallets.AssetHolder1](../../../pkg/models/wallets/assetholder1.md) | :heavy_check_mark:                                                    | N/A                                                                   |
| `ExpirableFunds`                                                      | map[string][*big.Int](https://pkg.go.dev/math/big#Int)                | :heavy_check_mark:                                                    | N/A                                                                   |
| `ExpiredFunds`                                                        | map[string][*big.Int](https://pkg.go.dev/math/big#Int)                | :heavy_check_mark:                                                    | N/A                                                                   |
| `HoldFunds`                                                           | map[string][*big.Int](https://pkg.go.dev/math/big#Int)                | :heavy_check_mark:                                                    | N/A                                                                   |