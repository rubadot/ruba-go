<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	rubago "github.com/Rubadot/ruba-go"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := rubago.New(
		rubago.WithSecurity(os.Getenv("RUBA_ACCESS_TOKEN")),
	)

	res, err := s.Organizations.List(ctx, nil, rubago.Pointer[int64](1), rubago.Pointer[int64](10), nil)
	if err != nil {
		log.Fatal(err)
	}
	if res.ListResourceOrganization != nil {
		for {
			// handle items

			res, err = res.Next()

			if err != nil {
				// handle error
			}

			if res == nil {
				break
			}
		}
	}
}

```
<!-- End SDK Example Usage [usage] -->