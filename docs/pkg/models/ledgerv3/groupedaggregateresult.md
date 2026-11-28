# GroupedAggregateResult


## Fields

| Field                                                                           | Type                                                                            | Required                                                                        | Description                                                                     |
| ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- |
| `Prefix`                                                                        | `string`                                                                        | :heavy_check_mark:                                                              | Account prefix this group matches                                               |
| `Volumes`                                                                       | [][ledgerv3.AggregatedVolume](../../../pkg/models/ledgerv3/aggregatedvolume.md) | :heavy_check_mark:                                                              | Aggregated volumes for accounts matching this prefix                            |