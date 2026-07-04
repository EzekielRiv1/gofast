package main

import (
	"fmt"
	"html/template"
	"log"
	"time"

	"github.com/EzekielRiv1/gofast"
)

func main() {
	app := gofast.New().WithLayout(gofast.MustLayout(layout))

	app.Get("/", func(*gofast.Context) gofast.Page {
		return gofast.Page{
			Title: "Gofast demo",
			Body:  page("Home", "This page was rendered in Go. Try moving around without a full reload."),
		}
	})

	app.Get("/counter", func(*gofast.Context) gofast.Page {
		return gofast.Page{
			Title: "Counter",
			Body:  page("Counter", fmt.Sprintf("Server time: %s", time.Now().Format(time.Kitchen))),
		}
	})

	log.Println("listening on http://localhost:8080")
	log.Fatal(app.ListenAndServe(":8080"))
}

func page(title, text string) template.HTML {
	return gofast.HTML(fmt.Sprintf(`<nav>
  <a href="/">Home</a>
  <a href="/counter">Counter</a>
</nav>
<section>
  <h1>%s</h1>
  <p>%s</p>
  <p><a href="/counter">Refresh the counter page</a></p>
</section>`, template.HTMLEscapeString(title), template.HTMLEscapeString(text)))
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
