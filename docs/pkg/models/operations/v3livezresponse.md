# V3LivezResponse


## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ContentType`                                                              | `string`                                                                   | :heavy_check_mark:                                                         | HTTP response content type for this operation                              |
| `HealthResponse`                                                           | [*ledgerv3.HealthResponse](../../../pkg/models/ledgerv3/healthresponse.md) | :heavy_minus_sign:                                                         | Process is alive                                                           |
| `StatusCode`                                                               | `int`                                                                      | :heavy_check_mark:                                                         | HTTP response status code for this operation                               |
| `RawResponse`                                                              | [*http.Response](https://pkg.go.dev/net/http#Response)                     | :heavy_check_mark:                                                         | Raw HTTP response; suitable for custom response parsing                    |