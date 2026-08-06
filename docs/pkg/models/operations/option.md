# Options

## Global Options

Global options are passed when initializing the SDK client and apply to all operations.

### WithServerURL

WithServerURL allows providing an alternative server URL.

```go
v5.WithServerURL("https://api.example.com")
```

### WithTemplatedServerURL

WithTemplatedServerURL allows providing an alternative server URL with templated parameters.

```go
v5.WithTemplatedServerURL("https://{host}:{port}", map[string]string{
    "host": "api.example.com",
    "port": "8080",
})
```

### WithServerIndex

WithServerIndex allows the overriding of the default server by index.

```go
v5.WithServerIndex(1)
```

### WithEnvironment

WithEnvironment allows setting the environment variable for url substitution.

```go
v5.WithEnvironment(/* ... */)
```

### WithOrganization

WithOrganization allows setting the organization variable for url substitution.

```go
v5.WithOrganization(/* ... */)
```

### WithClient

WithClient allows the overriding of the default HTTP client used by the SDK.

```go
v5.WithClient(httpClient)
```

### WithSecurity

WithSecurity configures the SDK to use the provided security details.

```go
v5.WithSecurity(/* ... */)
```

### WithSecuritySource

WithSecuritySource configures the SDK to invoke the provided function on each method call to determine authentication.

```go
v5.WithSecuritySource(/* ... */)
```

### WithRetryConfig

WithRetryConfig allows setting the default retry configuration used by the SDK for all supported operations.

```go
v5.WithRetryConfig(retry.Config{
    Strategy: "backoff",
    Backoff: retry.BackoffStrategy{
        InitialInterval: 500 * time.Millisecond,
        MaxInterval: 60 * time.Second,
        Exponent: 1.5,
        MaxElapsedTime: 5 * time.Minute,
    },
    RetryConnectionErrors: true,
})
```

### WithTimeout

WithTimeout sets the default request timeout for all operations.

```go
v5.WithTimeout(30 * time.Second)
```

## Per-Method Options

Per-method options are passed as the last argument to individual methods and override any global settings for that request.

### WithServerURL

WithServerURL allows providing an alternative server URL for a single request.

```go
operations.WithServerURL("http://api.example.com")
```

### WithTemplatedServerURL

WithTemplatedServerURL allows providing an alternative server URL with templated parameters for a single request.

```go
operations.WithTemplatedServerURL("http://{host}:{port}", map[string]string{
    "host": "api.example.com",
    "port": "8080",
})
```

### WithRetries

WithRetries allows customizing the default retry configuration for a single request.

```go
operations.WithRetries(retry.Config{
    Strategy: "backoff",
    Backoff: retry.BackoffStrategy{
        InitialInterval: 500 * time.Millisecond,
        MaxInterval: 60 * time.Second,
        Exponent: 1.5,
        MaxElapsedTime: 5 * time.Minute,
    },
    RetryConnectionErrors: true,
})
```

### WithOperationTimeout

WithOperationTimeout allows setting the request timeout for a single request.

```go
operations.WithOperationTimeout(30 * time.Second)
```

### WithSetHeaders

WithSetHeaders allows setting custom headers on a per-request basis. If the request already contains headers matching the provided keys, they will be overwritten.

```go
operations.WithSetHeaders(map[string]string{
    "X-Cache-TTL": "60",
})
```

### WithURLOverride

WithURLOverride allows overriding the default URL for an operation.

```go
operations.WithURLOverride("/custom/path")
```

### WithAcceptHeaderOverride

WithAcceptHeaderOverride allows overriding the `Accept` header for operations that support multiple response content types.

```go
operations.WithAcceptHeaderOverride(operations.AcceptHeaderEnumApplicationJson)
```