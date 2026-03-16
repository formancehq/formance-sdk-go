# TasksCursorData


## Supported Types

### TaskStripe

```go
tasksCursorData := shared.CreateTasksCursorDataTaskStripe(shared.TaskStripe{/* values here */})
```

### TaskWise

```go
tasksCursorData := shared.CreateTasksCursorDataTaskWise(shared.TaskWise{/* values here */})
```

### TaskCurrencyCloud

```go
tasksCursorData := shared.CreateTasksCursorDataTaskCurrencyCloud(shared.TaskCurrencyCloud{/* values here */})
```

### TaskDummyPay

```go
tasksCursorData := shared.CreateTasksCursorDataTaskDummyPay(shared.TaskDummyPay{/* values here */})
```

### TaskModulr

```go
tasksCursorData := shared.CreateTasksCursorDataTaskModulr(shared.TaskModulr{/* values here */})
```

### TaskBankingCircle

```go
tasksCursorData := shared.CreateTasksCursorDataTaskBankingCircle(shared.TaskBankingCircle{/* values here */})
```

### TaskMangoPay

```go
tasksCursorData := shared.CreateTasksCursorDataTaskMangoPay(shared.TaskMangoPay{/* values here */})
```

### TaskMoneycorp

```go
tasksCursorData := shared.CreateTasksCursorDataTaskMoneycorp(shared.TaskMoneycorp{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch tasksCursorData.Type {
	case shared.TasksCursorDataTypeTaskStripe:
		// tasksCursorData.TaskStripe is populated
	case shared.TasksCursorDataTypeTaskWise:
		// tasksCursorData.TaskWise is populated
	case shared.TasksCursorDataTypeTaskCurrencyCloud:
		// tasksCursorData.TaskCurrencyCloud is populated
	case shared.TasksCursorDataTypeTaskDummyPay:
		// tasksCursorData.TaskDummyPay is populated
	case shared.TasksCursorDataTypeTaskModulr:
		// tasksCursorData.TaskModulr is populated
	case shared.TasksCursorDataTypeTaskBankingCircle:
		// tasksCursorData.TaskBankingCircle is populated
	case shared.TasksCursorDataTypeTaskMangoPay:
		// tasksCursorData.TaskMangoPay is populated
	case shared.TasksCursorDataTypeTaskMoneycorp:
		// tasksCursorData.TaskMoneycorp is populated
}
```
