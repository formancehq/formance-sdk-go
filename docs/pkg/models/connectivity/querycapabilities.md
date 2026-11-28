# QueryCapabilities

The filterable keys per resource, serialized from the same allowlist
the server enforces -- it cannot drift from what queries accept.



## Fields

| Field                                                                                                               | Type                                                                                                                | Required                                                                                                            | Description                                                                                                         |
| ------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------- |
| `Resources`                                                                                                         | map[string]map[string][connectivity.QueryFieldCapability](../../../pkg/models/connectivity/queryfieldcapability.md) | :heavy_check_mark:                                                                                                  | N/A                                                                                                                 |