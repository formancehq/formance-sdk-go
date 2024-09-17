<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	formancesdkgo "github.com/formancehq/formance-sdk-go/v3"
	"github.com/formancehq/formance-sdk-go/v3/pkg/models/shared"
	"log"
)

func main() {
	s := formancesdkgo.New(
		formancesdkgo.WithSecurity(shared.Security{
			ClientID:     "<YOUR_CLIENT_ID_HERE>",
			ClientSecret: "<YOUR_CLIENT_SECRET_HERE>",
		}),
	)

	ctx := context.Background()
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