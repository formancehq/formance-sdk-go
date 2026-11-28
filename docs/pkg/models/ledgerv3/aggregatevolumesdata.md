# AggregateVolumesData


## Fields

| Field                                                                                       | Type                                                                                        | Required                                                                                    | Description                                                                                 |
| ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- |
| `Groups`                                                                                    | [][ledgerv3.GroupedAggregateResult](../../../pkg/models/ledgerv3/groupedaggregateresult.md) | :heavy_minus_sign:                                                                          | Per-prefix aggregated volumes (present when groupByPrefixes is set)                         |
| `Volumes`                                                                                   | [][ledgerv3.AggregatedVolume](../../../pkg/models/ledgerv3/aggregatedvolume.md)             | :heavy_check_mark:                                                                          | Flat aggregated volumes per asset                                                           |