# WalletSubject


## Fields

| Field                                                                            | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `Balance`                                                                        | `*string`                                                                        | :heavy_minus_sign:                                                               | Name of the balance within the wallet. Defaults to the main balance when omitted |
| `Identifier`                                                                     | `string`                                                                         | :heavy_check_mark:                                                               | Identifier of the wallet                                                         |
| `Type`                                                                           | `string`                                                                         | :heavy_check_mark:                                                               | Discriminator identifying this subject as a wallet                               |