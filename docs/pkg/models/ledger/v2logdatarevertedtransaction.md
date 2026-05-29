# V2LogDataRevertedTransaction

Payload for REVERTED_TRANSACTION log entries. Contains both the original reverted transaction and the new reverting transaction.


## Fields

| Field                                                                     | Type                                                                      | Required                                                                  | Description                                                               |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `V2LogTransaction`                                                        | [ledger.V2LogTransaction](../../../pkg/models/ledger/v2logtransaction.md) | :heavy_check_mark:                                                        | Transaction structure as it appears in log payloads                       |
| `V2LogTransaction1`                                                       | [ledger.V2LogTransaction](../../../pkg/models/ledger/v2logtransaction.md) | :heavy_check_mark:                                                        | Transaction structure as it appears in log payloads                       |