---
id: installation
title: Installation
---

Install Gofast in an existing Go module:

```bash
go get github.com/ezekielriv1/gofast@v0.1.0
```

## Requirements

- Go 1.26 or newer.
- A Go module for your application.

## Check the install

Create a small `main.go`:

```go
package main

import "github.com/ezekielriv1/gofast"

func main() {
	app := gofast.New()

	app.Get("/", func(ctx *gofast.Context) gofast.Page {
		return ctx.HTMLPage("Home", gofast.HTML("<h1>Gofast is running</h1>"))
	})

	_ = app.ListenAndServe(":8080")
}
```

Run it:

```bash
go run .
```

Open `http://localhost:8080`. You should see `Gofast is running`.

## Next step

Continue with [Create an App](create-an-app) to add route parameters, named URLs, and templates.
