# Gofast

Gofast is a small Go-first framework for server-rendered applications that feel like single-page apps.

Go owns routing, rendering, layouts, and server behavior. The browser gets a tiny JavaScript layer that intercepts local links, asks the server for rendered fragments, and swaps the page body without a full reload.

## Features

- Explicit Go route handlers.
- Route parameters like `/projects/:projectID`.
- Named routes and URL generation.
- Server-rendered pages with Go `html/template`.
- SPA-style same-origin navigation with normal server fallback.

## Quick start

```go
package main

import "github.com/EzekielRiv1/gofast"

func main() {
	app := gofast.New()

	app.Get("/", func(ctx *gofast.Context) gofast.Page {
		return ctx.HTMLPage("Home", gofast.HTML("<h1>Hello from Go</h1>"))
	})

	_ = app.ListenAndServe(":8080")
}
```

Run the example:

```powershell
cd examples/basic
go run .
```

Then open `http://localhost:8080`.

## Documentation

The public documentation lives in `docs/` and is built with Docusaurus.
