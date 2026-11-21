# CreatedTransactionRule

Rule scoped to MirrorCreatedTransaction. Inside `match`, `log` is the created-transaction proto directly (e.g. `log.metadata["k"]`).



## Fields

| Field                                                                                           | Type                                                                                            | Required                                                                                        | Description                                                                                     |
| ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- |
| `Actions`                                                                                       | [][ledgerv3.CreatedTransactionAction](../../../pkg/models/ledgerv3/createdtransactionaction.md) | :heavy_minus_sign:                                                                              | N/A                                                                                             |
| `Match`                                                                                         | `*string`                                                                                       | :heavy_minus_sign:                                                                              | N/A                                                                                             |