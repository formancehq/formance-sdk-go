# SnoozeAlertRequest

Mute an alert's notifications until `until` (which must be in the future).


## Fields

| Field                                     | Type                                      | Required                                  | Description                               | Example                                   |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| `By`                                      | `string`                                  | :heavy_check_mark:                        | N/A                                       | ops@buildr.com                            |
| `Note`                                    | `*string`                                 | :heavy_minus_sign:                        | N/A                                       |                                           |
| `Until`                                   | [time.Time](https://pkg.go.dev/time#Time) | :heavy_check_mark:                        | N/A                                       |                                           |