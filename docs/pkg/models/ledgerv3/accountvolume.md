# AccountVolume


## Fields

| Field                                                                             | Type                                                                              | Required                                                                          | Description                                                                       |
| --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- |
| `Asset`                                                                           | `string`                                                                          | :heavy_check_mark:                                                                | Asset identifier (e.g. "USD/2")                                                   |
| `Color`                                                                           | `string`                                                                          | :heavy_check_mark:                                                                | Color of the bucket. Empty string is the uncolored bucket.                        |
| `Volumes`                                                                         | [ledgerv3.VolumesWithBalance](../../../pkg/models/ledgerv3/volumeswithbalance.md) | :heavy_check_mark:                                                                | N/A                                                                               |