# StageDelay


## Fields

| Field                                                                                     | Type                                                                                      | Required                                                                                  | Description                                                                               |
| ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- |
| `Duration`                                                                                | `*string`                                                                                 | :heavy_minus_sign:                                                                        | Wait for this long, expressed as a Go duration such as 10s. Mutually exclusive with until |
| `Until`                                                                                   | [*time.Time](https://pkg.go.dev/time#Time)                                                | :heavy_minus_sign:                                                                        | Wait until this absolute point in time. Mutually exclusive with duration                  |