# CreatedTransactionRuleRead

Rule scoped to MirrorCreatedTransaction. Inside `match`, `log` is the created-transaction proto directly (e.g. `log.metadata["k"]`).



## Fields

| Field                                                                                                   | Type                                                                                                    | Required                                                                                                | Description                                                                                             |
| ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- |
| `Actions`                                                                                               | [][ledgerv3.CreatedTransactionActionRead](../../../pkg/models/ledgerv3/createdtransactionactionread.md) | :heavy_minus_sign:                                                                                      | N/A                                                                                                     |
| `Match`                                                                                                 | `*string`                                                                                               | :heavy_minus_sign:                                                                                      | N/A                                                                                                     |