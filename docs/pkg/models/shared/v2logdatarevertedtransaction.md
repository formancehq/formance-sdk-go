# V2LogDataRevertedTransaction

Payload for REVERTED_TRANSACTION log entries. Contains both the original reverted transaction and the new reverting transaction.


## Fields

| Field                                                                     | Type                                                                      | Required                                                                  | Description                                                               |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `RevertedTransaction`                                                     | [shared.V2LogTransaction](../../../pkg/models/shared/v2logtransaction.md) | :heavy_check_mark:                                                        | Transaction structure as it appears in log payloads                       |
| `Transaction`                                                             | [shared.V2LogTransaction](../../../pkg/models/shared/v2logtransaction.md) | :heavy_check_mark:                                                        | Transaction structure as it appears in log payloads                       |