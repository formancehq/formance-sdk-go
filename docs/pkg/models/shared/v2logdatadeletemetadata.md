# V2LogDataDeleteMetadata

Payload for DELETE_METADATA log entries. Contains the target entity and the metadata key that was deleted.


## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `Key`                                                         | `string`                                                      | :heavy_check_mark:                                            | The metadata key that was deleted                             |
| `TargetID`                                                    | [shared.TargetID](../../../pkg/models/shared/targetid.md)     | :heavy_check_mark:                                            | N/A                                                           |
| `TargetType`                                                  | [shared.TargetType](../../../pkg/models/shared/targettype.md) | :heavy_check_mark:                                            | Type of the target entity                                     |