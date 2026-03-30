# Hold


## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `Subject`                                                              | [*orchestration.Subject](../../../pkg/models/orchestration/subject.md) | :heavy_minus_sign:                                                     | N/A                                                                    |
| `Description`                                                          | `string`                                                               | :heavy_check_mark:                                                     | N/A                                                                    |
| `ID`                                                                   | `string`                                                               | :heavy_check_mark:                                                     | The unique ID of the hold.                                             |
| `Metadata`                                                             | map[string]`string`                                                    | :heavy_check_mark:                                                     | Metadata associated with the hold.                                     |
| `WalletID`                                                             | `string`                                                               | :heavy_check_mark:                                                     | The ID of the wallet the hold is associated with.                      |