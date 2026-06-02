# TasksCursorData


## Supported Types

### TaskStripe

```go
tasksCursorData := shared.CreateTasksCursorDataTaskStripe(payments.TaskStripe{/* values here */})
```

### TaskWise

```go
tasksCursorData := shared.CreateTasksCursorDataTaskWise(payments.TaskWise{/* values here */})
```

### TaskCurrencyCloud

```go
tasksCursorData := shared.CreateTasksCursorDataTaskCurrencyCloud(payments.TaskCurrencyCloud{/* values here */})
```

### TaskDummyPay

```go
tasksCursorData := shared.CreateTasksCursorDataTaskDummyPay(payments.TaskDummyPay{/* values here */})
```

### TaskModulr

```go
tasksCursorData := shared.CreateTasksCursorDataTaskModulr(payments.TaskModulr{/* values here */})
```

### TaskBankingCircle

```go
tasksCursorData := shared.CreateTasksCursorDataTaskBankingCircle(payments.TaskBankingCircle{/* values here */})
```

### TaskMangoPay

```go
tasksCursorData := shared.CreateTasksCursorDataTaskMangoPay(payments.TaskMangoPay{/* values here */})
```

### TaskMoneycorp

```go
tasksCursorData := shared.CreateTasksCursorDataTaskMoneycorp(payments.TaskMoneycorp{/* values here */})
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
