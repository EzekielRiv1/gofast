---
id: routing
title: Routing
---

Gofast routes are explicit Go handlers.

```go
app.Get("/account", func(ctx *gofast.Context) gofast.Page {
	return ctx.HTMLPage("Account", gofast.HTML("<h1>Your account</h1>"))
})
```

`Get` accepts only `GET` requests. Other methods receive `405 Method Not Allowed`.

If a path is not registered, Gofast renders the configured not found page.

## Route parameters

Use `:name` segments when a route needs data from the URL.

```go
app.Get("/projects/:projectID", func(ctx *gofast.Context) gofast.Page {
	projectID := html.EscapeString(ctx.Param("projectID"))

	return ctx.HTMLPage("Project", gofast.HTML("<h1>Project "+projectID+"</h1>"))
})
```

Route parameters are matched one path segment at a time. `/projects/:projectID` matches `/projects/gofast`, but not `/projects/gofast/issues`.

This snippet uses `html.EscapeString` from Go's standard library. Escape dynamic route values when you place them into trusted HTML. For normal pages, prefer template rendering because Go's `html/template` escapes dynamic data by default.

## Named routes

Name a route when you want Gofast to generate links for you.

```go
app.Get("/projects/:projectID", showProject).Name("project.show")
```

Then build URLs with the route name and params:

```go
path, err := app.URL("project.show", gofast.Params{
	"projectID": "gofast",
})
```

Inside handlers, use the context helper:

```go
href := ctx.MustURL("project.show", gofast.Params{
	"projectID": ctx.Param("projectID"),
})
```

Generated URLs escape path parameters for you. If a required parameter is missing, `URL` returns an error and `MustURL` panics.

## Query values

Use `ctx.Query` for query string values.

```go
tab := ctx.Query("tab")
```

Query values are not part of route matching. They are extra request data available to your handler.

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
