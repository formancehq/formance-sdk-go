# V3ConnectorBase

Summary of a connector, without its configuration


## Fields

| Field                                             | Type                                              | Required                                          | Description                                       |
| ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- |
| `CreatedAt`                                       | [*time.Time](https://pkg.go.dev/time#Time)        | :heavy_minus_sign:                                | When the connector was installed                  |
| `ID`                                              | `*string`                                         | :heavy_minus_sign:                                | Unique identifier of the connector                |
| `Name`                                            | `*string`                                         | :heavy_minus_sign:                                | Human-readable name of the connector instance     |
| `Provider`                                        | `*string`                                         | :heavy_minus_sign:                                | Name of the payment provider behind the connector |
| `Reference`                                       | `*string`                                         | :heavy_minus_sign:                                | Stable reference identifying the connector        |