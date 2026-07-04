package gofast

import (
	"context"
	"html/template"
	"net/http"
	"strings"
)

const navigateHeader = "X-Gofast-Navigate"

// Handler renders a page for the current request.
type Handler func(*Context) Page

// Context carries request-specific framework helpers.
type Context struct {
	Response http.ResponseWriter
	Request  *http.Request
}

// Page is the rendered result for a route.
type Page struct {
	Title  string
	Body   template.HTML
	Status int
}

// App is a small Go-first application router and renderer.
type App struct {
	routes   map[string]Handler
	layout   Layout
	notFound Handler
}

// New creates an application with the default document layout.
func New() *App {
	app := &App{
		routes: make(map[string]Handler),
		layout: DefaultLayout(),
	}
	app.notFound = func(*Context) Page {
		return Page{
			Title:  "Not found",
			Body:   "<h1>Not found</h1><p>The page could not be found.</p>",
			Status: http.StatusNotFound,
		}
	}
	return app
}

// WithLayout replaces the document layout used for full page requests.
func (a *App) WithLayout(layout Layout) *App {
	a.layout = layout
	return a
}

// NotFound replaces the default 404 page.
func (a *App) NotFound(handler Handler) {
	a.notFound = handler
}

// Route registers an exact path handler.
func (a *App) Route(path string, handler Handler) {
	if path == "" || !strings.HasPrefix(path, "/") {
		panic("gofast: route path must begin with /")
	}
	a.routes[path] = handler
}

// Get registers an exact GET route.
func (a *App) Get(path string, handler Handler) {
	a.Route(path, func(ctx *Context) Page {
		if ctx.Request.Method != http.MethodGet {
			ctx.Response.Header().Set("Allow", http.MethodGet)
			return Page{
				Title:  "Method not allowed",
				Body:   "<h1>Method not allowed</h1>",
				Status: http.StatusMethodNotAllowed,
			}
		}
		return handler(ctx)
	})
}

// ServeHTTP implements http.Handler.
func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler, ok := a.routes[r.URL.Path]
	if !ok {
		handler = a.notFound
	}
	a.render(w, r, handler)
}

// ListenAndServe starts an HTTP server for the application.
func (a *App) ListenAndServe(addr string) error {
	return http.ListenAndServe(addr, a)
}

// HTML marks a trusted HTML fragment for rendering.
func HTML(markup string) template.HTML {
	return template.HTML(markup)
}

// RequestContext returns the request context for convenience in route helpers.
func (c *Context) RequestContext() context.Context {
	return c.Request.Context()
}

func (a *App) render(w http.ResponseWriter, r *http.Request, handler Handler) {
	page := handler(&Context{Response: w, Request: r})
	if page.Status == 0 {
		page.Status = http.StatusOK
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if r.Header.Get(navigateHeader) == "true" {
		w.Header().Set("X-Gofast-Title", page.Title)
		w.WriteHeader(page.Status)
		_, _ = w.Write([]byte(page.Body))
		return
	}

	w.WriteHeader(page.Status)
	_ = a.layout.Execute(w, page)
}
