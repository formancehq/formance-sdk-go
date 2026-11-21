# V2UpdateAccount

Metadata to set on a ledger account


## Fields

| Field                                   | Type                                    | Required                                | Description                             |
| --------------------------------------- | --------------------------------------- | --------------------------------------- | --------------------------------------- |
| `ID`                                    | `string`                                | :heavy_check_mark:                      | Address of the ledger account to update |
| `Ledger`                                | `string`                                | :heavy_check_mark:                      | Name of the ledger holding the account  |
| `Metadata`                              | map[string]`string`                     | :heavy_check_mark:                      | Metadata to set on the account          |