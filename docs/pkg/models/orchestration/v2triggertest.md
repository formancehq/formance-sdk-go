# V2TriggerTest

Result of evaluating a trigger against a sample event, without running the workflow


## Fields

| Field                                                                                | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `Filter`                                                                             | [*orchestration.Filter](../../../pkg/models/orchestration/filter.md)                 | :heavy_minus_sign:                                                                   | How the trigger's filter evaluated against the sample event                          |
| `Variables`                                                                          | map[string][orchestration.Variables](../../../pkg/models/orchestration/variables.md) | :heavy_minus_sign:                                                                   | The variables the trigger would build from the sample event                          |