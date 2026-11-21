# AuditItem

One order within a proposal.


## Fields

| Field                                                                                    | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `LogSequence`                                                                            | `*int64`                                                                                 | :heavy_minus_sign:                                                                       | Sequence of the log produced by this order (0 for idempotent replay or failed proposals) |
| `OrderIndex`                                                                             | `*int64`                                                                                 | :heavy_minus_sign:                                                                       | Zero-based position of this order within the proposal                                    |
| `SerializedOrder`                                                                        | `*string`                                                                                | :heavy_minus_sign:                                                                       | Hex-encoded deterministic serialized bytes of the order at apply time                    |