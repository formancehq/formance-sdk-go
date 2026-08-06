# Snooze

A time-boxed, operator-initiated mute of an alert's notifications. While
`until` is in the future the alert keeps failing and keeps counting
against period-green — only its notifications are suppressed.



## Fields

| Field                                     | Type                                      | Required                                  | Description                               | Example                                   |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| `At`                                      | [time.Time](https://pkg.go.dev/time#Time) | :heavy_check_mark:                        | N/A                                       |                                           |
| `By`                                      | `string`                                  | :heavy_check_mark:                        | N/A                                       | ops@buildr.com                            |
| `Note`                                    | `*string`                                 | :heavy_minus_sign:                        | N/A                                       |                                           |
| `Until`                                   | [time.Time](https://pkg.go.dev/time#Time) | :heavy_check_mark:                        | N/A                                       |                                           |