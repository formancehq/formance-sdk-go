# V3InfoResponse


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ContentType`                                                      | `string`                                                           | :heavy_check_mark:                                                 | HTTP response content type for this operation                      |
| `ServerInfo`                                                       | [*ledgerv3.ServerInfo](../../../pkg/models/ledgerv3/serverinfo.md) | :heavy_minus_sign:                                                 | Server build information                                           |
| `StatusCode`                                                       | `int`                                                              | :heavy_check_mark:                                                 | HTTP response status code for this operation                       |
| `RawResponse`                                                      | [*http.Response](https://pkg.go.dev/net/http#Response)             | :heavy_check_mark:                                                 | Raw HTTP response; suitable for custom response parsing            |