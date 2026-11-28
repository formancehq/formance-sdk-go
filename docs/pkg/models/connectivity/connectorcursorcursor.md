# ConnectorCursorCursor


## Fields

| Field                                                                     | Type                                                                      | Required                                                                  | Description                                                               |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `Data`                                                                    | [][connectivity.Connector](../../../pkg/models/connectivity/connector.md) | :heavy_check_mark:                                                        | N/A                                                                       |
| `HasMore`                                                                 | `bool`                                                                    | :heavy_check_mark:                                                        | N/A                                                                       |
| `Next`                                                                    | `*string`                                                                 | :heavy_minus_sign:                                                        | Cursor of the next page; absent on the last page                          |
| `PageSize`                                                                | `int`                                                                     | :heavy_check_mark:                                                        | N/A                                                                       |