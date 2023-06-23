# PaymentsAccount


## Fields

| Field                                                             | Type                                                              | Required                                                          | Description                                                       |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `CreatedAt`                                                       | [time.Time](https://pkg.go.dev/time#Time)                         | :heavy_check_mark:                                                | N/A                                                               |
| `ID`                                                              | *string*                                                          | :heavy_check_mark:                                                | N/A                                                               |
| `Provider`                                                        | [Connector](../../models/shared/connector.md)                     | :heavy_check_mark:                                                | N/A                                                               |
| `Reference`                                                       | *string*                                                          | :heavy_check_mark:                                                | N/A                                                               |
| `Type`                                                            | [PaymentsAccountType](../../models/shared/paymentsaccounttype.md) | :heavy_check_mark:                                                | N/A                                                               |