# PreparedQueryFilterInput

A filter supplied to a prepared-query create/update request, accepted in EITHER of two interchangeable representations that parse to the same filter and are validated against the same per-target rules (EN-1511):

- the structured JSON QueryFilter object (see `QueryFilter`); or
- the textual filter grammar carried as a JSON string, e.g.


  `"metadata[status] == \"active\""`.


`null` represents a match-all query. An empty object and an empty or whitespace-only textual expression are invalid. Responses echo a stored non-null filter in the structured `QueryFilter` form and omit this field for match-all queries.


## Supported Types

### QueryFilter

```go
preparedQueryFilterInput := shared.CreatePreparedQueryFilterInputQueryFilter(ledgerv3.QueryFilter{/* values here */})
```

### 

```go
preparedQueryFilterInput := shared.CreatePreparedQueryFilterInputStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch preparedQueryFilterInput.Type {
	case shared.PreparedQueryFilterInputTypeQueryFilter:
		// preparedQueryFilterInput.QueryFilter is populated
	case shared.PreparedQueryFilterInputTypeStr:
		// preparedQueryFilterInput.Str is populated
}
```
