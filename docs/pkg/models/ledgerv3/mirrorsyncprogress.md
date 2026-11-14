# MirrorSyncProgress

Sync progress for MIRROR ledgers (populated at read time)


## Fields

| Field                                                                        | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `Cursor`                                                                     | `int64`                                                                      | :heavy_check_mark:                                                           | Last ingested v2 log ID                                                      |
| `Error`                                                                      | [*ledgerv3.MirrorSyncError](../../../pkg/models/ledgerv3/mirrorsyncerror.md) | :heavy_minus_sign:                                                           | Most recent mirror sync error (null if healthy)                              |
| `RemainingLogs`                                                              | `int64`                                                                      | :heavy_check_mark:                                                           | Number of remaining logs to sync (sourceLogCount - cursor)                   |
| `SourceLogCount`                                                             | `int64`                                                                      | :heavy_check_mark:                                                           | Latest known log ID in the v2 source                                         |
| `State`                                                                      | [ledgerv3.State](../../../pkg/models/ledgerv3/state.md)                      | :heavy_check_mark:                                                           | Current sync state (SYNCING = catching up, FOLLOWING = up to date)           |