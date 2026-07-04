---
id: create-an-app
title: Create Your First App
---

This page builds a small Gofast app with one route, one route parameter, one generated URL, and one template.

## 1. Create the app

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

Run it and open `http://localhost:8080`.

## 2. Add a route parameter

```go
app.Get("/projects/:projectID", func(ctx *gofast.Context) gofast.Page {
	projectID := html.EscapeString(ctx.Param("projectID"))

	return ctx.HTMLPage("Project", gofast.HTML("<h1>Project "+projectID+"</h1>"))
}).Name("project.show")
```

Open `http://localhost:8080/projects/gofast`.

This snippet uses `html.EscapeString` from Go's standard library before placing the route value into trusted HTML.

## 3. Generate a URL

```go
href := app.MustURL("project.show", gofast.Params{
	"projectID": "gofast",
})
```

Use named routes for links instead of repeating path strings throughout your app.

## 4. Render a template

```go
app := gofast.New().WithViews(gofast.MustViews("views/*.html"))

app.Get("/", func(ctx *gofast.Context) gofast.Page {
	return ctx.Render("Home", "pages/home", map[string]string{
		"Name": "Gofast",
	})
})
```

```html
{{ define "pages/home" }}
  <h1>Hello, {{ .Name }}</h1>
{{ end }}
```

## What you learned

- `app.Get` registers a route.
- `:projectID` captures one URL segment.
- `.Name` lets Gofast generate URLs for that route.
- `ctx.Render` renders a named Go template into a page.

For a complete working version, see `examples/basic`.
