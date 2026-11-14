# V2BulkElementRevertTransactionData


## Fields

| Field                                              | Type                                               | Required                                           | Description                                        | Example                                            |
| -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- |
| `AtEffectiveDate`                                  | `*bool`                                            | :heavy_minus_sign:                                 | N/A                                                |                                                    |
| `Force`                                            | `*bool`                                            | :heavy_minus_sign:                                 | N/A                                                |                                                    |
| `ID`                                               | [*big.Int](https://pkg.go.dev/math/big#Int)        | :heavy_check_mark:                                 | N/A                                                |                                                    |
| `Metadata`                                         | map[string]`string`                                | :heavy_minus_sign:                                 | Arbitrary key/value pairs attached to the resource | {<br/>"admin": "true"<br/>}                        |