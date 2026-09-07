# V3Pool

A named group of accounts whose balances are aggregated together


## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `CreatedAt`                                                                | [time.Time](https://pkg.go.dev/time#Time)                                  | :heavy_check_mark:                                                         | When the pool was created                                                  |
| `ID`                                                                       | `string`                                                                   | :heavy_check_mark:                                                         | Unique identifier of the pool                                              |
| `Name`                                                                     | `string`                                                                   | :heavy_check_mark:                                                         | Human-readable name of the pool                                            |
| `PoolAccounts`                                                             | []`string`                                                                 | :heavy_check_mark:                                                         | Accounts currently in the pool                                             |
| `Query`                                                                    | map[string]`any`                                                           | :heavy_minus_sign:                                                         | Filter selecting the accounts a dynamic pool contains                      |
| `Type`                                                                     | [*payments.V3PoolTypeEnum](../../../pkg/models/payments/v3pooltypeenum.md) | :heavy_minus_sign:                                                         | Whether a pool holds a fixed account list or is driven by a query          |