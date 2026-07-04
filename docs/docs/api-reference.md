---
id: api-reference
title: API Reference
---

This page summarizes the public APIs available today. For guided usage, start with [Create Your First App](create-an-app).

## App

### `gofast.New() *App`

Creates an app with the default layout and default not found page.

### `app.Get(path string, handler Handler) *Route`

Registers a GET route.

```go
app.Get("/projects/:projectID", showProject).Name("project.show")
```

### `app.Route(path string, handler Handler) *Route`

Registers a route without restricting the HTTP method.

### `app.WithLayout(layout Layout) *App`

Sets the document layout used for full page requests.

### `app.WithViews(views *Views) *App`

Sets the template registry used by `ctx.Render`.

### `app.URL(name string, params Params) (string, error)`

Builds a URL for a named route.

### `app.MustURL(name string, params Params) string`

Builds a URL for a named route and panics if the URL cannot be built.

### `app.NotFound(handler Handler)`

Sets the page rendered when no route matches.

### `app.ListenAndServe(addr string) error`

Starts an HTTP server for the app.

## Route

### `route.Name(name string) *Route`

Names a route for URL generation.

```go
app.Get("/projects/:projectID", showProject).Name("project.show")
```

## Context

### `ctx.Param(name string) string`

Returns a route parameter.

### `ctx.Query(name string) string`

Returns the first query string value for `name`.

### `ctx.Params() Params`

Returns a copy of all matched route parameters.

### `ctx.RequestContext() context.Context`

Returns the request context.

### `ctx.HTMLPage(title string, body template.HTML) Page`

Creates a page from trusted HTML.

### `ctx.Render(title, templateName string, data any) Page`

Renders a named template from the configured `Views` registry.

### `ctx.URL(name string, params Params) (string, error)`

Builds a URL using the current app.

### `ctx.MustURL(name string, params Params) string`

Builds a URL using the current app and panics on error.

## Views

### `gofast.MustViews(patterns ...string) *Views`

Parses templates from file patterns.

### `gofast.MustViewsFS(files fs.FS, patterns ...string) *Views`

Parses templates from an `fs.FS`.

### `views.Render(name string, data any) (template.HTML, error)`

Renders a named template.

## Params

```go
type Params map[string]string
```

Use `Params` when generating URLs for named routes.

```go
href := ctx.MustURL("project.show", gofast.Params{
	"projectID": "gofast",
})
```

## Page

```go
type Page struct {
	Title  string
	Body   template.HTML
	Status int
}
```

`Status` defaults to `200 OK` when left as `0`.

## Layout

### `gofast.MustLayout(markup string) Layout`

Parses a document layout or panics.

### `gofast.BrowserScriptForLayout() string`

Returns the browser navigation script for custom layouts.
