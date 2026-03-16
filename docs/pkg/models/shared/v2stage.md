# V2Stage


## Supported Types

### V2StageSend

```go
v2Stage := shared.CreateV2StageV2StageSend(shared.V2StageSend{/* values here */})
```

### V2StageDelay

```go
v2Stage := shared.CreateV2StageV2StageDelay(shared.V2StageDelay{/* values here */})
```

### V2StageWaitEvent

```go
v2Stage := shared.CreateV2StageV2StageWaitEvent(shared.V2StageWaitEvent{/* values here */})
```

### V2Update

```go
v2Stage := shared.CreateV2StageV2Update(shared.V2Update{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch v2Stage.Type {
	case shared.V2StageTypeV2StageSend:
		// v2Stage.V2StageSend is populated
	case shared.V2StageTypeV2StageDelay:
		// v2Stage.V2StageDelay is populated
	case shared.V2StageTypeV2StageWaitEvent:
		// v2Stage.V2StageWaitEvent is populated
	case shared.V2StageTypeV2Update:
		// v2Stage.V2Update is populated
}
```
