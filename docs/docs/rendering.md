---
id: rendering
title: Rendering
---

Handlers return a `gofast.Page`. A page has a title, body, and optional HTTP status.

```go
return ctx.HTMLPage("Dashboard", gofast.HTML("<h1>Dashboard</h1>"))
```

For normal browser requests, Gofast renders a full HTML document with a `main#gofast-app` element. For client-side navigation requests, it renders only the page body.

## Template rendering

Use `Views` when you want pages and partials rendered with Go's `html/template`.

```go
app := gofast.New().WithViews(gofast.MustViews("views/*.html"))

app.Get("/", func(ctx *gofast.Context) gofast.Page {
	return ctx.Render("Home", "pages/home", map[string]string{
		"Name": "Gofast",
	})
})
```

Define the template in `views/pages.html`:

```html
{{ define "pages/home" }}
  <h1>Hello, {{ .Name }}</h1>
{{ end }}
```

`ctx.Render` returns a `Page`. If the app has no view registry or the template fails to render, Gofast returns a `500` page instead of panicking during the request.

## Status codes

Set `Page.Status` when the response should use a specific HTTP status.

```go
return gofast.Page{
	Title:  "Created",
	Body:   gofast.HTML("<h1>Created</h1>"),
	Status: 201,
}
```

## Custom layouts

Use `WithLayout` when you want your own document shell:

```go
app := gofast.New().WithLayout(gofast.MustLayout(`<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>{{ .Title }}</title>
  <script type="module">` + gofast.BrowserScriptForLayout() + `</script>
</head>
<body>
  <main id="gofast-app">{{ .Body }}</main>
</body>
</html>`))
```

The layout receives the current `Page` as template data. Keep `main#gofast-app` in the document if you want SPA-style navigation to work.

## When to use raw HTML

Use `gofast.HTML` for small trusted fragments, examples, or output that has already been escaped. For normal pages, prefer `html/template` through `Views` so dynamic data is escaped by default.

## Common mistakes

- Do not pass unescaped user input to `gofast.HTML`.
- Keep `main id="gofast-app"` in custom layouts if you want browser navigation to update in place.
- Define templates with `{{ define "name" }}` when you plan to render them by name.
