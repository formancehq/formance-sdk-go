# TaskResponseData


## Supported Types

### TaskStripe

```go
taskResponseData := shared.CreateTaskResponseDataTaskStripe(payments.TaskStripe{/* values here */})
```

### TaskWise

```go
taskResponseData := shared.CreateTaskResponseDataTaskWise(payments.TaskWise{/* values here */})
```

### TaskCurrencyCloud

```go
taskResponseData := shared.CreateTaskResponseDataTaskCurrencyCloud(payments.TaskCurrencyCloud{/* values here */})
```

### TaskDummyPay

```go
taskResponseData := shared.CreateTaskResponseDataTaskDummyPay(payments.TaskDummyPay{/* values here */})
```

### TaskModulr

```go
taskResponseData := shared.CreateTaskResponseDataTaskModulr(payments.TaskModulr{/* values here */})
```

### TaskBankingCircle

```go
taskResponseData := shared.CreateTaskResponseDataTaskBankingCircle(payments.TaskBankingCircle{/* values here */})
```

### TaskMangoPay

```go
taskResponseData := shared.CreateTaskResponseDataTaskMangoPay(payments.TaskMangoPay{/* values here */})
```

### TaskMoneycorp

```go
taskResponseData := shared.CreateTaskResponseDataTaskMoneycorp(payments.TaskMoneycorp{/* values here */})
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
