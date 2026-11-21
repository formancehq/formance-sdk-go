# V2ConfigInfo


## Fields

| Field                                                 | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ExperimentalFeatures`                                | []`string`                                            | :heavy_minus_sign:                                    | Experimental feature flags enabled on this deployment |
| `Server`                                              | `string`                                              | :heavy_check_mark:                                    | Name of the server serving the API                    |
| `Version`                                             | `string`                                              | :heavy_check_mark:                                    | Version of the ledger service                         |