# MetadataValue

A typed metadata value (string, integer, boolean, or null). Integer numbers are decoded exactly, including int64 bounds and values beyond 2^53. Integral decimal and exponent spellings are accepted; fractions and out-of-range values are rejected. String values must not contain NUL bytes. A value is limited to 16384 bytes by default; the effective ceiling comes from the cluster policy (`--metadata-max-value-bytes`). Exceeding it returns 400 with reason `METADATA_LIMIT_EXCEEDED` for a new or expired idempotency key. A live retained key with matching content returns its historical outcome; different content returns 409 with reason `IDEMPOTENCY_KEY_CONFLICT`. Requests without a key retain the limit error. Fresh policy refusals are audited without business effects; keyed refusals replay while retained, even if the policy is later loosened.


## Supported Types

### 

```go
metadataValue := shared.CreateMetadataValueStr(string{/* values here */})
```

### 

```go
metadataValue := shared.CreateMetadataValueInteger(int64{/* values here */})
```

### 

```go
metadataValue := shared.CreateMetadataValueBoolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch metadataValue.Type {
	case shared.MetadataValueTypeStr:
		// metadataValue.Str is populated
	case shared.MetadataValueTypeInteger:
		// metadataValue.Integer is populated
	case shared.MetadataValueTypeBoolean:
		// metadataValue.Boolean is populated
}
```
