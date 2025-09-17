<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	"github.com/formancehq/formance-sdk-go/v3"
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/shared"
	"log"
)

func main() {
	ctx := context.Background()

	s := v3.New(
		v3.WithSecurity(shared.Security{
			ClientID:     v3.Pointer("<YOUR_CLIENT_ID_HERE>"),
			ClientSecret: v3.Pointer("<YOUR_CLIENT_SECRET_HERE>"),
		}),
	)

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