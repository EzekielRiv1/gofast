---
id: rendering
title: Rendering
---

Handlers return a `gofast.Page`.

```go
return gofast.Page{
	Title: "Dashboard",
	Body:  gofast.HTML("<h1>Dashboard</h1>"),
}
```

For normal browser requests, Gofast renders a full HTML document with a `main#gofast-app` element. For client-side navigation requests, it renders only the page body.

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
