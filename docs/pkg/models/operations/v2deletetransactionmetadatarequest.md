# V2DeleteTransactionMetadataRequest


## Fields

| Field                                       | Type                                        | Required                                    | Description                                 | Example                                     |
| ------------------------------------------- | ------------------------------------------- | ------------------------------------------- | ------------------------------------------- | ------------------------------------------- |
| `IdempotencyKey`                            | **string*                                   | :heavy_minus_sign:                          | Use an idempotency key                      |                                             |
| `ID`                                        | [*big.Int](https://pkg.go.dev/math/big#Int) | :heavy_check_mark:                          | Transaction ID.                             | 1234                                        |
| `Key`                                       | *string*                                    | :heavy_check_mark:                          | The key to remove.                          | foo                                         |
| `Ledger`                                    | *string*                                    | :heavy_check_mark:                          | Name of the ledger.                         | ledger001                                   |