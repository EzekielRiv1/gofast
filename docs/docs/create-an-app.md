---
id: create-an-app
title: Create Your First App
---

This guide creates a small Gofast app with:

- one Go server file;
- one template file;
- one CSS file;
- one JavaScript file;
- one normal page route;
- one route with a URL parameter.

## Project structure

Start with this structure:

```text
myapp/
  go.mod
  main.go
  views/
    pages.html
  assets/
    app.css
    app.js
```

Gofast does not require this exact structure, but it is a good starting point:

| Path | Purpose |
| --- | --- |
| `main.go` | Creates the app, registers routes, serves assets, and starts the server |
| `views/pages.html` | Stores server-rendered page templates |
| `assets/app.css` | Stores page styles |
| `assets/app.js` | Stores page-specific browser behavior |

## 1. Install Gofast

From inside your app folder:

```bash
go mod init example.com/myapp
go get github.com/ezekielriv1/gofast@v0.1.0
```

## 2. Create `main.go`

```go
package main

import (
	"log"
	"net/http"

	"github.com/ezekielriv1/gofast"
)

func main() {
	app := gofast.New().
		WithLayout(gofast.MustLayout(layout)).
		WithViews(gofast.MustViews("views/*.html"))

	app.Get("/", func(ctx *gofast.Context) gofast.Page {
		return ctx.Render("Home", "pages/home", map[string]string{
			"Title":      "Home",
			"Message":    "This page was rendered by Go.",
			"HomeURL":    ctx.MustURL("home", nil),
			"ProjectURL": ctx.MustURL("project.show", gofast.Params{"projectID": "gofast"}),
		})
	}).Name("home")

	app.Get("/projects/:projectID", func(ctx *gofast.Context) gofast.Page {
		projectID := ctx.Param("projectID")

		return ctx.Render("Project", "pages/home", map[string]string{
			"Title":      "Project: " + projectID,
			"Message":    "The project ID came from the URL.",
			"HomeURL":    ctx.MustURL("home", nil),
			"ProjectURL": ctx.MustURL("project.show", gofast.Params{"projectID": projectID}),
		})
	}).Name("project.show")

	mux := http.NewServeMux()
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	mux.Handle("/", app)

	log.Println("listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

var layout = `<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>{{ .Title }}</title>
  <link rel="stylesheet" href="/assets/app.css">
  <script type="module">` + gofast.BrowserScriptForLayout() + `</script>
  <script type="module" src="/assets/app.js"></script>
</head>
<body>
  <main id="gofast-app">{{ .Body }}</main>
</body>
</html>`
```

The `layout` string is the full HTML document. Gofast uses it for normal page loads. During internal navigation, Gofast sends only the new page body and swaps it into `#gofast-app`.

The `mux` serves `/assets/app.css` and `/assets/app.js` with Go's standard library, then sends all other requests to Gofast.

## 3. Create `views/pages.html`

```html
{{ define "pages/home" }}
<nav class="nav">
  <a href="{{ .HomeURL }}">Home</a>
  <a href="{{ .ProjectURL }}">Project</a>
</nav>

<section class="page">
  <h1>{{ .Title }}</h1>
  <p>{{ .Message }}</p>
  <button data-action="hello">Say hello</button>
</section>
{{ end }}
```

Templates are rendered on the server with Go's `html/template`. Dynamic values like `{{ .Title }}` and `{{ .Message }}` are escaped by default.

## 4. Create `assets/app.css`

```css
body {
  margin: 0;
  font-family: system-ui, sans-serif;
  color: #17202a;
  background: #f7f8fa;
}

main {
  max-width: 720px;
  margin: 48px auto;
  padding: 0 24px;
}

.nav {
  display: flex;
  gap: 12px;
  margin-bottom: 24px;
}

.page {
  padding: 24px;
  border: 1px solid #d9e0e8;
  border-radius: 8px;
  background: white;
}
```

Use CSS the same way you would in a normal server-rendered Go app. Gofast does not require a CSS framework.

## 5. Create `assets/app.js`

```js
document.addEventListener("click", (event) => {
  const button = event.target.closest("[data-action='hello']");
  if (!button) return;

  button.textContent = "Hello from app.js";
});
```

Use this file for page-specific browser behavior. Gofast already provides the navigation script through `gofast.BrowserScriptForLayout()`, so you do not need to write your own link interception code.

## 6. Run the app

```bash
go run .
```

Open `http://localhost:8080`.

Try these checks:

- Click `Project`. The URL should change to `/projects/gofast`.
- Click `Home`. The URL should change back to `/`.
- Click `Say hello`. The button text should update.
- Reload either page directly. It should still render from the server.

## What each piece does

| Piece | What it controls |
| --- | --- |
| `app.Get("/", ...)` | The home route |
| `app.Get("/projects/:projectID", ...)` | A route with one URL parameter |
| `.Name("project.show")` | The name used to generate links |
| `ctx.MustURL(...)` | Builds links from route names |
| `ctx.Render(...)` | Renders a named template into a page |
| `gofast.BrowserScriptForLayout()` | Enables internal navigation without a full reload |
| `http.FileServer` | Serves your CSS and JavaScript files |

## Next steps

- Learn how [routing](routing) works.
- Learn how [rendering](rendering) works.
- Learn how [browser navigation](browser-layer) works.
