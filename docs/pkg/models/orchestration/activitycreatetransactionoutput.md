# ActivityCreateTransactionOutput

Result of the activity that writes a transaction to a ledger


## Fields

| Field                                                                         | Type                                                                          | Required                                                                      | Description                                                                   |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `Data`                                                                        | [orchestration.Transaction](../../../pkg/models/orchestration/transaction.md) | :heavy_check_mark:                                                            | A transaction recorded in the ledger, carrying its postings and metadata      |