# SetMetadataAction

`value` is a literal metadata string; `value_expr` is a CEL expression typed against the current variant and evaluated at commit time (e.g. `log.reference`). Exactly one of the two should be set; if neither is set the value defaults to the empty string. Optional `type` coerces the produced string into a typed MetadataValue via the platform conversion matrix — one of "string" (default), "int64", "bool", "uint64", "int8"/"int16"/"int32", "uint8"/"uint16"/"uint32", "datetime".



## Fields

| Field              | Type               | Required           | Description        |
| ------------------ | ------------------ | ------------------ | ------------------ |
| `Key`              | `string`           | :heavy_check_mark: | N/A                |
| `Type`             | `*string`          | :heavy_minus_sign: | N/A                |
| `Value`            | `*string`          | :heavy_minus_sign: | N/A                |
| `ValueExpr`        | `*string`          | :heavy_minus_sign: | N/A                |