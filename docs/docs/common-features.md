---
id: common-features
title: Common Features
---

## Shared navigation

Put shared navigation in your layout or in helper functions that return HTML.

```go
func shell(body template.HTML) template.HTML {
	return gofast.HTML(`<nav><a href="/">Home</a></nav>` + string(body))
}
```

## Server-generated pages

Keep application logic in Go. Query data, authorize requests, and render the resulting HTML in the route handler or a small view helper.

```go
app.Get("/projects", func(ctx *gofast.Context) gofast.Page {
	return gofast.Page{
		Title: "Projects",
		Body:  renderProjects(loadProjects(ctx.RequestContext())),
	}
})
```

## Custom status codes

Set `Page.Status` when a route needs a specific HTTP status.

```go
return gofast.Page{
	Title:  "Created",
	Body:   gofast.HTML("<h1>Created</h1>"),
	Status: 201,
}
```
