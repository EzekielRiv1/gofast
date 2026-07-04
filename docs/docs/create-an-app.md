---
id: create-an-app
title: Create an App
---

A minimal app creates a `gofast.App`, registers routes, and starts the HTTP server.

```go
package main

import "github.com/EzekielRiv1/gofast"

func main() {
	app := gofast.New()

	app.Get("/", func(*gofast.Context) gofast.Page {
		return gofast.Page{
			Title: "Home",
			Body:  gofast.HTML("<h1>Hello from Go</h1>"),
		}
	})

	_ = app.ListenAndServe(":8080")
}
```

Run it and open `http://localhost:8080`.

For a complete example, see `examples/basic`.
