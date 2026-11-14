# GetVersionsResponse


## Fields

| Field                                                       | Type                                                        | Required                                                    | Description                                                 |
| ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- |
| `Env`                                                       | `string`                                                    | :heavy_check_mark:                                          | Environment the stack is deployed in                        |
| `Region`                                                    | `string`                                                    | :heavy_check_mark:                                          | Region the stack is deployed in                             |
| `Versions`                                                  | [][gateway.Version](../../../pkg/models/gateway/version.md) | :heavy_check_mark:                                          | Version and health of each component in the stack           |