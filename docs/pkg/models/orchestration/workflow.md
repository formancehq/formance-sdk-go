# Workflow

A workflow definition and the stages it runs


## Fields

| Field                                                                               | Type                                                                                | Required                                                                            | Description                                                                         |
| ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- |
| `Config`                                                                            | [orchestration.WorkflowConfig](../../../pkg/models/orchestration/workflowconfig.md) | :heavy_check_mark:                                                                  | The stages a workflow runs, in order                                                |
| `CreatedAt`                                                                         | [time.Time](https://pkg.go.dev/time#Time)                                           | :heavy_check_mark:                                                                  | When the workflow was created                                                       |
| `ID`                                                                                | `string`                                                                            | :heavy_check_mark:                                                                  | Unique identifier of the workflow                                                   |
| `UpdatedAt`                                                                         | [time.Time](https://pkg.go.dev/time#Time)                                           | :heavy_check_mark:                                                                  | When the workflow was last modified                                                 |