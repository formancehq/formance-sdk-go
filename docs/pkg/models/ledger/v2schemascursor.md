# V2SchemasCursor

Paginated cursor over the ledger's schemas


## Fields

| Field                                                       | Type                                                        | Required                                                    | Description                                                 | Example                                                     |
| ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- |
| `Data`                                                      | [][ledger.V2Schema](../../../pkg/models/ledger/v2schema.md) | :heavy_check_mark:                                          | The schemas on this page                                    |                                                             |
| `HasMore`                                                   | `bool`                                                      | :heavy_check_mark:                                          | Whether further pages are available                         |                                                             |
| `Next`                                                      | `*string`                                                   | :heavy_minus_sign:                                          | Cursor for the next page, absent on the last page           | aW0gdmVuaWFtLCBxdWlzIG5vc3RydWQ=                            |
| `PageSize`                                                  | `int64`                                                     | :heavy_check_mark:                                          | Number of items requested per page                          |                                                             |
| `Previous`                                                  | `*string`                                                   | :heavy_minus_sign:                                          | Cursor for the previous page, absent on the first page      | YXVsdCBhbmQgYSBtYXhpbXVtIG1heF9yZXN1bHRzLol=                |