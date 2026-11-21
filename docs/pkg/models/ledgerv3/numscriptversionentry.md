# NumscriptVersionEntry


## Fields

| Field                                          | Type                                           | Required                                       | Description                                    | Example                                        |
| ---------------------------------------------- | ---------------------------------------------- | ---------------------------------------------- | ---------------------------------------------- | ---------------------------------------------- |
| `CreatedAt`                                    | [*time.Time](https://pkg.go.dev/time#Time)     | :heavy_minus_sign:                             | When this version was stored (ISO 8601 format) |                                                |
| `Version`                                      | `string`                                       | :heavy_check_mark:                             | Stored semver version string (e.g. "1.0.0")    | 1.0.0                                          |