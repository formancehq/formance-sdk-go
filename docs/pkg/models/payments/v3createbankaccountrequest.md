# V3CreateBankAccountRequest


## Fields

| Field                                                               | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `AccountNumber`                                                     | `*string`                                                           | :heavy_minus_sign:                                                  | Domestic account number. Supply this or an IBAN                     |
| `Country`                                                           | `*string`                                                           | :heavy_minus_sign:                                                  | Country the account is held in, as an ISO 3166-1 alpha-2 code       |
| `Iban`                                                              | `*string`                                                           | :heavy_minus_sign:                                                  | International bank account number. Supply this or an account number |
| `Metadata`                                                          | map[string]`string`                                                 | :heavy_minus_sign:                                                  | Arbitrary key/value pairs attached to the resource                  |
| `Name`                                                              | `string`                                                            | :heavy_check_mark:                                                  | Human-readable name for the bank account                            |
| `SwiftBicCode`                                                      | `*string`                                                           | :heavy_minus_sign:                                                  | SWIFT/BIC code identifying the bank                                 |