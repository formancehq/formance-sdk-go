# MigrationInfo


## Fields

| Field                                                | Type                                                 | Required                                             | Description                                          | Example                                              |
| ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- |
| `Date`                                               | [*time.Time](https://pkg.go.dev/time#Time)           | :heavy_minus_sign:                                   | When the migration was applied                       |                                                      |
| `Name`                                               | `*string`                                            | :heavy_minus_sign:                                   | Name of the migration                                | migrations:001                                       |
| `State`                                              | [*ledger.State](../../../pkg/models/ledger/state.md) | :heavy_minus_sign:                                   | Current state of the migration                       |                                                      |
| `Version`                                            | `*string`                                            | :heavy_minus_sign:                                   | Sequence number of the migration                     | 11                                                   |