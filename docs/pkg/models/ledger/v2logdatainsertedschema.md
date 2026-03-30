# V2LogDataInsertedSchema

Payload for INSERTED_SCHEMA log entries. Contains the schema that was inserted into the ledger.


## Fields

| Field                                                             | Type                                                              | Required                                                          | Description                                                       |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `V2SchemaData`                                                    | [ledger.V2SchemaData](../../../pkg/models/ledger/v2schemadata.md) | :heavy_check_mark:                                                | Complete schema structure with metadata                           |