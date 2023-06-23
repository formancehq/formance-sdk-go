# ExpandedDebitHold


## Fields

| Field                                             | Type                                              | Required                                          | Description                                       | Example                                           |
| ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- |
| `Description`                                     | *string*                                          | :heavy_check_mark:                                | N/A                                               |                                                   |
| `Destination`                                     | [*Subject](../../models/shared/subject.md)        | :heavy_minus_sign:                                | N/A                                               |                                                   |
| `ID`                                              | *string*                                          | :heavy_check_mark:                                | The unique ID of the hold.                        |                                                   |
| `Metadata`                                        | map[string]*string*                               | :heavy_check_mark:                                | Metadata associated with the hold.                |                                                   |
| `OriginalAmount`                                  | *int64*                                           | :heavy_check_mark:                                | Original amount on hold                           | 100                                               |
| `Remaining`                                       | *int64*                                           | :heavy_check_mark:                                | Remaining amount on hold                          | 10                                                |
| `WalletID`                                        | *string*                                          | :heavy_check_mark:                                | The ID of the wallet the hold is associated with. |                                                   |