<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	"github.com/formancehq/formance-sdk-go/v5"
	"log"
)

func main() {
	ctx := context.Background()

	s := v5.New()

	res, err := s.GetVersions(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res.GetVersionsResponse != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->