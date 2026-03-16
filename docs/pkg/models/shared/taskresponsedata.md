# TaskResponseData


## Supported Types

### TaskStripe

```go
taskResponseData := shared.CreateTaskResponseDataTaskStripe(shared.TaskStripe{/* values here */})
```

### TaskWise

```go
taskResponseData := shared.CreateTaskResponseDataTaskWise(shared.TaskWise{/* values here */})
```

### TaskCurrencyCloud

```go
taskResponseData := shared.CreateTaskResponseDataTaskCurrencyCloud(shared.TaskCurrencyCloud{/* values here */})
```

### TaskDummyPay

```go
taskResponseData := shared.CreateTaskResponseDataTaskDummyPay(shared.TaskDummyPay{/* values here */})
```

### TaskModulr

```go
taskResponseData := shared.CreateTaskResponseDataTaskModulr(shared.TaskModulr{/* values here */})
```

### TaskBankingCircle

```go
taskResponseData := shared.CreateTaskResponseDataTaskBankingCircle(shared.TaskBankingCircle{/* values here */})
```

### TaskMangoPay

```go
taskResponseData := shared.CreateTaskResponseDataTaskMangoPay(shared.TaskMangoPay{/* values here */})
```

### TaskMoneycorp

```go
taskResponseData := shared.CreateTaskResponseDataTaskMoneycorp(shared.TaskMoneycorp{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch taskResponseData.Type {
	case shared.TaskResponseDataTypeTaskStripe:
		// taskResponseData.TaskStripe is populated
	case shared.TaskResponseDataTypeTaskWise:
		// taskResponseData.TaskWise is populated
	case shared.TaskResponseDataTypeTaskCurrencyCloud:
		// taskResponseData.TaskCurrencyCloud is populated
	case shared.TaskResponseDataTypeTaskDummyPay:
		// taskResponseData.TaskDummyPay is populated
	case shared.TaskResponseDataTypeTaskModulr:
		// taskResponseData.TaskModulr is populated
	case shared.TaskResponseDataTypeTaskBankingCircle:
		// taskResponseData.TaskBankingCircle is populated
	case shared.TaskResponseDataTypeTaskMangoPay:
		// taskResponseData.TaskMangoPay is populated
	case shared.TaskResponseDataTypeTaskMoneycorp:
		// taskResponseData.TaskMoneycorp is populated
}
```
