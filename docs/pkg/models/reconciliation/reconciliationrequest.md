# ReconciliationRequest


## Fields

| Field                                                 | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ReconciledAtLedger`                                  | [time.Time](https://pkg.go.dev/time#Time)             | :heavy_check_mark:                                    | Point in time at which the ledger side is evaluated   | 2021-01-01T00:00:00.000Z                              |
| `ReconciledAtPayments`                                | [time.Time](https://pkg.go.dev/time#Time)             | :heavy_check_mark:                                    | Point in time at which the payments side is evaluated | 2021-01-01T00:00:00.000Z                              |