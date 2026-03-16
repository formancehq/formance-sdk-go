# Stage


## Supported Types

### StageSend

```go
stage := shared.CreateStageStageSend(shared.StageSend{/* values here */})
```

### StageDelay

```go
stage := shared.CreateStageStageDelay(shared.StageDelay{/* values here */})
```

### StageWaitEvent

```go
stage := shared.CreateStageStageWaitEvent(shared.StageWaitEvent{/* values here */})
```

### Update

```go
stage := shared.CreateStageUpdate(shared.Update{/* values here */})
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
