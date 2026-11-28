# Snooze

A time-boxed, operator-initiated mute of an alert's notifications. While
`until` is in the future the alert keeps failing and keeps counting
against period-green — only its notifications are suppressed.



## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           | Example                                                               |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `At`                                                                  | [time.Time](https://pkg.go.dev/time#Time)                             | :heavy_check_mark:                                                    | When the alert was snoozed                                            |                                                                       |
| `By`                                                                  | `string`                                                              | :heavy_check_mark:                                                    | Who snoozed the alert                                                 | ops@buildr.com                                                        |
| `Note`                                                                | `*string`                                                             | :heavy_minus_sign:                                                    | Free-text note left by whoever snoozed the alert                      |                                                                       |
| `Until`                                                               | [time.Time](https://pkg.go.dev/time#Time)                             | :heavy_check_mark:                                                    | When the mute expires and notifications resume. Must be in the future |                                                                       |