# Info


## Fields

| Field                                                        | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `BuildDate`                                                  | `string`                                                     | :heavy_check_mark:                                           | Source commit timestamp (`org.opencontainers.image.created`) |
| `Commit`                                                     | `string`                                                     | :heavy_check_mark:                                           | Git commit SHA of the build                                  |
| `Version`                                                    | `string`                                                     | :heavy_check_mark:                                           | Build version of the API binary (the image tag)              |