# Ruba for Go

Use typed clients to connect Go services to Ruba payments, billing, customers, subscriptions, orders, and webhooks.

## Add it to your project

```bash
go get github.com/Rubadot/ruba-go
```

## Make a request

```go
package main

import (
	"context"
	"log"
	"os"

	rubago "github.com/Rubadot/ruba-go"
)

func main() {
	client := rubago.New(
		rubago.WithSecurity(os.Getenv("RUBA_ACCESS_TOKEN")),
	)

	result, err := client.Organizations.List(
		context.Background(),
		nil,
		rubago.Pointer[int64](1),
		rubago.Pointer[int64](10),
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(result.ListResourceOrganization)
}
```

Set a custom server when you need the sandbox. Authentication uses a Ruba access token sent as an HTTP bearer token.

## Learn more

Start with the [Ruba documentation](https://docs.getruba.com), then use the [API reference](https://docs.getruba.com/api-reference/introduction) for endpoint details. Go SDK problems belong in this repository's [issue tracker](https://github.com/Rubadot/ruba-go/issues).

## License

MIT
