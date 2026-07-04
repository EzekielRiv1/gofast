---
id: intro
title: What is Gofast?
slug: /
---

Gofast is a reusable framework for building server-rendered Go applications that behave like single-page apps.

Your application remains a normal Go HTTP server. Routes return HTML, layouts are rendered on the server, and the browser only runs a small navigation layer. When a user clicks an internal link, Gofast asks the server for the next page fragment and swaps it into the current document.

Start with Gofast when you want the ergonomics of SPA navigation without moving your application model into JavaScript.

Use Gofast when you want:

- Go routing and rendering.
- HTML that stays close to the final output.
- SPA-style navigation without adopting a frontend framework.
- A small JavaScript surface that can be understood and replaced.

## Install

```bash
go get github.com/EzekielRiv1/gofast
```

Then create a Go server, register routes, and return rendered HTML from handlers.

```go
app := gofast.New()

app.Get("/", func(*gofast.Context) gofast.Page {
	return gofast.Page{
		Title: "Home",
		Body:  gofast.HTML("<h1>Hello from Go</h1>"),
	}
})
```
