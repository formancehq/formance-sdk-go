# V2Exporter


## Fields

| Field                                          | Type                                           | Required                                       | Description                                    |
| ---------------------------------------------- | ---------------------------------------------- | ---------------------------------------------- | ---------------------------------------------- |
| `Config`                                       | map[string]`any`                               | :heavy_check_mark:                             | Driver-specific configuration for the exporter |
| `CreatedAt`                                    | [time.Time](https://pkg.go.dev/time#Time)      | :heavy_check_mark:                             | N/A                                            |
| `Driver`                                       | `string`                                       | :heavy_check_mark:                             | Name of the exporter driver to use             |
| `ID`                                           | `string`                                       | :heavy_check_mark:                             | N/A                                            |