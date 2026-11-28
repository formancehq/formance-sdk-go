# Payload

Named system payload variants are preserved as supplied, including
nested fields and future sibling variants. Only apply is typed here;
other variants remain open-ended (EN-2781).



## Fields

| Field                                                    | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `AdditionalProperties`                                   | map[string]`any`                                         | :heavy_minus_sign:                                       | N/A                                                      |
| `Apply`                                                  | [*ledgerv3.Apply](../../../pkg/models/ledgerv3/apply.md) | :heavy_minus_sign:                                       | N/A                                                      |