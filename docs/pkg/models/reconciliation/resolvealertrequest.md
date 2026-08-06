# ResolveAlertRequest

Mark an alert resolved. When `transactionRefs` is non-empty the
resolution kind is recorded as `fixed_by_booking`.



## Fields

| Field              | Type               | Required           | Description        |
| ------------------ | ------------------ | ------------------ | ------------------ |
| `By`               | `string`           | :heavy_check_mark: | N/A                |
| `Note`             | `*string`          | :heavy_minus_sign: | N/A                |
| `TransactionRefs`  | []`string`         | :heavy_minus_sign: | N/A                |