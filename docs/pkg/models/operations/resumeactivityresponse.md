# ResumeActivityResponse


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ContentType`                                                      | `string`                                                           | :heavy_check_mark:                                                 | HTTP response content type for this operation                      |
| `Error`                                                            | [*orchestration.Error](../../../pkg/models/orchestration/error.md) | :heavy_minus_sign:                                                 | General error                                                      |
| `StatusCode`                                                       | `int`                                                              | :heavy_check_mark:                                                 | HTTP response status code for this operation                       |
| `RawResponse`                                                      | [*http.Response](https://pkg.go.dev/net/http#Response)             | :heavy_check_mark:                                                 | Raw HTTP response; suitable for custom response parsing            |