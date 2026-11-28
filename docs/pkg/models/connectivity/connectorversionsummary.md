# ConnectorVersionSummary


## Fields

| Field                                                   | Type                                                    | Required                                                | Description                                             |
| ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- |
| `Digest`                                                | `*string`                                               | :heavy_minus_sign:                                      | Image digest pinning this version immutably, when known |
| `ReleaseDate`                                           | [*time.Time](https://pkg.go.dev/time#Time)              | :heavy_minus_sign:                                      | N/A                                                     |
| `Version`                                               | `string`                                                | :heavy_check_mark:                                      | Version tag (e.g. `v1.4.2`)                             |