# SnoozeAlertRequest

Mute an alert's notifications until `until` (which must be in the future).


## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           | Example                                                               |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `By`                                                                  | `string`                                                              | :heavy_check_mark:                                                    | Who is snoozing the alert                                             | ops@buildr.com                                                        |
| `Note`                                                                | `*string`                                                             | :heavy_minus_sign:                                                    | Free-text note to record with the snooze                              |                                                                       |
| `Until`                                                               | [time.Time](https://pkg.go.dev/time#Time)                             | :heavy_check_mark:                                                    | When the mute expires and notifications resume. Must be in the future |                                                                       |