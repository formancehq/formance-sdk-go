# V2Hold

Funds locked by a pending wallet debit, later either confirmed or voided


## Fields

| Field                                                                            | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `Description`                                                                    | `string`                                                                         | :heavy_check_mark:                                                               | Human-readable reason the funds were held                                        |
| `Destination`                                                                    | [*orchestration.V2Subject](../../../pkg/models/orchestration/v2subject.md)       | :heavy_minus_sign:                                                               | The counterparty of a wallet movement, either a ledger account or another wallet |
| `ID`                                                                             | `string`                                                                         | :heavy_check_mark:                                                               | The unique ID of the hold.                                                       |
| `Metadata`                                                                       | map[string]`string`                                                              | :heavy_check_mark:                                                               | Metadata associated with the hold.                                               |
| `WalletID`                                                                       | `string`                                                                         | :heavy_check_mark:                                                               | The ID of the wallet the hold is associated with.                                |