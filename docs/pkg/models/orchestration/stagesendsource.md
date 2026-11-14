# StageSendSource

Where a send stage takes the funds from


## Fields

| Field                                                                                                | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `Account`                                                                                            | [*orchestration.StageSendSourceAccount](../../../pkg/models/orchestration/stagesendsourceaccount.md) | :heavy_minus_sign:                                                                                   | Take the funds from a ledger account                                                                 |
| `Payment`                                                                                            | [*orchestration.StageSendSourcePayment](../../../pkg/models/orchestration/stagesendsourcepayment.md) | :heavy_minus_sign:                                                                                   | Take the funds from a payment                                                                        |
| `Wallet`                                                                                             | [*orchestration.StageSendSourceWallet](../../../pkg/models/orchestration/stagesendsourcewallet.md)   | :heavy_minus_sign:                                                                                   | Take the funds from a wallet                                                                         |