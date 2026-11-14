# Monetary

An amount together with the asset it is denominated in


## Fields

| Field                                       | Type                                        | Required                                    | Description                                 |
| ------------------------------------------- | ------------------------------------------- | ------------------------------------------- | ------------------------------------------- |
| `Amount`                                    | [*big.Int](https://pkg.go.dev/math/big#Int) | :heavy_check_mark:                          | The amount of the monetary value.           |
| `Asset`                                     | `string`                                    | :heavy_check_mark:                          | The asset of the monetary value.            |