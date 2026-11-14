# InspectSummaryResponse


## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `Cardinality`                                                 | `int64`                                                       | :heavy_check_mark:                                            | Number of distinct values                                     |
| `EntitiesWithKey`                                             | `int64`                                                       | :heavy_check_mark:                                            | Number of entities that have this metadata key set (non-null) |
| `EntitiesWithNull`                                            | `int64`                                                       | :heavy_check_mark:                                            | Number of entities with a null value for this key             |
| `Max`                                                         | `any`                                                         | :heavy_minus_sign:                                            | Maximum value (null if index is empty)                        |
| `Min`                                                         | `any`                                                         | :heavy_minus_sign:                                            | Minimum value (null if index is empty)                        |