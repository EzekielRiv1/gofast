---
id: routing
title: Routing
---

Gofast routes are explicit Go handlers.

```go
app.Get("/account", func(ctx *gofast.Context) gofast.Page {
	return gofast.Page{
		Title: "Account",
		Body:  gofast.HTML("<h1>Your account</h1>"),
	}
})
```

`Get` accepts only `GET` requests. Other methods receive `405 Method Not Allowed`.

Routes currently match exact paths. This keeps behavior predictable while the framework is small. If a path is not registered, Gofast renders the configured not found page.

Customize the not found page with `NotFound`:

```go
app.NotFound(func(*gofast.Context) gofast.Page {
	return gofast.Page{
		Title:  "Missing page",
		Body:   gofast.HTML("<h1>Missing page</h1>"),
		Status: 404,
	}
})
```
