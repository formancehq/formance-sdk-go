# Stage


## Supported Types

### StageSend

```go
stage := shared.CreateStageStageSend(orchestration.StageSend{/* values here */})
```

### StageDelay

```go
stage := shared.CreateStageStageDelay(orchestration.StageDelay{/* values here */})
```

### StageWaitEvent

```go
stage := shared.CreateStageStageWaitEvent(orchestration.StageWaitEvent{/* values here */})
```

### Update

```go
stage := shared.CreateStageUpdate(orchestration.Update{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch stage.Type {
	case shared.StageTypeStageSend:
		// stage.StageSend is populated
	case shared.StageTypeStageDelay:
		// stage.StageDelay is populated
	case shared.StageTypeStageWaitEvent:
		// stage.StageWaitEvent is populated
	case shared.StageTypeUpdate:
		// stage.Update is populated
}
```
