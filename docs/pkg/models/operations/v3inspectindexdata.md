# V3InspectIndexData


## Supported Types

### InspectSummaryResponse

```go
v3InspectIndexData := operations.CreateV3InspectIndexDataInspectSummaryResponse(ledgerv3.InspectSummaryResponse{/* values here */})
```

### InspectDistinctValuesResponse

```go
v3InspectIndexData := operations.CreateV3InspectIndexDataInspectDistinctValuesResponse(ledgerv3.InspectDistinctValuesResponse{/* values here */})
```

### InspectFacetsResponse

```go
v3InspectIndexData := operations.CreateV3InspectIndexDataInspectFacetsResponse(ledgerv3.InspectFacetsResponse{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch v3InspectIndexData.Type {
	case operations.V3InspectIndexDataTypeInspectSummaryResponse:
		// v3InspectIndexData.InspectSummaryResponse is populated
	case operations.V3InspectIndexDataTypeInspectDistinctValuesResponse:
		// v3InspectIndexData.InspectDistinctValuesResponse is populated
	case operations.V3InspectIndexDataTypeInspectFacetsResponse:
		// v3InspectIndexData.InspectFacetsResponse is populated
}
```
