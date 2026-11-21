# FacetDistribution

Distinct-connector counts per facet value, computed from the
namespaced `facet:value` tags of the connectors matching the query.



## Fields

| Field                                                               | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `Facets`                                                            | map[string]map[string]`int64`                                       | :heavy_check_mark:                                                  | Facet name to value-to-count map, e.g. `{"provider": {"psp": 6}}`.<br/> |
| `Total`                                                             | `int64`                                                             | :heavy_check_mark:                                                  | Published connectors matching the query.                            |