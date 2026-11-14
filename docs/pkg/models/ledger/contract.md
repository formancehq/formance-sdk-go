# Contract


## Fields

| Field                                                | Type                                                 | Required                                             | Description                                          | Example                                              |
| ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- |
| `Account`                                            | `*string`                                            | :heavy_minus_sign:                                   | Account address, or pattern, the contract applies to | users:001                                            |
| `Expr`                                               | [ledger.Expr](../../../pkg/models/ledger/expr.md)    | :heavy_check_mark:                                   | Expression the account's balances must satisfy       |                                                      |