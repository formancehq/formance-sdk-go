# V2DeleteAccountMetadataRequest


## Fields

| Field                  | Type                   | Required               | Description            | Example                |
| ---------------------- | ---------------------- | ---------------------- | ---------------------- | ---------------------- |
| `IdempotencyKey`       | **string*              | :heavy_minus_sign:     | Use an idempotency key |                        |
| `Address`              | *string*               | :heavy_check_mark:     | Account address        |                        |
| `Key`                  | *string*               | :heavy_check_mark:     | The key to remove.     | foo                    |
| `Ledger`               | *string*               | :heavy_check_mark:     | Name of the ledger.    | ledger001              |