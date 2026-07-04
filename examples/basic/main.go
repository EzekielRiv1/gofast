package main

import (
	"fmt"
	"log"
	"time"

	"github.com/EzekielRiv1/gofast"
)

func main() {
	app := gofast.New().
		WithLayout(gofast.MustLayout(layout)).
		WithViews(gofast.MustViews("views/*.html"))

	app.Get("/", func(ctx *gofast.Context) gofast.Page {
		return ctx.Render("Gofast demo", "pages/basic", map[string]string{
			"Title":      "Home",
			"Text":       "This page was rendered from a Go template. Try moving around without a full reload.",
			"HomeURL":    ctx.MustURL("home", nil),
			"CounterURL": ctx.MustURL("counter.show", gofast.Params{"name": "clock"}),
			"ProjectURL": ctx.MustURL("project.show", gofast.Params{"projectID": "gofast"}),
		})
	}).Name("home")

	app.Get("/counter/:name", func(ctx *gofast.Context) gofast.Page {
		return ctx.Render("Counter", "pages/basic", map[string]string{
			"Title":      "Counter: " + ctx.Param("name"),
			"Text":       fmt.Sprintf("Server time: %s", time.Now().Format(time.Kitchen)),
			"HomeURL":    ctx.MustURL("home", nil),
			"CounterURL": ctx.MustURL("counter.show", gofast.Params{"name": ctx.Param("name")}),
			"ProjectURL": ctx.MustURL("project.show", gofast.Params{"projectID": "gofast"}),
		})
	}).Name("counter.show")

	app.Get("/projects/:projectID", func(ctx *gofast.Context) gofast.Page {
		projectID := ctx.Param("projectID")
		return ctx.Render("Project", "pages/basic", map[string]string{
			"Title":      "Project: " + projectID,
			"Text":       "The project ID came from the URL route parameter.",
			"HomeURL":    ctx.MustURL("home", nil),
			"CounterURL": ctx.MustURL("counter.show", gofast.Params{"name": "clock"}),
			"ProjectURL": ctx.MustURL("project.show", gofast.Params{"projectID": projectID}),
		})
	}).Name("project.show")

	log.Println("listening on http://localhost:8080")
	log.Fatal(app.ListenAndServe(":8080"))
}

var layout = `<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>{{ .Title }}</title>
  <style>
    body { font-family: system-ui, sans-serif; margin: 0; color: #17202a; background: #f7f8fa; }
    main { max-width: 720px; margin: 48px auto; padding: 0 24px; }
    nav { display: flex; gap: 12px; margin-bottom: 32px; }
    a { color: #0b6bcb; }
    section { background: white; border: 1px solid #d9e0e8; border-radius: 8px; padding: 24px; }
  </style>
  <script type="module">` + gofast.BrowserScriptForLayout() + `</script>
</head>
<body>
  <main id="gofast-app">{{ .Body }}</main>
</body>
</html>`
