# PaymentsCursorCursor

Paginated cursor wrapping the list of payments


## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   | Example                                                       |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `Data`                                                        | [][payments.Payment](../../../pkg/models/payments/payment.md) | :heavy_check_mark:                                            | N/A                                                           |                                                               |
| `HasMore`                                                     | `bool`                                                        | :heavy_check_mark:                                            | Whether further pages are available                           | false                                                         |
| `Next`                                                        | `*string`                                                     | :heavy_minus_sign:                                            | Cursor for the next page, absent on the last page             |                                                               |
| `PageSize`                                                    | `int64`                                                       | :heavy_check_mark:                                            | Number of items requested per page                            | 15                                                            |
| `Previous`                                                    | `*string`                                                     | :heavy_minus_sign:                                            | Cursor for the previous page, absent on the first page        | YXVsdCBhbmQgYSBtYXhpbXVtIG1heF9yZXN1bHRzLol=                  |