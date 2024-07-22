<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	formancesdkgo "github.com/formancehq/formance-sdk-go/v2"
	"github.com/formancehq/formance-sdk-go/v2/pkg/models/shared"
	"log"
	"os"
)

func main() {
	s := formancesdkgo.New(
		formancesdkgo.WithSecurity(shared.Security{
			Authorization: os.Getenv("AUTHORIZATION"),
		}),
	)

	ctx := context.Background()
	res, err := s.GetOIDCWellKnowns(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->