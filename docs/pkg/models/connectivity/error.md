# Error

Request or server error. The HTTP status and code identify the failure.


## Fields

| Field                                                                               | Type                                                                                | Required                                                                            | Description                                                                         |
| ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- |
| `Details`                                                                           | `*string`                                                                           | :heavy_minus_sign:                                                                  | Additional details; when structured, this string contains a JSON object.            |
| `ErrorCode`                                                                         | `*string`                                                                           | :heavy_minus_sign:                                                                  | Machine-readable error code (e.g. `not_found`, `unauthorized`, `connector_missing`) |
| `ErrorMessage`                                                                      | `*string`                                                                           | :heavy_minus_sign:                                                                  | N/A                                                                                 |