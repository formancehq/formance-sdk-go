# V3RemoveEventsSinkRequest


## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `IdempotencyKey`                                              | `*string`                                                     | :heavy_minus_sign:                                            | Idempotency key for the request                               |
| `ControllerID`                                                | `*string`                                                     | :heavy_minus_sign:                                            | Optional opaque owner identity used as a removal precondition |
| `SinkName`                                                    | `string`                                                      | :heavy_check_mark:                                            | Name of the event sink to remove                              |