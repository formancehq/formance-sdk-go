# Hold


## Fields

| Field                                                    | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `Asset`                                                  | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `Description`                                            | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `Destination`                                            | [*shared.Subject](../../../pkg/models/shared/subject.md) | :heavy_minus_sign:                                       | N/A                                                      |
| `ID`                                                     | *string*                                                 | :heavy_check_mark:                                       | The unique ID of the hold.                               |
| `Metadata`                                               | map[string]*string*                                      | :heavy_check_mark:                                       | Metadata associated with the hold.                       |
| `WalletID`                                               | *string*                                                 | :heavy_check_mark:                                       | The ID of the wallet the hold is associated with.        |