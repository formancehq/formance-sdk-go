# V2LogDataDeleteMetadata

Payload for DELETE_METADATA log entries. Contains the target entity and the metadata key that was deleted.


## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `Key`                                                         | `string`                                                      | :heavy_check_mark:                                            | The metadata key that was deleted                             |
| `TargetID`                                                    | [ledger.TargetID](../../../pkg/models/ledger/targetid.md)     | :heavy_check_mark:                                            | N/A                                                           |
| `TargetType`                                                  | [ledger.TargetType](../../../pkg/models/ledger/targettype.md) | :heavy_check_mark:                                            | Type of the target entity                                     |